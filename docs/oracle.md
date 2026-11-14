# The oracle

Two runs of the same correct program never produce byte-identical databases. Uuids differ, sequences skip values that a rolled-back duplicate consumed, and timestamps move. The oracle factors that out using only the PostgreSQL catalog. Code: [internal/oracle/oracle.go](../internal/oracle/oracle.go).

## Generated values

A column is **generated** if its default calls a non-immutable function (`gen_random_uuid`, `nextval`, `now`, and so on), uses `CURRENT_TIMESTAMP` or a similar keyword, or is an identity column. Its values are never compared directly.

## Edges

Something that points at a generated value becomes an **edge** to the row that owns it:

- a single-column foreign key to a generated column;
- any string in a `text` or `jsonb` column that exactly equals a generated value.

The second rule is what keeps the cached response in an idempotency table tied to the charge it embeds. A row that contains its own generated id records `<self>` instead.

## Signatures

Each row gets a hash built bottom-up:

1. Take its own non-generated, non-edge columns. Timestamp columns are dropped. RFC 3339 strings inside JSON become `<ts>`. Edge values become `<ref:table>`.
2. Add the sorted signatures of every row that points at it, each labelled with the referencing table and column.

This is AHU tree hashing, the standard technique for tree isomorphism. Two states are equal when every table holds the same multiset of row signatures.

Because of that, two customers identical except for their ids, each with a different charge, compare equal in any insert order. Any of these is a diff:

- moving a charge from one customer to the other;
- adding a row;
- changing one amount;
- pointing a cached response at a different row.

A reference cycle between rows is reported as an error.

## Guard against false alarms

Before exploring, stutter replays the trace fault-free twice. If the two runs disagree, the target has nondeterminism the catalog cannot see, so stutter names the unstable columns instead of producing false alarms.

On minipay this caught `idempotency_keys.request_hash`, a hash of a request body that contains a generated customer id. One flag marks it: `-ignore idempotency_keys.request_hash`.

## Limits

- **Sound, not complete.** Equivalent states always hash equal, which is why false alarms stay at zero. Two different states can still collide when a row references two parents and a bug only re-pairs them.
- **Exact strings only.** Integer ids stored as JSON numbers, and partial matches inside longer strings, are not edges.
- **Single-column foreign keys only.**
- **Timestamps are ignored,** so a bug whose only effect is a wrong timestamp is missed.
