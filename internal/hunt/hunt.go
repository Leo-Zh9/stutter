// Package hunt runs every single fault a trace admits against a target,
// shrinks each failure, groups identical minimal reproductions, and writes a
// report plus a curl script per finding.
package hunt

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/Leo-Zh9/stutter/internal/agentstub"
	"github.com/Leo-Zh9/stutter/internal/config"
	"github.com/Leo-Zh9/stutter/internal/explore"
	"github.com/Leo-Zh9/stutter/internal/oracle"
	"github.com/Leo-Zh9/stutter/internal/pgenv"
	"github.com/Leo-Zh9/stutter/internal/replay"
	"github.com/Leo-Zh9/stutter/internal/target"
	"github.com/Leo-Zh9/stutter/internal/trace"
)

type Finding struct {
	ID         int               `json:"id"`
	Minimal    replay.Schedule   `json:"minimal"`
	Steps      []string          `json:"steps"`
	Delivery   []string          `json:"delivery_order"`
	Diff       []oracle.DiffLine `json:"diff"`
	OneMinimal bool              `json:"one_minimal_verified"`
	// Unjudged counts single removals that could not fail (their fault did
	// not fire or their run left the model script): 1-minimal holds over
	// the others.
	Unjudged  int      `json:"one_minimal_unjudged_removals"`
	Replayed  int      `json:"reproduced_in_replays"`
	CurlRepro int      `json:"reproduced_by_repro_sh"`
	ReplayOf  int      `json:"replays"`
	FoundBy   []string `json:"found_by_single_faults"`
	Dir       string   `json:"dir"`
}

type Result struct {
	Target    string              `json:"target"`
	When      string              `json:"when"`
	Requests  int                 `json:"trace_requests"`
	Preflight string              `json:"preflight"`
	Schedules int                 `json:"single_fault_schedules"`
	ByType    map[string]*Tally   `json:"by_fault_type"`
	NoOpPairs []string            `json:"no_op_reorder_pairs"`
	Crash     *explore.CrashSpace `json:"crash,omitempty"`
	Inference *oracle.Inference   `json:"id_inference,omitempty"`
	Findings  []Finding           `json:"findings"`
	Seconds   float64             `json:"seconds"`
}

type Tally struct {
	Run      int `json:"run"`
	Fired    int `json:"n_fired"`
	NotFired int `json:"n_not_fired"`
	Failed   int `json:"failed"`
	NoOp     int `json:"no_op_passes,omitempty"`
	// Diverged runs left the model script; they are neither passes nor
	// failures. Nondetermined crash faults gave different outcomes in their
	// two runs (async checkpoint writes can race the kill); they are not
	// findings.
	Diverged      int `json:"diverged,omitempty"`
	Nondetermined int `json:"nondetermined,omitempty"`
	// Why groups the faults that did not fire by reason (held commits only).
	Why map[string][]string `json:"not_fired_reasons,omitempty"`
}

const replays = 10

