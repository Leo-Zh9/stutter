# stutter hunt: getpaykit/paykit @ c7572ff

Generated 2026-11-14T23:42:00Z. Trace: 16 requests. Pre-flight: every recorded webhook delivery got 2xx.

| Fault type | Schedules | Fired | Not fired | Changed the final state | No-op passes |
|---|---|---|---|---|---|
| concurrent | 16 | 16 | 0 | 0 | 0 |
| crash | 32 | 32 | 0 | 16 | 0 |
| lost | 16 | 16 | 0 | 0 | 0 |
| reorder | 23 | 23 | 0 | 13 | 0 |

No-op reorder pairs (dropping both events leaves the state unchanged, so their pass says nothing about ordering): 0.

## Effects: 26 distinct end-state changes

1. Findings [1]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"trialing","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team1","providerEventId":"evt_0000","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_one_a","providerSubscriptionScheduleId":null,"status":"trialing"}},"status":"processing","stripe_event_id":"evt_0000","type":"subscription.updated"}
- paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team1","providerEventId":"evt_0000","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_one_a","providerSubscriptionScheduleId":null,"status":"trialing"}},"status":"processed","stripe_event_id":"evt_0000","type":"subscription.updated"}
```

2. Findings [2]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team1","providerEventId":"evt_0001","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_one_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processing","stripe_event_id":"evt_0001","type":"subscription.updated"}
- paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team1","providerEventId":"evt_0001","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_one_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processed","stripe_event_id":"evt_0001","type":"subscription.updated"}
```

3. Findings [3]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"past_due","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team1","providerEventId":"evt_0002","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_one_a","providerSubscriptionScheduleId":null,"status":"past_due"}},"status":"processing","stripe_event_id":"evt_0002","type":"subscription.updated"}
- paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team1","providerEventId":"evt_0002","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_one_a","providerSubscriptionScheduleId":null,"status":"past_due"}},"status":"processed","stripe_event_id":"evt_0002","type":"subscription.updated"}
```

4. Findings [4]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team1","providerEventId":"evt_0003","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_one_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processing","stripe_event_id":"evt_0003","type":"subscription.updated"}
- paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team1","providerEventId":"evt_0003","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_one_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processed","stripe_event_id":"evt_0003","type":"subscription.updated"}
```

5. Findings [5]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":false,"canceled":true,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"canceled","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team1","providerEventId":"evt_0004","providerSubscriptionId":"sub_one_a"},"status":"processing","stripe_event_id":"evt_0004","type":"subscription.deleted"}
- paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team1","providerEventId":"evt_0004","providerSubscriptionId":"sub_one_a"},"status":"processed","stripe_event_id":"evt_0004","type":"subscription.deleted"}
```

6. Findings [6]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team2","providerEventId":"evt_0005","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_two_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processing","stripe_event_id":"evt_0005","type":"subscription.updated"}
- paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team2","providerEventId":"evt_0005","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_two_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processed","stripe_event_id":"evt_0005","type":"subscription.updated"}
```

7. Findings [7]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":true,"canceled":true,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team2","providerEventId":"evt_0006","subscription":{"cancelAtPeriodEnd":true,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_two_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processing","stripe_event_id":"evt_0006","type":"subscription.updated"}
- paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team2","providerEventId":"evt_0006","subscription":{"cancelAtPeriodEnd":true,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_two_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processed","stripe_event_id":"evt_0006","type":"subscription.updated"}
```

8. Findings [8]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":false,"canceled":true,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"canceled","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team2","providerEventId":"evt_0007","providerSubscriptionId":"sub_two_a"},"status":"processing","stripe_event_id":"evt_0007","type":"subscription.deleted"}
- paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team2","providerEventId":"evt_0007","providerSubscriptionId":"sub_two_a"},"status":"processed","stripe_event_id":"evt_0007","type":"subscription.deleted"}
```

