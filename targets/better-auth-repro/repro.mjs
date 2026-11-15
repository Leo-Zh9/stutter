// Standalone reproduction for @better-auth/stripe: webhooks delivered out of order.
// Each scenario runs twice on a fresh in-memory database, once in Stripe's order and once
// with two events swapped, and prints the subscription rows. Nothing calls Stripe's API.
import { stripe } from "@better-auth/stripe";
import { betterAuth } from "better-auth";
import { memoryAdapter } from "better-auth/adapters/memory";
import Stripe from "stripe";

const secret = "whsec_test";
const stripeClient = new Stripe("sk_test_unused");
const t0 = 1790813160;

function subscription(id, price, status) {
  return {
    object: "subscription",
    id,
    customer: "cus_1",
    status,
    metadata: {},
    cancel_at: null,
    cancel_at_period_end: false,
    canceled_at: status === "canceled" ? t0 + 120 : null,
    ended_at: status === "canceled" ? t0 + 120 : null,
    trial_start: null,
    trial_end: null,
    schedule: null,
    items: {
      data: [
        {
          price: { id: price, recurring: { interval: "month" } },
          quantity: 1,
          current_period_start: t0,
          current_period_end: t0 + 30 * 86400,
        },
      ],
    },
  };
}

function event(id, type, created, object) {
  return { id, object: "event", api_version: "2025-10-29.clover", type, created, data: { object } };
}

function app() {
  const db = {
    user: [
      {
        id: "user_1",
        name: "User 1",
        email: "user1@example.com",
        emailVerified: true,
        stripeCustomerId: "cus_1",
        createdAt: new Date(t0 * 1000),
        updatedAt: new Date(t0 * 1000),
      },
    ],
    session: [],
    account: [],
    verification: [],
    subscription: [],
  };
  const auth = betterAuth({
    baseURL: "http://localhost:3000",
    secret: "a-test-secret-that-is-long-enough-123",
    database: memoryAdapter(db),
    logger: { disabled: true },
    plugins: [
      stripe({
        stripeClient,
        stripeWebhookSecret: secret,
        subscription: {
          enabled: true,
          plans: [
            { name: "starter", priceId: "price_starter" },
            { name: "pro", priceId: "price_pro" },
          ],
        },
      }),
    ],
  });
  return { db, auth };
}

async function deliver({ auth }, e) {
  const payload = JSON.stringify(e);
  const signature = stripeClient.webhooks.generateTestHeaderString({ payload, secret });
  const res = await auth.handler(
    new Request("http://localhost:3000/api/auth/stripe/webhook", {
      method: "POST",
      headers: { "content-type": "application/json", "stripe-signature": signature },
      body: payload,
    }),
  );
  return res.status;
}

function rows(db) {
  return db.subscription
    .map((r) => `${r.stripeSubscriptionId} plan=${r.plan} status=${r.status}`)
    .sort();
}

async function run(events) {
  const a = app();
  const log = [];
  for (const e of events) {
    const status = await deliver(a, e);
    log.push(`${e.id} ${e.type} ${e.data.object.id} -> HTTP ${status}`);
  }
  return { log, rows: rows(a.db) };
}

const scenarios = [
  {
    name: "1. An early updated for a new subscription takes over the customer's other row",
    story:
      "The customer has a Starter subscription and adds a Pro one. The Pro subscription's updated event arrives before its created event.",
    events: [
      event("evt_1", "customer.subscription.created", t0, subscription("sub_a", "price_starter", "active")),
      event("evt_2", "customer.subscription.created", t0 + 180, subscription("sub_b", "price_pro", "active")),
      event("evt_3", "customer.subscription.updated", t0 + 240, subscription("sub_b", "price_pro", "active")),
    ],
    swap: [1, 2],
  },
  {
    name: "2. A deleted that arrives before its created leaves a live subscription",
    story: "The subscription is created, then canceled. The deleted event arrives first.",
    events: [
      event("evt_1", "customer.subscription.created", t0, subscription("sub_a", "price_starter", "active")),
      event("evt_2", "customer.subscription.deleted", t0 + 120, subscription("sub_a", "price_starter", "canceled")),
    ],
    swap: [0, 1],
  },
];

let failed = 0;
for (const s of scenarios) {
  const swapped = [...s.events];
  const [i, j] = s.swap;
  [swapped[i], swapped[j]] = [swapped[j], swapped[i]];
  const inOrder = await run(s.events);
  const reordered = await run(swapped);
  const same = JSON.stringify(inOrder.rows) === JSON.stringify(reordered.rows);
  if (!same) failed++;
  console.log(`\n${s.name}\n${s.story}\n`);
  console.log("In order:");
  for (const l of inOrder.log) console.log(`  ${l}`);
  for (const r of inOrder.rows) console.log(`  row: ${r}`);
  console.log(`Swapped ${s.events[i].id} and ${s.events[j].id}:`);
  for (const l of reordered.log) console.log(`  ${l}`);
  for (const r of reordered.rows) console.log(`  row: ${r}`);
  console.log(same ? "OK: same rows" : "BUG: delivery order changed the saved rows");
}
process.exit(failed ? 1 : 0);
