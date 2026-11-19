# Faults and client rules

Faults are injected in the replayer's own transport, the only place where a lost response is observable. Code: [internal/replay/replay.go](../internal/replay/replay.go).

| Fault | What happens | Real-world cause |
|---|---|---|
| `lost(#n)` | Request n reaches the service and completes. The sender reads the response, discards it, and retries. | Timeout or dropped connection after the server committed |
| `concurrent(#n)` | A duplicate of request n is sent while the first attempt is still inside its transaction. | Client timeout and retry while the first attempt is slow |
| `reorder(#a,#b)` | Two webhook deliveries about the same object arrive swapped. | Providers do not guarantee delivery order |
| `crash(#n@commitK)` | The service is killed right after the K-th commit of request n, before it learns the commit succeeded. It restarts on the same database and the client retries. | Process crash or deploy between two of a handler's own transactions |
| `concurrent(#n@commitK)` | The acknowledgement of the K-th commit of request n is held while a duplicate is sent; it is released once the duplicate has answered. Opt-in (`"holds": true`). | Client timeout and retry while the first attempt waits between two awaits of an async handler |

`reorder` applies to pairs of webhooks about the same object. The object is either named by `webhooks.object_path` in the config (for Stripe, `data.object.customer`), or taken as the first id minted earlier in the trace.

## Crash points

A black-box tool cannot place a crash inside a service, but it can see every commit acknowledgement before the service does. For crash schedules, the service's database URL points at a wire-protocol proxy ([internal/pgproxy](../internal/pgproxy/pgproxy.go)). The proxy forwards bytes unchanged and parses only the backend stream.

**What counts as a commit point.** A `ReadyForQuery` counts when it ends a committed write, and each one counts at most once:

- an explicit transaction: status goes from `T` to `I` with `COMMIT` reported since the previous `ReadyForQuery`;
- an autocommit or implicit transaction: status stays `I`, and a write tag (`INSERT`, `UPDATE`, `DELETE`, `MERGE`, `COPY`) was reported since the previous `ReadyForQuery`.

A failed `COMMIT` reports `ROLLBACK`, so it is not a commit point. Every `Sync` produces exactly one `ReadyForQuery`, in order, so this holds for the simple, extended and pipelined protocols alike.

**How a crash runs:**

1. Two observed fault-free runs count the commit points of every request. Requests whose count differs between the two are listed, and never faulted.
2. For `crash(#n@commitK)`, the proxy holds back the `ReadyForQuery` of commit K of request n. The commit is durable, but the service never hears it succeeded, and its connection is kept open and silent so it cannot react.
3. stutter SIGKILLs the service's process group, drops the proxied connections, restarts the service on the same database, and retries the request under the normal client rules.
4. If the request finishes with fewer than K commits, the fault did not fire.

**Not seen as commit points:** `COMMIT AND CHAIN` (status stays `T`), procedures that commit inside a `CALL`, and writes inside a function called from a plain `SELECT`.

## The lock barrier

Sending a duplicate at the same moment rarely lands inside the race window, because a handler's read and its write are microseconds apart. For `concurrent`, the replayer:

1. Opens a transaction holding `LOCK TABLE ... IN EXCLUSIVE MODE` on every table in the configured schemas. That mode lets plain `SELECT`s through but blocks every write.
2. Sends attempt A and waits until its backend appears in `pg_stat_activity` with `wait_event_type = 'Lock'`. Then it sends attempt B.
3. Waits until both backends are waiting, then commits, which releases the locks.

At that point both attempts have done their reads and neither has written, which is exactly where check-then-act bugs live.

If the barrier does not fill within 500 ms, the fault is recorded as **not fired**. A not-fired run never counts as a pass, and every rate reports fired and not-fired runs separately.

`pg_stat_activity` is polled from a separate connection, because inside a transaction PostgreSQL returns a frozen snapshot of it.

## Holding at a commit

The barrier holds attempt A at its first write. In an async service (FastAPI `async def` routes, one event loop per worker) that is often the wrong place:

