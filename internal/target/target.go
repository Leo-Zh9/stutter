// Package target starts and stops the service under test as a subprocess.
package target

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

// Spec says how to launch the service. Command and Env may use these
// placeholders, filled per run:
//
//	{DSN}           libpq keyword string (unix socket)
//	{DATABASE_URL}  postgres:// URL over TCP
//	{ADDR}          127.0.0.1:<port> the service must listen on
//	{HOST} {PORT}   the two halves of {ADDR}
//
// The service must answer GET HealthPath with a status below 500 once ready.
type Spec struct {
	Command      []string
	Env          []string
	Dir          string
	HealthPath   string
	StartTimeout time.Duration
}

// DB is the database a run hands to the service.
type DB struct {
	DSN string
	URL string
}

type Process struct {
	BaseURL string
	cmd     *exec.Cmd
	exited  chan struct{}
}

func freePort() (string, error) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return "", err
	}
	defer l.Close()
	_, port, _ := net.SplitHostPort(l.Addr().String())
	return port, nil
}

func fill(s string, r *strings.Replacer) string { return r.Replace(s) }

func Start(ctx context.Context, s Spec, db DB) (*Process, error) {
	timeout := s.StartTimeout
	if timeout == 0 {
		timeout = 5 * time.Second
	}
	var lastErr error
	for attempt := 0; attempt < 3; attempt++ { // a freed port can be taken before the service binds it
		port, err := freePort()
		if err != nil {
			return nil, err
		}
		addr := "127.0.0.1:" + port
		r := strings.NewReplacer("{DSN}", db.DSN, "{DATABASE_URL}", db.URL, "{ADDR}", addr, "{HOST}", "127.0.0.1", "{PORT}", port)
		args := make([]string, len(s.Command))
		for i, a := range s.Command {
			args[i] = fill(a, r)
		}
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Dir = s.Dir
		cmd.Env = os.Environ()
		for _, e := range s.Env {
			cmd.Env = append(cmd.Env, fill(e, r))
		}
		// Own process group, so a crash or stop reaches every worker the
		// service forks (node, next, gunicorn).
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
		var stderr strings.Builder
		cmd.Stderr = &stderr
		if err := cmd.Start(); err != nil {
			return nil, err
		}
		p := &Process{BaseURL: "http://" + addr, cmd: cmd, exited: make(chan struct{})}
		go func() { cmd.Wait(); close(p.exited) }()
		deadline := time.Now().Add(timeout)
		healthy := false
		for time.Now().Before(deadline) && !healthy {
			select {
			case <-p.exited:
				lastErr = fmt.Errorf("target exited during startup: %s", tail(stderr.String()))
				deadline = time.Time{}
				continue
			default:
			}
			if resp, err := http.Get(p.BaseURL + s.HealthPath); err == nil {
				resp.Body.Close()
				healthy = resp.StatusCode < 500
			}
			if !healthy {
				time.Sleep(2 * time.Millisecond)
			}
		}
		if healthy {
			return p, nil
		}
		p.Stop()
		if lastErr == nil {
			lastErr = fmt.Errorf("target not healthy after %s: %s", timeout, tail(stderr.String()))
		}
	}
	return nil, lastErr
}

func tail(s string) string {
	if len(s) > 2000 {
		return "..." + s[len(s)-2000:]
	}
	return s
}

// Stop SIGKILLs the whole process group and waits for the leader to exit.
func (p *Process) Stop() {
	if p == nil || p.cmd.Process == nil {
		return
	}
	syscall.Kill(-p.cmd.Process.Pid, syscall.SIGKILL)
	select {
	case <-p.exited:
	case <-time.After(5 * time.Second):
	}
}
