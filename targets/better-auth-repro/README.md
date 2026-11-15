# @better-auth/stripe: out-of-order webhooks

A standalone reproduction of two bugs stutter found in `@better-auth/stripe` through [kriasoft/react-starter-kit](../../docs/targets/react-starter-kit.md). It needs no Postgres, no app and no Stripe account: Better Auth runs on its in-memory adapter, and the webhooks are signed locally with the Stripe SDK's test helper.

```bash
npm ci
npm run repro     # exits 1 while either bug is present
```

Pinned versions: `better-auth`, `@better-auth/core` and `@better-auth/stripe` 1.7.25 (the latest release on 2026-11-15), `stripe` 23.1.0.

Each scenario runs twice on a fresh database, in Stripe's order and with two events swapped, and prints the saved rows. Every delivery returns 200 in both runs, so Stripe would never retry.

| Scenario | In order | Swapped |
|---|---|---|
| 1. A customer with a Starter subscription adds a Pro one; the Pro `updated` arrives before its `created` | `sub_a` starter active, `sub_b` pro active | only `sub_b` pro active: the `sub_a` row was rewritten into `sub_b` |
| 2. A subscription is created and canceled; the `deleted` arrives first | `sub_a` canceled | `sub_a` active |

Causes, in `packages/stripe/src/hooks.ts` on `main` at `8382ce6`:

1. `onSubscriptionUpdated` finds no row for the event's subscription, falls back to the customer's rows (`:353-377`), takes the only one, and writes the event onto it, including `stripeSubscriptionId` (`:423`). The later `created` then finds that row and skips.
2. `onSubscriptionDeleted` logs "Subscription not found" and returns (`:557-560`), so the later `created` inserts a live row.

Check of cause 1: restricting the fallback to rows with no `stripeSubscriptionId` (one `.filter` on the `findMany` result in the built package) makes scenario 1 print the same rows in both orders. Scenario 2 is unchanged by that patch.

Subscriptions here carry no Better Auth metadata, as when they are created in the Stripe Dashboard, which `onSubscriptionCreated` supports by creating the row itself.
