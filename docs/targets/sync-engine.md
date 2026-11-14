# stripe/sync-engine, before and after PR #177

| | |
|---|---|
| Repo | [stripe/sync-engine](https://github.com/stripe/sync-engine), 1,088 stars, 118 forks (read 2026-11-14) |
| Fix | [PR #177](https://github.com/stripe/sync-engine/pull/177), "chore: add upsert timestamp protection", merged 2025-08-26 as `d486b41dc1b18626ef798890b5bbc7f11f476517` |
| Pre-fix | `96e6b398d45d0c57ff2193fa4123fa988b1c4e45` (2025-08-19), the merge's first parent |
| Code changed | none, at either commit |
| Stack | Fastify 5.4, node-postgres 8.16, `stripe@18.2.0` in the engine package (pins `2025-05-28.basil`), Node 24.21 |
| Configs | [stutter.json](../../targets/sync-engine/stutter.json) (pre-fix), [stutter-fix.json](../../targets/sync-engine/stutter-fix.json) (fix) |
| Results | [hunt-pre/REPORT.md](../../targets/sync-engine/hunt-pre/REPORT.md), [hunt-fix/REPORT.md](../../targets/sync-engine/hunt-fix/REPORT.md) |

## This is a case study, not a blind result

PR #177 was found by searching merged fix PRs, before the workload was frozen. Its description says that an older webhook arriving out of order could overwrite newer data. So the bug was known, and so was the commit that fixed it. The question here is narrower: does stutter, run on the same frozen workload used for the two starters, fail the commit before the fix and pass the fix?

The expectation, written before any run, was that the pre-fix commit fails on reorder faults and the fix is clean. That is what happened. The details that did not match are listed at the end.

## Setup

- **Fetch and build.** [scripts/fetch-targets.sh](../../scripts/fetch-targets.sh) clones the repo twice, into `.targets/sync-engine-pre` and `.targets/sync-engine-fix`, then runs `pnpm install --frozen-lockfile` and `pnpm build` in each. The build is `tsup` for the engine and `tsc` for the Fastify app. Neither calls Stripe.
- **Start.** `node packages/fastify-app/dist/src/server.js`. The server runs its own migrations at boot, then listens on `PORT`. Webhooks go to `/webhooks`, health is `/health`.
- **Offline.** Three switches are set: `REVALIDATE_ENTITY_VIA_STRIPE_API=false`, `BACKFILL_RELATED_ENTITIES=false` (its default is `true`) and `AUTO_EXPAND_LISTS=false`. At both commits, the subscription path is then: verify the signature locally (`constructEventAsync`), use the payload as is (`fetchOrUseWebhookData`), skip `backfillCustomers`, skip `expandEntity`, and write to Postgres. No other call is made.
- **Network check.** Before recording, the 16 events were delivered once to each commit, with a preload script logging every socket connect and DNS lookup the node processes made. The server's only connects were to `127.0.0.1:54329`, the project-local Postgres. Its only lookup was `0.0.0.0`, for its own listen address.
- **Database.** Each template is built by the engine's own `runMigrations`, called with `node -e`. That function logs migration errors and does not throw, so the templates were checked by hand. Pre-fix has 33 migrations and 6 foreign keys. The fix has 35 migrations and none, because its migration 0034 drops them all.
- **Fixtures.** Before the fix, foreign keys require the rows the events point at: `subscriptions.customer` to `customers` (migration 0004), `subscription_items.price` to `prices` (0013), and `prices.product` to `products` (0003). With backfill off, nothing fetches them. [fixtures.sql](../../targets/sync-engine/fixtures.sql) adds five customers, two products and two prices. The same file is used at both commits.
- **Events.** The frozen workload, [lifecycles.mjs](../../targets/stripe/lifecycles.mjs), unchanged: 16 `customer.subscription.*` events across five lifecycles, 60 s apart. Its sha256 was checked before recording and matched `bc121c1f…cdb6` ([record.log](../../targets/sync-engine/record.log)).
- **API version.** `--api basil`. sync-engine's own `STRIPE_API_VERSION` default is `2020-08-27`, but that only sets the version of outbound API calls, which are off here. The SDK the engine verifies webhooks with, `stripe@18.2.0`, pins `2025-05-28.basil`. PR #173, merged before the pre-fix commit, added `current_period_*` to subscription items, which is where basil puts them. The generator's basil shape is labelled `2025-04-30.basil`.
- **Signing.** `--sdk` is `.targets/sync-engine-pre/packages/sync-engine/node_modules/stripe`, the copy the engine imports. `--customers` is left at its default, which matches the fixture.
- **Recording.** One trace, recorded against the pre-fix commit, is replayed against both. All 16 deliveries got HTTP 200 in order. The network-check runs recorded byte-identical traces at both commits.
- **Faults.** Same as the starters: stutter re-signs each delivery at send time, retries any non-2xx up to 5 times, and reorders pairs of events for the same Stripe customer. Both hunts ran with `-lanes 4`.

## Results

| Fault | Schedules | Fired | Pre-fix (96e6b39): changed the final state | Fix (d486b41): changed the final state |
|---|---|---|---|---|
| lost (redelivery) | 16 | 16 | 0 | 0 |
| concurrent redelivery | 16 | 16 | 0 | 0 |
| reorder | 23 | 23 | 13 (7 more are no-op pairs) | 0 (11 are no-op pairs) |
| crash after commit | 32 | 32 | 0 | 0 |

Schedules and fired counts are the same at both commits. Pre-fix exits 1 with 13 findings in 87 s. The fix exits 0 in 16 s.

Duplicates and crashes are harmless at both commits, because every write is an upsert keyed on the Stripe id. **Ordering is not, before the fix.** The 13 failing reorders shrink to 8 distinct end-state changes, which fall into three effects. Every finding is 1-minimal, reproduced in 10 of 10 fresh replays, and reproduced by its curl script `repro.sh` in 10 of 10 trials on fresh databases.

1. **A canceled subscription comes back** (findings 1 to 6 and 10 to 13). If `customer.subscription.deleted` arrives before an earlier event for the same subscription, the older event overwrites `canceled` and clears `canceled_at` and `ended_at`. Customer 1's subscription ends `trialing` (finding 1), `active` (2, 4) or `past_due` (3). Customer 2's first subscription ends `active` (5, 10, 11), or `active` with `cancel_at_period_end: true` (6, 12, 13). Minimal case: two deliveries, `#4 deleted(sub_one_a)` then `#1 updated(sub_one_a, active)`. Findings 10 to 13 need three: the swap is with the new subscription's event `#8` or `#9`, and the old `deleted` (`#7`) stays in between, so `#5` or `#6` lands after it.
2. **An upgrade is undone** (finding 7). If `created(sub_three_a, basic)` arrives after `updated(sub_three_a, plus)`, the item's price goes back to `price_prod_basic`, and so does the `items` copy on the subscription row. Minimal case: `#11` then `#10`.
3. **A subscription goes back to trialing** (findings 8, 9). A late `created(trialing)` overwrites a later status: a trial that ended `unpaid` shows `trialing` (8), and a trial that converted to `active` shows `trialing` (9). Minimal cases: `#13` then `#12`, and `#15` then `#14`.

At the fix, the 9 failing pairs that swap two events of one subscription all pass. The other 4 (findings 10 to 13) are no-op pairs there: the fix's later `deleted` wins whether or not those two events arrive, so their passing says nothing. Their effect is the same as findings 5 and 6, whose pairs do pass at the fix.

## Root cause and fix

Before the fix, every webhook write goes through `upsertMany`, whose SQL is `insert ... on conflict (id) do update set` every column, with no condition ([`postgres.ts:96-119`](https://github.com/stripe/sync-engine/blob/96e6b398d45d0c57ff2193fa4123fa988b1c4e45/packages/sync-engine/src/database/postgres.ts#L96-L119)). The last event to arrive wins, whatever its age.

PR #177 adds a `last_synced_at` column to every table (migration 0033). Each webhook passes its `event.created` as that value ([`stripeSync.ts:69`](https://github.com/stripe/sync-engine/blob/d486b41dc1b18626ef798890b5bbc7f11f476517/packages/sync-engine/src/stripeSync.ts#L69)), and the upsert only updates when the stored value is null or older ([`postgres.ts:186`](https://github.com/stripe/sync-engine/blob/d486b41dc1b18626ef798890b5bbc7f11f476517/packages/sync-engine/src/database/postgres.ts#L186)). Subscription items get the same guard.

Two limits of this check apply to the fix:

- The workload's events are 60 s apart. Stripe notes that distinct events can share a `created` second. With the fix's strict `<`, an equal timestamp keeps the first event to arrive. This workload never produces a tie, so it cannot show that case.
- stutter's oracle ignores timestamp columns, so `last_synced_at` itself is not compared. Every other column of every row is.

## What did not match the expectation

- **Nothing in the headline.** Pre-fix failed only on reorders, and the fix was clean on all 87 single faults.
- **No-op pairs grow at the fix.** 7 pre-fix, 11 at the fix. The extra 4 are findings 10 to 13, explained above. So the fix's clean reorder result rests on 12 pairs that are not no-ops, not 23.
- **Effects beyond cancellation.** The PR describes older events overwriting newer data in general. The undone upgrade and the return to `trialing` are instances of that, but were not called out in advance.
- **The fix changes more than ordering.** Migration 0034 drops all foreign keys. That is why the fixture is only needed before the fix.
- **API version label.** The generator's basil events say `2025-04-30.basil`. The engine's SDK pins `2025-05-28.basil`. Both are basil, with `current_period_*` on items, and the engine does not check the label.
