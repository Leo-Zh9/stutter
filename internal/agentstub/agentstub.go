// Package agentstub is a scripted, OpenAI-compatible model server. stutter
// starts one per replay and points the service under test at it, so an
// agent's runs are deterministic and cost nothing.
//
// The model is a pure function of the conversation it is sent. It finds the
// last user message, picks the script turn whose "user" text that message
// contains, and answers with step k of that turn, where k is the number of
// assistant messages after the user message. A node that is re-run after a
// crash or resume sends the same conversation, so it gets the same answer,
// down to the tool-call ids: exactly what makes a re-run node repeat its tool
// call.
//
//	{"turns": [
//	  {"user": "refund order 42",
//	   "steps": [{"tool_calls": [{"name": "issue_refund", "args": {"order_id": 42}}]},
//	             {"content": "Refund issued."}]}
//	]}
//
// A tool argument may be {"$ref": "tool:N:dotted.path"}, the value at that
// path in the JSON content of the Nth tool result (0-based, or "last") after
// the user message. A conversation no turn matches, a step past the end of a
// turn, or a call to a tool the request did not offer is a script miss: the
// server answers 500 and counts it. A miss in a fault-free run means the
// script is wrong; in a faulted run it means the run diverged.
package agentstub

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Script struct {
	Turns []Turn `json:"turns"`
	// Side answers model calls that are not steps of the conversation: a
	// router's structured choice, a title generator, a summariser. They are
	// checked in order before the turns, and do not advance the step count
	// (their output must not be appended to the conversation as an
	// assistant message, or the script must count it as a step instead).
	Side []SideRule `json:"side,omitempty"`
	// Dims is the length of embedding vectors (default 1536).
	Dims int `json:"embedding_dims,omitempty"`
}

type Turn struct {
	User  string `json:"user"`
	Steps []Step `json:"steps"`
}

type Step struct {
	Content   string     `json:"content,omitempty"`
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`
}

// SideRule matches a request by what it offers or says, not by its step.
type SideRule struct {
	When    When   `json:"when"`
	Content string `json:"content"`
}

// When is a conjunction; empty fields match anything. At least one is set.
type When struct {
	SystemContains string `json:"system_contains,omitempty"` // a system message contains this
	NoTools        bool   `json:"no_tools,omitempty"`        // the request offers no tools
	ResponseFormat bool   `json:"response_format,omitempty"` // the request asks for structured output
	Offers         string `json:"offers,omitempty"`          // the request offers this tool
}

type ToolCall struct {
	Name string         `json:"name"`
	Args map[string]any `json:"args"`
}

// Load reads a script file.
func Load(path string) (*Script, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var s Script
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&s); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	if err := s.Validate(); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return &s, nil
}

// Validate checks the script's shape.
func (s *Script) Validate() error {
	for i, t := range s.Turns {
		if t.User == "" || len(t.Steps) == 0 {
			return fmt.Errorf("turn %d needs user text and at least one step", i)
		}
		if last := t.Steps[len(t.Steps)-1]; len(last.ToolCalls) > 0 || last.Content == "" {
			return fmt.Errorf("turn %d: the last step must be a plain content reply", i)
		}
	}
	for i, r := range s.Side {
		if r.When == (When{}) {
			return fmt.Errorf("side rule %d matches everything; give it a when", i)
		}
	}
	return nil
}

// Call is one model request the stub answered (or missed).
type Call struct {
	Turn  int    `json:"turn"` // -1 on a miss
	Step  int    `json:"step"`
	Tools []string `json:"tools,omitempty"` // tool calls issued
	Miss  string `json:"miss,omitempty"`
}

// Server serves one script on a local port.
type Server struct {
	URL    string // base URL to hand the service, ending in /v1
	script *Script
	l      net.Listener
	srv    *http.Server
	mu     sync.Mutex
	calls  []Call
	// seen counts how often each conversation was answered in this replay.
	// A real model mints fresh ids every call, so a node that is re-run with
	// the same prompt must get new tool-call ids; only the n-th answer to a
	// given prompt is stable across replays.
	seen map[string]int
}

// Start listens on 127.0.0.1 and serves the script.
func Start(s *Script) (*Server, error) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, err
	}
	st := &Server{URL: "http://" + l.Addr().String() + "/v1", script: s, l: l, seen: map[string]int{}}
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/chat/completions", st.chat)
	mux.HandleFunc("/chat/completions", st.chat)
	mux.HandleFunc("/v1/embeddings", st.embeddings)
	mux.HandleFunc("/embeddings", st.embeddings)
	mux.HandleFunc("/v1/models", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"object":"list","data":[{"id":"stutter-scripted","object":"model","owned_by":"stutter"}]}`)
	})
	st.srv = &http.Server{Handler: mux}
	go st.srv.Serve(l)
	return st, nil
}

