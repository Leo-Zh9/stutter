# stutter hunt: miniagent

Generated 2026-11-18T10:26:01Z. Trace: 6 requests. Pre-flight: every recorded webhook delivery got 2xx.

| Fault type | Schedules | Fired | Not fired | Diverged | Nondetermined | Changed the final state | No-op passes |
|---|---|---|---|---|---|---|---|
| concurrent | 6 | 6 | 0 | 0 | 0 | 2 | 0 |
| crash | 63 | 60 | 0 | 0 | 3 | 10 | 0 |
| lost | 6 | 6 | 0 | 0 | 0 | 0 | 0 |

No-op reorder pairs (dropping both events leaves the state unchanged, so their pass says nothing about ordering): 0.

## Effects: 5 distinct end-state changes

1. Findings [1 2 8]:

```
+ approvals {"amount_cents":4200,"order_number":1001,"status":"pending","thread_id":"t-ada"}
```

2. Findings [3 4 5]:

```
+ approvals {"amount_cents":2000,"order_number":1003,"status":"pending","thread_id":"t-linus"}
```

3. Findings [6 10]:

```
+ approvals {"amount_cents":4200,"order_number":1001,"status":"approved","thread_id":"t-ada"} (+1 referencing rows)
+ outbox {"kind":"refund_email","order_number":1001}
+ refunds {"amount_cents":4200,"approval_id":"<ref:approvals>","order_number":1001}
```

4. Findings [7 11 12]:

```
+ approvals {"amount_cents":2000,"order_number":1003,"status":"rejected","thread_id":"t-linus"}
```

5. Findings [9]:

```
+ approvals {"amount_cents":4200,"order_number":1001,"status":"approved","thread_id":"t-ada"}
```

## Findings: 12 distinct minimal reproductions

### Finding 1

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#0@commit10).

```
#0 POST /chat {"thread_id":"t-ada","message_id":"m-1","text":"Hi, please refund order 1001"}
fault crash(#0@commit10)
```

Delivery order: send, kill the service after its commit 10, restart, resend #0

State diff (+ only in the faulted run, - only in the in-order run):

```
+ approvals x1 {"amount_cents":4200,"order_number":1001,"status":"pending","thread_id":"t-ada"}
```

Repro: `finding-1/repro.sh`

### Finding 2

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#0@commit11).

```
#0 POST /chat {"thread_id":"t-ada","message_id":"m-1","text":"Hi, please refund order 1001"}
fault crash(#0@commit11)
```

Delivery order: send, kill the service after its commit 11, restart, resend #0

State diff (+ only in the faulted run, - only in the in-order run):

```
+ approvals x1 {"amount_cents":4200,"order_number":1001,"status":"pending","thread_id":"t-ada"}
```

Repro: `finding-2/repro.sh`

### Finding 3

Reproduced in 7 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#4@commit9).

```
#4 POST /chat {"thread_id":"t-linus","message_id":"m-4","text":"OK, refund order 1003 for the broken part"}
fault crash(#4@commit9)
```

Delivery order: send, kill the service after its commit 9, restart, resend #4

State diff (+ only in the faulted run, - only in the in-order run):

```
+ approvals x1 {"amount_cents":2000,"order_number":1003,"status":"pending","thread_id":"t-linus"}
```

Repro: `finding-3/repro.sh`

### Finding 4

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#4@commit10).

```
#4 POST /chat {"thread_id":"t-linus","message_id":"m-4","text":"OK, refund order 1003 for the broken part"}
fault crash(#4@commit10)
```

Delivery order: send, kill the service after its commit 10, restart, resend #4

State diff (+ only in the faulted run, - only in the in-order run):

```
+ approvals x1 {"amount_cents":2000,"order_number":1003,"status":"pending","thread_id":"t-linus"}
```

Repro: `finding-4/repro.sh`

### Finding 5

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#4@commit11).

```
#4 POST /chat {"thread_id":"t-linus","message_id":"m-4","text":"OK, refund order 1003 for the broken part"}
fault crash(#4@commit11)
```

Delivery order: send, kill the service after its commit 11, restart, resend #4

State diff (+ only in the faulted run, - only in the in-order run):

