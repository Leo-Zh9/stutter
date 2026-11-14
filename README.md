# stutter

stutter finds retry bugs in HTTP services backed by PostgreSQL: double charges, webhooks applied twice or out of order, and work lost when a process dies between two commits.

It records a service's traffic, replays it on fresh copies of the database while injecting the faults real networks and deploys produce, and flags any fault that changes the final data. Every failure is shrunk to a minimal reproduction and written out as a curl script.

The rule it checks: **a request that was retried, redelivered or reordered must leave the database exactly as if nothing had gone wrong.** You don't write assertions. stutter compares each faulted run with a fault-free replay of the same trace.

## What it found in real code

| Target | Result |
|---|---|
| [nextjs/saas-starter](docs/targets/saas-starter.md) (16k stars), unmodified | Out-of-order Stripe webhooks reactivate canceled teams, cancel teams that are paying, or point a team at its old plan. 7 failing reorders, 3 effects, every one reproduced 10/10 by curl. |
| [boxyhq/saas-starter-kit](docs/targets/boxyhq.md) (5k stars), unmodified | Out-of-order webhooks recreate canceled subscriptions as active, and turn a trial that ended unpaid into an active subscription. 11 failing reorders, 3 effects, every one reproduced 10/10 by curl. |
| [Formance Ledger v2.4.18](docs/targets/formance.md), unmodified release binary | Clean: 0 of 129 single faults (lost, concurrent, crash) and 0 of 286 more schedules changed the final state. |
| Formance with one defence removed (4 builds) | Every removed defence is caught by the fault its design predicts. Dropping only the in-code lookup is masked by the unique index, as designed. |

Both starters' root cause is a webhook handler that applies every subscription event unconditionally, while Stripe documents that delivery order is not guaranteed. The hypotheses came from reading the code. stutter confirmed them, enumerated every pair, minimized them and produced the reproductions. Issue drafts are in each target's doc; none has been filed.

```
$ ./bin/stutter hunt -config targets/saas-starter/stutter.json
...
#7 webhook customer.subscription.deleted evt_0007 customer=cus_team2 id=sub_two_a status=canceled
#9 webhook customer.subscription.updated evt_0009 customer=cus_team2 id=sub_two_b status=active
fault reorder(#7,#9)

+ teams x1 {"name":"Team 2",...,"stripe_subscription_id":null,"subscription_status":"canceled"}
- teams x1 {"name":"Team 2",...,"stripe_subscription_id":"sub_two_b","subscription_status":"active"}
```

In this finding, Team 2 canceled one subscription and started another. When the old subscription's `deleted` event arrives late, the paying team is marked canceled.

## Quick start

