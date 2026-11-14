package oracle

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
)

// Band mirrors pgenv.Band: the range one sequence mints into, and the shift
// applied to it in the inference baseline.
type Band struct {
	Seq    string `json:"sequence"`
	Start  int64  `json:"start"`
	Width  int64  `json:"width"`
	Offset int64  `json:"offset"`
}

func (b Band) Contains(v int64) bool { return v >= b.Start && v < b.Start+b.Width }

// Inference says which columns and JSON paths hold ids minted by which
// sequence, and which differ between fault-free runs for other reasons.
//
// It compares two fault-free replays of the same trace: A on a normal clone,
// B on a clone where every banded sequence was advanced by its own offset.
// A path whose values in B are exactly A's values plus sequence k's offset
// (as a multiset; no row matching needed) holds sequence k's ids. Anything
// else that differs is volatile (hashes over ids, app-side clocks) and is
// left out of the comparison and listed.
type Inference struct {
	Derived  map[string]int    `json:"derived"`  // path -> band index
	Owners   map[int]string    `json:"owners"`   // band index -> owning table.column
	Volatile []string          `json:"volatile"` // paths dropped from the comparison
	Bands    []Band            `json:"bands"`
	Unowned  map[string]string `json:"-"`
}

// SeqOwners reads OWNED BY links (serial and identity columns, and
// sequences declared "owned by" a column).
func SeqOwners(ctx context.Context, conn *pgx.Conn, schemas []string) (map[string]string, error) {
	rows, err := conn.Query(ctx, `
		SELECT sn.nspname || '.' || s.relname,
		       CASE WHEN tn.nspname = 'public' THEN t.relname ELSE tn.nspname || '.' || t.relname END || '.' || a.attname
		FROM pg_depend d
		JOIN pg_class s ON s.oid = d.objid AND s.relkind = 'S'
		JOIN pg_namespace sn ON sn.oid = s.relnamespace
		JOIN pg_class t ON t.oid = d.refobjid
		JOIN pg_namespace tn ON tn.oid = t.relnamespace
		JOIN pg_attribute a ON a.attrelid = t.oid AND a.attnum = d.refobjsubid
		WHERE d.classid = 'pg_class'::regclass AND d.refclassid = 'pg_class'::regclass
		  AND d.deptype IN ('a', 'i') AND sn.nspname = ANY($1)`, schemas)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := map[string]string{}
	for rows.Next() {
		var seq, col string
		if err := rows.Scan(&seq, &col); err != nil {
			return nil, err
		}
		out[seq] = col
	}
	return out, rows.Err()
}

// leaves flattens a state into path -> multiset of leaf values. Scalar
// columns are "table.col"; JSON inside a column adds ".key" per object level
// and "[]" per array level.
func leaves(st *State, skip func(table string, c Column) bool) map[string][]string {
	out := map[string][]string{}
	var walk func(path string, v any)
	walk = func(path string, v any) {
		switch x := v.(type) {
		case map[string]any:
			for k, vv := range x {
				walk(path+"."+k, vv)
			}
		case []any:
			for _, vv := range x {
				walk(path+"[]", vv)
			}
		default:
			out[path] = append(out[path], fmt.Sprint(x))
		}
	}
	for _, t := range st.Tables {
		for _, c := range t.Columns {
			if skip(t.Name, c) {
				continue
			}
			for _, row := range st.Rows[t.Name] {
				walk(t.Name+"."+c.Name, row[c.Name])
			}
		}
	}
	return out
}

func Infer(a, b *State, bands []Band, owners map[string]string) (*Inference, error) {
	skip := func(_ string, c Column) bool { return c.Temporal }
	la, lb := leaves(a, skip), leaves(b, skip)
	inf := &Inference{Derived: map[string]int{}, Owners: map[int]string{}, Bands: bands}
	var paths []string
	for p := range la {
		paths = append(paths, p)
	}
	for p := range lb {
		if _, ok := la[p]; !ok {
			paths = append(paths, p)
		}
	}
	sort.Strings(paths)
	for _, p := range paths {
		va, vb := la[p], lb[p]
		if sameMultiset(va, vb) {
			continue
		}
		if k, ok := shiftedBy(va, vb, bands); ok {
			inf.Derived[p] = k
			continue
		}
		inf.Volatile = append(inf.Volatile, p)
	}
	for k, band := range bands {
		col, ok := owners[band.Seq]
		if !ok {
			// No OWNED BY link. Only an error if some column holds its ids.
			for p, kk := range inf.Derived {
				if kk == k {
					return nil, fmt.Errorf("sequence %s has no owning column (no OWNED BY link) but %s holds its values", band.Seq, p)
				}
			}
			continue
		}
		if kk, derived := inf.Derived[col]; !derived || kk != k {
			if inBand(la[col], band) == 0 {
				continue // the sequence minted nothing during this trace
			}
			return nil, fmt.Errorf("sequence %s is owned by %s, but that column does not shift with it", band.Seq, col)
		}
		seen := map[string]bool{}
		for _, v := range la[col] {
			if seen[v] {
				return nil, fmt.Errorf("owner %s of %s repeats value %s: cannot resolve references to it", col, band.Seq, v)
			}
			seen[v] = true
		}
		inf.Owners[k] = col
	}
	return inf, nil
}

func sameMultiset(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	m := map[string]int{}
	for _, x := range a {
		m[x]++
	}
	for _, x := range b {
		m[x]--
		if m[x] < 0 {
			return false
		}
	}
	return true
}

// shiftedBy reports the band k such that b is a with every in-band value
// shifted by k's offset. Values outside every band (rows the template
// already held) must be identical in a and b. All in-band values must belong
// to the same band, and there must be at least one.
func shiftedBy(a, b []string, bands []Band) (int, bool) {
	if len(a) == 0 || len(a) != len(b) {
		return 0, false
	}
	ia, ok1 := ints(a)
	ib, ok2 := ints(b)
	if !ok1 || !ok2 {
		return 0, false
	}
	bandOf := func(v int64) int {
		for k, band := range bands {
			if band.Contains(v) {
				return k
			}
		}
		return -1
	}
	k := -1
	var want []string
	for _, v := range ia {
		bk := bandOf(v)
		if bk < 0 {
			want = append(want, strconv.FormatInt(v, 10))
			continue
		}
		if k >= 0 && bk != k {
			return 0, false
		}
		k = bk
		want = append(want, strconv.FormatInt(v+bands[bk].Offset, 10))
	}
	if k < 0 {
		return 0, false
	}
	got := make([]string, len(ib))
	for i, v := range ib {
		got[i] = strconv.FormatInt(v, 10)
	}
	return k, sameMultiset(want, got)
}

// inBand counts values of a path that fall inside band k.
func inBand(xs []string, band Band) int {
	n := 0
	for _, x := range xs {
		if v, err := strconv.ParseInt(x, 10, 64); err == nil && band.Contains(v) {
			n++
		}
	}
	return n
}

func ints(xs []string) ([]int64, bool) {
	out := make([]int64, len(xs))
	for i, x := range xs {
		if x == "<nil>" {
			return nil, false
		}
		n, err := strconv.ParseInt(strings.TrimSpace(x), 10, 64)
		if err != nil {
			return nil, false
		}
		out[i] = n
	}
	return out, true
}

// numberOf returns an integer leaf value.
func numberOf(v any) (int64, bool) {
	switch x := v.(type) {
	case json.Number:
		n, err := x.Int64()
		return n, err == nil
	case float64:
		return int64(x), float64(int64(x)) == x
	}
	return 0, false
}
