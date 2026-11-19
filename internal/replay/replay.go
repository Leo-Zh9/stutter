// Package replay re-sends a recorded trace to a fresh copy of the service,
// acting as a well-behaved client and webhook provider, and injects faults
// into its own transport.
package replay

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
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

	"github.com/Leo-Zh9/stutter/internal/agentstub"
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
	// ConcurrentAt: the ack of commit K of request Seq is held while a
	// duplicate is sent; it fires only if the duplicate is answered while
	// the first attempt is still waiting (opt-in, see client.held).
	ConcurrentAt FaultType = "concurrent@commit"
)

// Why a ConcurrentAt fault did not fire.
const (
	WhyFinished = "finished before commit K"
	WhyGuard    = "resume guard"
	WhyNotHeld  = "first attempt was not held"
	WhyTimeout  = "no response within HoldTimeout"
)

type Fault struct {
	Type FaultType `json:"type"`
	Seq  int       `json:"seq"`
	Seq2 int       `json:"seq2,omitempty"`
	K    int       `json:"k,omitempty"` // crash, concurrent@commit: which commit of the request
}

func (f Fault) String() string {
	if f.Type == Reorder {
		return fmt.Sprintf("reorder(#%d,#%d)", f.Seq, f.Seq2)
	}
	if f.Type == Crash {
		return fmt.Sprintf("crash(#%d@commit%d)", f.Seq, f.K)
	}
	if f.Type == ConcurrentAt {
		return fmt.Sprintf("concurrent(#%d@commit%d)", f.Seq, f.K)
	}
	return fmt.Sprintf("%s(#%d)", f.Type, f.Seq)
}

// ParseFault reads the notation String writes: lost(#2), concurrent(#2),
// reorder(#3,#5), crash(#2@commit1), concurrent(#2@commit1).
func ParseFault(s string) (Fault, error) {
	var f Fault
	var a, b int
	switch {
	case strings.HasPrefix(s, "concurrent(#") && strings.Contains(s, "@commit"):
		_, err := fmt.Sscanf(s, "concurrent(#%d@commit%d)", &a, &b)
		return Fault{Type: ConcurrentAt, Seq: a, K: b}, err
	case strings.HasPrefix(s, "lost(#"):
		_, err := fmt.Sscanf(s, "lost(#%d)", &a)
		return Fault{Type: Lost, Seq: a}, err
	case strings.HasPrefix(s, "concurrent(#"):
		_, err := fmt.Sscanf(s, "concurrent(#%d)", &a)
		return Fault{Type: Concurrent, Seq: a}, err
	case strings.HasPrefix(s, "reorder(#"):
		_, err := fmt.Sscanf(s, "reorder(#%d,#%d)", &a, &b)
		return Fault{Type: Reorder, Seq: a, Seq2: b}, err
	case strings.HasPrefix(s, "crash(#"):
		_, err := fmt.Sscanf(s, "crash(#%d@commit%d)", &a, &b)
		return Fault{Type: Crash, Seq: a, K: b}, err
	}
	return f, fmt.Errorf("unknown fault %q", s)
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
	// HoldTimeout bounds how long a ConcurrentAt duplicate may take while
	// the first attempt is held; default 10s (a hunt derives it from the
	// slowest observed request).
	HoldTimeout time.Duration
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
	// commit points per request in Result.Commits, without crashing. It
	// also reports, for each resume request, whether an interrupt was
	// pending before and after it (Result.Pending).
	ObserveCommits bool
	// Rules widen which API calls the client may retry (see trace.Rule).
	Rules []trace.Rule
	// Headers are sent on every request: static values, never recorded.
	Headers map[string]string
	// Model, when set, is served by a scripted model stub for each run;
	// the service finds it through the {MODEL_URL} placeholder.
	Model *agentstub.Script
	// InterruptSQL names the interrupt pending on thread $1 (any text that
	// identifies it), or returns NULL or '' when none is pending. Default:
	// LangGraph's Postgres checkpointer.
	InterruptSQL string
}

