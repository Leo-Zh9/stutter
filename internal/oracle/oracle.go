// Package oracle snapshots a database and decides whether two snapshots hold
// the same data once values the database generated for itself (ids,
// timestamps) are factored out.
//
// The rule: a column is "generated" when its default calls a non-immutable
// function (gen_random_uuid, nextval, now, ...) or it is an identity column.
// Generated values are never compared directly. Instead, any place that
// points at a generated value (a foreign key, or a string inside a text or
// jsonb column that exactly equals one) becomes an edge to the row that owns
// it. Each row then gets a signature built bottom-up from its own data plus
// the sorted signatures of the rows that point at it (AHU tree hashing), and
// two states are equal when every table holds the same multiset of
// signatures. Swapping which uuid a row received, or the order rows were
// inserted, cannot change a signature; adding, removing or changing a row
// does.
package oracle

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

type Column struct {
	Name      string
	Type      string
	Generated bool
	Temporal  bool
	FKTable   string // set when this column references a generated column of FKTable
	FKColumn  string
}

type Table struct {
	Name    string
	Columns []Column
}

type Row map[string]any

type State struct {
	Tables []Table
	Rows   map[string][]Row
	// Opaque values were minted outside the database for this run (a
	// scripted model's message and tool-call ids). Like generated ids, they
	// differ between equivalent runs, so any string equal to one compares
	// as its label instead of its value.
	Opaque map[string]string
}

var temporalTypes = map[string]bool{
	"timestamp with time zone": true, "timestamp without time zone": true,
	"date": true, "time with time zone": true, "time without time zone": true, "interval": true,
}

var precision = regexp.MustCompile(`\(\d+\)`)

// isTemporal matches date and time types with or without a precision
// modifier: Prisma declares timestamp(3) without time zone.
func isTemporal(typ string) bool {
	return temporalTypes[strings.TrimSpace(precision.ReplaceAllString(typ, ""))] ||
		temporalTypes[strings.Join(strings.Fields(precision.ReplaceAllString(typ, "")), " ")]
}

var funcCall = regexp.MustCompile(`([A-Za-z_][A-Za-z0-9_]*)\s*\(`)
var sqlValueFuncs = regexp.MustCompile(`(?i)\b(CURRENT_TIMESTAMP|CURRENT_DATE|CURRENT_TIME|LOCALTIMESTAMP|LOCALTIME)\b`)

