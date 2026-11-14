// Package replay re-sends a recorded trace to a fresh copy of the service,
// acting as a well-behaved client and webhook provider, and injects faults
// into its own transport.
package replay

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/Leo-Zh9/stutter/internal/oracle"
	"github.com/Leo-Zh9/stutter/internal/pgenv"
	"github.com/Leo-Zh9/stutter/internal/target"
	"github.com/Leo-Zh9/stutter/internal/trace"
)

type FaultType string

const (
	// Lost: the request reaches the service and completes, but the sender
	// never sees the response (a timeout after commit), so it retries.
	Lost FaultType = "lost"
	// Concurrent: a duplicate is sent while the first attempt is still
	// inside its transaction.
	Concurrent FaultType = "concurrent"
	// Reorder: two webhook deliveries about the same object arrive swapped.
	Reorder FaultType = "reorder"
)

type Fault struct {
	Type FaultType `json:"type"`
	Seq  int       `json:"seq"`
	Seq2 int       `json:"seq2,omitempty"`
}

func (f Fault) String() string {
	if f.Type == Reorder {
		return fmt.Sprintf("reorder(#%d,#%d)", f.Seq, f.Seq2)
	}
	return fmt.Sprintf("%s(#%d)", f.Type, f.Seq)
}

// Schedule is a subset of the trace (by seq, in trace order) plus faults.
type Schedule struct {
	Seqs   []int   `json:"seqs"`
	Faults []Fault `json:"faults"`
}

func (s Schedule) HasConcurrent() bool {
	for _, f := range s.Faults {
		if f.Type == Concurrent {
			return true
		}
	}
	return false
}

func Full(tr []trace.Entry, faults ...Fault) Schedule {
	s := Schedule{Faults: faults}
	for _, e := range tr {
		s.Seqs = append(s.Seqs, e.Seq)
	}
	return s
}

type Config struct {
	Env            *pgenv.Env
	Template       string
	Target         target.Spec
	Barrier        bool          // make Concurrent deterministic with the lock barrier
	BarrierTimeout time.Duration // default 500ms
	MaxAttempts    int           // per request, default 3
	Oracle         oracle.Options
	// KeyHeader names the idempotency header (default Idempotency-Key).
	KeyHeader string
	Webhooks  WebhookPolicy
	// Arrival, if set, is called each time a request reaches the transport.
	Arrival func(seq int)
}

// WebhookPolicy describes the provider that sends webhook entries.
type WebhookPolicy struct {
	Signing     string // "stripe": sign each delivery at send time
	Secret      string
	Header      string // default Stripe-Signature
	RetryAny    bool   // provider retries any non-2xx (Stripe), not only 5xx
	MaxAttempts int    // provider attempts per delivery; default MaxAttempts
}

type Result struct {
	Canon    *oracle.Canon
	Fired    int
	NotFired int
	Err      error
	Took     time.Duration
}

// Run replays one schedule on a fresh database clone and returns the
// canonical end state.
func Run(ctx context.Context, cfg Config, tr []trace.Entry, s Schedule, seed int64) Result {
	start := time.Now()
	res := run(ctx, cfg, tr, s, seed)
	res.Took = time.Since(start)
	return res
}

