# Prior art

| Work | What it checks | How stutter differs |
|---|---|---|
| Toxiproxy, xk6-disruptor, Filibuster | Inject network and service faults | Judging data corruption is left to hand-written assertions; stutter's oracle needs none |
| Jepsen | Consistency models of databases, over operation histories | stutter tests the API built on top, against its own fault-free behaviour |
| Wasabi (SOSP 2024) | Retry bugs inside Java systems, with crash, hang and missing-backoff oracles | Black-box HTTP plus Postgres, with an oracle on the data itself |
| RESTler | Use-after-free, namespace and resource-hierarchy rules, invalid ids | No duplicate, lost-response or reordering faults |
| Schemathesis | 5xx responses, use-after-free, resource availability | Same |
| Sahin and Arcuri, [arXiv 2608.16239](https://arxiv.org/abs/2608.16239) | HTTP-semantics oracles on responses in EvoMaster, including non-idempotent `PUT` | stutter judges database state, not responses, and models faults at the commit boundary |

What stutter adds is the combination of three things:

- a database-state differential oracle;
- faults placed at the commit boundary (lost-after-commit, and a lock barrier for true concurrency) plus webhook reordering;
- shrinking to a minimal reproduction.

The differential idea on its own is ordinary metamorphic testing.

Checked 2026-11-14: [RESTler checkers](https://github.com/microsoft/restler-fuzzer/blob/main/docs/user-guide/Checkers.md), [Schemathesis checks](https://schemathesis.readthedocs.io/en/stable/reference/checks/).
