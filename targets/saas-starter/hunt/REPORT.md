# stutter hunt: nextjs/saas-starter @ 6e33e58

Generated 2026-11-14T12:06:45Z. Trace: 16 requests. Pre-flight: every recorded webhook delivery got 2xx.

| Fault type | Schedules | Fired | Not fired | Changed the final state | No-op passes |
|---|---|---|---|---|---|
| concurrent | 16 | 9 | 7 | 0 | 0 |
| crash | 9 | 9 | 0 | 0 | 0 |
| lost | 16 | 16 | 0 | 0 | 0 |
| reorder | 23 | 23 | 0 | 7 | 12 |

No-op reorder pairs (dropping both events leaves the state unchanged, so their pass says nothing about ordering): 12.

## Effects: 3 distinct end-state changes

1. Findings [1 2 5 6]:

```
+ teams {"name":"Team 1","plan_name":null,"stripe_customer_id":"cus_team1","stripe_product_id":"prod_basic","stripe_subscription_id":"sub_one_a","subscription_status":"active"}
- teams {"name":"Team 1","plan_name":null,"stripe_customer_id":"cus_team1","stripe_product_id":null,"stripe_subscription_id":null,"subscription_status":"canceled"}
```

2. Findings [3]:

```
+ teams {"name":"Team 2","plan_name":null,"stripe_customer_id":"cus_team2","stripe_product_id":"prod_basic","stripe_subscription_id":"sub_two_a","subscription_status":"active"}
- teams {"name":"Team 2","plan_name":null,"stripe_customer_id":"cus_team2","stripe_product_id":"prod_plus","stripe_subscription_id":"sub_two_b","subscription_status":"active"}
```

3. Findings [4 7]:

```
+ teams {"name":"Team 2","plan_name":null,"stripe_customer_id":"cus_team2","stripe_product_id":null,"stripe_subscription_id":null,"subscription_status":"canceled"}
- teams {"name":"Team 2","plan_name":null,"stripe_customer_id":"cus_team2","stripe_product_id":"prod_plus","stripe_subscription_id":"sub_two_b","subscription_status":"active"}
```

## Findings: 7 distinct minimal reproductions

### Finding 1

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#1,#4).

```
#1 webhook customer.subscription.updated evt_0001 customer=cus_team1 id=sub_one_a status=active
#4 webhook customer.subscription.deleted evt_0004 customer=cus_team1 id=sub_one_a status=canceled
fault reorder(#1,#4)
```

