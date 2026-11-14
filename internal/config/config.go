// Package config loads stutter.json, the description of one target service.
//
//	{
//	  "name": "minipay",
//	  "start": ["bin/minipay"],
//	  "env": ["DATABASE_URL={DSN}", "ADDR={ADDR}"],
//	  "health": "/healthz",
//	  "prepare": {"schema_cmd": ["bin/minipay", "-schema"]},
//	  "schemas": ["public"],
//	  "idempotency_header": "Idempotency-Key",
//	  "webhooks": {"prefix": "/webhooks", "signing": "stripe", "secret": "whsec_test", "retry": "non-2xx", "max_attempts": 5},
//	  "ignore": ["idempotency_keys.request_hash"],
//	  "trace": "traces/minipay.jsonl"
//	}
//
// Relative paths resolve against the config file's directory.
package config

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/Leo-Zh9/stutter/internal/oracle"
	"github.com/Leo-Zh9/stutter/internal/pgenv"
	"github.com/Leo-Zh9/stutter/internal/replay"
	"github.com/Leo-Zh9/stutter/internal/target"
	"github.com/Leo-Zh9/stutter/internal/trace"
)

type Prepare struct {
	// Exactly one way to build the schema: a SQL file, a command that prints
	// SQL, or a command that migrates the database named by {DATABASE_URL} /
	// {DSN} in its env (prisma migrate deploy, drizzle-kit migrate, ...).
	SchemaFile string   `json:"schema_file,omitempty"`
	SchemaCmd  []string `json:"schema_cmd,omitempty"`
	MigrateCmd []string `json:"migrate_cmd,omitempty"`
	MigrateEnv []string `json:"migrate_env,omitempty"`
	// Or let the service build its own schema: start it on the template
	// (e.g. with --auto-upgrade), send these requests, stop it.
	SetupRequests []SetupRequest `json:"setup_requests,omitempty"`
	// Optional SQL run after the schema exists: labelled fixture rows.
	SeedFile string `json:"seed_file,omitempty"`
	// Give every sequence its own numeric band so integer ids minted in
	// application SQL can be mapped and attributed (see docs/oracle.md).
	BandSequences bool `json:"band_sequences,omitempty"`
}

type SetupRequest struct {
	Method string          `json:"method"`
	Path   string          `json:"path"`
	Body   json.RawMessage `json:"body,omitempty"`
	Expect int             `json:"expect,omitempty"` // default: any 2xx
}

type Webhooks struct {
	Prefix      string `json:"prefix"`
	Object      string `json:"object_path,omitempty"` // JSON path grouping reorderable webhooks
	Signing     string `json:"signing,omitempty"` // "" or "stripe"
	Secret      string `json:"secret,omitempty"`
	Header      string `json:"header,omitempty"` // default Stripe-Signature
	Retry       string `json:"retry,omitempty"`  // "5xx" (default) or "non-2xx"
	MaxAttempts int    `json:"max_attempts,omitempty"`
}

type Config struct {
	Name         string   `json:"name"`
	Start        []string `json:"start"`
	Dir          string   `json:"dir,omitempty"`
	Env          []string `json:"env"`
	Health       string   `json:"health"`
	StartTimeout string   `json:"start_timeout,omitempty"`
	Prepare      Prepare  `json:"prepare"`
	Schemas      []string `json:"schemas,omitempty"`
	KeyHeader    string   `json:"idempotency_header,omitempty"`
	Webhooks     Webhooks `json:"webhooks"`
	Ignore       []string `json:"ignore,omitempty"`
	Trace        string   `json:"trace,omitempty"`

	base  string
	bands map[string][]pgenv.Band
	// Skipped lists sequences left out of banding, for the report.
	Skipped []string `json:"-"`
}

func Load(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var c Config
	dec := json.NewDecoder(strings.NewReader(string(b)))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&c); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	abs, _ := filepath.Abs(filepath.Dir(path))
	c.base = abs
	if len(c.Start) == 0 {
		return nil, fmt.Errorf("%s: start is required", path)
	}
	if c.Health == "" {
		c.Health = "/"
	}
	if c.KeyHeader == "" {
		c.KeyHeader = trace.DefaultKeyHeader
	}
	if len(c.Schemas) == 0 {
		c.Schemas = []string{"public"}
	}
	if c.Webhooks.Signing != "" && c.Webhooks.Signing != "stripe" {
		return nil, fmt.Errorf("%s: unknown webhooks.signing %q", path, c.Webhooks.Signing)
	}
	n := 0
	for _, set := range []bool{c.Prepare.SchemaFile != "", len(c.Prepare.SchemaCmd) > 0, len(c.Prepare.MigrateCmd) > 0, len(c.Prepare.SetupRequests) > 0} {
		if set {
			n++
		}
	}
	if n != 1 {
		return nil, fmt.Errorf("%s: prepare needs exactly one of schema_file, schema_cmd, migrate_cmd, setup_requests", path)
	}
	return &c, nil
}

// TemplateName names this config's template database for one purpose
// ("hunt", "explore", "record"). It includes a hash of the config's path, so
// two targets running at once never rebuild each other's template.
func (c *Config) TemplateName(kind string) string {
	sum := sha256.Sum256([]byte(c.base + "|" + c.Name))
	return fmt.Sprintf("tpl_%s_%s", kind, hex.EncodeToString(sum[:])[:10])
}

