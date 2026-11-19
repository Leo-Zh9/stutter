# langchain-ai/paid-media-agent (agent blind target)

| | |
|---|---|
| Repo | [langchain-ai/paid-media-agent](https://github.com/langchain-ai/paid-media-agent) |
| Pinned | `d4ee460b729ae6a9a7c0fb0e84e7ab07756fd5ee`, from the [agent screen](../evidence/agent-screen.md) |
| Code changed | none |
| App run | `paid-media-agent serve` (self-hosted runtime, HTTP API), PostgreSQL app tables and checkpointer, `PAID_MEDIA_DATA_MODE=sample` |
| Config | [stutter.json](../../targets/agents/paid-media-agent/stutter.json), [model.json](../../targets/agents/paid-media-agent/model.json), [workload.py](../../targets/agents/paid-media-agent/workload.py) |
| Results | [hunt/REPORT.md](../../targets/agents/paid-media-agent/hunt/REPORT.md), [hunt/results.json](../../targets/agents/paid-media-agent/hunt/results.json), [trace.jsonl](../../targets/agents/paid-media-agent/trace.jsonl) |

## Blind protocol

- **Pinned SHA.** No other commit was checked out.
- **Push before any source reading.** Config, model script, trace and results were pushed as `de16b2d`. Application source was read only after that push.
- **What was read before the push.** `README.md`, `.env.example`, `docs/self-hosting.md`, `docs/architecture/writes-and-approvals.md`, `docker-compose.yml`, `pyproject.toml` (to install), `config/write-policy.example.toml`, `config/accounts.example.toml`, the CLI's `--help`, and a list of the repo's file names. At runtime: HTTP responses, the rows the app wrote to scratch databases, the tool schemas it sent to the scripted model, and the server's log.
- **One disclosure.** While a first hunt was being debugged, the server log showed a Python traceback for a duplicate reject (`InvalidTransition: cannot apply reject in state rejected`), with source file paths and line numbers in its frames. That is runtime output, not source reading, but it named code locations before the push.

## Setup

- **Install.** `uv sync --extra self-host --frozen`.
- **Model and data.** `PAID_MEDIA_MODEL=openai:gpt-4o-mini` with `PAID_MEDIA_MODEL_BASE_URL` set to stutter's scripted model. Sample data mode, so the ad platforms are the app's own in-process fixture provider. No ad platform was called. Static bearer token `PAID_MEDIA_API_TOKENS`. Self-approval allowed so one operator can propose and approve.
- **Workload.** Three conversations through `POST /threads/{id}/messages`. In each, the model lists campaigns, calls `propose_change`, then `execute_change`, which pauses for approval. Then `POST /proposals/{id}/approve` (Google budget 180 → 150), `reject` (Meta pause) and `approve` (Reddit budget 75 → 90).
- **Retries.** The message endpoint takes no idempotency key and has no retry rule, so only the approve and reject requests (`#1`, `#3`, `#5`) are retried. Each one resumes a LangGraph thread found from the proposal id through `model.interrupt_sql`.
- **Oracle.** Checkpoint tables are ignored. The app mints `proposal_id`, `routing_id` and `claim_id` with `uuid4()` in Python. Approval claims carry a random nonce and a signature, and both records carry a `payload_digest` over content that includes the proposal id. These were handled with two options added to stutter during this run: `generated` for app-minted ids, and `ignore` paths inside JSON columns. `record.history` (timestamped log lines) is ignored too.
- **Mid-run fix in stutter.** The first full hunt aborted with `connection reset by peer`. After a 500, the app closed the kept-alive connection the replayer reused. The replayer now opens a fresh connection for each request.
- **Faults.** `-lanes 4`, crash points included. The hunt took 4713 s.

## Results

| Fault | Schedules | Fired | Not fired | Nondetermined | Changed the final state |
|---|---|---|---|---|---|
| lost | 3 | 3 | 0 | 0 | 0 |
| concurrent | 3 | 0 | 3 | 0 | 0 |
| crash after commit | 126 | 119 | 0 | 7 | 92 |

92 findings, 9 distinct end-state changes. Every finding is 1-minimal and reproduced in 10 of 10 fresh replays. `concurrent` never fired, so concurrency was not tested on this app.

## Effects and root causes

Source was read after the push. Paths are under `src/paid_media_agent/`.

The approve path commits in this order. Reads count, because each repository call is its own transaction. First the proposal reads, then the history update, then the claim insert (commit 4). Next the thread-owner claim, the checkpoint's resume write, the approval node's reads and its checkpoint writes. After that, five reads in the executor, `executing`, the claim marked used, the provider call, `verifying` and `verified`. The receipt insert comes last.

### 1. A retried approval is answered but never executed (counted)

**Effect.** Findings 7–12 and 14 (`#1`, commits 10–15 and 17) and 24–31 (`#5`, commits 10–17). The proposal ends `awaiting_approval` with no `pma_receipts` row, where the fault-free run ends `verified` with one. The budget change was never sent: the `executing` transition, which comes before the provider call, never committed.

**Cause.** `Runner.resume` (`surfaces/runner.py`) reads the thread's snapshot and returns early if `snapshot.interrupts` is empty. Once the approval node's writes reach the checkpoint, the interrupt counts as answered, but the `execute_change` tool task is still pending in `snapshot.next`. The crash lands between the two. The retry then finds no interrupt, returns 200, and runs nothing. Nothing else ever runs the pending task. The proposal stays `awaiting_approval`, and each further approve mints another claim and still does nothing.

**Fix.** In `resume`, if no interrupt is pending but `snapshot.next` is not empty, call `ainvoke(None, config)` to finish the pending task. The executor's state checks and single-use claim already stop it from running twice.

### 2. The terminal state is saved before the receipt (counted)

**Effect.** Finding 17 (`#1`, commit 21). The proposal ends `verified`, as in the fault-free run, but no `pma_receipts` row exists. The change was applied and verified, and its receipt is lost.

**Cause.** In `tools/writes.py`, every terminal path saves the proposal's new state, then writes the receipt in a separate transaction. For example, `mark(VERIFIED)` comes first and `_receipt(...)` second. A crash between the two leaves a terminal proposal with no receipt. `GET /proposals/{id}` then returns no receipt. A repeated execution is refused with `not_awaiting_approval`, which contradicts `docs/architecture/writes-and-approvals.md`: "Repeated execution returns the existing receipt".

**Fix.** Write the receipt and the terminal state in one transaction, or write the receipt first. The receipt upsert is already idempotent.

### 3. Approve mints a new claim on every call (not counted)

**Effect.** Findings 1–6 (`#1`, commits 4–9) and 18–23 (`#5`, commits 4–9): one extra `pma_approvals` row. The proposal still ends `verified` with one receipt. Effect 1 above also carries the extra claim.

**Cause.** `ProposalService.approve` creates and saves a new signed claim whenever the proposal is `awaiting_approval`. The executor uses the newest unused claim, and the older one is never read again.

**Why not counted.** The extra row changes nothing the app does. That is the protocol's "re-run with no app-visible effect". The fix would be to return the existing unused claim for the same proposal and revision.

### 4. Crash during execution leaves `executing` or `verifying` (not counted: documented design)

**Effect.** Findings 13, 15, 16, 32 (`executing`) and 33 (`verifying`). There is no receipt. The retried approve gets 403 `not_awaiting_approval`.

**Why not counted.** The docs say this is intended: "If a process stops during execution, inspect the provider through read-only tools before creating a new proposal. The agent does not automatically retry an interrupted mutation." The app chooses to stop rather than risk a second mutation, so there is no app-level fix to propose. With a real provider, `verifying` means the change is live but unrecorded until an operator reconciles it.

### 5. The fixture provider's operation counter (not counted: harness artifact)

**Effect.** Findings 34–92: the Reddit receipt's `provider_operation_ref` is `fixture-op-1` instead of `fixture-op-2`.

**Cause.** The sample-mode `FakeWriteProvider` numbers operations from an in-memory list. stutter's restart resets it, so any crash after the Google mutation and before the Reddit one shifts the number. Nothing in the app is wrong.

### Also seen, not a finding

A duplicate reject returns 500: `reject` calls `transition()`, which raises `InvalidTransition`, and the API only maps `WriteDenied`. The business tables end the same as in the fault-free run, so the hunt did not flag it.

## Issue drafts

Not filed. This project does not file upstream.

> **An approval retried after a crash returns 200 but never executes the change**
>
> `Runner.resume` returns early when `snapshot.interrupts` is empty. If the process dies after the approval node's checkpoint writes but before the `execute_change` task runs, the retried `POST /proposals/{id}/approve` mints a new claim, finds no interrupt, and returns without running the pending task. The proposal stays `awaiting_approval` for good, and no change or receipt is ever made.
>
> Reproduced against d4ee460 with the self-hosted runtime on PostgreSQL in sample mode: kill the process after the approve request's tenth database commit, restart, and resend the approve.
>
> Suggested fix: if `snapshot.next` is not empty, call `ainvoke(None, config)` to finish the pending task.

> **A crash between the terminal state and the receipt loses the receipt**
>
> `ProposalExecutor` saves `verified` (or `failed`, `unknown`, ...) and then writes the receipt in a separate transaction. A crash in between leaves a verified proposal with no receipt. Repeated execution is then refused rather than returning the receipt, unlike what `writes-and-approvals.md` says.
>
> Suggested fix: save the receipt and the terminal state in one transaction.

## Counted under the protocol

Two distinct root causes. Cause 1 drops a row, and the approved change never runs. Cause 2 drops the receipt row, though the change itself ran once. The claim's wording ("a tool's side effect ran twice or never") fits cause 1 directly, and fits cause 2 only through the record it loses. [agent-results.md](../evidence/agent-results.md) gives the count both ways.

## Held-commit re-run (after the blind phase)

Not part of the blind protocol. After the blind results were pushed, `concurrent(#n@commitK)` was added to stutter ([faults.md](../faults.md#holding-at-a-commit)) because `concurrent` had never fired here. The same pinned commit, config (plus `"holds": true`), model script and trace were then hunted again with `-lanes 4`. The owner's go-ahead to fix concurrency covered this re-run. The source had already been read by then. Results: [hunt-holds/REPORT.md](../../targets/agents/paid-media-agent/hunt-holds/REPORT.md), [hunt-holds/results.json](../../targets/agents/paid-media-agent/hunt-holds/results.json). The hunt took 3916 s.

| Fault | Schedules | Fired | Not fired | Nondetermined | Changed the final state |
|---|---|---|---|---|---|
| lost | 3 | 3 | 0 | 0 | 0 |
| concurrent | 3 | 0 | 3 | 0 | 0 |
| concurrent@commit | 126 | 71 | 53 | 2 | 0 |
| crash after commit | 126 | 116 | 0 | 10 | 89 |

**Crash row versus the blind hunt.** Four points that failed in the blind hunt did not fail this time: `crash(#1@commit10)`, `#1@commit11`, `#1@commit19` and `#5@commit17`. One that had not failed before did: `#5@commit21`. Re-run alone with the new build (`stutter check -n 4`), all five failed 4 of 4. So this is run-to-run variance on requests whose checkpoint writes race in the background, not a change in the crash fault. Both counted causes are still found: cause 1 at `#1` commits 12–15 and 17 and at `#5` commits 10–16; cause 2 at `#1` commit 21. There are 9 distinct end-state changes, as before.

**Why neither concurrency fault reaches the window that matters.** `async def approve` calls the sync `ProposalService.approve` on a sync `psycopg_pool.ConnectionPool`. The executor's repository calls are sync too. Each one blocks the event loop of the single uvicorn worker. While A is inside them, no other request runs.

- The barrier holds A at its first write, inside such a call, so B is never served. That is why it never fired. Once it timed out, B ran later and minted an extra claim (`+ pma_approvals`), which is the not-counted cause 3.
- The held commit did not fire at any sync commit: 53 points, all `no response within HoldTimeout`. These are `#1` commits 1–8, 12–22, 24, 46 and 47; `#3` commits 1–7, 31 and 32; and `#5` commits 1–8, 12–18, 20–24, 46 and 47.
- It fired at 71 checkpoint commits, and all passed. But those commits are LangGraph's background checkpoint writes, which A does not wait for. A kept running while one was held. At the two approve points probed (`#1@commit9`, `#1@commit30`), the duplicate got `403 not_awaiting_approval: verified`: by the time B ran, A had already executed and verified the change. At `#3@commit10`, the duplicate reject got 500, the unmapped `InvalidTransition` already noted under "Also seen".

So a duplicate approve cannot run between A's claim and its `executing` transition in one worker process, under either fault. Whether two workers or replicas could execute twice in that window was not tested. Two held commits were nondetermined. The report counts them but does not name them.

The duplicates' responses quoted here came from `stutter check` runs of a throwaway debug build that printed B's response. That build was not committed.

**Counting.** No new finding. M is unchanged.
