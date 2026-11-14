-- MUTANT drop-index: remove the unique index that backs idempotency keys.
DROP INDEX _default.logs_idempotency_key;
