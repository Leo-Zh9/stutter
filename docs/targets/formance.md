# Formance Ledger v2.4.18

| | |
|---|---|
| Repo | [formancehq/ledger](https://github.com/formancehq/ledger), MIT. v2.4 is the PostgreSQL line; `main` is v3, which does not use PostgreSQL |
| Version | release binary `ledger_darwin-arm64.tar.gz` from v2.4.18 (2026-11-05), checked against `ledger_checksums.txt`, commit `83cd8dc7a` |
| Code changed | none for the clean result; mutants are separate builds from the v2.4.18 source with one patch each |
| Config | [targets/formance/stutter.json](../../targets/formance/stutter.json) |
| Results | [targets/formance/hunt/REPORT.md](../../targets/formance/hunt/REPORT.md), [mutants](../../targets/formance/mutants) |

## Why this target

A double-entry ledger is the case idempotency keys exist for. Formance documents the `Idempotency-Key` header for transactions, reverts and metadata. Its design layers two defences:

1. `fetchLogWithIK` returns the stored result before doing any work;
2. a unique index `logs_idempotency_key` catches concurrent duplicates, after which `forgeLogRetry` fetches the winner's log.

The prediction before running was a clean result. This target tests whether the oracle can stay silent on real, complex code, and whether stutter still catches each defence when it is removed.

## What it took to run unmodified

- **pgcrypto.** The project-local PostgreSQL (from `pgserver`) ships no contrib modules, and Formance's migrations create `pgcrypto`. The v2.4.18 source calls exactly one function from it: `digest(..., 'sha256')`, in 7 migrations. stutter installs a 20-line SQL extension ([scripts/pgcrypto-shim](../../scripts/pgcrypto-shim)) that implements `digest()` on PostgreSQL's built-in `sha224/256/384/512` and `md5`. Its output is byte-identical to openssl on test vectors. Formance's code and binary are untouched.
- **Setup through the service.** Formance migrates itself on start (`--auto-upgrade`). The template is built by starting it once, creating the `default` ledger, and posting one setup transaction.
- **Why the setup transaction.** A ledger's first write moves it from `initializing` to `in_use` and resets its id sequences to `max(id)+1` (`internal/controller/system/state_tracker.go:96-123`). Without the setup transaction, that reset would undo the sequence bands below.
- **Integer ids from application SQL.** Transaction and log ids come from `nextval()` in Formance's own queries. They appear as integers in JSON responses, request paths (`/transactions/{id}/revert`), jsonb log payloads, and as decimal strings in revert metadata. This is what [sequence bands and id inference](../oracle.md#integer-ids-minted-in-application-sql) were built for.
- **Unix socket.** Formance opens up to 20 connections per process. Restarting it per replay over TCP left about 9,000 sockets in TIME_WAIT and exhausted ephemeral ports, so it connects through `{DATABASE_URL_SOCKET}`.
- **Workload.** A labelled driver, [workload.mjs](../../targets/formance/workload.mjs), sends 43 requests, each with an `Idempotency-Key`:
  - 30 transactions using the order/fee/seller numscript from Formance's own `tools/generator/examples/example1.js`;
  - 5 metadata writes and 5 reverts by transaction id;
  - 3 account metadata writes.
  - `_bulk` is not used, because Formance ignores the header there (idempotency is per element), so retrying a bulk call would be stutter's own bug.

## What inference found

[Inference](../oracle.md#integer-ids-minted-in-application-sql) is reported in [results.json](../../targets/formance/hunt/results.json):

- **Owners, from `OWNED BY`:** `transactions.id`, `logs.id`, `moves.seq`, `accounts_metadata.seq`, `transactions_metadata.seq`.
- **References:** `moves.transactions_id`, `transactions_metadata.transactions_id`, `logs.data.transaction.id`, `logs.data.revertedTransaction.id`, and the revert links stored as strings in three metadata paths.
- **Volatile, dropped and listed:** `logs.hash` and `logs.memento` (bytes over the ids), `logs.idempotency_hash`, timestamps inside log JSON, and `logs.data.targetId`.

`targetId` holds a transaction id for transaction metadata but an account address for account metadata. Because the path mixes types, it is treated as volatile and not as a reference. That is a stated limit.

## Results

| Build | lost | concurrent | crash | Changed the final state |
|---|---|---|---|---|
| release binary, unmodified | 0 / 43 | 0 / 43 | 0 / 43 | none |
| source build, unmodified (control) | 0 / 43 | 0 / 43 | 0 / 43 | none |
| unique index dropped | 0 / 43 | **38 / 43** | 0 / 43 | concurrent duplicates double transactions |
| lookup removed | 0 / 43 | 0 / 43 | 0 / 43 | none: the unique index still catches every duplicate |
| lookup removed and index dropped | **38 / 43** | **38 / 43** | **38 / 43** | every retry redoes the work |
| log written in a second transaction | 0 / 43 | **30 / 43** | **35 / 86** | a crash between the two commits, or a racing duplicate, doubles the work |

Each cell shows schedules that changed the final state out of schedules that fired; every schedule fired. There are 43 crash points per build, one commit per request. The split mutant has 86, because each request now commits twice.

- **Unmodified Formance:** 0 of 129 single faults changed the final state. Neither did any of 286 further schedules (86 singles plus 200 random combinations of 2 to 4 faults, all fired), run with `stutter explore` ([output](../../targets/formance/explore/random-200.txt)).
- **The missed cases** in the mutant rows are the 5 reverts in the trace. A duplicate revert is refused by Formance's own `ErrAlreadyReverted` check (`internal/controller/ledger/errors.go:89`), independent of the idempotency key. The split mutant's 13 surviving concurrent cases are the 5 reverts plus 8 metadata writes, which are idempotent by nature.
- **The split mutant is the same bug class as minipay's `key-after-commit`.** A lost response cannot see it, because the retry finds the log. A crash at commit 1 can, and so can a duplicate racing between the two commits.
- **One partial result.** In the "lookup removed and index dropped" build, two concurrent findings reproduced in only 6 and 4 of 10 replays. One of them did not shrink to a 1-minimal schedule. The barrier holds the first write, and in that build the race window extends past it. Every lost and crash finding reproduced 10 of 10.
