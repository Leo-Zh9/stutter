# Prior art

| Work | What it checks | How stutter differs |
|---|---|---|
| Toxiproxy, xk6-disruptor, Filibuster | Inject network and service faults | Judging data corruption is left to hand-written assertions; stutter's oracle needs none |
| Jepsen | Consistency models of databases, over operation histories | stutter tests the API built on top, against its own fault-free behaviour |
| Wasabi (SOSP 2024) | Retry bugs inside Java systems, with crash, hang and missing-backoff oracles | Black-box HTTP plus Postgres, with an oracle on the data itself |
| ALICE (OSDI 2014), CrashMonkey (OSDI 2018) | Crash consistency of applications and file systems, by replaying their file-system calls up to each possible crash point | stutter applies the same idea one layer up: it crashes a black-box service right after each database commit it makes |
| RESTler | Use-after-free, namespace and resource-hierarchy rules, invalid ids | No duplicate, lost-response or reordering faults |
| Schemathesis | 5xx responses, use-after-free, resource availability | Same |
| Sahin and Arcuri, [arXiv 2608.16239](https://arxiv.org/abs/2608.16239) | HTTP-semantics oracles on responses in EvoMaster, including non-idempotent `PUT` | stutter judges database state, not responses, and models faults at the commit boundary |

What stutter adds is the combination of three things:

- a database-state differential oracle;
- faults placed at the commit boundary: lost-after-commit, a lock barrier for true concurrency, and crash points at commit acknowledgements seen through a wire-protocol proxy, plus webhook reordering;
- shrinking to a minimal reproduction.

The differential idea on its own is ordinary metamorphic testing.

Checked 2026-11-14: [RESTler checkers](https://github.com/microsoft/restler-fuzzer/blob/main/docs/user-guide/Checkers.md), [Schemathesis checks](https://schemathesis.readthedocs.io/en/stable/reference/checks/).

## Agent mode

Checked 2026-11-18.

| Work | What it checks | How stutter's agent mode differs |
|---|---|---|
| [RetryRealistic](https://neurips.cc/virtual/2026/173038) (NeurIPS 2026) | Whether the *model* recovers honestly from 9 tool fault types, in a synthetic sandbox with 8 tools | The model is pinned by a script; the question is whether real *app code* guards its side effects under crash and resume |
| [Verified tool calls](https://arxiv.org/abs/2608.02645) (arXiv 2608.02645) | A tool wrapper with postcondition checks, verify-before-retry and idempotency keys, evaluated in simulation | A mitigation, not a tester; its techniques are the fixes stutter's findings point to |
| [AgentCheck](https://arxiv.org/abs/2607.11098) (arXiv 2607.11098) | A workbench that records an agent's MCP tool responses, replays them perturbed by 12 fault types, and goes live after the agent diverges | It perturbs what tools *return* to the model; stutter faults what the app *does* (retries, double approvals, crashes at commits) and judges the database. No checkpoints, no hunt on real apps. Replaying a fixed model is not claimed as new |
| [ACRFence](https://arxiv.org/abs/2603.20625) (arXiv 2603.20625) | Defines action-replay attacks on agent checkpoint-restore (a restored agent re-generates a payment) and a framework-agnostic mitigation that enforces replay-or-fork semantics, shown in a proof of concept | A mitigation for the framework layer; stutter tests whether an app's own code survives the re-runs |
| [Agent tool-call failure testing](https://www.kunalganglani.com/blog/agent-tool-call-failure-testing) | A CI recipe: scripted HTTP faults per scenario step and a ledger assertion of one charge per key | Hand-written assertions on your own agent; no crash points, no shrinking |
| [Exactly-once and agent retries](https://tianpan.co/blog/2026-07-01-exactly-once-was-hard-before-your-agent-could-retry-itself) | Design patterns: workflow-position idempotency keys, an action ledger | Advice; stutter checks whether an app follows it |
| LangGraph [#8039](https://github.com/langchain-ai/langgraph/issues/8039), [#9185](https://github.com/langchain-ai/langgraph/issues/9185), [#7417](https://github.com/langchain-ai/langgraph/issues/7417), [#9183](https://github.com/langchain-ai/langgraph/issues/9183) | The framework re-runs nodes on resume and after crashes; documented and filed | stutter treats that as given and tests the app's code built on it |

What agent mode adds is the same combination as before, applied to agents: a database-state oracle with no assertions, faults at the commit boundary (including the checkpointer's commits) and at approvals, and shrinking to a minimal reproduction, plus a scripted model that makes all of it deterministic and free.