9. Findings [9]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_b","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_two_b","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_b","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_b","stripe_subscription_schedule_id":null}
+ paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team2","providerEventId":"evt_0008","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_plus","productId":"prod_plus"},"providerSubscriptionId":"sub_two_b","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processing","stripe_event_id":"evt_0008","type":"subscription.updated"}
- paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team2","providerEventId":"evt_0008","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_plus","productId":"prod_plus"},"providerSubscriptionId":"sub_two_b","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processed","stripe_event_id":"evt_0008","type":"subscription.updated"}
```

10. Findings [10]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_b","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_two_b","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_b","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_b","stripe_subscription_schedule_id":null}
+ paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team2","providerEventId":"evt_0009","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_plus","productId":"prod_plus"},"providerSubscriptionId":"sub_two_b","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processing","stripe_event_id":"evt_0009","type":"subscription.updated"}
- paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team2","providerEventId":"evt_0009","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_plus","productId":"prod_plus"},"providerSubscriptionId":"sub_two_b","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processed","stripe_event_id":"evt_0009","type":"subscription.updated"}
```

11. Findings [11]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team3","id":"psub_three_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_three_a","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team3","id":"psub_three_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_three_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team3","providerEventId":"evt_0010","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_three_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processing","stripe_event_id":"evt_0010","type":"subscription.updated"}
- paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team3","providerEventId":"evt_0010","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_three_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processed","stripe_event_id":"evt_0010","type":"subscription.updated"}
```

12. Findings [12]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team3","id":"psub_three_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_three_a","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team3","id":"psub_three_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_three_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team3","providerEventId":"evt_0011","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_plus","productId":"prod_plus"},"providerSubscriptionId":"sub_three_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processing","stripe_event_id":"evt_0011","type":"subscription.updated"}
- paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team3","providerEventId":"evt_0011","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_plus","productId":"prod_plus"},"providerSubscriptionId":"sub_three_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processed","stripe_event_id":"evt_0011","type":"subscription.updated"}
```

13. Findings [13]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team4","id":"psub_four_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_four_a","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team4","id":"psub_four_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"trialing","stripe_subscription_id":"sub_four_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team4","providerEventId":"evt_0012","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_four_a","providerSubscriptionScheduleId":null,"status":"trialing"}},"status":"processing","stripe_event_id":"evt_0012","type":"subscription.updated"}
- paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team4","providerEventId":"evt_0012","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_four_a","providerSubscriptionScheduleId":null,"status":"trialing"}},"status":"processed","stripe_event_id":"evt_0012","type":"subscription.updated"}
```

14. Findings [14]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team4","id":"psub_four_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_four_a","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team4","id":"psub_four_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"unpaid","stripe_subscription_id":"sub_four_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team4","providerEventId":"evt_0013","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_four_a","providerSubscriptionScheduleId":null,"status":"unpaid"}},"status":"processing","stripe_event_id":"evt_0013","type":"subscription.updated"}
- paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team4","providerEventId":"evt_0013","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_four_a","providerSubscriptionScheduleId":null,"status":"unpaid"}},"status":"processed","stripe_event_id":"evt_0013","type":"subscription.updated"}
```

15. Findings [15]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team5","id":"psub_five_a","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_five_a","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team5","id":"psub_five_a","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"trialing","stripe_subscription_id":"sub_five_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team5","providerEventId":"evt_0014","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_plus","productId":"prod_plus"},"providerSubscriptionId":"sub_five_a","providerSubscriptionScheduleId":null,"status":"trialing"}},"status":"processing","stripe_event_id":"evt_0014","type":"subscription.updated"}
- paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team5","providerEventId":"evt_0014","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_plus","productId":"prod_plus"},"providerSubscriptionId":"sub_five_a","providerSubscriptionScheduleId":null,"status":"trialing"}},"status":"processed","stripe_event_id":"evt_0014","type":"subscription.updated"}
```

16. Findings [16]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team5","id":"psub_five_a","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_five_a","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team5","id":"psub_five_a","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_five_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team5","providerEventId":"evt_0015","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_plus","productId":"prod_plus"},"providerSubscriptionId":"sub_five_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processing","stripe_event_id":"evt_0015","type":"subscription.updated"}
- paykit_webhook_event {"error":null,"payload":{"providerCustomerId":"cus_team5","providerEventId":"evt_0015","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_plus","productId":"prod_plus"},"providerSubscriptionId":"sub_five_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processed","stripe_event_id":"evt_0015","type":"subscription.updated"}
```

