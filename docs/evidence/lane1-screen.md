# Lane 1: a documented screen of past fixes

## What this is, and what it is not

This is a documented screen. It is not a pre-registration.

Most of the screening happened before these rules were frozen. About 14 of the 18 repositories below were searched first, and stripe/sync-engine #177 was found and picked as the case study during that search. The rules were then written to describe the screen, so a reader should treat #177 as a case study, not as a sample.

What the screen does show is every fix found, and why each one was or was not run. Nothing found is left out.

## Universe

The 18 repositories, frozen as listed:

medusajs/medusa, saleor/saleor, getlago/lago-api, juspay/hyperswitch, calcom/cal.diy (formerly calcom/cal.com), documenso/documenso, dubinc/dub, twentyhq/twenty, makeplane/plane, TryGhost/Ghost, stripe/sync-engine (formerly supabase/stripe-sync-engine), vendurehq/vendure (formerly vendure-ecommerce/vendure), pretix/pretix, polarsource/polar, formancehq/ledger, lukevella/rallly, mfts/papermark, elie222/inbox-zero.

## Queries

Per repository R and keyword k:

```bash
gh search prs "<k>" --repo R --merged --limit 50 --json number,title
```

k is one of: `idempotency`, `idempotent`, `duplicate`, `race`, `"out of order"`, `webhook`, `concurrent`, `stale`, `twice`, `retry`.

Titles were kept if they matched `idempot|duplicat|double|race|concurren|out.of.order|twice|retry|replay|lock|stale`, case-insensitive. Kept PRs were then read by body.

The first pass used only the first seven keywords. It did not save the raw match lists.

## Inclusion rules

A PR is included if all four hold:

- **(I1)** Its title, body or linked issue names a defect caused by a retried, redelivered, concurrently duplicated or reordered request, or by a crash between two commits.
- **(I2)** The duplicated request is a webhook or carries an Idempotency-Key.
- **(I3)** PostgreSQL is the store the defect touches.
- **(I4)** The fix is merged.

When unsure, include.

Y is the included PRs that can run here: no Docker, no Redis, no PostgreSQL contrib extension, and no fake of a provider's API. Included PRs that cannot run here are listed by name. A runnable PR that fails to start is counted as not run, never dropped.

## Results so far

| Repo | PR | Defect | Fault | Status |
|---|---|---|---|---|
| stripe/sync-engine | [#177](https://github.com/stripe/sync-engine/pull/177) | An older webhook overwrites newer data | reorder | **Included, runnable.** Case study. Pre-fix 96e6b39, fix d486b41. |
| documenso/documenso | [#2977](https://github.com/documenso/documenso/pull/2977) | Per-event Stripe handlers replaced by a sync | reorder, redelivery | Included, not runnable. A migration needs `pg_trgm`. The fix calls the Stripe API. |
| polarsource/polar | [#13325](https://github.com/polarsource/polar/pull/13325) | Concurrent webhook race in a payment intent upsert | concurrent | Included, not runnable. Needs Redis. Webhooks are processed by an async worker. |
| vendurehq/vendure | [#2462](https://github.com/vendurehq/vendure/pull/2462) | A duplicate Mollie `paid` webhook throws an error | redelivery | Included, not runnable. The handler calls the Mollie API. The effect is an HTTP status, which a state oracle would likely miss. |
| medusajs/medusa | [#1262](https://github.com/medusajs/medusa/pull/1262) | Duplicate tax lines after a failed cart completion | none in model | **Excluded (I2).** The fix's test sends two unkeyed `/complete` calls, and the first returns 400, so the duplicate is already in the fault-free run. On the keyed path, the tax lines and the recovery point are written in one transaction, so a crash cannot split them. |
| medusajs/medusa | [#2341](https://github.com/medusajs/medusa/pull/2341) | Webhook and storefront both complete a cart | two different requests | Excluded (I1). stutter only duplicates the same request. |
| medusajs/medusa | [#16097](https://github.com/medusajs/medusa/pull/16097) | Concurrent captures exceed the authorized amount | unkeyed concurrent | Excluded (I2). |
| calcom/cal.diy | [#13865](https://github.com/calcom/cal.diy/pull/13865) | Duplicate bookings | unkeyed | Excluded (I2). |
| saleor/saleor | [#15956](https://github.com/saleor/saleor/pull/15956), [#17776](https://github.com/saleor/saleor/pull/17776) | Races between different checkout mutations | two different requests | Excluded (I1). |
| vendurehq/vendure | [#2402](https://github.com/vendurehq/vendure/pull/2402) | Stripe webhook not wrapped in one transaction | none in model | Excluded (I1). The stated cause is read-replica lag. |
| vendurehq/vendure | [#2691](https://github.com/vendurehq/vendure/pull/2691) | Duplicate payments from several browser tabs | two different requests | Excluded (I1). |
| formancehq/ledger | [#919](https://github.com/formancehq/ledger/pull/919) | Key reused with a different body returns 500 | none in model | Excluded (I1). A key reused with a different body is wrong with or without faults. |
| lukevella/rallly | [#1894](https://github.com/lukevella/rallly/pull/1894) | Race between checkout and a webhook creating a customer | two different requests | Excluded (I1). |
| twentyhq/twenty | [#28084](https://github.com/twentyhq/twenty/pull/28084) | Subscription creation made idempotent | not yet read | Body not screened. The repo needs Redis. |
| getlago/lago-api | [#5708](https://github.com/getlago/lago-api/pull/5708), [#5523](https://github.com/getlago/lago-api/pull/5523) | Double charge from open checkout sessions; double invoice finalization | not yet read | Bodies not screened. The repo needs Ruby 3, Redis and Sidekiq. |

Other title matches in the screened repositories were dropped by title as not about duplicated or reordered requests. Those lists were not saved, so the rerun below replaces them.

## To finish

1. Screen the four repositories not yet searched: juspay/hyperswitch, dubinc/dub, makeplane/plane, TryGhost/Ghost.
2. Rerun all 10 keywords on every repository, including `stale`, `twice` and `retry`, which the first pass did not use. Commit the raw match lists.
3. Rerun mfts/papermark, which returned no matches in the first pass. That may have been a failed search.
4. Read the bodies of twentyhq/twenty #28084 and getlago/lago-api #5708 and #5523, and place them under I1 to I4.