// DefaultInterruptSQL reads LangGraph's Postgres checkpointer: a dynamic
// interrupt() leaves an __interrupt__ write on the thread's latest root
// checkpoint, and that checkpoint's id names the interrupt. It fails safe: a
// node that interrupts twice keeps its old __interrupt__ row after the first
// resume, so that resume still reads as pending, and its retry is refused.
const DefaultInterruptSQL = `SELECT max(c.checkpoint_id) FROM checkpoints c
	WHERE c.thread_id = $1 AND c.checkpoint_ns = ''
	  AND c.checkpoint_id = (SELECT max(c2.checkpoint_id) FROM checkpoints c2
	                         WHERE c2.thread_id = $1 AND c2.checkpoint_ns = '')
	  AND EXISTS (SELECT 1 FROM checkpoint_writes w
	              WHERE w.thread_id = $1 AND w.checkpoint_ns = '' AND w.checkpoint_id = c.checkpoint_id
	                AND w.channel = '__interrupt__')`

// Policy is the client's retry rights under this config.
func (c Config) Policy() trace.Policy {
	return trace.Policy{KeyHeader: c.KeyHeader, Rules: c.Rules}
}

// PendingObs is what an observation run saw around one resume request.
type PendingObs struct {
	Before bool `json:"before"`
	After  bool `json:"after"`
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
	Canon     *oracle.Canon
	State     *oracle.State
	Commits   map[int]int        // commit points per request (proxy runs only)
	BgCommits int                // commits seen while no request was in flight
	Pending   map[int]PendingObs // resume requests (observation runs only)
	Fired     int
	NotFired  int
	// Why lists, per ConcurrentAt fault that did not fire, the reason.
	Why []string
	// Slowest is the slowest single request (observation runs only).
	Slowest time.Duration
	// ModelCalls and ModelMisses count scripted model requests; a miss
	// means the conversation left the script. MissReason is the first.
	ModelCalls  int
	ModelMisses int
	MissReason  string
	Err         error
	Took        time.Duration
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
	if cfg.HoldTimeout == 0 {
		cfg.HoldTimeout = 10 * time.Second
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
	holds := false
	for i := range s.Faults {
		if s.Faults[i].Type == Crash {
			crash = &s.Faults[i]
		}
		holds = holds || s.Faults[i].Type == ConcurrentAt
	}
	dbFor := target.DB{DSN: cfg.Env.DSN(db), URL: cfg.Env.URL(db), SocketURL: cfg.Env.SocketURL(db)}
	var px *pgproxy.Proxy
	if crash != nil || holds || cfg.ObserveCommits {
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
	var stub *agentstub.Server
	if cfg.Model != nil {
		stub, err = agentstub.Start(cfg.Model)
		if err != nil {
			return Result{Err: err}
		}
		defer stub.Close()
		dbFor.Vars = map[string]string{"{MODEL_URL}": stub.URL}
	}
	proc, err := target.Start(ctx, cfg.Target, dbFor)
	if err != nil {
		return Result{Err: err}
	}
	defer func() { proc.Stop() }()
	policy := cfg.Policy()

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
	hold := map[int]int{} // seq -> commit whose ack is held
	var res Result
	for _, f := range s.Faults {
		switch f.Type {
		case Lost:
			if include[f.Seq] && policy.Retryable(bySeq[f.Seq]) {
				lost[f.Seq] = true
			} else {
				res.NotFired++
			}
		case Concurrent:
			if include[f.Seq] && policy.Retryable(bySeq[f.Seq]) {
				conc[f.Seq] = true
			} else {
				res.NotFired++
			}
		case ConcurrentAt:
			if include[f.Seq] && policy.Retryable(bySeq[f.Seq]) {
				hold[f.Seq] = f.K
			} else {
				res.NotFired++
			}
		case Crash:
			if !include[f.Seq] || !policy.Retryable(bySeq[f.Seq]) {
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
		cfg: cfg, base: proc.BaseURL,
		// A fresh connection per request: a service may close a kept-alive
		// connection after an error, and reusing it would turn the next
		// request into a transport error the service never caused.
		http: &http.Client{Timeout: 10 * time.Second, Transport: &http.Transport{DisableKeepAlives: true}},
		rng:  rand.New(rand.NewSource(seed)), subst: map[string]string{}, nums: map[int64]int64{}, dbName: db,
	}
	if len(conc) > 0 && cfg.Barrier {
		conn, err := pgx.Connect(ctx, cfg.Env.DSN(db))
		if err != nil {
			return Result{Err: err}
		}
		defer conn.Close(context.Background())
		c.barrier = conn
		// Tables the oracle ignores are left unlocked: a framework's own
		// bookkeeping (a checkpointer's writes) commits before the handler
		// reaches its check-then-act, so holding it would release both
		// attempts before the race.
		rows, err := conn.Query(ctx, `SELECT quote_ident(schemaname) || '.' || quote_ident(tablename),
			CASE WHEN schemaname = 'public' THEN tablename ELSE schemaname || '.' || tablename END
			FROM pg_tables WHERE schemaname = ANY($1) ORDER BY 1`, schemas(cfg))
		if err != nil {
			return Result{Err: err}
		}
		for rows.Next() {
			var t, label string
			rows.Scan(&t, &label)
			if !cfg.Oracle.IgnoreTables[label] {
				c.tables = append(c.tables, t)
			}
		}
		rows.Close()
	}

	c.interruptSQL = cfg.InterruptSQL
	if c.interruptSQL == "" {
		c.interruptSQL = DefaultInterruptSQL
	}
	if cfg.ObserveCommits {
		res.Pending = map[int]PendingObs{}
	}
	for _, q := range order {
		e := bySeq[q]
		var r response
		thread, isResume := policy.Resume(e)
		faulted := lost[q] || conc[q] || hold[q] > 0 || (crash != nil && crash.Seq == q)
		if isResume && (faulted || cfg.ObserveCommits) {
			c.resume = &resumeState{thread: thread, before: c.interrupt(thread)}
			if faulted && c.resume.before == "" {
				// Nothing to answer: a duplicate would not be a duplicate
				// of an approval. The fault does not fire.
				res.NotFired++
				if hold[q] > 0 {
					res.Why = append(res.Why, WhyGuard)
				}
				lost[q], conc[q] = false, false
				delete(hold, q)
				if crash != nil && crash.Seq == q {
					crash = nil
				}
				faulted = false
			}
		} else {
			c.resume = nil
		}
		if px != nil {
			px.Begin(q)
		}
		fired, why, t0 := false, "", time.Now()
		switch {
		case crash != nil && crash.Seq == q:
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
		case hold[q] > 0:
			r, fired, why, err = c.held(ctx, e, px, hold[q])
		case conc[q]:
			r, fired, err = c.concurrent(ctx, e)
		case lost[q]:
			c.once(ctx, e) // completes on the server; the sender never sees it
			fired = true
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
		if cfg.ObserveCommits {
			res.Slowest = max(res.Slowest, time.Since(t0))
		}
		if c.resume != nil && c.resume.refused {
			fired = false // the retry would have answered a different interrupt
			if hold[q] > 0 && why == "" {
				why = WhyGuard
			}
		}
		if faulted {
			if fired {
				res.Fired++
			} else {
				res.NotFired++
				if why != "" {
					res.Why = append(res.Why, why)
				}
			}
		}
		if cfg.ObserveCommits && isResume {
			after := c.interrupt(thread)
			res.Pending[q] = PendingObs{Before: c.resume.before != "", After: after != "" && after != c.resume.before}
		}
		c.resume = nil
		c.learn(e, r)
	}
	proc.Stop()
	if stub != nil {
		calls := stub.Calls()
		res.ModelCalls = len(calls)
		for _, cl := range calls {
			if cl.Miss != "" {
				if res.ModelMisses == 0 {
					res.MissReason = cl.Miss
				}
				res.ModelMisses++
			}
		}
	}
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
	st = st.Without(cfg.Oracle.IgnoreTables)
	if stub != nil {
		st.Opaque = stub.Minted()
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
	// resume is set while a resume request is in flight.
	resume       *resumeState
	interruptSQL string
}

// resumeState guards the retries of a request that answers an interrupt.
// A retry is sent only while no interrupt is pending (the first attempt
// finished the run) or the same one still is (it did not take effect).
// If a different interrupt is pending, the retry would approve something
// the user never saw: the client's mistake, so it is refused and the fault
// counts as not fired.
type resumeState struct {
	thread  string
	before  string
	refused bool
}

// interrupt names the interrupt pending on a thread, "" if none. Any error
// reads as an unknown interrupt, which refuses retries: fail safe.
func (c *client) interrupt(thread string) string {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, c.cfg.Env.DSN(c.dbName))
	if err != nil {
		return "?unreadable"
	}
	defer conn.Close(ctx)
	var id *string
	if err := conn.QueryRow(ctx, c.interruptSQL, c.subst1(thread)).Scan(&id); err != nil {
		return "?unreadable"
	}
	if id == nil {
		return ""
	}
	return *id
}

// mayRetry applies the resume guard before a retry.
func (c *client) mayRetry() bool {
	if c.resume == nil {
		return true
	}
	if now := c.interrupt(c.resume.thread); now != "" && now != c.resume.before {
		c.resume.refused = true
		return false
	}
	return true
}

// once sends a single attempt and reads the full response.
func (c *client) once(ctx context.Context, e trace.Entry) (response, error) {
	return c.onceWith(ctx, e, c.http)
}

func (c *client) onceWith(ctx context.Context, e trace.Entry, hc *http.Client) (response, error) {
	body := c.rewrite(e.Body)
	req, err := http.NewRequestWithContext(ctx, e.Method, c.base+c.rewritePath(e.Path), bytes.NewReader(body))
	if err != nil {
		return response{}, err
	}
	for k, v := range e.Headers {
		req.Header.Set(k, v)
	}
	for k, v := range c.cfg.Headers {
		req.Header.Set(k, v)
	}
	if e.Kind == trace.Webhook && c.cfg.Webhooks.Signing == "stripe" {
		req.Header.Set(c.cfg.Webhooks.Header, StripeSignature(c.cfg.Webhooks.Secret, time.Now(), body))
	}
	if c.cfg.Arrival != nil {
		c.cfg.Arrival(e.Seq)
	}
	resp, err := hc.Do(req)
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
	// A refused retry ends the request as the client last saw it (nothing,
	// after a lost response); the run goes on and the fault does not count.
	var last response
	var lastErr error
	for ; ; attempt++ {
		if attempt >= 2 && !c.mayRetry() {
			return last, lastErr
		}
		r, err := c.once(ctx, e)
		last, lastErr = r, err
		if c.done(e, r, err) || !c.cfg.Policy().Retryable(e) || attempt >= max {
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

// held sends request e with the proxy armed to hold back the ack of commit k
// on A's own connection, then sends a duplicate B while it is held. It fires
// only if the hold was reached, B got an HTTP response within HoldTimeout,
// and A had not answered by then: B ran its whole handler while A was parked
// right after commit k. Otherwise why says what happened. A sync database
// call inside an async handler freezes the event loop while held, so B
// cannot run, and that shows up as not fired, never as a pass. A is released
// in every case, B is always waited for, and A's response is the one the
// client keeps, as under the barrier.
func (c *client) held(ctx context.Context, e trace.Entry, px *pgproxy.Proxy, k int) (response, bool, string, error) {
	// A may be parked for HoldTimeout, so the default 10s client timeout
	// would cut it off and pose as a transport error.
	hc := &http.Client{Timeout: c.cfg.HoldTimeout + 20*time.Second, Transport: &http.Transport{DisableKeepAlives: true}}
	heldCh, release := px.ArmHold(e.Seq, k)
	defer release()
	type out struct {
		r   response
		err error
	}
	aCh, bCh := make(chan out, 1), make(chan out, 1)
	go func() { r, err := c.onceWith(ctx, e, hc); aCh <- out{r, err} }()
	var a out
	aDone, bWait, fired, why := false, false, false, ""
	select {
	case a = <-aCh:
		aDone, why = true, WhyFinished
	case <-heldCh:
		if !c.mayRetry() {
			why = WhyGuard // B would answer an interrupt A just raised
			break
		}
		bWait = true
		go func() { r, err := c.onceWith(ctx, e, hc); bCh <- out{r, err} }()
		timer := time.NewTimer(c.cfg.HoldTimeout)
		select {
		case a = <-aCh:
			aDone, why = true, WhyNotHeld
		case b := <-bCh:
			bWait, fired = false, b.err == nil
			if !fired {
				why = WhyTimeout
			}
		case <-timer.C:
			why = WhyTimeout
		}
		timer.Stop()
	}
	release()
	if !aDone {
		a = <-aCh
	}
	if bWait {
		<-bCh
	}
	if a.err == nil && c.done(e, a.r, nil) {
		return a.r, fired, why, nil
	}
	r, err := c.send(ctx, e, 2) // A's own retry policy continues
	return r, fired, why, err
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
		if y, ok := live.(string); ok && x != y && looksLikeID(x) && looksLikeID(y) {
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

// looksLikeID keeps prose out of the id map: an agent's reply text differs
// between runs too, and mapping it would rewrite unrelated requests. Ids are
// non-empty and contain no whitespace; timestamps are not ids.
func looksLikeID(s string) bool {
	return s != "" && !strings.ContainsAny(s, " \t\n\r") && !looksLikeTime(s)
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

// subst1 maps one recorded id to its live value.
func (c *client) subst1(s string) string {
	if y, ok := c.subst[s]; ok {
		return y
	}
	return s
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
	Rules     []trace.Rule
	// Pending is what an observation run saw around each resume request.
	// A resume admits concurrent only if an interrupt was pending before
	// it, and lost or crash only if, in addition, none was pending after
	// it: otherwise the retry would answer the next interrupt, which is
	// the client's mistake, not the service's. Unobserved resumes admit
	// nothing.
	Pending map[int]PendingObs
	// WebhookObject is a dotted JSON path naming what a webhook is about
	// (e.g. "data.object.customer"). Two webhooks with the same value may be
	// reordered. Empty: group by the first id minted earlier in the trace.
	WebhookObject string
}

// Faults lists every single fault the trace admits: Lost and Concurrent on
// each retryable request, Reorder on each pair of webhooks about the same
// object.
func Faults(tr []trace.Entry, sp Space) []Fault {
	var out []Fault
	for _, e := range tr {
		if sp.Admits(e, Lost) {
			out = append(out, Fault{Type: Lost, Seq: e.Seq})
		}
	}
	for _, e := range tr {
		if sp.Admits(e, Concurrent) {
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

// Admits reports whether fault type t may be placed on request e.
func (sp Space) Admits(e trace.Entry, t FaultType) bool {
	p := trace.Policy{KeyHeader: sp.KeyHeader, Rules: sp.Rules}
	if !p.Retryable(e) {
		return false
	}
	h := sp.KeyHeader
	if h == "" {
		h = trace.DefaultKeyHeader
	}
	if e.Retryable(h) {
		return true // keyed by header, or a webhook
	}
	if r := p.RuleFor(e); r != nil && r.Resume != nil {
		obs, ok := sp.Pending[e.Seq]
		if !ok || !obs.Before {
			return false
		}
		return t == Concurrent || t == ConcurrentAt || !obs.After
	}
	return true
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
