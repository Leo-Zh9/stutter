// minipay is a small payments API used as stutter's measurement target.
//
// It is correct by default. Setting BUG=<name> switches on exactly one
// seeded bug, each a documented real-world idempotency or webhook mistake:
//
//	key-ignored               idempotency key never checked; a retry redoes the charge
//	key-after-commit          key saved in its own transaction after the charge commits
//	webhook-no-dedupe         webhook handler keeps no event log; redelivery re-applies it
//	webhook-stale-overwrite   status webhook applied without comparing event time
//	webhook-check-then-insert event dedupe is SELECT-then-INSERT with no unique index
//	key-unscoped              key not bound to the request body (control: no fault needed)
//
// Run `minipay -schema` to print the schema for the selected variant.
package main

import (
	"context"
	"crypto/sha256"
	_ "embed"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed schema.sql
var schemaSQL string

var Bugs = []string{
	"key-ignored",
	"key-after-commit",
	"webhook-no-dedupe",
	"webhook-stale-overwrite",
	"webhook-check-then-insert",
	"key-unscoped",
}

func variantSchema(bug string) string {
	if bug != "webhook-check-then-insert" {
		return schemaSQL
	}
	var out []string
	skip := false
	for _, line := range strings.Split(schemaSQL, "\n") {
		if strings.Contains(line, "@unique-event-id") {
			skip = true
			continue
		}
		if skip {
			skip = false
			continue
		}
		out = append(out, line)
	}
	return strings.Join(out, "\n")
}

type server struct {
	db  *pgxpool.Pool
	bug string
}

func main() {
	printSchema := flag.Bool("schema", false, "print the schema for the BUG variant and exit")
	flag.Parse()
	bug := os.Getenv("BUG")
	if bug != "" && !contains(Bugs, bug) {
		log.Fatalf("unknown BUG %q", bug)
	}
	if *printSchema {
		fmt.Print(variantSchema(bug))
		return
	}
	cfg, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	cfg.MaxConns = 8
	db, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		log.Fatal(err)
	}
	s := &server{db: db, bug: bug}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(204) })
	mux.HandleFunc("POST /customers", s.createCustomer)
	mux.HandleFunc("POST /charges", s.createCharge)
	mux.HandleFunc("POST /webhooks", s.webhook)
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = "127.0.0.1:8080"
	}
	log.Fatal(http.ListenAndServe(addr, mux))
}

type reply struct {
	code int
	body any
}

func writeJSON(w http.ResponseWriter, code int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(body)
}

func errReply(code int, msg string) reply {
	return reply{code, map[string]string{"error": msg}}
}

// idempotent runs fn exactly once per Idempotency-Key, in one transaction with
// the key row, and replays the stored response for any later attempt. The key
// row is inserted first, so a concurrent duplicate blocks on the primary key
// until the first attempt commits, then sees its stored response.
func (s *server) idempotent(ctx context.Context, key string, body []byte, fn func(pgx.Tx) reply) reply {
	sum := sha256.Sum256(body)
	hash := hex.EncodeToString(sum[:])
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return errReply(500, err.Error())
	}
	defer tx.Rollback(ctx)
	tag, err := tx.Exec(ctx, `INSERT INTO idempotency_keys (key, request_hash, response_code, response_body)
		VALUES ($1, $2, 0, '{}') ON CONFLICT (key) DO NOTHING`, key, hash)
	if err != nil {
		return errReply(500, err.Error())
	}
	if tag.RowsAffected() == 0 {
		tx.Rollback(ctx)
		return s.storedReply(ctx, key, hash)
	}
	rep := fn(tx)
	if rep.code >= 400 {
		return rep // rolled back: nothing stored, a retry recomputes the same error
	}
	b, _ := json.Marshal(rep.body)
	if _, err := tx.Exec(ctx, `UPDATE idempotency_keys SET response_code = $2, response_body = $3 WHERE key = $1`,
		key, rep.code, b); err != nil {
		return errReply(500, err.Error())
	}
	if err := tx.Commit(ctx); err != nil {
		return errReply(500, err.Error())
	}
	return rep
}

