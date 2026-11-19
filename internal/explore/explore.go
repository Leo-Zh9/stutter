// Package explore generates fault schedules, runs them across parallel lanes
// and compares each run with a fault-free baseline of the same trace.
package explore

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/Leo-Zh9/stutter/internal/oracle"
	"github.com/Leo-Zh9/stutter/internal/replay"
	"github.com/Leo-Zh9/stutter/internal/shrink"
	"github.com/Leo-Zh9/stutter/internal/trace"
)

type Runner struct {
	Cfg   replay.Config
	Trace []trace.Entry
	Lanes int
	Seed  int64

	mu        sync.Mutex
	baselines map[string]*oracle.Canon
	runs      int
}

// Outcome of one faulted run.
type Outcome struct {
	Schedule replay.Schedule `json:"schedule"`
	Fired    bool            `json:"fired"` // every fault in the schedule fired
	// Diverged: a scripted model was asked something its script does not
	// cover, in this run or in its baseline. Like a fault that did not
	// fire, it is never a pass and never a failure.
	Diverged bool              `json:"diverged,omitempty"`
	Failed   bool              `json:"failed"`
	Diff     []oracle.DiffLine `json:"diff,omitempty"`
	Err      string            `json:"error,omitempty"`
	// Why says why a held-commit fault did not fire (replay.Why*).
	Why []string `json:"why,omitempty"`
}

func key(seqs []int) string { return fmt.Sprint(seqs) }

// Runs counts every replay this runner has executed, baselines included.
func (r *Runner) Runs() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.runs
}

func (r *Runner) count() {
	r.mu.Lock()
	r.runs++
	r.mu.Unlock()
}

// ErrDiverged marks a baseline whose conversation left the model script, for
// example because shrinking dropped the request that started it. It is not
// cached, and a schedule compared against it neither passes nor fails.
var ErrDiverged = errors.New("baseline left the model script")

// Baseline replays the given subset of the trace with no faults. Results are
// cached per subset.
func (r *Runner) Baseline(ctx context.Context, seqs []int) (*oracle.Canon, error) {
	r.mu.Lock()
	if r.baselines == nil {
		r.baselines = map[string]*oracle.Canon{}
	}
	if c, ok := r.baselines[key(seqs)]; ok {
		r.mu.Unlock()
		return c, nil
	}
	r.mu.Unlock()
	res := replay.Run(ctx, r.Cfg, r.Trace, replay.Schedule{Seqs: seqs}, r.Seed)
	r.count()
	if res.Err != nil {
		return nil, res.Err
	}
	if res.ModelMisses > 0 {
		return nil, fmt.Errorf("%w: %s", ErrDiverged, res.MissReason)
	}
	r.mu.Lock()
	r.baselines[key(seqs)] = res.Canon
	r.mu.Unlock()
	return res.Canon, nil
}

// CheckStable replays the full trace fault-free twice. Different end states
// mean the target has nondeterminism the catalog rule cannot see; the
// offending columns are named so they can be marked in Options.Ignore.
func (r *Runner) CheckStable(ctx context.Context) error {
	if len(r.Cfg.Bands) > 0 {
		return r.inferIDs(ctx)
	}
	s := replay.Full(r.Trace)
	a := replay.Run(ctx, r.Cfg, r.Trace, s, r.Seed)
	b := replay.Run(ctx, r.Cfg, r.Trace, s, r.Seed+1)
	r.count()
	r.count()
	if a.Err != nil {
		return a.Err
	}
	if b.Err != nil {
		return b.Err
	}
	for _, x := range []replay.Result{a, b} {
		if x.ModelMisses > 0 {
			return fmt.Errorf("the fault-free run left the model script (fix the script): %s", x.MissReason)
		}
	}
	if d := oracle.Diff(a.Canon, b.Canon); len(d) > 0 {
		return fmt.Errorf("fault-free replays disagree; unstable columns: %s",
			strings.Join(oracle.UnstableColumns(a.Canon, b.Canon), ", "))
	}
	r.mu.Lock()
	if r.baselines == nil {
		r.baselines = map[string]*oracle.Canon{}
	}
	r.baselines[key(s.Seqs)] = a.Canon
	r.mu.Unlock()
	return nil
}