// Path resolves a config-relative path.
func (c *Config) Path(p string) string {
	if p == "" || filepath.IsAbs(p) {
		return p
	}
	return filepath.Join(c.base, p)
}

func (c *Config) resolveCmd(cmd []string) []string {
	out := append([]string(nil), cmd...)
	if len(out) > 0 && strings.Contains(out[0], "/") {
		out[0] = c.Path(out[0])
	}
	return out
}

func (c *Config) dir() string {
	if c.Dir == "" {
		return c.base
	}
	return c.Path(c.Dir)
}

// Spec is the launch spec, with extra env (e.g. BUG=...) appended.
func (c *Config) Spec(extraEnv []string) target.Spec {
	var to time.Duration
	if c.StartTimeout != "" {
		to, _ = time.ParseDuration(c.StartTimeout)
	}
	return target.Spec{
		Command: c.resolveCmd(c.Start), Env: append(append([]string(nil), c.Env...), extraEnv...),
		Dir: c.dir(), HealthPath: c.Health, StartTimeout: to,
	}
}

// Space is the fault space the config admits.
func (c *Config) Space() replay.Space {
	return replay.Space{KeyHeader: c.KeyHeader, WebhookObject: c.Webhooks.Object}
}

// Replay builds a replay config for a template database.
func (c *Config) Replay(env *pgenv.Env, template string, extraEnv []string) replay.Config {
	ign := map[string]bool{}
	for _, col := range c.Ignore {
		ign[col] = true
	}
	return replay.Config{
		Env: env, Template: template, Barrier: true, Target: c.Spec(extraEnv), Bands: c.bands[template],
		Oracle:    oracle.Options{Ignore: ign, Schemas: c.Schemas},
		KeyHeader: c.KeyHeader,
		Webhooks: replay.WebhookPolicy{
			Signing: c.Webhooks.Signing, Secret: c.Webhooks.Secret, Header: c.Webhooks.Header,
			RetryAny: c.Webhooks.Retry == "non-2xx", MaxAttempts: c.Webhooks.MaxAttempts,
		},
	}
}

// PrepareTemplate builds the template database: schema, then seed.
func (c *Config) PrepareTemplate(ctx context.Context, env *pgenv.Env, name string, extraEnv []string) error {
	p := c.Prepare
	schema := ""
	switch {
	case p.SchemaFile != "":
		b, err := os.ReadFile(c.Path(p.SchemaFile))
		if err != nil {
			return err
		}
		schema = string(b)
	case len(p.SchemaCmd) > 0:
		args := c.resolveCmd(p.SchemaCmd)
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Dir = c.dir()
		cmd.Env = append(os.Environ(), extraEnv...)
		out, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("schema_cmd: %w", err)
		}
		schema = string(out)
	}
	if err := env.CreateTemplate(ctx, name, schema); err != nil {
		return err
	}
	if len(p.MigrateCmd) > 0 {
		r := strings.NewReplacer("{DATABASE_URL_SOCKET}", env.SocketURL(name), "{DSN}", env.DSN(name), "{DATABASE_URL}", env.URL(name))
		args := c.resolveCmd(p.MigrateCmd)
		for i := range args {
			args[i] = r.Replace(args[i])
		}
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Dir = c.dir()
		cmd.Env = append(os.Environ(), extraEnv...)
		for _, e := range p.MigrateEnv {
			cmd.Env = append(cmd.Env, r.Replace(e))
		}
		if out, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("migrate_cmd: %w\n%s", err, out)
		}
	}
	if len(p.SetupRequests) > 0 {
		if err := c.runSetup(ctx, env, name, extraEnv); err != nil {
			return fmt.Errorf("setup_requests: %w", err)
		}
	}
	if p.SeedFile != "" {
		b, err := os.ReadFile(c.Path(p.SeedFile))
		if err != nil {
			return err
		}
		if err := env.ExecIn(ctx, name, string(b)); err != nil {
			return fmt.Errorf("seed: %w", err)
		}
	}
	if p.BandSequences {
		bands, skipped, err := env.ApplyBands(ctx, name, c.Schemas)
		if err != nil {
			return fmt.Errorf("band sequences: %w", err)
		}
		if c.bands == nil {
			c.bands = map[string][]pgenv.Band{}
		}
		c.bands[name], c.Skipped = bands, skipped
	}
	return nil
}

func (c *Config) runSetup(ctx context.Context, env *pgenv.Env, db string, extraEnv []string) error {
	proc, err := target.Start(ctx, c.Spec(extraEnv), target.DB{DSN: env.DSN(db), URL: env.URL(db), SocketURL: env.SocketURL(db)})
	if err != nil {
		return err
	}
	defer proc.Stop()
	for _, r := range c.Prepare.SetupRequests {
		req, err := http.NewRequestWithContext(ctx, r.Method, proc.BaseURL+r.Path, bytes.NewReader(r.Body))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		resp.Body.Close()
		ok := resp.StatusCode >= 200 && resp.StatusCode < 300
		if r.Expect != 0 {
			ok = resp.StatusCode == r.Expect
		}
		if !ok {
			return fmt.Errorf("%s %s: HTTP %d", r.Method, r.Path, resp.StatusCode)
		}
	}
	return nil
}