Delivery order: send #4, send #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ teams x1 {"name":"Team 1","plan_name":null,"stripe_customer_id":"cus_team1","stripe_product_id":"prod_basic","stripe_subscription_id":"sub_one_a","subscription_status":"active"}
- teams x1 {"name":"Team 1","plan_name":null,"stripe_customer_id":"cus_team1","stripe_product_id":null,"stripe_subscription_id":null,"subscription_status":"canceled"}
```

Repro: `finding-1/repro.sh`

### Finding 2

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#3,#4).

```
#3 webhook customer.subscription.updated evt_0003 customer=cus_team1 id=sub_one_a status=active
#4 webhook customer.subscription.deleted evt_0004 customer=cus_team1 id=sub_one_a status=canceled
fault reorder(#3,#4)
```

Delivery order: send #4, send #3

State diff (+ only in the faulted run, - only in the in-order run):

```
+ teams x1 {"name":"Team 1","plan_name":null,"stripe_customer_id":"cus_team1","stripe_product_id":"prod_basic","stripe_subscription_id":"sub_one_a","subscription_status":"active"}
- teams x1 {"name":"Team 1","plan_name":null,"stripe_customer_id":"cus_team1","stripe_product_id":null,"stripe_subscription_id":null,"subscription_status":"canceled"}
```

Repro: `finding-2/repro.sh`

### Finding 3

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#6,#9).

```
#6 webhook customer.subscription.updated evt_0006 customer=cus_team2 id=sub_two_a status=active
#9 webhook customer.subscription.updated evt_0009 customer=cus_team2 id=sub_two_b status=active
fault reorder(#6,#9)
```

Delivery order: send #9, send #6

State diff (+ only in the faulted run, - only in the in-order run):

```
+ teams x1 {"name":"Team 2","plan_name":null,"stripe_customer_id":"cus_team2","stripe_product_id":"prod_basic","stripe_subscription_id":"sub_two_a","subscription_status":"active"}
- teams x1 {"name":"Team 2","plan_name":null,"stripe_customer_id":"cus_team2","stripe_product_id":"prod_plus","stripe_subscription_id":"sub_two_b","subscription_status":"active"}
```

Repro: `finding-3/repro.sh`

### Finding 4

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#7,#9).

```
#7 webhook customer.subscription.deleted evt_0007 customer=cus_team2 id=sub_two_a status=canceled
#9 webhook customer.subscription.updated evt_0009 customer=cus_team2 id=sub_two_b status=active
fault reorder(#7,#9)
```

Delivery order: send #9, send #7

State diff (+ only in the faulted run, - only in the in-order run):

```
+ teams x1 {"name":"Team 2","plan_name":null,"stripe_customer_id":"cus_team2","stripe_product_id":null,"stripe_subscription_id":null,"subscription_status":"canceled"}
- teams x1 {"name":"Team 2","plan_name":null,"stripe_customer_id":"cus_team2","stripe_product_id":"prod_plus","stripe_subscription_id":"sub_two_b","subscription_status":"active"}
```

Repro: `finding-4/repro.sh`

### Finding 5

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#0,#4).

```
#0 webhook customer.subscription.created evt_0000 customer=cus_team1 id=sub_one_a status=trialing
#1 webhook customer.subscription.updated evt_0001 customer=cus_team1 id=sub_one_a status=active
#4 webhook customer.subscription.deleted evt_0004 customer=cus_team1 id=sub_one_a status=canceled
fault reorder(#0,#4)
```

Delivery order: send #4, send #1, send #0

State diff (+ only in the faulted run, - only in the in-order run):

```
+ teams x1 {"name":"Team 1","plan_name":null,"stripe_customer_id":"cus_team1","stripe_product_id":"prod_basic","stripe_subscription_id":"sub_one_a","subscription_status":"active"}
- teams x1 {"name":"Team 1","plan_name":null,"stripe_customer_id":"cus_team1","stripe_product_id":null,"stripe_subscription_id":null,"subscription_status":"canceled"}
```

Repro: `finding-5/repro.sh`

### Finding 6

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#2,#4).

```
#2 webhook customer.subscription.updated evt_0002 customer=cus_team1 id=sub_one_a status=past_due
#3 webhook customer.subscription.updated evt_0003 customer=cus_team1 id=sub_one_a status=active
#4 webhook customer.subscription.deleted evt_0004 customer=cus_team1 id=sub_one_a status=canceled
fault reorder(#2,#4)
```

Delivery order: send #4, send #3, send #2

State diff (+ only in the faulted run, - only in the in-order run):

```
+ teams x1 {"name":"Team 1","plan_name":null,"stripe_customer_id":"cus_team1","stripe_product_id":"prod_basic","stripe_subscription_id":"sub_one_a","subscription_status":"active"}
- teams x1 {"name":"Team 1","plan_name":null,"stripe_customer_id":"cus_team1","stripe_product_id":null,"stripe_subscription_id":null,"subscription_status":"canceled"}
```

Repro: `finding-6/repro.sh`

### Finding 7

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
+ teams x1 {"name":"Team 2","plan_name":null,"stripe_customer_id":"cus_team2","stripe_product_id":null,"stripe_subscription_id":null,"subscription_status":"canceled"}
- teams x1 {"name":"Team 2","plan_name":null,"stripe_customer_id":"cus_team2","stripe_product_id":"prod_plus","stripe_subscription_id":"sub_two_b","subscription_status":"active"}
```

Repro: `finding-7/repro.sh`

Hunt time: 104 s.