// Check runs one schedule once and compares it with its baseline.
func (r *Runner) Check(ctx context.Context, s replay.Schedule, seed int64) Outcome {
	base, err := r.Baseline(ctx, s.Seqs)
	if errors.Is(err, ErrDiverged) {
		return Outcome{Schedule: s, Diverged: true}
	}
	if err != nil {
		return Outcome{Schedule: s, Err: err.Error()}
	}
	res := replay.Run(ctx, r.Cfg, r.Trace, s, seed)
	r.count()
	if res.Err != nil {
		return Outcome{Schedule: s, Err: res.Err.Error()}
	}
	if res.ModelMisses > 0 {
		return Outcome{Schedule: s, Diverged: true}
	}
	o := Outcome{Schedule: s, Fired: res.NotFired == 0, Why: res.Why}
	o.Diff = oracle.Diff(base, res.Canon)
	o.Failed = o.Fired && len(o.Diff) > 0
	return o
}

// Parallel runs fn over n items on the runner's lanes, preserving order.
func (r *Runner) Parallel(n int, fn func(i int) Outcome) []Outcome {
	out := make([]Outcome, n)
	lanes := r.Lanes
	if lanes < 1 {
		lanes = 1
	}
	sem := make(chan struct{}, lanes)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(i int) {
			defer wg.Done()
			defer func() { <-sem }()
			out[i] = fn(i)
		}(i)
	}
	wg.Wait()
	return out
}

// Schedules returns every single-fault schedule in a fixed order, then
// `random` multi-fault schedules of 2 to 4 faults drawn with the seed.
func Schedules(tr []trace.Entry, sp replay.Space, seed int64, random int) []replay.Schedule {
	singles := replay.Faults(tr, sp)
	var out []replay.Schedule
	for _, f := range singles {
		out = append(out, replay.Full(tr, f))
	}
	rng := rand.New(rand.NewSource(seed))
	for len(out) < len(singles)+random && len(singles) >= 2 {
		k := 2 + rng.Intn(3)
		used := map[int]bool{}
		var fs []replay.Fault
		for tries := 0; len(fs) < k && tries < 100; tries++ {
			f := singles[rng.Intn(len(singles))]
			if used[f.Seq] || (f.Type == replay.Reorder && used[f.Seq2]) {
				continue // one fault per request keeps every fault observable
			}
			used[f.Seq] = true
			if f.Type == replay.Reorder {
				used[f.Seq2] = true
			}
			fs = append(fs, f)
		}
		out = append(out, replay.Full(tr, fs...))
	}
	return out
}

// FirstFailure runs schedules in order, one lane-sized batch at a time, and
// stops after the batch that holds the first failure. It returns the index
// of the first failing schedule (-1 if none) and every outcome it ran.
func (r *Runner) FirstFailure(ctx context.Context, ss []replay.Schedule) (int, []Outcome) {
	var all []Outcome
	batch := r.Lanes
	if batch < 1 {
		batch = 1
	}
	for lo := 0; lo < len(ss); lo += batch {
		hi := min(lo+batch, len(ss))
		outs := r.Parallel(hi-lo, func(i int) Outcome { return r.Check(ctx, ss[lo+i], r.Seed+int64(lo+i)) })
		all = append(all, outs...)
		for i, o := range outs {
			if o.Failed {
				return lo + i, all
			}
		}
	}
	return -1, all
}

// ShrinkStats describes one shrink.
type ShrinkStats struct {
	ElementsBefore int `json:"elements_before"`
	ElementsAfter  int `json:"elements_after"`
	Candidates     int `json:"candidates_tested"`
	NotFired       int `json:"candidates_not_fired"`
}

// Fails is the shrink predicate: the schedule fails against the baseline of
// its own (possibly reduced) trace. A schedule with a Concurrent fault is
// run k times and fails only if all k runs fired and failed; a run whose
// fault did not fire counts as passing, which can only keep extra elements.
func (r *Runner) Fails(ctx context.Context, s replay.Schedule, k int) (fails, notFired bool) {
	if len(s.Faults) == 0 {
		return false, false
	}
	if !s.HasConcurrent() {
		k = 1
	}
	for i := 0; i < k; i++ {
		o := r.Check(ctx, s, r.Seed+int64(i))
		if !o.Fired {
			return false, true
		}
		if !o.Failed {
			return false, false
		}
	}
	return true, false
}

// elements encode a schedule as ints: 0..len(Seqs)-1 are requests, the rest
// are faults. Dropping a request drops every fault that names it.
func decode(s replay.Schedule, elems []int) replay.Schedule {
	n := len(s.Seqs)
	keep := map[int]bool{}
	var out replay.Schedule
	for _, e := range elems {
		if e < n {
			keep[s.Seqs[e]] = true
		}
	}
	for _, q := range s.Seqs {
		if keep[q] {
			out.Seqs = append(out.Seqs, q)
		}
	}
	for _, e := range elems {
		if e >= n {
			f := s.Faults[e-n]
			if keep[f.Seq] && (f.Type != replay.Reorder || keep[f.Seq2]) {
				out.Faults = append(out.Faults, f)
			}
		}
	}
	return out
}

