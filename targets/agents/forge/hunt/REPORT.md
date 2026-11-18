# stutter hunt: nihalashetty/Forge @ 25074df

Generated 2026-11-18T16:53:17Z. Trace: 4 requests. Pre-flight: every recorded webhook delivery got 2xx.

| Fault type | Schedules | Fired | Not fired | Diverged | Nondetermined | Changed the final state | No-op passes |
|---|---|---|---|---|---|---|---|
| concurrent | 4 | 0 | 4 | 0 | 0 | 0 | 0 |
| crash | 47 | 47 | 0 | 0 | 0 | 42 | 0 |
| lost | 4 | 4 | 0 | 0 | 0 | 0 | 0 |

No-op reorder pairs (dropping both events leaves the state unchanged, so their pass says nothing about ordering): 0.

## Effects: 8 distinct end-state changes

1. Findings [1]:

```
+ runs {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":null,"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"queued","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

2. Findings [2 3 4 5 6 7 8]:

```
+ runs {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":null,"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

3. Findings [9]:

```
+ runs {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"interrupted","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

4. Findings [10 11 12 13 14 15 16]:

```
+ runs {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":null,"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

5. Findings [17 36]:

```
+ runs {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"interrupted","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

6. Findings [18 19 20 21 22 23 24 25 26]:

```
- memories {"kind":"note","project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","scope":"default","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","text":"Customer ACME prefers to be contacted by email."}
+ runs {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"decision":"approve","messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] approve","name":null,"tool_calls":null,"type":"human"},{"content":"","name":null,"tool_calls":[{"args":{"text":"Customer ACME prefers to be contacted by email."},"display_name":"remember","name":"remember"}],"type":"ai"},{"content":"Saved to long-term memory.","name":"remember","tool_calls":null,"type":"tool"},{"content":"Saved: ACME prefers email.","name":null,"tool_calls":null,"type":"ai"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0.000009,"total_tokens":41,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

7. Findings [27 28 29 30 31 32 33 34 35]:

```
+ runs {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"decision":"approve","messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] approve","name":null,"tool_calls":null,"type":"human"},{"content":"","name":null,"tool_calls":[{"args":{"text":"Customer ACME prefers to be contacted by email."},"display_name":"remember","name":"remember"}],"type":"ai"},{"content":"Saved to long-term memory.","name":"remember","tool_calls":null,"type":"tool"},{"content":"Saved: ACME prefers email.","name":null,"tool_calls":null,"type":"ai"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0.000009,"total_tokens":41,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

8. Findings [37 38 39 40 41 42]:

```
+ runs {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":{"decision":"reject","messages":[{"content":"Remember that customer Globex wants invoices in EUR.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] reject","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

## Findings: 42 distinct minimal reproductions

### Finding 1

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#0@commit1).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
fault crash(#0@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #0

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":null,"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"queued","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-1/repro.sh`

### Finding 2

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#0@commit2).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
fault crash(#0@commit2)
```

Delivery order: send, kill the service after its commit 2, restart, resend #0

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":null,"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-2/repro.sh`

### Finding 3

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#0@commit3).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
fault crash(#0@commit3)
```

Delivery order: send, kill the service after its commit 3, restart, resend #0

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":null,"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-3/repro.sh`

### Finding 4

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#0@commit4).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
fault crash(#0@commit4)
```

Delivery order: send, kill the service after its commit 4, restart, resend #0

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":null,"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-4/repro.sh`

### Finding 5

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#0@commit5).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
fault crash(#0@commit5)
```

Delivery order: send, kill the service after its commit 5, restart, resend #0

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":null,"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-5/repro.sh`

### Finding 6

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#0@commit6).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
fault crash(#0@commit6)
```

Delivery order: send, kill the service after its commit 6, restart, resend #0

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":null,"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-6/repro.sh`

### Finding 7

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#0@commit7).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
fault crash(#0@commit7)
```

Delivery order: send, kill the service after its commit 7, restart, resend #0

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":null,"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-7/repro.sh`

### Finding 8

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#0@commit8).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
fault crash(#0@commit8)
```

Delivery order: send, kill the service after its commit 8, restart, resend #0

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":null,"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-8/repro.sh`

### Finding 9

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#0@commit9).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
fault crash(#0@commit9)
```

Delivery order: send, kill the service after its commit 9, restart, resend #0

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"interrupted","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-9/repro.sh`

### Finding 10

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#2@commit2).

```
#2 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=globex-1
fault crash(#2@commit2)
```

Delivery order: send, kill the service after its commit 2, restart, resend #2

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":null,"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-10/repro.sh`

### Finding 11

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#2@commit3).

```
#2 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=globex-1
fault crash(#2@commit3)
```

Delivery order: send, kill the service after its commit 3, restart, resend #2

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":null,"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-11/repro.sh`

### Finding 12

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#2@commit4).

```
#2 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=globex-1
fault crash(#2@commit4)
```

Delivery order: send, kill the service after its commit 4, restart, resend #2

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":null,"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-12/repro.sh`

### Finding 13

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#2@commit5).

```
#2 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=globex-1
fault crash(#2@commit5)
```

Delivery order: send, kill the service after its commit 5, restart, resend #2

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":null,"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-13/repro.sh`

### Finding 14

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#2@commit6).

```
#2 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=globex-1
fault crash(#2@commit6)
```

Delivery order: send, kill the service after its commit 6, restart, resend #2

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":null,"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-14/repro.sh`

### Finding 15

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#2@commit7).

```
#2 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=globex-1
fault crash(#2@commit7)
```

Delivery order: send, kill the service after its commit 7, restart, resend #2

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":null,"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-15/repro.sh`

### Finding 16

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#2@commit8).

```
#2 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=globex-1
fault crash(#2@commit8)
```

Delivery order: send, kill the service after its commit 8, restart, resend #2

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":null,"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-16/repro.sh`

### Finding 17

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#2@commit9).

```
#2 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=globex-1
fault crash(#2@commit9)
```

Delivery order: send, kill the service after its commit 9, restart, resend #2

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"interrupted","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-17/repro.sh`

### Finding 18

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit2).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
#1 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-2
fault crash(#1@commit2)
```

Delivery order: send #0, send, kill the service after its commit 2, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
- memories x1 {"kind":"note","project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","scope":"default","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","text":"Customer ACME prefers to be contacted by email."}
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"decision":"approve","messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] approve","name":null,"tool_calls":null,"type":"human"},{"content":"","name":null,"tool_calls":[{"args":{"text":"Customer ACME prefers to be contacted by email."},"display_name":"remember","name":"remember"}],"type":"ai"},{"content":"Saved to long-term memory.","name":"remember","tool_calls":null,"type":"tool"},{"content":"Saved: ACME prefers email.","name":null,"tool_calls":null,"type":"ai"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0.000009,"total_tokens":41,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-18/repro.sh`

### Finding 19

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit3).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
#1 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-2
fault crash(#1@commit3)
```

Delivery order: send #0, send, kill the service after its commit 3, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
- memories x1 {"kind":"note","project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","scope":"default","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","text":"Customer ACME prefers to be contacted by email."}
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"decision":"approve","messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] approve","name":null,"tool_calls":null,"type":"human"},{"content":"","name":null,"tool_calls":[{"args":{"text":"Customer ACME prefers to be contacted by email."},"display_name":"remember","name":"remember"}],"type":"ai"},{"content":"Saved to long-term memory.","name":"remember","tool_calls":null,"type":"tool"},{"content":"Saved: ACME prefers email.","name":null,"tool_calls":null,"type":"ai"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0.000009,"total_tokens":41,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-19/repro.sh`

### Finding 20

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit4).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
#1 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-2
fault crash(#1@commit4)
```

Delivery order: send #0, send, kill the service after its commit 4, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
- memories x1 {"kind":"note","project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","scope":"default","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","text":"Customer ACME prefers to be contacted by email."}
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"decision":"approve","messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] approve","name":null,"tool_calls":null,"type":"human"},{"content":"","name":null,"tool_calls":[{"args":{"text":"Customer ACME prefers to be contacted by email."},"display_name":"remember","name":"remember"}],"type":"ai"},{"content":"Saved to long-term memory.","name":"remember","tool_calls":null,"type":"tool"},{"content":"Saved: ACME prefers email.","name":null,"tool_calls":null,"type":"ai"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0.000009,"total_tokens":41,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-20/repro.sh`

### Finding 21

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit5).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
#1 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-2
fault crash(#1@commit5)
```

Delivery order: send #0, send, kill the service after its commit 5, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
- memories x1 {"kind":"note","project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","scope":"default","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","text":"Customer ACME prefers to be contacted by email."}
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"decision":"approve","messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] approve","name":null,"tool_calls":null,"type":"human"},{"content":"","name":null,"tool_calls":[{"args":{"text":"Customer ACME prefers to be contacted by email."},"display_name":"remember","name":"remember"}],"type":"ai"},{"content":"Saved to long-term memory.","name":"remember","tool_calls":null,"type":"tool"},{"content":"Saved: ACME prefers email.","name":null,"tool_calls":null,"type":"ai"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0.000009,"total_tokens":41,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-21/repro.sh`

### Finding 22

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit6).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
#1 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-2
fault crash(#1@commit6)
```

Delivery order: send #0, send, kill the service after its commit 6, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
- memories x1 {"kind":"note","project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","scope":"default","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","text":"Customer ACME prefers to be contacted by email."}
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"decision":"approve","messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] approve","name":null,"tool_calls":null,"type":"human"},{"content":"","name":null,"tool_calls":[{"args":{"text":"Customer ACME prefers to be contacted by email."},"display_name":"remember","name":"remember"}],"type":"ai"},{"content":"Saved to long-term memory.","name":"remember","tool_calls":null,"type":"tool"},{"content":"Saved: ACME prefers email.","name":null,"tool_calls":null,"type":"ai"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0.000009,"total_tokens":41,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-22/repro.sh`

### Finding 23

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit7).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
#1 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-2
fault crash(#1@commit7)
```

Delivery order: send #0, send, kill the service after its commit 7, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
- memories x1 {"kind":"note","project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","scope":"default","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","text":"Customer ACME prefers to be contacted by email."}
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"decision":"approve","messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] approve","name":null,"tool_calls":null,"type":"human"},{"content":"","name":null,"tool_calls":[{"args":{"text":"Customer ACME prefers to be contacted by email."},"display_name":"remember","name":"remember"}],"type":"ai"},{"content":"Saved to long-term memory.","name":"remember","tool_calls":null,"type":"tool"},{"content":"Saved: ACME prefers email.","name":null,"tool_calls":null,"type":"ai"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0.000009,"total_tokens":41,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-23/repro.sh`

### Finding 24

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit8).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
#1 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-2
fault crash(#1@commit8)
```

Delivery order: send #0, send, kill the service after its commit 8, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
- memories x1 {"kind":"note","project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","scope":"default","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","text":"Customer ACME prefers to be contacted by email."}
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"decision":"approve","messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] approve","name":null,"tool_calls":null,"type":"human"},{"content":"","name":null,"tool_calls":[{"args":{"text":"Customer ACME prefers to be contacted by email."},"display_name":"remember","name":"remember"}],"type":"ai"},{"content":"Saved to long-term memory.","name":"remember","tool_calls":null,"type":"tool"},{"content":"Saved: ACME prefers email.","name":null,"tool_calls":null,"type":"ai"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0.000009,"total_tokens":41,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-24/repro.sh`

### Finding 25

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit9).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
#1 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-2
fault crash(#1@commit9)
```

Delivery order: send #0, send, kill the service after its commit 9, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
- memories x1 {"kind":"note","project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","scope":"default","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","text":"Customer ACME prefers to be contacted by email."}
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"decision":"approve","messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] approve","name":null,"tool_calls":null,"type":"human"},{"content":"","name":null,"tool_calls":[{"args":{"text":"Customer ACME prefers to be contacted by email."},"display_name":"remember","name":"remember"}],"type":"ai"},{"content":"Saved to long-term memory.","name":"remember","tool_calls":null,"type":"tool"},{"content":"Saved: ACME prefers email.","name":null,"tool_calls":null,"type":"ai"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0.000009,"total_tokens":41,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-25/repro.sh`

### Finding 26

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit10).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
#1 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-2
fault crash(#1@commit10)
```

Delivery order: send #0, send, kill the service after its commit 10, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
- memories x1 {"kind":"note","project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","scope":"default","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","text":"Customer ACME prefers to be contacted by email."}
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"decision":"approve","messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] approve","name":null,"tool_calls":null,"type":"human"},{"content":"","name":null,"tool_calls":[{"args":{"text":"Customer ACME prefers to be contacted by email."},"display_name":"remember","name":"remember"}],"type":"ai"},{"content":"Saved to long-term memory.","name":"remember","tool_calls":null,"type":"tool"},{"content":"Saved: ACME prefers email.","name":null,"tool_calls":null,"type":"ai"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0.000009,"total_tokens":41,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-26/repro.sh`

### Finding 27

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit11).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
#1 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-2
fault crash(#1@commit11)
```

Delivery order: send #0, send, kill the service after its commit 11, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"decision":"approve","messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] approve","name":null,"tool_calls":null,"type":"human"},{"content":"","name":null,"tool_calls":[{"args":{"text":"Customer ACME prefers to be contacted by email."},"display_name":"remember","name":"remember"}],"type":"ai"},{"content":"Saved to long-term memory.","name":"remember","tool_calls":null,"type":"tool"},{"content":"Saved: ACME prefers email.","name":null,"tool_calls":null,"type":"ai"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0.000009,"total_tokens":41,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-27/repro.sh`

### Finding 28

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit12).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
#1 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-2
fault crash(#1@commit12)
```

Delivery order: send #0, send, kill the service after its commit 12, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"decision":"approve","messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] approve","name":null,"tool_calls":null,"type":"human"},{"content":"","name":null,"tool_calls":[{"args":{"text":"Customer ACME prefers to be contacted by email."},"display_name":"remember","name":"remember"}],"type":"ai"},{"content":"Saved to long-term memory.","name":"remember","tool_calls":null,"type":"tool"},{"content":"Saved: ACME prefers email.","name":null,"tool_calls":null,"type":"ai"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0.000009,"total_tokens":41,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-28/repro.sh`

### Finding 29

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit13).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
#1 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-2
fault crash(#1@commit13)
```

Delivery order: send #0, send, kill the service after its commit 13, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"decision":"approve","messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] approve","name":null,"tool_calls":null,"type":"human"},{"content":"","name":null,"tool_calls":[{"args":{"text":"Customer ACME prefers to be contacted by email."},"display_name":"remember","name":"remember"}],"type":"ai"},{"content":"Saved to long-term memory.","name":"remember","tool_calls":null,"type":"tool"},{"content":"Saved: ACME prefers email.","name":null,"tool_calls":null,"type":"ai"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0.000009,"total_tokens":41,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-29/repro.sh`

### Finding 30

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit14).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
#1 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-2
fault crash(#1@commit14)
```

Delivery order: send #0, send, kill the service after its commit 14, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"decision":"approve","messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] approve","name":null,"tool_calls":null,"type":"human"},{"content":"","name":null,"tool_calls":[{"args":{"text":"Customer ACME prefers to be contacted by email."},"display_name":"remember","name":"remember"}],"type":"ai"},{"content":"Saved to long-term memory.","name":"remember","tool_calls":null,"type":"tool"},{"content":"Saved: ACME prefers email.","name":null,"tool_calls":null,"type":"ai"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0.000009,"total_tokens":41,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-30/repro.sh`

### Finding 31

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit15).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
#1 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-2
fault crash(#1@commit15)
```

Delivery order: send #0, send, kill the service after its commit 15, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"decision":"approve","messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] approve","name":null,"tool_calls":null,"type":"human"},{"content":"","name":null,"tool_calls":[{"args":{"text":"Customer ACME prefers to be contacted by email."},"display_name":"remember","name":"remember"}],"type":"ai"},{"content":"Saved to long-term memory.","name":"remember","tool_calls":null,"type":"tool"},{"content":"Saved: ACME prefers email.","name":null,"tool_calls":null,"type":"ai"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0.000009,"total_tokens":41,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-31/repro.sh`

### Finding 32

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit16).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
#1 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-2
fault crash(#1@commit16)
```

Delivery order: send #0, send, kill the service after its commit 16, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"decision":"approve","messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] approve","name":null,"tool_calls":null,"type":"human"},{"content":"","name":null,"tool_calls":[{"args":{"text":"Customer ACME prefers to be contacted by email."},"display_name":"remember","name":"remember"}],"type":"ai"},{"content":"Saved to long-term memory.","name":"remember","tool_calls":null,"type":"tool"},{"content":"Saved: ACME prefers email.","name":null,"tool_calls":null,"type":"ai"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0.000009,"total_tokens":41,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-32/repro.sh`

### Finding 33

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit17).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
#1 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-2
fault crash(#1@commit17)
```

Delivery order: send #0, send, kill the service after its commit 17, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"decision":"approve","messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] approve","name":null,"tool_calls":null,"type":"human"},{"content":"","name":null,"tool_calls":[{"args":{"text":"Customer ACME prefers to be contacted by email."},"display_name":"remember","name":"remember"}],"type":"ai"},{"content":"Saved to long-term memory.","name":"remember","tool_calls":null,"type":"tool"},{"content":"Saved: ACME prefers email.","name":null,"tool_calls":null,"type":"ai"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0.000009,"total_tokens":41,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-33/repro.sh`

### Finding 34

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit18).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
#1 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-2
fault crash(#1@commit18)
```

Delivery order: send #0, send, kill the service after its commit 18, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"decision":"approve","messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] approve","name":null,"tool_calls":null,"type":"human"},{"content":"","name":null,"tool_calls":[{"args":{"text":"Customer ACME prefers to be contacted by email."},"display_name":"remember","name":"remember"}],"type":"ai"},{"content":"Saved to long-term memory.","name":"remember","tool_calls":null,"type":"tool"},{"content":"Saved: ACME prefers email.","name":null,"tool_calls":null,"type":"ai"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0.000009,"total_tokens":41,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-34/repro.sh`

### Finding 35

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit19).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
#1 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-2
fault crash(#1@commit19)
```

Delivery order: send #0, send, kill the service after its commit 19, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer ACME prefers email.","role":"user"}]},"output":{"decision":"approve","messages":[{"content":"Remember that customer ACME prefers email.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] approve","name":null,"tool_calls":null,"type":"human"},{"content":"","name":null,"tool_calls":[{"args":{"text":"Customer ACME prefers to be contacted by email."},"display_name":"remember","name":"remember"}],"type":"ai"},{"content":"Saved to long-term memory.","name":"remember","tool_calls":null,"type":"tool"},{"content":"Saved: ACME prefers email.","name":null,"tool_calls":null,"type":"ai"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0.000009,"total_tokens":41,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-35/repro.sh`

### Finding 36

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#2@commit10).

```
#0 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=acme-1
#2 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=globex-1
fault crash(#2@commit10)
```

Delivery order: send #0, send, kill the service after its commit 10, restart, resend #2

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"interrupted","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-36/repro.sh`

### Finding 37

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit2).

```
#2 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=globex-1
#3 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=globex-2
fault crash(#3@commit2)
```

Delivery order: send #2, send, kill the service after its commit 2, restart, resend #3

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":{"decision":"reject","messages":[{"content":"Remember that customer Globex wants invoices in EUR.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] reject","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-37/repro.sh`

### Finding 38

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit3).

```
#2 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=globex-1
#3 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=globex-2
fault crash(#3@commit3)
```

Delivery order: send #2, send, kill the service after its commit 3, restart, resend #3

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":{"decision":"reject","messages":[{"content":"Remember that customer Globex wants invoices in EUR.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] reject","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-38/repro.sh`

### Finding 39

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit4).

```
#2 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=globex-1
#3 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=globex-2
fault crash(#3@commit4)
```

Delivery order: send #2, send, kill the service after its commit 4, restart, resend #3

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":{"decision":"reject","messages":[{"content":"Remember that customer Globex wants invoices in EUR.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] reject","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-39/repro.sh`

### Finding 40

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit5).

```
#2 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=globex-1
#3 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=globex-2
fault crash(#3@commit5)
```

Delivery order: send #2, send, kill the service after its commit 5, restart, resend #3

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":{"decision":"reject","messages":[{"content":"Remember that customer Globex wants invoices in EUR.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] reject","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-40/repro.sh`

### Finding 41

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit6).

```
#2 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=globex-1
#3 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=globex-2
fault crash(#3@commit6)
```

Delivery order: send #2, send, kill the service after its commit 6, restart, resend #3

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":{"decision":"reject","messages":[{"content":"Remember that customer Globex wants invoices in EUR.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] reject","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-41/repro.sh`

### Finding 42

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit7).

```
#2 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=globex-1
#3 POST /v1/projects/74ed7105-e6a4-47b7-a571-172b4eb7af47/run key=globex-2
fault crash(#3@commit7)
```

Delivery order: send #2, send, kill the service after its commit 7, restart, resend #3

State diff (+ only in the faulted run, - only in the in-order run):

```
+ runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"running","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
- runs x1 {"error":null,"input":{"messages":[{"content":"Remember that customer Globex wants invoices in EUR.","role":"user"}]},"output":{"decision":"reject","messages":[{"content":"Remember that customer Globex wants invoices in EUR.","name":null,"tool_calls":null,"type":"human"},{"content":"[human decision] reject","name":null,"tool_calls":null,"type":"human"}]},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","source":"api","status":"done","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","thread_id":"<ref:threads>","total_cost_usd":0,"total_tokens":0,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"}
+ threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
- threads x1 {"metadata":{},"project_id":"74ed7105-e6a4-47b7-a571-172b4eb7af47","status":"active","tenant_id":"a6a2db13-c587-435e-8cf5-8d465b144df6","title":"Playground run","user_external_id":null,"workflow_id":"db5a4284-e08c-4a4c-a4f0-712383cd6444"} (+1 referencing rows)
```

Repro: `finding-42/repro.sh`

Hunt time: 4548 s.
