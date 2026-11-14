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
	"path/filepath"
	"strings"

	"github.com/Leo-Zh9/stutter/internal/bench"
	"github.com/Leo-Zh9/stutter/internal/config"
	"github.com/Leo-Zh9/stutter/internal/explore"
	"github.com/Leo-Zh9/stutter/internal/hunt"
	"github.com/Leo-Zh9/stutter/internal/pgenv"
	"github.com/Leo-Zh9/stutter/internal/replay"
	"github.com/Leo-Zh9/stutter/internal/target"
	"github.com/Leo-Zh9/stutter/internal/trace"
)

const usage = `stutter: find retry bugs by replaying traffic under faults

  stutter record  -config stutter.json [-out TRACE] -- WORKLOAD...
  stutter explore -config stutter.json [-trace TRACE] [-env K=V,...] [-random N] [-lanes N] [-no-barrier]
  stutter hunt    -config stutter.json [-out DIR]   every single fault, shrink, report, curl repros
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
	case "hunt":
		huntCmd(ctx, root, os.Args[2:])
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
	tpl := cfg.TemplateName("record")
	must(cfg.PrepareTemplate(ctx, env, tpl, extra))
	db, err := env.Clone(ctx, tpl)
	must(err)
	defer env.Drop(ctx, db)
	proc, err := target.Start(ctx, cfg.Spec(extra), target.DB{DSN: env.DSN(db), URL: env.URL(db), SocketURL: env.SocketURL(db)})
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
	tpl := cfg.TemplateName("explore")
	must(cfg.PrepareTemplate(ctx, env, tpl, extra))
	tr, err := trace.Read(*tracePath)
	must(err)
	rc := cfg.Replay(env, tpl, extra)
	rc.Barrier = !*noBarrier
	r := &explore.Runner{Cfg: rc, Trace: tr, Lanes: *lanes, Seed: *seed}
	must(r.CheckStable(ctx))
	scheds := explore.Schedules(tr, cfg.Space(), *seed, *random)
	fmt.Printf("%d requests, %d schedules\n", len(tr), len(scheds))
	idx, outs := r.FirstFailure(ctx, scheds)
	for _, o := range outs {
		if o.Err != "" {
			fail(o.Err)
		}
	}
	fired := 0
	for _, o := range outs {
		if o.Fired {
			fired++
		}
	}
	if idx < 0 {
		fmt.Printf("no fault changed the final state in %d runs (%d fired, %d did not fire)\n", len(outs), fired, len(outs)-fired)
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

func huntCmd(ctx context.Context, root string, args []string) {
	fs := flag.NewFlagSet("hunt", flag.ExitOnError)
	cfgPath := fs.String("config", "stutter.json", "target config")
	out := fs.String("out", "", "output directory (default: hunt/ next to the config)")
	lanes := fs.Int("lanes", 6, "parallel replays")
	seed := fs.Int64("seed", 1, "seed")
	noCrash := fs.Bool("no-crash", false, "skip crash-point faults")
	fs.Parse(args)
	cfg, err := config.Load(*cfgPath)
	must(err)
	if *out == "" {
		*out = cfg.Path("hunt")
	}
	env, err := pgenv.FromRepo(root)
	must(err)
	tpl := cfg.TemplateName("hunt")
	must(cfg.PrepareTemplate(ctx, env, tpl, nil))
	tr, err := trace.Read(cfg.Path(cfg.Trace))
	must(err)
	r := &explore.Runner{Cfg: cfg.Replay(env, tpl, nil), Trace: tr, Lanes: *lanes, Seed: *seed}
	say := func(s string) { fmt.Fprintln(os.Stderr, "hunt:", s) }
	res, err := hunt.Run(ctx, cfg, r, *out, !*noCrash, say)
	must(err)
	for i := range res.Findings {
		if hasCrash(res.Findings[i].Minimal) {
			res.Findings[i].CurlRepro = -1 // a crash point cannot be reproduced with curl
			res.Findings[i].Dir, _ = filepath.Rel(*out, res.Findings[i].Dir)
			continue
		}
		say(fmt.Sprintf("finding %d: running repro.sh 10 times", i+1))
		n, err := hunt.VerifyRepro(ctx, cfg, env, tpl, res.Findings[i].Dir, 10)
		must(err)
		res.Findings[i].CurlRepro = n
		res.Findings[i].Dir, _ = filepath.Rel(*out, res.Findings[i].Dir)
	}
	must(hunt.Write(*out, res))
	fmt.Print(hunt.Markdown(res))
	if len(res.Findings) > 0 {
		os.Exit(1)
	}
}

func hasCrash(s replay.Schedule) bool {
	for _, f := range s.Faults {
		if f.Type == replay.Crash {
			return true
		}
	}
	return false
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
