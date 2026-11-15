# stutter hunt: kriasoft/react-starter-kit @ dbdb0d4

Generated 2026-11-14T23:36:43Z. Trace: 16 requests. Pre-flight: every recorded webhook delivery got 2xx.

| Fault type | Schedules | Fired | Not fired | Changed the final state | No-op passes |
|---|---|---|---|---|---|
| concurrent | 16 | 16 | 0 | 0 | 0 |
| crash | 16 | 16 | 0 | 0 | 0 |
| lost | 16 | 16 | 0 | 0 | 0 |
| reorder | 23 | 23 | 0 | 15 | 3 |

No-op reorder pairs (dropping both events leaves the state unchanged, so their pass says nothing about ordering): 3.

## Effects: 12 distinct end-state changes

1. Findings [1]:

```
+ subscription {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team1","seats":1,"status":"trialing","stripe_customer_id":"cus_team1","stripe_schedule_id":null,"stripe_subscription_id":"sub_one_a"}
- subscription {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team1","seats":1,"status":"canceled","stripe_customer_id":"cus_team1","stripe_schedule_id":null,"stripe_subscription_id":"sub_one_a"}
```

2. Findings [2 11]:

```
+ subscription {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team2","seats":1,"status":"active","stripe_customer_id":"cus_team2","stripe_schedule_id":null,"stripe_subscription_id":"sub_two_a"}
- subscription {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team2","seats":1,"status":"canceled","stripe_customer_id":"cus_team2","stripe_schedule_id":null,"stripe_subscription_id":"sub_two_a"}
```

3. Findings [3]:

```
+ subscription {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team2","seats":1,"status":"active","stripe_customer_id":"cus_team2","stripe_schedule_id":null,"stripe_subscription_id":"sub_two_a"}
- subscription {"billing_interval":"month","cancel_at_period_end":false,"plan":"pro","reference_id":"usr_team2","seats":1,"status":"active","stripe_customer_id":"cus_team2","stripe_schedule_id":null,"stripe_subscription_id":"sub_two_b"}
```

4. Findings [4 13]:

```
+ subscription {"billing_interval":"month","cancel_at_period_end":true,"plan":"starter","reference_id":"usr_team2","seats":1,"status":"active","stripe_customer_id":"cus_team2","stripe_schedule_id":null,"stripe_subscription_id":"sub_two_a"}
- subscription {"billing_interval":"month","cancel_at_period_end":false,"plan":"pro","reference_id":"usr_team2","seats":1,"status":"active","stripe_customer_id":"cus_team2","stripe_schedule_id":null,"stripe_subscription_id":"sub_two_b"}
```

5. Findings [5]:

```
+ subscription {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team3","seats":1,"status":"active","stripe_customer_id":"cus_team3","stripe_schedule_id":null,"stripe_subscription_id":"sub_three_a"}
- subscription {"billing_interval":"month","cancel_at_period_end":false,"plan":"pro","reference_id":"usr_team3","seats":1,"status":"active","stripe_customer_id":"cus_team3","stripe_schedule_id":null,"stripe_subscription_id":"sub_three_a"}
```

6. Findings [6]:

```
+ subscription {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team4","seats":1,"status":"trialing","stripe_customer_id":"cus_team4","stripe_schedule_id":null,"stripe_subscription_id":"sub_four_a"}
- subscription {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team4","seats":1,"status":"unpaid","stripe_customer_id":"cus_team4","stripe_schedule_id":null,"stripe_subscription_id":"sub_four_a"}
```

7. Findings [7]:

```
+ subscription {"billing_interval":"month","cancel_at_period_end":false,"plan":"pro","reference_id":"usr_team5","seats":1,"status":"trialing","stripe_customer_id":"cus_team5","stripe_schedule_id":null,"stripe_subscription_id":"sub_five_a"}
- subscription {"billing_interval":"month","cancel_at_period_end":false,"plan":"pro","reference_id":"usr_team5","seats":1,"status":"active","stripe_customer_id":"cus_team5","stripe_schedule_id":null,"stripe_subscription_id":"sub_five_a"}
```

8. Findings [8 10]:

```
+ subscription {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team1","seats":1,"status":"active","stripe_customer_id":"cus_team1","stripe_schedule_id":null,"stripe_subscription_id":"sub_one_a"}
- subscription {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team1","seats":1,"status":"canceled","stripe_customer_id":"cus_team1","stripe_schedule_id":null,"stripe_subscription_id":"sub_one_a"}
```

