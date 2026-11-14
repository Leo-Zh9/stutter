// Package pgproxy sits between a service and PostgreSQL and sees each commit
// acknowledgement before the service does.
//
// It forwards bytes unchanged in both directions and parses only the backend
// stream. A commit point is a ReadyForQuery that ends a committed write:
//
//   - status T -> I with a CommandComplete "COMMIT" since the previous
//     ReadyForQuery (an explicit transaction committed), or
//   - status I -> I with a write tag (INSERT, UPDATE, DELETE, MERGE, COPY)
//     since the previous ReadyForQuery (an autocommit or implicit transaction).
//
// Each ReadyForQuery counts at most once, so a multi-statement implicit
// transaction is one commit point. A failed COMMIT reports ROLLBACK and is
// not a commit point. Every Sync yields exactly one ReadyForQuery, in order,
// so this holds for the simple, extended and pipelined protocols alike.
//
// To crash, the proxy holds back the ReadyForQuery of the chosen commit: the
// commit is durable, but the service never learns it happened. The caller
// then kills the service.
//
// Not seen as commit points: COMMIT AND CHAIN (status stays T), procedures
// that COMMIT internally via CALL, and writes inside a function called from
// a plain SELECT.
package pgproxy

import (
	"encoding/binary"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const (
	sslRequest    = 80877103
	gssEncRequest = 80877104
	cancelRequest = 80877102
)

// Proxy listens on a unix socket (and optionally TCP) and forwards to the
// PostgreSQL unix socket upstream.
type Proxy struct {
	Dir      string // directory holding .s.PGSQL.<port>; use as libpq host
	Port     string
	TCPAddr  string // 127.0.0.1:<port>, when TCP is enabled
	upstream string

	mu      sync.Mutex
	conns   map[net.Conn]bool
	lns     []net.Listener
	window  int  // seq of the request in flight, -1 when idle
	commits map[int]int
	bg      int
	crashAt struct {
		seq, k int
		armed  bool
	}
	crashed chan struct{}
	held    bool
	drop    chan struct{}
}

func (p *Proxy) dropped() <-chan struct{} {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.drop
}

// New starts a proxy to the PostgreSQL socket in upstreamDir/.s.PGSQL.<port>.
func New(upstreamDir, port string, tcp bool) (*Proxy, error) {
	dir, err := os.MkdirTemp("", "stp")
	if err != nil {
		return nil, err
	}
	p := &Proxy{Dir: dir, Port: port, upstream: filepath.Join(upstreamDir, ".s.PGSQL."+port),
		conns: map[net.Conn]bool{}, window: -1, commits: map[int]int{}, crashed: make(chan struct{}), drop: make(chan struct{})}
	ul, err := net.Listen("unix", filepath.Join(dir, ".s.PGSQL."+port))
	if err != nil {
		return nil, err
	}
	p.lns = append(p.lns, ul)
	go p.serve(ul)
	if tcp {
		tl, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			p.Close()
			return nil, err
		}
		p.TCPAddr = tl.Addr().String()
		p.lns = append(p.lns, tl)
		go p.serve(tl)
	}
	return p, nil
}

func (p *Proxy) serve(l net.Listener) {
	for {
		c, err := l.Accept()
		if err != nil {
			return
		}
		go p.handle(c)
	}
}

// Begin marks request seq as in flight; commits until End belong to it.
func (p *Proxy) Begin(seq int) {
	p.mu.Lock()
	p.window = seq
	p.mu.Unlock()
}

func (p *Proxy) End() {
	p.mu.Lock()
	p.window = -1
	p.mu.Unlock()
}

// ArmCrash makes the proxy hold back the ack of commit k (1-based) of
// request seq and close Crashed.
func (p *Proxy) ArmCrash(seq, k int) {
	p.mu.Lock()
	p.crashAt.seq, p.crashAt.k, p.crashAt.armed = seq, k, true
	p.mu.Unlock()
}

// Resume clears a crash so the restarted service is served normally.
func (p *Proxy) Resume() {
	p.mu.Lock()
	p.held = false
	p.crashAt.armed = false
	p.mu.Unlock()
}

// Crashed is closed once the armed commit has been held back.
func (p *Proxy) Crashed() <-chan struct{} { return p.crashed }

