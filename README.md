# stutter

stutter finds retry bugs in HTTP services backed by PostgreSQL: double charges, webhooks applied twice, stale events overwriting newer state.

It records a service's traffic, replays it on a fresh database while injecting the faults real networks produce, and flags any fault that changes the final data. Every failure is shrunk to a minimal reproduction.

The rule it checks: **a request that was retried correctly must leave the database exactly as if it had succeeded the first time.** You don't write any assertions. stutter compares each faulted run with a fault-free replay of the same trace.

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

In this run, a client sent a duplicate charge while the first attempt was still inside its transaction, and the customer was charged twice.

## Quick start

Requires Go 1.25+ and [uv](https://docs.astral.sh/uv/). No Docker: PostgreSQL 16 runs as a project-local cluster from the `pgserver` package.

```bash
make demo                                  # start Postgres, build, run the full suite (about 90 s)
make test                                  # unit and integration tests
make explore BUG=webhook-stale-overwrite   # hunt one seeded bug; exits 1 when it finds one
make stop                                  # stop the local Postgres
```

To point stutter at your own service, describe it in a `stutter.json`: see [docs/usage.md](docs/usage.md).

## How it works

1. **Record.** A reverse proxy writes every request and response to a JSONL trace.
2. **Reset.** Each replay runs on its own clone of a template database and its own copy of the service, 8 in parallel.
3. **Replay with faults.** The replayer acts as a well-behaved client: same body and same `Idempotency-Key` on every attempt, and no retries of unkeyed calls.
4. **Compare.** The final database is snapshotted and compared with the fault-free run, after factoring out generated ids and timestamps.
5. **Shrink.** Delta debugging reduces a failing schedule until removing any single request or fault makes it pass.

| Fault | What happens |
|---|---|
| `lost(#n)` | The request completes on the server; the client never sees the response and retries |
| `concurrent(#n)` | A duplicate arrives while the first attempt is still inside its transaction |
| `reorder(#a,#b)` | Two webhooks about the same object arrive swapped |

Concurrent duplicates rarely hit the race window by chance. stutter forces them into it with a lock barrier: it holds `EXCLUSIVE` locks on every table, so both attempts finish their reads and then wait on their first write. It releases the locks only once Postgres shows both backends waiting.

## Results

Target: minipay ([targets/minipay](targets/minipay)), a **simulated** payments API with one seeded bug per variant, each drawn from a documented real-world mistake. The trace is 66 requests from a seeded synthetic workload. Every variant gets the same 337 schedules: all 137 single faults, then 200 random combinations. Numbers are from [results/results.json](results/results.json).

| Variant | Bug | Found by | Shrunk to |
|---|---|---|---|
| correct | none | 0 false alarms in 337 runs | - |
| key-ignored | idempotency key never checked | lost | 3 elements |
| key-after-commit | key saved in a second transaction after the charge | concurrent | 3 |
| webhook-no-dedupe | no event log; a redelivered refund applies twice | lost | 4 |
| webhook-stale-overwrite | status webhook ignores event time | reorder | 5 |
| webhook-check-then-insert | dedupe is SELECT then INSERT, no unique index | concurrent | 4 |
| key-unscoped (control) | key not bound to the request body | not detected, by design | - |

- **5 of 5 in-model bugs found. 0 false alarms.** Every shrunk schedule was rechecked by brute force and is 1-minimal.
- **The barrier matters.** The two race bugs triggered 100 of 100 times with the barrier, and 1 to 7 of 100 without it across three runs ([results/runs](results/runs)).
- **The control marks the method's limit.** A key reused with a different body is wrong with or without faults, so comparing against a fault-free run cannot see it.
- **Speed.** 35.7 replays per second on 8 lanes; median replay 55 ms; full suite about 90 s on a 10-core Apple Silicon Mac.

## Limits

- One simulated target so far.
- Process crashes are not injected yet. A crash inside one transaction is uninteresting, because Postgres rolls it back.
- Integer ids inside JSON and multi-column foreign keys are not matched as references.
- A bug whose only effect is a wrong timestamp is missed.

Details: [docs/usage.md](docs/usage.md), [docs/oracle.md](docs/oracle.md), [docs/faults.md](docs/faults.md), [docs/shrinking.md](docs/shrinking.md), [docs/guarantees.md](docs/guarantees.md), [docs/prior-art.md](docs/prior-art.md).

## Prior art

- **Fault injectors** (Toxiproxy, Filibuster) cause failures but leave judging the damage to hand-written assertions.
- **API fuzzers** check other properties:
  - RESTler and Schemathesis: crashes and use-after-delete.
  - EvoMaster: HTTP semantics on responses.
- **Jepsen** checks databases, not services.

stutter combines a database-state oracle, faults placed at the commit boundary and shrinking. See [docs/prior-art.md](docs/prior-art.md).

## Layout

```
cmd/stutter        CLI: record, explore, bench
internal/config    stutter.json loading, template preparation
internal/trace     trace format, recording proxy
internal/replay    client rules, faults, lock barrier, id mapping
internal/oracle    catalog snapshot, canonical form, diff
internal/explore   schedules, parallel lanes, baselines, shrink driver
internal/shrink    ddmin
internal/bench     measurement suite, writes results/
internal/pgenv     template databases and clones
targets/minipay    simulated payments API and its workload generator
```