// Close stops the server.
func (s *Server) Close() { s.srv.Close() }

// Calls returns every request answered so far, in arrival order.
func (s *Server) Calls() []Call {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]Call(nil), s.calls...)
}

// Misses counts script misses so far.
func (s *Server) Misses() int {
	n := 0
	for _, c := range s.Calls() {
		if c.Miss != "" {
			n++
		}
	}
	return n
}

func (s *Server) record(c Call) {
	s.mu.Lock()
	s.calls = append(s.calls, c)
	s.mu.Unlock()
}

type message struct {
	Role       string          `json:"role"`
	Content    json.RawMessage `json:"content"`
	ToolCalls  []any           `json:"tool_calls,omitempty"`
	ToolCallID string          `json:"tool_call_id,omitempty"`
}

type chatRequest struct {
	Model         string    `json:"model"`
	Messages      []message `json:"messages"`
	Tools         []struct {
		Function struct {
			Name string `json:"name"`
		} `json:"function"`
	} `json:"tools"`
	Stream         bool            `json:"stream"`
	ResponseFormat json.RawMessage `json:"response_format"`
	StreamOptions *struct {
		IncludeUsage bool `json:"include_usage"`
	} `json:"stream_options"`
}

// text flattens a message's content: a string, or an array of parts whose
// text parts are joined.
func text(raw json.RawMessage) string {
	if len(raw) == 0 || string(raw) == "null" {
		return ""
	}
	var s string
	if json.Unmarshal(raw, &s) == nil {
		return s
	}
	var parts []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	}
	if json.Unmarshal(raw, &parts) == nil {
		var b strings.Builder
		for _, p := range parts {
			if p.Text != "" {
				if b.Len() > 0 {
					b.WriteByte('\n')
				}
				b.WriteString(p.Text)
			}
		}
		return b.String()
	}
	return string(raw)
}

// Answer is what the model says to one conversation.
type Answer struct {
	Key        string // hash of the conversation and what answered it
	Turn, Step int    // Turn is -1 for a side rule; Step is then its index
	Content    string
	ToolCalls  []ResolvedCall
}

type ResolvedCall struct {
	ID, Name, Args string
}

// Request is what Decide looks at.
type Request struct {
	Messages       []message
	Offered        map[string]bool
	ResponseFormat bool
}

// Decide computes the scripted answer for a request, or a miss reason. Tool
// call ids are left for the server to assign.
func (s *Script) Decide(req Request) (Answer, string) {
	msgs := req.Messages
	conv, _ := json.Marshal(msgs)
	for i, r := range s.Side {
		if r.When.matches(req) {
			h := sha256.Sum256(append([]byte(fmt.Sprintf("side%d|", i)), conv...))
			return Answer{Key: hex.EncodeToString(h[:12]), Turn: -1, Step: i, Content: r.Content}, ""
		}
	}
	last := -1
	for i, m := range msgs {
		if m.Role == "user" {
			last = i
		}
	}
	if last < 0 {
		return Answer{}, "no user message"
	}
	user := text(msgs[last].Content)
	turn := s.match(user)
	if turn < 0 {
		return Answer{}, fmt.Sprintf("no turn matches user message %q", clip(user))
	}
	steps := s.Turns[turn].Steps
	h := sha256.Sum256(conv)
	key := hex.EncodeToString(h[:12])
	// The same request again after a completed run of it: a model that can
	// see the action already happened says so instead of acting twice. The
	// missing dedupe still shows in the transcript, but the script never
	// manufactures a second side effect.
	if s.completedBefore(msgs[:last], turn) {
		return Answer{Key: key, Turn: turn, Step: len(steps) - 1, Content: steps[len(steps)-1].Content}, ""
	}
	step := 0
	var tools []string
	for _, m := range msgs[last+1:] {
		switch m.Role {
		case "assistant":
			step++
		case "tool":
			tools = append(tools, text(m.Content))
		}
	}
	if step >= len(steps) {
		return Answer{}, fmt.Sprintf("turn %d has %d steps; conversation is at step %d", turn, len(steps), step)
	}
	st := steps[step]
	a := Answer{Key: key, Turn: turn, Step: step, Content: st.Content}
	for _, tc := range st.ToolCalls {
		if len(req.Offered) > 0 && !req.Offered[tc.Name] {
			return Answer{}, fmt.Sprintf("turn %d step %d calls %q, which the request does not offer", turn, step, tc.Name)
		}
		args, err := resolve(tc.Args, tools)
		if err != nil {
			return Answer{}, fmt.Sprintf("turn %d step %d: %v", turn, step, err)
		}
		b, _ := json.Marshal(args)
		a.ToolCalls = append(a.ToolCalls, ResolvedCall{Name: tc.Name, Args: string(b)})
	}
	return a, ""
}