9. Findings [9]:

```
+ subscription {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team1","seats":1,"status":"past_due","stripe_customer_id":"cus_team1","stripe_schedule_id":null,"stripe_subscription_id":"sub_one_a"}
- subscription {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team1","seats":1,"status":"canceled","stripe_customer_id":"cus_team1","stripe_schedule_id":null,"stripe_subscription_id":"sub_one_a"}
```

10. Findings [12]:

```
+ subscription {"billing_interval":"month","cancel_at_period_end":true,"plan":"starter","reference_id":"usr_team2","seats":1,"status":"active","stripe_customer_id":"cus_team2","stripe_schedule_id":null,"stripe_subscription_id":"sub_two_a"}
- subscription {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team2","seats":1,"status":"canceled","stripe_customer_id":"cus_team2","stripe_schedule_id":null,"stripe_subscription_id":"sub_two_a"}
```

11. Findings [14]:

```
- subscription {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team2","seats":1,"status":"active","stripe_customer_id":"cus_team2","stripe_schedule_id":null,"stripe_subscription_id":"sub_two_a"}
```

12. Findings [15]:

```
- subscription {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team2","seats":1,"status":"canceled","stripe_customer_id":"cus_team2","stripe_schedule_id":null,"stripe_subscription_id":"sub_two_a"}
```

