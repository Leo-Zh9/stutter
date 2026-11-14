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
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/Leo-Zh9/stutter/internal/oracle"
	"github.com/Leo-Zh9/stutter/internal/pgenv"
	"github.com/Leo-Zh9/stutter/internal/pgproxy"
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
	// Crash: the service is killed right after commit K of request Seq,
	// before it hears the commit succeeded; it restarts and the sender
	// retries (see package pgproxy).
	Crash FaultType = "crash"
)

type Fault struct {
	Type FaultType `json:"type"`
	Seq  int       `json:"seq"`
	Seq2 int       `json:"seq2,omitempty"`
	K    int       `json:"k,omitempty"` // crash: which commit of the request
}

func (f Fault) String() string {
	if f.Type == Reorder {
		return fmt.Sprintf("reorder(#%d,#%d)", f.Seq, f.Seq2)
	}
	if f.Type == Crash {
		return fmt.Sprintf("crash(#%d@commit%d)", f.Seq, f.K)
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
	// Bands, when set, are the numeric ranges each sequence mints into.
	// Integers inside a band are mapped from recorded to live ids.
	Bands []pgenv.Band
	// Shift advances every banded sequence by its offset right after the
	// clone: the inference baseline (see oracle.Infer).
	Shift bool
	// KeyHeader names the idempotency header (default Idempotency-Key).
	KeyHeader string
	Webhooks  WebhookPolicy
	// Arrival, if set, is called each time a request reaches the transport.
	Arrival func(seq int)
	// ObserveCommits routes the service through the wire proxy and reports
	// commit points per request in Result.Commits, without crashing.
	ObserveCommits bool
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
	State    *oracle.State
	Commits  map[int]int // commit points per request (proxy runs only)
	BgCommits int        // commits seen while no request was in flight
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
	if cfg.Shift {
		if err := cfg.Env.ShiftBands(ctx, db, cfg.Bands); err != nil {
			return Result{Err: err}
		}
	}
	var crash *Fault
	for i := range s.Faults {
		if s.Faults[i].Type == Crash {
			crash = &s.Faults[i]
		}
	}
	dbFor := target.DB{DSN: cfg.Env.DSN(db), URL: cfg.Env.URL(db), SocketURL: cfg.Env.SocketURL(db)}
	var px *pgproxy.Proxy
	if crash != nil || cfg.ObserveCommits {
		px, err = pgproxy.New(cfg.Env.SocketDir, cfg.Env.Port, true)
		if err != nil {
			return Result{Err: err}
		}
		defer px.Close()
		dbFor = target.DB{
			DSN:       fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", px.Dir, px.Port, cfg.Env.User, db),
			URL:       fmt.Sprintf("postgres://%s@%s/%s?sslmode=disable", cfg.Env.User, px.TCPAddr, db),
			SocketURL: fmt.Sprintf("postgres://%s@/%s?host=%s&port=%s&sslmode=disable", cfg.Env.User, db, px.Dir, px.Port),
		}
	}
	proc, err := target.Start(ctx, cfg.Target, dbFor)
	if err != nil {
		return Result{Err: err}
	}
	defer func() { proc.Stop() }()

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
		case Crash:
			if !include[f.Seq] || !bySeq[f.Seq].Retryable(cfg.KeyHeader) {
				res.NotFired++
				crash = nil
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
		rng: rand.New(rand.NewSource(seed)), subst: map[string]string{}, nums: map[int64]int64{}, dbName: db,
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
		if px != nil {
			px.Begin(q)
		}
		switch {
		case crash != nil && crash.Seq == q:
			var fired bool
			r, fired, err = c.crashed(ctx, e, px, crash.K, func() error {
				proc.Stop()
				px.DropConns()
				px.Resume()
				np, err := target.Start(ctx, cfg.Target, dbFor)
				if err != nil {
					return err
				}
				proc, c.base = np, np.BaseURL
				return nil
			})
			if fired {
				res.Fired++
			} else {
				res.NotFired++
			}
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
		if px != nil {
			px.End()
		}
		if err != nil {
			return Result{Err: fmt.Errorf("seq %d: %w", q, err)}
		}
		c.learn(e, r)
	}
	proc.Stop()
	if px != nil {
		res.Commits, res.BgCommits = px.Commits()
	}

	conn, err := pgx.Connect(ctx, cfg.Env.DSN(db))
	if err != nil {
		return Result{Err: err}
	}
	defer conn.Close(context.Background())
	st, err := oracle.Snapshot(ctx, conn, schemas(cfg))
	if err != nil {
		return Result{Err: err}
	}
	res.State = st
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
	nums    map[int64]int64 // recorded -> live integer ids, in-band only
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

// crashed sends request e with the proxy armed to hold back commit k. If
// the crash point is reached, the service is killed and restarted, and the
// request is retried exactly as after a lost response. If the request
// finishes first, the fault did not fire (the request has fewer commits).
func (c *client) crashed(ctx context.Context, e trace.Entry, px *pgproxy.Proxy, k int, restart func() error) (response, bool, error) {
	px.ArmCrash(e.Seq, k)
	type out struct {
		r   response
		err error
	}
	ch := make(chan out, 1)
	go func() { r, err := c.once(ctx, e); ch <- out{r, err} }()
	select {
	case o := <-ch:
		px.Resume()
		if o.err == nil && c.done(e, o.r, nil) {
			return o.r, false, nil
		}
		r, err := c.send(ctx, e, 2)
		return r, false, err
	case <-px.Crashed():
		if err := restart(); err != nil {
			return response{}, true, err
		}
		<-ch // the killed attempt's connection error
		r, err := c.send(ctx, e, 2)
		return r, true, err
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
	rec, err1 := decodeNum(e.RespBody)
	live, err2 := decodeNum(r.body)
	if err1 != nil || err2 != nil {
		return
	}
	c.pair(rec, live)
}

func decodeNum(b []byte) (any, error) {
	var v any
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()
	err := dec.Decode(&v)
	return v, err
}

func (c *client) inBand(v int64) bool {
	for _, b := range c.cfg.Bands {
		if b.Contains(v) {
			return true
		}
	}
	return false
}

func (c *client) pair(rec, live any) {
	switch x := rec.(type) {
	case string:
		if y, ok := live.(string); ok && x != y && !looksLikeTime(x) {
			c.subst[x] = y
		}
	case json.Number:
		y, ok := live.(json.Number)
		if !ok {
			return
		}
		a, e1 := x.Int64()
		b, e2 := y.Int64()
		if e1 == nil && e2 == nil && a != b && c.inBand(a) {
			c.nums[a] = b
		}
	case map[string]any:
		if y, ok := live.(map[string]any); ok {
			for k, v := range x {
				c.pair(v, y[k])
			}
		}
	case []any:
		if y, ok := live.([]any); ok {
			for i := range x {
				if i < len(y) {
					c.pair(x[i], y[i])
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
	if len(body) == 0 || (len(c.subst) == 0 && len(c.nums) == 0) {
		return body
	}
	var v any
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.UseNumber()
	if dec.Decode(&v) != nil {
		return body
	}
	v = c.substAll(v)
	b, _ := json.Marshal(v)
	return b
}

func (c *client) substAll(v any) any {
	switch x := v.(type) {
	case string:
		if y, ok := c.subst[x]; ok {
			return y
		}
		return x
	case json.Number:
		if n, err := x.Int64(); err == nil {
			if y, ok := c.nums[n]; ok {
				return json.Number(strconv.FormatInt(y, 10))
			}
		}
		return x
	case map[string]any:
		for k, vv := range x {
			x[k] = c.substAll(vv)
		}
		return x
	case []any:
		for i := range x {
			x[i] = c.substAll(x[i])
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
		} else if n, err := strconv.ParseInt(s, 10, 64); err == nil {
			if y, ok := c.nums[n]; ok {
				parts[i] = strconv.FormatInt(y, 10)
			}
		}
	}
	return strings.Join(parts, "/")
}

// Space says which faults a trace admits.
type Space struct {
	KeyHeader string // API calls without it are never faulted
	// WebhookObject is a dotted JSON path naming what a webhook is about
	// (e.g. "data.object.customer"). Two webhooks with the same value may be
	// reordered. Empty: group by the first id minted earlier in the trace.
	WebhookObject string
}

// Faults lists every single fault the trace admits: Lost and Concurrent on
// each retryable request, Reorder on each pair of webhooks about the same
// object.
func Faults(tr []trace.Entry, sp Space) []Fault {
	keyHeader := sp.KeyHeader
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
		obj := ""
		if sp.WebhookObject != "" {
			obj, _ = jsonPath(v, sp.WebhookObject).(string)
		} else {
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
				obj = ids[0]
			}
		}
		if obj != "" {
			objOf[e.Seq] = obj
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

// jsonPath follows a dotted path through decoded JSON objects.
func jsonPath(v any, path string) any {
	for _, k := range strings.Split(path, ".") {
		m, ok := v.(map[string]any)
		if !ok {
			return nil
		}
		v = m[k]
	}
	return v
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