func encodeAll(s replay.Schedule) []int {
	var out []int
	for i := 0; i < len(s.Seqs)+len(s.Faults); i++ {
		out = append(out, i)
	}
	return out
}

// Shrink reduces a failing schedule to a 1-minimal one.
func (r *Runner) Shrink(ctx context.Context, s replay.Schedule, k int) (replay.Schedule, ShrinkStats) {
	st := ShrinkStats{ElementsBefore: len(s.Seqs) + len(s.Faults)}
	var mu sync.Mutex
	test := func(cands [][]int) []bool {
		res := make([]bool, len(cands))
		r.Parallel(len(cands), func(i int) Outcome {
			f, nf := r.Fails(ctx, decode(s, cands[i]), k)
			mu.Lock()
			st.Candidates++
			if nf {
				st.NotFired++
			}
			mu.Unlock()
			res[i] = f
			return Outcome{}
		})
		return res
	}
	min := decode(s, shrink.DDMin(encodeAll(s), test))
	st.ElementsAfter = len(min.Seqs) + len(min.Faults)
	return min, st
}

// OneMinimal re-checks a schedule by brute force: it must fail, and removing
// any single request or fault must make it pass. A removal whose fault did
// not fire, or whose run left the model script, cannot fail, so it counts as
// passing here; MinimalCheck reports how many removals were like that.
func (r *Runner) OneMinimal(ctx context.Context, s replay.Schedule, k int) bool {
	ok, _ := r.MinimalCheck(ctx, s, k)
	return ok
}

// MinimalCheck is OneMinimal plus the number of single removals that could
// not be judged (not fired, or diverged): "1-minimal over fired candidates".
func (r *Runner) MinimalCheck(ctx context.Context, s replay.Schedule, k int) (bool, int) {
	if f, _ := r.Fails(ctx, s, k); !f {
		return false, 0
	}
	all := encodeAll(s)
	cands := make([][]int, len(all))
	for i := range all {
		for _, e := range all {
			if e != i {
				cands[i] = append(cands[i], e)
			}
		}
	}
	ok, unjudged := true, 0
	var mu sync.Mutex
	r.Parallel(len(cands), func(i int) Outcome {
		f, nf := r.Fails(ctx, decode(s, cands[i]), k)
		mu.Lock()
		if f {
			ok = false
		}
		if nf {
			unjudged++
		}
		mu.Unlock()
		return Outcome{}
	})
	return ok, unjudged
}

// inferIDs runs the fault-free trace three times: twice normally (A, A2) and
// once with every banded sequence shifted by its own offset (B). From A and
// B it infers which columns and JSON paths hold whose sequence ids, then
// requires A, A2 and B to canonicalize to the same state. A mismatch means
// nondeterminism the inference could not explain, and nothing runs.
func (r *Runner) inferIDs(ctx context.Context) error {
	s := replay.Full(r.Trace)
	shifted := r.Cfg
	shifted.Shift = true
	a := replay.Run(ctx, r.Cfg, r.Trace, s, r.Seed)
	a2 := replay.Run(ctx, r.Cfg, r.Trace, s, r.Seed+1)
	b := replay.Run(ctx, shifted, r.Trace, s, r.Seed)
	for i := 0; i < 3; i++ {
		r.count()
	}
	for _, x := range []replay.Result{a, a2, b} {
		if x.Err != nil {
			return x.Err
		}
	}
	conn, err := pgx.Connect(ctx, r.Cfg.Env.DSN(r.Cfg.Template))
	if err != nil {
		return err
	}
	owners, err := oracle.SeqOwners(ctx, conn, r.Cfg.Oracle.Schemas)
	conn.Close(ctx)
	if err != nil {
		return err
	}
	var bands []oracle.Band
	for _, bd := range r.Cfg.Bands {
		bands = append(bands, oracle.Band{Seq: bd.Seq, Start: bd.Start, Width: bd.Width, Offset: bd.Offset})
	}
	inf, err := oracle.Infer(a.State, b.State, bands, owners)
	if err != nil {
		return fmt.Errorf("id inference: %w", err)
	}
	r.Cfg.Oracle.Inf = inf
	ca, err := oracle.Canonicalize(a.State, r.Cfg.Oracle)
	if err != nil {
		return err
	}
	for name, x := range map[string]replay.Result{"second fault-free run": a2, "shifted-sequence run": b} {
		cx, err := oracle.Canonicalize(x.State, r.Cfg.Oracle)
		if err != nil {
			return err
		}
		if d := oracle.Diff(ca, cx); len(d) > 0 {
			return fmt.Errorf("%s disagrees after id inference; unstable columns: %s; first diff: %s %s %s",
				name, strings.Join(oracle.UnstableColumns(ca, cx), ", "), d[0].Sign, d[0].Table, d[0].Row)
		}
	}
	r.mu.Lock()
	r.baselines = map[string]*oracle.Canon{key(s.Seqs): ca}
	r.mu.Unlock()
	return nil
}

