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
	"time"
)

// Spec says how to launch the service. The service must read DATABASE_URL
// and ADDR from its environment and answer GET HealthPath once ready.
type Spec struct {
	Command    []string
	Env        []string
	HealthPath string
}

type Process struct {
	BaseURL string
	cmd     *exec.Cmd
}

func freeAddr() (string, error) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return "", err
	}
	defer l.Close()
	return l.Addr().String(), nil
}

func Start(ctx context.Context, s Spec, dsn string) (*Process, error) {
	var lastErr error
	for attempt := 0; attempt < 3; attempt++ { // a freed port can be taken before we bind it
		addr, err := freeAddr()
		if err != nil {
			return nil, err
		}
		cmd := exec.Command(s.Command[0], s.Command[1:]...)
		cmd.Env = append(append(os.Environ(), s.Env...), "DATABASE_URL="+dsn, "ADDR="+addr)
		var stderr strings.Builder
		cmd.Stderr = &stderr
		if err := cmd.Start(); err != nil {
			return nil, err
		}
		p := &Process{BaseURL: "http://" + addr, cmd: cmd}
		exited := make(chan struct{})
		go func() { cmd.Wait(); close(exited) }()
		deadline := time.Now().Add(5 * time.Second)
		for time.Now().Before(deadline) {
			select {
			case <-exited:
				lastErr = fmt.Errorf("target exited during startup: %s", stderr.String())
				goto next
			default:
			}
			if resp, err := http.Get(p.BaseURL + s.HealthPath); err == nil {
				resp.Body.Close()
				if resp.StatusCode < 500 {
					p.cmd = cmd
					return p, nil
				}
			}
			time.Sleep(2 * time.Millisecond)
		}
		cmd.Process.Kill()
		lastErr = fmt.Errorf("target not healthy after 5s: %s", stderr.String())
	next:
	}
	return nil, lastErr
}

func (p *Process) Stop() {
	if p == nil || p.cmd.Process == nil {
		return
	}
	p.cmd.Process.Kill()
}
