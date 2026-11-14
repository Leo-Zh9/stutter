# Lane 1 results: the finished screen

This finishes the screen described in [lane1-screen.md](lane1-screen.md). That file stays as frozen. Raw search output is in [lane1-raw/](lane1-raw), one JSON file per repository, with each query's command and run time. The queries ran on 2026-11-14 between 23:14 and 23:53 UTC.

## Funnel

All counts come from the raw files.

| Stage | Count | Notes |
|---|---|---|
| Queries | 200 | 18 repositories × 10 keywords as frozen, plus 20 repeats under new names (below). |
| Failed queries | 20 | `dubinc/dub` and `mfts/papermark` were renamed. GitHub search does not follow renames, so every query on the old name failed. The errors are kept in the raw files. The repeats under `dubinc/dub.link` and `papermark/papermark` succeeded. |
| Hits | 7,021 | All results from the 180 queries that succeeded. |
| Unique PRs | 5,303 | Hits deduplicated per repository. |
| Kept by title | 860 | Titles that match the frozen regex. |
| Screened | 860 | 108 of them (Ghost 84, Dub 24) fail I3 at repository level. Ghost depends on `mysql2` and `better-sqlite3`, with no PostgreSQL driver. Dub's Prisma datasource is `mysql`. 44 PRs were read in full. The rest were judged by title. |
| Included (I1 to I4) | 9 | Plus #177, which the title filter drops (below). |
| In model | 1 | Plus #177. "In model" means a stutter fault can trigger the defect and its effect lands in database state. 3 more are unknown. |
| Runnable here | 0 in model | #177 is runnable. Of the included set, only formance #1624 is runnable, and it is not in model. |
| Run | 0 | |
| Caught | 0 | #177 is pending. |

## Two limits of the frozen queries

1. **The 50-result cap truncates them.** 117 of the 180 queries that succeeded returned exactly 50 results, the frozen `--limit`, so more matches existed. Older PRs in large repositories are not reached. medusa #1262 and #2341, found in the first pass with a date filter, do not appear in this rerun at all.
2. **The title filter drops the case study.** stripe/sync-engine #177 appears in the hits for `out of order` and `webhook`. Its title, "chore: add upsert timestamp protection", does not match the frozen regex. It is listed below as included, but it entered through its body, not through the filter.

## Included PRs

