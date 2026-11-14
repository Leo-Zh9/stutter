package oracle

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
)

// A miniature of minipay's catalog, built by hand so the canonicalizer can
// be tested without a database.
func schema() []Table {
	return []Table{
		{Name: "customers", Columns: []Column{
			{Name: "id", Generated: true}, {Name: "email"}, {Name: "balance"},
			{Name: "created_at", Temporal: true, Generated: true},
		}},
		{Name: "charges", Columns: []Column{
			{Name: "id", Generated: true}, {Name: "customer_id", FKTable: "customers", FKColumn: "id"},
			{Name: "amount"}, {Name: "status"},
		}},
		{Name: "idem", Columns: []Column{{Name: "key"}, {Name: "response"}}},
	}
}

type world struct {
	customers [][2]any // email, balance
	charges   [][3]any // customer index, amount, status
}

// build materializes a world with fresh random ids, a random insert order and
// random timestamps, the way two runs of the same correct program would.
func build(w world, rng *rand.Rand) *State {
	st := &State{Tables: schema(), Rows: map[string][]Row{}}
	ids := make([]string, len(w.customers))
	for i := range ids {
		ids[i] = fmt.Sprintf("cus_%08x", rng.Uint32())
	}
	for _, i := range rng.Perm(len(w.customers)) {
		c := w.customers[i]
		st.Rows["customers"] = append(st.Rows["customers"], Row{
			"id": ids[i], "email": c[0], "balance": c[1], "created_at": fmt.Sprintf("2026-11-%02dT00:00:00Z", 1+rng.Intn(28)),
		})
	}
	for _, i := range rng.Perm(len(w.charges)) {
		ch := w.charges[i]
		id := fmt.Sprintf("ch_%08x", rng.Uint32())
		st.Rows["charges"] = append(st.Rows["charges"], Row{"id": id, "customer_id": ids[ch[0].(int)], "amount": ch[1], "status": ch[2]})
		// A cached response that embeds the generated id and a timestamp.
		st.Rows["idem"] = append(st.Rows["idem"], Row{"key": fmt.Sprintf("k%d", i), "response": map[string]any{
			"id": id, "customer": ids[ch[0].(int)], "created": fmt.Sprintf("2026-11-14T09:%02d:00Z", rng.Intn(60)),
		}})
	}
	return st
}

func canon(t *testing.T, st *State) *Canon {
	t.Helper()
	c, err := Canonicalize(st, Options{})
	if err != nil {
		t.Fatal(err)
	}
	return c
}

var base = world{
	customers: [][2]any{{"a@x", 100}, {"b@x", 300}, {"a@x", 100}}, // two customers identical except id
	charges: [][3]any{
		{0, 100, "pending"}, {1, 300, "succeeded"}, {1, 300, "succeeded"}, // tied charges
		{2, 50, "disputed"},
	},
}

// Shuffled insert order, different uuids and timestamps: no diff.
func TestEquivalentStatesAreEqual(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	want := canon(t, build(base, rng))
	for i := 0; i < 200; i++ {
		if d := Diff(want, canon(t, build(base, rng))); len(d) > 0 {
			t.Fatalf("equivalent states differ: %v", d)
		}
	}
}

// Two parents identical except for their key, each with a different child.
// Whatever order they are inserted in, the states are equal; moving a child
// to the other parent is a real change and must show.
func TestIsomorphicParents(t *testing.T) {
	w := world{customers: [][2]any{{"a@x", 0}, {"a@x", 0}}, charges: [][3]any{{0, 100, "p"}, {1, 200, "p"}}}
	moved := world{customers: w.customers, charges: [][3]any{{0, 100, "p"}, {0, 200, "p"}}}
	rng := rand.New(rand.NewSource(2))
	a := canon(t, build(w, rng))
	for i := 0; i < 100; i++ {
		if d := Diff(a, canon(t, build(w, rng))); len(d) > 0 {
			t.Fatalf("insert order changed the canonical form: %v", d)
		}
	}
	if d := Diff(a, canon(t, build(moved, rng))); len(d) == 0 {
		t.Fatal("moving a charge between customers was not detected")
	}
}

func TestRealChangesAreDetected(t *testing.T) {
	rng := rand.New(rand.NewSource(3))
	a := canon(t, build(base, rng))
	changed := base
	changed.charges = append([][3]any{{0, 101, "pending"}}, base.charges[1:]...)
	if len(Diff(a, canon(t, build(changed, rng)))) == 0 {
		t.Fatal("changed amount not detected")
	}
	extra := base
	extra.charges = append(append([][3]any{}, base.charges...), [3]any{1, 300, "succeeded"})
	if len(Diff(a, canon(t, build(extra, rng)))) == 0 {
		t.Fatal("duplicate charge not detected")
	}
	bal := base
	bal.customers = [][2]any{{"a@x", 100}, {"b@x", 600}, {"a@x", 100}}
	if len(Diff(a, canon(t, build(bal, rng)))) == 0 {
		t.Fatal("changed balance not detected")
	}
}