17. Findings [17]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"trialing","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":false,"canceled":true,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"canceled","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
```

18. Findings [18 20]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":false,"canceled":true,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"canceled","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
```

19. Findings [19]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"past_due","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":false,"canceled":true,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"canceled","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
```

20. Findings [21]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":true,"canceled":true,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
```

21. Findings [22]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":false,"canceled":true,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"canceled","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
```

22. Findings [23 24]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_b","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"ended","stripe_subscription_id":"sub_two_b","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_b","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_b","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"ended","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
```

23. Findings [25]:

```
+ paykit_subscription {"cancel_at_period_end":true,"canceled":true,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":true,"canceled":true,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"canceled","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
```

24. Findings [26 27]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_b","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"ended","stripe_subscription_id":"sub_two_b","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_b","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_b","stripe_subscription_schedule_id":null}
```

25. Findings [28]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team4","id":"psub_four_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"trialing","stripe_subscription_id":"sub_four_a","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team4","id":"psub_four_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"unpaid","stripe_subscription_id":"sub_four_a","stripe_subscription_schedule_id":null}
```

26. Findings [29]:

```
+ paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team5","id":"psub_five_a","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"trialing","stripe_subscription_id":"sub_five_a","stripe_subscription_schedule_id":null}
- paykit_subscription {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team5","id":"psub_five_a","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_five_a","stripe_subscription_schedule_id":null}
```

## Findings: 29 distinct minimal reproductions

### Finding 1

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#0@commit1).

```
#0 webhook customer.subscription.created evt_0000 customer=cus_team1 id=sub_one_a status=trialing
fault crash(#0@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #0

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"trialing","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team1","providerEventId":"evt_0000","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_one_a","providerSubscriptionScheduleId":null,"status":"trialing"}},"status":"processing","stripe_event_id":"evt_0000","type":"subscription.updated"}
- paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team1","providerEventId":"evt_0000","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_one_a","providerSubscriptionScheduleId":null,"status":"trialing"}},"status":"processed","stripe_event_id":"evt_0000","type":"subscription.updated"}
```

Repro: `finding-1/repro.sh`

### Finding 2

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#1@commit1).

```
#1 webhook customer.subscription.updated evt_0001 customer=cus_team1 id=sub_one_a status=active
fault crash(#1@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team1","providerEventId":"evt_0001","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_one_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processing","stripe_event_id":"evt_0001","type":"subscription.updated"}
- paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team1","providerEventId":"evt_0001","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_one_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processed","stripe_event_id":"evt_0001","type":"subscription.updated"}
```

Repro: `finding-2/repro.sh`

### Finding 3

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#2@commit1).

```
#2 webhook customer.subscription.updated evt_0002 customer=cus_team1 id=sub_one_a status=past_due
fault crash(#2@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #2

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"past_due","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team1","providerEventId":"evt_0002","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_one_a","providerSubscriptionScheduleId":null,"status":"past_due"}},"status":"processing","stripe_event_id":"evt_0002","type":"subscription.updated"}
- paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team1","providerEventId":"evt_0002","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_one_a","providerSubscriptionScheduleId":null,"status":"past_due"}},"status":"processed","stripe_event_id":"evt_0002","type":"subscription.updated"}
```

Repro: `finding-3/repro.sh`

### Finding 4

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#3@commit1).

```
#3 webhook customer.subscription.updated evt_0003 customer=cus_team1 id=sub_one_a status=active
fault crash(#3@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #3

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team1","providerEventId":"evt_0003","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_one_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processing","stripe_event_id":"evt_0003","type":"subscription.updated"}
- paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team1","providerEventId":"evt_0003","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_one_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processed","stripe_event_id":"evt_0003","type":"subscription.updated"}
```

