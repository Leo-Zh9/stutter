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
