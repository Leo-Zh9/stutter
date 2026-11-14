// Package pgenv manages throwaway PostgreSQL databases. Every replay runs on
// a fresh clone of a template database, so runs never see each other's rows.
package pgenv

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync/atomic"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Env struct {
	SocketDir string
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
	e := &Env{SocketDir: host, User: "postgres", prefix: fmt.Sprintf("r%d_", os.Getpid())}
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

func (e *Env) DSN(db string) string {
	return fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable", e.SocketDir, e.User, db)
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
	_, err = conn.Exec(ctx, schema)
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
