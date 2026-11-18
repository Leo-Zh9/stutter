# stutter hunt: miniagent

Generated 2026-11-18T10:13:28Z. Trace: 6 requests. Pre-flight: every recorded webhook delivery got 2xx.

| Fault type | Schedules | Fired | Not fired | Diverged | Nondetermined | Changed the final state | No-op passes |
|---|---|---|---|---|---|---|---|
| concurrent | 6 | 6 | 0 | 0 | 0 | 1 | 0 |
| crash | 61 | 61 | 0 | 0 | 0 | 0 | 0 |
| lost | 6 | 6 | 0 | 0 | 0 | 0 | 0 |

No-op reorder pairs (dropping both events leaves the state unchanged, so their pass says nothing about ordering): 0.

## Effects: 1 distinct end-state changes

1. Findings [1]:

```
+ approvals {"amount_cents":4200,"order_number":1001,"status":"approved","thread_id":"t-ada"} (+2 referencing rows)
- approvals {"amount_cents":4200,"order_number":1001,"status":"approved","thread_id":"t-ada"} (+1 referencing rows)
+ outbox {"kind":"refund_email","order_number":1001}
+ refunds {"amount_cents":4200,"approval_id":"<ref:approvals>","order_number":1001}
```

## Findings: 1 distinct minimal reproductions

### Finding 1

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: concurrent(#1).

```
#0 POST /chat {"thread_id":"t-ada","message_id":"m-1","text":"Hi, please refund order 1001"}
#1 POST /threads/t-ada/approve {"decision":"approve"}
fault concurrent(#1)
```

Delivery order: send #0, send-concurrently #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ approvals x1 {"amount_cents":4200,"order_number":1001,"status":"approved","thread_id":"t-ada"} (+2 referencing rows)
- approvals x1 {"amount_cents":4200,"order_number":1001,"status":"approved","thread_id":"t-ada"} (+1 referencing rows)
+ outbox x1 {"kind":"refund_email","order_number":1001}
+ refunds x1 {"amount_cents":4200,"approval_id":"<ref:approvals>","order_number":1001}
```

Repro: `finding-1/repro.sh`

Hunt time: 161 s.
