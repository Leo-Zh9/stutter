# Blind screen: reasons (screener only)

Screener: Agent B. Screening day: 2026-11-14 (UTC). Protocol: docs/evidence/blind-protocol.md as committed in c7ee40d.
Search: run verbatim from the stutter repo root, output docs/evidence/blind-search.json (31 repositories).
Each pinned SHA was read with `gh api repos/OWNER/REPO/commits/DEFAULT_BRANCH --jq .sha`, and every shallow clone was checked to have exactly that HEAD.
File:line references are at the pinned SHA.

## Rulings on ambiguous cases (applied literally)

- R1. **Category.** Criteria are tested in the protocol's order: P, a, b, c, d. The first one that fails is the recorded category. Later ones are noted where they are obvious, but they are not counted.
- R2. **Libraries and SDKs with no in-repo app** fail (P), because "PostgreSQL is its store" does not hold. Where a library ships demo apps, the demo that serves the Stripe route is the app judged. Wiring a different database into it would mean patching target code, which the protocol forbids.
- R3. **Runtimes and package managers.** (P) names the store, the extensions, Docker, Redis and "other service". It does not name language runtimes or package managers. So a missing runtime (Bun) is not a (P) failure under the text as written. The hunter needs that runtime installed, which is Leo's call, or else the target is reported as not run.
- R4. **Neon serverless drivers.** `@neondatabase/serverless` and `drizzle-orm/neon-*` speak WebSocket or HTTP to a Neon proxy, not the Postgres wire protocol. Reaching a local PostgreSQL through them needs another service (a Neon proxy), so they fail (P).
- R5. **Rank** is the 1-based position in blind-search.json. Ranks 3 (nextjs/saas-starter) and 29 (stripe/sync-engine) are excluded by the protocol and not judged. Every result had isFork=false, so no forks of the excluded targets appeared. stripe/ai ships copies of nextjs/saas-starter as benchmark environments, but it is not a fork, so it was judged as a whole (see rank 15).
- R6. **Stopping.** The list was exhausted with 3 passes, fewer than the 4 sought.
- R7. **(a) "verifies a Stripe signature"** is met by any HMAC check of the `Stripe-Signature` header, with or without stripe-node.

## Correction to my round-1 attack

In round 1 I said that next-saas-stripe-starter "HEAD has no webhook route". That was wrong. At c6df229 it serves the Stripe webhook through the Better Auth Stripe plugin: lib/auth/server.ts:4 and :131-133, mounted by app/api/auth/[...all]/route.ts. It still fails here, on (P), for an unrelated reason (rank 9).

## Entries

1. **better-auth/better-auth** @ 8382ce6: **FAIL (P)** (R2).
   - The only in-repo app with the Stripe plugin is demo/nextjs. Its store is SQLite, MySQL or libSQL: demo/nextjs/lib/auth.ts:44-84 (NodeSqliteDialect :46, MysqlDialect :62, LibsqlDialect :71), with `stripe(` at :249.
   - The PG demo (demo/expo/src/lib/auth.ts:3-6) has no Stripe plugin.
   - Note: the plugin's handlers would satisfy (b) to (d).
2. **kriasoft/react-starter-kit** @ dbdb0d4: **PASS**.
   - (P) PostgreSQL through Drizzle and Hyperdrive. Local connection strings are in .env:14-19. db/migrations has no `CREATE EXTENSION`. No Redis, KV, queue or secondaryStorage: apps/api/lib/auth.ts and wrangler.jsonc define only hyperdrive bindings (:28-30). Runtime is Bun (package.json:92, apps/api/package.json:13, Bun.argv in apps/api/dev.ts), see R3.
   - (a) The Better Auth Stripe plugin at apps/api/lib/auth.ts:121-125 verifies the signature in @better-auth/stripe@1.7.4 (locked in bun.lock), dist/index.mjs:1559-1560.
   - (b) The subscription handlers in dist/index.mjs:246-460 make only adapter calls, with no Stripe client calls. Plans come from config (apps/api/lib/auth.ts:126-142). The only callback, authorizeReference, is not on the webhook path.
   - (c) adapter.create and adapter.update on created, updated and deleted.
   - (d) Rows are found by `stripeSubscriptionId`, or by `stripeCustomerId` on user or organization, in dist/index.mjs:154-180 and :246-420. The metadata `subscriptionId` path is skipped when metadata is empty.