// Snapshot reads every ordinary table in the given schemas. Tables outside
// public are named schema.table.
func Snapshot(ctx context.Context, conn *pgx.Conn, schemas []string) (*State, error) {
	if len(schemas) == 0 {
		schemas = []string{"public"}
	}
	nonImmutable := map[string]bool{}
	rows, err := conn.Query(ctx, `SELECT DISTINCT proname FROM pg_proc WHERE provolatile IN ('v','s')`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var n string
		rows.Scan(&n)
		nonImmutable[n] = true
	}
	rows.Close()

	rows, err = conn.Query(ctx, `
		SELECT CASE WHEN n.nspname = 'public' THEN c.relname ELSE n.nspname || '.' || c.relname END,
		       a.attname, format_type(a.atttypid, a.atttypmod),
		       COALESCE(pg_get_expr(d.adbin, d.adrelid), ''), a.attidentity <> ''
		FROM pg_class c
		JOIN pg_namespace n ON n.oid = c.relnamespace AND n.nspname = ANY($1)
		JOIN pg_attribute a ON a.attrelid = c.oid AND a.attnum > 0 AND NOT a.attisdropped
		LEFT JOIN pg_attrdef d ON d.adrelid = c.oid AND d.adnum = a.attnum
		WHERE c.relkind IN ('r','p') AND NOT c.relispartition
		ORDER BY 1, a.attnum`, schemas)
	if err != nil {
		return nil, err
	}
	byName := map[string]*Table{}
	var order []string
	for rows.Next() {
		var tbl, col, typ, def string
		var identity bool
		if err := rows.Scan(&tbl, &col, &typ, &def, &identity); err != nil {
			return nil, err
		}
		t, ok := byName[tbl]
		if !ok {
			t = &Table{Name: tbl}
			byName[tbl] = t
			order = append(order, tbl)
		}
		gen := identity || sqlValueFuncs.MatchString(def)
		for _, m := range funcCall.FindAllStringSubmatch(def, -1) {
			if nonImmutable[strings.ToLower(m[1])] {
				gen = true
			}
		}
		t.Columns = append(t.Columns, Column{Name: col, Type: typ, Generated: gen, Temporal: isTemporal(typ)})
	}
	rows.Close()

	// Single-column foreign keys.
	rows, err = conn.Query(ctx, `
		SELECT CASE WHEN cn.nspname = 'public' THEN cl.relname ELSE cn.nspname || '.' || cl.relname END, a.attname,
		       CASE WHEN pn.nspname = 'public' THEN pc.relname ELSE pn.nspname || '.' || pc.relname END, pa.attname
		FROM pg_constraint k
		JOIN pg_class cl ON cl.oid = k.conrelid
		JOIN pg_namespace cn ON cn.oid = cl.relnamespace
		JOIN pg_class pc ON pc.oid = k.confrelid
		JOIN pg_namespace pn ON pn.oid = pc.relnamespace
		JOIN pg_attribute a ON a.attrelid = k.conrelid AND a.attnum = k.conkey[1]
		JOIN pg_attribute pa ON pa.attrelid = k.confrelid AND pa.attnum = k.confkey[1]
		WHERE k.contype = 'f' AND array_length(k.conkey, 1) = 1`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var ct, cc, pt, pcol string
		if err := rows.Scan(&ct, &cc, &pt, &pcol); err != nil {
			return nil, err
		}
		if p, ok := byName[pt]; ok && p.col(pcol) != nil && p.col(pcol).Generated && byName[ct] != nil {
			if c := byName[ct].col(cc); c != nil {
				c.FKTable, c.FKColumn = pt, pcol
			}
		}
	}
	rows.Close()

	st := &State{Rows: map[string][]Row{}}
	for _, name := range order {
		st.Tables = append(st.Tables, *byName[name])
		sch, tbl := "public", name
		if i := strings.Index(name, "."); i >= 0 {
			sch, tbl = name[:i], name[i+1:]
		}
		q := fmt.Sprintf(`SELECT to_jsonb(t)::text FROM %s t`, pgx.Identifier{sch, tbl}.Sanitize())
		r, err := conn.Query(ctx, q)
		if err != nil {
			return nil, err
		}
		for r.Next() {
			var s string
			if err := r.Scan(&s); err != nil {
				return nil, err
			}
			dec := json.NewDecoder(strings.NewReader(s))
			dec.UseNumber()
			var row Row
			if err := dec.Decode(&row); err != nil {
				return nil, err
			}
			st.Rows[name] = append(st.Rows[name], row)
		}
		r.Close()
		if err := r.Err(); err != nil {
			return nil, err
		}
	}
	return st, nil
}

func (t *Table) col(name string) *Column {
	for i := range t.Columns {
		if t.Columns[i].Name == name {
			return &t.Columns[i]
		}
	}
	return nil
}

// Options lets a user mark extra columns as nondeterministic ("table.column").
type Options struct {
	// Ignore drops a column ("table.column") or a path inside a JSON
	// column ("table.column.key.key", "[]" for array elements).
	Ignore map[string]bool
	// Generated marks columns whose ids the application mints itself (a
	// uuid4 in app code, no column default): they are treated exactly like
	// columns the catalog shows as generated.
	Generated map[string]bool
	Schemas   []string // schemas to snapshot; default public
	// IgnoreTables are left out of the snapshot entirely (framework
	// bookkeeping such as a checkpointer's tables), and out of the lock
	// barrier's table list.
	IgnoreTables map[string]bool
	// Inf, when set, adds inferred integer ids (see Infer): owners act as
	// generated columns, derived paths become edges, volatile paths drop.
	Inf *Inference
}

type rowRef struct {
	table string
	idx   int
}

