# Faults and client rules

Faults are injected in the replayer's own transport, the only place where a lost response is observable. Code: [internal/replay/replay.go](../internal/replay/replay.go).

| Fault | What happens | Real-world cause |
|---|---|---|
| `lost(#n)` | Request n reaches the service and completes. The sender reads the response, discards it, and retries. | Timeout or dropped connection after the server committed |
| `concurrent(#n)` | A duplicate of request n is sent while the first attempt is still inside its transaction. | Client timeout and retry while the first attempt is slow |
| `reorder(#a,#b)` | Two webhook deliveries about the same object arrive swapped. | Providers do not guarantee delivery order |
| `crash(#n@commitK)` | The service is killed right after the K-th commit of request n, before it learns the commit succeeded. It restarts on the same database and the client retries. | Process crash or deploy between two of a handler's own transactions |

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
- **The barrier only holds the first write.** It controls the race at a handler's first write, not the gap between a multi-phase handler's later transactions. In minipay's `phases-no-recovery`, a concurrent duplicate hits that gap only some of the time, so stutter reports those detections but shrinks the crash finding instead, which is deterministic.
- **The test database runs with durability off.** `fsync`, `synchronous_commit` and `full_page_writes` are off. Those settings only matter after an OS crash; they do not change isolation or locking.