4. **vercel/next-forge** @ f189de7: **FAIL (P)** (R4).
   - packages/database/index.ts:3-13 uses PrismaNeon with a WebSocket.
   - Also noted: (c) fails. apps/api/app/webhooks/payments/route.ts:84-96 has no customer.subscription.* case and no database write.
5. **yournextstore/yournextstore** @ b0caa2d: **FAIL (P)**. No PostgreSQL store: data comes from the hosted commerce-kit API (lib/commerce.ts:1, package.json "commerce-kit"). There is no Stripe webhook route.
6. **sadmann7/skateshop** @ e954d54: **SCREENED OUT (c)**.
   - PostgreSQL through postgres-js and Drizzle.
   - src/app/api/webhooks/stripe/route.ts:36-231 handles checkout.session.completed, invoice.payment_succeeded, payment_intent.*, application_fee.created and charge.succeeded. customer.subscription.* falls through to `default` at :229-230, which only logs.
7. **stripe/stripe-node** @ 6237106: **FAIL (P)** (R2). It is the Stripe SDK (package.json "name": "stripe"), with no app and no store.
8. **async-labs/builderbook** @ a71c72e: **FAIL (P)**. MongoDB: book/9-end/package.json:33, :44, :45 (connect-mongo, mongodb, mongoose).
9. **mickasmt/next-saas-stripe-starter** @ c6df229: **FAIL (P)** (R4). lib/db/index.ts:1-2 and :16 use the Neon serverless Pool over WebSocket. (a) holds (see the correction above).
10. **nextify-limited/saasfly** @ 49f7e28: **FAIL (P)** (R4).
    - packages/db/index.ts:1 and :10 use `@vercel/postgres-kysely`, a Neon serverless driver.
    - Also noted: (b) fails (packages/stripe/src/webhooks.ts:11 and :42 call stripe.subscriptions.retrieve), and (c) fails (customer.subscription.updated only logs, :84-87).
11. **juicycleff/ultimate-backend** @ 0ae62a4: **FAIL (P)**. MongoDB, Redis, Consul and NATS: package.json:21, :39, :45, :56, :88, :92 and :113.
12. **adrianhajdin/ecommerce_sanity_stripe** @ c4d1c05: **FAIL (P)**. Sanity CMS is the store (package.json:13-18). There is no PostgreSQL and no webhook route (pages/api has only hello.js and stripe.js).
13. **talivia-group/talivia** @ 9b7ad2a: **PASS**.
    - (P) Prisma provider is postgresql (prisma/schema.prisma:8). No `CREATE EXTENSION` in prisma/migrations. ClickHouse, Kafka and Redis are optional, enabled only by env: src/lib/clickhouse.ts:22, src/lib/kafka.ts:14, src/lib/redis.ts:106.
    - (a) src/app/api/payments/stripe/[websiteId]/webhook/route.ts:762-785 verifies an HMAC Stripe signature (src/lib/stripe-webhook.ts), falling back to env STRIPE_WEBHOOK_SECRET at :775-776.
    - (b) route.ts:737-744 sends subscription events to processSubscriptionLifecycle (:427-481), and on to recordSubscriptionState (src/queries/prisma/payment.ts:443ff). Everything is Prisma, except currency conversion: src/queries/prisma/websiteCurrency.ts:62-82 converts only when currencies differ. The workload sends usd, DEFAULT_CURRENCY is USD (src/lib/constants.ts:23), and the same-currency branch is at src/lib/currency-exchange.ts:171-172. fetchWebsite uses Redis only if it is enabled (src/lib/load.ts:9-13). With empty metadata, findSessionContext returns null (payment.ts:103-104).
    - (c) providerEvent create and update (route.ts:813-836), and subscription create or updateMany (payment.ts ~:553-620).
    - (d) The website is found by websiteId in the URL path (fixture), and the subscription by `websiteId_providerName_providerSubscriptionId` from the payload's subscription.id (payment.ts ~:552). No metadata is needed.
