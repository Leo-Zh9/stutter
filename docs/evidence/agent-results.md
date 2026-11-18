# Agent blind results

This reports the agent hunts under [agent-protocol.md](agent-protocol.md), frozen in `2bb512e`. Every number comes from a file in this repository. Each target's doc has the details: [Forge](../targets/forge.md), [paid-media-agent](../targets/paid-media-agent.md), [deer-flow](../../targets/agents/deer-flow/README.md).

## Order of events

Times are from `git log` and the local reflog of `origin/main`, UTC, 2026-11-18.

| Time | Commit | Event |
|---|---|---|
| 10:02:56 | `2bb512e` | Protocol frozen, before any search |
| 10:16:50 | `1cabe04` | Frozen search output (235 repositories) and mechanical pre-filter (41 kept) |
| 10:34:13 | `88970d1` | Screen: 41 judged, 3 pass. Reasons held back under sha256 `980b062c…` |
| 10:43:20 | `05e6ced` | Pushed with agent mode. Running third-party code then waited for the owner's approval |
| 18:13:25 | `de16b2d` | Hunt results for Forge and paid-media-agent, deer-flow's not-run write-up, pushed. No application source read before this |
| 18:13:44 | `88249b7` | Screener's reasons published as [agent-reasons.md](agent-reasons.md). It hashes to `980b062cf17ffa4f0883a9b39dbf056bfd31734bf96054f121106c7640cad6b5`, the hash published at 10:34 |

## Counts

| | Count | Repositories |
|---|---|---|
| Judged | 41 | All of [agent-prefilter.md](agent-prefilter.md) |
| Passed (K) | **3** | bytedance/deer-flow, nihalashetty/Forge, langchain-ai/paid-media-agent |
| Found defects | 2 | Forge, paid-media-agent |
| Clean (C) | 0 | |
| Not reached (R) | 0 | Both hunts fired crash faults on requests that write business tables |
| Not run (N) | 1 | deer-flow: the only tool that writes its tables is offered only to browser login sessions |
| Distinct root causes (M) | **4** | Forge 2, paid-media-agent 2. **2** if only a tool's side effect running twice or never counts (see below) |

The list ran out with 3 passes. The protocol seeks 4, so K is 3.

## Per target

| | Forge | paid-media-agent |
|---|---|---|
| Requests in the trace | 4 | 6 |
| Single-fault schedules | 55 | 132 |
| `concurrent` fired | 0 of 4 | 0 of 3 |
| Crash schedules that changed the end state | 42 of 47 | 92 of 126 (7 nondetermined) |
| Distinct end-state changes | 8 | 9 |
| Findings, all 1-minimal, all 10 of 10 | 42 | 92 |
| Root causes counted | 2 | 2 |
| Not counted | — | extra unused approval claim (no app-visible effect); `executing`/`verifying` after a crash (documented design); fixture provider's in-memory counter (harness artifact) |

Root causes:

1. **Forge: a retried approval is refused, and the approved action is dropped.** The resume commits `running` before acting, and a retry only accepts `interrupted` runs. The approved `remember` call never runs.
2. **Forge: the Idempotency-Key is remembered in process memory and only after the run completes.** A retry after a crash starts a second run and leaves a second pending approval. In this workload the duplicated run did not repeat a tool call, because the gate comes first.
3. **paid-media-agent: a retried approval returns 200 but never executes.** `resume` returns when no interrupt is pending and ignores the tool task still pending after it. The budget change is never sent.
4. **paid-media-agent: the terminal state is saved before the receipt.** A crash in between loses the receipt of a change that was applied and verified.

## The claim

> Against frozen per-target model scripts, stutter ran on 3 open-source LangGraph apps chosen by a fixed, published repo search and pass/fail criteria. It found 4 distinct root causes where a tool's side effect ran twice or never (2 if the wording is read strictly; see below). 0 apps were clean, 0 were not reached and 1 was not run.

**Where the rule and the wording disagree.** The protocol counts a finding when a business-table row is duplicated or dropped. The claim's template says "a tool's side effect ran twice or never". Causes 1 and 3 fit both: an approved tool call never ran. Cause 2 duplicates a run and a pending approval, not a tool call. Cause 4 loses the receipt of a tool call that ran once. Reported under the rule as written, M is 4. Under the template's wording alone, it is 2.

## Disclosures

- **stutter changed during the hunts.** Three changes, all in `f454ed1`:
  - `generated` declares ids an app mints in its own code.
  - `ignore` accepts paths inside JSON columns.
  - The replayer opens a fresh connection per request.

  Each was needed to get a stable baseline or a complete run. None changes how faults are chosen or judged. Every ignore and generated column is listed in the target's `stutter.json` and doc.
- **Runtime output that named code.** A paid-media-agent server traceback, seen before the push, named source files and lines.
- **Fixture data.** paid-media-agent ran in its sample data mode. Its ad platforms are in-process fixtures, so one crash effect there is an artifact and is not counted.
- **deer-flow's docs and gateway disagree.** The auth design says internal-token requests take no part in CSRF. The gateway required a CSRF pair for them. This did not change the outcome: no static-token caller is offered the table-writing tool.
- **No issues were filed upstream.** Drafts are in each target's doc.
