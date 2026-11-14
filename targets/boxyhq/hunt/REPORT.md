# stutter hunt: boxyhq/saas-starter-kit @ abc9b68

Generated 2026-11-14T12:09:40Z. Trace: 16 requests. Pre-flight: every recorded webhook delivery got 2xx.

| Fault type | Schedules | Fired | Not fired | Changed the final state | No-op passes |
|---|---|---|---|---|---|
| concurrent | 16 | 16 | 0 | 0 | 0 |
| crash | 16 | 16 | 0 | 0 | 0 |
| lost | 16 | 16 | 0 | 0 | 0 |
| reorder | 23 | 23 | 0 | 11 | 7 |

No-op reorder pairs (dropping both events leaves the state unchanged, so their pass says nothing about ordering): 7.

## Effects: 3 distinct end-state changes

1. Findings [1 2 3 4]:

```
+ Subscription {"active":true,"customerId":"cus_team1","id":"sub_one_a","priceId":"price_prod_basic"}
```

2. Findings [5 6 8 9 10 11]:

```
+ Subscription {"active":true,"customerId":"cus_team2","id":"sub_two_a","priceId":"price_prod_basic"}
```

3. Findings [7]:

```
+ Subscription {"active":true,"customerId":"cus_team4","id":"sub_four_a","priceId":"price_prod_basic"}
- Subscription {"active":false,"customerId":"cus_team4","id":"sub_four_a","priceId":"price_prod_basic"}
```

## Findings: 11 distinct minimal reproductions

### Finding 1

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#0,#4).

```
#0 webhook customer.subscription.created evt_0000 customer=cus_team1 id=sub_one_a status=trialing
#4 webhook customer.subscription.deleted evt_0004 customer=cus_team1 id=sub_one_a status=canceled
fault reorder(#0,#4)
```

Delivery order: send #4, send #0

State diff (+ only in the faulted run, - only in the in-order run):

```
+ Subscription x1 {"active":true,"customerId":"cus_team1","id":"sub_one_a","priceId":"price_prod_basic"}
```

Repro: `finding-1/repro.sh`

### Finding 2

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#1,#4).

```
#1 webhook customer.subscription.updated evt_0001 customer=cus_team1 id=sub_one_a status=active
#4 webhook customer.subscription.deleted evt_0004 customer=cus_team1 id=sub_one_a status=canceled
fault reorder(#1,#4)
```

