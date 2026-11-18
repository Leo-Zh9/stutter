package replay

import (
	"testing"

	"github.com/Leo-Zh9/stutter/internal/trace"
)

// A resume admits concurrent only with an interrupt pending before it, and
// lost or crash only if none is pending after it; an unobserved resume
// admits nothing.
func TestResumeAdmission(t *testing.T) {
	rules := []trace.Rule{{Method: "POST", Path: "/threads/*/approve", Resume: &trace.ResumeRule{Thread: "path:2"}}}
	e := func(seq int) trace.Entry {
		return trace.Entry{Seq: seq, Kind: trace.API, Method: "POST", Path: "/threads/t/approve"}
	}
	sp := Space{Rules: rules, Pending: map[int]PendingObs{
		1: {Before: true, After: false}, // completes the run
		2: {Before: true, After: true},  // leads to a second approval
		3: {Before: false, After: false},
	}}
	cases := []struct {
		seq        int
		lost, conc bool
	}{{1, true, true}, {2, false, true}, {3, false, false}, {4, false, false}}
	for _, c := range cases {
		if got := sp.Admits(e(c.seq), Lost); got != c.lost {
			t.Errorf("seq %d lost: %v", c.seq, got)
		}
		if got := sp.Admits(e(c.seq), Crash); got != c.lost {
			t.Errorf("seq %d crash: %v", c.seq, got)
		}
		if got := sp.Admits(e(c.seq), Concurrent); got != c.conc {
			t.Errorf("seq %d concurrent: %v", c.seq, got)
		}
	}
}