14. **mrvautin/expressCart** @ b31302f: **FAIL (P)**. MongoDB: package.json:37, :59-60.
15. **stripe/ai** @ 74f7f39: **FAIL (P)** (R2, R5).
    - It is a toolkit, plugin and benchmark repo (README.md:5). It has no app of its own with a PostgreSQL store.
    - The only PG plus Stripe-webhook apps inside it are benchmarks/saas-starter-*/environment. These are copies of nextjs/saas-starter (that environment's README.md:1 reads "Next.js SaaS Starter"), an excluded, already-read target.
16. **stripe/react-stripe-js** @ d6670a2: **FAIL (P)** (R2). A React component library with no store.
17. **ParabolInc/parabol** @ 9edd33a: **FAIL (P)**. Redis is required: .env.example:116-117 sets REDIS_URL redis://localhost:6379. The webhook (packages/server/billing/stripeWebhookHandler.ts:1-5) runs through the uWS and GraphQL server stack.
18. **agentic-commerce-protocol/agentic-commerce-protocol** @ 7fdd78d: **FAIL (P)** (R2). A protocol spec repository (README.md:1) with no app and no store.
19. **reliverse/relivator** @ a1871b0: **FAIL (a)**. It has no Stripe code: a grep for "stripe" across the repo returns no files, and package.json:20-21 shows billing is Polar. Its store is PostgreSQL (package.json:45, :55).
20. **dev-xo/remix-saas** @ cdedef9: **FAIL (P)**. SQLite: prisma/schema.prisma:9, package.json:31 (adapter-libsql).
21. **adrianhajdin/ecommerce** @ 7c70b5c: **FAIL (P)**. MongoDB: package.json:26 (@payloadcms/db-mongodb).
22. **stripe/stripe-react-native** @ 13f0e79: **FAIL (P)** (R2). A mobile SDK with no store.
23. **nextacular/nextacular** @ 06bee75: **SCREENED OUT (c)**.
    - PostgreSQL through Prisma (prisma/schema.prisma:6), with no extensions.
    - src/pages/api/payments/hooks.ts:44-55 handles only charge.succeeded, found by metadata. Every other event type, including customer.subscription.*, gets a 400 and no write.
24. **adrianhajdin/ai_saas_app** @ f82bfdd: **FAIL (P)**. MongoDB: package.json:25-26.
25. **moinulmoin/chadnext** @ 1a1f25a: **SCREENED OUT (c)**.
    - PostgreSQL through Prisma (prisma/schema.prisma:8).
    - src/app/api/webhooks/stripe/route.ts:30-86 handles only checkout.session.* and invoice.payment_succeeded. customer.subscription.* gets a 200 with no write.
26. **RevenueCat/react-native-purchases** @ fbebad4: **FAIL (P)** (R2). A mobile SDK (package.json:2, :5).
27. **vue-stripe/vue-stripe** @ 9a8b075: **FAIL (P)** (R2). A Vue component library monorepo (package.json:2).
28. **ShenSeanChen/launch-mvp-stripe-nextjs-supabase** @ 513a8f0: **FAIL (P)**. The store is reached only through the Supabase client: app/api/stripe/webhook/route.ts:4, :39, :176 and :192 use supabaseAdmin (PostgREST over HTTP), and package.json:17-18 lists the Supabase packages. That needs the Supabase services.
30. **getpaykit/paykit** @ c7572ff: **PASS**.
    - (P) apps/demo is the in-repo app. Its PG pools are in apps/demo/src/server/db.ts:1, with AUTH_DATABASE_URL and PAYKIT_DATABASE_URL in apps/demo/src/env.js:13 and :15. No `CREATE EXTENSION` in packages/paykit/src/database/migrations or apps/demo. No Redis.
    - (a) apps/demo/src/app/paykit/[[...slug]]/route.ts mounts paykit.handler. Signature verification is at packages/paykit/src/stripe/stripe-provider.ts:1008-1020 (constructEventAsync).
    - (b) stripe-provider.ts:1021-1026 and :472-540 build subscription events straight from the payload. The comment at :481-483 says Stripe is deliberately not re-fetched. createCheckoutCompletedEvents returns [] for non-checkout events (:375-377). Actions go through webhook.service.ts:125-138 to subscription.service.ts:497ff, which are database calls only. emitCustomerUpdated calls only user `on` hooks (customer.service.ts ~:473-481), and the demo configures none (apps/demo/src/lib/paykit.ts:8-26).
    - (c) webhookEvent insert and update (webhook.service.ts:33, :54, :83), and subscription insert, end or sync (subscription.service.ts:497-620).
    - (d) findCustomerByProviderCustomerId on the payload's customer (subscription.service.ts:501-504), and getSubscriptionByProviderSubscriptionId (:517-520).
31. **Skolaczk/next-starter** @ a4d25f2: **FAIL (P)** (R4). src/lib/schema.ts:1-2 and :13-14 use neon-http.

## Counts

- Judged: 29 (31 minus 2 excluded).
- Passed: 3 (ranks 2, 13, 30).
- FAIL (P): 22.
- FAIL (a): 1.
- FAIL (b): 0.
- SCREENED OUT (c): 3.
- SCREENED OUT (d): 0.