## Findings: 15 distinct minimal reproductions

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
+ subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team1","seats":1,"status":"trialing","stripe_customer_id":"cus_team1","stripe_schedule_id":null,"stripe_subscription_id":"sub_one_a"}
- subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team1","seats":1,"status":"canceled","stripe_customer_id":"cus_team1","stripe_schedule_id":null,"stripe_subscription_id":"sub_one_a"}
```

Repro: `finding-1/repro.sh`

### Finding 2

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#5,#7).

```
#5 webhook customer.subscription.created evt_0005 customer=cus_team2 id=sub_two_a status=active
#7 webhook customer.subscription.deleted evt_0007 customer=cus_team2 id=sub_two_a status=canceled
fault reorder(#5,#7)
```

Delivery order: send #7, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team2","seats":1,"status":"active","stripe_customer_id":"cus_team2","stripe_schedule_id":null,"stripe_subscription_id":"sub_two_a"}
- subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team2","seats":1,"status":"canceled","stripe_customer_id":"cus_team2","stripe_schedule_id":null,"stripe_subscription_id":"sub_two_a"}
```

Repro: `finding-2/repro.sh`

### Finding 3

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#5,#9).

```
#5 webhook customer.subscription.created evt_0005 customer=cus_team2 id=sub_two_a status=active
#9 webhook customer.subscription.updated evt_0009 customer=cus_team2 id=sub_two_b status=active
fault reorder(#5,#9)
```

Delivery order: send #9, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team2","seats":1,"status":"active","stripe_customer_id":"cus_team2","stripe_schedule_id":null,"stripe_subscription_id":"sub_two_a"}
- subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"pro","reference_id":"usr_team2","seats":1,"status":"active","stripe_customer_id":"cus_team2","stripe_schedule_id":null,"stripe_subscription_id":"sub_two_b"}
```

Repro: `finding-3/repro.sh`

### Finding 4

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#6,#8).

```
#6 webhook customer.subscription.updated evt_0006 customer=cus_team2 id=sub_two_a status=active
#8 webhook customer.subscription.created evt_0008 customer=cus_team2 id=sub_two_b status=active
fault reorder(#6,#8)
```

Delivery order: send #8, send #6

State diff (+ only in the faulted run, - only in the in-order run):

```
+ subscription x1 {"billing_interval":"month","cancel_at_period_end":true,"plan":"starter","reference_id":"usr_team2","seats":1,"status":"active","stripe_customer_id":"cus_team2","stripe_schedule_id":null,"stripe_subscription_id":"sub_two_a"}
- subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"pro","reference_id":"usr_team2","seats":1,"status":"active","stripe_customer_id":"cus_team2","stripe_schedule_id":null,"stripe_subscription_id":"sub_two_b"}
```

Repro: `finding-4/repro.sh`

### Finding 5

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#10,#11).

```
#10 webhook customer.subscription.created evt_0010 customer=cus_team3 id=sub_three_a status=active
#11 webhook customer.subscription.updated evt_0011 customer=cus_team3 id=sub_three_a status=active
fault reorder(#10,#11)
```

Delivery order: send #11, send #10

State diff (+ only in the faulted run, - only in the in-order run):

```
+ subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team3","seats":1,"status":"active","stripe_customer_id":"cus_team3","stripe_schedule_id":null,"stripe_subscription_id":"sub_three_a"}
- subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"pro","reference_id":"usr_team3","seats":1,"status":"active","stripe_customer_id":"cus_team3","stripe_schedule_id":null,"stripe_subscription_id":"sub_three_a"}
```

Repro: `finding-5/repro.sh`

### Finding 6

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#12,#13).

```
#12 webhook customer.subscription.created evt_0012 customer=cus_team4 id=sub_four_a status=trialing
#13 webhook customer.subscription.updated evt_0013 customer=cus_team4 id=sub_four_a status=unpaid
fault reorder(#12,#13)
```

Delivery order: send #13, send #12

State diff (+ only in the faulted run, - only in the in-order run):

```
+ subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team4","seats":1,"status":"trialing","stripe_customer_id":"cus_team4","stripe_schedule_id":null,"stripe_subscription_id":"sub_four_a"}
- subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team4","seats":1,"status":"unpaid","stripe_customer_id":"cus_team4","stripe_schedule_id":null,"stripe_subscription_id":"sub_four_a"}
```

Repro: `finding-6/repro.sh`

### Finding 7

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#14,#15).

```
#14 webhook customer.subscription.created evt_0014 customer=cus_team5 id=sub_five_a status=trialing
#15 webhook customer.subscription.updated evt_0015 customer=cus_team5 id=sub_five_a status=active
fault reorder(#14,#15)
```

Delivery order: send #15, send #14

State diff (+ only in the faulted run, - only in the in-order run):

```
+ subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"pro","reference_id":"usr_team5","seats":1,"status":"trialing","stripe_customer_id":"cus_team5","stripe_schedule_id":null,"stripe_subscription_id":"sub_five_a"}
- subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"pro","reference_id":"usr_team5","seats":1,"status":"active","stripe_customer_id":"cus_team5","stripe_schedule_id":null,"stripe_subscription_id":"sub_five_a"}
```

Repro: `finding-7/repro.sh`

### Finding 8

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#1,#4).

```
#0 webhook customer.subscription.created evt_0000 customer=cus_team1 id=sub_one_a status=trialing
#1 webhook customer.subscription.updated evt_0001 customer=cus_team1 id=sub_one_a status=active
#4 webhook customer.subscription.deleted evt_0004 customer=cus_team1 id=sub_one_a status=canceled
fault reorder(#1,#4)
```

Delivery order: send #0, send #4, send #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team1","seats":1,"status":"active","stripe_customer_id":"cus_team1","stripe_schedule_id":null,"stripe_subscription_id":"sub_one_a"}
- subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team1","seats":1,"status":"canceled","stripe_customer_id":"cus_team1","stripe_schedule_id":null,"stripe_subscription_id":"sub_one_a"}
```

Repro: `finding-8/repro.sh`

### Finding 9

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#2,#4).

```
#0 webhook customer.subscription.created evt_0000 customer=cus_team1 id=sub_one_a status=trialing
#2 webhook customer.subscription.updated evt_0002 customer=cus_team1 id=sub_one_a status=past_due
#4 webhook customer.subscription.deleted evt_0004 customer=cus_team1 id=sub_one_a status=canceled
fault reorder(#2,#4)
```

Delivery order: send #0, send #4, send #2

State diff (+ only in the faulted run, - only in the in-order run):

```
+ subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team1","seats":1,"status":"past_due","stripe_customer_id":"cus_team1","stripe_schedule_id":null,"stripe_subscription_id":"sub_one_a"}
- subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team1","seats":1,"status":"canceled","stripe_customer_id":"cus_team1","stripe_schedule_id":null,"stripe_subscription_id":"sub_one_a"}
```

Repro: `finding-9/repro.sh`

### Finding 10

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#3,#4).

```
#0 webhook customer.subscription.created evt_0000 customer=cus_team1 id=sub_one_a status=trialing
#3 webhook customer.subscription.updated evt_0003 customer=cus_team1 id=sub_one_a status=active
#4 webhook customer.subscription.deleted evt_0004 customer=cus_team1 id=sub_one_a status=canceled
fault reorder(#3,#4)
```

Delivery order: send #0, send #4, send #3

State diff (+ only in the faulted run, - only in the in-order run):

```
+ subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team1","seats":1,"status":"active","stripe_customer_id":"cus_team1","stripe_schedule_id":null,"stripe_subscription_id":"sub_one_a"}
- subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team1","seats":1,"status":"canceled","stripe_customer_id":"cus_team1","stripe_schedule_id":null,"stripe_subscription_id":"sub_one_a"}
```

Repro: `finding-10/repro.sh`

### Finding 11

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
+ subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team2","seats":1,"status":"active","stripe_customer_id":"cus_team2","stripe_schedule_id":null,"stripe_subscription_id":"sub_two_a"}
- subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team2","seats":1,"status":"canceled","stripe_customer_id":"cus_team2","stripe_schedule_id":null,"stripe_subscription_id":"sub_two_a"}
```

Repro: `finding-11/repro.sh`

### Finding 12

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#6,#7).

```
#5 webhook customer.subscription.created evt_0005 customer=cus_team2 id=sub_two_a status=active
#6 webhook customer.subscription.updated evt_0006 customer=cus_team2 id=sub_two_a status=active
#7 webhook customer.subscription.deleted evt_0007 customer=cus_team2 id=sub_two_a status=canceled
fault reorder(#6,#7)
```

Delivery order: send #5, send #7, send #6

State diff (+ only in the faulted run, - only in the in-order run):

```
+ subscription x1 {"billing_interval":"month","cancel_at_period_end":true,"plan":"starter","reference_id":"usr_team2","seats":1,"status":"active","stripe_customer_id":"cus_team2","stripe_schedule_id":null,"stripe_subscription_id":"sub_two_a"}
- subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team2","seats":1,"status":"canceled","stripe_customer_id":"cus_team2","stripe_schedule_id":null,"stripe_subscription_id":"sub_two_a"}
```

Repro: `finding-12/repro.sh`

### Finding 13

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#6,#9).

```
#6 webhook customer.subscription.updated evt_0006 customer=cus_team2 id=sub_two_a status=active
#8 webhook customer.subscription.created evt_0008 customer=cus_team2 id=sub_two_b status=active
#9 webhook customer.subscription.updated evt_0009 customer=cus_team2 id=sub_two_b status=active
fault reorder(#6,#9)
```

Delivery order: send #9, send #8, send #6

State diff (+ only in the faulted run, - only in the in-order run):

```
+ subscription x1 {"billing_interval":"month","cancel_at_period_end":true,"plan":"starter","reference_id":"usr_team2","seats":1,"status":"active","stripe_customer_id":"cus_team2","stripe_schedule_id":null,"stripe_subscription_id":"sub_two_a"}
- subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"pro","reference_id":"usr_team2","seats":1,"status":"active","stripe_customer_id":"cus_team2","stripe_schedule_id":null,"stripe_subscription_id":"sub_two_b"}
```

Repro: `finding-13/repro.sh`

### Finding 14

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#8,#9).

```
#5 webhook customer.subscription.created evt_0005 customer=cus_team2 id=sub_two_a status=active
#8 webhook customer.subscription.created evt_0008 customer=cus_team2 id=sub_two_b status=active
#9 webhook customer.subscription.updated evt_0009 customer=cus_team2 id=sub_two_b status=active
fault reorder(#8,#9)
```

Delivery order: send #5, send #9, send #8

State diff (+ only in the faulted run, - only in the in-order run):

```
- subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team2","seats":1,"status":"active","stripe_customer_id":"cus_team2","stripe_schedule_id":null,"stripe_subscription_id":"sub_two_a"}
```

Repro: `finding-14/repro.sh`

### Finding 15

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#7,#9).

```
#5 webhook customer.subscription.created evt_0005 customer=cus_team2 id=sub_two_a status=active
#7 webhook customer.subscription.deleted evt_0007 customer=cus_team2 id=sub_two_a status=canceled
#8 webhook customer.subscription.created evt_0008 customer=cus_team2 id=sub_two_b status=active
#9 webhook customer.subscription.updated evt_0009 customer=cus_team2 id=sub_two_b status=active
fault reorder(#7,#9)
```

Delivery order: send #5, send #9, send #8, send #7

State diff (+ only in the faulted run, - only in the in-order run):

```
- subscription x1 {"billing_interval":"month","cancel_at_period_end":false,"plan":"starter","reference_id":"usr_team2","seats":1,"status":"canceled","stripe_customer_id":"cus_team2","stripe_schedule_id":null,"stripe_subscription_id":"sub_two_a"}
```

Repro: `finding-15/repro.sh`

Hunt time: 1125 s.
