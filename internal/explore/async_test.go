package explore_test

import (
	"context"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"testing"

	"github.com/Leo-Zh9/stutter/internal/config"
	"github.com/Leo-Zh9/stutter/internal/explore"
	"github.com/Leo-Zh9/stutter/internal/replay"
	"github.com/Leo-Zh9/stutter/internal/testenv"
	"github.com/Leo-Zh9/stutter/internal/trace"
)

// asyncRunner prepares miniasync (async FastAPI routes, sync and async
// psycopg pools) for one variant and observes its commit points. It runs on
// miniagent's virtualenv and is skipped without it, like agentRunner.
func asyncRunner(t *testing.T, bug string) (*explore.Runner, explore.CrashSpace) {
	t.Helper()
	root := testenv.Root()
	if _, err := os.Stat(filepath.Join(root, "targets", "miniagent", ".venv", "bin", "python")); err != nil {
		if os.Getenv("STUTTER_REQUIRE_AGENT") == "1" {
			t.Fatal("targets/miniagent/.venv missing: run make miniagent")
		}
		t.Skip("targets/miniagent/.venv missing: run make miniagent")
	}
	env := testenv.Env(t)
	cfg, err := config.Load(filepath.Join(root, "targets", "miniasync", "stutter.json"))
	if err != nil {
		t.Fatal(err)
	}
	var extra []string
	if bug != "" {
		extra = []string{"BUG=" + bug}
	}
	name := "tpl_test_async_" + strings.ReplaceAll(bug, "-", "_")
	if err := cfg.PrepareTemplate(context.Background(), env, name, extra); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { env.Drop(context.Background(), name) })
	tr, err := trace.Read(cfg.Path(cfg.Trace))
	if err != nil {
		t.Fatal(err)
	}
	r := &explore.Runner{Cfg: cfg.Replay(env, name, extra), Trace: tr, Lanes: 6, Seed: 1}
	if err := r.CheckStable(context.Background()); err != nil {
		t.Fatal(err)
	}
	cs, err := r.CrashFaults(context.Background(), cfg.KeyHeader)
	if err != nil {
		t.Fatal(err)
	}
	var got []string
	for _, f := range cs.Holds {
		got = append(got, f.String())
	}
	want := "concurrent(#0@commit1) concurrent(#0@commit2) concurrent(#0@commit3) concurrent(#1@commit1) concurrent(#1@commit2) concurrent(#2@commit1) concurrent(#2@commit2)"
	if strings.Join(got, " ") != want || cs.HoldPoints != 7 {
		t.Fatalf("held commits %v (%d points), want %s", got, cs.HoldPoints, want)
	}
	return r, cs
}

func hold(seq, k int) replay.Fault { return replay.Fault{Type: replay.ConcurrentAt, Seq: seq, K: k} }

func wantNotFired(t *testing.T, o explore.Outcome, why string) {
	t.Helper()
	if o.Err != "" {
		t.Fatal(o.Err)
	}
	if o.Fired || !slices.Equal(o.Why, []string{why}) {
		t.Errorf("%v: fired=%v why=%q, want not fired: %q", o.Schedule.Faults, o.Fired, o.Why, why)
	}
}

// The barrier and a hold at the sync commit both freeze the event loop, so
// neither fires; holding the async commit lets the duplicate run between the
// check and the insert, every time, and the finding shrinks to that request.
func TestHeldCommitFindsAsyncRace(t *testing.T) {
	r, _ := asyncRunner(t, "dup-charge")
	ctx := context.Background()
	if o := r.Check(ctx, replay.Full(r.Trace, replay.Fault{Type: replay.Concurrent, Seq: 0}), 1); o.Err != "" || o.Fired {
		t.Fatalf("barrier on the sync write: fired=%v err=%s", o.Fired, o.Err)
	}
	wantNotFired(t, r.Check(ctx, replay.Full(r.Trace, hold(0, 1)), 1), replay.WhyTimeout)
	s := replay.Full(r.Trace, hold(0, 2))
	for i := 0; i < 3; i++ {
		o := r.Check(ctx, s, int64(i))
		if o.Err != "" || !o.Fired || !o.Failed {
			t.Fatalf("run %d: fired=%v failed=%v why=%v err=%s", i, o.Fired, o.Failed, o.Why, o.Err)
		}
		if len(o.Diff) != 1 || o.Diff[0].Table != "charges" || o.Diff[0].Sign != "+" {
			t.Fatalf("run %d: want one extra charge, got %v", i, o.Diff)
		}
	}
	min, _ := r.Shrink(ctx, s, 3)
	if !slices.Equal(min.Seqs, []int{0}) || len(min.Faults) != 1 || min.Faults[0] != hold(0, 2) {
		t.Fatalf("want #0 + concurrent(#0@commit2), got %+v", min)
	}
}

// On the correct variant no hold fails, the async commit still fires, and
// every other hold says why it did not: the frozen loop or the in-process
// lock (no response), the resume guard (the duplicate is never sent), or a
// held write the first attempt does not wait on.
func TestHeldCommitCorrectHasNoFalseAlarms(t *testing.T) {
	r, cs := asyncRunner(t, "")
	ctx := context.Background()
	outs := r.Parallel(len(cs.Holds), func(i int) explore.Outcome {
		return r.Check(ctx, replay.Full(r.Trace, cs.Holds[i]), int64(i))
	})
	for _, o := range outs {
		if o.Err != "" {
			t.Fatal(o.Err)
		}
		if o.Failed {
			t.Errorf("false alarm on %v: %v", o.Schedule.Faults, o.Diff)
		}
	}
	why := map[replay.Fault]string{
		hold(0, 1): replay.WhyTimeout, hold(1, 1): replay.WhyGuard, hold(1, 2): replay.WhyGuard,
		hold(2, 1): replay.WhyTimeout, hold(2, 2): replay.WhyNotHeld,
	}
	for _, o := range outs {
		f := o.Schedule.Faults[0]
		if w, ok := why[f]; ok {
			wantNotFired(t, o, w)
		} else if !o.Fired {
			t.Errorf("%v did not fire: %v", f, o.Why)
		}
	}

	// The guard refuses before the duplicate leaves: the resume arrives once.
	var mu sync.Mutex
	arrived := 0
	cfg := r.Cfg
	cfg.Arrival = func(seq int) {
		mu.Lock()
		if seq == 1 {
			arrived++
		}
		mu.Unlock()
	}
	res := replay.Run(ctx, cfg, r.Trace, replay.Full(r.Trace, hold(1, 1)), 1)
	if res.Err != nil {
		t.Fatal(res.Err)
	}
	if arrived != 1 || res.Fired != 0 || !slices.Equal(res.Why, []string{replay.WhyGuard}) {
		t.Fatalf("arrivals=%d fired=%d why=%v", arrived, res.Fired, res.Why)
	}
}
