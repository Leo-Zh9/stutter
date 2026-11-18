package explore_test

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Leo-Zh9/stutter/internal/config"
	"github.com/Leo-Zh9/stutter/internal/explore"
	"github.com/Leo-Zh9/stutter/internal/replay"
	"github.com/Leo-Zh9/stutter/internal/testenv"
	"github.com/Leo-Zh9/stutter/internal/trace"
)

// agentRunner prepares miniagent (FastAPI + LangGraph + PostgresSaver) for
// one variant. It needs targets/miniagent/.venv (make miniagent); without it
// the test is skipped unless STUTTER_REQUIRE_AGENT=1.
func agentRunner(t *testing.T, bug string) (*explore.Runner, *config.Config) {
	t.Helper()
	dir := filepath.Join(testenv.Root(), "targets", "miniagent")
	if _, err := os.Stat(filepath.Join(dir, ".venv", "bin", "python")); err != nil {
		if os.Getenv("STUTTER_REQUIRE_AGENT") == "1" {
			t.Fatal("targets/miniagent/.venv missing: run make miniagent")
		}
		t.Skip("targets/miniagent/.venv missing: run make miniagent")
	}
	env := testenv.Env(t)
	cfg, err := config.Load(filepath.Join(dir, "stutter.json"))
	if err != nil {
		t.Fatal(err)
	}
	var extra []string
	if bug != "" {
		extra = []string{"BUG=" + bug}
	}
	name := "tpl_test_agent_" + strings.ReplaceAll(bug, "-", "_")
	if err := cfg.PrepareTemplate(context.Background(), env, name, extra); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { env.Drop(context.Background(), name) })
	tr, err := trace.Read(cfg.Path(cfg.Trace))
	if err != nil {
		t.Fatal(err)
	}
	r := &explore.Runner{Cfg: cfg.Replay(env, name, extra), Trace: tr, Lanes: 6, Seed: 1}
	if err := r.CheckStable(context.Background()); err != nil {
		t.Fatal(err)
	}
	return r, cfg
}

// The correct agent survives every lost and concurrent fault, and every
// crash point of both approvals, with every fault firing.
func TestAgentCorrectHasNoFalseAlarms(t *testing.T) {
	r, cfg := agentRunner(t, "")
	ctx := context.Background()
	cs, err := r.CrashFaults(ctx, cfg.KeyHeader)
	if err != nil {
		t.Fatal(err)
	}
	sp := cfg.Space()
	sp.Pending = cs.Pending
	singles := replay.Faults(r.Trace, sp)
	if len(singles) != 12 {
		t.Fatalf("want lost and concurrent on all 4 chats and both approvals, got %v", singles)
	}
	for _, f := range cs.Faults {
		if f.Seq == 1 || f.Seq == 5 {
			singles = append(singles, f)
		}
	}
	outs := r.Parallel(len(singles), func(i int) explore.Outcome {
		return r.Check(ctx, replay.Full(r.Trace, singles[i]), int64(i))
	})
	for _, o := range outs {
		switch {
		case o.Err != "":
			t.Fatal(o.Err)
		case o.Diverged || !o.Fired:
			t.Errorf("%v did not fire (diverged=%v)", o.Schedule.Faults, o.Diverged)
		case o.Failed:
			t.Errorf("false alarm on %v: %v", o.Schedule.Faults, o.Diff)
		}
	}
}

// A double-clicked approval refunds twice when the approval is checked,
// then acted on; the barrier makes it happen every time.
func TestAgentApprovalRaceIsDetected(t *testing.T) {
	r, _ := agentRunner(t, "approval-no-claim")
	s := replay.Full(r.Trace, replay.Fault{Type: replay.Concurrent, Seq: 1})
	for i := 0; i < 3; i++ {
		o := r.Check(context.Background(), s, int64(i))
		if !o.Fired || !o.Failed {
			t.Fatalf("run %d: fired=%v failed=%v", i, o.Fired, o.Failed)
		}
	}
	min, _ := r.Shrink(context.Background(), s, 3)
	if len(min.Seqs) != 2 || len(min.Faults) != 1 {
		t.Fatalf("want chat + approve + concurrent, got %+v", min)
	}
}

// A refund claimed in one commit and written in the next is lost when the
// service dies between them: some crash point of the approval drops it.
func TestAgentCrashBetweenClaimAndRefundIsDetected(t *testing.T) {
	r, cfg := agentRunner(t, "claim-before-refund")
	ctx := context.Background()
	cs, err := r.CrashFaults(ctx, cfg.KeyHeader)
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, f := range cs.Faults {
		if f.Seq != 1 {
			continue
		}
		if o := r.Check(ctx, replay.Full(r.Trace, f), 1); o.Failed {
			found = true
			for _, d := range o.Diff {
				if d.Table == "refunds" && d.Sign == "-" {
					return
				}
			}
		}
	}
	t.Fatalf("no crash point of the approval lost the refund (any failure: %v)", found)
}
