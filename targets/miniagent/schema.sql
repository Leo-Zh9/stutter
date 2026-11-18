-- miniagent's own tables. LangGraph's checkpoint tables are created by
-- PostgresSaver.setup() and are ignored by stutter's oracle.
CREATE TABLE IF NOT EXISTS orders (
  number      integer PRIMARY KEY,
  customer    text NOT NULL,
  total_cents integer NOT NULL,
  status      text NOT NULL
);
CREATE TABLE IF NOT EXISTS chat_requests (
  message_id text PRIMARY KEY,
  thread_id  text NOT NULL,
  status     text NOT NULL DEFAULT 'running',
  response   jsonb,
  created_at timestamptz NOT NULL DEFAULT now()
);
CREATE TABLE IF NOT EXISTS approvals (
  id           uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  thread_id    text NOT NULL,
  tool_call_id text NOT NULL,
  order_number integer NOT NULL REFERENCES orders(number),
  amount_cents integer NOT NULL,
  status       text NOT NULL DEFAULT 'pending',
  created_at   timestamptz NOT NULL DEFAULT now()
);
CREATE TABLE IF NOT EXISTS refunds (
  id           uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  approval_id  uuid NOT NULL REFERENCES approvals(id),
  order_number integer NOT NULL REFERENCES orders(number),
  amount_cents integer NOT NULL,
  created_at   timestamptz NOT NULL DEFAULT now()
);
CREATE TABLE IF NOT EXISTS outbox (
  id           uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  kind         text NOT NULL,
  order_number integer NOT NULL,
  created_at   timestamptz NOT NULL DEFAULT now()
);
CREATE TABLE IF NOT EXISTS order_notes (
  id           uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  order_number integer NOT NULL REFERENCES orders(number),
  note         text NOT NULL,
  dedupe_key   text NOT NULL UNIQUE,
  created_at   timestamptz NOT NULL DEFAULT now()
);
