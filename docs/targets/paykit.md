# getpaykit/paykit (blind target)

| | |
|---|---|
| Repo | [getpaykit/paykit](https://github.com/getpaykit/paykit), default branch `main` |
| Pinned | `c7572ff7e0116428a43a2f98ad25dc2c26c8927e`, from the blind screen |
| Code changed | none |
| App run | `apps/demo` (Next.js 16.3.5), which mounts the `paykitjs` handler at `/paykit/[[...slug]]`. Webhooks go to `/paykit/webhook`. |
| Stack | `paykitjs` 0.2.0, drizzle-orm 0.45, node-postgres, `stripe@19.3.1` under `packages/paykit`, Node 24.21 |
| Config | [stutter.json](../../targets/paykit/stutter.json), [fixtures.sql](../../targets/paykit/fixtures.sql) |
| Results | [hunt/REPORT.md](../../targets/paykit/hunt/REPORT.md), [hunt/results.json](../../targets/paykit/hunt/results.json), [record.log](../../targets/paykit/record.log), [hunt/workload-check.txt](../../targets/paykit/hunt/workload-check.txt) |

## Blind protocol

- **Pinned SHA.** `c7572ff7e0116428a43a2f98ad25dc2c26c8927e`. No other commit was checked out.
- **Push before any source reading.** Config, fixtures, trace and hunt results were pushed as `937067a` ("paykit blind hunt: config, fixtures, trace and results"). The handler and the code it calls were read only after that push.
- **What was read before the push.**
  - Docs: `README.md`, `AGENTS.md` and `CLAUDE.md` (which only includes `AGENTS.md`), `.env.example`, and `apps/web/content/docs/` pages `webhook-events.mdx`, `installation.mdx`, `quickstart.mdx`, `database.mdx`.
  - Schema: the four migration files `packages/paykit/src/database/migrations/0000` to `0003`. `schema.ts` was not opened.
  - Build config: `package.json` of `packages/paykit` and `apps/demo`, the root `packageManager` field, `pnpm-workspace.yaml`, `apps/demo/drizzle.config.ts`, `apps/demo/next.config.js`, and `scripts/load-root-env.js` (see deviations).
  - A file-name listing of the repo and of `e2e/`, and the version field of the installed `stripe` package.
  - Runtime output only: HTTP statuses, and the rows the app wrote to scratch databases during smoke runs, including `paykit_webhook_event.status`, `error` and `payload`.
- **Generator flags.** `--api basil` (`stripe` is pinned `^19.1.0` in `packages/paykit/package.json`). `--sdk /Users/leozhang/stutter/.targets/paykit/packages/paykit/node_modules/stripe` (the library's own copy, 19.3.1). `--customers` left at its default, `cus_team1` to `cus_team5`, which the fixtures hold.
- **Workload hash.** `targets/stripe/lifecycles.mjs: OK` before recording ([record.log](../../targets/paykit/record.log)) and before the hunt ([workload-check.txt](../../targets/paykit/hunt/workload-check.txt)).

## Setup

- **Fetch and build.** [scripts/fetch-targets.sh](../../scripts/fetch-targets.sh) clones the repo into `.targets/paykit`, runs `pnpm@12.4.2 install --frozen-lockfile`, then a compile-only `next build` of `apps/demo` with env validation skipped.
- **Start.** `next start` in `apps/demo`, with `PAYKIT_DATABASE_URL` and `AUTH_DATABASE_URL` set to the run's database and `STRIPE_WEBHOOK_SECRET=whsec_stutter_local`. Health is `GET /paykit/webhook`, which answers 404.
- **Database.** The template is built by `drizzle-kit migrate` with the demo's own `drizzle.config.ts`, which applies the library's four migrations.
- **Offline.** Before recording, the 16 events were delivered once directly to the app with a preload script that logged every socket connect and DNS lookup of every node process. The only connect was to `127.0.0.1:54329`, the project-local Postgres. The only lookup was `127.0.0.1`. Nothing called Stripe.
- **Fixtures.** Five `paykit_customer` rows with `stripe_customer_id` set, two `paykit_product` rows with `stripe_price_id` `price_prod_basic` and `price_prod_plus`, and one `paykit_subscription` row per Stripe subscription in the workload, keyed by `stripe_subscription_id`, status `incomplete`. A first smoke run without the subscription rows got 200 for all 16 events, logged them as `processed`, and wrote no subscription state. With the rows, the events update them.
- **Recording.** All 16 deliveries got HTTP 200 in order.
- **Oracle.** The first hunt run stopped at the stability check: `paykit_webhook_event.id` and `paykit_webhook_event.trace_id` are random ids minted in application code. Those two columns were added to `ignore`. Every other column of every row is compared, except timestamps.
- **Faults.** stutter re-signs each delivery at send time, retries any non-2xx up to 5 times, and reorders pairs of events for the same Stripe customer. `-lanes 2`, crash points included. The hunt took 547 s.

## Results

| Fault | Schedules | Fired | Changed the final state |
|---|---|---|---|
| lost (redelivery) | 16 | 16 | 0 |
| concurrent redelivery | 16 | 16 | 0 |
| reorder | 23 | 23 | 13 (0 no-op pairs) |
| crash after commit | 32 | 32 | 16 |

29 findings, which shrink to 26 distinct end-state changes. Every finding is 1-minimal and reproduced in 10 of 10 fresh replays. The 13 reorder findings are also reproduced by their curl script `repro.sh` in 10 of 10 trials on fresh databases. The 16 crash findings cannot be scripted with curl.

Duplicates are harmless: a redelivery hits the unique `stripe_event_id` and is skipped. Ordering and crashes are not.

## Effects

**Reorder, one subscription (9 findings, 8 effects).**

1. **A canceled subscription comes back.** An older event for the same subscription arriving after `customer.subscription.deleted` overwrites `canceled`. Customer 1's subscription ends `trialing` (finding 17), `active` (18, 20) or `past_due` (19). Customer 2's first subscription ends `active` (22), or `active` with `cancel_at_period_end` (25). Minimal case: `#4 deleted(sub_one_a)` then `#0 created(sub_one_a)`.
2. **A pending cancellation is lost** (finding 21). `#6 updated(sub_two_a, cancel_at_period_end: true)` then `#5 created(sub_two_a)` leaves `cancel_at_period_end` and `canceled` false.
3. **A subscription goes back to trialing** (28, 29). A late `created(trialing)` overwrites `unpaid` (28) or `active` (29).

**Reorder, two subscriptions of one customer (4 findings, 2 effects).**

4. **The customer's new subscription is ended by an event about the old one** (23, 24, 26, 27). Customer 2 cancels `sub_two_a` and starts `sub_two_b`. If an event for `sub_two_a` arrives after an event for `sub_two_b`, `sub_two_b` ends with status `ended`. With `#5 created(sub_two_a)` late (23, 24), `sub_two_a` also stays `active`. With `#6 updated(sub_two_a)` late (26, 27), only `sub_two_b` changes.

**Crash after commit 1 (16 findings, 16 effects).**

5. **The event is dropped and never applied.** For every one of the 16 events, a crash right after the first commit leaves the webhook row in `processing`, and the subscription change the event carries is never made. The retry gets 200.

## Root causes (M = 3)

1. **Subscription events are applied whatever their age.** For `created` and `updated`, the handler finds the row by `stripe_subscription_id` ([`subscription.service.ts:559-562`](https://github.com/getpaykit/paykit/blob/c7572ff7e0116428a43a2f98ad25dc2c26c8927e/packages/paykit/src/subscription/subscription.service.ts#L559-L562)) and calls `syncSubscriptionFromProvider` ([`:602-605`](https://github.com/getpaykit/paykit/blob/c7572ff7e0116428a43a2f98ad25dc2c26c8927e/packages/paykit/src/subscription/subscription.service.ts#L602-L605)). That function sets `status`, `canceled` and `cancel_at_period_end` from the payload with no condition ([`:1593-1605`](https://github.com/getpaykit/paykit/blob/c7572ff7e0116428a43a2f98ad25dc2c26c8927e/packages/paykit/src/subscription/subscription.service.ts#L1593-L1605)). The event's `created` time is not carried into the action at all ([`stripe-provider.ts:517-535`](https://github.com/getpaykit/paykit/blob/c7572ff7e0116428a43a2f98ad25dc2c26c8927e/packages/paykit/src/stripe/stripe-provider.ts#L517-L535)), and a row already `canceled` is not protected. The last event to arrive wins. Effects 1 to 3, findings 17 to 22, 25, 28, 29.
2. **Competing subscriptions are ended in arrival order.** After any non-terminal event, the handler ends every other active subscription of the customer in the same product group ([`subscription.service.ts:642-646`](https://github.com/getpaykit/paykit/blob/c7572ff7e0116428a43a2f98ad25dc2c26c8927e/packages/paykit/src/subscription/subscription.service.ts#L642-L646), [`:688-703`](https://github.com/getpaykit/paykit/blob/c7572ff7e0116428a43a2f98ad25dc2c26c8927e/packages/paykit/src/subscription/subscription.service.ts#L688-L703)). It does not ask which subscription is newer. So a late event about the old subscription ends the current one. A per-row event-time guard (the fix for cause 1) would not stop this: in findings 26 and 27 the late event is the only event for `sub_two_a` in the run. Effect 4, findings 23, 24, 26, 27.
3. **The claim is committed before the work, and a stuck claim blocks redelivery.** `beginWebhookEvent` inserts the event row with status `processing` as its own autocommit statement ([`webhook.service.ts:31-44`](https://github.com/getpaykit/paykit/blob/c7572ff7e0116428a43a2f98ad25dc2c26c8927e/packages/paykit/src/webhook/webhook.service.ts#L31-L44)), outside the transaction that applies the event ([`:163`](https://github.com/getpaykit/paykit/blob/c7572ff7e0116428a43a2f98ad25dc2c26c8927e/packages/paykit/src/webhook/webhook.service.ts#L163)). If the process dies between the two, a redelivery hits the unique index and may reclaim the row only if it is `failed` or has been `processing` for more than 5 minutes ([`:53-70`](https://github.com/getpaykit/paykit/blob/c7572ff7e0116428a43a2f98ad25dc2c26c8927e/packages/paykit/src/webhook/webhook.service.ts#L53-L70)). Otherwise it is skipped as a duplicate ([`:151-154`](https://github.com/getpaykit/paykit/blob/c7572ff7e0116428a43a2f98ad25dc2c26c8927e/packages/paykit/src/webhook/webhook.service.ts#L151-L154)) and answered 200, so Stripe stops retrying. Nothing else reprocesses a `processing` row. stutter retries at once, inside the 5-minute window. A redelivery that came after 5 minutes would reclaim the row and apply the event. Effect 5, findings 1 to 16.

## Failing pairs

| Finding | Pair (sent second first) | Cause |
|---|---|---|
| 17 | `#4 deleted(sub_one_a)`, `#0 created` | 1 |
| 18 | `#4`, `#1 updated(active)` | 1 |
| 19 | `#4`, `#2 updated(past_due)` | 1 |
| 20 | `#4`, `#3 updated(active)` | 1 |
| 21 | `#6 updated(sub_two_a, cancel_at_period_end)`, `#5 created` | 1 |
| 22 | `#7 deleted(sub_two_a)`, `#5 created` | 1 |
| 23 | `#8 created(sub_two_b)`, `#5 created(sub_two_a)` | 2 |
| 24 | `#9 updated(sub_two_b)`, `#5 created(sub_two_a)` | 2 |
| 25 | `#7 deleted(sub_two_a)`, `#6 updated` | 1 |
| 26 | `#8 created(sub_two_b)`, `#6 updated(sub_two_a)` | 2 |
| 27 | `#9 updated(sub_two_b)`, `#6 updated(sub_two_a)` | 2 |
| 28 | `#13 updated(sub_four_a, unpaid)`, `#12 created` | 1 |
| 29 | `#15 updated(sub_five_a, active)`, `#14 created` | 1 |

The other 10 reorder schedules fired and left the final state unchanged. None is a no-op pair. No pair for customer 3 (`sub_three_a`) fails.

## Repro counts

- Reorder findings 17 to 29: 10 of 10 fresh replays, and `repro.sh` (curl, fresh databases) 10 of 10.
- Crash findings 1 to 16: 10 of 10 fresh replays. Not scriptable with curl.
- All 29 verified 1-minimal.

## Deviations from the protocol

- **Subscription rows in the fixture.** The fixture holds `paykit_subscription` rows, not only customers and products. They are derived from the schema, and they stand for the state after checkout. Without them, the 16 events changed no subscription state. The fixture's comment says PayKit "does not create subscription rows from a webhook". After the push, the code shows that is only partly true: it inserts a row when the stored product is a plan in the app's in-code catalog ([`subscription.service.ts:577-596`](https://github.com/getpaykit/paykit/blob/c7572ff7e0116428a43a2f98ad25dc2c26c8927e/packages/paykit/src/subscription/subscription.service.ts#L577-L596)). The demo's catalog is `free`, `pro` and `ultra`, and the fixture's products are `basic` and `plus`, so it never does here. The comment was left as pushed.
- **Branches the hunt did not reach.** Because the fixture's products are not catalog plans and neither is a free default, the "schedule the default plan" and "activate scheduled plan" branches did nothing in every run. Defects there would not show.
- **`ignore` columns.** Two columns were added to `ignore` after the first hunt run named them as unstable. That run produced no results.
- **Hash before the second hunt run.** The workload hash was checked before the first hunt run, not again before the second. It was checked again after the second run and still matched.
- **One file outside the read list.** `scripts/load-root-env.js` (it loads a root `.env` if one exists) was read before the push, because the demo's drizzle and next configs import it. It is a build helper, not application source, but the list does not name it.
- **Smoke runs before recording.** The workload was delivered directly three times to scratch databases to choose fixtures and to run the network check. Those runs are not part of the trace. The scratch databases were dropped.
- **`--sdk` as an absolute path.** A relative path did not resolve from the generator's own directory. It is the same copy.
- **`fetch-targets.sh` header.** The list of targets at the top of the script was not updated, because another hunter had uncommitted edits on those lines. Only the paykit section was committed.
- **API version.** The rule gives `--api basil` for `stripe` 18 or later, and the generator's events say `2025-04-30.basil`. But `stripe@19.3.1` pins `2025-10-29.clover`. The rule was applied as written. The handler reads `current_period_*` from the subscription items, where basil puts them, and does not check the label.
