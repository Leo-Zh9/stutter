# stutter

stutter finds retry bugs in HTTP services that write to PostgreSQL. It records a service's traffic, replays it against a fresh copy of the database while injecting the faults real networks produce, and flags any fault that changes the final data. Each failure is then shrunk to the smallest set of requests and faults that still breaks.

The check rests on one rule: **a request that was retried correctly must leave the database exactly as if it had succeeded the first time.** stutter checks that rule automatically by comparing the faulted run's final state with a fault-free replay of the same trace. Nobody writes per-endpoint assertions.

```
$ make explore BUG=key-after-commit
66 requests, 337 schedules
run 72 of 337 changed the final state; shrinking
shrunk 67 -> 3 elements (140 candidates)

  #2 POST /customers key=key_cus_2
  #5 POST /charges key=key_ch_0
  fault concurrent(#5)

  + charges x1 {"amount_cents":581,"currency":"usd","customer_id":"<ref:customers>","idempotency_key":"key_ch_0",...}
  + customers x1 {"balance_cents":1162,"email":"buyer2@example.test"} (+4 referencing rows)
  - customers x1 {"balance_cents":581,"email":"buyer2@example.test"} (+3 referencing rows)
```

Read that as: create a customer, charge them once, and have the client send a duplicate while the first attempt is still inside its transaction. The customer is charged twice.

## Quick start