Delivery order: send #4, send #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ Subscription x1 {"active":true,"customerId":"cus_team1","id":"sub_one_a","priceId":"price_prod_basic"}
```

Repro: `finding-2/repro.sh`

### Finding 3

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#2,#4).

```
#2 webhook customer.subscription.updated evt_0002 customer=cus_team1 id=sub_one_a status=past_due
#4 webhook customer.subscription.deleted evt_0004 customer=cus_team1 id=sub_one_a status=canceled
fault reorder(#2,#4)
```

Delivery order: send #4, send #2

State diff (+ only in the faulted run, - only in the in-order run):

```
+ Subscription x1 {"active":true,"customerId":"cus_team1","id":"sub_one_a","priceId":"price_prod_basic"}
```

Repro: `finding-3/repro.sh`

### Finding 4

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#3,#4).

```
#3 webhook customer.subscription.updated evt_0003 customer=cus_team1 id=sub_one_a status=active
#4 webhook customer.subscription.deleted evt_0004 customer=cus_team1 id=sub_one_a status=canceled
fault reorder(#3,#4)
```

Delivery order: send #4, send #3

State diff (+ only in the faulted run, - only in the in-order run):

```
+ Subscription x1 {"active":true,"customerId":"cus_team1","id":"sub_one_a","priceId":"price_prod_basic"}
```

Repro: `finding-4/repro.sh`

### Finding 5

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#5,#7).

```
#5 webhook customer.subscription.created evt_0005 customer=cus_team2 id=sub_two_a status=active
#7 webhook customer.subscription.deleted evt_0007 customer=cus_team2 id=sub_two_a status=canceled
fault reorder(#5,#7)
```

Delivery order: send #7, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ Subscription x1 {"active":true,"customerId":"cus_team2","id":"sub_two_a","priceId":"price_prod_basic"}
```

Repro: `finding-5/repro.sh`

### Finding 6

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#6,#7).

```
#6 webhook customer.subscription.updated evt_0006 customer=cus_team2 id=sub_two_a status=active
#7 webhook customer.subscription.deleted evt_0007 customer=cus_team2 id=sub_two_a status=canceled
fault reorder(#6,#7)
```

Delivery order: send #7, send #6

State diff (+ only in the faulted run, - only in the in-order run):

```
+ Subscription x1 {"active":true,"customerId":"cus_team2","id":"sub_two_a","priceId":"price_prod_basic"}
```

Repro: `finding-6/repro.sh`

### Finding 7

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#12,#13).

```
#12 webhook customer.subscription.created evt_0012 customer=cus_team4 id=sub_four_a status=trialing
#13 webhook customer.subscription.updated evt_0013 customer=cus_team4 id=sub_four_a status=unpaid
fault reorder(#12,#13)
```

Delivery order: send #13, send #12

State diff (+ only in the faulted run, - only in the in-order run):

```
+ Subscription x1 {"active":true,"customerId":"cus_team4","id":"sub_four_a","priceId":"price_prod_basic"}
- Subscription x1 {"active":false,"customerId":"cus_team4","id":"sub_four_a","priceId":"price_prod_basic"}
```

Repro: `finding-7/repro.sh`

### Finding 8

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#5,#8).

```
#5 webhook customer.subscription.created evt_0005 customer=cus_team2 id=sub_two_a status=active
#7 webhook customer.subscription.deleted evt_0007 customer=cus_team2 id=sub_two_a status=canceled
#8 webhook customer.subscription.created evt_0008 customer=cus_team2 id=sub_two_b status=active
fault reorder(#5,#8)
```

Delivery order: send #8, send #7, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ Subscription x1 {"active":true,"customerId":"cus_team2","id":"sub_two_a","priceId":"price_prod_basic"}
```

Repro: `finding-8/repro.sh`

### Finding 9

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#5,#9).

```
#5 webhook customer.subscription.created evt_0005 customer=cus_team2 id=sub_two_a status=active
#7 webhook customer.subscription.deleted evt_0007 customer=cus_team2 id=sub_two_a status=canceled
#9 webhook customer.subscription.updated evt_0009 customer=cus_team2 id=sub_two_b status=active
fault reorder(#5,#9)
```

Delivery order: send #9, send #7, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ Subscription x1 {"active":true,"customerId":"cus_team2","id":"sub_two_a","priceId":"price_prod_basic"}
```

Repro: `finding-9/repro.sh`

### Finding 10

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#6,#8).

```
#6 webhook customer.subscription.updated evt_0006 customer=cus_team2 id=sub_two_a status=active
#7 webhook customer.subscription.deleted evt_0007 customer=cus_team2 id=sub_two_a status=canceled
#8 webhook customer.subscription.created evt_0008 customer=cus_team2 id=sub_two_b status=active
fault reorder(#6,#8)
```

Delivery order: send #8, send #7, send #6

State diff (+ only in the faulted run, - only in the in-order run):

```
+ Subscription x1 {"active":true,"customerId":"cus_team2","id":"sub_two_a","priceId":"price_prod_basic"}
```

Repro: `finding-10/repro.sh`

### Finding 11

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#6,#9).

```
#6 webhook customer.subscription.updated evt_0006 customer=cus_team2 id=sub_two_a status=active
#7 webhook customer.subscription.deleted evt_0007 customer=cus_team2 id=sub_two_a status=canceled
#9 webhook customer.subscription.updated evt_0009 customer=cus_team2 id=sub_two_b status=active
fault reorder(#6,#9)
```

Delivery order: send #9, send #7, send #6

State diff (+ only in the faulted run, - only in the in-order run):

```
+ Subscription x1 {"active":true,"customerId":"cus_team2","id":"sub_two_a","priceId":"price_prod_basic"}
```

Repro: `finding-11/repro.sh`

Hunt time: 428 s.
