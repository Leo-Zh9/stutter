package replay_test

import (
	"context"
	"encoding/json"
	"sync"
	"testing"
	"time"

	"github.com/Leo-Zh9/stutter/internal/replay"
	"github.com/Leo-Zh9/stutter/internal/testenv"
	"github.com/Leo-Zh9/stutter/internal/trace"
)

type arrivals struct {
	mu  sync.Mutex
	seq []int
}

func (a *arrivals) add(s int) { a.mu.Lock(); a.seq = append(a.seq, s); a.mu.Unlock() }
func (a *arrivals) count(s int) int {
	a.mu.Lock()
	defer a.mu.Unlock()
	n := 0
	for _, x := range a.seq {
		if x == s {
			n++
		}
	}
	return n
}

func firstOf(tr []trace.Entry, kind trace.Kind, path string) trace.Entry {
	for _, e := range tr {
		if e.Kind == kind && (path == "" || e.Path == path) {
			return e
		}
	}
	panic("not found")
}

// Lost sends the request, which completes, then retries it: two arrivals.
func TestLostFires(t *testing.T) {
	env := testenv.Env(t)
	cfg := testenv.Config(t, env, "")
	tr := testenv.Trace(t)
	var a arrivals
	cfg.Arrival = a.add
	e := firstOf(tr, trace.API, "/charges")
	res := replay.Run(context.Background(), cfg, tr, replay.Full(tr, replay.Fault{Type: replay.Lost, Seq: e.Seq}), 1)
	if res.Err != nil {
		t.Fatal(res.Err)
	}
	if got := a.count(e.Seq); got != 2 || res.Fired != 1 || res.NotFired != 0 {
		t.Fatalf("arrivals=%d fired=%d notfired=%d", got, res.Fired, res.NotFired)
	}
}

// A client may not retry an unkeyed API call, so faults on one never fire
// and it is sent exactly once.
func TestUnkeyedRequestsAreNeverFaulted(t *testing.T) {
	env := testenv.Env(t)
	cfg := testenv.Config(t, env, "")
	tr := testenv.Trace(t)
	e := firstOf(tr, trace.API, "/charges")
	delete(tr[e.Seq].Headers, "Idempotency-Key")
	var a arrivals
	cfg.Arrival = a.add
	for _, ft := range []replay.FaultType{replay.Lost, replay.Concurrent} {
		a.seq = nil
		res := replay.Run(context.Background(), cfg, tr, replay.Full(tr, replay.Fault{Type: ft, Seq: e.Seq}), 1)
		if res.Err != nil {
			t.Fatal(res.Err)
		}
		if a.count(e.Seq) != 1 || res.Fired != 0 || res.NotFired != 1 {
			t.Fatalf("%s on unkeyed call: arrivals=%d fired=%d", ft, a.count(e.Seq), res.Fired)
		}
	}
	for _, f := range replay.Faults(tr, "Idempotency-Key") {
		if f.Seq == e.Seq && f.Type != replay.Reorder {
			t.Fatalf("enumeration offered %v on an unkeyed call", f)
		}
	}
}

// Concurrent with the barrier: both attempts reach the database and wait on
// the lock before it is released (Fired counts only that case).
func TestConcurrentBarrierFires(t *testing.T) {
	env := testenv.Env(t)
	cfg := testenv.Config(t, env, "")
	tr := testenv.Trace(t)
	var a arrivals
	cfg.Arrival = a.add
	for _, e := range []trace.Entry{firstOf(tr, trace.API, "/charges"), firstOf(tr, trace.Webhook, "")} {
		a.seq = nil
		res := replay.Run(context.Background(), cfg, tr, replay.Full(tr, replay.Fault{Type: replay.Concurrent, Seq: e.Seq}), 1)
		if res.Err != nil {
			t.Fatal(res.Err)
		}
		if res.Fired != 1 || a.count(e.Seq) != 2 {
			t.Fatalf("seq %d: fired=%d arrivals=%d", e.Seq, res.Fired, a.count(e.Seq))
		}
	}
}

