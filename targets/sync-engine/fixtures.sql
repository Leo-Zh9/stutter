-- Labelled fixture: the five customers, two products and two prices that the
-- generated subscription events reference. With backfill turned off nothing
-- fetches them from Stripe, and before the fix (migration 0034 drops them)
-- foreign keys require them: subscriptions.customer -> customers (0004),
-- subscription_items.price -> prices (0013), prices.product -> products (0003).
INSERT INTO stripe.customers (id, object, name, livemode) VALUES
  ('cus_team1', 'customer', 'Team 1', false), ('cus_team2', 'customer', 'Team 2', false),
  ('cus_team3', 'customer', 'Team 3', false), ('cus_team4', 'customer', 'Team 4', false),
  ('cus_team5', 'customer', 'Team 5', false);
INSERT INTO stripe.products (id, object, active, name, livemode) VALUES
  ('prod_basic', 'product', true, 'Basic', false),
  ('prod_plus', 'product', true, 'Plus', false);
INSERT INTO stripe.prices (id, object, active, currency, unit_amount, type, recurring, livemode, product) VALUES
  ('price_prod_basic', 'price', true, 'usd', 800, 'recurring', '{"interval":"month","interval_count":1,"usage_type":"licensed"}', false, 'prod_basic'),
  ('price_prod_plus', 'price', true, 'usd', 2400, 'recurring', '{"interval":"month","interval_count":1,"usage_type":"licensed"}', false, 'prod_plus');
