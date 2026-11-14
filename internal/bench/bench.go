// Package bench runs stutter's measurement suite against minipay and writes
// results/results.json and results/RESULTS.md. Every number in the README
// comes from these files.
package bench

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/Leo-Zh9/stutter/internal/explore"
	"github.com/Leo-Zh9/stutter/internal/oracle"
	"github.com/Leo-Zh9/stutter/internal/pgenv"
	"github.com/Leo-Zh9/stutter/internal/replay"
	"github.com/Leo-Zh9/stutter/internal/target"
	"github.com/Leo-Zh9/stutter/internal/trace"
)

type Variant struct {
	Bug      string `json:"bug"`
	Class    string `json:"class"` // correct | in-model | control
	Pattern  string `json:"pattern"`
	Expected string `json:"expected_fault"`
}

var Variants = []Variant{
	{"", "correct", "key row inserted first in the same transaction as the work; event ids behind a unique index", "none (false-alarm measurement)"},
	{"key-ignored", "in-model", "idempotency key accepted but never checked", "lost"},
	{"key-after-commit", "in-model", "key recorded in a second transaction after the charge commits", "concurrent"},
	{"webhook-no-dedupe", "in-model", "webhook handler keeps no event log", "lost (redelivery)"},
	{"webhook-stale-overwrite", "in-model", "status webhook applied without comparing event time", "reorder"},
	{"webhook-check-then-insert", "in-model", "event dedupe is SELECT then INSERT with no unique index", "concurrent (redelivery)"},
	{"key-unscoped", "control", "key not bound to the request body; wrong with or without faults", "none: outside a differential oracle"},
}

type Config struct {
	Root     string // repo root
	Minipay  string // path to built minipay binary
	Trace    string
	Lanes    int
	Random   int // random multi-fault schedules after the single-fault sweep
	Trials   int // per arm for the barrier experiment
	Seed     int64
	OutDir   string
	Progress func(string)
}

type Rate struct {
	Detected int `json:"detected"`
	NFired   int `json:"n_fired"`
	NNotFire int `json:"n_not_fired"`
}

type VariantResult struct {
	Variant
	Schedules    int                 `json:"schedules_available"`
	Ran          int                 `json:"schedules_run"`
	Detected     bool                `json:"detected"`
	DetectedAt   int                 `json:"detected_at_run,omitempty"` // 1-based
	NFired       int                 `json:"n_fired"`
	NNotFired    int                 `json:"n_not_fired"`
	FirstFailing *replay.Schedule    `json:"-"`
	FirstFaults  []string            `json:"first_failing_faults,omitempty"`
	Diff         []oracle.DiffLine   `json:"diff,omitempty"`
	Shrunk       []string            `json:"shrunk,omitempty"`
	Shrink       *explore.ShrinkStats `json:"shrink,omitempty"`
	OneMinimal   *bool               `json:"one_minimal_verified,omitempty"`
	ShrinkSec    float64             `json:"shrink_seconds,omitempty"`
}

type Trigger struct {
	Bug        string `json:"bug"`
	Fault      string `json:"fault"`
	BarrierOn  Rate   `json:"barrier_on"`
	BarrierOff Rate   `json:"barrier_off"`
	Trials     int    `json:"trials_per_arm"`
}

type Results struct {
	When        string          `json:"when"`
	Machine     string          `json:"machine"`
	Postgres    string          `json:"postgres"`
	Lanes       int             `json:"lanes"`
	TraceEntries int            `json:"trace_entries"`
	Retryable   int             `json:"retryable_entries"`
	Webhooks    int             `json:"webhook_entries"`
	SingleFaults int            `json:"single_fault_schedules"`
	RandomFaults int            `json:"random_multi_fault_schedules"`
	Determinism struct {
		Replays  int `json:"fault_free_replays"`
		Distinct int `json:"distinct_digests"`
	} `json:"determinism"`
	FalseAlarms struct {
		Runs    int `json:"schedules_run"`
		NFired  int `json:"n_fired"`
		NotFire int `json:"n_not_fired"`
		Flagged int `json:"flagged"`
	} `json:"false_alarms_correct_variant"`
	RunsPerSec  float64         `json:"replays_per_second"`
	MedianRunMs float64         `json:"median_replay_ms"`
	Variants    []VariantResult `json:"variants"`
	Triggers    []Trigger       `json:"barrier_trigger_rate"`
	TotalSec    float64         `json:"total_seconds"`
}