// Canon is a canonical state: per table, a multiset of row signatures.
type Canon struct {
	Tables map[string]map[string]int
	Show   map[string]string         // signature hash -> readable own content
	Cols   map[string]map[string]int // "table.col" -> normalized value -> count
}

type edge struct {
	label string // "child_table.column"
	child rowRef
}

// Canonicalize builds the canonical form of a state.
func Canonicalize(st *State, opt Options) (*Canon, error) {
	tables := map[string]Table{}
	for _, t := range st.Tables {
		tables[t.Name] = t
	}
	// Index every generated value.
	genStr := map[string]rowRef{}
	genAny := map[string]map[string]rowRef{} // "table.col" -> value -> row
	for _, t := range st.Tables {
		for i, row := range st.Rows[t.Name] {
			for _, c := range t.Columns {
				if !(c.Generated || opt.Generated[t.Name+"."+c.Name]) || row[c.Name] == nil {
					continue
				}
				k := t.Name + "." + c.Name
				if genAny[k] == nil {
					genAny[k] = map[string]rowRef{}
				}
				v := fmt.Sprint(row[c.Name])
				genAny[k][v] = rowRef{t.Name, i}
				if s, ok := row[c.Name].(string); ok {
					genStr[s] = rowRef{t.Name, i}
				}
			}
		}
	}

	// Inferred integer ids: owner column values, per band.
	var volatile map[string]bool
	ownerOf := map[string]int{}
	genNum := map[int]map[int64]rowRef{}
	if inf := opt.Inf; inf != nil {
		volatile = map[string]bool{}
		for _, v := range inf.Volatile {
			volatile[v] = true
		}
		for k, col := range inf.Owners {
			ownerOf[col] = k
			genNum[k] = map[int64]rowRef{}
		}
		for _, t := range st.Tables {
			for i, row := range st.Rows[t.Name] {
				for _, c := range t.Columns {
					if k, ok := ownerOf[t.Name+"."+c.Name]; ok {
						if n, ok := numberOf(row[c.Name]); ok {
							genNum[k][n] = rowRef{t.Name, i}
						}
					}
				}
			}
		}
	}
	if len(opt.Ignore) > 0 {
		v := map[string]bool{}
		for k := range volatile {
			v[k] = true
		}
		for k := range opt.Ignore {
			v[k] = true
		}
		volatile = v
	}
	numRef := func(path string, v any, me rowRef, label string, children map[rowRef][]edge) (string, bool) {
		if opt.Inf == nil {
			return "", false
		}
		k, ok := opt.Inf.Derived[path]
		if !ok {
			return "", false
		}
		if _, isOwner := ownerOf[path]; isOwner {
			return "", false
		}
		n, ok := numberOf(v)
		if !ok {
			return "", false
		}
		p, ok := genNum[k][n]
		if !ok {
			return "<dangling:" + opt.Inf.Bands[k].Seq + ">", true
		}
		if p == me {
			return "<self>", true
		}
		children[p] = append(children[p], edge{label, me})
		return "<ref:" + p.table + ">", true
	}

	own := map[rowRef]string{}
	cols := map[string]map[string]int{}
	children := map[rowRef][]edge{}
	for _, t := range st.Tables {
		for i, row := range st.Rows[t.Name] {
			me := rowRef{t.Name, i}
			fields := map[string]any{}
			for _, c := range t.Columns {
				label := t.Name + "." + c.Name
				if c.Generated || opt.Generated[label] || c.Temporal || opt.Ignore[label] || volatile[label] {
					continue
				}
				if _, isOwner := ownerOf[label]; isOwner {
					continue // an inferred id column: generated in all but name
				}
				v := row[c.Name]
				if r, ok := numRef(label, v, me, label, children); ok {
					fields[c.Name] = r
					continue
				}
				if c.FKTable != "" && v != nil {
					if p, ok := genAny[c.FKTable+"."+c.FKColumn][fmt.Sprint(v)]; ok {
						fields[c.Name] = "<ref:" + p.table + ">"
						if p != me {
							children[p] = append(children[p], edge{label, me})
						}
						continue
					}
				}
				fields[c.Name] = normalizeAt(label, v, volatile, func(path string, x any) (string, bool) {
					return numRef(path, x, me, label, children)
				}, func(s string) (string, bool) {
					p, ok := genStr[s]
					if !ok {
						if l, ok := st.Opaque[s]; ok {
							return l, true
						}
						return "", false
					}
					if p == me {
						return "<self>", true
					}
					children[p] = append(children[p], edge{label, me})
					return "<ref:" + p.table + ">", true
				})
			}
			for k, v := range fields {
				j := marshal(v)
				key := t.Name + "." + k
				if cols[key] == nil {
					cols[key] = map[string]int{}
				}
				cols[key][j]++
			}
			own[me] = t.Name + marshal(fields) // map keys marshal sorted
		}
	}

	sig := map[rowRef]string{}
	visiting := map[rowRef]bool{}
	var compute func(r rowRef) (string, error)
	compute = func(r rowRef) (string, error) {
		if s, ok := sig[r]; ok {
			return s, nil
		}
		if visiting[r] {
			return "", fmt.Errorf("reference cycle through %s row %d", r.table, r.idx)
		}
		visiting[r] = true
		var kids []string
		for _, e := range children[r] {
			cs, err := compute(e.child)
			if err != nil {
				return "", err
			}
			kids = append(kids, e.label+"="+cs)
		}
		sort.Strings(kids)
		h := sha256.New()
		h.Write([]byte(own[r]))
		for _, k := range kids {
			h.Write([]byte{0})
			h.Write([]byte(k))
		}
		s := hex.EncodeToString(h.Sum(nil))[:16]
		visiting[r] = false
		sig[r] = s
		return s, nil
	}

	out := &Canon{Tables: map[string]map[string]int{}, Show: map[string]string{}, Cols: cols}
	for _, t := range st.Tables {
		out.Tables[t.Name] = map[string]int{}
		for i := range st.Rows[t.Name] {
			r := rowRef{t.Name, i}
			s, err := compute(r)
			if err != nil {
				return nil, err
			}
			out.Tables[t.Name][s]++
			show := strings.TrimPrefix(own[r], t.Name)
			if n := len(children[r]); n > 0 {
				show += fmt.Sprintf(" (+%d referencing rows)", n)
			}
			out.Show[s] = show
		}
	}
	return out, nil
}