Requirements: Go 1.25+ (tested with 1.27.4) and [uv](https://docs.astral.sh/uv/). No Docker is needed. PostgreSQL 16 comes from the `pgserver` package and runs as a project-local cluster on a unix socket.

```bash
make demo      # start Postgres, build, run the full measurement suite (about 90 s)
make test      # unit and integration tests
make explore BUG=webhook-stale-overwrite   # hunt one seeded bug; exits 1 when it finds one
make stop      # stop the local Postgres
```

`make demo` writes `results/results.json` and `results/RESULTS.md`. Every number below comes from those files.

## How it works

```
                record                                       replay (per run, per lane)
 workload ──► stutter proxy ──► service ──► Postgres     trace ──► replayer ──► service ──► fresh DB clone
                   │                                                │  faults         │
                   ▼                                                │  lock barrier ◄─┘
             trace.jsonl ───────────────────────────────────────────┘                  │
                                                                     snapshot ◄────────┘
                                                                        │
                                                    canonicalize, diff against fault-free baseline
                                                                        │
                                                              ddmin shrink on failure
```

1. **Record.** `stutter record` puts a reverse proxy in front of the service and writes each request and response to a JSONL trace. It keeps the body and the `Idempotency-Key` and `Content-Type` headers, and never records credentials.
2. **Reset.** Every replay runs on its own database, cloned from a template with `CREATE DATABASE ... TEMPLATE ... STRATEGY FILE_COPY`, and on its own copy of the service process. Eight lanes run in parallel.
3. **Replay.** The replayer acts as a well-behaved client and webhook provider (rules below). Ids the service minted during recording are mapped to the ids it mints live. A recorded `ch_abc` in a later webhook body becomes this run's `ch_xyz`.
4. **Compare.** After the last request the replayer stops the service, snapshots every table, and diffs the result against the fault-free run of the same trace.
5. **Shrink.** A failing schedule goes through delta debugging over its requests and faults together.

### Faults

| Fault | What happens | Real-world cause |
|---|---|---|
| `lost(#n)` | Request n reaches the service and completes. The sender reads the response, throws it away, and retries. | Timeout or dropped connection after the server committed |
| `concurrent(#n)` | A duplicate of request n is sent while the first attempt is still inside its transaction. | Client timeout and retry while the first attempt is slow |
| `reorder(#a,#b)` | Two webhook deliveries about the same object arrive swapped. | Providers do not guarantee delivery order |

Faults are injected in the replayer's own transport, which is the only place where a lost response is observable. Process crashes are out of scope (see Limits).

**The lock barrier.** Sending a duplicate "at the same time" rarely lands inside the race window, because a handler's read and write are microseconds apart. So for `concurrent`, the replayer:

1. Opens a transaction that holds `LOCK TABLE ... IN EXCLUSIVE MODE` on every table. That mode lets plain `SELECT`s through but blocks any write.
2. Sends attempt A, waits until its backend shows up in `pg_stat_activity` waiting on a lock, then sends attempt B.
3. Waits until both backends are waiting, then commits to release the locks.

Both attempts have now done their reads and neither has written, which is exactly the window where check-then-act bugs live. If the barrier does not fill within 500 ms, the fault is recorded as **not fired**. A not-fired run is never counted as a pass, and every rate below reports fired and not-fired runs separately.

### Client rules (so the tool never blames itself)

- Every attempt of a request is byte-identical and carries the same `Idempotency-Key`.
- Retries happen only after a lost response, a 5xx, or a 409 ("key in progress"), at most 3 attempts in total, with seeded backoff.
- **An API call without an `Idempotency-Key` is never faulted.** A correct client has no right to retry a non-idempotent call, so a duplicate there would be the tool's bug, not the service's. Webhooks can always be redelivered.
- Under `concurrent`, the client keeps attempt A's response, so later id mapping does not depend on which attempt the scheduler finished first.

### The oracle

Two runs of the same correct program never produce byte-identical databases: uuids differ, sequences skip values a rolled-back duplicate consumed, timestamps move. The oracle factors that out using the catalog alone:

- **Generated values.** A column is generated if its default calls a non-immutable function (`gen_random_uuid`, `nextval`, `now`, ...) or it is an identity column. Its values are never compared directly.
- **Edges.** A foreign key to a generated column is an edge. So is any string in a `text` or `jsonb` column that exactly equals a generated value, which is how the cached response in an idempotency table stays tied to the charge it embeds.
- **Signatures.** Each row gets a hash built bottom-up from its own non-generated data plus the sorted signatures of the rows that point at it. This is AHU tree hashing, the technique used for tree isomorphism. Timestamp columns are dropped, and RFC 3339 strings inside JSON become `<ts>`.
- **Comparison.** Two states are equal when every table holds the same multiset of row signatures.

So two customers who are identical except for their ids, each with a different charge, compare equal in any insert order. Moving a charge from one to the other, adding a row, or changing one amount does not compare equal.

**Guard against false alarms.** Before exploring, stutter replays the trace fault-free twice. If those runs disagree, the target has nondeterminism the catalog cannot see, and stutter names the column instead of producing false alarms. On minipay this caught `idempotency_keys.request_hash`, a hash of a request body that contains a generated customer id. One config line marks it (`-ignore idempotency_keys.request_hash`).

### Shrinking

Shrinking uses Zeller's ddmin over a combined list of trace requests and faults. Dropping a request drops the faults that name it, and every candidate is compared against the fault-free baseline of its own reduced trace (baselines are cached), so removing the request that created an id stays sound. Candidates are evaluated in parallel, and ddmin always takes the lowest-index failing candidate, so the result matches a sequential run.

A candidate with a `concurrent` fault is run 3 times and counts as failing only if all 3 runs fired and failed. A candidate whose fault did not fire counts as passing, which can only keep extra elements, never drop a needed one. The result is **1-minimal over fired candidates**: removing any single remaining request or fault makes the failure go away. The bench re-checks this by brute force for every bug.

## Results

The target is **minipay** (`targets/minipay`), a small payments API written for this project: customers, charges with idempotency keys, and provider webhooks. **It is simulated.** Each variant switches on exactly one bug from a documented real-world class. The trace is 66 requests from a seeded synthetic workload (5 customers, 40 charges, one client that reuses a key with a different body, and 20 webhooks), not production traffic.

Every variant faces the same 337 schedules: all 137 single faults the trace admits (66 lost, 66 concurrent, 5 reorder), then 200 random schedules of 2 to 4 faults.

| Variant | The bug | Found by | First failing run | Shrunk to | 1-minimal |
|---|---|---|---|---|---|
| correct | none | (0 false alarms in 337 fired runs) | - | - | - |
| key-ignored | key accepted, never checked | lost | 6 of 337 | 67 -> 3 elements | yes |
| key-after-commit | key recorded in a second transaction after the charge commits (the race that inserting the key first avoids, as in [Brandur Leach's Stripe-style design](https://brandur.org/idempotency-keys)) | concurrent | 72 of 337 | 67 -> 3 | yes |
| webhook-no-dedupe | handler keeps no event log; a redelivered refund is applied twice | lost | 50 of 337 | 67 -> 4 | yes |
| webhook-stale-overwrite | status webhook applied without comparing event time | reorder | 133 of 337 | 67 -> 5 | yes |
| webhook-check-then-insert | event dedupe is `SELECT` then `INSERT` with no unique index | concurrent | 113 of 337 | 67 -> 4 | yes |
| key-unscoped (control) | key not bound to the request body | **not detected, by design** | - | - | - |

- **5 of 5 in-model bugs detected. 0 false alarms in 337 fired runs on the correct variant** (0 runs did not fire).
- **Race bugs need the right fault.** `key-after-commit` and `webhook-check-then-insert` survive every `lost` fault, because a sequential retry finds the key or event row. Only a true concurrent duplicate breaks them, and a test pins that.
- **The control marks the method's edge.** `key-unscoped` returns the first charge for a reused key with a different body. That is wrong with or without faults, so the faulted and fault-free runs agree and a differential oracle cannot see it. Catching it needs a spec-level check (EvoMaster-style HTTP semantics oracles are the right tool).

**Does the barrier matter?** Each of the two race bugs gets its triggering `concurrent` fault replayed 100 times with the barrier and 100 times with both attempts fired back to back from two goroutines, on a fresh clone each time:

| Bug | Barrier on | Barrier off (run 1) | Barrier off (run 2) |
|---|---|---|---|
| key-after-commit | 100 / 100 fired, 100 detected | 7 / 100 | 1 / 100 |
| webhook-check-then-insert | 100 / 100 fired, 100 detected | 2 / 100 | 1 / 100 |

Without the barrier, hitting the race depends on loopback timing and lane contention, so that arm varies between runs. Both full runs are kept in `results/runs/`.

**Speed.** 35.7 replays per second across 8 lanes on an Apple Silicon Mac (10 cores). Each replay clones the database, starts the service, replays 66 requests, snapshots and drops the database. The median single replay takes 55 ms. Shrinks took 2 to 18 s. The full suite takes about 90 s. Twenty fault-free replays produced one canonical state digest.

The test cluster runs with `fsync`, `synchronous_commit` and `full_page_writes` off. Those settings only matter after an OS crash. They do not change isolation or locking, which is all the faults depend on.

## Guarantees and the tests that hold them

| Guarantee | Test |
|---|---|
| Equivalent states compare equal under any insert order, uuids, sequence gaps and timestamps | `TestEquivalentStatesAreEqual`, `TestBurnedIdsAreNotADiff` |
| Isomorphic parents do not swap labels; moving a child between them is a diff | `TestIsomorphicParents` |
| One changed value, one extra row, or a re-pointed reference is a diff | `TestRealChangesAreDetected`, `TestBurnedIdsAreNotADiff` |
| The generated-value rule matches the real catalog (uuid, serial, identity, now; not constants) | `TestCatalogGeneratedRule` |
| Each fault fires: lost sends twice, concurrent puts two backends behind the barrier, reorder swaps delivery | `TestLostFires`, `TestConcurrentBarrierFires`, `TestReorderSwapsDelivery` |
| A barrier that never fills is not fired, never a pass | `TestBarrierTimeoutIsNotFired` |
| Unkeyed API calls are never faulted | `TestUnkeyedRequestsAreNeverFaulted` |
| Fault-free replay is deterministic | `TestFaultFreeReplayIsDeterministic` |
| ddmin output fails and is 1-minimal (300 random failure predicates, checked by brute force) | `TestDDMinOneMinimalProperty` |
| No false alarms on the correct variant; every in-model bug is found by its fault type and shrinks 1-minimal; the control stays undetected | `internal/explore/explore_test.go` |

Each of these tests was checked by mutation: the mechanism it guards was broken on purpose, and the test failed. That covered dropping child signatures, ignoring string edges, treating identity columns as data, counting a barrier timeout as fired, skipping the discarded first attempt, treating unkeyed calls as retryable, not swapping on reorder, and two ddmin defects.

## Limits

- **Black box, crashes excluded.** stutter sees HTTP and the database. It cannot crash the service between two of its own transactions. A `SIGKILL` inside one transaction is uninteresting, because Postgres rolls it back.
- **The oracle is sound, not complete.** False alarms stay at zero because equivalent states always hash equal. Two different states can still collide when a row references two parents and a bug only re-pairs them.
- **Edges are exact strings.** Integer ids stored as JSON numbers inside `text` or `jsonb` are not matched, and neither are partial matches inside longer strings. Foreign keys are single-column.
- **Timestamps are ignored,** so a bug whose only effect is a wrong timestamp is missed.
- **Barrier coverage.** A handler whose first write goes outside the `public` schema, or a service with a connection pool of 1, cannot fill the barrier. That shows up as not fired, not as a pass.
- **One simulated target.** Open-source services that claim idempotency (Medusa, Lago, Saleor) need Docker or Redis or call a live payment provider. Self-contained ones make no idempotency claim to test.

## Prior art

- **Fault injectors** (Toxiproxy, xk6-disruptor, Filibuster) produce failures, but deciding whether data was corrupted is left to hand-written assertions.
- **Jepsen** checks the consistency models of databases over operation histories, not an API's retry behaviour.
- **Wasabi** (SOSP 2024) finds retry bugs inside Java systems, with crash, hang and missing-backoff oracles.
- **REST fuzzers** check other properties:
  - RESTler: use-after-free, namespace, resource hierarchy, invalid ids.
  - Schemathesis: 5xx, use-after-free, resource availability.
  - Sahin and Arcuri ([arXiv 2608.16239](https://arxiv.org/abs/2608.16239)): HTTP-semantics oracles on responses in EvoMaster, including non-idempotent `PUT`.

What stutter adds is the combination: a database-state differential oracle, faults placed at the commit boundary (lost-after-commit, plus the lock barrier for true concurrency), webhook reordering, and shrinking to a minimal reproduction. The differential idea itself is ordinary metamorphic testing.

## Layout

```
cmd/stutter         CLI: record, explore, bench
internal/trace      trace format and recording proxy
internal/replay     client rules, faults, lock barrier, id mapping
internal/oracle     catalog snapshot, AHU canonicalization, diff
internal/explore    schedules, parallel lanes, baselines, shrink driver
internal/shrink     ddmin
internal/bench      measurement suite -> results/
internal/pgenv      template databases and clones
targets/minipay     simulated payments API with seeded bugs, plus its workload generator
traces/             the recorded trace used by the bench
```
