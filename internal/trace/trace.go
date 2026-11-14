// Package trace defines the recorded request log and the recording proxy.
package trace

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"sync"
)

// Kind separates calls a client makes (api) from deliveries a provider makes
// to the service (webhook). They have different retry rights.
type Kind string

const (
	API     Kind = "api"
	Webhook Kind = "webhook"
)

// Entry is one recorded request and the response the service gave.
type Entry struct {
	Seq        int               `json:"seq"`
	Kind       Kind              `json:"kind"`
	Method     string            `json:"method"`
	Path       string            `json:"path"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       json.RawMessage   `json:"body,omitempty"`
	Status     int               `json:"status"`
	RespBody   json.RawMessage   `json:"resp_body,omitempty"`
}

// Header names copied into the trace. Credentials are never recorded.
var keptHeaders = []string{"Content-Type", "Idempotency-Key"}

// Key returns the idempotency key, if the request carried one.
func (e Entry) Key() string { return e.Headers["Idempotency-Key"] }

// Retryable reports whether a well-behaved sender may resend this request:
// a provider may always redeliver a webhook, and a client may retry an API
// call only when it carries an idempotency key.
func (e Entry) Retryable() bool {
	return e.Kind == Webhook || e.Key() != ""
}

func Read(path string) ([]Entry, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var out []Entry
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 1<<20), 1<<24)
	for sc.Scan() {
		if len(bytes.TrimSpace(sc.Bytes())) == 0 {
			continue
		}
		var e Entry
		if err := json.Unmarshal(sc.Bytes(), &e); err != nil {
			return nil, err
		}
		out = append(out, e)
	}
	return out, sc.Err()
}

func Write(path string, entries []Entry) error {
	var buf bytes.Buffer
	for _, e := range entries {
		b, err := json.Marshal(e)
		if err != nil {
			return err
		}
		buf.Write(b)
		buf.WriteByte('\n')
	}
	return os.WriteFile(path, buf.Bytes(), 0o644)
}

// Recorder is a reverse proxy that forwards to an upstream service and logs
// every request/response pair as an Entry.
type Recorder struct {
	proxy         *httputil.ReverseProxy
	webhookPrefix string
	mu            sync.Mutex
	entries       []Entry
}

func NewRecorder(upstream *url.URL, webhookPrefix string) *Recorder {
	r := &Recorder{proxy: httputil.NewSingleHostReverseProxy(upstream), webhookPrefix: webhookPrefix}
	return r
}

func (r *Recorder) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	body, _ := io.ReadAll(req.Body)
	req.Body = io.NopCloser(bytes.NewReader(body))
	rec := &capture{ResponseWriter: w, status: 200}
	r.proxy.ServeHTTP(rec, req)

	e := Entry{Kind: API, Method: req.Method, Path: req.URL.RequestURI(), Headers: map[string]string{}, Status: rec.status}
	if r.webhookPrefix != "" && strings.HasPrefix(req.URL.Path, r.webhookPrefix) {
		e.Kind = Webhook
	}
	for _, h := range keptHeaders {
		if v := req.Header.Get(h); v != "" {
			e.Headers[h] = v
		}
	}
	e.Body = asJSON(body)
	e.RespBody = asJSON(rec.buf.Bytes())
	r.mu.Lock()
	e.Seq = len(r.entries)
	r.entries = append(r.entries, e)
	r.mu.Unlock()
}

func (r *Recorder) Entries() []Entry {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]Entry(nil), r.entries...)
}

// asJSON keeps a body as raw JSON when it is JSON, else as a JSON string.
func asJSON(b []byte) json.RawMessage {
	b = bytes.TrimSpace(b)
	if len(b) == 0 {
		return nil
	}
	if json.Valid(b) {
		return json.RawMessage(append([]byte(nil), b...))
	}
	s, _ := json.Marshal(string(b))
	return s
}

type capture struct {
	http.ResponseWriter
	status int
	buf    bytes.Buffer
}

func (c *capture) WriteHeader(code int) { c.status = code; c.ResponseWriter.WriteHeader(code) }
func (c *capture) Write(b []byte) (int, error) {
	c.buf.Write(b)
	return c.ResponseWriter.Write(b)
}