func (s *Script) match(user string) int {
	for i, t := range s.Turns {
		if strings.Contains(user, t.User) {
			return i
		}
	}
	return -1
}

// completedBefore reports whether an earlier user message of the same turn
// was followed by every step of that turn.
func (s *Script) completedBefore(msgs []message, turn int) bool {
	n := len(s.Turns[turn].Steps)
	for i := 0; i < len(msgs); i++ {
		if msgs[i].Role != "user" || s.match(text(msgs[i].Content)) != turn {
			continue
		}
		answered := 0
		for _, m := range msgs[i+1:] {
			if m.Role == "user" {
				break
			}
			if m.Role == "assistant" {
				answered++
			}
		}
		if answered >= n {
			return true
		}
	}
	return false
}

func (w When) matches(r Request) bool {
	if w.NoTools && len(r.Offered) > 0 {
		return false
	}
	if w.ResponseFormat && !r.ResponseFormat {
		return false
	}
	if w.Offers != "" && !r.Offered[w.Offers] {
		return false
	}
	if w.SystemContains != "" {
		found := false
		for _, m := range r.Messages {
			if m.Role == "system" || m.Role == "developer" {
				if strings.Contains(text(m.Content), w.SystemContains) {
					found = true
				}
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func clip(s string) string {
	if len(s) > 120 {
		return s[:120] + "..."
	}
	return s
}

func resolve(v any, tools []string) (any, error) {
	switch x := v.(type) {
	case map[string]any:
		if ref, ok := x["$ref"].(string); ok && len(x) == 1 {
			return deref(ref, tools)
		}
		out := map[string]any{}
		for k, vv := range x {
			r, err := resolve(vv, tools)
			if err != nil {
				return nil, err
			}
			out[k] = r
		}
		return out, nil
	case []any:
		out := make([]any, len(x))
		for i := range x {
			r, err := resolve(x[i], tools)
			if err != nil {
				return nil, err
			}
			out[i] = r
		}
		return out, nil
	}
	return v, nil
}

// deref resolves "tool:N:dotted.path" against the tool results so far.
func deref(ref string, tools []string) (any, error) {
	parts := strings.SplitN(ref, ":", 3)
	if len(parts) < 2 || parts[0] != "tool" {
		return nil, fmt.Errorf("bad $ref %q", ref)
	}
	n := len(tools) - 1
	if parts[1] != "last" {
		var err error
		if n, err = strconv.Atoi(parts[1]); err != nil {
			return nil, fmt.Errorf("bad $ref %q", ref)
		}
	}
	if n < 0 || n >= len(tools) {
		return nil, fmt.Errorf("$ref %q: only %d tool results so far", ref, len(tools))
	}
	var v any
	if err := json.Unmarshal([]byte(tools[n]), &v); err != nil {
		return nil, fmt.Errorf("$ref %q: tool result is not JSON", ref)
	}
	if len(parts) == 3 && parts[2] != "" {
		for _, k := range strings.Split(parts[2], ".") {
			m, ok := v.(map[string]any)
			if !ok {
				return nil, fmt.Errorf("$ref %q: no %q", ref, k)
			}
			v = m[k]
		}
	}
	if v == nil {
		return nil, fmt.Errorf("$ref %q resolved to null", ref)
	}
	return v, nil
}

func (s *Server) chat(w http.ResponseWriter, r *http.Request) {
	var req chatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":{"message":"bad request"}}`, 400)
		return
	}
	offered := map[string]bool{}
	for _, t := range req.Tools {
		offered[t.Function.Name] = true
	}
	rf := len(req.ResponseFormat) > 0 && string(req.ResponseFormat) != "null" && !strings.Contains(string(req.ResponseFormat), `"text"`)
	a, miss := s.script.Decide(Request{Messages: req.Messages, Offered: offered, ResponseFormat: rf})
	if miss != "" {
		s.record(Call{Turn: -1, Miss: miss})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]any{"error": map[string]any{"message": "stutter script miss: " + miss, "type": "server_error"}})
		return
	}
	s.mu.Lock()
	n := s.seen[a.Key]
	s.seen[a.Key]++
	s.mu.Unlock()
	h := sha256.Sum256([]byte(fmt.Sprintf("%s|%d", a.Key, n)))
	answerID := hex.EncodeToString(h[:12])
	c := Call{Turn: a.Turn, Step: a.Step}
	for i := range a.ToolCalls {
		th := sha256.Sum256([]byte(fmt.Sprintf("%s|%d", answerID, i)))
		a.ToolCalls[i].ID = "call_" + hex.EncodeToString(th[:12])
		c.Tools = append(c.Tools, a.ToolCalls[i].Name)
	}
	s.record(c)
	model := req.Model
	if model == "" {
		model = "stutter-scripted"
	}
	id := "chatcmpl-" + answerID
	finish := "stop"
	if len(a.ToolCalls) > 0 {
		finish = "tool_calls"
	}
	usage := map[string]int{"prompt_tokens": 10, "completion_tokens": 5, "total_tokens": 15}
	if !req.Stream {
		msg := map[string]any{"role": "assistant", "content": nil}
		if a.Content != "" || len(a.ToolCalls) == 0 {
			msg["content"] = a.Content
		}
		if len(a.ToolCalls) > 0 {
			var tcs []any
			for _, tc := range a.ToolCalls {
				tcs = append(tcs, map[string]any{"id": tc.ID, "type": "function", "function": map[string]any{"name": tc.Name, "arguments": tc.Args}})
			}
			msg["tool_calls"] = tcs
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"id": id, "object": "chat.completion", "created": 1700000000, "model": model,
			"choices": []any{map[string]any{"index": 0, "message": msg, "finish_reason": finish, "logprobs": nil}},
			"usage":   usage,
		})
		return
	}
	w.Header().Set("Content-Type", "text/event-stream")
	fl, _ := w.(http.Flusher)
	send := func(delta map[string]any, finish any, extra map[string]any) {
		ch := map[string]any{"id": id, "object": "chat.completion.chunk", "created": 1700000000, "model": model,
			"choices": []any{map[string]any{"index": 0, "delta": delta, "finish_reason": finish, "logprobs": nil}}}
		for k, v := range extra {
			ch[k] = v
		}
		b, _ := json.Marshal(ch)
		fmt.Fprintf(w, "data: %s\n\n", b)
		if fl != nil {
			fl.Flush()
		}
	}
	send(map[string]any{"role": "assistant", "content": ""}, nil, nil)
	if a.Content != "" {
		send(map[string]any{"content": a.Content}, nil, nil)
	}
	for i, tc := range a.ToolCalls {
		send(map[string]any{"tool_calls": []any{map[string]any{"index": i, "id": tc.ID, "type": "function",
			"function": map[string]any{"name": tc.Name, "arguments": ""}}}}, nil, nil)
		send(map[string]any{"tool_calls": []any{map[string]any{"index": i, "function": map[string]any{"arguments": tc.Args}}}}, nil, nil)
	}
	send(map[string]any{}, finish, nil)
	if req.StreamOptions != nil && req.StreamOptions.IncludeUsage {
		b, _ := json.Marshal(map[string]any{"id": id, "object": "chat.completion.chunk", "created": 1700000000, "model": model, "choices": []any{}, "usage": usage})
		fmt.Fprintf(w, "data: %s\n\n", b)
	}
	io.WriteString(w, "data: [DONE]\n\n")
	if fl != nil {
		fl.Flush()
	}
}

