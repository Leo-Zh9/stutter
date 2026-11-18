package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Leo-Zh9/stutter/internal/trace"
)

func write(t *testing.T, body string) string {
	p := filepath.Join(t.TempDir(), "stutter.json")
	if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	return p
}

func TestLoadDefaultsAndPaths(t *testing.T) {
	p := write(t, `{"start":["./bin/svc","--port","{PORT}"],"env":["DATABASE_URL={DATABASE_URL}"],"prepare":{"schema_file":"schema.sql"}}`)
	c, err := Load(p)
	if err != nil {
		t.Fatal(err)
	}
	if c.KeyHeader != "Idempotency-Key" || c.Health != "/" || len(c.Schemas) != 1 || c.Schemas[0] != "public" {
		t.Fatalf("defaults: %+v", c)
	}
	spec := c.Spec([]string{"BUG=x"})
	if !filepath.IsAbs(spec.Command[0]) || spec.Command[2] != "{PORT}" || spec.Env[len(spec.Env)-1] != "BUG=x" {
		t.Fatalf("spec: %+v", spec)
	}
}

func TestLoadRejectsBadConfigs(t *testing.T) {
	for name, body := range map[string]string{
		"no start":        `{"prepare":{"schema_file":"s.sql"}}`,
		"no prepare":      `{"start":["x"]}`,
		"two prepares":    `{"start":["x"],"prepare":{"schema_file":"s.sql","migrate_cmd":["m"]}}`,
		"unknown field":   `{"start":["x"],"prepare":{"schema_file":"s.sql"},"tracee":"t"}`,
		"unknown signing": `{"start":["x"],"prepare":{"schema_file":"s.sql"},"webhooks":{"signing":"github"}}`,
	} {
		if _, err := Load(write(t, body)); err == nil {
			t.Errorf("%s: accepted", name)
		} else if !strings.Contains(err.Error(), "stutter.json") {
			t.Errorf("%s: error does not name the file: %v", name, err)
		}
	}
}

// Two targets running at once must never share a template database: one
// rebuilding it under the other once corrupted a whole run.
func TestTemplateNamesAreDistinct(t *testing.T) {
	a, _ := Load(write(t, `{"name":"a","start":["x"],"prepare":{"schema_file":"s.sql"}}`))
	b, _ := Load(write(t, `{"name":"b","start":["x"],"prepare":{"schema_file":"s.sql"}}`))
	if a.TemplateName("hunt") == b.TemplateName("hunt") || a.TemplateName("hunt") == a.TemplateName("explore") {
		t.Fatalf("template names collide: %s %s %s", a.TemplateName("hunt"), b.TemplateName("hunt"), a.TemplateName("explore"))
	}
}

func TestResumeThreadMustNameAWildcard(t *testing.T) {
	rule := func(path, thread string) trace.Rule {
		return trace.Rule{Path: path, Resume: &trace.ResumeRule{Thread: thread}}
	}
	for _, bad := range []trace.Rule{rule("/threads/*/approve", "path:3"), rule("/threads/*/approve", "path:9"),
		rule("/threads/*/approve", "path:x"), rule("/approve", "thread"), rule("/approve", "body:")} {
		if checkThread(bad) == nil {
			t.Errorf("accepted %+v", bad.Resume)
		}
	}
	for _, good := range []trace.Rule{rule("/threads/*/approve", "path:2"), rule("/approve", "body:thread_id")} {
		if err := checkThread(good); err != nil {
			t.Errorf("rejected %+v: %v", good.Resume, err)
		}
	}
}
