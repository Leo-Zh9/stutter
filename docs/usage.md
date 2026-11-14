# Running stutter on your own service

stutter needs four things from a service:

- it reads its database address from the environment or its command line;
- it listens on an address it is given;
- it answers a health path with a status below 500;
- it keeps its state in PostgreSQL.

Describe the service in a `stutter.json`.

```json
{
  "name": "myapp",
  "start": ["node", "server.js", "--port", "{PORT}"],
  "env": ["DATABASE_URL={DATABASE_URL}"],
  "health": "/healthz",
  "start_timeout": "10s",
  "prepare": {"migrate_cmd": ["npx", "prisma", "migrate", "deploy"], "migrate_env": ["DATABASE_URL={DATABASE_URL}"], "seed_file": "fixtures.sql"},
  "schemas": ["public"],
  "idempotency_header": "Idempotency-Key",
  "webhooks": {"prefix": "/api/webhooks/stripe", "signing": "stripe", "secret": "whsec_test_only", "retry": "non-2xx", "max_attempts": 5},
  "ignore": ["sessions.token"],
  "trace": "trace.jsonl"
}
```

| Field | Meaning |
|---|---|
| `start`, `env` | Command and extra environment. Placeholders are filled per run (see below). |
| `health` | Polled until it answers below 500. Defaults to `/`. |
| `start_timeout` | How long to wait for health. Defaults to `5s`. |
| `prepare` | Builds the template database, using exactly one of `schema_file`, `schema_cmd` (prints SQL) or `migrate_cmd` (migrates `{DATABASE_URL}`). `seed_file` then adds fixture rows. |
| `schemas` | Schemas to snapshot and to lock for the barrier. Tables outside `public` appear as `schema.table`. |
| `idempotency_header` | API calls without this header are never retried or faulted. |
| `webhooks.prefix` | Requests under this path are provider deliveries, which may always be redelivered. |
| `webhooks.signing` | `stripe` signs each delivery at send time (`t=...,v1=HMAC-SHA256`), since recorded signatures expire after 300 s. |
| `webhooks.retry` | `5xx` (default) or `non-2xx`. Stripe retries any non-2xx. A diff only counts after the provider's retries are used up. |
| `ignore` | `table.column` values the oracle should skip. stutter names any that are needed: see the false-alarm guard in [oracle.md](oracle.md). |

Placeholders, usable in `start`, `env` and `prepare.migrate_*`:

| Placeholder | Value |
|---|---|
| `{DSN}` | libpq keyword string over the project-local unix socket |
| `{DATABASE_URL}` | `postgres://postgres@127.0.0.1:54329/<db>?sslmode=disable` |
| `{DATABASE_URL_SOCKET}` | the same over the unix socket; prefer it for services with large pools, since TCP connections closed by killed processes pile up in TIME_WAIT and can exhaust ephemeral ports |
| `{ADDR}` | `127.0.0.1:<port>` the service must listen on |
| `{HOST}`, `{PORT}` | the two halves of `{ADDR}` |

The service runs in its own process group, and all of it is killed between runs.

Record a trace by pointing your workload, test suite or a script at the proxy:

```bash
./bin/stutter record -config stutter.json -- ./my-workload --base-url {proxy}
```

Then explore it:

```bash
./bin/stutter explore -config stutter.json
```

Exit codes: `0` no fault changed the final state, `1` a failure was found and shrunk, `2` stutter itself failed.
