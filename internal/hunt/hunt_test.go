package hunt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/Leo-Zh9/stutter/internal/config"
	"github.com/Leo-Zh9/stutter/internal/replay"
	"github.com/Leo-Zh9/stutter/internal/trace"
)

// The generated repro.sh must send the faulted order with each request's own
// method, and sign Stripe deliveries so a real verifier accepts them.
func TestReproScript(t *testing.T) {
	for _, signing := range []string{"stripe", ""} {
		dir := t.TempDir()
		cfgPath := filepath.Join(dir, "stutter.json")
		body := `{"name":"t","start":["x"],"prepare":{"schema_file":"s.sql"},"webhooks":{"prefix":"/hook","signing":"` + signing + `","secret":"whsec_t"}}`
		os.WriteFile(cfgPath, []byte(body), 0o644)
		cfg, err := config.Load(cfgPath)
		if err != nil {
			t.Fatal(err)
		}
		tr := []trace.Entry{
			{Seq: 0, Kind: trace.Webhook, Method: "POST", Path: "/hook", Body: []byte(`{"id":"evt_a"}`)},
			{Seq: 1, Kind: trace.Webhook, Method: "PUT", Path: "/hook", Body: []byte(`{"id":"evt_b","note":"100% \"quoted\""}`)},
		}
		s := replay.Schedule{Seqs: []int{0, 1}, Faults: []replay.Fault{{Type: replay.Reorder, Seq: 0, Seq2: 1}}}
		if err := WriteRepro(filepath.Join(dir, "f"), cfg, tr, s); err != nil {
			t.Fatal(err)
		}
		var mu sync.Mutex
		var got []string
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			b, _ := io.ReadAll(r.Body)
			if signing == "stripe" {
				parts := strings.Split(r.Header.Get("Stripe-Signature"), ",")
				m := hmac.New(sha256.New, []byte("whsec_t"))
				m.Write([]byte(strings.TrimPrefix(parts[0], "t=") + "." + string(b)))
				if len(parts) != 2 || parts[1] != "v1="+hex.EncodeToString(m.Sum(nil)) {
					w.WriteHeader(400)
					return
				}
			}
			mu.Lock()
			got = append(got, r.Method+" "+string(b))
			mu.Unlock()
		}))
		cmd := exec.Command("bash", filepath.Join(dir, "f", "repro.sh"), "faulted")
		cmd.Env = append(os.Environ(), "BASE_URL="+srv.URL)
		out, err := cmd.CombinedOutput()
		srv.Close()
		if err != nil {
			t.Fatalf("signing=%q: %v\n%s", signing, err, out)
		}
		want := []string{`PUT {"id":"evt_b","note":"100% \"quoted\""}`, `POST {"id":"evt_a"}`}
		if strings.Join(got, "|") != strings.Join(want, "|") {
			t.Fatalf("signing=%q: got %q\nscript output:\n%s", signing, got, out)
		}
	}
}