- **A sync database call inside an async route freezes the event loop.** If A is held inside it, the whole worker stops, B is never served, and the barrier times out. paid-media-agent calls a sync `psycopg_pool.ConnectionPool` from async routes, so `concurrent` never fired there.
- **An in-process lock around the first write serializes B behind A.** B makes its reads and then waits on the lock A holds (Forge's per-tenant `asyncio` lock). In one worker that is a real guarantee; it says nothing about the gap after the lock.

In both, after the barrier timed out B ran anyway, and on some requests the end state changed. Those runs were discarded as not fired: unjudged, not passes.

`concurrent(#n@commitK)` holds A at a commit point instead, through the same wire proxy as crashes:

1. The two observation runs that count commit points also list a held commit for every commit K of every request that admits a concurrent duplicate and whose count agrees between the runs. The hunt's `HoldTimeout` is three times the slowest observed request, at least 1 s.
2. The proxy delays the `ReadyForQuery` of commit K on A's connection only. The commit is durable; A just has not heard. Every other connection, and every later commit, goes through.
3. When the hold is reached, the resume guard runs (as before a retry), then B is sent. stutter waits up to `HoldTimeout` for B's response.
4. It releases A, waits for B, and keeps A's response, as under the barrier. If A's response is a 5xx, a 409 or a transport error, A's own retry policy continues.

It **fires** only if the hold was reached, B got an HTTP response, and A had not answered by then. Otherwise it is not fired, with one of four reasons, reported per fault in `not_fired_reasons` and in REPORT.md:

| Reason | Meaning |
|---|---|
| `finished before commit K` | A answered before reaching commit K. |
| `resume guard` | At the hold, a different interrupt was already pending, so B would answer something the user never saw. B is never sent. Also used when nothing was pending before the request. |
| `first attempt was not held` | A answered while its commit K was held: that commit was not on A's path (a background task's write, an async checkpoint). |
| `no response within HoldTimeout` | B got no HTTP response in time: the held commit froze the event loop (a sync call in an async route), or B waits on a lock A holds. |

**Deterministic only where the held commit blocks A's own coroutine.** A commit made by a background task, such as LangGraph's checkpoint writes under the default `async` durability, may or may not be the one held, and A may move on. Those points are best-effort: like crash faults, every held-commit fault that fires is run twice, and two runs that disagree are reported as nondetermined, never as findings.

**Cost and opt-in.** One run per commit point (two when it fires), and a run that cannot serve B takes `HoldTimeout` longer. So it is off unless the config sets `"holds": true`. Without it, hunts run and report exactly as before. A hold is never placed in random schedules or `explore`. Its repro.sh says it needs stutter's wire proxy, like a crash finding.

The fixture `targets/miniasync` shows all four outcomes and the race the barrier cannot reach: `/pay` makes a sync write first (the barrier and `@commit1` do not fire), then checks for the charge and inserts it with awaits in between (`@commit2` fires; under `BUG=dup-charge` it charges twice every time).

## Client rules

These rules stop the tool from blaming the service for its own misbehaviour:

- **Same request every time.** Every attempt is byte-identical and carries the same `Idempotency-Key`.
- **When it retries.** Only after a lost response, a 5xx, or a 409 ("key in progress"), at most 3 attempts in total, with seeded backoff.
- **Unkeyed calls are left alone.** An API call without an `Idempotency-Key` is never faulted, because a correct client has no right to retry it. Webhooks can always be redelivered.
- **Which response it keeps.** Under `concurrent`, the client keeps attempt A's response, so later id mapping does not depend on which attempt the scheduler finished first.

## Id mapping

Ids the service minted during recording differ from the ids it mints on replay. After each response, the replayer pairs the recorded and live responses by JSON path. Every string that differs, other than timestamps, is mapped, and later request bodies and path segments are rewritten. So a recorded `ch_abc` in a later webhook becomes this run's `ch_xyz`.

## Limits

- **Crashes only between commits.** Killing the service inside a transaction is uninteresting, because Postgres rolls it back. stutter only crashes right after a commit.
- **The barrier can't always fill.** It never fills for a handler whose first write goes outside the configured schemas, for a service that answers a duplicate before it touches Postgres (an in-memory or Redis idempotency store), or for a service with a connection pool of 1. All of these show up as not fired.
- **Holds are opt-in and stop at commit points.** A held commit can only park A right after one of its commits. A race between two reads with no commit between them stays out of reach, and so does a model call or other await that does not commit.
- **The barrier only holds the first write.** It controls the race at a handler's first write, not the gap between a multi-phase handler's later transactions. In minipay's `phases-no-recovery`, a concurrent duplicate hits that gap only some of the time, so stutter reports those detections but shrinks the crash finding instead, which is deterministic.
- **The test database runs with durability off.** `fsync`, `synchronous_commit` and `full_page_writes` are off. Those settings only matter after an OS crash; they do not change isolation or locking.
