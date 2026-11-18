# Agent mode

Agent mode points stutter at an LLM agent app: a service whose requests run a LangGraph graph that calls a model, executes tools, pauses for human approval with `interrupt()`, and checkpoints to PostgreSQL. It asks the same question stutter asks of a billing webhook: **does a retried, duplicated or crashed request leave the app's tables exactly as a fault-free run does?** For agents, the side effects that matter are the ones tools write: a refund issued twice, an approval acted on twice, a note added twice, an action approved but never carried out.

## Why agents need this

LangGraph re-runs a node from its start when a run resumes, after an approval or after a crash ([durable execution](https://docs.langchain.com/oss/python/durable-execution)). Its default durability, `"async"`, writes a step's checkpoint in the background while the next step already runs. So a tool can commit its side effect before the checkpoint that records it is durable. Whether that is safe is entirely up to the app's own code: how it claims an approval, how it dedupes a retried message, whether its side effect and its claim share a transaction. The framework's behaviour is documented and filed upstream ([#8039](https://github.com/langchain-ai/langgraph/issues/8039), [#9185](https://github.com/langchain-ai/langgraph/issues/9185)); what nobody tests is the app built on top.

## The scripted model

A real model makes runs nondeterministic and costs money. Recording its answers breaks as soon as a fault changes the prompt. stutter instead serves a **scripted model**: an OpenAI-compatible server ([internal/agentstub](../internal/agentstub/agentstub.go)) started for every replay, which the app reaches through the `{MODEL_URL}` placeholder.

```json
{"turns": [
  {"user": "refund order 1001",
   "steps": [
     {"tool_calls": [{"name": "lookup_order", "args": {"order_id": 1001}}]},
     {"tool_calls": [{"name": "issue_refund", "args": {"order_id": 1001, "amount_cents": {"$ref": "tool:0:total_cents"}}}]},
     {"content": "I have handled the refund for order 1001."}
   ]}
]}
```

- **It is a pure function of the conversation.** It finds the last user message, picks the turn whose `user` text that message contains, and answers with step *k*, where *k* is the number of assistant messages since. A node LangGraph re-runs sends the same conversation and gets the same tool call.
- **Ids behave like a real model's.** A tool-call id is fresh every time the model is called, even with the same prompt, so an app that dedupes on `tool_call_id` is tested against what a real model would do when its node re-runs. Only the n-th answer to a given prompt is stable across replays, which keeps fault-free runs identical. Because they differ between equivalent runs, the oracle treats every id the stub minted like a database-generated value: an app that stores one in its tables compares as `<model-tool-call-id>` or `<model-message-id>`, not as the raw string.
- **It never manufactures a duplicate.** If the conversation already holds a completed run of the same request, the model answers the final reply with no tool calls, as a model that can see the action already happened would.
- **`$ref`** reads an earlier tool result in the same turn (`tool:0:total_cents`, or `tool:last:...`).
- **`side` rules** answer calls that are not steps of the conversation: a router's structured output, a title generator (`{"when": {"no_tools": true, "system_contains": "title"}, "content": "..."}`).
- **Embeddings** are deterministic unit vectors per input string.
- **A script miss** (no matching turn, a step past the end, a tool the request does not offer) is answered with a 500 and counted. In a fault-free run it is an error: fix the script. In a faulted run the schedule is **diverged**, which, like a fault that did not fire, is never a pass and never a failure.

The model's choices are fixed by the script, not sampled. That is the point: the script is the workload, and the app's code is what is tested.

## Configuration

```json
{
  "name": "miniagent",
  "start": [".venv/bin/python", "app.py", "--port", "{PORT}"],
  "env": ["DATABASE_URL={DATABASE_URL_SOCKET}", "OPENAI_BASE_URL={MODEL_URL}", "OPENAI_API_KEY=sk-stutter-scripted"],
  "ignore_tables": ["checkpoints", "checkpoint_blobs", "checkpoint_writes", "checkpoint_migrations"],
  "retry_rules": [
    {"method": "POST", "path": "/chat", "body_key": "message_id"},
    {"method": "POST", "path": "/threads/*/approve", "resume": {"thread": "path:2"}}
  ],
  "headers": {"Authorization": "Bearer ${TEST_TOKEN}"},
  "model": {"script": "model.json"}
}
```

| Field | Meaning |
|---|---|
| `model.script` | The scripted model. `{MODEL_URL}` is filled per run. |
| `ignore_tables` | Left out of the oracle and of the lock barrier: the checkpointer's own tables. |
| `retry_rules` | Which API calls a well-behaved client may retry, beyond the `Idempotency-Key` header. |
| `retry_rules[].body_key` | A per-action id in the body (a client message id). Its presence makes the request retryable. A thread id is not one: two messages to a thread are two messages. |
| `retry_rules[].resume` | A request that answers a pending `interrupt()`, keyed only by thread (`path:N` or `body:dotted.path`). |
| `headers` | Static headers sent on every replayed request, `${ENV}` expanded, never recorded. |
| `model.interrupt_sql` | Overrides how the pending interrupt on a thread is named (default: LangGraph's Postgres checkpointer: its latest checkpoint id, if that checkpoint has an `__interrupt__` write). |

## Faults

| Fault | Agent meaning |
|---|---|
| `lost(#n)` | The chat or approve response is lost and the client resends it. |
| `concurrent(#n)` | A double-clicked approval, or a chat message sent twice at once. The lock barrier holds both attempts at their first write to an app table; the checkpointer's tables stay unlocked so its writes cannot release them early. |
| `crash(#n@commitK)` | The service is killed right after the K-th commit of a request, a checkpoint write or an app write, then restarted, and the client retries. |

**Resumes are retried only when it is safe to.** An approval carries no per-action id, but an interrupt can be answered only once, so a second copy is a duplicate by definition. Two rules keep stutter from blaming the app for a client's mistake:

- **When faults are listed,** two fault-free observation runs record whether an interrupt was pending before and after each resume. A resume admits `concurrent` only if one was pending before it, and `lost` or `crash` only if, in addition, no new one is pending after it.
- **While a run is in flight,** stutter names the interrupt pending before the first attempt, and checks again before every retry. A retry is sent only if no interrupt is pending (the first attempt finished the run) or the same one still is (it did not take effect). If a different one is pending, the retry would approve something the user never saw, so it is refused and the fault counts as not fired. This also holds for reduced schedules during shrinking, where the observation runs no longer apply.

The interrupt is named by the thread's latest checkpoint id, if that checkpoint holds an `__interrupt__` write. The check fails safe: an unreadable answer refuses the retry.

**Crash-twice.** Under async durability, other connections keep committing while stutter holds the crashed request's commit acknowledgement, so the next node can race the kill. Every crash fault is therefore run twice. If the two runs disagree, the fault is reported as **nondetermined** and is never a finding.

## Results on miniagent

See [targets/miniagent](../targets/miniagent/README.md).

## Reproducing a finding by hand

Each finding's `repro.sh` sends the faulted order with curl. An agent app needs its model while it runs, so start the scripted model first and point the app at it:

```bash
./bin/stutter model -config targets/miniagent/stutter.json    # prints http://127.0.0.1:PORT/v1
```

Then start the app with `OPENAI_BASE_URL` set to that URL on a fresh database, and run `BASE_URL=http://127.0.0.1:APP_PORT ./repro.sh`. `stutter hunt` does exactly this for every scriptable finding (crash findings are not scriptable: curl cannot stop a process at a commit). Races are hit only some of the time without stutter's lock barrier.

## Limits

- **LangGraph's Postgres checkpointer, OpenAI-compatible models.** The interrupt query and `ignore_tables` assume the Python or JS `PostgresSaver` schema; `interrupt_sql` overrides the former. Static breakpoints (`interrupt_before`) leave no `__interrupt__` write and are not faulted as resumes.
- **The async-durability race is reported, not controlled.** Whether a tool's commit lands before the agent's checkpoint write is decided inside the app. stutter's crash points are numbered by arrival, so a defect that needs one particular order is hit only sometimes; it shows as fewer than 10 of 10 reproductions, and crash-twice can miss it.
- **Scripts are written by hand,** from the tool schemas the app sends to the model. A conversation the script does not cover is a miss, never a guess.
