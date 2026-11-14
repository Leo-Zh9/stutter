-- Labelled fixture: five teams that each already have a Stripe customer, the
-- state saas-starter reaches after checkout creates the customer.
INSERT INTO teams (name, stripe_customer_id) VALUES
  ('Team 1', 'cus_team1'), ('Team 2', 'cus_team2'), ('Team 3', 'cus_team3'),
  ('Team 4', 'cus_team4'), ('Team 5', 'cus_team5');
