# miniagent

A small support agent built the way real LangGraph apps are, used to measure stutter's agent mode. It is **simulated**: one app, one seeded defect per variant, each a mistake real agent code makes.

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
| `write-before-interrupt` (control) | `issue_refund` inserts its approval row before `interrupt()`. LangGraph re-runs the node on resume, so the row is written twice in every run, faulted or not. | none: a differential oracle cannot see a defect the fault-free run has too |

Run them all with `make agent-bench` (results in [results/agents.md](../../results/agents.md)), or one with:

```bash
./bin/stutter hunt -config targets/miniagent/stutter.json -env BUG=approval-no-claim -out targets/miniagent/hunt-approval-no-claim
```
