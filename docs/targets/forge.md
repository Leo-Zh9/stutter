# nihalashetty/Forge (agent blind target)

| | |
|---|---|
| Repo | [nihalashetty/Forge](https://github.com/nihalashetty/Forge) |
| Pinned | `25074df07b9872912902b323706a01dd47cd8477`, from the [agent screen](../evidence/agent-screen.md) |
| Code changed | none |
| App run | `apps/api` (FastAPI), `uvicorn forge.main:app`, PostgreSQL app database and PostgreSQL checkpointer, no Redis |
| Config | [stutter.json](../../targets/agents/forge/stutter.json), [fixtures.sql](../../targets/agents/forge/fixtures.sql), [model.json](../../targets/agents/forge/model.json), [workload.py](../../targets/agents/forge/workload.py) |
| Results | [hunt/REPORT.md](../../targets/agents/forge/hunt/REPORT.md), [hunt/results.json](../../targets/agents/forge/hunt/results.json), [trace.jsonl](../../targets/agents/forge/trace.jsonl) |

## Blind protocol

- **Pinned SHA.** No other commit was checked out.
- **Push before any source reading.** Config, fixtures, model script, trace and results were pushed as `de16b2d`. Application source was read only after that push.
- **What was read before the push.** `README.md`, `.env.example`, `apps/api/README.md`, `docs/MANUAL.md`, `apps/api/pyproject.toml` (to install), the list of migration files, the head of `migrations/versions/0001_baseline.py` and of `migrations/env.py`, the workflow and node JSON schemas under `packages/schemas/forge/`, and the served `/openapi.json`. At runtime: HTTP responses, the rows the app wrote to scratch databases, and the tool schemas it sent to the scripted model.

## Setup

- **Install.** Python 3.12 venv in `apps/api`, `uv pip install -e ".[providers,postgres]"`. The database is built by the app's own `alembic upgrade head`.
- **Model.** `OPENAI_BASE_URL` points the app's `openai:gpt-4o-mini` at stutter's scripted model. The project's embedding model is `openai:text-embedding-3-small`, so the `remember` tool's embedding call also goes to the stub. The first attempt used the default local embedder, which failed because `fastembed` is not installed.
- **Fixtures.** One tenant, the bootstrap admin, and a project built through Forge's own API on a fresh database, then dumped: a workflow `start → gate (human input, output "decision") → saver (agent, tool remember) → end`, with the edge from `gate` going to `saver` on `approve` and to `end` on `reject`. The workflow is the project's API workflow.
- **Workload.** Four calls to `POST /v1/projects/{id}/run` with `stream: false`, each with its own `Idempotency-Key`, the retry header `docs/MANUAL.md` documents. `#0` ACME note → pauses at the gate. `#1` resume `approve` → the agent calls `remember`, which writes one `memories` row. `#2` Globex note → pauses. `#3` resume `reject` → the run ends with no write.
- **Oracle.** Checkpoint tables, `alembic_version`, and the telemetry tables `traces`, `spans`, `audit_logs` are ignored. These were set before the first hunt. The first hunt stopped at the stability check on `memories.id`, `runs.id`, `runs.thread_id`, `threads.id` and `threads.lg_thread_id`, which the app mints with `uuid4()` in Python. The four ids were declared with the new `generated` option, added to stutter for this run. `runs.thread_id` then becomes an edge.
- **Faults.** `-lanes 3`, crash points included. The hunt took 4548 s.

## Results

| Fault | Schedules | Fired | Not fired | Changed the final state |
|---|---|---|---|---|
| lost | 4 | 4 | 0 | 0 |
| concurrent | 4 | 0 | 4 | 0 |
| crash after commit | 47 | 47 | 0 | 42 |

42 findings, 8 distinct end-state changes. Every finding is 1-minimal and reproduced in 10 of 10 fresh replays. Crash findings cannot be scripted with curl.

`concurrent` never fired: the duplicate never reached its first write within 500 ms. Concurrency was not tested on this app.

## Effects

**A retried new turn starts a second run** (findings 1–17 and 36; effects 1–5). A crash during `#0` (after any of commits 1–9) or `#2` (commits 2–10), then the client's retry with the same `Idempotency-Key`, leaves one extra `threads` row and one extra `runs` row. The extra run is stuck at `queued` (commit 1), `running` (most commits) or `interrupted` (the last commit). In the `interrupted` case the one request has left two approvals pending. The workload approves the second, and the first is never answered.

**A retried approval is refused, and the approved action is dropped** (findings 18–26; effect 6). A crash during `#1` (after commits 2–10), then the retry, ends with no `memories` row. The run stays at `running` and never reaches `done`.

**A retried resume leaves the run stuck at `running`** (findings 27–35 and 37–42; effects 7–8). If the crash comes after the memory is written (`#1`, commits 11–19), the memory exists but the run stays `running`. On the rejected path (`#3`, commits 2–7) nothing is written either way, and the run stays `running` instead of `done`.

## Root causes

Read after the push.

### 1. The `Idempotency-Key` is remembered only in process memory, and only after the run completes

[`forge/routers/project_run.py`](https://github.com/nihalashetty/Forge/blob/25074df07b9872912902b323706a01dd47cd8477/apps/api/forge/routers/project_run.py) looks up the key in `idempotency.get(ik)` before creating the run. It stores the result with `idempotency.put(ik, result)` only after `run_to_completion` returns. Without `FORGE_REDIS_URL`, `idempotency` is the process-local `IdempotencyCache` in `forge/util/ratelimit.py`, so a restart forgets every key. With Redis, the key survives a restart, but nothing is stored while the run is in flight. A retry after a crash at any point before `put` therefore creates a new `Thread` and `Run`. The first run is left behind. The code's comment states the intent, "so the side-effecting run never fires twice", and `docs/MANUAL.md` defines the header as one that keeps "a retried request" from running twice.

In this workload the first run had not yet reached a tool when it was left behind, because the gate comes first. In a workflow whose first tool comes before its first gate, the same retry runs that tool a second time. That is inferred from the code and was not observed.

**Fix.** Store the key in PostgreSQL in the same transaction that inserts the `Run`, with a unique `(tenant_id, project_id, key) → run_id`. On a retry, return the stored result if the run finished, return its interrupt if it is `interrupted`, and answer 409 "in progress" if it is still running, instead of creating a new run.

### 2. A resume flips the run to `running` before acting, and a retry only accepts `interrupted`

`RunService.resume` in [`forge/services/runs.py`](https://github.com/nihalashetty/Forge/blob/25074df07b9872912902b323706a01dd47cd8477/apps/api/forge/services/runs.py) sets `run.status = "running"` and commits before it calls `graph.ainvoke(Command(resume=value))`. The router's resume path looks for a run with `Run.status == "interrupted"`. After a crash, the run is `running` with no live driver, so the client's retry gets `409 no interrupted run to resume on this thread`. Nothing else re-drives the run. The reaper (`reap_stale_runs`, every 300 s from `forge/main.py`) marks a `running` run `error` after 3600 s. The human's approval is then lost for good, although the LangGraph checkpoint still holds the paused graph.

Resume requests are also excluded from the idempotency key (`body.resume is None` in the key condition), so the key does not help here.

**Fix.** Make the resume retryable. Let the resume path accept a `running` run whose driver is gone: the thread lock is free and the run is older than its lease. Re-drive it from the checkpoint: resume with the decision if the interrupt is still pending, or `ainvoke(None)` to finish if the graph had already passed the gate. Or store the decision with the status change in one commit, and re-drive `running` runs at startup.

## Issue drafts

Not filed. This project does not file upstream.

> **Retry with the same Idempotency-Key after a crash starts a second run**
>
> `POST /v1/projects/{id}/run` with `stream: false` and an `Idempotency-Key` checks the key in `idempotency` before creating the run, but stores it only after `run_to_completion` returns, and by default in process memory. If the API process dies while the run is in flight (deploy, OOM), the client's retry with the same key creates a new thread and run. The first run stays `queued`, `running` or `interrupted`. With a human-input node, the one request now has two pending approvals, and any tool before the gate runs twice. A crash is not needed: a retry sent while the first run is still in flight gets 200 with a second `run_id`.
>
> Reproduced against 25074df with PostgreSQL app and checkpoint databases: kill the process after the first request's run commits, restart it, and resend the request with the same key. The `runs` table now holds two runs for the one request.
>
> Suggested fix: persist `(tenant, project, key) → run_id` in the transaction that inserts the run, and answer retries from that row.

> **A crash during a HITL resume makes the approval impossible to retry**
>
> `RunService.resume` commits `status = "running"` before invoking the graph. If the process dies after that commit, the retry of the same resume request gets `409 no interrupted run to resume on this thread`, because the router only resumes `interrupted` runs. The reaper later marks the run `error`. An approved action never runs, even though the checkpoint still holds the paused graph.
>
> Reproduced against 25074df: a workflow `human_input → agent(remember)`, approve, kill the process right after the resume's first commits, restart and retry. No memory row is written, and the run stays `running`.
>
> Suggested fix: let a resume take over a `running` run whose driver is gone, and re-drive it from the checkpoint.

## Counted under the protocol

Two distinct root causes. Cause 2 drops a business row: the approved `memories` row is never written. Cause 1 duplicates rows in `runs` and `threads` and leaves a second pending approval. In this workload the duplicated run did not repeat a tool call. The claim's wording ("a tool's side effect ran twice or never") fits cause 2 directly. Cause 1 fits it only through the run and approval it duplicates. [agent-results.md](../evidence/agent-results.md) gives the count both ways.

## Held-commit re-run (after the blind phase)

Not part of the blind protocol. After the blind results were pushed, `concurrent(#n@commitK)` was added to stutter ([faults.md](../faults.md#holding-at-a-commit)) because `concurrent` had never fired here. The same pinned commit, config (plus `"holds": true`), fixtures, model script and trace were then hunted again with `-lanes 3`. The owner's go-ahead to fix concurrency covered this re-run. The source had already been read by then. Results: [hunt-holds/REPORT.md](../../targets/agents/forge/hunt-holds/REPORT.md), [hunt-holds/results.json](../../targets/agents/forge/hunt-holds/results.json). The hunt took 3785 s.

| Fault | Schedules | Fired | Not fired | Nondetermined | Changed the final state |
|---|---|---|---|---|---|
| lost | 4 | 4 | 0 | 0 | 0 |
| concurrent | 4 | 0 | 4 | 0 | 0 |
| concurrent@commit | 47 | 31 | 15 | 1 | 4 |
| crash after commit | 47 | 47 | 0 | 0 | 42 |

The lost, concurrent and crash rows match the blind hunt: the same 42 crash findings, and the same requests and counts.

**Why the barrier never fired.** For a new turn, `run_admission` (`forge/services/quota.py`) takes a per-tenant `asyncio` lock around the quota check and `create_run`. A resume takes the per-thread lock. The barrier holds A at its first write, inside that lock. B makes its reads, then waits on the lock, so it never reaches a write within 500 ms. Once the barrier timed out, B ran anyway. On `#0` and `#2` it left an extra run and thread, but those runs counted as not fired and were thrown away.

**What the held commit found.** Four findings (18, 19, 45, 46), all 1-minimal and reproduced in 10 of 10 replays: `concurrent(#0@commit2)`, `#0@commit9`, `#2@commit3` and `#2@commit10`. Each leaves one extra `runs` row at `interrupted` and one extra `threads` row: the effect the barrier threw away. No crash is involved. While the first attempt is still in flight, the duplicate with the same `Idempotency-Key` gets 200 with a different `run_id` and its own pending approval. The process never restarts, so this would happen with Redis too.

**Root cause.** Cause 1 above, not a new one. The key is stored only after `run_to_completion` returns, so nothing answers a duplicate while the run is in flight. The blind write-up inferred this from the code; this re-run observed it.

**Where it did not fire.** On `#0` at commits 1 and 3–8, and on `#2` at commits 2 and 4–9, the duplicate got no response within the hold timeout. That fits the per-tenant lock around commit 1, and LangGraph's checkpoint writes going through one shared saver connection, which the held commit blocks. Neither was traced further. `#2@commit1` was not on A's path: A answered while it was held (one of the 22 background commits seen while idle). On the resumes `#1` and `#3`, every held commit that fired passed. In the one point probed, `#1@commit5`, the duplicate got `409 no interrupted run to resume on this thread`. One held commit was nondetermined. The report counts it but does not name it.

The duplicates' responses quoted here (`200` with a new `run_id`, `409`) came from `stutter check` runs of a throwaway debug build that printed B's response. That build was not committed.

**Counting.** M is unchanged. The issue draft for cause 1 now says a crash is not needed.
