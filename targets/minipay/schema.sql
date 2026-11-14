-- minipay schema. Ids are opaque strings minted by column defaults, the way
-- most payment APIs expose them (cus_..., ch_..., we_...).
CREATE TABLE customers (
  id            text PRIMARY KEY DEFAULT 'cus_' || replace(gen_random_uuid()::text, '-', ''),
  email         text NOT NULL,
  balance_cents bigint NOT NULL DEFAULT 0,
  created_at    timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE charges (
  id              text PRIMARY KEY DEFAULT 'ch_' || replace(gen_random_uuid()::text, '-', ''),
  customer_id     text NOT NULL REFERENCES customers(id),
  amount_cents    bigint NOT NULL,
  currency        text NOT NULL,
  status          text NOT NULL DEFAULT 'pending',
  status_at       timestamptz,
  refunded_cents  bigint NOT NULL DEFAULT 0,
  idempotency_key text,
  created_at      timestamptz NOT NULL DEFAULT now()
);

-- Idempotency records. The key is client-supplied, so it is data, not a
-- generated value. response_body embeds generated ids (the charge id).
CREATE TABLE idempotency_keys (
  key           text PRIMARY KEY,
  request_hash  text NOT NULL,
  response_code int NOT NULL,
  response_body jsonb NOT NULL,
  created_at    timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE webhook_events (
  id          text PRIMARY KEY DEFAULT 'we_' || replace(gen_random_uuid()::text, '-', ''),
  event_id    text NOT NULL,
  type        text NOT NULL,
  charge_id   text REFERENCES charges(id),
  payload     jsonb NOT NULL,
  received_at timestamptz NOT NULL DEFAULT now()
);
-- @unique-event-id (removed in the B5 variant)
CREATE UNIQUE INDEX webhook_events_event_id ON webhook_events(event_id);
