// Package shrink reduces a failing fault schedule with delta debugging.
package shrink

// DDMin is Zeller's ddmin. elems is a failing configuration; test evaluates
// a batch of candidate configurations (it may run them in parallel) and
// reports which fail. The search always takes the lowest-index failing
// candidate, so the result is the same as a sequential run.
//
// On return the configuration is 1-minimal: removing any single element
// makes test report it as passing. That holds because the loop only stops
// at granularity n == len(c), where every complement is c minus one element
// and all of them passed.
func DDMin(elems []int, test func([][]int) []bool) []int {
	c := append([]int(nil), elems...)
	n := 2
	for len(c) >= 2 {
		subsets := split(c, n)
		if i := firstTrue(test(subsets)); i >= 0 {
			c = subsets[i]
			n = max(n-1, 2)
			continue
		}
		var comps [][]int
		for i := range subsets {
			comps = append(comps, without(c, subsets[i]))
		}
		if i := firstTrue(test(comps)); i >= 0 {
			c = comps[i]
			n = max(n-1, 2)
			continue
		}
		if n >= len(c) {
			break
		}
		n = min(len(c), 2*n)
	}
	return c
}

func split(c []int, n int) [][]int {
	var out [][]int
	start := 0
	for i := 0; i < n; i++ {
		end := start + (len(c)-start)/(n-i)
		out = append(out, append([]int(nil), c[start:end]...))
		start = end
	}
	return out
}

func without(c, drop []int) []int {
	d := map[int]bool{}
	for _, x := range drop {
		d[x] = true
	}
	var out []int
	for _, x := range c {
		if !d[x] {
			out = append(out, x)
		}
	}
	return out
}

func firstTrue(bs []bool) int {
	for i, b := range bs {
		if b {
			return i
		}
	}
	return -1
}