func schemaFor(bin, bug string) (string, error) {
	cmd := exec.Command(bin, "-schema")
	cmd.Env = append(os.Environ(), "BUG="+bug)
	out, err := cmd.Output()
	return string(out), err
}

func tplName(bug string) string {
	if bug == "" {
		return "tpl_correct"
	}
	return "tpl_" + strings.ReplaceAll(bug, "-", "_")
}

func Run(ctx context.Context, cfg Config) (*Results, error) {
	t0 := time.Now()
	say := cfg.Progress
	if say == nil {
		say = func(string) {}
	}
	env, err := pgenv.FromRepo(cfg.Root)
	if err != nil {
		return nil, err
	}
	defer env.Close()
	tr, err := trace.Read(cfg.Trace)
	if err != nil {
		return nil, err
	}
	res := &Results{When: time.Now().UTC().Format(time.RFC3339), Lanes: cfg.Lanes, TraceEntries: len(tr)}
	res.Machine = fmt.Sprintf("%s/%s, %d CPUs", runtime.GOOS, runtime.GOARCH, runtime.NumCPU())
	env.Admin().QueryRow(ctx, "SHOW server_version").Scan(&res.Postgres)
	for _, e := range tr {
		if e.Retryable() {
			res.Retryable++
		}
		if e.Kind == trace.Webhook {
			res.Webhooks++
		}
	}
	scheds := explore.Schedules(tr, cfg.Seed, cfg.Random)
	res.SingleFaults = len(replay.Faults(tr))
	res.RandomFaults = len(scheds) - res.SingleFaults

	runner := func(bug string, barrier bool) *explore.Runner {
		return &explore.Runner{
			Cfg: replay.Config{
				Env: env, Template: tplName(bug), Barrier: barrier,
				// A hash of a body that embeds a generated customer id: derived
				// nondeterminism the catalog rule cannot see. CheckStable names it.
				Oracle: oracle.Options{Ignore: map[string]bool{"idempotency_keys.request_hash": true}},
				Target: target.Spec{Command: []string{cfg.Minipay}, Env: []string{"BUG=" + bug}, HealthPath: "/healthz"},
			},
			Trace: tr, Lanes: cfg.Lanes, Seed: cfg.Seed,
		}
	}
	for _, v := range Variants {
		schema, err := schemaFor(cfg.Minipay, v.Bug)
		if err != nil {
			return nil, err
		}
		if err := env.CreateTemplate(ctx, tplName(v.Bug), schema); err != nil {
			return nil, err
		}
	}

	// Determinism: fault-free replays of the full trace on the correct variant.
	say("determinism: 20 fault-free replays")
	rc := runner("", true)
	digests := map[string]bool{}
	outs := rc.Parallel(20, func(i int) explore.Outcome {
		r := replay.Run(ctx, rc.Cfg, tr, replay.Full(tr), cfg.Seed+int64(i))
		if r.Err != nil {
			return explore.Outcome{Err: r.Err.Error()}
		}
		return explore.Outcome{Err: "", Diff: []oracle.DiffLine{{Row: r.Canon.Digest()}}}
	})
	for _, o := range outs {
		if o.Err != "" {
			return nil, fmt.Errorf("fault-free replay: %s", o.Err)
		}
		digests[o.Diff[0].Row] = true
	}
	res.Determinism.Replays, res.Determinism.Distinct = 20, len(digests)

	// False alarms: every schedule against the correct variant.
	say(fmt.Sprintf("false alarms: %d schedules on the correct variant", len(scheds)))
	if err := rc.CheckStable(ctx); err != nil {
		return nil, err
	}
	var durs []time.Duration
	tSweep := time.Now()
	before := rc.Runs()
	outs = rc.Parallel(len(scheds), func(i int) explore.Outcome {
		s := time.Now()
		o := rc.Check(ctx, scheds[i], cfg.Seed+int64(i))
		_ = s
		return o
	})
	sweep := time.Since(tSweep)
	runs := rc.Runs() - before
	res.RunsPerSec = float64(runs) / sweep.Seconds()
	res.FalseAlarms.Runs = len(outs)
	for _, o := range outs {
		if o.Err != "" {
			return nil, fmt.Errorf("correct variant: %s", o.Err)
		}
		if o.Fired {
			res.FalseAlarms.NFired++
		} else {
			res.FalseAlarms.NotFire++
		}
		if o.Failed {
			res.FalseAlarms.Flagged++
		}
	}
	// Median single-replay latency, measured on its own lane-free loop.
	for i := 0; i < 15; i++ {
		r := replay.Run(ctx, rc.Cfg, tr, replay.Full(tr), cfg.Seed)
		durs = append(durs, r.Took)
	}
	sort.Slice(durs, func(i, j int) bool { return durs[i] < durs[j] })
	res.MedianRunMs = float64(durs[len(durs)/2].Microseconds()) / 1000
	res.Variants = append(res.Variants, VariantResult{
		Variant: Variants[0], Schedules: len(scheds), Ran: len(outs),
		NFired: res.FalseAlarms.NFired, NNotFired: res.FalseAlarms.NotFire, Detected: res.FalseAlarms.Flagged > 0,
	})

	for _, v := range Variants[1:] {
		say("variant " + v.Bug)
		r := runner(v.Bug, true)
		if err := r.CheckStable(ctx); err != nil {
			return nil, fmt.Errorf("%s: %w", v.Bug, err)
		}
		idx, outs := r.FirstFailure(ctx, scheds)
		vr := VariantResult{Variant: v, Schedules: len(scheds), Ran: len(outs)}
		for _, o := range outs {
			if o.Err != "" {
				return nil, fmt.Errorf("%s: %s", v.Bug, o.Err)
			}
			if o.Fired {
				vr.NFired++
			} else {
				vr.NNotFired++
			}
		}
		if idx >= 0 {
			vr.Detected, vr.DetectedAt = true, idx+1
			first := outs[idx]
			vr.FirstFailing = &first.Schedule
			for _, f := range first.Schedule.Faults {
				vr.FirstFaults = append(vr.FirstFaults, f.String())
			}
			ts := time.Now()
			min, st := r.Shrink(ctx, first.Schedule, 3)
			vr.ShrinkSec = time.Since(ts).Seconds()
			vr.Shrink = &st
			vr.Shrunk = Describe(tr, min)
			ok := r.OneMinimal(ctx, min, 3)
			vr.OneMinimal = &ok
			fin := r.Check(ctx, min, cfg.Seed)
			vr.Diff = fin.Diff
			if min.HasConcurrent() {
				for _, f := range min.Faults {
					if f.Type == replay.Concurrent {
						say("barrier trigger rate " + v.Bug)
						res.Triggers = append(res.Triggers, triggerRate(ctx, cfg, runner(v.Bug, true), runner(v.Bug, false), tr, f))
					}
				}
			}
		}
		res.Variants = append(res.Variants, vr)
	}
	res.TotalSec = time.Since(t0).Seconds()
	return res, nil
}

