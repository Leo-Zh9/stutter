# kriasoft/react-starter-kit, blind result: found

| | |
|---|---|
| Repo | [kriasoft/react-starter-kit](https://github.com/kriasoft/react-starter-kit), default branch `main` |
| Pinned SHA | `dbdb0d4991bbaae1671ff62180089f2733dd406a` |
| Stack | Hono API run by `apps/api/dev.ts` under Bun 1.4.2, Wrangler `getPlatformProxy` with a local Hyperdrive binding (workerd), Drizzle 0.45 with postgres.js, Better Auth 1.7.4 with `@better-auth/stripe` 1.7.4, `stripe@22.6.2` in `apps/api` |
| Config | [stutter.json](../../targets/react-starter-kit/stutter.json), [fixtures.sql](../../targets/react-starter-kit/fixtures.sql) |
| Results | [hunt/REPORT.md](../../targets/react-starter-kit/hunt/REPORT.md), [hunt/results.json](../../targets/react-starter-kit/hunt/results.json), [hunt.log](../../targets/react-starter-kit/hunt.log), [record.log](../../targets/react-starter-kit/record.log), [netcheck/](../../targets/react-starter-kit/netcheck) |
| Outcome | Found. 15 of 23 reorder pairs change the final state. Lost, concurrent and crash faults: 0 of 48. M = 4. |

## This is a blind result

The repository came from the screener as a name and a pinned SHA. The hunter configured it from its docs, env file and schema only, recorded the frozen workload, hunted, and pushed the results before reading any application source.

The webhook handler is not in this repository. The app mounts the `@better-auth/stripe` plugin (`apps/api/lib/auth.ts:121-158`) and configures two plans. Every webhook write is done by that plugin, version 1.7.4, as installed by the app's own lockfile. All four root causes below are in that package.

## Blind protocol

- **Pinned SHA.** `dbdb0d4991bbaae1671ff62180089f2733dd406a`, used and no other.
- **Push before any source reading.** Config, fixtures, trace, network check and results were committed as `1c2d81d29f3cd183efa327363c4eee66a0d66b94` and pushed to `origin/main`. The handler and the plugin were read only after that push.
- **What was read before the push.**
  - Docs: `README.md`, `apps/api/README.md`, `docs/billing/index.md`, `docs/billing/webhooks.md`, `docs/billing/plans.md`, `docs/specs/billing.md`, `docs/getting-started/environment-variables.md`.
  - Env example: the committed `.env` (the repo has no `.env.example`; its README says `.env` holds the shared placeholders).
  - Schema: `db/migrations/0000_init.sql`, `db/migrations/meta/_journal.json`, `db/schema/subscription.ts`, and `db/schema/id.ts` (the id generator the schema files import).
  - Package and build config: `package.json`, `apps/api/package.json`, `db/package.json`, the first 30 lines of `apps/email/package.json`, `db/drizzle.config.ts`, `apps/api/wrangler.jsonc`.
  - A tree listing (`git ls-files`). The webhook path `/api/auth/stripe/webhook` is stated in the billing docs.
  - Not read on purpose, but seen: one Bun error at startup printed a stack trace with six lines of `apps/api/lib/db.ts` (lines 13 to 18, the start of `createDb`) and a frame naming `apps/api/dev.ts:66`. See the deviations.
  - No code search was run. Nothing about the target was read online.
- **Generator flags.**
  - `--api basil`: `apps/api/package.json` pins `stripe` at `^22.6.2` (18 or later).
  - `--sdk /Users/leozhang/stutter/.targets/react-starter-kit/apps/api/node_modules/stripe`: the app's own copy, `stripe@22.6.2`.
  - `--customers cus_team1,cus_team2,cus_team3,cus_team4,cus_team5`: the ids the fixture holds (the same as the generator default).
- **Workload hash.** Checked before the recording and before each hunt. All three checks print `targets/stripe/lifecycles.mjs: OK` ([record.log](../../targets/react-starter-kit/record.log), [hunt.log](../../targets/react-starter-kit/hunt.log)).

## Setup

- **Fetch and build.** [scripts/fetch-targets.sh](../../scripts/fetch-targets.sh) clones the pinned SHA into `.targets/react-starter-kit`, runs `npx -y bun@1.4.2 install --frozen-lockfile`, then builds `@repo/email` with `tsc`. The API's dev server imports the built email package, and fails to start without it.
- **Start.** `apps/api/dev.ts` under `npx --offline -y bun@1.4.2 run`, in `apps/api` (which has no `.env` for Bun to load). It reads `PORT` and serves `/health`, which queries Postgres. It reaches Postgres through Wrangler's local Hyperdrive binding, read from `CLOUDFLARE_HYPERDRIVE_LOCAL_CONNECTION_STRING_HYPERDRIVE_{CACHED,UNCACHED}`.
- **Password in the URL.** With stutter's password-less `{DATABASE_URL}`, the binding came up undefined and `/health` returned 500. The env example's URLs carry a password, and with any password the binding works (the local cluster trusts TCP). So `start` is a `sh -c` wrapper that inserts `:postgres` into the URL. The same wrapper covers the crash proxy's URL, which has the same prefix.
- **Database.** The template is built by the app's own `drizzle-kit migrate` (one migration, `0000_init`). No extensions.
- **Fixtures.** Five `user` rows with `stripe_customer_id` `cus_team1` to `cus_team5`. That is the state after sign-up, since the docs set `createCustomerOnSignUp: true`. No organizations and no subscription rows.
- **Plans.** `STRIPE_STARTER_PRICE_ID=price_prod_basic` and `STRIPE_PRO_PRICE_ID=price_prod_plus`, the two prices the workload sends. Without them billing is off and the route returns 404.
- **Offline.** Stripe and Resend keys are placeholders. `WRANGLER_SEND_METRICS=false`. Bun is started with `npx --offline`, from the npm cache filled by the first `npx -y bun@1.4.2`, so npx does not contact the registry on each start.
- **Network check.** Before recording, the 16 events were delivered once to a server started with the same command plus a preload ([netcheck/netwatch-bun.cjs](../../targets/react-starter-kit/netcheck/netwatch-bun.cjs)) that logs `fetch`, `net`/`tls` connects, `http(s)` requests, `Bun.connect` and DNS lookups. The bun process made 16 connects to `127.0.0.1:54329` and one connect to an inherited file descriptor (`{"fd":13}`), and nothing else. workerd is a native binary the preload cannot see, so `lsof` polled the whole process tree every 0.2 s: every socket was on `127.0.0.1` ([netcheck/lsof-summary.txt](../../targets/react-starter-kit/netcheck/lsof-summary.txt)). All 16 deliveries got 200.
- **Recording.** 16 deliveries, all HTTP 200. The pre-flight in the report confirms every recorded delivery got 2xx.
- **Oracle.** `subscription.id` is minted in app code (`sub_` plus a cuid2), not by a column default. The first hunt stopped at the stability check, naming it. It was added to `ignore`. Nothing references it.
- **Faults.** stutter re-signs each delivery, retries any non-2xx up to 5 times, and reorders pairs of events for the same Stripe customer. `-lanes 2`.

## Results

| Fault | Schedules | Fired | Not fired | Changed the final state |
|---|---|---|---|---|
| lost (redelivery) | 16 | 16 | 0 | 0 |
| concurrent redelivery | 16 | 16 | 0 | 0 |
| reorder | 23 | 23 | 0 | 15 (3 more are no-op pairs) |
| crash after commit | 16 | 16 | 0 | 0 |

The no-op pairs are `reorder(#1,#2)`, `reorder(#1,#3)` and `reorder(#2,#3)`. The hunt exits 1 with 15 findings in 1125 s.

Duplicates and crashes are harmless: `created` skips a subscription that already exists, and `updated` and `deleted` write the same values again. **Ordering is not.** The 15 failing reorders shrink to 12 distinct end-state changes, which fall into three effects. Every finding is 1-minimal, reproduced in 10 of 10 fresh replays, and reproduced by its curl script `repro.sh` in 10 of 10 trials on fresh databases.

### Effects

1. **A canceled subscription comes back** (findings 1, 2, 8, 9, 10, 11, 12). Customer 1's subscription ends `trialing` (1), `active` (8, 10) or `past_due` (9) instead of `canceled`. Customer 2's first subscription ends `active` (2, 11), or `active` with `cancel_at_period_end: true` (12). Minimal cases: `#4 deleted(sub_one_a)` then `#0 created(sub_one_a)`, and `#0`, `#4 deleted`, then `#1 updated(active)`.
2. **An event is lost and an older state stays** (findings 3, 5, 6, 7). An `updated` that arrives before its `created` is dropped. Customer 3's upgrade to Pro is lost and the row stays `starter` (5). Customer 4's trial that ended `unpaid` shows `trialing` (6). Customer 5's trial that converted to `active` shows `trialing` (7). In finding 3, customer 2 ends on the old Starter subscription instead of the new Pro one.
3. **One subscription's row is taken over by another** (findings 4, 13, 14, 15). An `updated` for a subscription with no row rewrites the customer's other row, including its Stripe subscription id. In findings 14 and 15, `updated(sub_two_b)` turns the `sub_two_a` row into `sub_two_b`, `created(sub_two_b)` then sees it exists and skips, and customer 2 ends with one row instead of two: the canceled `sub_two_a` row is gone (15), or the active one is (14). In findings 4 and 13, the late `updated(sub_two_a)` turns the paying `sub_two_b` Pro row into `sub_two_a` Starter, `cancel_at_period_end: true`.

## Root causes (read after the push)

All four are in `@better-auth/stripe` 1.7.4, `node_modules/@better-auth/stripe/dist/index.mjs` as installed. The route returns 200 whenever a handler returns, and each handler catches its own errors, so Stripe never redelivers (`index.mjs:409`, `:457`, `:1592`). M = 4.

1. **`updated` is applied unconditionally** (`index.mjs:364-389`). It overwrites status, periods, cancellation fields and plan with the event's values, with no check of `event.created` or of a terminal status. A late `updated` undoes a later `deleted`. Failing pairs: `reorder(#1,#4)`, `reorder(#2,#4)`, `reorder(#3,#4)`, `reorder(#6,#7)` (findings 8, 9, 10, 12).
2. **`deleted` for a row that does not exist yet is dropped** (`index.mjs:418-456`). It logs "Subscription not found" and returns 200. The later `created` then inserts a live subscription. Failing pairs: `reorder(#0,#4)`, `reorder(#5,#7)`, `reorder(#5,#8)` (findings 1, 2, 11).
3. **`updated` for a customer with no row crashes and is acknowledged** (`index.mjs:342-357`, `:387`). The fallback lookup by customer returns no rows, `subscription` is `undefined`, and `subscription.id` throws a `TypeError`. The `catch` at `:409` logs it and the route returns 200. Nothing creates the row from the event. Failing pairs: `reorder(#5,#9)`, `reorder(#10,#11)`, `reorder(#12,#13)`, `reorder(#14,#15)` (findings 3, 5, 6, 7).
4. **The customer fallback picks another subscription's row** (`index.mjs:342-357`, `:381`). When no row has the event's subscription id, `updated` takes the customer's only row, or its active one, and writes the event onto it, including `stripeSubscriptionId`. With two subscriptions per customer, an early or late event for one rewrites the other. Failing pairs: `reorder(#6,#8)`, `reorder(#6,#9)`, `reorder(#8,#9)`, `reorder(#7,#9)` (findings 4, 13, 14, 15).

Causes 2 and 3 have the same shape, an early event for an unknown subscription that is dropped with a 200, but they sit in different handlers and fail differently (a deliberate warning, and an unhandled `undefined`). Each needs its own change, so they are counted separately. Counted as one, M would be 3.

The causes were checked against the running app after reading. Re-running findings 1, 5, 4 and 14 by `repro.sh` on fresh databases gave the predicted rows. The server logged "Subscription not found for subscriptionId: sub_one_a" for finding 1, and "TypeError: undefined is not an object (evaluating 'subscription.id')" for finding 5.

## Repro counts

| Finding | Pair | Root cause | Fresh replays | repro.sh |
|---|---|---|---|---|
| 1 | reorder(#0,#4) | 2 | 10 / 10 | 10 / 10 |
| 2 | reorder(#5,#7) | 2 | 10 / 10 | 10 / 10 |
| 3 | reorder(#5,#9) | 3 | 10 / 10 | 10 / 10 |
| 4 | reorder(#6,#8) | 4 | 10 / 10 | 10 / 10 |
| 5 | reorder(#10,#11) | 3 | 10 / 10 | 10 / 10 |
| 6 | reorder(#12,#13) | 3 | 10 / 10 | 10 / 10 |
| 7 | reorder(#14,#15) | 3 | 10 / 10 | 10 / 10 |
| 8 | reorder(#1,#4) | 1 | 10 / 10 | 10 / 10 |
| 9 | reorder(#2,#4) | 1 | 10 / 10 | 10 / 10 |
| 10 | reorder(#3,#4) | 1 | 10 / 10 | 10 / 10 |
| 11 | reorder(#5,#8) | 2 | 10 / 10 | 10 / 10 |
| 12 | reorder(#6,#7) | 1 | 10 / 10 | 10 / 10 |
| 13 | reorder(#6,#9) | 4 | 10 / 10 | 10 / 10 |
| 14 | reorder(#8,#9) | 4 | 10 / 10 | 10 / 10 |
| 15 | reorder(#7,#9) | 4 | 10 / 10 | 10 / 10 |

## Deviations from the protocol

- **Source text seen before the push.** A Bun error at startup (the Hyperdrive binding was undefined) printed a stack trace quoting `apps/api/lib/db.ts` lines 13 to 18 and naming `apps/api/dev.ts:66`. These lines are not the webhook handler and did not shape the config, but they are application source.
- **Files read that the list does not name.** `apps/api/wrangler.jsonc` (read as build config) and `db/schema/id.ts` (in the schema directory, the id generator).
- **The handler is a dependency.** The root causes are in `@better-auth/stripe`, which the app installs, not in code the repository wrote. A claim of root causes "in an app" should say so.
- **Black-box trials before the push.** The password requirement, the `/health` path and the missing `@repo/email` build were found by starting the server and reading its errors. The network check delivered the workload once and read the `subscription` table. None of this read handler source.
- **The `start` command wraps the URL.** `sh -c` adds a password to `{DATABASE_URL}`. It is part of `stutter.json`, the allowed config.
- **`ignore: ["subscription.id"]`.** Added after the first hunt stopped at the stability check. That hunt ran no faults.
- **Two failed recording attempts.** `--sdk` given as a relative path, then as `./`, failed to load before any delivery, because `require` resolves against `lifecycles.mjs`. The third attempt used the absolute path. Each attempt was preceded by a passing hash check ([record.log](../../targets/react-starter-kit/record.log)).
- **`npx --offline`.** Start and migrate use `npx --offline -y bun@1.4.2`, the same Bun from the same npm source, without a registry lookup per start. During the first manual start, without `--offline`, the npx process held a connection to `104.16.7.34:443` (the npm registry). That was the launcher, not the app.
- **Network check limits.** It was one delivery of the workload, outside stutter. workerd was covered only by `lsof` polling every 0.2 s, which can miss a very short connection.
- **Leftover servers.** Two servers started at 18:35, during the recording and the hunt that stopped at the stability check, were still running after the hunt (stutter exited without stopping them). Each had its own database (`r37297_1`, `r37454_1`) and port. They were killed after the hunt.
- **`-lanes 2`**, because other hunters shared the cluster.

## Upstream

Checked 2026-11-15 against better-auth `main` at `8382ce6`. Cause 1 was already reported as better-auth/better-auth#12255 (fix in #12267). Cause 3's `TypeError` was replaced by a logged warning in #12473 (shipped in 1.7.25); the event is still acknowledged and dropped, which #10534 and #11943 cover. Causes 4 and 2 had no report and are filed as [better-auth/better-auth#12642](https://github.com/better-auth/better-auth/issues/12642), with a standalone reproduction on the in-memory adapter in [targets/better-auth-repro](../../targets/better-auth-repro).