// Run hunts one target. The runner must already have a template.
func Run(ctx context.Context, cfg *config.Config, r *explore.Runner, outDir string, crash bool, say func(string)) (*Result, error) {
	t0 := time.Now()
	tr := r.Trace
	res := &Result{Target: cfg.Name, When: time.Now().UTC().Format(time.RFC3339), Requests: len(tr), ByType: map[string]*Tally{}}

	// Pre-flight: the recorded, in-order run must have accepted every
	// webhook, or a malformed event could pose as a finding.
	hooks := 0
	for _, e := range tr {
		if e.Kind != trace.Webhook {
			continue
		}
		hooks++
		if e.Status < 200 || e.Status >= 300 {
			return nil, fmt.Errorf("preflight: recorded delivery #%d got HTTP %d; fix the trace first", e.Seq, e.Status)
		}
	}
	res.Preflight = "every recorded webhook delivery got 2xx"
	if hooks == 0 {
		res.Preflight = "no webhook deliveries in the trace"
	}
	say("stability check")
	if err := r.CheckStable(ctx); err != nil {
		return nil, err
	}
	full, err := r.Baseline(ctx, replay.Full(tr).Seqs)
	if err != nil {
		return nil, err
	}

	res.Inference = r.Inference()
	// One observation pass serves crash points, held commits and resume
	// rules: it counts commit points per request and records whether an
	// interrupt was pending around each resume.
	sp := cfg.Space()
	var cs explore.CrashSpace
	if crash || cfg.Holds || hasResume(cfg.RetryRules) {
		say("observing commit points and pending interrupts")
		var err error
		if cs, err = r.CrashFaults(ctx, cfg.KeyHeader); err != nil {
			return nil, err
		}
		sp.Pending = cs.Pending
	}
	singles := replay.Faults(tr, sp)
	if crash {
		res.Crash = &cs
		singles = append(singles, cs.Faults...)
	}
	if cfg.Holds {
		singles = append(singles, cs.Holds...)
	} else {
		cs.HoldPoints, cs.HoldNondetermined = 0, nil // not run, so not reported
	}
	res.Schedules = len(singles)
	say(fmt.Sprintf("%d single-fault schedules", len(singles)))
	nondet := make([]bool, len(singles))
	outs := r.Parallel(len(singles), func(i int) explore.Outcome {
		o := r.Check(ctx, replay.Full(tr, singles[i]), r.Seed+int64(i))
		twice := singles[i].Type == replay.Crash || singles[i].Type == replay.ConcurrentAt
		if !twice || o.Err != "" || o.Diverged || !o.Fired {
			return o
		}
		// Run twice: with async checkpoint writes the next node can race
		// the kill or the hold, so a crash or held-commit outcome counts
		// only if a second run agrees with the first.
		o2 := r.Check(ctx, replay.Full(tr, singles[i]), r.Seed+int64(i)+7919)
		if o2.Err != "" || o2.Diverged {
			return o2
		}
		if o2.Diverged || !o2.Fired || o2.Failed != o.Failed || fmt.Sprint(o2.Diff) != fmt.Sprint(o.Diff) {
			nondet[i] = true
		}
		return o
	})
	var failing []explore.Outcome
	for i, o := range outs {
		if o.Err != "" {
			return nil, fmt.Errorf("%v: %s", singles[i], o.Err)
		}
		ft := string(singles[i].Type)
		tl := res.ByType[ft]
		if tl == nil {
			tl = &Tally{}
			res.ByType[ft] = tl
		}
		tl.Run++
		if o.Diverged {
			tl.Diverged++
			continue
		}
		if !o.Fired {
			tl.NotFired++
			for _, w := range o.Why {
				if tl.Why == nil {
					tl.Why = map[string][]string{}
				}
				tl.Why[w] = append(tl.Why[w], singles[i].String())
			}
			continue
		}
		if nondet[i] {
			tl.Nondetermined++
			continue
		}
		tl.Fired++
		if o.Failed {
			tl.Failed++
			failing = append(failing, o)
		}
	}

	// A passing reorder pair is a no-op pair when dropping both events leaves
	// the end state unchanged: the handler ignored them, so the pass says
	// nothing about ordering. Failing pairs are never reclassified.
	failed := map[string]bool{}
	for _, o := range failing {
		failed[o.Schedule.Faults[0].String()] = true
	}
	for _, f := range singles {
		if f.Type != replay.Reorder || failed[f.String()] {
			continue
		}
		var seqs []int
		for _, e := range tr {
			if e.Seq != f.Seq && e.Seq != f.Seq2 {
				seqs = append(seqs, e.Seq)
			}
		}
		b, err := r.Baseline(ctx, seqs)
		if errors.Is(err, explore.ErrDiverged) {
			continue // cannot tell whether the pair is a no-op
		}
		if err != nil {
			return nil, err
		}
		if len(oracle.Diff(full, b)) == 0 {
			res.NoOpPairs = append(res.NoOpPairs, f.String())
			res.ByType[string(replay.Reorder)].NoOp++
		}
	}

	// Shrink every failure and group identical minimal schedules.
	say(fmt.Sprintf("%d failing schedules; shrinking each", len(failing)))
	type group struct {
		min     replay.Schedule
		foundBy []string
	}
	groups := map[string]*group{}
	var order []string
	for _, o := range failing {
		min, _ := r.Shrink(ctx, o.Schedule, 3)
		k := fmt.Sprint(min)
		if groups[k] == nil {
			groups[k] = &group{min: min}
			order = append(order, k)
		}
		groups[k].foundBy = append(groups[k].foundBy, o.Schedule.Faults[0].String())
	}
	sort.SliceStable(order, func(i, j int) bool {
		a, b := groups[order[i]].min, groups[order[j]].min
		return len(a.Seqs)+len(a.Faults) < len(b.Seqs)+len(b.Faults)
	})
	for i, k := range order {
		g := groups[k]
		say(fmt.Sprintf("finding %d: verifying", i+1))
		f := Finding{ID: i + 1, Minimal: g.min, FoundBy: g.foundBy, ReplayOf: replays}
		f.Steps = describe(tr, g.min, cfg.KeyHeader)
		f.Delivery = deliveryOrder(tr, g.min)
		f.OneMinimal, f.Unjudged = r.MinimalCheck(ctx, g.min, 3)
		f.Diff = r.Check(ctx, g.min, r.Seed).Diff
		for n := 0; n < replays; n++ {
			if o := r.Check(ctx, g.min, r.Seed+100+int64(n)); o.Failed {
				f.Replayed++
			}
		}
		f.Dir = filepath.Join(outDir, fmt.Sprintf("finding-%d", f.ID))
		if err := WriteRepro(f.Dir, cfg, tr, g.min); err != nil {
			return nil, err
		}
		res.Findings = append(res.Findings, f)
	}
	res.Seconds = time.Since(t0).Seconds()
	return res, nil
}

