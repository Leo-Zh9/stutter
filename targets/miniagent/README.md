# miniagent

A small support agent built the way real LangGraph apps are, used to measure stutter's agent mode. It is **simulated**: one app, one seeded defect per variant, each a mistake real agent code makes. The "correct" variant was not correct at first: stutter's first hunt of it found two real defects in it, described below.

- FastAPI serves a LangGraph 1.2 graph: an `agent` node calling an OpenAI-compatible model with three tools bound, and the prebuilt `ToolNode`.
- `PostgresSaver` checkpoints every step at LangGraph's default durability, `"async"`.
- `issue_refund` pauses with `interrupt()` until a human approves through `POST /threads/{id}/approve`.
- `POST /chat` carries a client `message_id`. The correct app claims it first, holds an advisory lock while the run is in flight, returns the stored reply to a late duplicate, and finishes a run that a crash cut short.

The model is stutter's scripted model ([model.json](model.json)); no API key is used. The workload ([workload.py](workload.py)) runs four conversations: a refund that is approved, a note, a status question, and a refund that is rejected.

| Variant (`BUG=`) | Defect | The fault that should find it |
|---|---|---|
| (unset) | none: the false-alarm measurement | none |
| `approval-no-claim` | The refund tool reads the approval's status, then refunds and marks it approved. Two approvals that both read `pending` both refund. This is the double-clicked-approval bug. | `concurrent` on the approve request |
| `chat-no-dedupe` | `/chat` ignores `message_id`. A retried message starts the run again and opens a second approval request. | `lost` or `concurrent` on a chat request |
| `claim-before-refund` | The approval is claimed in one commit and the refund written in the next. A crash between them leaves the approval `approved` with no refund, and the retry finds nothing left to claim. The refund is never issued. | `crash` on the approve request |
| `note-keyed-on-tool-call` | `add_order_note` dedupes on the model's `tool_call_id`. If the agent node is re-run, a real model mints a new id, and the note is written twice. | `crash` on the chat request, if the agent's checkpoint write loses the race with the tool's commit |
| `write-before-interrupt` | `issue_refund` inserts its approval row before `interrupt()`. LangGraph re-runs the node on resume, so the row is written twice even in a fault-free run. | A differential oracle cannot see the duplicate every run has. It sees the extra ones: a crash or a double-clicked approval adds re-runs of the node, and each adds a row. |

Run them all with `make agent-bench` (results in [results/agents.md](../../results/agents.md)), or one with:

```bash
./bin/stutter hunt -config targets/miniagent/stutter.json -env BUG=approval-no-claim -out targets/miniagent/hunt-approval-no-claim
```

## Results

From [results/agents.md](../../results/agents.md) (`make agent-bench`, 27 min on a 10-core Apple Silicon Mac). Each cell: single-fault schedules that changed the app's tables / schedules that fired deterministically.

| Variant | lost | concurrent | crash | Distinct minimal reproductions | Reproduced in 10 fresh replays |
|---|---|---|---|---|---|
| correct | 0 / 6 | 0 / 6 | 0 / 61 | 0 | |
| approval-no-claim | 0 / 6 | **1 / 6** | 0 / 61 | 1 | 10 |
| chat-no-dedupe | 0 / 6 | **2 / 5** | 0 / 53 | 2 | 10, 10 |
| claim-before-refund | 0 / 6 | 0 / 6 | **1 / 62** | 1 | 10 |
| note-keyed-on-tool-call | 0 / 6 | 0 / 6 | **1 / 61** | 1 | **4** |
| write-before-interrupt | 0 / 6 | **2 / 6** | **10 / 60** (3 nondetermined) | 12 | 10 in 11 of them, 7 in one |

- **No false alarms.** On the correct variant, all 73 single faults fired, none diverged from the script, none was nondetermined, and none changed the app's tables.
- **Each defect is found only by the fault its design predicts.** The double-clicked approval needs `concurrent` on the approve request; every one of its 61 crash points and 6 lost responses passes. The refund that is never issued needs a crash between two commits of the approval; no retry finds it.
- **Minimal and scriptable.** The approval race shrinks to two requests and one fault: `POST /chat`, `POST /threads/t-ada/approve` sent twice at once. Its curl `repro.sh`, run against the scripted model, reproduced it 10 of 10 times.
- **Async durability is a real race, and stutter says so.** `note-keyed-on-tool-call` duplicates a note only when the agent node's checkpoint write loses the race with the tool's commit, so its crash finding reproduces 4 of 10 times and does not shrink stably. Under the hunt protocol's 10-of-10 rule it would not count. Crash-twice did not catch it: both of its first two runs happened to fail.
- **chat-no-dedupe is found by `concurrent`, not `lost`.** A retried message after the first finished re-runs the turn, but the approval is keyed by order, so no second approval appears. Two copies sent at once both pass the approval's check-then-insert.

### What the hunt found in the "correct" variant

The first version of the correct variant failed 16 of 73 schedules. Both causes are real mistakes, and both are now fixed in the variant:

1. **Recovery that trusted `get_state().next`.** After a crash right after the input checkpoint, the thread's latest checkpoint has one pending task whose writes are already saved, and LangGraph reports `next == ()`. The recovery code read that as "finished" and stored the user's own message as the reply: the refund was never requested. Recovery now resumes whenever `get_state().tasks` is not empty.
2. **An approval keyed on `tool_call_id`.** Under async durability the tool can insert its approval before the agent's checkpoint write is durable. A crash in that window re-runs the agent, a real model mints a new tool-call id, and a second approval appears. Approvals are now keyed by the business action (this order in this thread).
