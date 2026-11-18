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
	"strconv"
	"strings"
	"time"

	"github.com/Leo-Zh9/stutter/internal/agentstub"
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
	Signing     string `json:"signing,omitempty"`     // "" or "stripe"
	Secret      string `json:"secret,omitempty"`
	Header      string `json:"header,omitempty"` // default Stripe-Signature
	Retry       string `json:"retry,omitempty"`  // "5xx" (default) or "non-2xx"
	MaxAttempts int    `json:"max_attempts,omitempty"`
}

type Model struct {
	Script string `json:"script"` // path to the script JSON
	// InterruptSQL overrides the query that names the interrupt pending on
	// thread $1 (NULL or '' if none); default: LangGraph's Postgres
	// checkpointer (see replay.DefaultInterruptSQL).
	InterruptSQL string `json:"interrupt_sql,omitempty"`
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
	// Generated lists columns holding ids the app mints itself
	// ("table.column"), compared like database-generated ids.
	Generated []string `json:"generated,omitempty"`
	Trace     string   `json:"trace,omitempty"`
	// IgnoreTables are left out of the oracle and the lock barrier: a
	// framework's own bookkeeping, such as a checkpointer's tables.
	IgnoreTables []string `json:"ignore_tables,omitempty"`
	// Headers are sent on every replayed request (a static test token, for
	// example). Values may use ${ENV} references; they are never recorded.
	Headers map[string]string `json:"headers,omitempty"`
	// RetryRules widen which API calls the client may retry (trace.Rule).
	RetryRules []trace.Rule `json:"retry_rules,omitempty"`
	// Model is a scripted model (agentstub) served to the service at
	// {MODEL_URL} on every run: no recording, no paid calls.
	Model *Model `json:"model,omitempty"`

	script *agentstub.Script

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
	if c.Model != nil {
		sc, err := agentstub.Load(c.Path(c.Model.Script))
		if err != nil {
			return nil, err
		}
		c.script = sc
	}
	for i, r := range c.RetryRules {
		if r.Path == "" || (r.BodyKey == "") == (r.Resume == nil) {
			return nil, fmt.Errorf("%s: retry rule %d needs a path and exactly one of body_key or resume", path, i)
		}
		if r.Resume != nil {
			if err := checkThread(r); err != nil {
				return nil, fmt.Errorf("%s: retry rule %d: %w", path, i, err)
			}
		}
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

// checkThread requires a resume rule's thread to name a wildcard path
// segment ("path:N" where segment N of the rule's path is "*") or a body
// field ("body:dotted.path"): anything else would quietly make every resume
// look like the same thread, or none.
func checkThread(r trace.Rule) error {
	t := r.Resume.Thread
	switch {
	case strings.HasPrefix(t, "body:") && len(t) > len("body:"):
		return nil
	case strings.HasPrefix(t, "path:"):
		n, err := strconv.Atoi(strings.TrimPrefix(t, "path:"))
		segs := strings.Split(strings.Trim(r.Path, "/"), "/")
		if err != nil || n < 1 || n > len(segs) || segs[n-1] != "*" {
			return fmt.Errorf("resume thread %q must name a * segment of %q", t, r.Path)
		}
		return nil
	}
	return fmt.Errorf("resume thread %q must be path:N or body:dotted.path", t)
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
	return replay.Space{KeyHeader: c.KeyHeader, WebhookObject: c.Webhooks.Object, Rules: c.RetryRules}
}

// Replay builds a replay config for a template database.
func (c *Config) Replay(env *pgenv.Env, template string, extraEnv []string) replay.Config {
	ign := map[string]bool{}
	for _, col := range c.Ignore {
		ign[col] = true
	}
	gen := map[string]bool{}
	for _, col := range c.Generated {
		gen[col] = true
	}
	it := map[string]bool{}
	for _, t := range c.IgnoreTables {
		it[t] = true
	}
	rc := replay.Config{
		Env: env, Template: template, Barrier: true, Target: c.Spec(extraEnv), Bands: c.bands[template],
		Oracle:    oracle.Options{Ignore: ign, Generated: gen, Schemas: c.Schemas, IgnoreTables: it},
		KeyHeader: c.KeyHeader,
		Rules:     c.RetryRules,
		Headers:   c.ResolvedHeaders(),
		Model:     c.script,
		Webhooks: replay.WebhookPolicy{
			Signing: c.Webhooks.Signing, Secret: c.Webhooks.Secret, Header: c.Webhooks.Header,
			RetryAny: c.Webhooks.Retry == "non-2xx", MaxAttempts: c.Webhooks.MaxAttempts,
		},
	}
	if c.Model != nil {
		rc.InterruptSQL = c.Model.InterruptSQL
	}
	return rc
}

// ResolvedHeaders fills ${ENV} references in the static headers.
func (c *Config) ResolvedHeaders() map[string]string {
	if len(c.Headers) == 0 {
		return nil
	}
	out := map[string]string{}
	for k, v := range c.Headers {
		out[k] = os.ExpandEnv(v)
	}
	return out
}

// Script is the loaded model script, if any.
func (c *Config) Script() *agentstub.Script { return c.script }

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
		r := strings.NewReplacer("{DATABASE_URL_SOCKET}", env.SocketURL(name), "{DATABASE_SOCKET_NOSCHEME}", strings.TrimPrefix(env.SocketURL(name), "postgres://"), "{DSN}", env.DSN(name), "{DATABASE_URL}", env.URL(name))
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
	dbFor := target.DB{DSN: env.DSN(db), URL: env.URL(db), SocketURL: env.SocketURL(db)}
	if c.script != nil {
		stub, err := agentstub.Start(c.script)
		if err != nil {
			return err
		}
		defer stub.Close()
		dbFor.Vars = map[string]string{"{MODEL_URL}": stub.URL}
	}
	proc, err := target.Start(ctx, c.Spec(extraEnv), dbFor)
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
