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

	"github.com/Leo-Zh9/stutter/internal/config"
	"github.com/Leo-Zh9/stutter/internal/pgenv"
	"github.com/Leo-Zh9/stutter/internal/replay"
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
	buildErr  error
	buildOut  string
)

// Minipay builds bin/minipay, the binary targets/minipay/stutter.json starts.
func Minipay(t *testing.T) string {
	t.Helper()
	bin := filepath.Join(Root(), "bin", "minipay")
	buildOnce.Do(func() {
		cmd := exec.Command("go", "build", "-o", bin, "./targets/minipay")
		cmd.Dir = Root()
		out, err := cmd.CombinedOutput()
		buildErr, buildOut = err, string(out)
	})
	if buildErr != nil {
		t.Fatalf("build minipay: %v %s", buildErr, buildOut)
	}
	return bin
}

func Trace(t *testing.T) []trace.Entry {
	t.Helper()
	tr, err := trace.Read(filepath.Join(Root(), "traces", "minipay.jsonl"))
	if err != nil {
		t.Fatal(err)
	}
	return tr
}

// Target loads targets/minipay/stutter.json.
func Target(t *testing.T) *config.Config {
	t.Helper()
	Minipay(t)
	c, err := config.Load(filepath.Join(Root(), "targets", "minipay", "stutter.json"))
	if err != nil {
		t.Fatal(err)
	}
	return c
}

// Config builds a replay config for one minipay variant on its own template,
// through the same stutter.json path the CLI uses.
func Config(t *testing.T, env *pgenv.Env, bug string) replay.Config {
	t.Helper()
	c := Target(t)
	name := "tpl_test_" + strings.NewReplacer("-", "_", "/", "_").Replace(strings.ToLower(t.Name()))
	if bug != "" {
		name += "_" + strings.ReplaceAll(bug, "-", "_")
	}
	if len(name) > 60 {
		name = name[:60]
	}
	extra := []string{"BUG=" + bug}
	if err := c.PrepareTemplate(context.Background(), env, name, extra); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { env.Drop(context.Background(), name) })
	return c.Replay(env, name, extra)
}
