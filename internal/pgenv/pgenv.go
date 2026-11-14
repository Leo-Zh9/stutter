// Package pgenv manages throwaway PostgreSQL databases. Every replay runs on
// a fresh clone of a template database, so runs never see each other's rows.
package pgenv

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Env struct {
	SocketDir string
	Port      string
	User      string
	admin     *pgxpool.Pool
	counter   atomic.Int64
	prefix    string
}

// FromRepo finds the socket started by scripts/bootstrap.sh. STUTTER_PGHOST
// overrides it.
func FromRepo(root string) (*Env, error) {
	host := os.Getenv("STUTTER_PGHOST")
	if host == "" {
		host = filepath.Join(root, ".pgsock")
	}
	port := os.Getenv("STUTTER_PGPORT")
	if port == "" {
		if b, err := os.ReadFile(filepath.Join(root, ".pgport")); err == nil {
			port = strings.TrimSpace(string(b))
		}
	}
	if port == "" {
		port = "5432"
	}
	e := &Env{SocketDir: host, Port: port, User: "postgres", prefix: fmt.Sprintf("r%d_", os.Getpid())}
	cfg, err := pgxpool.ParseConfig(e.DSN("postgres"))
	if err != nil {
		return nil, err
	}
	cfg.MaxConns = 16
	e.admin, err = pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return nil, err
	}
	if err := e.admin.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("postgres not reachable at %s (run scripts/bootstrap.sh): %w", host, err)
	}
	return e, nil
}

// DSN is a libpq keyword string over the unix socket.
func (e *Env) DSN(db string) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", e.SocketDir, e.Port, e.User, db)
}

// URL is a postgres:// URL over TCP on 127.0.0.1, for drivers and ORMs that
// only take URLs (Prisma, postgres.js).
func (e *Env) URL(db string) string {
	return fmt.Sprintf("postgres://%s@127.0.0.1:%s/%s?sslmode=disable", e.User, e.Port, db)
}

// SocketURL is a postgres:// URL over the unix socket. Services that open
// many short-lived connections should use it: on TCP, killed processes leave
// thousands of TIME_WAIT sockets and exhaust ephemeral ports.
func (e *Env) SocketURL(db string) string {
	return fmt.Sprintf("postgres://%s@/%s?host=%s&port=%s&sslmode=disable", e.User, db, e.SocketDir, e.Port)
}

func (e *Env) Admin() *pgxpool.Pool { return e.admin }

// CreateTemplate (re)creates a template database from schema SQL.
func (e *Env) CreateTemplate(ctx context.Context, name, schema string) error {
	if _, err := e.admin.Exec(ctx, "DROP DATABASE IF EXISTS "+pgx.Identifier{name}.Sanitize()+" WITH (FORCE)"); err != nil {
		return err
	}
	if _, err := e.admin.Exec(ctx, "CREATE DATABASE "+pgx.Identifier{name}.Sanitize()); err != nil {
		return err
	}
	conn, err := pgx.Connect(ctx, e.DSN(name))
	if err != nil {
		return err
	}
	defer conn.Close(ctx)
	if strings.TrimSpace(schema) == "" {
		return nil
	}
	_, err = conn.Exec(ctx, schema)
	return err
}

// ExecIn runs SQL in the named database.
func (e *Env) ExecIn(ctx context.Context, db, sql string) error {
	conn, err := pgx.Connect(ctx, e.DSN(db))
	if err != nil {
		return err
	}
	defer conn.Close(ctx)
	_, err = conn.Exec(ctx, sql)
	return err
}

// Clone copies the template into a new database and returns its name.
// FILE_COPY is faster than WAL_LOG for small templates with fsync off.
func (e *Env) Clone(ctx context.Context, template string) (string, error) {
	name := fmt.Sprintf("%s%d", e.prefix, e.counter.Add(1))
	_, err := e.admin.Exec(ctx, fmt.Sprintf("CREATE DATABASE %s TEMPLATE %s STRATEGY FILE_COPY",
		pgx.Identifier{name}.Sanitize(), pgx.Identifier{template}.Sanitize()))
	return name, err
}

func (e *Env) Drop(ctx context.Context, name string) error {
	_, err := e.admin.Exec(ctx, "DROP DATABASE IF EXISTS "+pgx.Identifier{name}.Sanitize()+" WITH (FORCE)")
	return err
}

func (e *Env) Close() { e.admin.Close() }

// Band is the numeric range one sequence mints into. Giving every sequence
// its own range lets stutter tell a minted id from any other number, and
// tell whose id it is.
type Band struct {
	Seq    string `json:"sequence"` // schema.name
	Type   string `json:"type"`     // bigint | integer
	Start  int64  `json:"start"`
	Width  int64  `json:"width"`
	Offset int64  `json:"offset"` // shift used by the inference baseline
}

func (b Band) Contains(v int64) bool { return v >= b.Start && v < b.Start+b.Width }

// ApplyBands moves every sequence in the given schemas to the start of its
// own band and returns the bands, plus sequences left out (smallint, or no
// room for a band below max_value). Sequence k (sorted by name) gets:
//
//	bigint:  start (k+1)*10^12, width 10^12, inference offset (k+1)*10^6
//	integer: start (k+1)*10^7,  width 10^7,  inference offset (k+1)*10^5
func (e *Env) ApplyBands(ctx context.Context, db string, schemas []string) ([]Band, []string, error) {
	conn, err := pgx.Connect(ctx, e.DSN(db))
	if err != nil {
		return nil, nil, err
	}
	defer conn.Close(ctx)
	rows, err := conn.Query(ctx, `SELECT schemaname || '.' || sequencename, data_type::text, max_value
		FROM pg_sequences WHERE schemaname = ANY($1) ORDER BY 1`, schemas)
	if err != nil {
		return nil, nil, err
	}
	type seq struct {
		name, typ string
		max       int64
	}
	var seqs []seq
	for rows.Next() {
		var s seq
		if err := rows.Scan(&s.name, &s.typ, &s.max); err != nil {
			return nil, nil, err
		}
		seqs = append(seqs, s)
	}
	rows.Close()
	var bands []Band
	var skipped []string
	ib, bb := int64(0), int64(0)
	for _, s := range seqs {
		var b Band
		switch s.typ {
		case "bigint":
			bb++
			b = Band{Seq: s.name, Type: s.typ, Start: bb * 1e12, Width: 1e12, Offset: bb * 1e6}
		case "integer":
			ib++
			b = Band{Seq: s.name, Type: s.typ, Start: ib * 1e7, Width: 1e7, Offset: ib * 1e5}
		default:
			skipped = append(skipped, s.name+" ("+s.typ+")")
			continue
		}
		if b.Start+b.Width > s.max {
			skipped = append(skipped, s.name+" (no room below max_value)")
			continue
		}
		if _, err := conn.Exec(ctx, `SELECT setval($1::regclass, $2, false)`, s.name, b.Start); err != nil {
			return nil, nil, err
		}
		bands = append(bands, b)
	}
	return bands, skipped, nil
}

// ShiftBands advances each banded sequence by its inference offset.
func (e *Env) ShiftBands(ctx context.Context, db string, bands []Band) error {
	conn, err := pgx.Connect(ctx, e.DSN(db))
	if err != nil {
		return err
	}
	defer conn.Close(ctx)
	for _, b := range bands {
		if _, err := conn.Exec(ctx, `SELECT setval($1::regclass, $2, false)`, b.Seq, b.Start+b.Offset); err != nil {
			return err
		}
	}
	return nil
}
