// Package testenv gives integration tests a Postgres environment, a built
// minipay binary and the committed trace.
package testenv

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"

	"github.com/Leo-Zh9/stutter/internal/oracle"
	"github.com/Leo-Zh9/stutter/internal/pgenv"
	"github.com/Leo-Zh9/stutter/internal/replay"
	"github.com/Leo-Zh9/stutter/internal/target"
	"github.com/Leo-Zh9/stutter/internal/trace"
)

func Root() string {
	_, f, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(f), "..", "..")
}

// Env connects to the bootstrap cluster. Without it the test is skipped,
// unless STUTTER_REQUIRE_PG=1 (set by `make test`), where it fails.
func Env(t *testing.T) *pgenv.Env {
	t.Helper()
	e, err := pgenv.FromRepo(Root())
	if err != nil {
		if os.Getenv("STUTTER_REQUIRE_PG") == "1" {
			t.Fatal(err)
		}
		t.Skip(err)
	}
	t.Cleanup(e.Close)
	return e
}

var (
	buildOnce sync.Once
	binPath   string
	buildErr  error
)

func Minipay(t *testing.T) string {
	t.Helper()
	buildOnce.Do(func() {
		dir, err := os.MkdirTemp("", "stutter-test")
		if err != nil {
			buildErr = err
			return
		}
		binPath = filepath.Join(dir, "minipay")
		cmd := exec.Command("go", "build", "-o", binPath, "./targets/minipay")
		cmd.Dir = Root()
		out, err := cmd.CombinedOutput()
		if err != nil {
			buildErr = err
			binPath = string(out)
		}
	})
	if buildErr != nil {
		t.Fatalf("build minipay: %v %s", buildErr, binPath)
	}
	return binPath
}

func Trace(t *testing.T) []trace.Entry {
	t.Helper()
	tr, err := trace.Read(filepath.Join(Root(), "traces", "minipay.jsonl"))
	if err != nil {
		t.Fatal(err)
	}
	return tr
}

// Config builds a replay config for one minipay variant on its own template.
func Config(t *testing.T, env *pgenv.Env, bug string) replay.Config {
	t.Helper()
	bin := Minipay(t)
	cmd := exec.Command(bin, "-schema")
	cmd.Env = append(os.Environ(), "BUG="+bug)
	schema, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
	}
	name := "tpl_test_" + strings.NewReplacer("-", "_", "/", "_").Replace(strings.ToLower(t.Name()))
	if bug != "" {
		name += "_" + strings.ReplaceAll(bug, "-", "_")
	}
	if len(name) > 60 {
		name = name[:60]
	}
	if err := env.CreateTemplate(context.Background(), name, string(schema)); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { env.Drop(context.Background(), name) })
	return replay.Config{
		Env: env, Template: name, Barrier: true,
		Target: target.Spec{Command: []string{bin}, Env: []string{"BUG=" + bug}, HealthPath: "/healthz"},
		Oracle: oracle.Options{Ignore: map[string]bool{"idempotency_keys.request_hash": true}},
	}
}