// embeddings answers with a deterministic unit vector per input string.
func (s *Server) embeddings(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Model          string          `json:"model"`
		Input          json.RawMessage `json:"input"`
		EncodingFormat string          `json:"encoding_format"`
		Dimensions     int             `json:"dimensions"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":{"message":"bad request"}}`, 400)
		return
	}
	var inputs []string
	var one string
	if json.Unmarshal(req.Input, &one) == nil {
		inputs = []string{one}
	} else if json.Unmarshal(req.Input, &inputs) != nil {
		var toks [][]int
		json.Unmarshal(req.Input, &toks)
		for _, t := range toks {
			inputs = append(inputs, fmt.Sprint(t))
		}
	}
	dims := req.Dimensions
	if dims == 0 {
		dims = s.script.Dims
	}
	if dims == 0 {
		dims = 1536
	}
	var data []any
	for i, in := range inputs {
		v := Vector(in, dims)
		var emb any = v
		if req.EncodingFormat == "base64" {
			buf := make([]byte, 4*len(v))
			for j, f := range v {
				binary.LittleEndian.PutUint32(buf[4*j:], math.Float32bits(f))
			}
			emb = base64.StdEncoding.EncodeToString(buf)
		}
		data = append(data, map[string]any{"object": "embedding", "index": i, "embedding": emb})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"object": "list", "data": data, "model": req.Model,
		"usage": map[string]int{"prompt_tokens": len(inputs), "total_tokens": len(inputs)}})
}

// Vector is a deterministic unit vector for a string.
func Vector(s string, dims int) []float32 {
	v := make([]float32, dims)
	var norm float64
	seed := sha256.Sum256([]byte(s))
	block := seed
	for i := 0; i < dims; i++ {
		if i%8 == 0 && i > 0 {
			block = sha256.Sum256(block[:])
		}
		u := binary.LittleEndian.Uint32(block[(i%8)*4:])
		f := float64(u)/float64(math.MaxUint32)*2 - 1
		v[i] = float32(f)
		norm += f * f
	}
	n := float32(math.Sqrt(norm))
	for i := range v {
		v[i] /= n
	}
	return v
}