// A request whose handler never writes cannot fill the barrier. That must be
// reported as not fired, never as a pass.
func TestBarrierTimeoutIsNotFired(t *testing.T) {
	env := testenv.Env(t)
	cfg := testenv.Config(t, env, "")
	tr := testenv.Trace(t)
	hook := firstOf(tr, trace.Webhook, "")
	var ev map[string]any
	json.Unmarshal(hook.Body, &ev)
	ev["data"].(map[string]any)["charge"] = "ch_does_not_exist" // handler reads, finds nothing, never writes
	tr[hook.Seq].Body, _ = json.Marshal(ev)
	res := replay.Run(context.Background(), cfg, tr, replay.Full(tr, replay.Fault{Type: replay.Concurrent, Seq: hook.Seq}), 1)
	if res.Err != nil {
		t.Fatal(res.Err)
	}
	if res.Fired != 0 || res.NotFired != 1 {
		t.Fatalf("fired=%d notfired=%d", res.Fired, res.NotFired)
	}
}

// Reorder swaps two deliveries about the same object.
func TestReorderSwapsDelivery(t *testing.T) {
	env := testenv.Env(t)
	cfg := testenv.Config(t, env, "")
	tr := testenv.Trace(t)
	var f replay.Fault
	for _, x := range replay.Faults(tr, "Idempotency-Key") {
		if x.Type == replay.Reorder {
			f = x
			break
		}
	}
	if f.Type == "" {
		t.Fatal("trace has no reorderable pair")
	}
	var a arrivals
	cfg.Arrival = a.add
	res := replay.Run(context.Background(), cfg, tr, replay.Full(tr, f), 1)
	if res.Err != nil || res.Fired != 1 {
		t.Fatalf("err=%v fired=%d", res.Err, res.Fired)
	}
	pos := map[int]int{}
	for i, s := range a.seq {
		pos[s] = i
	}
	if pos[f.Seq2] > pos[f.Seq] {
		t.Fatalf("#%d still delivered before #%d", f.Seq, f.Seq2)
	}
}

// Fault-free replays of the same trace end in one canonical state.
func TestFaultFreeReplayIsDeterministic(t *testing.T) {
	env := testenv.Env(t)
	cfg := testenv.Config(t, env, "")
	tr := testenv.Trace(t)
	digests := map[string]bool{}
	for i := 0; i < 5; i++ {
		res := replay.Run(context.Background(), cfg, tr, replay.Full(tr), int64(i))
		if res.Err != nil {
			t.Fatal(res.Err)
		}
		digests[res.Canon.Digest()] = true
	}
	if len(digests) != 1 {
		t.Fatalf("%d distinct end states", len(digests))
	}
}

// A provider profile like Stripe's retries any non-2xx; the default client
// rule retries only 5xx and 409. A webhook for an unknown charge gets 404.
func TestProviderRetryProfile(t *testing.T) {
	env := testenv.Env(t)
	tr := testenv.Trace(t)
	hook := firstOf(tr, trace.Webhook, "")
	var ev map[string]any
	json.Unmarshal(hook.Body, &ev)
	ev["data"].(map[string]any)["charge"] = "ch_does_not_exist"
	tr[hook.Seq].Body, _ = json.Marshal(ev)
	for _, tc := range []struct {
		retryAny bool
		want     int
	}{{false, 1}, {true, 5}} {
		cfg := testenv.Config(t, env, "")
		cfg.Webhooks.RetryAny, cfg.Webhooks.MaxAttempts = tc.retryAny, 5
		var a arrivals
		cfg.Arrival = a.add
		if res := replay.Run(context.Background(), cfg, tr, replay.Full(tr), 1); res.Err != nil {
			t.Fatal(res.Err)
		}
		if got := a.count(hook.Seq); got != tc.want {
			t.Fatalf("retryAny=%v: %d deliveries, want %d", tc.retryAny, got, tc.want)
		}
	}
}

// Stripe's scheme: v1 is hex HMAC-SHA256 over "<t>.<body>" with the secret.
func TestStripeSignature(t *testing.T) {
	got := replay.StripeSignature("whsec_test", time.Unix(1700000000, 0), []byte(`{"id":"evt_1"}`))
	// Computed independently: printf '1700000000.{"id":"evt_1"}' | openssl dgst -sha256 -hmac whsec_test
	want := "t=1700000000,v1=" + stripeWant
	if got != want {
		t.Fatalf("got %s want %s", got, want)
	}
}
