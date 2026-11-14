-- Labelled fixture: five teams that each already have a Stripe customer
-- (billingId), the state the kit reaches after checkout creates it.
INSERT INTO "Team" (id, name, slug, "billingId", "billingProvider") VALUES
  ('team_1', 'Team 1', 'team-1', 'cus_team1', 'stripe'),
  ('team_2', 'Team 2', 'team-2', 'cus_team2', 'stripe'),
  ('team_3', 'Team 3', 'team-3', 'cus_team3', 'stripe'),
  ('team_4', 'Team 4', 'team-4', 'cus_team4', 'stripe'),
  ('team_5', 'Team 5', 'team-5', 'cus_team5', 'stripe');