// normalizeAt is normalize with JSON paths: leaves at volatile paths are
// dropped, and integer leaves at inferred id paths become references.
func normalizeAt(path string, v any, volatile map[string]bool, num func(string, any) (string, bool), ref func(string) (string, bool)) any {
	switch x := v.(type) {
	case map[string]any:
		m := make(map[string]any, len(x))
		for k, vv := range x {
			p := path + "." + k
			if volatile[p] {
				continue
			}
			m[k] = normalizeAt(p, vv, volatile, num, ref)
		}
		return m
	case []any:
		a := make([]any, len(x))
		for i, vv := range x {
			a[i] = normalizeAt(path+"[]", vv, volatile, num, ref)
		}
		return a
	case json.Number, float64:
		if r, ok := num(path, x); ok {
			return r
		}
		return x
	case string:
		// Ids are sometimes stored as decimal strings (Formance's revert
		// links in metadata). Only paths the inference attributed count.
		if _, err := strconv.ParseInt(x, 10, 64); err == nil {
			if r, ok := num(path, json.Number(x)); ok {
				return r
			}
		}
		return normalize(v, ref)
	default:
		return normalize(v, ref)
	}
}

// normalize rewrites string leaves: generated ids become references and
// RFC 3339 timestamps become <ts>. Objects and arrays are walked.
func normalize(v any, ref func(string) (string, bool)) any {
	switch x := v.(type) {
	case string:
		if r, ok := ref(x); ok {
			return r
		}
		if isTimestamp(x) {
			return "<ts>"
		}
		return x
	case map[string]any:
		m := make(map[string]any, len(x))
		for k, vv := range x {
			m[k] = normalize(vv, ref)
		}
		return m
	case []any:
		a := make([]any, len(x))
		for i, vv := range x {
			a[i] = normalize(vv, ref)
		}
		return a
	default:
		return v
	}
}

