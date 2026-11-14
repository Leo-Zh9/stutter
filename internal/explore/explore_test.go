package explore_test

import (
	"context"
	"testing"

	"github.com/Leo-Zh9/stutter/internal/explore"
	"github.com/Leo-Zh9/stutter/internal/replay"
	"github.com/Leo-Zh9/stutter/internal/testenv"
)

func runner(t *testing.T, bug string) *explore.Runner {
	env := testenv.Env(t)
	r := &explore.Runner{Cfg: testenv.Config(t, env, bug), Trace: testenv.Trace(t), Lanes: 8, Seed: 1}
	if err := r.CheckStable(context.Background()); err != nil {
		t.Fatal(err)
	}
	return r
}

// The correct variant: no single fault the trace admits, plus a random
// multi-fault budget, changes the final state.
func TestCorrectVariantHasNoFalseAlarms(t *testing.T) {
	r := runner(t, "")
	ss := explore.Schedules(r.Trace, replay.Space{}, 1, 30)
	cs, err := r.CrashFaults(context.Background(), "")
	if err != nil || cs.Points == 0 || len(cs.Nondetermined) > 0 {
		t.Fatalf("crash space: %+v %v", cs, err)
	}
	for _, f := range cs.Faults {
		ss = append(ss, replay.Full(r.Trace, f))
	}
	outs := r.Parallel(len(ss), func(i int) explore.Outcome { return r.Check(context.Background(), ss[i], int64(i)) })
	for _, o := range outs {
		if o.Err != "" {
			t.Fatal(o.Err)
		}
		if !o.Fired {
			t.Errorf("schedule %v did not fire on the correct variant", o.Schedule.Faults)
		}
		if o.Failed {
			t.Errorf("false alarm on %v: %v", o.Schedule.Faults, o.Diff)
		}
	}
}

// Every in-model bug is found by the fault type that models it, and the
// shrunk schedule still fails and is 1-minimal (re-checked by brute force).
func TestEveryInModelBugIsDetectedAndShrunk(t *testing.T) {
	cases := map[string]replay.FaultType{
		"key-ignored":               replay.Lost,
		"key-after-commit":          replay.Concurrent,
		"webhook-no-dedupe":         replay.Lost,
		"webhook-stale-overwrite":   replay.Reorder,
		"webhook-check-then-insert": replay.Concurrent,
		"outbox-after-commit":       replay.Crash,
		// phases-no-recovery: see TestCrashBugs. Concurrent duplicates hit it
		// only some of the time, so its first failure does not shrink stably.
	}
	for bug, want := range cases {
		t.Run(bug, func(t *testing.T) {
			r := runner(t, bug)
			ss := explore.Schedules(r.Trace, replay.Space{}, 1, 0)
			cs, err := r.CrashFaults(context.Background(), "")
			if err != nil {
				t.Fatal(err)
			}
			for _, f := range cs.Faults {
				ss = append(ss, replay.Full(r.Trace, f))
			}
			idx, outs := r.FirstFailure(context.Background(), ss)
			if idx < 0 {
				t.Fatal("not detected")
			}
			min, _ := r.Shrink(context.Background(), outs[idx].Schedule, 3)
			if len(min.Faults) != 1 || min.Faults[0].Type != want {
				t.Fatalf("shrunk to %v, want one %s fault", min.Faults, want)
			}
			if !r.OneMinimal(context.Background(), min, 3) {
				t.Fatalf("%v is not 1-minimal", min)
			}
		})
	}
}

// key-after-commit is invisible to Lost: a sequential retry finds the key.
// webhook-check-then-insert is invisible to Lost too. Only Concurrent sees them.
func TestRaceBugsNeedConcurrency(t *testing.T) {
	for _, bug := range []string{"key-after-commit", "webhook-check-then-insert"} {
		r := runner(t, bug)
		for _, s := range explore.Schedules(r.Trace, replay.Space{}, 1, 0) {
			if s.Faults[0].Type != replay.Lost {
				continue
			}
			if o := r.Check(context.Background(), s, 1); o.Failed {
				t.Fatalf("%s: Lost alone detected it (%v); the bug is not a pure race", bug, s.Faults)
			}
		}
	}
}

// The control: a key not bound to its body is wrong without any fault, so a
// differential oracle cannot see it. The test pins that limit.
func TestControlIsOutsideTheOracle(t *testing.T) {
	r := runner(t, "key-unscoped")
	if idx, _ := r.FirstFailure(context.Background(), explore.Schedules(r.Trace, replay.Space{}, 1, 0)); idx >= 0 {
		t.Fatal("control was detected; the oracle claim in the README is wrong")
	}
}

// outbox-after-commit survives every lost and concurrent fault: only killing
// the service between its two commits exposes it. phases-no-recovery is
// exposed by a crash between phases and also by a concurrent duplicate that
// takes the key's row lock between phases.
func TestCrashBugs(t *testing.T) {
	r := runner(t, "outbox-after-commit")
	ss := explore.Schedules(r.Trace, replay.Space{}, 1, 0)
	outs := r.Parallel(len(ss), func(i int) explore.Outcome { return r.Check(context.Background(), ss[i], int64(i)) })
	for _, o := range outs {
		if o.Failed {
			t.Errorf("outbox-after-commit: %v alone detected it", o.Schedule.Faults)
		}
	}
	for _, bug := range []string{"outbox-after-commit", "phases-no-recovery"} {
		r := runner(t, bug)
		cs, err := r.CrashFaults(context.Background(), "")
		if err != nil {
			t.Fatal(err)
		}
		var crash []replay.Schedule
		for _, f := range cs.Faults {
			crash = append(crash, replay.Full(r.Trace, f))
		}
		idx, outs := r.FirstFailure(context.Background(), crash)
		if idx < 0 {
			t.Fatalf("%s: no crash point detected it", bug)
		}
		min, _ := r.Shrink(context.Background(), outs[idx].Schedule, 3)
		if len(min.Faults) != 1 || min.Faults[0].Type != replay.Crash || !r.OneMinimal(context.Background(), min, 3) {
			t.Fatalf("%s: shrunk to %v", bug, min)
		}
	}
}