| Repo | PR | Defect | Fault that would expose it | In model | Runnable here |
|---|---|---|---|---|---|
| stripe/sync-engine | [#177](https://github.com/stripe/sync-engine/pull/177) | An older webhook overwrites newer data | reorder | Yes | Yes. Result: **pending (case study, docs/targets/sync-engine.md)** |
| documenso/documenso | [#2977](https://github.com/documenso/documenso/pull/2977) | Per-event Stripe subscription handlers, replaced by a sync | reorder, redelivery | Yes | No. A migration needs `pg_trgm`, and the fix calls the Stripe API. |
| formancehq/ledger | [#1624](https://github.com/formancehq/ledger/pull/1624), backport of [#1601](https://github.com/formancehq/ledger/pull/1601) | A revert whose key is still in flight returns 500 instead of 409 | concurrent | No. Only the status code changes, and both are retried, so the final state is the same. | Yes. This is the v2.4 PostgreSQL line, and v2.4.18 already contains the fix. |
| vendurehq/vendure | [#2462](https://github.com/vendurehq/vendure/pull/2462) | A duplicate Mollie `paid` webhook returns an error | lost (redelivery) | No. Only the response changes. | No. The handler calls the Mollie API. |
| polarsource/polar | [#12344](https://github.com/polarsource/polar/pull/12344) | A concurrent duplicate webhook returns 500 on a unique index | concurrent | No. The unique index holds, so the final state is the same. | No. Needs Redis and an async worker. |
| polarsource/polar | [#13325](https://github.com/polarsource/polar/pull/13325) | Two different webhooks for one payment intent race to insert | none: stutter does not send two different webhooks at once | No | No. Same as above. |
| juspay/hyperswitch | [#14489](https://github.com/juspay/hyperswitch/pull/14489), backport [#14507](https://github.com/juspay/hyperswitch/pull/14507) | A concurrent duplicate dispute webhook returns 500 | concurrent | No. Only the response changes. | No. Rust and Redis. |
| juspay/hyperswitch | [#10523](https://github.com/juspay/hyperswitch/pull/10523) | Concurrent payout webhooks for the same payout | concurrent | Unknown. The body does not say whether these are duplicates or what the effect is. | No. Rust and Redis. |
| juspay/hyperswitch | [#319](https://github.com/juspay/hyperswitch/pull/319) | A repeated client-supplied `payment_id` was not rejected | concurrent or lost, if `payment_id` acts as a key | Unknown. Included because unsure under I2. | No. Rust and Redis. |
| pretix/pretix | [#1572](https://github.com/pretix/pretix/pull/1572) | Race when failing a payment | reorder, if the race is between webhooks | Unknown. The PR has no body. Included because unsure. | No. Its Stripe webhooks call `stripe.Charge.retrieve` and `stripe.PaymentIntent.retrieve`. |

## PRs read in full and excluded

| Repo | PR | Rule | Reason |
|---|---|---|---|
| medusajs/medusa | [#9231](https://github.com/medusajs/medusa/pull/9231) | I2 | The `/store/carts/:id/complete` route takes no idempotency key. The payment webhook runs in an async subscriber. |
| medusajs/medusa | [#2358](https://github.com/medusajs/medusa/pull/2358) | I1 | A failing stage commits partial data. That happens with or without a fault. |
| medusajs/medusa | [#2612](https://github.com/medusajs/medusa/pull/2612) | I2 | Adds idempotency to a route that had none. Before the fix, a retry was the client's fault. |
| medusajs/medusa | [#16097](https://github.com/medusajs/medusa/pull/16097) | I2 | Concurrent captures on an unkeyed admin route. |
| medusajs/medusa | [#16964](https://github.com/medusajs/medusa/pull/16964), [#16951](https://github.com/medusajs/medusa/pull/16951) | I1 | Retrying after a provider failure. No duplicated request causes it. |
| medusajs/medusa | [#14073](https://github.com/medusajs/medusa/pull/14073), [#10135](https://github.com/medusajs/medusa/pull/10135) | I1 | Avoid needless provider calls. No state defect. |
| medusajs/medusa | [#13841](https://github.com/medusajs/medusa/pull/13841) | I1 | Every update recreates addresses, with or without faults. |
| medusajs/medusa | [#2947](https://github.com/medusajs/medusa/pull/2947) | I2 | Retries of internal queue jobs, not HTTP requests. |
| getlago/lago-api | [#5708](https://github.com/getlago/lago-api/pull/5708) | I1 | Two different payment paths, checkout and auto-charge, race. |
| getlago/lago-api | [#5523](https://github.com/getlago/lago-api/pull/5523), [#6485](https://github.com/getlago/lago-api/pull/6485) | I1 | Background jobs race. |
| getlago/lago-api | [#5495](https://github.com/getlago/lago-api/pull/5495), [#7000](https://github.com/getlago/lago-api/pull/7000) | I2 | Concurrent duplicates of unkeyed API calls. |
| getlago/lago-api | [#6911](https://github.com/getlago/lago-api/pull/6911) | I1 | A replay was answered to a different request's key. That is wrong with or without faults. |
| getlago/lago-api | [#6444](https://github.com/getlago/lago-api/pull/6444), [#3578](https://github.com/getlago/lago-api/pull/3578) | I1 | Lookup scoping, and an internal guard. |
| juspay/hyperswitch | [#2916](https://github.com/juspay/hyperswitch/pull/2916) | I1 | A webhook races a different API call. |
| juspay/hyperswitch | [#14359](https://github.com/juspay/hyperswitch/pull/14359), [#8034](https://github.com/juspay/hyperswitch/pull/8034), [#14931](https://github.com/juspay/hyperswitch/pull/14931), [#9963](https://github.com/juspay/hyperswitch/pull/9963), [#13386](https://github.com/juspay/hyperswitch/pull/13386) | I1 | Internal scheduling, outgoing webhooks, or races between different operations. |
| documenso/documenso | [#1721](https://github.com/documenso/documenso/pull/1721) | I2 | The frontend sent duplicate calls, which fired outgoing webhooks twice. |
| twentyhq/twenty | [#28084](https://github.com/twentyhq/twenty/pull/28084) | I2 | A double click or two tabs send unkeyed requests. The key in the fix is outgoing, to Stripe. |
| twentyhq/twenty | [#27953](https://github.com/twentyhq/twenty/pull/27953) | I2 | Queue job retries. |
| twentyhq/twenty | [#25304](https://github.com/twentyhq/twenty/pull/25304) | I1 | Adds logging. The PR says a unique index already dedupes the redeliveries. |
| makeplane/plane | [#6893](https://github.com/makeplane/plane/pull/6893) | I2 | A double submit in the UI. |
| vendurehq/vendure | [#2691](https://github.com/vendurehq/vendure/pull/2691) | I1 | Several browser tabs send different requests. |
| pretix/pretix | [#5320](https://github.com/pretix/pretix/pull/5320) | I1 | Changes which state transitions are allowed. |
| polarsource/polar | [#14274](https://github.com/polarsource/polar/pull/14274) | I1 | A failure webhook races a later client request. |
| lukevella/rallly | [#1894](https://github.com/lukevella/rallly/pull/1894) | I1 | Checkout and a webhook both create a customer. |
| elie222/inbox-zero | [#4268](https://github.com/elie222/inbox-zero/pull/4268) | I3 | The lost message comes from a processing lock held in Redis. |

PRs read in the first pass that the frozen queries do not return are not counted in the funnel. Their status from [lane1-screen.md](lane1-screen.md) stands: medusa #1262 (I2), #2341 (I1); vendure #2402 (I1); formance #919 (I1).

## Title-level exclusions

The other 708 kept titles were judged by title alone. That is 860 kept, less 108 excluded at repository level and 44 read in full. They fall into these groups:

- tests, docs, CI, dependency and UI changes;
- "stale" or "lock" in an unrelated sense, such as stale caches, lockfiles, blocklists and UI state;
- duplication features, such as "duplicate a document" or "duplicate a poll";
- races between different operations or background jobs;
- duplicates of unkeyed client requests, such as Saleor's GraphQL mutations (for example [#19889](https://github.com/saleor/saleor/pull/19889)) and cal.com's server-derived booking keys (for example [#13865](https://github.com/calcom/cal.diy/pull/13865)).

The per-PR titles are in the raw files.

## What this supports

On this screen, one past fix is both in model and runnable here: stripe/sync-engine #177, and it entered through its body rather than the title filter. The rest fail for one of three reasons:

- the duplicated request carries no key;
- the stack needs Rust, Ruby, Redis, a PostgreSQL contrib extension or a provider's API;
- the effect is only an HTTP status, which a database-state oracle cannot see.

The last group includes the one other runnable PR, formance #1624. Lane 1 therefore stays a case study plus this table. It is not a "caught X of Y" rate.