Repro: `finding-4/repro.sh`

### Finding 5

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#4@commit1).

```
#4 webhook customer.subscription.deleted evt_0004 customer=cus_team1 id=sub_one_a status=canceled
fault crash(#4@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #4

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":true,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"canceled","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team1","providerEventId":"evt_0004","providerSubscriptionId":"sub_one_a"},"status":"processing","stripe_event_id":"evt_0004","type":"subscription.deleted"}
- paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team1","providerEventId":"evt_0004","providerSubscriptionId":"sub_one_a"},"status":"processed","stripe_event_id":"evt_0004","type":"subscription.deleted"}
```

Repro: `finding-5/repro.sh`

### Finding 6

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#5@commit1).

```
#5 webhook customer.subscription.created evt_0005 customer=cus_team2 id=sub_two_a status=active
fault crash(#5@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team2","providerEventId":"evt_0005","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_two_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processing","stripe_event_id":"evt_0005","type":"subscription.updated"}
- paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team2","providerEventId":"evt_0005","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_two_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processed","stripe_event_id":"evt_0005","type":"subscription.updated"}
```

Repro: `finding-6/repro.sh`

### Finding 7

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#6@commit1).

```
#6 webhook customer.subscription.updated evt_0006 customer=cus_team2 id=sub_two_a status=active
fault crash(#6@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #6

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":true,"canceled":true,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team2","providerEventId":"evt_0006","subscription":{"cancelAtPeriodEnd":true,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_two_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processing","stripe_event_id":"evt_0006","type":"subscription.updated"}
- paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team2","providerEventId":"evt_0006","subscription":{"cancelAtPeriodEnd":true,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_two_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processed","stripe_event_id":"evt_0006","type":"subscription.updated"}
```

Repro: `finding-7/repro.sh`

### Finding 8

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#7@commit1).

```
#7 webhook customer.subscription.deleted evt_0007 customer=cus_team2 id=sub_two_a status=canceled
fault crash(#7@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #7

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":true,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"canceled","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team2","providerEventId":"evt_0007","providerSubscriptionId":"sub_two_a"},"status":"processing","stripe_event_id":"evt_0007","type":"subscription.deleted"}
- paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team2","providerEventId":"evt_0007","providerSubscriptionId":"sub_two_a"},"status":"processed","stripe_event_id":"evt_0007","type":"subscription.deleted"}
```

Repro: `finding-8/repro.sh`

### Finding 9

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#8@commit1).

```
#8 webhook customer.subscription.created evt_0008 customer=cus_team2 id=sub_two_b status=active
fault crash(#8@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #8

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_b","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_two_b","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_b","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_b","stripe_subscription_schedule_id":null}
+ paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team2","providerEventId":"evt_0008","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_plus","productId":"prod_plus"},"providerSubscriptionId":"sub_two_b","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processing","stripe_event_id":"evt_0008","type":"subscription.updated"}
- paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team2","providerEventId":"evt_0008","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_plus","productId":"prod_plus"},"providerSubscriptionId":"sub_two_b","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processed","stripe_event_id":"evt_0008","type":"subscription.updated"}
```

Repro: `finding-9/repro.sh`

### Finding 10

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#9@commit1).

```
#9 webhook customer.subscription.updated evt_0009 customer=cus_team2 id=sub_two_b status=active
fault crash(#9@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #9

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_b","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_two_b","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_b","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_b","stripe_subscription_schedule_id":null}
+ paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team2","providerEventId":"evt_0009","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_plus","productId":"prod_plus"},"providerSubscriptionId":"sub_two_b","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processing","stripe_event_id":"evt_0009","type":"subscription.updated"}
- paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team2","providerEventId":"evt_0009","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_plus","productId":"prod_plus"},"providerSubscriptionId":"sub_two_b","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processed","stripe_event_id":"evt_0009","type":"subscription.updated"}
```

