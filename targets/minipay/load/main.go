// minipay-load drives a seeded synthetic workload against minipay. Pointed at
// `stutter record`, it produces the trace that stutter replays. The traffic is
// generated, not production traffic.
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	target := flag.String("target", "http://127.0.0.1:8080", "base URL (the recorder's listen address)")
	seed := flag.Int64("seed", 1, "workload seed")
	customers := flag.Int("customers", 5, "customers to create")
	charges := flag.Int("charges", 40, "charges to create")
	webhooks := flag.Int("webhooks", 20, "webhook deliveries to send")
	flag.Parse()
	rng := rand.New(rand.NewSource(*seed))

	post := func(path, key string, body any) map[string]any {
		b, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST", *target+path, bytes.NewReader(b))
		req.Header.Set("Content-Type", "application/json")
		if key != "" {
			req.Header.Set("Idempotency-Key", key)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		raw, _ := io.ReadAll(resp.Body)
		var out map[string]any
		_ = json.Unmarshal(raw, &out)
		return out
	}

	var cus []string
	for i := 0; i < *customers; i++ {
		r := post("/customers", fmt.Sprintf("key_cus_%d", i), map[string]any{"email": fmt.Sprintf("buyer%d@example.test", i)})
		cus = append(cus, r["id"].(string))
	}
	type charge struct {
		id     string
		amount int64
	}
	var chs []charge
	var firstBody map[string]any
	for i := 0; i < *charges; i++ {
		amt := int64(500 + rng.Intn(9500))
		body := map[string]any{"customer": cus[rng.Intn(len(cus))], "amount": amt, "currency": "usd"}
		if i == 0 {
			firstBody = body
		}
		r := post("/charges", fmt.Sprintf("key_ch_%d", i), body)
		chs = append(chs, charge{r["id"].(string), amt})
	}
	// A client mistake seen in practice: reusing a key with a different body.
	// A correct server rejects it with 422. This request exists so the
	// key-unscoped control variant has something to get wrong.
	reuse := map[string]any{"customer": firstBody["customer"], "amount": firstBody["amount"].(int64) + 100, "currency": "usd"}
	post("/charges", "key_ch_0", reuse)

	// Payouts through the multi-phase key (three commits each).
	for i := 0; i < 8; i++ {
		post("/payouts", fmt.Sprintf("key_po_%d", i), map[string]any{"customer": cus[rng.Intn(len(cus))], "amount": int64(100 + rng.Intn(900))})
	}

	// Provider events, delivered in creation order during recording. Some
	// charges get succeeded then disputed (order matters), others get a
	// partial refund (applying it twice is the bug to find).
	base := time.Date(2026, 11, 1, 0, 0, 0, 0, time.UTC)
	n := 0
	send := func(typ string, c charge, amount int64) {
		ev := map[string]any{
			"id":      fmt.Sprintf("evt_%04d", n),
			"type":    typ,
			"created": base.Add(time.Duration(n) * time.Minute).Format(time.RFC3339),
			"data":    map[string]any{"charge": c.id, "amount": amount},
		}
		n++
		post("/webhooks", "", ev)
	}
	perm := rng.Perm(len(chs))
	i := 0
	for n < *webhooks && i < len(perm) {
		c := chs[perm[i]]
		switch i % 3 {
		case 0:
			send("charge.succeeded", c, 0)
			if n < *webhooks {
				send("charge.disputed", c, 0)
			}
		case 1:
			send("charge.succeeded", c, 0)
		case 2:
			send("charge.refunded", c, c.amount/4)
		}
		i++
	}
}
