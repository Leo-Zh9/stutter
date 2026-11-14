# Faults and client rules

Faults are injected in the replayer's own transport, the only place where a lost response is observable. Code: [internal/replay/replay.go](../internal/replay/replay.go).

| Fault | What happens | Real-world cause |
|---|---|---|
| `lost(#n)` | Request n reaches the service and completes. The sender reads the response, discards it, and retries. | Timeout or dropped connection after the server committed |
| `concurrent(#n)` | A duplicate of request n is sent while the first attempt is still inside its transaction. | Client timeout and retry while the first attempt is slow |
| `reorder(#a,#b)` | Two webhook deliveries about the same object arrive swapped. | Providers do not guarantee delivery order |

`reorder` applies to pairs of webhooks whose bodies reference the same id minted earlier in the trace.

## The lock barrier

Sending a duplicate at the same moment rarely lands inside the race window, because a handler's read and its write are microseconds apart. For `concurrent`, the replayer:

1. Opens a transaction holding `LOCK TABLE ... IN EXCLUSIVE MODE` on every table in `public`. That mode lets plain `SELECT`s through but blocks every write.
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

- **No crash points yet.** Killing the service inside one transaction is uninteresting, because Postgres rolls it back. Crashing it between two of its own transactions needs a view of commits from outside the service.
- **The barrier can't always fill.** It never fills for a handler whose first write goes outside the `public` schema, or for a service with a connection pool of 1. Both show up as not fired.
- **The test database runs with durability off.** `fsync`, `synchronous_commit` and `full_page_writes` are off. Those settings only matter after an OS crash; they do not change isolation or locking.