Repro: `finding-10/repro.sh`

### Finding 11

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#10@commit1).

```
#10 webhook customer.subscription.created evt_0010 customer=cus_team3 id=sub_three_a status=active
fault crash(#10@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #10

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team3","id":"psub_three_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_three_a","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team3","id":"psub_three_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_three_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team3","providerEventId":"evt_0010","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_three_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processing","stripe_event_id":"evt_0010","type":"subscription.updated"}
- paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team3","providerEventId":"evt_0010","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_three_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processed","stripe_event_id":"evt_0010","type":"subscription.updated"}
```

Repro: `finding-11/repro.sh`

### Finding 12

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#11@commit1).

```
#11 webhook customer.subscription.updated evt_0011 customer=cus_team3 id=sub_three_a status=active
fault crash(#11@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #11

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team3","id":"psub_three_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_three_a","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team3","id":"psub_three_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_three_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team3","providerEventId":"evt_0011","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_plus","productId":"prod_plus"},"providerSubscriptionId":"sub_three_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processing","stripe_event_id":"evt_0011","type":"subscription.updated"}
- paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team3","providerEventId":"evt_0011","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_plus","productId":"prod_plus"},"providerSubscriptionId":"sub_three_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processed","stripe_event_id":"evt_0011","type":"subscription.updated"}
```

Repro: `finding-12/repro.sh`

### Finding 13

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#12@commit1).

```
#12 webhook customer.subscription.created evt_0012 customer=cus_team4 id=sub_four_a status=trialing
fault crash(#12@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #12

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team4","id":"psub_four_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_four_a","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team4","id":"psub_four_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"trialing","stripe_subscription_id":"sub_four_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team4","providerEventId":"evt_0012","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_four_a","providerSubscriptionScheduleId":null,"status":"trialing"}},"status":"processing","stripe_event_id":"evt_0012","type":"subscription.updated"}
- paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team4","providerEventId":"evt_0012","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_four_a","providerSubscriptionScheduleId":null,"status":"trialing"}},"status":"processed","stripe_event_id":"evt_0012","type":"subscription.updated"}
```

Repro: `finding-13/repro.sh`

### Finding 14

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#13@commit1).

```
#13 webhook customer.subscription.updated evt_0013 customer=cus_team4 id=sub_four_a status=unpaid
fault crash(#13@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #13

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team4","id":"psub_four_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_four_a","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team4","id":"psub_four_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"unpaid","stripe_subscription_id":"sub_four_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team4","providerEventId":"evt_0013","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_four_a","providerSubscriptionScheduleId":null,"status":"unpaid"}},"status":"processing","stripe_event_id":"evt_0013","type":"subscription.updated"}
- paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team4","providerEventId":"evt_0013","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_basic","productId":"prod_basic"},"providerSubscriptionId":"sub_four_a","providerSubscriptionScheduleId":null,"status":"unpaid"}},"status":"processed","stripe_event_id":"evt_0013","type":"subscription.updated"}
```

Repro: `finding-14/repro.sh`

### Finding 15

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#14@commit1).

```
#14 webhook customer.subscription.created evt_0014 customer=cus_team5 id=sub_five_a status=trialing
fault crash(#14@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #14

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team5","id":"psub_five_a","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_five_a","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team5","id":"psub_five_a","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"trialing","stripe_subscription_id":"sub_five_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team5","providerEventId":"evt_0014","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_plus","productId":"prod_plus"},"providerSubscriptionId":"sub_five_a","providerSubscriptionScheduleId":null,"status":"trialing"}},"status":"processing","stripe_event_id":"evt_0014","type":"subscription.updated"}
- paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team5","providerEventId":"evt_0014","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_plus","productId":"prod_plus"},"providerSubscriptionId":"sub_five_a","providerSubscriptionScheduleId":null,"status":"trialing"}},"status":"processed","stripe_event_id":"evt_0014","type":"subscription.updated"}
```

Repro: `finding-15/repro.sh`

### Finding 16

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#15@commit1).

