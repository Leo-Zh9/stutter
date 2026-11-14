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
	ss := explore.Schedules(r.Trace, "", 1, 30)
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
	}
	for bug, want := range cases {
		t.Run(bug, func(t *testing.T) {
			r := runner(t, bug)
			idx, outs := r.FirstFailure(context.Background(), explore.Schedules(r.Trace, "", 1, 0))
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
		for _, s := range explore.Schedules(r.Trace, "", 1, 0) {
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
	if idx, _ := r.FirstFailure(context.Background(), explore.Schedules(r.Trace, "", 1, 0)); idx >= 0 {
		t.Fatal("control was detected; the oracle claim in the README is wrong")
	}
}
