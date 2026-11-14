// Command stutter records API traffic and replays it under retry faults,
// flagging any fault that changes the final database state.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/Leo-Zh9/stutter/internal/bench"
	"github.com/Leo-Zh9/stutter/internal/config"
	"github.com/Leo-Zh9/stutter/internal/explore"
	"github.com/Leo-Zh9/stutter/internal/pgenv"
	"github.com/Leo-Zh9/stutter/internal/target"
	"github.com/Leo-Zh9/stutter/internal/trace"
)

const usage = `stutter: find retry bugs by replaying traffic under faults

  stutter record  -config stutter.json [-out TRACE] -- WORKLOAD...
  stutter explore -config stutter.json [-trace TRACE] [-env K=V,...] [-random N] [-lanes N] [-no-barrier]
  stutter bench   [-config targets/minipay/stutter.json] [-out results]

In WORKLOAD, {proxy} is replaced with the recording proxy's base URL.
Exit codes: 0 no failure found, 1 failure found, 2 tool error.
`

func main() {
	log.SetFlags(0)
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(2)
	}
	root, _ := os.Getwd()
	ctx := context.Background()
	switch os.Args[1] {
	case "record":
		record(ctx, root, os.Args[2:])
	case "explore":
		exploreCmd(ctx, root, os.Args[2:])
	case "bench":
		benchCmd(ctx, root, os.Args[2:])
	default:
		fmt.Fprint(os.Stderr, usage)
		os.Exit(2)
	}
}

func splitEnv(s string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, ",")
}

func record(ctx context.Context, root string, args []string) {
	fs := flag.NewFlagSet("record", flag.ExitOnError)
	cfgPath := fs.String("config", "stutter.json", "target config")
	out := fs.String("out", "", "trace file to write (default: the config's trace)")
	envs := fs.String("env", "", "extra env for the service, comma separated")
	fs.Parse(args)
	workload := fs.Args()
	if len(workload) == 0 {
		fail("record needs a workload command after --")
	}
	cfg, err := config.Load(*cfgPath)
	must(err)
	if *out == "" {
		*out = cfg.Path(cfg.Trace)
	}
	env, err := pgenv.FromRepo(root)
	must(err)
	extra := splitEnv(*envs)
	must(cfg.PrepareTemplate(ctx, env, "tpl_record", extra))
	db, err := env.Clone(ctx, "tpl_record")
	must(err)
	defer env.Drop(ctx, db)
	proc, err := target.Start(ctx, cfg.Spec(extra), target.DB{DSN: env.DSN(db), URL: env.URL(db)})
	must(err)
	defer proc.Stop()
	up, _ := url.Parse(proc.BaseURL)
	rec := trace.NewRecorder(up, cfg.Webhooks.Prefix, cfg.KeyHeader)
	l, err := net.Listen("tcp", "127.0.0.1:0")
	must(err)
	go http.Serve(l, rec)
	proxy := "http://" + l.Addr().String()
	for i, a := range workload {
		workload[i] = strings.ReplaceAll(a, "{proxy}", proxy)
	}
	cmd := exec.Command(workload[0], workload[1:]...)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	must(cmd.Run())
	entries := rec.Entries()
	must(trace.Write(*out, entries))
	fmt.Printf("recorded %d requests to %s\n", len(entries), *out)
}

func exploreCmd(ctx context.Context, root string, args []string) {
	fs := flag.NewFlagSet("explore", flag.ExitOnError)
	cfgPath := fs.String("config", "stutter.json", "target config")
	tracePath := fs.String("trace", "", "trace file (default: the config's trace)")
	random := fs.Int("random", 200, "random multi-fault schedules after the single-fault sweep")
	lanes := fs.Int("lanes", 8, "parallel replays")
	noBarrier := fs.Bool("no-barrier", false, "send concurrent duplicates without the lock barrier")
	envs := fs.String("env", "", "extra env for the service, comma separated (e.g. BUG=key-ignored)")
	seed := fs.Int64("seed", 1, "seed")
	fs.Parse(args)
	cfg, err := config.Load(*cfgPath)
	must(err)
	if *tracePath == "" {
		*tracePath = cfg.Path(cfg.Trace)
	}
	env, err := pgenv.FromRepo(root)
	must(err)
	extra := splitEnv(*envs)
	must(cfg.PrepareTemplate(ctx, env, "tpl_explore", extra))
	tr, err := trace.Read(*tracePath)
	must(err)
	rc := cfg.Replay(env, "tpl_explore", extra)
	rc.Barrier = !*noBarrier
	r := &explore.Runner{Cfg: rc, Trace: tr, Lanes: *lanes, Seed: *seed}
	must(r.CheckStable(ctx))
	scheds := explore.Schedules(tr, cfg.KeyHeader, *seed, *random)
	fmt.Printf("%d requests, %d schedules\n", len(tr), len(scheds))
	idx, outs := r.FirstFailure(ctx, scheds)
	for _, o := range outs {
		if o.Err != "" {
			fail(o.Err)
		}
	}
	if idx < 0 {
		fmt.Printf("no fault changed the final state in %d runs\n", len(outs))
		return
	}
	fmt.Printf("run %d of %d changed the final state; shrinking\n", idx+1, len(scheds))
	min, st := r.Shrink(ctx, outs[idx].Schedule, 3)
	fmt.Printf("shrunk %d -> %d elements (%d candidates)\n\n", st.ElementsBefore, st.ElementsAfter, st.Candidates)
	for _, l := range bench.Describe(tr, min, cfg.KeyHeader) {
		fmt.Println("  " + l)
	}
	fmt.Println()
	for _, d := range r.Check(ctx, min, *seed).Diff {
		fmt.Printf("  %s %s x%d %s\n", d.Sign, d.Table, d.Count, d.Row)
	}
	os.Exit(1)
}

func benchCmd(ctx context.Context, root string, args []string) {
	fs := flag.NewFlagSet("bench", flag.ExitOnError)
	cfgPath := fs.String("config", "targets/minipay/stutter.json", "minipay config")
	out := fs.String("out", "results", "output directory")
	lanes := fs.Int("lanes", 8, "parallel replays")
	random := fs.Int("random", 200, "random multi-fault schedules")
	trials := fs.Int("trials", 100, "trials per arm in the barrier experiment")
	seed := fs.Int64("seed", 1, "seed")
	fs.Parse(args)
	cfg, err := config.Load(*cfgPath)
	must(err)
	res, err := bench.Run(ctx, bench.Config{
		Root: root, Target: cfg, Lanes: *lanes, Random: *random,
		Trials: *trials, Seed: *seed, OutDir: *out,
		Progress: func(s string) { fmt.Fprintln(os.Stderr, "bench:", s) },
	})
	must(err)
	must(bench.Write(*out, res))
	fmt.Print(bench.Markdown(res))
}

func fail(msg string) {
	fmt.Fprintln(os.Stderr, "stutter:", msg)
	os.Exit(2)
}

func must(err error) {
	if err != nil {
		fail(err.Error())
	}
}
