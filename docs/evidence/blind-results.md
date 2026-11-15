# Blind results (Lane 2)

This reports the blind hunts under [blind-protocol.md](blind-protocol.md), frozen in `c7ee40d`, using its "Reporting" and "The claim" sections. Every number here comes from a file in this repository. Each target's own doc has the details: [talivia](../targets/talivia.md), [paykit](../targets/paykit.md), [react-starter-kit](../targets/react-starter-kit.md).

## Order of events

All times are from `git log` and from the local reflog of `origin/main`, in UTC on 2026-11-14 unless marked.

| Time | Commit | Event |
|---|---|---|
| 23:13:37 | `c7ee40d` | Workload hash and both protocols frozen. Pushed 23:13:39. |
| 23:21:24 | `dad6d52` | sync-engine #177 case study. Not blind (see the end of this file). |
| 23:26:13 | `23d3ffd` | Blind screen: 29 judged, 3 pass. Reasons held back under a published sha256. Pushed 23:26:14. |
| 23:42:03 | `e6b9d67` | talivia results committed. Reached `origin` 23:44:54, with `e4ebb06`, after one refused push. |
| 23:54:40 | `937067a` | paykit results. Pushed 23:54:47. |
| 23:55:58 | `1d72ab0` | Lane 1 results ([lane1-results.md](lane1-results.md)). |
| 00:05:45 (11-15) | `1c2d81d` | react-starter-kit results. Pushed 00:05:54. |
| 00:09:35 (11-15) | `ee49c13` | Screener's reasons published, after all three hunt pushes. |

The reasons file, [blind-reasons.md](blind-reasons.md), hashes to `fce0217c659c8c9bafa536a43ba4dd4b1ad505916421cfbea87dc83df5243264`. That is the hash [blind-screen.md](blind-screen.md) published before any hunt.

## Counts

| | Count | Repositories |
|---|---|---|
| Judged | 29 | Ranks 1 to 31 of [blind-search.json](blind-search.json), less the 2 excluded by the protocol (ranks 3 and 29) |
| Passed, so hunted (K) | **3** | talivia-group/talivia, getpaykit/paykit, kriasoft/react-starter-kit |
| Found defects | 2 | getpaykit/paykit, kriasoft/react-starter-kit |
| Clean | 1 | talivia-group/talivia |
| Not reached | 0 | Every hunt fired all 23 of its reorder pairs. |
| Not run | 0 | |
| Screened out: the workload cannot reach it | 3 | sadmann7/skateshop, nextacular/nextacular, moinulmoin/chadnext |
| Failed the criteria | 23 | |
| Distinct root causes (M) | **7** | paykit 3, react-starter-kit 4. 6 if react-starter-kit's causes 2 and 3 are counted as one (see below). |

The list ran out with 3 passes. The protocol seeks 4, so K is 3.

## Per target

| | talivia | paykit | react-starter-kit |
|---|---|---|---|
| Pinned SHA | `9b7ad2a6039bce79e20364dbcea44c66a765f398` | `c7572ff7e0116428a43a2f98ad25dc2c26c8927e` | `dbdb0d4991bbaae1671ff62180089f2733dd406a` |
| Results commit | `e6b9d67`, 23:42:03 | `937067a`, 23:54:40 | `1c2d81d`, 00:05:45 (11-15) |
| On `origin` | 23:44:54 | 23:54:47 | 00:05:54 (11-15) |
| Single-fault schedules | 109 | 87 | 71 |
| Failing schedules | 0 | 29 | 15 |
| Distinct end-state changes | 0 | 26 | 12 |
| Effects | 0 | 5 | 3 |
| M | 0 | 3 | 4 |
| Hunt time, lanes | 93 s, 2 | 547 s, 2 | 1125 s, 2 |
| Results | [results.json](../../targets/talivia/hunt/results.json) | [results.json](../../targets/paykit/hunt/results.json) | [results.json](../../targets/react-starter-kit/hunt/results.json) |

Fault tables. Each cell is schedules that changed the final state out of schedules that fired. No schedule failed to fire in any hunt.

| Fault | talivia | paykit | react-starter-kit |
|---|---|---|---|
| lost (redelivery) | 0 / 16 | 0 / 16 | 0 / 16 |
| concurrent redelivery | 0 / 16 | 0 / 16 | 0 / 16 |
| reorder | 0 / 23 (0 no-op pairs) | 13 / 23 (0 no-op pairs) | 15 / 23 (3 no-op pairs) |
| crash after commit | 0 / 54 | 16 / 32 | 0 / 16 |

Every finding in both failing hunts is 1-minimal and reproduced in 10 of 10 fresh replays. Every reorder finding is also reproduced by its curl script in 10 of 10 trials. paykit's 16 crash findings cannot be scripted with curl.

### talivia-group/talivia: clean, reached

0 of 109 single faults changed the final state. All 23 reorder pairs fired, and none is a no-op pair, so all 23 are real ordering tests. M = 0.