```
+ approvals x1 {"amount_cents":2000,"order_number":1003,"status":"pending","thread_id":"t-linus"}
```

Repro: `finding-5/repro.sh`

### Finding 6

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: concurrent(#1).

```
#0 POST /chat {"thread_id":"t-ada","message_id":"m-1","text":"Hi, please refund order 1001"}
#1 POST /threads/t-ada/approve {"decision":"approve"}
fault concurrent(#1)
```

Delivery order: send #0, send-concurrently #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ approvals x1 {"amount_cents":4200,"order_number":1001,"status":"approved","thread_id":"t-ada"} (+1 referencing rows)
+ outbox x1 {"kind":"refund_email","order_number":1001}
+ refunds x1 {"amount_cents":4200,"approval_id":"<ref:approvals>","order_number":1001}
```

Repro: `finding-6/repro.sh`

### Finding 7

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: concurrent(#5).

```
#4 POST /chat {"thread_id":"t-linus","message_id":"m-4","text":"OK, refund order 1003 for the broken part"}
#5 POST /threads/t-linus/approve {"decision":"reject"}
fault concurrent(#5)
```

Delivery order: send #4, send-concurrently #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ approvals x1 {"amount_cents":2000,"order_number":1003,"status":"rejected","thread_id":"t-linus"}
```

Repro: `finding-7/repro.sh`

### Finding 8

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: false (0 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#0@commit9).

```
#0 POST /chat {"thread_id":"t-ada","message_id":"m-1","text":"Hi, please refund order 1001"}
#5 POST /threads/t-linus/approve {"decision":"reject"}
fault crash(#0@commit9)
```

Delivery order: send, kill the service after its commit 9, restart, resend #0, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ approvals x1 {"amount_cents":4200,"order_number":1001,"status":"pending","thread_id":"t-ada"}
```

Repro: `finding-8/repro.sh`

### Finding 9

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit2).

```
#0 POST /chat {"thread_id":"t-ada","message_id":"m-1","text":"Hi, please refund order 1001"}
#1 POST /threads/t-ada/approve {"decision":"approve"}
fault crash(#1@commit2)
```

Delivery order: send #0, send, kill the service after its commit 2, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ approvals x1 {"amount_cents":4200,"order_number":1001,"status":"approved","thread_id":"t-ada"}
```

Repro: `finding-9/repro.sh`

### Finding 10

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit3).

```
#0 POST /chat {"thread_id":"t-ada","message_id":"m-1","text":"Hi, please refund order 1001"}
#1 POST /threads/t-ada/approve {"decision":"approve"}
fault crash(#1@commit3)
```

Delivery order: send #0, send, kill the service after its commit 3, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ approvals x1 {"amount_cents":4200,"order_number":1001,"status":"approved","thread_id":"t-ada"} (+1 referencing rows)
+ outbox x1 {"kind":"refund_email","order_number":1001}
+ refunds x1 {"amount_cents":4200,"approval_id":"<ref:approvals>","order_number":1001}
```

Repro: `finding-10/repro.sh`

### Finding 11

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#5@commit2).

```
#4 POST /chat {"thread_id":"t-linus","message_id":"m-4","text":"OK, refund order 1003 for the broken part"}
#5 POST /threads/t-linus/approve {"decision":"reject"}
fault crash(#5@commit2)
```

Delivery order: send #4, send, kill the service after its commit 2, restart, resend #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ approvals x1 {"amount_cents":2000,"order_number":1003,"status":"rejected","thread_id":"t-linus"}
```

Repro: `finding-11/repro.sh`

### Finding 12

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#5@commit3).

```
#4 POST /chat {"thread_id":"t-linus","message_id":"m-4","text":"OK, refund order 1003 for the broken part"}
#5 POST /threads/t-linus/approve {"decision":"reject"}
fault crash(#5@commit3)
```

Delivery order: send #4, send, kill the service after its commit 3, restart, resend #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ approvals x1 {"amount_cents":2000,"order_number":1003,"status":"rejected","thread_id":"t-linus"}
```

Repro: `finding-12/repro.sh`

Hunt time: 666 s.
