package agentstub

import (
	"bufio"
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

var script = &Script{Turns: []Turn{
	{User: "refund order 42", Steps: []Step{
		{ToolCalls: []ToolCall{{Name: "lookup", Args: map[string]any{"order": 42.0}}}},
		{ToolCalls: []ToolCall{{Name: "refund", Args: map[string]any{"id": map[string]any{"$ref": "tool:0:order.id"}}}}},
		{Content: "Refunded."},
	}},
}}

func msgs(t *testing.T, s string) []message {
	t.Helper()
	var m []message
	if err := json.Unmarshal([]byte(s), &m); err != nil {
		t.Fatal(err)
	}
	return m
}

func req(m []message, offered map[string]bool) Request { return Request{Messages: m, Offered: offered} }

func TestStepFollowsAssistantCount(t *testing.T) {
	offered := map[string]bool{"lookup": true, "refund": true}
	a, miss := script.Decide(req(msgs(t, `[{"role":"system","content":"x"},{"role":"user","content":"please refund order 42"}]`), offered))
	if miss != "" || a.Step != 0 || a.ToolCalls[0].Name != "lookup" {
		t.Fatalf("step 0: %+v %s", a, miss)
	}
	conv := `[{"role":"user","content":[{"type":"text","text":"refund order 42"}]},
		{"role":"assistant","content":null,"tool_calls":[{"id":"c1"}]},
		{"role":"tool","tool_call_id":"c1","content":"{\"order\":{\"id\":\"ord_9\"}}"}]`
	a, miss = script.Decide(req(msgs(t, conv), offered))
	if miss != "" || a.Step != 1 || a.ToolCalls[0].Args != `{"id":"ord_9"}` {
		t.Fatalf("step 1 with $ref: %+v %s", a, miss)
	}
}

func TestMisses(t *testing.T) {
	if _, miss := script.Decide(req(msgs(t, `[{"role":"user","content":"hello"}]`), nil)); miss == "" {
		t.Fatal("an unmatched user message must miss")
	}
	if _, miss := script.Decide(req(msgs(t, `[{"role":"user","content":"refund order 42"}]`), map[string]bool{"other": true})); miss == "" {
		t.Fatal("a tool the request did not offer must miss")
	}
	past := `[{"role":"user","content":"refund order 42"},{"role":"assistant"},{"role":"assistant"},{"role":"assistant"}]`
	if _, miss := script.Decide(req(msgs(t, past), nil)); miss == "" {
		t.Fatal("a step past the end of the turn must miss")
	}
}

func TestServerStreamsToolCalls(t *testing.T) {
	s, err := Start(script)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()
	body := `{"model":"gpt-x","stream":true,"stream_options":{"include_usage":true},
		"tools":[{"type":"function","function":{"name":"lookup"}}],
		"messages":[{"role":"user","content":"refund order 42"}]}`
	resp, err := http.Post(s.URL+"/chat/completions", "application/json", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	var args, name string
	sc := bufio.NewScanner(resp.Body)
	done := false
	for sc.Scan() {
		line := strings.TrimPrefix(sc.Text(), "data: ")
		if line == "[DONE]" {
			done = true
			continue
		}
		if line == "" {
			continue
		}
		var ch struct {
			Choices []struct {
				Delta struct {
					ToolCalls []struct {
						Function struct{ Name, Arguments string }
					} `json:"tool_calls"`
				}
			}
		}
		json.Unmarshal([]byte(line), &ch)
		for _, c := range ch.Choices {
			for _, tc := range c.Delta.ToolCalls {
				name += tc.Function.Name
				args += tc.Function.Arguments
			}
		}
	}
	if !done || name != "lookup" || args != `{"order":42}` {
		t.Fatalf("stream: done=%v name=%q args=%q", done, name, args)
	}
	if len(s.Calls()) != 1 || s.Misses() != 0 {
		t.Fatalf("calls %+v", s.Calls())
	}
}

func TestEmbeddingsAreDeterministic(t *testing.T) {
	s, _ := Start(&Script{Dims: 8})
	defer s.Close()
	get := func() []byte {
		resp, err := http.Post(s.URL+"/embeddings", "application/json", bytes.NewReader([]byte(`{"input":["a","b"]}`)))
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()
		var b bytes.Buffer
		b.ReadFrom(resp.Body)
		return b.Bytes()
	}
	if !bytes.Equal(get(), get()) {
		t.Fatal("embeddings must be deterministic")
	}
}

func TestRepeatAfterCompletedRunDoesNotActAgain(t *testing.T) {
	done := `[{"role":"user","content":"refund order 42"},{"role":"assistant"},{"role":"tool","content":"{\"order\":{\"id\":\"o\"}}"},
		{"role":"assistant"},{"role":"tool","content":"{}"},{"role":"assistant","content":"Refunded."},
		{"role":"user","content":"refund order 42"}]`
	a, miss := script.Decide(req(msgs(t, done), nil))
	if miss != "" || len(a.ToolCalls) != 0 || a.Content != "Refunded." {
		t.Fatalf("got %+v %s", a, miss)
	}
	partial := `[{"role":"user","content":"refund order 42"},{"role":"assistant"},{"role":"user","content":"refund order 42"}]`
	a, _ = script.Decide(req(msgs(t, partial), nil))
	if len(a.ToolCalls) != 1 || a.ToolCalls[0].Name != "lookup" {
		t.Fatalf("an unfinished earlier run does not count as done: %+v", a)
	}
}

func TestSideRulesAndFreshIDs(t *testing.T) {
	sc := &Script{Turns: script.Turns, Side: []SideRule{{When: When{NoTools: true, SystemContains: "title"}, Content: "Refund chat"}}}
	a, miss := sc.Decide(Request{Messages: msgs(t, `[{"role":"system","content":"write a title"},{"role":"user","content":"refund order 42"}]`)})
	if miss != "" || a.Turn != -1 || a.Content != "Refund chat" {
		t.Fatalf("side rule: %+v %s", a, miss)
	}
	s, _ := Start(script)
	defer s.Close()
	call := func() string {
		resp, err := http.Post(s.URL+"/chat/completions", "application/json",
			strings.NewReader(`{"messages":[{"role":"user","content":"refund order 42"}]}`))
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()
		var out struct {
			Choices []struct {
				Message struct {
					ToolCalls []struct{ ID string } `json:"tool_calls"`
				}
			}
		}
		json.NewDecoder(resp.Body).Decode(&out)
		return out.Choices[0].Message.ToolCalls[0].ID
	}
	if call() == call() {
		t.Fatal("a re-run of the same prompt must get a fresh tool-call id, as a real model would")
	}
}
