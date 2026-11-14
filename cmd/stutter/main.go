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
	"github.com/Leo-Zh9/stutter/internal/explore"
	"github.com/Leo-Zh9/stutter/internal/oracle"
	"github.com/Leo-Zh9/stutter/internal/pgenv"
	"github.com/Leo-Zh9/stutter/internal/replay"
	"github.com/Leo-Zh9/stutter/internal/target"
	"github.com/Leo-Zh9/stutter/internal/trace"
)

const usage = `stutter: find retry bugs by replaying traffic under faults

  stutter record  -schema FILE -target CMD -out TRACE [-webhooks /webhooks] -- WORKLOAD...
  stutter explore -schema FILE -target CMD -trace TRACE [-random N] [-lanes N] [-no-barrier]
  stutter bench   -minipay BIN -trace TRACE [-out results]

The service command must read DATABASE_URL and ADDR and answer GET /healthz.
In WORKLOAD, {proxy} is replaced with the recording proxy's base URL.
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

func targetSpec(cmd string, env []string) target.Spec {
	return target.Spec{Command: strings.Fields(cmd), Env: env, HealthPath: "/healthz"}
}

func record(ctx context.Context, root string, args []string) {
	fs := flag.NewFlagSet("record", flag.ExitOnError)
	schema := fs.String("schema", "", "schema SQL file")
	tcmd := fs.String("target", "", "service command")
	out := fs.String("out", "trace.jsonl", "trace file to write")
	hooks := fs.String("webhooks", "/webhooks", "path prefix of webhook deliveries")
	fs.Parse(args)
	workload := fs.Args()
	if *schema == "" || *tcmd == "" || len(workload) == 0 {
		log.Fatal("record needs -schema, -target and a workload command after --")
	}
	env, err := pgenv.FromRepo(root)
	must(err)
	sql, err := os.ReadFile(*schema)
	must(err)
	must(env.CreateTemplate(ctx, "tpl_record", string(sql)))
	db, err := env.Clone(ctx, "tpl_record")
	must(err)
	defer env.Drop(ctx, db)
	proc, err := target.Start(ctx, targetSpec(*tcmd, nil), env.DSN(db))
	must(err)
	defer proc.Stop()
	up, _ := url.Parse(proc.BaseURL)
	rec := trace.NewRecorder(up, *hooks)
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
	schema := fs.String("schema", "", "schema SQL file")
	tcmd := fs.String("target", "", "service command")
	tracePath := fs.String("trace", "", "trace file")
	random := fs.Int("random", 200, "random multi-fault schedules after the single-fault sweep")
	lanes := fs.Int("lanes", 8, "parallel replays")
	noBarrier := fs.Bool("no-barrier", false, "send concurrent duplicates without the lock barrier")
	envs := fs.String("env", "", "extra env for the service, comma separated (e.g. BUG=key-ignored)")
	seed := fs.Int64("seed", 1, "seed")
	ignore := fs.String("ignore", "", "nondeterministic columns to skip, comma separated table.column")
	fs.Parse(args)
	opts := oracle.Options{Ignore: map[string]bool{}}
	for _, c := range strings.Split(*ignore, ",") {
		if c != "" {
			opts.Ignore[c] = true
		}
	}
	env, err := pgenv.FromRepo(root)
	must(err)
	sql, err := os.ReadFile(*schema)
	must(err)
	must(env.CreateTemplate(ctx, "tpl_explore", string(sql)))
	tr, err := trace.Read(*tracePath)
	must(err)
	var extra []string
	if *envs != "" {
		extra = strings.Split(*envs, ",")
	}
	r := &explore.Runner{
		Cfg:   replay.Config{Env: env, Template: "tpl_explore", Barrier: !*noBarrier, Target: targetSpec(*tcmd, extra), Oracle: opts},
		Trace: tr, Lanes: *lanes, Seed: *seed,
	}
	must(r.CheckStable(ctx))
	scheds := explore.Schedules(tr, *seed, *random)
	fmt.Printf("%d requests, %d schedules\n", len(tr), len(scheds))
	idx, outs := r.FirstFailure(ctx, scheds)
	if idx < 0 {
		fmt.Printf("no fault changed the final state in %d runs\n", len(outs))
		return
	}
	fmt.Printf("run %d of %d changed the final state; shrinking\n", idx+1, len(scheds))
	min, st := r.Shrink(ctx, outs[idx].Schedule, 3)
	fmt.Printf("shrunk %d -> %d elements (%d candidates)\n\n", st.ElementsBefore, st.ElementsAfter, st.Candidates)
	for _, l := range bench.Describe(tr, min) {
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
	bin := fs.String("minipay", "bin/minipay", "minipay binary")
	tracePath := fs.String("trace", "traces/minipay.jsonl", "trace file")
	out := fs.String("out", "results", "output directory")
	lanes := fs.Int("lanes", 8, "parallel replays")
	random := fs.Int("random", 200, "random multi-fault schedules")
	trials := fs.Int("trials", 100, "trials per arm in the barrier experiment")
	seed := fs.Int64("seed", 1, "seed")
	fs.Parse(args)
	res, err := bench.Run(ctx, bench.Config{
		Root: root, Minipay: *bin, Trace: *tracePath, Lanes: *lanes, Random: *random,
		Trials: *trials, Seed: *seed, OutDir: *out,
		Progress: func(s string) { fmt.Fprintln(os.Stderr, "bench:", s) },
	})
	must(err)
	must(bench.Write(*out, res))
	fmt.Print(bench.Markdown(res))
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
