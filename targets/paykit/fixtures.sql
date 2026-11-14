-- Labelled fixture for getpaykit/paykit. Columns are derived from the drizzle
-- schema after migrations 0000-0003 (packages/paykit/src/database/schema.ts and
-- the migration SQL). It reproduces the state a PayKit app reaches after each
-- customer has checked out, which is what the subscription webhooks then sync:
--
--   * five customers, each mapped to a Stripe customer by stripe_customer_id
--     (the id the events carry in data.object.customer);
--   * the two products the events reference, mapped by stripe_price_id
--     (price_prod_basic / price_prod_plus) to a paykit_product row;
--   * one paykit_subscription row per Stripe subscription the workload touches,
--     keyed by stripe_subscription_id. PayKit's webhook path updates an existing
--     subscription row found by that id; it does not create subscription rows
--     from a webhook. Without these rows the events are logged but change no
--     subscription state. started status is "incomplete", the pre-activation
--     state; every run starts from this same template, so the initial value is
--     differenced out.
--
-- With these rows seeded nothing is fetched from Stripe.
INSERT INTO paykit_customer (id, email, name, stripe_customer_id, created_at, updated_at) VALUES
  ('user_team1', 'team1@example.com', 'Team 1', 'cus_team1', now(), now()),
  ('user_team2', 'team2@example.com', 'Team 2', 'cus_team2', now(), now()),
  ('user_team3', 'team3@example.com', 'Team 3', 'cus_team3', now(), now()),
  ('user_team4', 'team4@example.com', 'Team 4', 'cus_team4', now(), now()),
  ('user_team5', 'team5@example.com', 'Team 5', 'cus_team5', now(), now());
INSERT INTO paykit_product (internal_id, id, version, name, "group", is_default, price_amount, price_interval, price_currency, stripe_product_id, stripe_price_id, created_at, updated_at) VALUES
  ('prod_internal_basic', 'basic', 1, 'Basic', 'base', false, 800, 'month', 'usd', 'prod_basic', 'price_prod_basic', now(), now()),
  ('prod_internal_plus', 'plus', 1, 'Plus', 'base', false, 2400, 'month', 'usd', 'prod_plus', 'price_prod_plus', now(), now());
INSERT INTO paykit_subscription (id, customer_id, product_internal_id, stripe_subscription_id, status, canceled, cancel_at_period_end, quantity, created_at, updated_at) VALUES
  ('psub_one_a',   'user_team1', 'prod_internal_basic', 'sub_one_a',   'incomplete', false, false, 1, now(), now()),
  ('psub_two_a',   'user_team2', 'prod_internal_basic', 'sub_two_a',   'incomplete', false, false, 1, now(), now()),
  ('psub_two_b',   'user_team2', 'prod_internal_plus',  'sub_two_b',   'incomplete', false, false, 1, now(), now()),
  ('psub_three_a', 'user_team3', 'prod_internal_basic', 'sub_three_a', 'incomplete', false, false, 1, now(), now()),
  ('psub_four_a',  'user_team4', 'prod_internal_basic', 'sub_four_a',  'incomplete', false, false, 1, now(), now()),
  ('psub_five_a',  'user_team5', 'prod_internal_plus',  'sub_five_a',  'incomplete', false, false, 1, now(), now());