// A correct handler that rolled back a duplicate still burned a uuid and an
// identity value; the surviving ids differ from the baseline's, and a cached
// response embeds them. None of that is a diff.
func TestBurnedIdsAreNotADiff(t *testing.T) {
	tables := []Table{
		{Name: "orders", Columns: []Column{{Name: "id", Generated: true}, {Name: "sku"}}},
		{Name: "lines", Columns: []Column{{Name: "id", Generated: true}, {Name: "order_id", FKTable: "orders", FKColumn: "id"}, {Name: "qty"}}},
		{Name: "idem", Columns: []Column{{Name: "key"}, {Name: "response"}}},
	}
	mk := func(orderIDs []int64, lineIDs []int64, uuid string) *State {
		return &State{Tables: tables, Rows: map[string][]Row{
			"orders": {{"id": json.Number(fmt.Sprint(orderIDs[0])), "sku": "a"}, {"id": json.Number(fmt.Sprint(orderIDs[1])), "sku": "b"}},
			"lines": {
				{"id": json.Number(fmt.Sprint(lineIDs[0])), "order_id": json.Number(fmt.Sprint(orderIDs[0])), "qty": 1},
				{"id": json.Number(fmt.Sprint(lineIDs[1])), "order_id": json.Number(fmt.Sprint(orderIDs[1])), "qty": 2},
			},
			"idem": {{"key": "k1", "response": map[string]any{"order": uuid}}},
		}}
	}
	// orders.id is an identity column here; the uuid lives in a separate table.
	tables = append(tables, Table{Name: "tokens", Columns: []Column{{Name: "id", Generated: true}, {Name: "label"}}})
	a := mk([]int64{1, 2}, []int64{1, 2}, "tok_aaaa")
	a.Rows["tokens"] = []Row{{"id": "tok_aaaa", "label": "x"}}
	b := mk([]int64{2, 3}, []int64{3, 4}, "tok_bbbb") // sequences advanced by a rolled-back duplicate
	b.Rows["tokens"] = []Row{{"id": "tok_bbbb", "label": "x"}}
	ca, cb := canon(t, a), canon(t, b)
	if d := Diff(ca, cb); len(d) > 0 {
		t.Fatalf("burned ids produced a diff: %v", d)
	}
	// But pointing the cached response at the wrong token is a real change.
	c := mk([]int64{1, 2}, []int64{1, 2}, "tok_other")
	c.Rows["tokens"] = []Row{{"id": "tok_aaaa", "label": "x"}, {"id": "tok_other", "label": "y"}}
	if len(Diff(ca, canon(t, c))) == 0 {
		t.Fatal("re-pointed reference not detected")
	}
}

func TestReferenceCycleIsAnError(t *testing.T) {
	st := &State{
		Tables: []Table{{Name: "n", Columns: []Column{{Name: "id", Generated: true}, {Name: "next"}}}},
		Rows:   map[string][]Row{"n": {{"id": "n_1", "next": "n_2"}, {"id": "n_2", "next": "n_1"}}},
	}
	if _, err := Canonicalize(st, Options{}); err == nil {
		t.Fatal("expected a cycle error")
	}
}

func TestUnstableColumnsAreNamed(t *testing.T) {
	tables := []Table{{Name: "t", Columns: []Column{{Name: "id", Generated: true}, {Name: "nonce"}, {Name: "v"}}}}
	a := &State{Tables: tables, Rows: map[string][]Row{"t": {{"id": "x_1", "nonce": "r1", "v": 1}}}}
	b := &State{Tables: tables, Rows: map[string][]Row{"t": {{"id": "x_2", "nonce": "r9", "v": 1}}}}
	got := UnstableColumns(canon(t, a), canon(t, b))
	if len(got) != 1 || got[0] != "t.nonce" {
		t.Fatalf("got %v", got)
	}
	ca, _ := Canonicalize(a, Options{Ignore: map[string]bool{"t.nonce": true}})
	cb, _ := Canonicalize(b, Options{Ignore: map[string]bool{"t.nonce": true}})
	if len(Diff(ca, cb)) != 0 {
		t.Fatal("ignored column still diffed")
	}
}