Requires Go 1.25+ and [uv](https://docs.astral.sh/uv/). No Docker: PostgreSQL 16 runs as a project-local cluster.

```bash
make demo        # Postgres, build, the minipay benchmark (about 3 min)
make test        # unit and integration tests
./scripts/fetch-targets.sh && make hunt   # the real targets above (about 50 min; needs node, pnpm)
```

To point stutter at your own service, describe it in a `stutter.json` ([docs/usage.md](docs/usage.md)) and run `stutter record`, then `stutter hunt`.

## How it works

1. **Record.** A reverse proxy writes every request and response to a JSONL trace.
2. **Reset.** Each replay runs on its own clone of a template database, with its own copy of the service.
3. **Replay with faults.** The replayer acts as a well-behaved client or webhook provider: the same body and the same `Idempotency-Key` on every attempt, unkeyed API calls never retried, and Stripe-style redelivery of webhooks.
4. **Compare.** The final database is snapshotted and compared with the fault-free run. Ids, sequence values and timestamps the database generated are factored out using the catalog. For services that mint integer ids in their own SQL, a second fault-free run with shifted sequences is used to infer them.
5. **Shrink.** Delta debugging reduces a failure until removing any single request or fault makes it pass.

| Fault | What happens |
|---|---|
| `lost(#n)` | The request completes on the server; the client never sees the response and retries |
| `concurrent(#n)` | A duplicate arrives while the first attempt is inside its transaction, forced into the race window by a lock barrier |
| `reorder(#a,#b)` | Two webhooks about the same object arrive swapped |
| `crash(#n@commitK)` | A Postgres wire-protocol proxy holds back the K-th commit acknowledgement of request n; stutter kills the service, restarts it, and the client retries |

Details: [faults and crash points](docs/faults.md), [the oracle and id inference](docs/oracle.md), [shrinking](docs/shrinking.md), [guarantees and tests](docs/guarantees.md), [prior art](docs/prior-art.md).

## Benchmark: minipay

minipay ([targets/minipay](targets/minipay)) is a **simulated** payments API with one seeded bug per variant, each drawn from a documented real-world mistake. The trace is 74 requests from a synthetic workload. Every variant faces the same 242 single faults, including 89 crash points, and the correct variant also faces 200 random multi-fault schedules. Numbers are from [results/results.json](results/results.json).

| Variant | Bug | lost | concurrent | reorder | crash |
|---|---|---|---|---|---|
| correct | none | 0 / 74 | 0 / 74 | 0 / 5 | 0 / 89 |
| key-ignored | idempotency key never checked | 41 / 74 | 41 / 74 | 0 / 5 | 40 / 89 |
| key-after-commit | key saved in a second transaction | 0 / 74 | 40 / 73 | 0 / 5 | 40 / 89 |
| webhook-no-dedupe | no event log | 5 / 74 | 5 / 74 | 0 / 5 | 5 / 89 |
| webhook-stale-overwrite | status event ignores event time | 0 / 74 | 0 / 74 | 5 / 5 | 0 / 89 |
| webhook-check-then-insert | dedupe is SELECT then INSERT | 0 / 74 | 20 / 74 | 0 / 5 | 0 / 89 |
| outbox-after-commit | receipt queued after the charge commits | 0 / 74 | 0 / 74 | 0 / 5 | 40 / 89 |
| phases-no-recovery | multi-phase payout retried from the top | 0 / 74 | 1 / 74 | 0 / 5 | 8 / 89 |
| key-unscoped (control) | key not bound to the body | 0 / 74 | 0 / 74 | 0 / 5 | 0 / 89 |

Each cell shows schedules that changed the final state out of schedules that fired.

- **Detection.** All 7 in-model bugs are found, with 0 false alarms in 442 fired runs on the correct variant. Each shrinks to 3 to 5 elements, verified 1-minimal by brute force.
- **Crash points find what retries can't.** `outbox-after-commit` survives every lost and concurrent fault, because the retry replays a stored response. Only a crash between its two commits shows the lost receipt.
- **The lock barrier matters.** With the barrier, the two race bugs triggered 100 of 100 times. Without it, both attempts sent back to back from two goroutines triggered them 1 to 21 times in 100 across runs ([results/runs](results/runs)).
- **The control marks the method's limit.** A key reused with a different body is wrong with or without faults, so a differential oracle cannot see it.
- **Speed.** 28.3 replays per second on 8 lanes, median replay 76 ms, on a 10-core Apple Silicon Mac. The full benchmark takes about 150 s.

## Limits

- **PostgreSQL and HTTP only.** Services that answer a duplicate before touching Postgres (in-memory or Redis idempotency stores) cannot fill the barrier.
- **Crashes are placed only after commits.** Writes inside functions called from `SELECT`, `COMMIT AND CHAIN`, and commits inside procedures are not crash points.
- **The oracle is sound, not complete.** Equivalent states always compare equal, which is why false alarms stay at zero. A bug that only re-pairs rows referencing two parents can collide, a JSON path that mixes id and non-id values is dropped, and timestamps are ignored.
- **Concurrency is held at the first write.** A race between a multi-phase handler's later transactions is hit only some of the time, and stutter reports it as such.

## Layout

```
cmd/stutter        CLI: record, explore, hunt, bench
internal/config    stutter.json, template preparation, sequence bands
internal/trace     trace format, recording proxy
internal/replay    client rules, faults, lock barrier, crash restarts, id mapping
internal/pgproxy   wire-protocol proxy that sees and holds back commit acks
internal/oracle    catalog snapshot, id inference, canonical form, diff
internal/explore   schedules, lanes, baselines, crash enumeration, shrink driver
internal/hunt      every single fault, findings, curl repro scripts
internal/shrink    ddmin
internal/bench     minipay measurement suite, writes results/
targets/           minipay, the two SaaS starters, Formance and its mutants
```
