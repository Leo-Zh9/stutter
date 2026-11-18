# stutter hunt: miniagent

Generated 2026-11-18T10:16:41Z. Trace: 6 requests. Pre-flight: every recorded webhook delivery got 2xx.

| Fault type | Schedules | Fired | Not fired | Diverged | Nondetermined | Changed the final state | No-op passes |
|---|---|---|---|---|---|---|---|
| concurrent | 6 | 5 | 1 | 0 | 0 | 2 | 0 |
| crash | 53 | 53 | 0 | 0 | 0 | 0 | 0 |
| lost | 6 | 6 | 0 | 0 | 0 | 0 | 0 |

No-op reorder pairs (dropping both events leaves the state unchanged, so their pass says nothing about ordering): 0.

## Effects: 2 distinct end-state changes

1. Findings [1]:

```
+ approvals {"amount_cents":4200,"order_number":1001,"status":"pending","thread_id":"t-ada"}
```

2. Findings [2]:

```
+ approvals {"amount_cents":2000,"order_number":1003,"status":"pending","thread_id":"t-linus"}
```

## Findings: 2 distinct minimal reproductions

### Finding 1

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 3 of 10 times; 1-minimal: true (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: concurrent(#0).

```
#0 POST /chat {"thread_id":"t-ada","message_id":"m-1","text":"Hi, please refund order 1001"}
fault concurrent(#0)
```

Delivery order: send-concurrently #0

State diff (+ only in the faulted run, - only in the in-order run):

```
+ approvals x1 {"amount_cents":4200,"order_number":1001,"status":"pending","thread_id":"t-ada"}
```

Repro: `finding-1/repro.sh`

### Finding 2

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 5 of 10 times; 1-minimal: true (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: concurrent(#4).

```
#4 POST /chat {"thread_id":"t-linus","message_id":"m-4","text":"OK, refund order 1003 for the broken part"}
fault concurrent(#4)
```

Delivery order: send-concurrently #4

State diff (+ only in the faulted run, - only in the in-order run):

```
+ approvals x1 {"amount_cents":2000,"order_number":1003,"status":"pending","thread_id":"t-linus"}
```

Repro: `finding-2/repro.sh`

Hunt time: 167 s.