func run(ctx context.Context, cfg Config, tr []trace.Entry, s Schedule, seed int64) Result {
	if cfg.BarrierTimeout == 0 {
		cfg.BarrierTimeout = 500 * time.Millisecond
	}
	if cfg.MaxAttempts == 0 {
		cfg.MaxAttempts = 3
	}
	if cfg.KeyHeader == "" {
		cfg.KeyHeader = trace.DefaultKeyHeader
	}
	if cfg.Webhooks.MaxAttempts == 0 {
		cfg.Webhooks.MaxAttempts = cfg.MaxAttempts
	}
	if cfg.Webhooks.Header == "" {
		cfg.Webhooks.Header = "Stripe-Signature"
	}
	db, err := cfg.Env.Clone(ctx, cfg.Template)
	if err != nil {
		return Result{Err: err}
	}
	defer cfg.Env.Drop(context.Background(), db)
	proc, err := target.Start(ctx, cfg.Target, target.DB{DSN: cfg.Env.DSN(db), URL: cfg.Env.URL(db)})
	if err != nil {
		return Result{Err: err}
	}
	defer proc.Stop()

	bySeq := map[int]trace.Entry{}
	for _, e := range tr {
		bySeq[e.Seq] = e
	}
	include := map[int]bool{}
	for _, q := range s.Seqs {
		include[q] = true
	}
	order := append([]int(nil), s.Seqs...)
	pos := map[int]int{}
	for i, q := range order {
		pos[q] = i
	}
	lost := map[int]bool{}
	conc := map[int]bool{}
	var res Result
	for _, f := range s.Faults {
		switch f.Type {
		case Lost:
			if include[f.Seq] && bySeq[f.Seq].Retryable(cfg.KeyHeader) {
				lost[f.Seq] = true
			} else {
				res.NotFired++
			}
		case Concurrent:
			if include[f.Seq] && bySeq[f.Seq].Retryable(cfg.KeyHeader) {
				conc[f.Seq] = true
			} else {
				res.NotFired++
			}
		case Reorder:
			i, ok1 := pos[f.Seq]
			j, ok2 := pos[f.Seq2]
			if ok1 && ok2 {
				order[i], order[j] = order[j], order[i]
				pos[f.Seq], pos[f.Seq2] = j, i
				res.Fired++
			} else {
				res.NotFired++
			}
		}
	}

	c := &client{
		cfg: cfg, base: proc.BaseURL, http: &http.Client{Timeout: 10 * time.Second},
		rng: rand.New(rand.NewSource(seed)), subst: map[string]string{}, dbName: db,
	}
	if len(conc) > 0 && cfg.Barrier {
		conn, err := pgx.Connect(ctx, cfg.Env.DSN(db))
		if err != nil {
			return Result{Err: err}
		}
		defer conn.Close(context.Background())
		c.barrier = conn
		rows, err := conn.Query(ctx, `SELECT quote_ident(schemaname) || '.' || quote_ident(tablename)
			FROM pg_tables WHERE schemaname = ANY($1) ORDER BY 1`, schemas(cfg))
		if err != nil {
			return Result{Err: err}
		}
		for rows.Next() {
			var t string
			rows.Scan(&t)
			c.tables = append(c.tables, t)
		}
		rows.Close()
	}

	for _, q := range order {
		e := bySeq[q]
		var r response
		switch {
		case conc[q]:
			var fired bool
			r, fired, err = c.concurrent(ctx, e)
			if fired {
				res.Fired++
			} else {
				res.NotFired++
			}
		case lost[q]:
			c.once(ctx, e) // completes on the server; the sender never sees it
			res.Fired++
			r, err = c.send(ctx, e, 2)
		default:
			r, err = c.send(ctx, e, 1)
		}
		if err != nil {
			return Result{Err: fmt.Errorf("seq %d: %w", q, err)}
		}
		c.learn(e, r)
	}
	proc.Stop()

	conn, err := pgx.Connect(ctx, cfg.Env.DSN(db))
	if err != nil {
		return Result{Err: err}
	}
	defer conn.Close(context.Background())
	st, err := oracle.Snapshot(ctx, conn, schemas(cfg))
	if err != nil {
		return Result{Err: err}
	}
	res.Canon, res.Err = oracle.Canonicalize(st, cfg.Oracle)
	return res
}

func schemas(cfg Config) []string {
	if len(cfg.Oracle.Schemas) == 0 {
		return []string{"public"}
	}
	return cfg.Oracle.Schemas
}

type response struct {
	status int
	body   []byte
}

type client struct {
	cfg     Config
	base    string
	http    *http.Client
	rng     *rand.Rand
	rngMu   sync.Mutex
	subst   map[string]string
	barrier *pgx.Conn
	tables  []string
	dbName  string
}