func hasResume(rules []trace.Rule) bool {
	for _, r := range rules {
		if r.Resume != nil {
			return true
		}
	}
	return false
}

func describe(tr []trace.Entry, s replay.Schedule, keyHeader string) []string {
	by := map[int]trace.Entry{}
	for _, e := range tr {
		by[e.Seq] = e
	}
	var out []string
	for _, q := range s.Seqs {
		out = append(out, Describe(by[q], keyHeader))
	}
	for _, f := range s.Faults {
		out = append(out, "fault "+f.String())
	}
	return out
}

// Describe renders one trace entry; Stripe events show type, customer, subscription and status.
func Describe(e trace.Entry, keyHeader string) string {
	if e.Kind == trace.Webhook {
		var ev struct {
			ID   string `json:"id"`
			Type string `json:"type"`
			Data struct {
				Object map[string]any `json:"object"`
			} `json:"data"`
		}
		json.Unmarshal(e.Body, &ev)
		extra := ""
		if o := ev.Data.Object; o != nil {
			for _, k := range []string{"customer", "id", "status"} {
				if v, ok := o[k].(string); ok {
					extra += " " + k + "=" + v
				}
			}
		}
		return fmt.Sprintf("#%d webhook %s %s%s", e.Seq, ev.Type, ev.ID, extra)
	}
	if k := e.Key(keyHeader); k != "" {
		return fmt.Sprintf("#%d %s %s key=%s", e.Seq, e.Method, e.Path, k)
	}
	if len(e.Body) > 0 && len(e.Body) <= 160 {
		return fmt.Sprintf("#%d %s %s %s", e.Seq, e.Method, e.Path, e.Body)
	}
	return fmt.Sprintf("#%d %s %s", e.Seq, e.Method, e.Path)
}

// deliveryOrder is the faulted send order, as repro.sh performs it.
func deliveryOrder(tr []trace.Entry, s replay.Schedule) []string {
	steps := plan(s)
	var out []string
	for _, st := range steps {
		out = append(out, fmt.Sprintf("%s #%d", st.how, st.seq))
	}
	return out
}

type step struct {
	seq int
	how string // send | send-twice | send-concurrently
}

func plan(s replay.Schedule) []step {
	order := append([]int(nil), s.Seqs...)
	pos := map[int]int{}
	for i, q := range order {
		pos[q] = i
	}
	how := map[int]string{}
	for _, f := range s.Faults {
		switch f.Type {
		case replay.Reorder:
			i, j := pos[f.Seq], pos[f.Seq2]
			order[i], order[j] = order[j], order[i]
			pos[f.Seq], pos[f.Seq2] = j, i
		case replay.Lost:
			how[f.Seq] = "send-twice"
		case replay.Concurrent:
			how[f.Seq] = "send-concurrently"
		case replay.Crash:
			how[f.Seq] = fmt.Sprintf("send, kill the service after its commit %d, restart, resend", f.K)
		case replay.ConcurrentAt:
			how[f.Seq] = fmt.Sprintf("send, hold the ack of its commit %d, send a duplicate meanwhile, release", f.K)
		}
	}
	var out []step
	for _, q := range order {
		h := how[q]
		if h == "" {
			h = "send"
		}
		out = append(out, step{q, h})
	}
	return out
}

