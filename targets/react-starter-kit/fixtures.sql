-- Labelled fixture: five users that each already have a Stripe customer, the
-- state react-starter-kit reaches after sign-up, since its billing docs set
-- createCustomerOnSignUp: true. Columns follow db/migrations/0000_init.sql;
-- ids use the usr_ prefix of db/schema/id.ts.
INSERT INTO "user" (id, name, email, email_verified, stripe_customer_id) VALUES
  ('usr_team1', 'Team 1', 'team1@example.test', true, 'cus_team1'),
  ('usr_team2', 'Team 2', 'team2@example.test', true, 'cus_team2'),
  ('usr_team3', 'Team 3', 'team3@example.test', true, 'cus_team3'),
  ('usr_team4', 'Team 4', 'team4@example.test', true, 'cus_team4'),
  ('usr_team5', 'Team 5', 'team5@example.test', true, 'cus_team5');