// once sends a single attempt and reads the full response.
func (c *client) once(ctx context.Context, e trace.Entry) (response, error) {
	body := c.rewrite(e.Body)
	req, err := http.NewRequestWithContext(ctx, e.Method, c.base+c.rewritePath(e.Path), bytes.NewReader(body))
	if err != nil {
		return response{}, err
	}
	for k, v := range e.Headers {
		req.Header.Set(k, v)
	}
	if e.Kind == trace.Webhook && c.cfg.Webhooks.Signing == "stripe" {
		req.Header.Set(c.cfg.Webhooks.Header, StripeSignature(c.cfg.Webhooks.Secret, time.Now(), body))
	}
	if c.cfg.Arrival != nil {
		c.cfg.Arrival(e.Seq)
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return response{}, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	return response{resp.StatusCode, b}, err
}

// send is the client's retry policy: same request and same idempotency key
// on every attempt, retried only on a transport error, a 5xx, or 409 (key in
// progress), and only if the request is retryable at all. Unkeyed API calls
// are sent exactly once.
func (c *client) send(ctx context.Context, e trace.Entry, attempt int) (response, error) {
	max := c.cfg.MaxAttempts
	if e.Kind == trace.Webhook {
		max = c.cfg.Webhooks.MaxAttempts
	}
	for ; ; attempt++ {
		r, err := c.once(ctx, e)
		if c.done(e, r, err) || !e.Retryable(c.cfg.KeyHeader) || attempt >= max {
			return r, err
		}
		c.rngMu.Lock()
		d := time.Duration(1+c.rng.Intn(3)) * time.Millisecond * time.Duration(attempt)
		c.rngMu.Unlock()
		time.Sleep(d)
	}
}

// done reports whether a response ends the sender's retries. A client stops
// on anything but a 5xx or 409; a provider like Stripe retries any non-2xx.
func (c *client) done(e trace.Entry, r response, err error) bool {
	if err != nil {
		return false
	}
	if e.Kind == trace.Webhook && c.cfg.Webhooks.RetryAny {
		return r.status >= 200 && r.status < 300
	}
	return r.status < 500 && r.status != 409
}

// StripeSignature builds a Stripe-Signature header value: t=<unix>,v1=<hex
// HMAC-SHA256 of "<t>.<body>">. Stripe rejects signatures older than 300 s,
// so stutter signs at send time instead of replaying recorded headers.
func StripeSignature(secret string, now time.Time, body []byte) string {
	t := now.Unix()
	m := hmac.New(sha256.New, []byte(secret))
	fmt.Fprintf(m, "%d.", t)
	m.Write(body)
	return fmt.Sprintf("t=%d,v1=%s", t, hex.EncodeToString(m.Sum(nil)))
}

// concurrent sends attempt A, then a duplicate B while A is still inside its
// transaction. With the barrier on, the replayer holds EXCLUSIVE locks on
// every table: plain SELECTs proceed, but the first write of each attempt
// waits. Once both backends are waiting, the locks are released, so both
// attempts have done their reads before either writes. If both never wait
// within the timeout, the fault did not fire and the run must not count as
// a pass. A's response is the one the client keeps.
func (c *client) concurrent(ctx context.Context, e trace.Entry) (response, bool, error) {
	var tx pgx.Tx
	if c.barrier != nil {
		var err error
		tx, err = c.barrier.Begin(ctx)
		if err != nil {
			return response{}, false, err
		}
		if _, err := tx.Exec(ctx, "LOCK TABLE "+strings.Join(c.tables, ", ")+" IN EXCLUSIVE MODE"); err != nil {
			tx.Rollback(ctx)
			return response{}, false, err
		}
	}
	type out struct {
		r   response
		err error
	}
	aCh, bCh := make(chan out, 1), make(chan out, 1)
	go func() { r, err := c.once(ctx, e); aCh <- out{r, err} }()
	fired := true
	if tx != nil {
		deadline := time.Now().Add(c.cfg.BarrierTimeout)
		fired = c.waitWaiters(ctx, 1, deadline)
		go func() { r, err := c.once(ctx, e); bCh <- out{r, err} }()
		fired = fired && c.waitWaiters(ctx, 2, deadline)
		if err := tx.Commit(ctx); err != nil {
			return response{}, false, err
		}
	} else {
		go func() { r, err := c.once(ctx, e); bCh <- out{r, err} }()
	}
	a, b := <-aCh, <-bCh
	_ = b
	if a.err != nil {
		return response{}, fired, a.err
	}
	if !c.done(e, a.r, nil) {
		r, err := c.send(ctx, e, 2) // A's own retry policy continues
		return r, fired, err
	}
	return a.r, fired, nil
}

func (c *client) waitWaiters(ctx context.Context, n int, deadline time.Time) bool {
	for time.Now().Before(deadline) {
		var waiting int
		err := c.cfg.Env.Admin().QueryRow(ctx, `SELECT count(*) FROM pg_stat_activity
			WHERE datname = $1 AND wait_event_type = 'Lock' AND backend_type = 'client backend'`, c.dbName).Scan(&waiting)
		if err == nil && waiting >= n {
			return true
		}
		time.Sleep(500 * time.Microsecond)
	}
	return false
}

// learn maps ids from the recorded response onto the live response at the
// same JSON path, so later requests that reused a recorded id get the live
// one. Only exact string values are mapped; timestamps are skipped.
func (c *client) learn(e trace.Entry, r response) {
	if r.status != e.Status || len(e.RespBody) == 0 {
		return
	}
	var rec, live any
	if json.Unmarshal(e.RespBody, &rec) != nil || json.Unmarshal(r.body, &live) != nil {
		return
	}
	pair(rec, live, c.subst)
}

func pair(rec, live any, m map[string]string) {
	switch x := rec.(type) {
	case string:
		if y, ok := live.(string); ok && x != y && !looksLikeTime(x) {
			m[x] = y
		}
	case map[string]any:
		if y, ok := live.(map[string]any); ok {
			for k, v := range x {
				pair(v, y[k], m)
			}
		}
	case []any:
		if y, ok := live.([]any); ok {
			for i := range x {
				if i < len(y) {
					pair(x[i], y[i], m)
				}
			}
		}
	}
}

func looksLikeTime(s string) bool {
	_, err := time.Parse(time.RFC3339Nano, s)
	return err == nil
}

func (c *client) rewrite(body json.RawMessage) []byte {
	if len(body) == 0 || len(c.subst) == 0 {
		return body
	}
	var v any
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.UseNumber()
	if dec.Decode(&v) != nil {
		return body
	}
	v = subst(v, c.subst)
	b, _ := json.Marshal(v)
	return b
}

func subst(v any, m map[string]string) any {
	switch x := v.(type) {
	case string:
		if y, ok := m[x]; ok {
			return y
		}
		return x
	case map[string]any:
		for k, vv := range x {
			x[k] = subst(vv, m)
		}
		return x
	case []any:
		for i := range x {
			x[i] = subst(x[i], m)
		}
		return x
	}
	return v
}

func (c *client) rewritePath(p string) string {
	parts := strings.Split(p, "/")
	for i, s := range parts {
		if y, ok := c.subst[s]; ok {
			parts[i] = y
		}
	}
	return strings.Join(parts, "/")
}

// Faults lists every single fault the trace admits: Lost and Concurrent on
// each retryable request, Reorder on each pair of webhooks about the same
// object (the first string in data that a response minted).
func Faults(tr []trace.Entry, keyHeader string) []Fault {
	if keyHeader == "" {
		keyHeader = trace.DefaultKeyHeader
	}
	var out []Fault
	for _, e := range tr {
		if e.Retryable(keyHeader) {
			out = append(out, Fault{Type: Lost, Seq: e.Seq})
		}
	}
	for _, e := range tr {
		if e.Retryable(keyHeader) {
			out = append(out, Fault{Type: Concurrent, Seq: e.Seq})
		}
	}
	objOf := map[int]string{}
	minted := map[string]bool{}
	for _, e := range tr {
		var v any
		if json.Unmarshal(e.RespBody, &v) == nil {
			collectStrings(v, minted)
		}
	}
	var hooks []trace.Entry
	for _, e := range tr {
		if e.Kind != trace.Webhook {
			continue
		}
		var v any
		if json.Unmarshal(e.Body, &v) != nil {
			continue
		}
		refs := map[string]bool{}
		collectStrings(v, refs)
		var ids []string
		for s := range refs {
			if minted[s] && !looksLikeTime(s) {
				ids = append(ids, s)
			}
		}
		sort.Strings(ids)
		if len(ids) > 0 {
			objOf[e.Seq] = ids[0]
			hooks = append(hooks, e)
		}
	}
	for i := 0; i < len(hooks); i++ {
		for j := i + 1; j < len(hooks); j++ {
			if objOf[hooks[i].Seq] == objOf[hooks[j].Seq] {
				out = append(out, Fault{Type: Reorder, Seq: hooks[i].Seq, Seq2: hooks[j].Seq})
			}
		}
	}
	return out
}

func collectStrings(v any, out map[string]bool) {
	switch x := v.(type) {
	case string:
		out[x] = true
	case map[string]any:
		for _, vv := range x {
			collectStrings(vv, out)
		}
	case []any:
		for _, vv := range x {
			collectStrings(vv, out)
		}
	}
}
