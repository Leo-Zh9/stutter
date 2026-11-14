// Generates a labelled synthetic set of Stripe subscription lifecycles and
// delivers them, in creation order, to a webhook endpoint. Each delivery is
// signed with the target's own pinned Stripe SDK, so the recording run is
// accepted exactly as a real delivery would be.
//
// Shapes follow the Stripe API version the target pins:
//   basil  (2025-04-30.basil): current_period_* live on subscription items
//   acacia (2025-02-24.acacia): current_period_* live on the subscription
//
// usage: node lifecycles.mjs --sdk <path to node_modules/stripe> --url <endpoint>
//          --secret <whsec> --api basil|acacia [--customers cus_a,cus_b,...]
import { createRequire } from 'node:module';

const args = Object.fromEntries(process.argv.slice(2).reduce((acc, a, i, all) => {
  if (a.startsWith('--')) acc.push([a.slice(2), all[i + 1]]);
  return acc;
}, []));
const require = createRequire(import.meta.url);
const Stripe = require(args.sdk);
const stripe = new Stripe('sk_test_offline_never_used');
const api = args.api || 'basil';
const apiVersion = api === 'basil' ? '2025-04-30.basil' : '2025-02-24.acacia';
const customers = (args.customers || 'cus_team1,cus_team2,cus_team3,cus_team4,cus_team5').split(',');

const T0 = Date.UTC(2026, 9, 1) / 1000; // 2026-10-01, synthetic clock
let tick = 0;
let evtN = 0;
const MONTH = 30 * 24 * 3600;

const products = { basic: 'prod_basic', plus: 'prod_plus' };
function sub(customer, id, status, product, extra = {}) {
  const start = T0;
  const item = {
    id: `si_${id.slice(4)}`, object: 'subscription_item', created: start, metadata: {}, quantity: 1,
    subscription: id,
    plan: { id: `price_${product}`, object: 'plan', active: true, amount: product === products.plus ? 2400 : 800,
      currency: 'usd', interval: 'month', interval_count: 1, product, usage_type: 'licensed' },
    price: { id: `price_${product}`, object: 'price', active: true, currency: 'usd', product,
      recurring: { interval: 'month', interval_count: 1, usage_type: 'licensed' }, type: 'recurring',
      unit_amount: product === products.plus ? 2400 : 800 },
  };
  const s = {
    id, object: 'subscription', customer, status, livemode: false, metadata: {},
    cancel_at_period_end: false, canceled_at: null, ended_at: null, created: start, start_date: start,
    trial_start: null, trial_end: null, collection_method: 'charge_automatically', currency: 'usd',
    items: { object: 'list', data: [item], has_more: false, total_count: 1, url: `/v1/subscription_items?subscription=${id}` },
    ...extra,
  };
  if (api === 'basil') {
    item.current_period_start = start;
    item.current_period_end = start + MONTH;
  } else {
    s.current_period_start = start;
    s.current_period_end = start + MONTH;
    s.plan = item.plan;
  }
  return s;
}

const events = [];
function emit(type, object, previous) {
  tick += 60;
  events.push({
    id: `evt_${String(evtN++).padStart(4, '0')}`, object: 'event', api_version: apiVersion, created: T0 + tick,
    livemode: false, pending_webhooks: 1, request: { id: null, idempotency_key: null }, type,
    data: previous ? { object, previous_attributes: previous } : { object },
  });
}

// Lifecycle shapes, one per customer, all common in real billing:
const [c1, c2, c3, c4, c5] = customers;
// 1. trial, converts, a failed renewal recovers, then cancels
emit('customer.subscription.created', sub(c1, 'sub_one_a', 'trialing', products.basic));
emit('customer.subscription.updated', sub(c1, 'sub_one_a', 'active', products.basic), { status: 'trialing' });
emit('customer.subscription.updated', sub(c1, 'sub_one_a', 'past_due', products.basic), { status: 'active' });
emit('customer.subscription.updated', sub(c1, 'sub_one_a', 'active', products.basic), { status: 'past_due' });
emit('customer.subscription.deleted', sub(c1, 'sub_one_a', 'canceled', products.basic, { canceled_at: T0 + 9000, ended_at: T0 + 9000 }), null);
// 2. cancels, then resubscribes on a new subscription
emit('customer.subscription.created', sub(c2, 'sub_two_a', 'active', products.basic));
emit('customer.subscription.updated', sub(c2, 'sub_two_a', 'active', products.basic, { cancel_at_period_end: true }), { cancel_at_period_end: false });
emit('customer.subscription.deleted', sub(c2, 'sub_two_a', 'canceled', products.basic, { canceled_at: T0 + 9600, ended_at: T0 + 9600 }), null);
emit('customer.subscription.created', sub(c2, 'sub_two_b', 'active', products.plus));
emit('customer.subscription.updated', sub(c2, 'sub_two_b', 'active', products.plus), { items: {} });
// 3. upgrades basic -> plus
emit('customer.subscription.created', sub(c3, 'sub_three_a', 'active', products.basic));
emit('customer.subscription.updated', sub(c3, 'sub_three_a', 'active', products.plus), { items: {} });
// 4. trial ends unpaid
emit('customer.subscription.created', sub(c4, 'sub_four_a', 'trialing', products.basic));
emit('customer.subscription.updated', sub(c4, 'sub_four_a', 'unpaid', products.basic), { status: 'trialing' });
// 5. trial converts and stays active
emit('customer.subscription.created', sub(c5, 'sub_five_a', 'trialing', products.plus));
emit('customer.subscription.updated', sub(c5, 'sub_five_a', 'active', products.plus), { status: 'trialing' });

for (const ev of events) {
  const payload = JSON.stringify(ev);
  const header = stripe.webhooks.generateTestHeaderString({ payload, secret: args.secret });
  const res = await fetch(args.url, { method: 'POST', body: payload,
    headers: { 'Content-Type': 'application/json', 'Stripe-Signature': header } });
  if (res.status >= 300) {
    console.error(`${ev.id} ${ev.type}: HTTP ${res.status} ${await res.text()}`);
    process.exit(1);
  }
}
console.log(`delivered ${events.length} events (${apiVersion})`);