// WriteRepro writes repro.sh plus one body file per request. The script
// signs Stripe deliveries itself with openssl at send time.
// usage: BASE_URL=http://127.0.0.1:3000 ./repro.sh [faulted|clean]
func WriteRepro(dir string, cfg *config.Config, tr []trace.Entry, s replay.Schedule) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	by := map[int]trace.Entry{}
	for _, e := range tr {
		by[e.Seq] = e
	}
	for _, q := range s.Seqs {
		if err := os.WriteFile(filepath.Join(dir, fmt.Sprintf("%d.json", q)), by[q].Body, 0o644); err != nil {
			return err
		}
	}
	var b strings.Builder
	fmt.Fprintf(&b, `#!/usr/bin/env bash
# Minimal reproduction generated by stutter for %s.
# usage: BASE_URL=http://127.0.0.1:3000 ./repro.sh [faulted|clean]
#   faulted: the delivery order that breaks; clean: the original order.
set -euo pipefail
cd "$(dirname "$0")"
: "${BASE_URL:?set BASE_URL to the running service}"
SECRET=%q
send() {
  local f=$1 path=$2 method=${3:-POST} t sig
  local hdr=()
`, cfg.Name, cfg.Webhooks.Secret)
	var hk []string
	for k := range cfg.Headers {
		hk = append(hk, k)
	}
	sort.Strings(hk)
	b.WriteString("  local STATIC=(")
	for _, k := range hk {
		fmt.Fprintf(&b, " -H %q", k+": "+cfg.Headers[k])
	}
	b.WriteString(" )\n")
	if cfg.Webhooks.Signing == "stripe" {
		b.WriteString(`  t=$(date +%s)
  sig=$(printf '%s.%s' "$t" "$(cat "$f")" | openssl dgst -sha256 -hmac "$SECRET" | awk '{print $NF}')
  hdr=(-H "Stripe-Signature: t=$t,v1=$sig")
`)
	}
	b.WriteString(`  curl -sS -o /dev/null -w "$f %{http_code}\n" -X "$method" "$BASE_URL$path" -H 'Content-Type: application/json' ${hdr[@]+"${hdr[@]}"} ${STATIC[@]+"${STATIC[@]}"} --data-binary @"$f"
}
if [ "${1:-faulted}" = clean ]; then
`)
	for _, q := range s.Seqs {
		fmt.Fprintf(&b, "  send %d.json %q %s\n", q, by[q].Path, by[q].Method)
	}
	b.WriteString("else\n")
	for _, st := range plan(s) {
		p := fmt.Sprintf("%q %s", by[st.seq].Path, by[st.seq].Method)
		switch st.how {
		case "send":
			fmt.Fprintf(&b, "  send %d.json %s\n", st.seq, p)
		case "send-twice":
			fmt.Fprintf(&b, "  send %d.json %s\n  send %d.json %s\n", st.seq, p, st.seq, p)
		case "send-concurrently":
			fmt.Fprintf(&b, "  send %d.json %s & send %d.json %s & wait\n", st.seq, p, st.seq, p)
		default:
			fmt.Fprintf(&b, "  # %s: needs stutter's wire proxy; curl cannot stop a process at a commit\n  send %d.json %s\n", st.how, st.seq, p)
		}
	}
	b.WriteString("fi\n")
	return os.WriteFile(filepath.Join(dir, "repro.sh"), []byte(b.String()), 0o755)
}

