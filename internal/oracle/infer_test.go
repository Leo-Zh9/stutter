package oracle

import (
	"encoding/json"
	"fmt"
	"testing"
)

func num(v int64) json.Number { return json.Number(fmt.Sprint(v)) }

// Two tables whose id ranges overlap, a burned value, an id echoed in JSON,
// and rows the template held before banding.
func inferWorld(txStart, logStart int64, burn bool) *State {
	tables := []Table{
		{Name: "tx", Columns: []Column{{Name: "id"}, {Name: "amount"}}},
		{Name: "logs", Columns: []Column{{Name: "id"}, {Name: "data"}, {Name: "hash"}}},
	}
	st := &State{Tables: tables, Rows: map[string][]Row{}}
	st.Rows["tx"] = append(st.Rows["tx"], Row{"id": num(1), "amount": num(5)}) // template row, outside bands
	t := txStart
	for i := 0; i < 3; i++ {
		if burn && i == 1 {
			t++ // a rolled-back duplicate consumed a value
		}
		st.Rows["tx"] = append(st.Rows["tx"], Row{"id": num(t), "amount": num(int64(100 * (i + 1)))})
		st.Rows["logs"] = append(st.Rows["logs"], Row{"id": num(logStart + int64(i)),
			"data": map[string]any{"transaction": map[string]any{"id": num(t)}, "reverts": fmt.Sprint(t)}, "hash": fmt.Sprintf("h%d-%d", t, i)})
		t++
	}
	return st
}

var testBands = []Band{
	{Seq: "s.log_id", Start: 1e12, Width: 1e12, Offset: 1e6},
	{Seq: "s.tx_id", Start: 2e12, Width: 1e12, Offset: 2e6},
}
var testOwners = map[string]string{"s.log_id": "logs.id", "s.tx_id": "tx.id"}

func TestInferAttributesIdsAndVolatiles(t *testing.T) {
	a := inferWorld(2e12, 1e12, false)
	b := inferWorld(2e12+2e6, 1e12+1e6, false)
	inf, err := Infer(a, b, testBands, testOwners)
	if err != nil {
		t.Fatal(err)
	}
	if inf.Derived["tx.id"] != 1 || inf.Derived["logs.id"] != 0 || inf.Derived["logs.data.transaction.id"] != 1 || inf.Derived["logs.data.reverts"] != 1 {
		t.Fatalf("derived: %v", inf.Derived)
	}
	if inf.Owners[1] != "tx.id" || inf.Owners[0] != "logs.id" {
		t.Fatalf("owners: %v", inf.Owners)
	}
	if len(inf.Volatile) != 1 || inf.Volatile[0] != "logs.hash" {
		t.Fatalf("volatile: %v", inf.Volatile)
	}
	if _, ok := inf.Derived["tx.amount"]; ok {
		t.Fatal("amount treated as an id")
	}
	// A run that burned a transaction id is equivalent; re-pointing the log
	// at a different transaction is not.
	opt := Options{Inf: inf}
	ca, _ := Canonicalize(a, opt)
	cb, _ := Canonicalize(inferWorld(5e12/2, 1e12+7, true), opt)
	if d := Diff(ca, cb); len(d) > 0 {
		t.Fatalf("burned id produced a diff: %v", d)
	}
	bad := inferWorld(2e12, 1e12, false)
	bad.Rows["logs"][0]["data"] = map[string]any{"transaction": map[string]any{"id": num(2e12 + 1)}, "reverts": "2000000000000"}
	cbad, _ := Canonicalize(bad, opt)
	if len(Diff(ca, cbad)) == 0 {
		t.Fatal("re-pointed reference not detected")
	}
}

func TestInferRejectsUnownedIds(t *testing.T) {
	a := inferWorld(2e12, 1e12, false)
	b := inferWorld(2e12+2e6, 1e12+1e6, false)
	if _, err := Infer(a, b, testBands, map[string]string{"s.log_id": "logs.id"}); err == nil {
		t.Fatal("sequence without an owner was not rejected")
	}
}

// An OWNED BY column whose in-band values do not move with its sequence
// cannot be trusted as the owner: inference must stop, not guess.
func TestInferRejectsOwnerThatDoesNotShift(t *testing.T) {
	a := inferWorld(2e12, 1e12, false)
	b := inferWorld(2e12+2e6, 1e12+1e6, false)
	for _, st := range []*State{a, b} {
		st.Tables[0].Columns = append(st.Tables[0].Columns, Column{Name: "legacy"})
		for _, row := range st.Rows["tx"] {
			row["legacy"] = num(2e12 + 5) // in the band, identical in both runs
		}
	}
	owners := map[string]string{"s.log_id": "logs.id", "s.tx_id": "tx.legacy"}
	if _, err := Infer(a, b, testBands, owners); err == nil {
		t.Fatal("owner column that does not shift was accepted")
	}
}