```
#15 webhook customer.subscription.updated evt_0015 customer=cus_team5 id=sub_five_a status=active
fault crash(#15@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #15

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team5","id":"psub_five_a","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"incomplete","stripe_subscription_id":"sub_five_a","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team5","id":"psub_five_a","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_five_a","stripe_subscription_schedule_id":null}
+ paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team5","providerEventId":"evt_0015","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_plus","productId":"prod_plus"},"providerSubscriptionId":"sub_five_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processing","stripe_event_id":"evt_0015","type":"subscription.updated"}
- paykit_webhook_event x1 {"error":null,"payload":{"providerCustomerId":"cus_team5","providerEventId":"evt_0015","subscription":{"cancelAtPeriodEnd":false,"canceledAt":null,"currentPeriodEndAt":"<ts>","currentPeriodStartAt":"<ts>","endedAt":null,"providerProduct":{"priceId":"price_prod_plus","productId":"prod_plus"},"providerSubscriptionId":"sub_five_a","providerSubscriptionScheduleId":null,"status":"active"}},"status":"processed","stripe_event_id":"evt_0015","type":"subscription.updated"}
```

Repro: `finding-16/repro.sh`

### Finding 17

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#0,#4).

```
#0 webhook customer.subscription.created evt_0000 customer=cus_team1 id=sub_one_a status=trialing
#4 webhook customer.subscription.deleted evt_0004 customer=cus_team1 id=sub_one_a status=canceled
fault reorder(#0,#4)
```