func Markdown(r *Result) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# stutter hunt: %s\n\nGenerated %s. Trace: %d requests. Pre-flight: %s.\n\n", r.Target, r.When, r.Requests, r.Preflight)
	b.WriteString("| Fault type | Schedules | Fired | Not fired | Diverged | Nondetermined | Changed the final state | No-op passes |\n|---|---|---|---|---|---|---|---|\n")
	var types []string
	for t := range r.ByType {
		types = append(types, t)
	}
	sort.Strings(types)
	for _, t := range types {
		x := r.ByType[t]
		fmt.Fprintf(&b, "| %s | %d | %d | %d | %d | %d | %d | %d |\n", t, x.Run, x.Fired, x.NotFired, x.Diverged, x.Nondetermined, x.Failed, x.NoOp)
	}
	for _, t := range types {
		var whys []string
		for w := range r.ByType[t].Why {
			whys = append(whys, w)
		}
		sort.Strings(whys)
		for _, w := range whys {
			fmt.Fprintf(&b, "\nNot fired (%s), %s: %s.\n", t, w, strings.Join(r.ByType[t].Why[w], ", "))
		}
	}
	fmt.Fprintf(&b, "\nNo-op reorder pairs (dropping both events leaves the state unchanged, so their pass says nothing about ordering): %d.\n\n", len(r.NoOpPairs))
	// Group findings by the state change they cause.
	effects := map[string][]int{}
	var effOrder []string
	for _, f := range r.Findings {
		var lines []string
		for _, d := range f.Diff {
			lines = append(lines, fmt.Sprintf("%s %s %s", d.Sign, d.Table, d.Row))
		}
		k := strings.Join(lines, "\n")
		if effects[k] == nil {
			effOrder = append(effOrder, k)
		}
		effects[k] = append(effects[k], f.ID)
	}
	fmt.Fprintf(&b, "## Effects: %d distinct end-state changes\n\n", len(effOrder))
	for i, k := range effOrder {
		fmt.Fprintf(&b, "%d. Findings %v:\n\n```\n%s\n```\n\n", i+1, effects[k], k)
	}
	fmt.Fprintf(&b, "## Findings: %d distinct minimal reproductions\n\n", len(r.Findings))
	for _, f := range r.Findings {
		fmt.Fprintf(&b, "### Finding %d\n\nReproduced in %d of %d fresh replays and by repro.sh (curl, fresh databases; -1 = needs the wire proxy (crash or held commit), not scriptable) %d of %d times; 1-minimal: %v (%d single removals could not fire and count as passing). Reached from %d single-fault schedules: %s.\n\n```\n%s\n```\n\nDelivery order: %s\n\nState diff (+ only in the faulted run, - only in the in-order run):\n\n```\n",
			f.ID, f.Replayed, f.ReplayOf, f.CurlRepro, f.ReplayOf, f.OneMinimal, f.Unjudged, len(f.FoundBy), strings.Join(f.FoundBy, ", "), strings.Join(f.Steps, "\n"), strings.Join(f.Delivery, ", "))
		for _, d := range f.Diff {
			fmt.Fprintf(&b, "%s %s x%d %s\n", d.Sign, d.Table, d.Count, d.Row)
		}
		fmt.Fprintf(&b, "```\n\nRepro: `%s/repro.sh`\n\n", filepath.Base(f.Dir))
	}
	fmt.Fprintf(&b, "Hunt time: %.0f s.\n", r.Seconds)
	return b.String()
}

func Write(dir string, r *Result) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	b, _ := json.MarshalIndent(r, "", "  ")
	if err := os.WriteFile(filepath.Join(dir, "results.json"), append(b, '\n'), 0o644); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, "REPORT.md"), []byte(Markdown(r)), 0o644)
}

// VerifyRepro runs a finding's repro.sh, which uses curl and openssl rather
// than stutter's replayer, n times. Each trial gets two fresh clones of the
// template: one receives the clean order, one the faulted order. A trial
// reproduces when the two end states differ.
func VerifyRepro(ctx context.Context, cfg *config.Config, env *pgenv.Env, template, dir string, n int, extraEnv []string) (int, error) {
	opts := cfg.Replay(env, template, extraEnv).Oracle
	run := func(mode string) (*oracle.Canon, error) {
		db, err := env.Clone(ctx, template)
		if err != nil {
			return nil, err
		}
		defer env.Drop(context.Background(), db)
		dbFor := target.DB{DSN: env.DSN(db), URL: env.URL(db), SocketURL: env.SocketURL(db)}
		if cfg.Script() != nil {
			stub, err := agentstub.Start(cfg.Script())
			if err != nil {
				return nil, err
			}
			defer stub.Close()
			dbFor.Vars = map[string]string{"{MODEL_URL}": stub.URL}
		}
		proc, err := target.Start(ctx, cfg.Spec(extraEnv), dbFor)
		if err != nil {
			return nil, err
		}
		cmd := exec.Command(filepath.Join(dir, "repro.sh"), mode)
		cmd.Env = append(os.Environ(), "BASE_URL="+proc.BaseURL)
		out, err := cmd.CombinedOutput()
		proc.Stop()
		if err != nil {
			return nil, fmt.Errorf("repro.sh %s: %v\n%s", mode, err, out)
		}
		conn, err := pgx.Connect(ctx, env.DSN(db))
		if err != nil {
			return nil, err
		}
		defer conn.Close(context.Background())
		st, err := oracle.Snapshot(ctx, conn, opts.Schemas)
		if err != nil {
			return nil, err
		}
		return oracle.Canonicalize(st.Without(opts.IgnoreTables), opts)
	}
	hit := 0
	for i := 0; i < n; i++ {
		clean, err := run("clean")
		if err != nil {
			return hit, err
		}
		faulted, err := run("faulted")
		if err != nil {
			return hit, err
		}
		if len(oracle.Diff(clean, faulted)) > 0 {
			hit++
		}
	}
	return hit, nil
}
