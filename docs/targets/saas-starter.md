# nextjs/saas-starter

| | |
|---|---|
| Repo | [nextjs/saas-starter](https://github.com/nextjs/saas-starter), 16,252 stars, 2,747 forks (read 2026-11-14) |
| Version | pinned commit `6e33e58b1e553a41fe22e6b941a7229a002de361` (2025-12-11, latest on `main`); the repo has no releases |
| Code changed | none |
| Stack | Next.js 15.6 canary, Drizzle ORM on postgres.js, `stripe@18.1.0`, API version `2025-04-30.basil` |
| Config | [targets/saas-starter/stutter.json](../../targets/saas-starter/stutter.json) |
| Results | [targets/saas-starter/hunt/REPORT.md](../../targets/saas-starter/hunt/REPORT.md) |

## The contract checked, and where the hypothesis came from

**The contract is Stripe's.** In its [Event ordering](https://docs.stripe.com/webhooks#event-ordering) section, Stripe says it "doesn't guarantee the delivery of events in the order that they're generated". Its docs also say endpoints may receive the same event more than once, and that failed deliveries are retried for up to three days. A retried delivery can therefore land after newer events.

**The hypothesis came from reading the code.** In the plan written before any run, `handleSubscriptionChange` was flagged as order-dependent, because it writes the status from the event with no comparison against current state. stutter's part was to:

- confirm the hypothesis under Stripe's delivery contract;
- try every redelivery, concurrent redelivery and reorder pair;
- shrink each failure and produce a curl reproduction;
- report anything else it found.

The cross-subscription effect (2 and 3 below) was not in the plan. It came from a second reading of the same function before the run, which is why the lifecycles include a cancel-and-resubscribe case. Effect 3, the wrong plan, was not anticipated.

## Setup

- **Build.** `next build` prerenders `/pricing`, which calls the live Stripe API. Rather than give it a real key, stutter builds with `next build --experimental-build-mode=compile`, which skips prerendering, then runs `next start`. The webhook route is dynamic either way.
- **Database.** Each run gets a template built by the project's own `drizzle-kit migrate`. The template also holds a labelled fixture: five teams that already have a `stripe_customer_id`, the state checkout leaves behind ([fixtures.sql](../../targets/saas-starter/fixtures.sql)).
- **Events.** There are 16 synthetic `customer.subscription.*` events across five lifecycles: trial conversion, a failed renewal that recovers, cancel and resubscribe, upgrade, and a trial ending unpaid. They are generated in the `2025-04-30.basil` shape, with `current_period_*` on subscription items ([lifecycles.mjs](../../targets/stripe/lifecycles.mjs)).
- **Recording.** During recording, each event is signed with the project's own pinned Stripe SDK. Every one got HTTP 200 in order, which is the pre-flight check.
- **Faults.** On replay, stutter re-signs each delivery at send time, and retries any non-2xx up to 5 times as Stripe does. Reorder pairs are any two events for the same Stripe customer.

## Results

| Fault | Schedules | Fired | Changed the final state |
|---|---|---|---|
| lost (redelivery) | 16 | 16 | 0 |
| concurrent redelivery | 16 | 9 (7 not fired: events the handler ignores never write) | 0 |
| reorder | 23 | 23 | 7 (12 more are no-op pairs: the handler ignores both events) |
| crash after commit | 9 (one per delivery that writes) | 9 | 0 |

Duplicates and crashes are harmless here, because every update is a single idempotent write. **Ordering is not.** The 7 failing reorders shrink to 3 distinct end-state changes. Every one reproduced in 10 of 10 fresh replays, and by the curl script `repro.sh` in 10 of 10 trials on fresh databases.

1. **A canceled subscription comes back as active** (findings 1, 2, 5, 6). If `customer.subscription.deleted` arrives before an earlier `updated` with status `active`, the team ends `active` with the old subscription id. The customer keeps paid access without paying. Minimal case: `updated(active)` and `deleted` for the same subscription, delivered swapped.
2. **A paying customer is marked canceled** (findings 4, 7). A team cancels subscription A and later subscribes to B. If A's `deleted` arrives after B's `updated(active)`, the team ends `canceled` with no subscription, while B is live and billing. Minimal case: two deliveries, `#9 updated(sub_two_b, active)` then `#7 deleted(sub_two_a)`.
3. **The team points at the wrong subscription and plan** (finding 3). A stale `updated` for the old subscription A, arriving after B's, sets the team to A's id and product.

The no-op pairs, such as events involving `past_due`, are reported separately and are not counted as passes. The handler ignores `past_due`, so their passing says nothing about ordering.

## Root cause

`handleSubscriptionChange` in [`lib/payments/stripe.ts:117-147`](https://github.com/nextjs/saas-starter/blob/6e33e58b1e553a41fe22e6b941a7229a002de361/lib/payments/stripe.ts#L117-L147) applies every `customer.subscription.updated` and `.deleted` event unconditionally:

- the status is written straight from the event payload (lines 131 and 139);
- nothing checks that `subscription.id` is the team's current `stripeSubscriptionId`, which is what causes effects 2 and 3;
- nothing checks whether the event is older than the state already stored, which is what causes effect 1.

Stripe advises against using `created` to order events, because distinct events can share a second. The fix it recommends is to treat the event as a signal and retrieve the subscription from the API before writing. A smaller guard closes effects 2 and 3: ignore events for a subscription that isn't the team's current one, unless the event activates it.

Unrelated and already reported upstream (#294, #295): `planName` is never set, because webhook payloads carry `plan.product` as an id string. This shows up as `plan_name: null` in both runs, so it is not a stutter finding.

## Upstream issue draft (not posted)

> **Title:** Out-of-order subscription webhooks can reactivate canceled teams or cancel paying ones
>
> Stripe doesn't guarantee webhook delivery order, and it retries failed deliveries for up to three days, so an older event can arrive after a newer one. `handleSubscriptionChange` (lib/payments/stripe.ts:117) writes the event's status without checking the team's current subscription or whether the event is stale. Three effects reproduce on `main` (6e33e58) with signed test events and no other changes:
>
> 1. `customer.subscription.deleted` delivered before an earlier `customer.subscription.updated` (status `active`) for the same subscription leaves the team `active` after it canceled.
> 2. A team cancels subscription A and subscribes to B. If A's `deleted` arrives after B's `updated(active)`, the team ends `canceled` with no subscription while B is live.
> 3. A stale `updated` for A arriving after B's sets the team back to A's subscription and product.
>
> A reproduction script with the event bodies is attached: it starts from a team with `stripe_customer_id` set and sends two or three signed events. A possible fix is to retrieve the subscription from Stripe inside the handler (Stripe's recommendation in its event-ordering docs), or at least ignore events whose `subscription.id` differs from `teams.stripe_subscription_id` unless the event activates that subscription.