Delivery order: send #4, send #0

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"trialing","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":true,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"canceled","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
```

Repro: `finding-17/repro.sh`

### Finding 18

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#1,#4).

```
#1 webhook customer.subscription.updated evt_0001 customer=cus_team1 id=sub_one_a status=active
#4 webhook customer.subscription.deleted evt_0004 customer=cus_team1 id=sub_one_a status=canceled
fault reorder(#1,#4)
```

Delivery order: send #4, send #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":true,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"canceled","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
```

Repro: `finding-18/repro.sh`

### Finding 19

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#2,#4).

```
#2 webhook customer.subscription.updated evt_0002 customer=cus_team1 id=sub_one_a status=past_due
#4 webhook customer.subscription.deleted evt_0004 customer=cus_team1 id=sub_one_a status=canceled
fault reorder(#2,#4)
```

Delivery order: send #4, send #2

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"past_due","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":true,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"canceled","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
```

Repro: `finding-19/repro.sh`

### Finding 20

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#3,#4).

```
#3 webhook customer.subscription.updated evt_0003 customer=cus_team1 id=sub_one_a status=active
#4 webhook customer.subscription.deleted evt_0004 customer=cus_team1 id=sub_one_a status=canceled
fault reorder(#3,#4)
```

Delivery order: send #4, send #3

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":true,"customer_id":"user_team1","id":"psub_one_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"canceled","stripe_subscription_id":"sub_one_a","stripe_subscription_schedule_id":null}
```

Repro: `finding-20/repro.sh`

### Finding 21

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#5,#6).

```
#5 webhook customer.subscription.created evt_0005 customer=cus_team2 id=sub_two_a status=active
#6 webhook customer.subscription.updated evt_0006 customer=cus_team2 id=sub_two_a status=active
fault reorder(#5,#6)
```

Delivery order: send #6, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":true,"canceled":true,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
```

Repro: `finding-21/repro.sh`

### Finding 22

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#5,#7).

```
#5 webhook customer.subscription.created evt_0005 customer=cus_team2 id=sub_two_a status=active
#7 webhook customer.subscription.deleted evt_0007 customer=cus_team2 id=sub_two_a status=canceled
fault reorder(#5,#7)
```

Delivery order: send #7, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":true,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"canceled","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
```

Repro: `finding-22/repro.sh`

### Finding 23

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#5,#8).

```
#5 webhook customer.subscription.created evt_0005 customer=cus_team2 id=sub_two_a status=active
#8 webhook customer.subscription.created evt_0008 customer=cus_team2 id=sub_two_b status=active
fault reorder(#5,#8)
```

Delivery order: send #8, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_b","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"ended","stripe_subscription_id":"sub_two_b","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_b","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_b","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"ended","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
```

Repro: `finding-23/repro.sh`

### Finding 24

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#5,#9).

```
#5 webhook customer.subscription.created evt_0005 customer=cus_team2 id=sub_two_a status=active
#9 webhook customer.subscription.updated evt_0009 customer=cus_team2 id=sub_two_b status=active
fault reorder(#5,#9)
```

Delivery order: send #9, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_b","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"ended","stripe_subscription_id":"sub_two_b","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_b","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_b","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"ended","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
```

Repro: `finding-24/repro.sh`

### Finding 25

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#6,#7).

```
#6 webhook customer.subscription.updated evt_0006 customer=cus_team2 id=sub_two_a status=active
#7 webhook customer.subscription.deleted evt_0007 customer=cus_team2 id=sub_two_a status=canceled
fault reorder(#6,#7)
```

Delivery order: send #7, send #6

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":true,"canceled":true,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":true,"canceled":true,"customer_id":"user_team2","id":"psub_two_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"canceled","stripe_subscription_id":"sub_two_a","stripe_subscription_schedule_id":null}
```

Repro: `finding-25/repro.sh`

### Finding 26

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#6,#8).

```
#6 webhook customer.subscription.updated evt_0006 customer=cus_team2 id=sub_two_a status=active
#8 webhook customer.subscription.created evt_0008 customer=cus_team2 id=sub_two_b status=active
fault reorder(#6,#8)
```

Delivery order: send #8, send #6

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_b","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"ended","stripe_subscription_id":"sub_two_b","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_b","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_b","stripe_subscription_schedule_id":null}
```

Repro: `finding-26/repro.sh`

### Finding 27

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#6,#9).

```
#6 webhook customer.subscription.updated evt_0006 customer=cus_team2 id=sub_two_a status=active
#9 webhook customer.subscription.updated evt_0009 customer=cus_team2 id=sub_two_b status=active
fault reorder(#6,#9)
```

Delivery order: send #9, send #6

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_b","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"ended","stripe_subscription_id":"sub_two_b","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team2","id":"psub_two_b","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_two_b","stripe_subscription_schedule_id":null}
```

Repro: `finding-27/repro.sh`

### Finding 28

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#12,#13).

```
#12 webhook customer.subscription.created evt_0012 customer=cus_team4 id=sub_four_a status=trialing
#13 webhook customer.subscription.updated evt_0013 customer=cus_team4 id=sub_four_a status=unpaid
fault reorder(#12,#13)
```

Delivery order: send #13, send #12

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team4","id":"psub_four_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"trialing","stripe_subscription_id":"sub_four_a","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team4","id":"psub_four_a","product_internal_id":"prod_internal_basic","quantity":1,"scheduled_product_id":null,"status":"unpaid","stripe_subscription_id":"sub_four_a","stripe_subscription_schedule_id":null}
```

Repro: `finding-28/repro.sh`

### Finding 29

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: reorder(#14,#15).

```
#14 webhook customer.subscription.created evt_0014 customer=cus_team5 id=sub_five_a status=trialing
#15 webhook customer.subscription.updated evt_0015 customer=cus_team5 id=sub_five_a status=active
fault reorder(#14,#15)
```

Delivery order: send #15, send #14

State diff (+ only in the faulted run, - only in the in-order run):

```
+ paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team5","id":"psub_five_a","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"trialing","stripe_subscription_id":"sub_five_a","stripe_subscription_schedule_id":null}
- paykit_subscription x1 {"cancel_at_period_end":false,"canceled":false,"customer_id":"user_team5","id":"psub_five_a","product_internal_id":"prod_internal_plus","quantity":1,"scheduled_product_id":null,"status":"active","stripe_subscription_id":"sub_five_a","stripe_subscription_schedule_id":null}
```

Repro: `finding-29/repro.sh`

Hunt time: 547 s.
