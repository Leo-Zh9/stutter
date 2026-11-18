package trace

import (
	"encoding/json"
	"testing"
)

func TestRetryRules(t *testing.T) {
	p := Policy{Rules: []Rule{
		{Method: "POST", Path: "/chat", BodyKey: "message_id"},
		{Method: "POST", Path: "/threads/*/approve", Resume: &ResumeRule{Thread: "path:2"}},
	}}
	chat := func(body string) Entry {
		return Entry{Kind: API, Method: "POST", Path: "/chat", Body: json.RawMessage(body)}
	}
	if !p.Retryable(chat(`{"message_id":"m1","text":"hi"}`)) {
		t.Fatal("a chat message with its own id may be retried")
	}
	if p.Retryable(chat(`{"thread_id":"t1","text":"hi"}`)) || p.Retryable(chat(`{"message_id":""}`)) {
		t.Fatal("a chat message without a per-action id must not be retried")
	}
	ap := Entry{Kind: API, Method: "POST", Path: "/threads/t-9/approve?x=1"}
	if !p.Retryable(ap) {
		t.Fatal("a resume rule makes the approval retryable")
	}
	if th, ok := p.Resume(ap); !ok || th != "t-9" {
		t.Fatalf("thread from path: %q %v", th, ok)
	}
	if p.Retryable(Entry{Kind: API, Method: "GET", Path: "/chat"}) {
		t.Fatal("method must match")
	}
	if !p.Retryable(Entry{Kind: API, Headers: map[string]string{"idempotency-key": "k"}}) {
		t.Fatal("the header still works")
	}
}