func (s *server) storedReply(ctx context.Context, key, hash string) reply {
	var storedHash string
	var code int
	var body json.RawMessage
	err := s.db.QueryRow(ctx, `SELECT request_hash, response_code, response_body FROM idempotency_keys WHERE key = $1`, key).
		Scan(&storedHash, &code, &body)
	if err != nil {
		return errReply(500, err.Error())
	}
	if storedHash != hash && s.bug != "key-unscoped" {
		return errReply(422, "idempotency key reused with a different request body")
	}
	if code == 0 {
		return errReply(409, "a request with this idempotency key is in progress")
	}
	return reply{code, body}
}

func (s *server) createCustomer(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Email string `json:"email"`
	}
	raw, ok := readBody(w, r, &in)
	if !ok {
		return
	}
	key := r.Header.Get("Idempotency-Key")
	if key == "" {
		writeJSON(w, 400, map[string]string{"error": "Idempotency-Key required"})
		return
	}
	rep := s.idempotent(r.Context(), key, raw, func(tx pgx.Tx) reply {
		var id string
		var created time.Time
		if err := tx.QueryRow(r.Context(), `INSERT INTO customers (email) VALUES ($1) RETURNING id, created_at`, in.Email).
			Scan(&id, &created); err != nil {
			return errReply(500, err.Error())
		}
		return reply{200, map[string]any{"id": id, "object": "customer", "email": in.Email, "created": created.UTC().Format(time.RFC3339Nano)}}
	})
	writeJSON(w, rep.code, rep.body)
}

type chargeIn struct {
	Customer string `json:"customer"`
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

func insertCharge(ctx context.Context, tx pgx.Tx, in chargeIn, key string) reply {
	tag, err := tx.Exec(ctx, `UPDATE customers SET balance_cents = balance_cents + $2 WHERE id = $1`, in.Customer, in.Amount)
	if err != nil {
		return errReply(500, err.Error())
	}
	if tag.RowsAffected() == 0 {
		return errReply(404, "no such customer")
	}
	var id string
	var created time.Time
	if err := tx.QueryRow(ctx, `INSERT INTO charges (customer_id, amount_cents, currency, idempotency_key)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at`, in.Customer, in.Amount, in.Currency, key).Scan(&id, &created); err != nil {
		return errReply(500, err.Error())
	}
	return reply{200, map[string]any{
		"id": id, "object": "charge", "customer": in.Customer, "amount": in.Amount,
		"currency": in.Currency, "status": "pending", "created": created.UTC().Format(time.RFC3339Nano),
	}}
}

func (s *server) createCharge(w http.ResponseWriter, r *http.Request) {
	var in chargeIn
	raw, ok := readBody(w, r, &in)
	if !ok {
		return
	}
	if in.Amount <= 0 || in.Currency == "" {
		writeJSON(w, 400, map[string]string{"error": "amount and currency required"})
		return
	}
	key := r.Header.Get("Idempotency-Key")
	if key == "" {
		writeJSON(w, 400, map[string]string{"error": "Idempotency-Key required"})
		return
	}
	ctx := r.Context()
	var rep reply
	switch s.bug {
	case "key-ignored":
		rep = s.inTx(ctx, func(tx pgx.Tx) reply { return insertCharge(ctx, tx, in, key) })
	case "key-after-commit":
		// Check the key, do the work, then record the key in a second
		// transaction. A duplicate that arrives before the record exists
		// passes the check and charges again.
		sum := sha256.Sum256(raw)
		hash := hex.EncodeToString(sum[:])
		var exists bool
		if err := s.db.QueryRow(ctx, `SELECT EXISTS (SELECT 1 FROM idempotency_keys WHERE key = $1)`, key).Scan(&exists); err != nil {
			rep = errReply(500, err.Error())
			break
		}
		if exists {
			rep = s.storedReply(ctx, key, hash)
			break
		}
		rep = s.inTx(ctx, func(tx pgx.Tx) reply { return insertCharge(ctx, tx, in, key) })
		if rep.code < 400 {
			b, _ := json.Marshal(rep.body)
			if _, err := s.db.Exec(ctx, `INSERT INTO idempotency_keys (key, request_hash, response_code, response_body)
				VALUES ($1, $2, $3, $4) ON CONFLICT (key) DO NOTHING`, key, hash, rep.code, b); err != nil {
				rep = errReply(500, err.Error())
			}
		}
	default:
		rep = s.idempotent(ctx, key, raw, func(tx pgx.Tx) reply { return insertCharge(ctx, tx, in, key) })
	}
	writeJSON(w, rep.code, rep.body)
}

func (s *server) inTx(ctx context.Context, fn func(pgx.Tx) reply) reply {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return errReply(500, err.Error())
	}
	defer tx.Rollback(ctx)
	rep := fn(tx)
	if rep.code >= 400 {
		return rep
	}
	if err := tx.Commit(ctx); err != nil {
		return errReply(500, err.Error())
	}
	return rep
}

