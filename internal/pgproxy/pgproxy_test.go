package pgproxy_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/Leo-Zh9/stutter/internal/pgproxy"
	"github.com/Leo-Zh9/stutter/internal/testenv"
)

func setup(t *testing.T) (*pgproxy.Proxy, string) {
	env := testenv.Env(t)
	ctx := context.Background()
	if err := env.CreateTemplate(ctx, "tpl_test_pgproxy", `CREATE TABLE t (id serial PRIMARY KEY, v int UNIQUE)`); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { env.Drop(ctx, "tpl_test_pgproxy") })
	p, err := pgproxy.New(env.SocketDir, env.Port, false)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(p.Close)
	return p, fmt.Sprintf("host=%s port=%s user=postgres dbname=tpl_test_pgproxy sslmode=disable", p.Dir, p.Port)
}

// Each case runs in its own request window and must count exactly the
// expected commit points.
func TestCommitPoints(t *testing.T) {
	p, dsn := setup(t)
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close(ctx)
	cases := []struct {
		name string
		want int
		run  func() error
	}{
		{"explicit commit", 1, func() error {
			tx, err := conn.Begin(ctx)
			if err != nil {
				return err
			}
			tx.Exec(ctx, `INSERT INTO t (v) VALUES (1)`)
			tx.Exec(ctx, `UPDATE t SET v = 2 WHERE v = 1`)
			return tx.Commit(ctx)
		}},
		{"autocommit write", 1, func() error { _, err := conn.Exec(ctx, `INSERT INTO t (v) VALUES (3)`); return err }},
		{"read only", 0, func() error { _, err := conn.Exec(ctx, `SELECT count(*) FROM t`); return err }},
		{"rolled back", 0, func() error {
			tx, _ := conn.Begin(ctx)
			tx.Exec(ctx, `INSERT INTO t (v) VALUES (4)`)
			return tx.Rollback(ctx)
		}},
		{"failed commit", 0, func() error {
			tx, _ := conn.Begin(ctx)
			tx.Exec(ctx, `INSERT INTO t (v) VALUES (3)`) // unique violation aborts the transaction
			tx.Commit(ctx)                              // server answers ROLLBACK
			return nil
		}},
		{"multi-statement implicit transaction", 1, func() error {
			_, err := conn.PgConn().Exec(ctx, `INSERT INTO t (v) VALUES (10); INSERT INTO t (v) VALUES (11)`).ReadAll()
			return err
		}},
		{"pipelined batch, one sync", 1, func() error {
			b := &pgx.Batch{}
			b.Queue(`INSERT INTO t (v) VALUES (20)`)
			b.Queue(`INSERT INTO t (v) VALUES (21)`)
			return conn.SendBatch(ctx, b).Close()
		}},
	}
	for i, c := range cases {
		p.Begin(i)
		if err := c.run(); err != nil {
			t.Fatalf("%s: %v", c.name, err)
		}
		p.End()
		got, _ := p.Commits()
		if got[i] != c.want {
			t.Errorf("%s: %d commit points, want %d", c.name, got[i], c.want)
		}
	}
	if _, bg := p.Commits(); bg != 0 {
		t.Errorf("%d commits outside any window", bg)
	}
}

// An armed crash holds back the acknowledgement: the client never hears
// back, yet the row is committed.
func TestCrashHoldsBackTheAck(t *testing.T) {
	p, dsn := setup(t)
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		t.Fatal(err)
	}
	p.ArmCrash(7, 1)
	p.Begin(7)
	done := make(chan error, 1)
	go func() { _, err := conn.Exec(ctx, `INSERT INTO t (v) VALUES (99)`); done <- err }()
	select {
	case <-p.Crashed():
	case <-time.After(3 * time.Second):
		t.Fatal("crash point not reached")
	}
	select {
	case err := <-done:
		t.Fatalf("client got an answer (%v) after its ack was held back", err)
	case <-time.After(200 * time.Millisecond):
	}
	p.DropConns()
	p.Resume()
	p.End()
	check, err := pgx.Connect(ctx, dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer check.Close(ctx)
	var n int
	check.QueryRow(ctx, `SELECT count(*) FROM t WHERE v = 99`).Scan(&n)
	if n != 1 {
		t.Fatalf("held-back commit not durable: %d rows", n)
	}
}
