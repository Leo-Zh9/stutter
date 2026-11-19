-- miniasync. charges gets a unique index on key unless BUG=dup-charge (app.py --migrate).
CREATE TABLE accounts (
  id          text PRIMARY KEY,
  last_key    text,
  checked_key text
);
CREATE TABLE charges (
  id           serial PRIMARY KEY,
  key          text NOT NULL,
  account      text NOT NULL REFERENCES accounts(id),
  amount_cents integer NOT NULL
);
-- A step pending on thread t1 before the trace runs: the resume answers it.
CREATE TABLE steps (
  thread  text PRIMARY KEY,
  n       integer NOT NULL,
  resumed boolean NOT NULL DEFAULT false
);
CREATE TABLE notes (
  key  text PRIMARY KEY,
  text text NOT NULL
);
CREATE TABLE audit (
  key text PRIMARY KEY
);
INSERT INTO accounts (id) VALUES ('acct-1');
INSERT INTO steps (thread, n) VALUES ('t1', 0);