Read after the push, the handler shows why: an event log with a unique key on the Stripe event id, and a subscription write that only applies an event newer than the stored `last_event_at` ([talivia.md](../targets/talivia.md)).

### getpaykit/paykit: M = 3

Failing schedules: 13 reorders and 16 crashes, 29 in all.

Effects:

1. A canceled subscription comes back (reorder).
2. A pending cancellation is lost (reorder).
3. A subscription goes back to trialing (reorder).
4. A customer's new subscription is ended by an event about the old one (reorder).
5. The event is dropped and never applied (crash after the first commit, all 16 events).

Root causes:

1. Subscription events are applied whatever their age. 9 failing pairs (findings 17 to 22, 25, 28, 29).
2. Competing subscriptions are ended in arrival order, without asking which is newer. 4 failing pairs (findings 23, 24, 26, 27).
3. The claim on an event is committed before the work, and a stuck claim blocks redelivery. 16 crash schedules (findings 1 to 16).

**Cause 3 depends on timing.** After a crash, the event row stays `processing`. A redelivery within 5 minutes of the claim is skipped as a duplicate and answered 200, so Stripe stops retrying and the event is never applied. A redelivery after 5 minutes reclaims the row and applies the event. stutter retries at once, inside the window. Stripe [documents](https://docs.stripe.com/webhooks) automatic retries "for up to three days with an exponential back off in live mode", but does not publish the delay before the first retry. So this hunt does not show whether Stripe's own first retry lands inside the 5 minutes. A manual resend inside the window, from the Dashboard or the CLI, does trigger it.

### kriasoft/react-starter-kit: M = 4

Failing schedules: 15 reorders. Lost, concurrent and crash faults: 0 of 48.

**The causes are in a dependency.** The app writes no webhook code of its own. It mounts the `@better-auth/stripe` plugin, version 1.7.4, as its lockfile installs it, and all four causes are in that package's `dist/index.mjs`.

Effects:

1. A canceled subscription comes back.
2. An event is lost and an older state stays.
3. One subscription's row is taken over by another.

Root causes:

1. `updated` is applied unconditionally, so a late `updated` undoes a later `deleted`. 4 failing pairs (findings 8, 9, 10, 12).
2. `deleted` for a row that does not exist yet is logged and dropped with 200. 3 failing pairs (findings 1, 2, 11).
3. `updated` for a customer with no row throws a `TypeError`, which is caught, logged and answered 200. 4 failing pairs (findings 3, 5, 6, 7).
4. The customer-id fallback picks another subscription's row and rewrites it, including its Stripe subscription id. 4 failing pairs (findings 4, 13, 14, 15).

**Counting causes 2 and 3.** Both are an early event for an unknown subscription, dropped with a 200. They sit in different handlers and fail in different ways: one is a deliberate warning, the other an unhandled `undefined`. Under the protocol's definition, "one defect in the code that needs its own fix", each needs its own change, so M = 4. Counted as one, M would be 3, and the total would be 6.

### Follow-up, not blind: `@better-auth/stripe` 1.7.25

This was run after the blind result was pushed and the plugin's source had been read, so it is not blind and is not counted above. The question: do the four causes still reproduce on the latest release, 1.7.25?

- **Setup** ([ba1725-setup.log](../../targets/react-starter-kit/ba1725-setup.log)). A copy of the target at `.targets/react-starter-kit-ba1725`. The original is untouched. Bumping only the plugin does not load: 1.7.25 imports `HIDE_METADATA` from `@better-auth/core/api`, which core 1.7.4 does not export. So `apps/api` pins `@better-auth/stripe` and its required peer `better-auth` at 1.7.25. That moves `@better-auth/core` and the adapters to 1.7.25 and `better-call` to 1.4.5. The app's code, Bun, `stripe`, migrations, fixtures and trace are unchanged.
- **Run.** Config [stutter-ba1725.json](../../targets/react-starter-kit/stutter-ba1725.json), identical to `stutter.json` except the name and the directory. The existing [trace.jsonl](../../targets/react-starter-kit/trace.jsonl), replayed. Hash check, then `-lanes 4` ([hunt-ba1725.log](../../targets/react-starter-kit/hunt-ba1725.log), [results.json](../../targets/react-starter-kit/hunt-ba1725/results.json)).
- **Result.** The same 71 schedules, the same 15 failing reorders and the same 3 no-op pairs. Lost, concurrent and crash: 0 of 48. Each finding has the same pair, the same minimal schedule and the same end-state diff as on 1.7.4. All 15 are 1-minimal, reproduced in 10 of 10 replays and by `repro.sh` in 10 of 10. The hunt exits 1 in 913 s.

All four causes still reproduce on 1.7.25:

1. `updated` is still applied with no ordering check. Findings 8, 9, 10, 12 unchanged.
2. `deleted` for an unknown subscription is still logged and dropped with 200. Findings 1, 2, 11 unchanged. Finding 1 logs "Subscription not found for subscriptionId: sub_one_a".
3. `updated` for a customer with no row no longer throws. It logs "No subscription found for customerId" and returns, and the route still answers 200. The `TypeError` is gone, but the event is still dropped. Findings 3, 5, 6, 7 unchanged. Finding 5 logs the new warning for `cus_team3`.
4. The customer-id fallback is unchanged. Findings 4, 13, 14, 15 unchanged.

On 1.7.25, causes 2 and 3 now fail the same way, a logged warning and a 200. They are still in two handlers, and each still needs its own change.

## Protocol deviations

Each comes from that target's doc, where it is explained in full.

**The screen.**

- The list was exhausted with 3 passes, fewer than the 4 the protocol seeks. No refill was possible.

**talivia** ([talivia.md](../targets/talivia.md#deviations-from-the-protocol)).

- The first push of `e6b9d67` was refused by the run's permission system, not by the remote. The hunter did not force. A retry went through: `e6b9d67` reached `origin` at 23:44:54 together with `e4ebb06`, a doc written from allowed materials only. The handler was read after that.
- The doc says the protocol asks for a talivia section in `scripts/fetch-targets.sh`. The protocol does not. The section was left out at hunt time and is added now.
- The README says `APP_SECRET` encrypts saved provider credentials. The fixture's plaintext webhook secret was accepted anyway.
- Two application-minted UUID keys were added to `ignore` after the stability check named them.

**paykit** ([paykit.md](../targets/paykit.md#deviations-from-the-protocol)).

- `scripts/load-root-env.js` was read before the push. The demo's drizzle and next configs import it. It is a build helper that the read list does not name.
- The fixture holds subscription rows, not only customers and products. Without them the 16 events changed no subscription state.
- The "schedule the default plan" and "activate scheduled plan" branches were not reached, because the fixture's products are not plans in the app's catalog.
- The workload hash was checked before the first hunt run, which stopped at the stability check, but not again before the second. It was checked after the second and still matched.
- Two random ids were added to `ignore` after the first hunt run named them.
- Three direct deliveries to scratch databases, to choose fixtures and to check the network, before recording.
- `--sdk` was given as an absolute path, because a relative one did not resolve.
- The header list of `fetch-targets.sh` was not updated at hunt time. It is updated now.
- `--api basil` was applied as the rule says, though `stripe@19.3.1` pins `2025-10-29.clover`.

**react-starter-kit** ([react-starter-kit.md](../targets/react-starter-kit.md#deviations-from-the-protocol)).

- A Bun error at startup printed a stack trace that quoted `apps/api/lib/db.ts` lines 13 to 18 and named `apps/api/dev.ts:66`. These are application source, though not the handler.
- `apps/api/wrangler.jsonc` and `db/schema/id.ts` were read before the push. The read list does not name them.
- The handler is a dependency, `@better-auth/stripe`, not code the repository wrote.
- Black-box trials before the push: the password requirement, the `/health` path and the missing `@repo/email` build were found by starting the server and reading its errors.
- `start` is a `sh -c` wrapper that adds a password to the database URL.
- `subscription.id` was added to `ignore` after the first hunt stopped at the stability check.
- Two recording attempts failed to load the SDK before any delivery. Each was preceded by a passing hash check.
- Bun is started with `npx --offline`. During one manual start without it, npx held a connection to the npm registry.
- The network check covered workerd only by `lsof` polling every 0.2 s.
- Two servers from earlier runs were left running and were killed after the hunt.

**All three hunts** ran with `-lanes 2`, because other hunters shared the cluster.

## The claim

The protocol gives two forms. The first is for hunts that find defects. The second, with the outcome counts, is for hunts that find "fewer or none". Here two of three apps had defects and one was clean, so the second form is used: it is the one that reports the clean app. The first form also says "webhook-ordering root causes", and one of the 7, paykit's cause 3, was found by a crash fault, not a reorder.

> Against a fixed Stripe webhook workload (sha256 bc121c1f), used unmodified, stutter ran on 3 open-source apps screened by a fixed, published repo search and pass/fail criteria. It found 7 distinct root causes. 1 app was clean, 0 were not reached and 0 were not run.

Two qualifications go with it:

- The 4 react-starter-kit causes are in its dependency `@better-auth/stripe` 1.7.4, not in code the app wrote.
- If react-starter-kit's causes 2 and 3 are counted as one, M is 6. In the first form, restricted to ordering, M is 6, or 5 with that merge.

## Not blind: the sync-engine case study

[stripe/sync-engine #177](../targets/sync-engine.md) was found by searching merged fixes before the freeze, so the bug and its fix were known. It is not counted above. On the same workload, the commit before the fix (`96e6b39`) fails 13 of 23 reorders, and the fix (`d486b41`) is clean on all 87 single faults ([hunt-pre](../../targets/sync-engine/hunt-pre/results.json), [hunt-fix](../../targets/sync-engine/hunt-fix/results.json)). [Lane 1](lane1-results.md) screened 5,303 merged PRs and found it the only past fix that is both in model and runnable here.
