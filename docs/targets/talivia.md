# talivia-group/talivia, blind result: clean

| | |
|---|---|
| Repo | [talivia-group/talivia](https://github.com/talivia-group/talivia), default branch `main` |
| Pinned SHA | `9b7ad2a6039bce79e20364dbcea44c66a765f398` ("fix(app): improve mobile dashboard layout") |
| Stack | Next.js 16.2.6 (App Router, webpack), Prisma 7.6.0 with `@prisma/adapter-pg`, node-postgres 8.20, `stripe@22.3.0`, Node 24.21 |
| Config | [stutter.json](../../targets/talivia/stutter.json) |
| Results | [hunt/REPORT.md](../../targets/talivia/hunt/REPORT.md), [hunt/results.json](../../targets/talivia/hunt/results.json) |
| Outcome | Clean. 109 single-fault schedules fired, 0 changed the final state. The reorder lane is reached (23 pairs fired, 0 no-op). |

## This is a blind result

talivia is the open-source edition of a revenue-attribution analytics product (a Umami
derivative). It was screened by the frozen repo search and pass/fail criteria, and handed to
the hunter as a repository name and a pinned SHA only. The hunter saw no screener reasons.

## Blind protocol

- **Pinned SHA.** `9b7ad2a6039bce79e20364dbcea44c66a765f398`, used and no other.
- **What was read before the results were committed.** Only the allowed materials:
  `README.md`, `.env.example`, `docker-compose.yml`, `package.json`, `pnpm-lock.yaml`,
  `prisma.config.ts`, `prisma/schema.prisma`, the three `prisma/migrations/*/migration.sql`
  files, and a tree listing of `src/`. The webhook route path
  `/api/payments/stripe/[websiteId]/webhook` was taken from the directory listing. No
  application source was read: not the Stripe handler, not any library it calls, no `scripts/`
  logic. No code search was run.
- **Generator flags.**
  - `--api basil`: `package.json` pins `stripe` at `^22.3.0` (>= 18), so basil.
  - `--sdk /Users/leozhang/stutter/.targets/talivia/node_modules/stripe`: the copy the app
    installs, `stripe@22.3.0`.
  - `--customers`: left at the generator default `cus_team1..cus_team5`, which the fixture holds.
- **Push before any source reading.** The config, trace and results were committed as
  `e6b9d6727640510c961266e84c86b27edf902fd5` (local). The push of that commit to `origin/main`
  was refused by the run's permission system (classified as unsanctioned communication), so at
  the time this doc was written the commit had not reached the remote. No handler source has
  been read. See the deviations.

## Setup

- **Fetch and build.** The repo is cloned into `.targets/talivia` at the pinned SHA.
  `pnpm@10.10.0 install --frozen-lockfile` installs; `prisma generate` then an
  `esbuild` bundle build the Prisma client; `next build --webpack --experimental-build-mode=compile`
  compiles the app without prerendering any page that would call a live service.
- **Start.** `next start -p {PORT} -H 127.0.0.1`, with `DATABASE_URL`, `APP_SECRET` and
  `NODE_ENV=production`. Health is `/api/heartbeat`. The Stripe webhook route is
  `/api/payments/stripe/{websiteId}/webhook`; the fixture's website id
  `11111111-1111-4111-8111-111111111111` is the one in the URL.
- **Offline.** `.env.example` and the README config table list only three settings:
  `DATABASE_URL`, `APP_SECRET` and an optional `COINGECKO_API_KEY` (crypto rate conversion,
  left unset). No ClickHouse, Redis or Kafka: the OSS edition and its `docker-compose.yml` run
  on Postgres alone, and those clients in `package.json` are optional cloud features.
  `CHECKPOINT_DISABLE=1` and `PRISMA_HIDE_UPDATE_MESSAGE=1` turn off Prisma's CLI version
  telemetry.
- **Network check.** During recording a preload (`--require netlog.cjs`) logged every TCP
  connect and DNS lookup of each node process. The migrate CLI, the server and the generator
  connected only to `127.0.0.1:54329`, the project-local Postgres
  ([netlog-record.txt](../../targets/talivia/netlog-record.txt)). The subscription webhook path
  makes no Stripe or other outside call. One earlier run showed a single `checkpoint.prisma.io`
  connect from the Prisma migrate CLI (version telemetry); `CHECKPOINT_DISABLE=1` removes it,
  and it was never on the webhook path.
- **Database.** The template is built by `prisma migrate deploy` (the three OSS migrations).
  `relationMode = "prisma"`, so the schema declares no database-level foreign keys.
- **Fixtures.** [fixtures.sql](../../targets/talivia/fixtures.sql), derived from the schema: one
  `website` row whose id is in the webhook URL, and one Stripe `payment_provider_connection`
  (`provider_name='stripe'`, `connection_status='connected'`, `webhook_status='active'`,
  `webhook_secret_ref='whsec_stutter_local'`). The baseline migration already seeds the `admin`
  user. The README says `APP_SECRET` "encrypts saved provider credentials", so the webhook
  secret was expected to be stored encrypted; in fact the plaintext value in `webhook_secret_ref`
  was accepted and all 16 signed deliveries verified and returned HTTP 200.
- **Events.** The frozen workload, [lifecycles.mjs](../../targets/stripe/lifecycles.mjs),
  unchanged: 16 `customer.subscription.*` events across five lifecycles, 60 s apart. Its sha256
  was checked before recording and again before the hunt and matched `bc121c1f…cdb6`
  ([record.log](../../targets/talivia/record.log),
  [workload-check.txt](../../targets/talivia/workload-check.txt)).
- **Oracle ignore list.** The stability check reported two columns that differ between two
  fault-free replays: `provider_event.provider_event_id` and `subscription.subscription_id`.
  Both are random UUID primary keys minted in application code (Prisma `@default(uuid())`), not
  by Postgres, so the catalog cannot map them. They are surrogate keys; the meaningful identity
  is the natural key still compared on every row (`website_id, provider_name,
  provider_subscription_id` for subscriptions; `provider_name, provider_event_key` for events).
  They were added to `ignore` in `stutter.json`, which is the documented false-alarm mechanism.
  Every other column, including `status`, `lifecycle_status`, `canceled_at`, `ended_at`,
  `product_name` and `last_event_type`, is compared.
- **Recording.** One trace, 16 deliveries, all HTTP 200
  `{"ok":true,"status":"processed",...}`. The hunt ran with `-lanes 2` because other hunters
  shared the cluster.

## Results

| Fault | Schedules | Fired | Not fired | Changed the final state |
|---|---|---|---|---|
| lost (redelivery) | 16 | 16 | 0 | 0 |
| concurrent redelivery | 16 | 16 | 0 | 0 |
| reorder | 23 | 23 | 0 | 0 (0 no-op pairs) |
| crash after commit | 54 | 54 | 0 | 0 |

109 single-fault schedules, all fired, 0 failing. The hunt exits 0 in 93 s. **M = 0 distinct
root causes. 0 failing pairs, 0 effects, 0 reproductions.**

The reorder lane is reached, not skipped: 23 reorder pairs fired and 0 were no-op pairs, so
every one is a real ordering test and all passed. This is a stronger reached-and-clean result
than a run whose passing reorders are mostly no-ops.

## Why it is clean (inferred from the schema, not from the handler)

No handler source was read. From the allowed materials the clean result is consistent with two
guards the schema already carries:

- **Redelivery, concurrency and crash are idempotent.** `provider_event` has a unique key on
  `(provider_name, provider_event_key)` ([schema.prisma:797](../../.targets/talivia/prisma/schema.prisma)),
  an event-dedupe log, and `subscription` has a unique key on `(website_id, provider_name,
  provider_subscription_id)` ([schema.prisma:766](../../.targets/talivia/prisma/schema.prisma)).
  A replayed or duplicated delivery writes the same natural-keyed rows, so lost, concurrent and
  crash-after-commit all leave the final state unchanged.
- **Out-of-order delivery is guarded.** `subscription` carries `last_event_at` (not null) and
  `last_event_priority` ([schema.prisma:754-755](../../.targets/talivia/prisma/schema.prisma)).
  A handler that refuses to apply an event older than the stored `last_event_at` would make the
  final state independent of delivery order, which is exactly what stutter observed across all
  23 reorder pairs.

A manual in-order delivery to a scratch database confirmed the handler writes correct, distinct
final states (sub_one_a canceled, sub_two_a canceled, sub_two_b active, sub_three_a active after
the upgrade, sub_four_a unpaid, sub_five_a active) and 16 `provider_event` rows, so the clean
reorder result is genuine ordering handling, not a handler that drops writes.

Confirming these guards in the handler source is the step that follows the push; it was not done,
because the push is blocked (see deviations).

## Deviations from the protocol

- **The push is blocked.** `git push origin main` of commit
  `e6b9d6727640510c961266e84c86b27edf902fd5` was refused by the run's permission system, not by
  the remote. Per the protocol the hunter reports a rejected push and does not force. As a
  result the commit had not reached `origin` when this doc was written, and the handler source
  was not read (the protocol gates source reading on the push). This doc reports the clean
  result from allowed materials only.
- **fetch-targets.sh section not added.** The protocol asks for a `talivia` clone section in
  `scripts/fetch-targets.sh`. At hunt time that file had another hunter's uncommitted edits in
  the shared checkout, and editing it would have entangled their work in this hunter's commit
  ("touch only your own paths"). It was left untouched and should be added when the shared file
  is free.
- **Credential storage expectation.** The README says `APP_SECRET` encrypts saved provider
  credentials, so a working `webhook_secret_ref` was expected to be an encrypted blob. In
  practice the plaintext secret was accepted and every delivery verified, so the target was
  runnable within the allowed inputs after all.
- **Oracle ignore list.** Two application-minted UUID primary keys were added to `ignore`, as
  the stability check instructed. This is an allowed `stutter.json` input, not a workload or code
  change.
