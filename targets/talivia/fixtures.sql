-- Best-effort fixture for a blind configuration attempt.
-- The baseline migration already seeds the admin user
-- (00000000-0000-4000-8000-000000000001). Derived from prisma/schema.prisma and
-- the migration DDL: a website whose id is the one in the Stripe webhook URL, and
-- a Stripe payment_provider_connection. The webhook signing secret is a saved
-- provider credential that the app encrypts with APP_SECRET (README: "APP_SECRET
-- ... encrypts saved provider credentials"); webhook_secret_ref holds that
-- encrypted value. This file can only put a plaintext guess there, which is why
-- this is an attempt, not a working config.
INSERT INTO "website" (website_id, name, user_id, created_by, updated_at)
VALUES ('11111111-1111-4111-8111-111111111111', 'Stutter Target',
        '00000000-0000-4000-8000-000000000001',
        '00000000-0000-4000-8000-000000000001', CURRENT_TIMESTAMP);

INSERT INTO "payment_provider_connection"
  (connection_id, website_id, provider_name, connection_status,
   provider_account_id, webhook_secret_ref, provider_webhook_endpoint_id,
   webhook_status, updated_at)
VALUES
  ('22222222-2222-4222-8222-222222222222',
   '11111111-1111-4111-8111-111111111111', 'stripe', 'connected',
   'acct_stutter_local', 'whsec_stutter_local', 'we_stutter_local',
   'active', CURRENT_TIMESTAMP);