type event struct {
	ID      string    `json:"id"`
	Type    string    `json:"type"`
	Created time.Time `json:"created"`
	Data    struct {
		Charge string `json:"charge"`
		Amount int64  `json:"amount"`
	} `json:"data"`
}

func (s *server) webhook(w http.ResponseWriter, r *http.Request) {
	var ev event
	raw, ok := readBody(w, r, &ev)
	if !ok {
		return
	}
	ctx := r.Context()
	rep := s.inTx(ctx, func(tx pgx.Tx) reply {
		var exists bool
		if err := tx.QueryRow(ctx, `SELECT EXISTS (SELECT 1 FROM charges WHERE id = $1)`, ev.Data.Charge).Scan(&exists); err != nil {
			return errReply(500, err.Error())
		}
		if !exists {
			return errReply(404, "no such charge")
		}
		switch s.bug {
		case "webhook-no-dedupe":
			// No event log at all: every delivery is applied.
		case "webhook-check-then-insert":
			var seen bool
			if err := tx.QueryRow(ctx, `SELECT EXISTS (SELECT 1 FROM webhook_events WHERE event_id = $1)`, ev.ID).Scan(&seen); err != nil {
				return errReply(500, err.Error())
			}
			if seen {
				return reply{200, map[string]any{"received": true, "duplicate": true}}
			}
			if _, err := tx.Exec(ctx, `INSERT INTO webhook_events (event_id, type, charge_id, payload) VALUES ($1, $2, $3, $4)`,
				ev.ID, ev.Type, ev.Data.Charge, raw); err != nil {
				return errReply(500, err.Error())
			}
		default:
			tag, err := tx.Exec(ctx, `INSERT INTO webhook_events (event_id, type, charge_id, payload) VALUES ($1, $2, $3, $4)
				ON CONFLICT (event_id) DO NOTHING`, ev.ID, ev.Type, ev.Data.Charge, raw)
			if err != nil {
				return errReply(500, err.Error())
			}
			if tag.RowsAffected() == 0 {
				return reply{200, map[string]any{"received": true, "duplicate": true}}
			}
		}
		if err := s.apply(ctx, tx, ev); err != nil {
			return errReply(500, err.Error())
		}
		return reply{200, map[string]any{"received": true}}
	})
	writeJSON(w, rep.code, rep.body)
}

func (s *server) apply(ctx context.Context, tx pgx.Tx, ev event) error {
	switch ev.Type {
	case "charge.succeeded", "charge.disputed", "charge.failed":
		status := strings.TrimPrefix(ev.Type, "charge.")
		q := `UPDATE charges SET status = $2, status_at = $3 WHERE id = $1 AND (status_at IS NULL OR status_at < $3)`
		if s.bug == "webhook-stale-overwrite" {
			q = `UPDATE charges SET status = $2, status_at = $3 WHERE id = $1`
		}
		_, err := tx.Exec(ctx, q, ev.Data.Charge, status, ev.Created)
		return err
	case "charge.refunded":
		var customer string
		if err := tx.QueryRow(ctx, `UPDATE charges SET refunded_cents = refunded_cents + $2 WHERE id = $1 RETURNING customer_id`,
			ev.Data.Charge, ev.Data.Amount).Scan(&customer); err != nil {
			return err
		}
		_, err := tx.Exec(ctx, `UPDATE customers SET balance_cents = balance_cents - $2 WHERE id = $1`, customer, ev.Data.Amount)
		return err
	}
	return nil
}

func readBody(w http.ResponseWriter, r *http.Request, v any) ([]byte, bool) {
	defer r.Body.Close()
	raw, err := io.ReadAll(r.Body)
	if err != nil {
		writeJSON(w, 400, map[string]string{"error": err.Error()})
		return nil, false
	}
	if err := json.Unmarshal(raw, v); err != nil {
		writeJSON(w, 400, map[string]string{"error": "invalid JSON"})
		return nil, false
	}
	return raw, true
}

func contains(xs []string, x string) bool {
	for _, y := range xs {
		if y == x {
			return true
		}
	}
	return false
}