func isTimestamp(s string) bool {
	if len(s) < 20 || s[4] != '-' || s[10] != 'T' {
		return false
	}
	_, err := time.Parse(time.RFC3339Nano, s)
	return err == nil
}

// DiffLine is one row present more times in one state than the other.
type DiffLine struct {
	Table string `json:"table"`
	Sign  string `json:"sign"` // "+" only in the faulted run, "-" only in the baseline
	Count int    `json:"count"`
	Row   string `json:"row"`
}

// Diff compares baseline a with faulted b. An empty result means equal.
func Diff(a, b *Canon) []DiffLine {
	var out []DiffLine
	names := map[string]bool{}
	for t := range a.Tables {
		names[t] = true
	}
	for t := range b.Tables {
		names[t] = true
	}
	var sorted []string
	for t := range names {
		sorted = append(sorted, t)
	}
	sort.Strings(sorted)
	for _, t := range sorted {
		sigs := map[string]bool{}
		for s := range a.Tables[t] {
			sigs[s] = true
		}
		for s := range b.Tables[t] {
			sigs[s] = true
		}
		var keys []string
		for s := range sigs {
			keys = append(keys, s)
		}
		sort.Strings(keys)
		for _, s := range keys {
			d := b.Tables[t][s] - a.Tables[t][s]
			switch {
			case d > 0:
				out = append(out, DiffLine{t, "+", d, b.Show[s]})
			case d < 0:
				out = append(out, DiffLine{t, "-", -d, a.Show[s]})
			}
		}
	}
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Table != out[j].Table {
			return out[i].Table < out[j].Table
		}
		return out[i].Sign < out[j].Sign
	})
	return out
}

// Digest is a stable hash of a canonical state, for determinism checks.
func (c *Canon) Digest() string {
	var buf bytes.Buffer
	var names []string
	for t := range c.Tables {
		names = append(names, t)
	}
	sort.Strings(names)
	for _, t := range names {
		var keys []string
		for s := range c.Tables[t] {
			keys = append(keys, s)
		}
		sort.Strings(keys)
		for _, s := range keys {
			fmt.Fprintf(&buf, "%s %s %d\n", t, s, c.Tables[t][s])
		}
	}
	sum := sha256.Sum256(buf.Bytes())
	return hex.EncodeToString(sum[:])[:16]
}

// UnstableColumns names columns whose normalized values differ between two
// canonical states of the same fault-free trace: nondeterminism the catalog
// rule missed (an app-side clock or random value). Mark them in Options.Ignore.
func UnstableColumns(a, b *Canon) []string {
	var out []string
	keys := map[string]bool{}
	for k := range a.Cols {
		keys[k] = true
	}
	for k := range b.Cols {
		keys[k] = true
	}
	for k := range keys {
		ma, mb := a.Cols[k], b.Cols[k]
		same := len(ma) == len(mb)
		for v, n := range ma {
			if mb[v] != n {
				same = false
			}
		}
		if !same {
			out = append(out, k)
		}
	}
	sort.Strings(out)
	return out
}

func marshal(v any) string {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	enc.Encode(v)
	return strings.TrimSuffix(b.String(), "\n")
}

// Without returns the state minus the named tables, for tables that hold
// framework bookkeeping rather than application data (a checkpointer's
// tables, for example). Edges into a dropped table are dropped with it.
func (s *State) Without(tables map[string]bool) *State {
	if len(tables) == 0 {
		return s
	}
	out := &State{Rows: map[string][]Row{}, Opaque: s.Opaque}
	for _, t := range s.Tables {
		if tables[t.Name] {
			continue
		}
		cols := make([]Column, len(t.Columns))
		copy(cols, t.Columns)
		for i := range cols {
			if tables[cols[i].FKTable] {
				cols[i].FKTable, cols[i].FKColumn = "", ""
			}
		}
		out.Tables = append(out.Tables, Table{Name: t.Name, Columns: cols})
		out.Rows[t.Name] = s.Rows[t.Name]
	}
	return out
}
