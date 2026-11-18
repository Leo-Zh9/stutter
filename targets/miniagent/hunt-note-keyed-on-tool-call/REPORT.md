# stutter hunt: miniagent

Generated 2026-11-18T10:23:26Z. Trace: 6 requests. Pre-flight: every recorded webhook delivery got 2xx.

| Fault type | Schedules | Fired | Not fired | Diverged | Nondetermined | Changed the final state | No-op passes |
|---|---|---|---|---|---|---|---|
| concurrent | 6 | 6 | 0 | 0 | 0 | 0 | 0 |
| crash | 61 | 61 | 0 | 0 | 0 | 1 | 0 |
| lost | 6 | 6 | 0 | 0 | 0 | 0 | 0 |

No-op reorder pairs (dropping both events leaves the state unchanged, so their pass says nothing about ordering): 0.

## Effects: 1 distinct end-state changes

1. Findings [1]:

```
+ order_notes {"dedupe_key":"<model-tool-call-id>","note":"delayed in transit","order_number":1002}
```

## Findings: 1 distinct minimal reproductions

### Finding 1

Reproduced in 4 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: false (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#2@commit4).

```
#0 POST /chat {"thread_id":"t-ada","message_id":"m-1","text":"Hi, please refund order 1001"}
#1 POST /threads/t-ada/approve {"decision":"approve"}
#2 POST /chat {"thread_id":"t-grace","message_id":"m-2","text":"Please note that order 1002 was delayed"}
#5 POST /threads/t-linus/approve {"decision":"reject"}
fault crash(#2@commit4)
```

Delivery order: send #0, send #1, send, kill the service after its commit 4, restart, resend #2, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ order_notes x1 {"dedupe_key":"<model-tool-call-id>","note":"delayed in transit","order_number":1002}
```

Repro: `finding-1/repro.sh`

Hunt time: 153 s.