// triggerRate replays the full trace with one Concurrent fault, `Trials`
// times with the barrier and `Trials` times without it (both attempts sent
// back to back from two goroutines), on a fresh clone each time.
func triggerRate(ctx context.Context, cfg Config, on, off *explore.Runner, tr []trace.Entry, f replay.Fault) Trigger {
	t := Trigger{Bug: on.Cfg.Target.Env[0][4:], Fault: f.String(), Trials: cfg.Trials}
	s := replay.Full(tr, f)
	arm := func(r *explore.Runner) Rate {
		var rate Rate
		outs := r.Parallel(cfg.Trials, func(i int) explore.Outcome { return r.Check(ctx, s, cfg.Seed+int64(i)) })
		for _, o := range outs {
			if !o.Fired {
				rate.NNotFire++
				continue
			}
			rate.NFired++
			if o.Failed {
				rate.Detected++
			}
		}
		return rate
	}
	t.BarrierOn = arm(on)
	t.BarrierOff = arm(off)
	return t
}

// Describe renders a schedule as one line per request plus its faults.
func Describe(tr []trace.Entry, s replay.Schedule) []string {
	by := map[int]trace.Entry{}
	for _, e := range tr {
		by[e.Seq] = e
	}
	var out []string
	for _, q := range s.Seqs {
		out = append(out, describeEntry(by[q]))
	}
	for _, f := range s.Faults {
		out = append(out, "fault "+f.String())
	}
	return out
}

