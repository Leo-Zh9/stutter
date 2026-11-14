package shrink

import (
	"math/rand"
	"testing"
)

func seqTest(pred func([]int) bool) func([][]int) []bool {
	return func(cs [][]int) []bool {
		out := make([]bool, len(cs))
		for i, c := range cs {
			out[i] = pred(c)
		}
		return out
	}
}

func containsAll(c []int, need []int) bool {
	have := map[int]bool{}
	for _, x := range c {
		have[x] = true
	}
	for _, x := range need {
		if !have[x] {
			return false
		}
	}
	return true
}

func TestDDMinFindsExactCause(t *testing.T) {
	all := make([]int, 67)
	for i := range all {
		all[i] = i
	}
	got := DDMin(all, seqTest(func(c []int) bool { return containsAll(c, []int{3, 17, 66}) }))
	if len(got) != 3 || !containsAll(got, []int{3, 17, 66}) {
		t.Fatalf("got %v", got)
	}
}

// Property: for random failure causes, the result still fails and removing
// any single element makes it pass (1-minimality), checked by brute force.
func TestDDMinOneMinimalProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(7))
	for trial := 0; trial < 300; trial++ {
		n := 1 + rng.Intn(80)
		all := make([]int, n)
		for i := range all {
			all[i] = i
		}
		// Failure needs every element of one of several random causes:
		// a disjunction of conjunctions, so several minimal answers exist.
		var causes [][]int
		for k := 0; k < 1+rng.Intn(3); k++ {
			var c []int
			for j := 0; j < 1+rng.Intn(4); j++ {
				c = append(c, rng.Intn(n))
			}
			causes = append(causes, c)
		}
		pred := func(c []int) bool {
			for _, cause := range causes {
				if containsAll(c, cause) {
					return true
				}
			}
			return false
		}
		got := DDMin(all, seqTest(pred))
		if !pred(got) {
			t.Fatalf("trial %d: result %v does not fail", trial, got)
		}
		for i := range got {
			drop := append(append([]int(nil), got[:i]...), got[i+1:]...)
			if pred(drop) {
				t.Fatalf("trial %d: %v is not 1-minimal, %v still fails", trial, got, drop)
			}
		}
	}
}
