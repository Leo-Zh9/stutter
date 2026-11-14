# boxyhq/saas-starter-kit

| | |
|---|---|
| Repo | [boxyhq/saas-starter-kit](https://github.com/boxyhq/saas-starter-kit), 4,977 stars, 1,241 forks (read 2026-11-14) |
| Version | pinned commit `abc9b686823cbfb4973c79bc36fea37a3244be6c` (2026-05-08, latest on `main`) |
| Code changed | none |
| Stack | Next.js 15.5 (pages router), Prisma 6.10, `stripe@17.7.0`, API version `2025-02-24.acacia` |
| Config | [targets/boxyhq/stutter.json](../../targets/boxyhq/stutter.json) |
| Results | [targets/boxyhq/hunt/REPORT.md](../../targets/boxyhq/hunt/REPORT.md) |

## The contract checked, and where the hypothesis came from

The contract is the same as for [saas-starter](saas-starter.md): Stripe documents at-least-once, unordered webhook delivery, with retries for up to three days.

The hypothesis also came from reading the code before the run:

- `customer.subscription.deleted` hard-deletes the subscription row;
- a later `customer.subscription.updated` for a row that no longer exists falls back to `handleSubscriptionCreated`, which hard-codes `active: true`.

stutter's part was to confirm this under Stripe's delivery contract, try every redelivery, concurrent redelivery, reorder pair and crash point, shrink each failure, and produce curl reproductions. Effect 3 below, a trial that ended unpaid becoming active, was not anticipated.

## Setup

- **Build.** As with saas-starter, `next build --experimental-build-mode=compile` skips prerendering, so nothing calls the live Stripe API.
- **Database.** Each run gets a template built by the project's own `prisma migrate deploy` (12 migrations), plus a labelled fixture of five teams with a Stripe `billingId` ([fixtures.sql](../../targets/boxyhq/fixtures.sql)).
- **Events.** The same 16 lifecycle events as saas-starter, in the `2025-02-24.acacia` shape that `stripe@17.7.0` pins: `current_period_*` on the subscription object, which is where this handler reads them.
- **Recording.** The events are signed with the project's own SDK, and every one got HTTP 200 in order.

## Results

| Fault | Schedules | Fired | Changed the final state |
|---|---|---|---|
| lost (redelivery) | 16 | 16 | 0 |
| concurrent redelivery | 16 | 16 | 0 |
| reorder | 23 | 23 | 11 (7 more are no-op pairs) |
| crash after commit | 16 | 16 | 0 |

Duplicates and crashes are harmless here. A second `created` hits the primary key, the handler answers 400, and Stripe-style retries end in the same state. **Ordering is not.** The 11 failing reorders shrink to 3 distinct end-state changes. Every finding reproduced in 10 of 10 fresh replays, and by the curl script `repro.sh` in 10 of 10 trials on fresh databases:

1. **A canceled subscription comes back as active.** The `deleted` event arrives before an earlier `created` or `updated`. The delete finds nothing to delete, and the late event then creates the row with `active: true`.
2. **A resubscribed team keeps its old canceled subscription marked active.** The team cancels A and subscribes to B. If A's `deleted` arrives before A's own `created` or `updated`, A ends up recreated and active next to B. Features that check for any active subscription would see two.
3. **A trial that ended unpaid becomes active.** If `updated(unpaid)` arrives before `created(trialing)`, the handler finds no row and creates one with `active: true`, ignoring the event's status. The late `created` then fails on the primary key, and every Stripe retry fails the same way, so the subscription stays active.

## Root cause

[`pages/api/webhooks/stripe.ts`](https://github.com/boxyhq/saas-starter-kit/blob/abc9b686823cbfb4973c79bc36fea37a3244be6c/pages/api/webhooks/stripe.ts):

- **line 61:** `deleted` removes the row (`deleteMany`), so no record says the subscription ended, and a late event cannot tell it is stale;
- **lines 91-97:** `handleSubscriptionUpdated` treats a missing row as new and calls `handleSubscriptionCreated`;
- **line 124:** `handleSubscriptionCreated` writes `active: true` whatever the event's status.

Following Stripe's advice, the fix is to retrieve the subscription from the API and upsert its current status. At minimum: keep canceled subscriptions as rows with `active: false` rather than deleting them, and derive `active` from `status` in every path.

Also seen while reading the code, and not a stutter finding because it is the same with or without faults: a subscription created in `trialing` status is stored as `active: true`.

## Upstream issue draft (not posted)

> **Title:** Out-of-order Stripe webhooks can recreate canceled subscriptions as active
>
> Stripe doesn't guarantee webhook delivery order, and it retries failed deliveries for up to three days. In `pages/api/webhooks/stripe.ts`, `customer.subscription.deleted` deletes the row (line 61), and `handleSubscriptionUpdated` recreates a missing row through `handleSubscriptionCreated`, which hard-codes `active: true` (lines 91-97, 124). Three effects reproduce on `main` (abc9b68) with signed test events and no other changes:
>
> 1. `deleted` delivered before an earlier `updated` or `created` for the same subscription leaves an active subscription row after the cancel.
> 2. A team that cancels subscription A and subscribes to B ends with both A and B active if A's events arrive out of order.
> 3. `updated` with status `unpaid` delivered before `created` (`trialing`) creates the row as active. The late `created` then fails on the primary key on every retry, so it stays active.
>
> A reproduction script with the event bodies is attached. A possible fix is to retrieve the subscription from Stripe in the handler and upsert `active` from its current `status`, and to mark canceled subscriptions inactive instead of deleting them.