func describeEntry(e trace.Entry) string {
	if e.Kind == trace.Webhook {
		var ev struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		}
		json.Unmarshal(e.Body, &ev)
		return fmt.Sprintf("#%d webhook %s %s", e.Seq, ev.Type, ev.ID)
	}
	return fmt.Sprintf("#%d %s %s key=%s", e.Seq, e.Method, e.Path, e.Key())
}

func Write(dir string, r *Results) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	b, _ := json.MarshalIndent(r, "", "  ")
	if err := os.WriteFile(filepath.Join(dir, "results.json"), append(b, '\n'), 0o644); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, "RESULTS.md"), []byte(Markdown(r)), 0o644)
}

func Markdown(r *Results) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# stutter results\n\nGenerated by `make bench` on %s (%s, PostgreSQL %s, %d lanes). Target: minipay, a simulated payments API with seeded bugs. Trace: %d recorded requests (%d retryable, %d webhooks), synthetic workload.\n\n",
		r.When, r.Machine, r.Postgres, r.Lanes, r.TraceEntries, r.Retryable, r.Webhooks)
	fmt.Fprintf(&b, "Schedules per variant: %d single-fault (every lost, concurrent and reorder fault the trace admits) then %d random 2 to 4 fault schedules.\n\n", r.SingleFaults, r.RandomFaults)
	b.WriteString("## Detection\n\n| Variant | Class | Expected fault | Detected | First failing run | Fired / not fired | Shrunk to | 1-minimal |\n|---|---|---|---|---|---|---|---|\n")
	for _, v := range r.Variants {
		name := v.Bug
		if name == "" {
			name = "correct"
		}
		det := "no"
		at := "-"
		if v.Detected {
			det = "yes"
			at = fmt.Sprintf("%d of %d", v.DetectedAt, v.Schedules)
		}
		sh, om := "-", "-"
		if v.Shrink != nil {
			sh = fmt.Sprintf("%d -> %d elements", v.Shrink.ElementsBefore, v.Shrink.ElementsAfter)
		}
		if v.OneMinimal != nil {
			om = fmt.Sprint(*v.OneMinimal)
		}
		fmt.Fprintf(&b, "| %s | %s | %s | %s | %s | %d / %d | %s | %s |\n", name, v.Class, v.Expected, det, at, v.NFired, v.NNotFired, sh, om)
	}
	fa := r.FalseAlarms
	fmt.Fprintf(&b, "\nFalse alarms on the correct variant: %d of %d fired runs (%d schedules did not fire and are not counted as passes).\n\n", fa.Flagged, fa.NFired, fa.NotFire)
	b.WriteString("## Lock barrier: does the concurrent duplicate actually hit the race?\n\n| Bug | Fault | Barrier on | Barrier off |\n|---|---|---|---|\n")
	for _, t := range r.Triggers {
		fmt.Fprintf(&b, "| %s | %s | %d / %d fired (%d not fired) | %d / %d |\n", t.Bug, t.Fault,
			t.BarrierOn.Detected, t.BarrierOn.NFired, t.BarrierOn.NNotFire, t.BarrierOff.Detected, t.BarrierOff.NFired)
	}
	b.WriteString("\n## Minimal reproductions\n\n")
	for _, v := range r.Variants {
		if len(v.Shrunk) == 0 {
			continue
		}
		fmt.Fprintf(&b, "### %s\n\n```\n%s\n```\n\nState diff (+ only in the faulted run, - only in the fault-free run):\n\n```\n", v.Bug, strings.Join(v.Shrunk, "\n"))
		for _, d := range v.Diff {
			fmt.Fprintf(&b, "%s %s x%d %s\n", d.Sign, d.Table, d.Count, d.Row)
		}
		b.WriteString("```\n\n")
	}
	fmt.Fprintf(&b, "## Speed\n\n%.1f replays/s across %d lanes during the correct-variant sweep (each replay clones a template database, starts the service, replays the trace, snapshots and drops the database). Median single replay: %.0f ms.\n\nDeterminism: %d fault-free replays produced %d distinct canonical state digest(s).\n\nTotal bench time: %.0f s.\n",
		r.RunsPerSec, r.Lanes, r.MedianRunMs, r.Determinism.Replays, r.Determinism.Distinct, r.TotalSec)
	return b.String()
}
