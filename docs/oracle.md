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

## Integer ids minted in application SQL

Some services mint ids with `nextval()` in their own SQL, so no column default reveals them, and store them as integers inside JSON, or even as decimal strings. Formance Ledger does all three. For these targets, `prepare.band_sequences` turns on two more steps.

**Bands.** After the template is built, every sequence k is moved to the start of its own range:

- `bigint` sequences start at (k+1)·10^12;
- `integer` sequences start at (k+1)·10^7;
- `smallint` sequences, or ones with no room below `max_value`, are left out and listed.

An integer inside a band is therefore always a minted id, and its band says whose. The replayer maps in-band numbers from recorded to live values, both in request bodies and in URL path segments such as `/transactions/4000000000003/revert`.

**Inference.** Before exploring, stutter replays the trace fault-free twice:

- **A**: on a normal clone;
- **B**: on a clone where every banded sequence was advanced by its own offset, (k+1)·10^6 for `bigint`.

For every column and every JSON path, it compares the multiset of values in A and B. No row matching is needed. A path whose in-band values in B are exactly A's values plus sequence k's offset holds sequence k's ids. Values outside every band (rows the template already held) must be identical. Anything else that differs is volatile: a hash over ids, or an app-side clock. Volatile paths are left out of the comparison and listed.

**Owners and edges.** The column that owns sequence k is read from the sequence's `OWNED BY` link in `pg_depend` (serial and identity columns have one, and Formance declares its own). The owner then acts like a generated column, and every other path holding k's ids becomes an edge to the owner row.

**When inference stops.** It refuses to guess:

- a sequence that minted values with no owner;
- an owner whose values do not shift with its sequence;
- an owner whose values repeat.

Each of these stops the run with the sequence named.

**Self-check.** Afterwards, A, a second normal run, and B must all canonicalize to the same state, or nothing runs. On Formance this self-check caught revert links stored as decimal strings in metadata (`"com.formance.spec/state/reverts": "4000000000007"`), which the first version mapped only as JSON numbers.

## Guard against false alarms

Before exploring, stutter replays the trace fault-free twice. If the two runs disagree, the target has nondeterminism the catalog cannot see, so stutter names the unstable columns instead of producing false alarms.

On minipay this caught `idempotency_keys.request_hash`, a hash of a request body that contains a generated customer id. One flag marks it: `-ignore idempotency_keys.request_hash`.

## Limits

- **Sound, not complete.** Equivalent states always hash equal, which is why false alarms stay at zero. Two different states can still collide when a row references two parents and a bug only re-pairs them.
- **Exact values only.** Without `band_sequences`, integer ids inside JSON are not edges. Partial matches inside longer strings never are.
- **Single-column foreign keys only.**
- **Timestamps are ignored,** so a bug whose only effect is a wrong timestamp is missed.