// Inference returns what inferIDs learned, if it ran.
func (r *Runner) Inference() *oracle.Inference { return r.Cfg.Oracle.Inf }

// CrashSpace is the crash faults a trace admits.
type CrashSpace struct {
	Faults        []replay.Fault `json:"-"`
	Points        int            `json:"crash_points"`
	Nondetermined []int          `json:"crash_nondeterministic_requests"`
	Background    int            `json:"background_commits"`
	// Pending is what both observation runs agreed on around each resume
	// request; a resume they disagree on is left out (and so not faulted).
	Pending map[int]replay.PendingObs `json:"resume_pending,omitempty"`
	// Holds are concurrent(#n@commitK) for the same commit points, on
	// requests that admit a concurrent duplicate; a hunt runs them only
	// when the config opts in.
	Holds      []replay.Fault `json:"-"`
	HoldPoints int            `json:"concurrent_hold_points,omitempty"`
	// HoldNondetermined lists hold-only requests (not crash-admitted, so
	// not in Nondetermined) whose commit count differs between the runs.
	HoldNondetermined []int `json:"concurrent_hold_nondeterministic_requests,omitempty"`
}

// CrashFaults counts commit points per request over two fault-free runs
// through the wire proxy and returns crash(i,k) for every commit k of every
// retryable request whose count agrees between the two runs. Requests whose
// count differs are listed, never silently dropped. It also lists the held
// commits (CrashSpace.Holds) and sets the runner's HoldTimeout to three times
// the slowest observed request, at least 1s.
func (r *Runner) CrashFaults(ctx context.Context, keyHeader string) (CrashSpace, error) {
	if keyHeader == "" {
		keyHeader = trace.DefaultKeyHeader
	}
	obs := r.Cfg
	obs.ObserveCommits = true
	var counts [2]map[int]int
	var pend [2]map[int]replay.PendingObs
	var cs CrashSpace
	var slowest time.Duration
	for i := 0; i < 2; i++ {
		res := replay.Run(ctx, obs, r.Trace, replay.Full(r.Trace), r.Seed+int64(i))
		r.count()
		if res.Err != nil {
			return cs, res.Err
		}
		if res.ModelMisses > 0 {
			return cs, fmt.Errorf("the observation run left the model script: %s", res.MissReason)
		}
		counts[i] = res.Commits
		pend[i] = res.Pending
		cs.Background += res.BgCommits
		slowest = max(slowest, res.Slowest)
	}
	r.Cfg.HoldTimeout = max(3*slowest, time.Second)
	cs.Pending = map[int]replay.PendingObs{}
	for q, a := range pend[0] {
		if b, ok := pend[1][q]; ok && a == b {
			cs.Pending[q] = a
		}
	}
	sp := replay.Space{KeyHeader: keyHeader, Rules: r.Cfg.Rules, Pending: cs.Pending}
	for _, e := range r.Trace {
		crash, hold := sp.Admits(e, replay.Crash), sp.Admits(e, replay.ConcurrentAt)
		if !crash && !hold {
			continue
		}
		a, b := counts[0][e.Seq], counts[1][e.Seq]
		if a != b {
			if crash {
				cs.Nondetermined = append(cs.Nondetermined, e.Seq)
			} else {
				cs.HoldNondetermined = append(cs.HoldNondetermined, e.Seq)
			}
			continue
		}
		for k := 1; k <= a; k++ {
			if crash {
				cs.Faults = append(cs.Faults, replay.Fault{Type: replay.Crash, Seq: e.Seq, K: k})
				cs.Points++
			}
			if hold {
				cs.Holds = append(cs.Holds, replay.Fault{Type: replay.ConcurrentAt, Seq: e.Seq, K: k})
				cs.HoldPoints++
			}
		}
	}
	return cs, nil
}