// Commits returns commit points per request seq, and those seen while idle.
func (p *Proxy) Commits() (map[int]int, int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	out := make(map[int]int, len(p.commits))
	for k, v := range p.commits {
		out[k] = v
	}
	return out, p.bg
}

// DropConns closes every proxied connection (after the service is killed).
func (p *Proxy) DropConns() {
	p.mu.Lock()
	for c := range p.conns {
		c.Close()
	}
	p.conns = map[net.Conn]bool{}
	close(p.drop)
	p.drop = make(chan struct{})
	p.mu.Unlock()
}

func (p *Proxy) Close() {
	for _, l := range p.lns {
		l.Close()
	}
	p.DropConns()
	os.RemoveAll(p.Dir)
}

func (p *Proxy) track(c net.Conn, on bool) {
	p.mu.Lock()
	if on {
		p.conns[c] = true
	} else {
		delete(p.conns, c)
	}
	p.mu.Unlock()
}

func (p *Proxy) handle(client net.Conn) {
	defer client.Close()
	p.track(client, true)
	defer p.track(client, false)
	// Startup: answer SSL and GSS encryption requests with 'N' (no), pass
	// cancel requests and the startup message through.
	var first []byte
	for {
		hdr := make([]byte, 8)
		if _, err := io.ReadFull(client, hdr); err != nil {
			return
		}
		n := int(binary.BigEndian.Uint32(hdr[:4]))
		code := binary.BigEndian.Uint32(hdr[4:8])
		if n < 8 || n > 1<<20 {
			return
		}
		rest := make([]byte, n-8)
		if _, err := io.ReadFull(client, rest); err != nil {
			return
		}
		if code == sslRequest || code == gssEncRequest {
			if _, err := client.Write([]byte{'N'}); err != nil {
				return
			}
			continue
		}
		first = append(hdr, rest...)
		_ = code == cancelRequest
		break
	}
	server, err := net.Dial("unix", p.upstream)
	if err != nil {
		return
	}
	defer server.Close()
	p.track(server, true)
	defer p.track(server, false)
	if _, err := server.Write(first); err != nil {
		return
	}
	go func() {
		io.Copy(server, client)
		server.Close()
	}()
	p.backend(server, client)
}

// backend copies server -> client message by message, watching for commits.
func (p *Proxy) backend(server, client net.Conn) {
	status := byte('I')
	committed, wrote := false, false
	hdr := make([]byte, 5)
	for {
		if _, err := io.ReadFull(server, hdr); err != nil {
			return
		}
		n := int(binary.BigEndian.Uint32(hdr[1:5]))
		if n < 4 {
			return
		}
		body := make([]byte, n-4)
		if _, err := io.ReadFull(server, body); err != nil {
			return
		}
		switch hdr[0] {
		case 'C': // CommandComplete
			tag := strings.TrimRight(string(body), "\x00")
			switch {
			case tag == "COMMIT":
				committed = true
			case strings.HasPrefix(tag, "INSERT"), strings.HasPrefix(tag, "UPDATE"), strings.HasPrefix(tag, "DELETE"),
				strings.HasPrefix(tag, "MERGE"), strings.HasPrefix(tag, "COPY"):
				wrote = true
			}
		case 'Z': // ReadyForQuery
			next := body[0]
			point := (status == 'T' && next == 'I' && committed) || (status == 'I' && next == 'I' && wrote)
			status, committed, wrote = next, false, false
			if point && p.commitPoint() {
				// Hold back this acknowledgement: the transaction is
				// committed, and the service will never hear so. Keep the
				// connection open and silent until the caller has killed
				// the service; closing it would let the service see an
				// error and answer its client first.
				<-p.dropped()
				return
			}
		}
		if _, err := client.Write(append(append([]byte(nil), hdr...), body...)); err != nil {
			return
		}
	}
}

// commitPoint records one commit and reports whether to crash on it.
func (p *Proxy) commitPoint() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.held {
		return true // after the crash point, nothing more gets through
	}
	if p.window < 0 {
		p.bg++
		return false
	}
	p.commits[p.window]++
	if p.crashAt.armed && p.window == p.crashAt.seq && p.commits[p.window] == p.crashAt.k {
		p.held = true
		close(p.crashed)
		return true
	}
	return false
}
