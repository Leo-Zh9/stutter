# stutter hunt: langchain-ai/paid-media-agent @ d4ee460

Generated 2026-11-18T16:50:37Z. Trace: 6 requests. Pre-flight: every recorded webhook delivery got 2xx.

| Fault type | Schedules | Fired | Not fired | Diverged | Nondetermined | Changed the final state | No-op passes |
|---|---|---|---|---|---|---|---|
| concurrent | 3 | 0 | 3 | 0 | 0 | 0 | 0 |
| crash | 126 | 119 | 0 | 0 | 7 | 92 | 0 |
| lost | 3 | 3 | 0 | 0 | 0 | 0 | 0 |

No-op reorder pairs (dropping both events leaves the state unchanged, so their pass says nothing about ordering): 0.

## Effects: 9 distinct end-state changes

1. Findings [1 2 3 4 5 6]:

```
+ pma_approvals {"claim":{"account_ref":"demo-google","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"google_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+6 referencing rows)
- pma_proposals {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+4 referencing rows)
```

2. Findings [7 8 9 10 11 12 14]:

```
+ pma_approvals {"claim":{"account_ref":"demo-google","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"google_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"awaiting_approval"},"state":"awaiting_approval","thread_id":"t-google"} (+4 referencing rows)
- pma_proposals {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+4 referencing rows)
- pma_receipts {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":150.0}]}}
```

3. Findings [13 15 16]:

```
+ pma_proposals {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"executing"},"state":"executing","thread_id":"t-google"} (+2 referencing rows)
- pma_proposals {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+4 referencing rows)
- pma_receipts {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":150.0}]}}
```

4. Findings [17]:

```
+ pma_proposals {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+2 referencing rows)
- pma_proposals {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+4 referencing rows)
- pma_receipts {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":150.0}]}}
```

5. Findings [18 19 20 21 22 23]:

```
+ pma_approvals {"claim":{"account_ref":"demo-reddit","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"reddit_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+6 referencing rows)
- pma_proposals {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
```

6. Findings [24 25 26 27 28 29 30 31]:

```
+ pma_approvals {"claim":{"account_ref":"demo-reddit","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"reddit_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"awaiting_approval"},"state":"awaiting_approval","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_receipts {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

7. Findings [32]:

```
+ pma_proposals {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"executing"},"state":"executing","thread_id":"t-reddit"} (+2 referencing rows)
- pma_proposals {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_receipts {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

8. Findings [33]:

```
+ pma_proposals {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verifying"},"state":"verifying","thread_id":"t-reddit"} (+2 referencing rows)
- pma_proposals {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_receipts {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

9. Findings [34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92]:

```
+ pma_proposals {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

## Findings: 92 distinct minimal reproductions

### Finding 1

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit4).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
fault crash(#1@commit4)
```

Delivery order: send #0, send, kill the service after its commit 4, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-google","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"google_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+6 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+4 referencing rows)
```

Repro: `finding-1/repro.sh`

### Finding 2

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit5).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
fault crash(#1@commit5)
```

Delivery order: send #0, send, kill the service after its commit 5, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-google","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"google_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+6 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+4 referencing rows)
```

Repro: `finding-2/repro.sh`

### Finding 3

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit6).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
fault crash(#1@commit6)
```

Delivery order: send #0, send, kill the service after its commit 6, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-google","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"google_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+6 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+4 referencing rows)
```

Repro: `finding-3/repro.sh`

### Finding 4

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit7).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
fault crash(#1@commit7)
```

Delivery order: send #0, send, kill the service after its commit 7, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-google","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"google_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+6 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+4 referencing rows)
```

Repro: `finding-4/repro.sh`

### Finding 5

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit8).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
fault crash(#1@commit8)
```

Delivery order: send #0, send, kill the service after its commit 8, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-google","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"google_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+6 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+4 referencing rows)
```

Repro: `finding-5/repro.sh`

### Finding 6

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit9).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
fault crash(#1@commit9)
```

Delivery order: send #0, send, kill the service after its commit 9, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-google","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"google_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+6 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+4 referencing rows)
```

Repro: `finding-6/repro.sh`

### Finding 7

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit10).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
fault crash(#1@commit10)
```

Delivery order: send #0, send, kill the service after its commit 10, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-google","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"google_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"awaiting_approval"},"state":"awaiting_approval","thread_id":"t-google"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+4 referencing rows)
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":150.0}]}}
```

Repro: `finding-7/repro.sh`

### Finding 8

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit11).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
fault crash(#1@commit11)
```

Delivery order: send #0, send, kill the service after its commit 11, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-google","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"google_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"awaiting_approval"},"state":"awaiting_approval","thread_id":"t-google"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+4 referencing rows)
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":150.0}]}}
```

Repro: `finding-8/repro.sh`

### Finding 9

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit12).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
fault crash(#1@commit12)
```

Delivery order: send #0, send, kill the service after its commit 12, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-google","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"google_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"awaiting_approval"},"state":"awaiting_approval","thread_id":"t-google"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+4 referencing rows)
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":150.0}]}}
```

Repro: `finding-9/repro.sh`

### Finding 10

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit13).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
fault crash(#1@commit13)
```

Delivery order: send #0, send, kill the service after its commit 13, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-google","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"google_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"awaiting_approval"},"state":"awaiting_approval","thread_id":"t-google"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+4 referencing rows)
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":150.0}]}}
```

Repro: `finding-10/repro.sh`

### Finding 11

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit14).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
fault crash(#1@commit14)
```

Delivery order: send #0, send, kill the service after its commit 14, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-google","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"google_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"awaiting_approval"},"state":"awaiting_approval","thread_id":"t-google"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+4 referencing rows)
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":150.0}]}}
```

Repro: `finding-11/repro.sh`

### Finding 12

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit15).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
fault crash(#1@commit15)
```

Delivery order: send #0, send, kill the service after its commit 15, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-google","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"google_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"awaiting_approval"},"state":"awaiting_approval","thread_id":"t-google"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+4 referencing rows)
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":150.0}]}}
```

Repro: `finding-12/repro.sh`

### Finding 13

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit16).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
fault crash(#1@commit16)
```

Delivery order: send #0, send, kill the service after its commit 16, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"executing"},"state":"executing","thread_id":"t-google"} (+2 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+4 referencing rows)
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":150.0}]}}
```

Repro: `finding-13/repro.sh`

### Finding 14

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit17).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
fault crash(#1@commit17)
```

Delivery order: send #0, send, kill the service after its commit 17, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-google","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"google_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"awaiting_approval"},"state":"awaiting_approval","thread_id":"t-google"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+4 referencing rows)
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":150.0}]}}
```

Repro: `finding-14/repro.sh`

### Finding 15

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit18).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
fault crash(#1@commit18)
```

Delivery order: send #0, send, kill the service after its commit 18, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"executing"},"state":"executing","thread_id":"t-google"} (+2 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+4 referencing rows)
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":150.0}]}}
```

Repro: `finding-15/repro.sh`

### Finding 16

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit19).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
fault crash(#1@commit19)
```

Delivery order: send #0, send, kill the service after its commit 19, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"executing"},"state":"executing","thread_id":"t-google"} (+2 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+4 referencing rows)
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":150.0}]}}
```

Repro: `finding-16/repro.sh`

### Finding 17

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit21).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
fault crash(#1@commit21)
```

Delivery order: send #0, send, kill the service after its commit 21, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+2 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-google","after":[{"field":"daily_budget","unit":"account currency per day","value":150}],"before":[{"field":"daily_budget","unit":"account currency per day","value":180.0}],"canonical_args":{"campaign_id":"g-101","customer_id":"fixture-google-0001","daily_budget":150},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"google_ads","policy_digest":"239398093931b7df","proposal_id":"<self>","reason":"Brand search is over-pacing; the user asked for 150 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 180.","revision":1,"risk":"medium","risk_flags":["budget_delta"],"schema_hash":"5592572e1d29c911","target_ref":"g-101","thread_id":"t-google","tool_name":"google_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-google"} (+4 referencing rows)
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":150.0}]}}
```

Repro: `finding-17/repro.sh`

### Finding 18

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#5@commit4).

```
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#5@commit4)
```

Delivery order: send #4, send, kill the service after its commit 4, restart, resend #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-reddit","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"reddit_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+6 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
```

Repro: `finding-18/repro.sh`

### Finding 19

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#5@commit5).

```
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#5@commit5)
```

Delivery order: send #4, send, kill the service after its commit 5, restart, resend #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-reddit","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"reddit_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+6 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
```

Repro: `finding-19/repro.sh`

### Finding 20

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#5@commit6).

```
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#5@commit6)
```

Delivery order: send #4, send, kill the service after its commit 6, restart, resend #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-reddit","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"reddit_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+6 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
```

Repro: `finding-20/repro.sh`

### Finding 21

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#5@commit7).

```
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#5@commit7)
```

Delivery order: send #4, send, kill the service after its commit 7, restart, resend #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-reddit","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"reddit_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+6 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
```

Repro: `finding-21/repro.sh`

### Finding 22

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#5@commit8).

```
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#5@commit8)
```

Delivery order: send #4, send, kill the service after its commit 8, restart, resend #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-reddit","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"reddit_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+6 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
```

Repro: `finding-22/repro.sh`

### Finding 23

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#5@commit9).

```
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#5@commit9)
```

Delivery order: send #4, send, kill the service after its commit 9, restart, resend #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-reddit","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"reddit_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+6 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
```

Repro: `finding-23/repro.sh`

### Finding 24

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#5@commit10).

```
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#5@commit10)
```

Delivery order: send #4, send, kill the service after its commit 10, restart, resend #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-reddit","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"reddit_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"awaiting_approval"},"state":"awaiting_approval","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-24/repro.sh`

### Finding 25

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#5@commit11).

```
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#5@commit11)
```

Delivery order: send #4, send, kill the service after its commit 11, restart, resend #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-reddit","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"reddit_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"awaiting_approval"},"state":"awaiting_approval","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-25/repro.sh`

### Finding 26

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#5@commit12).

```
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#5@commit12)
```

Delivery order: send #4, send, kill the service after its commit 12, restart, resend #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-reddit","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"reddit_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"awaiting_approval"},"state":"awaiting_approval","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-26/repro.sh`

### Finding 27

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#5@commit13).

```
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#5@commit13)
```

Delivery order: send #4, send, kill the service after its commit 13, restart, resend #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-reddit","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"reddit_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"awaiting_approval"},"state":"awaiting_approval","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-27/repro.sh`

### Finding 28

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#5@commit14).

```
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#5@commit14)
```

Delivery order: send #4, send, kill the service after its commit 14, restart, resend #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-reddit","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"reddit_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"awaiting_approval"},"state":"awaiting_approval","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-28/repro.sh`

### Finding 29

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#5@commit15).

```
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#5@commit15)
```

Delivery order: send #4, send, kill the service after its commit 15, restart, resend #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-reddit","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"reddit_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"awaiting_approval"},"state":"awaiting_approval","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-29/repro.sh`

### Finding 30

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#5@commit16).

```
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#5@commit16)
```

Delivery order: send #4, send, kill the service after its commit 16, restart, resend #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-reddit","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"reddit_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"awaiting_approval"},"state":"awaiting_approval","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-30/repro.sh`

### Finding 31

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#5@commit17).

```
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#5@commit17)
```

Delivery order: send #4, send, kill the service after its commit 17, restart, resend #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_approvals x1 {"claim":{"account_ref":"demo-reddit","approved_at":"<ts>","approver_ref":"operator","claim_id":"<self>","expires_at":"<ts>","proposal_id":"<ref:pma_proposals>","requester_ref":"operator","revision":1,"tool_name":"reddit_ads__update_campaign_budget"},"proposal_id":"<ref:pma_proposals>","revision":1}
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"awaiting_approval"},"state":"awaiting_approval","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-31/repro.sh`

### Finding 32

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#5@commit18).

```
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#5@commit18)
```

Delivery order: send #4, send, kill the service after its commit 18, restart, resend #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"executing"},"state":"executing","thread_id":"t-reddit"} (+2 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-32/repro.sh`

### Finding 33

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#5@commit19).

```
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#5@commit19)
```

Delivery order: send #4, send, kill the service after its commit 19, restart, resend #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verifying"},"state":"verifying","thread_id":"t-reddit"} (+2 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-33/repro.sh`

### Finding 34

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit24).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#1@commit24)
```

Delivery order: send #0, send, kill the service after its commit 24, restart, resend #1, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-34/repro.sh`

### Finding 35

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit25).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#1@commit25)
```

Delivery order: send #0, send, kill the service after its commit 25, restart, resend #1, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-35/repro.sh`

### Finding 36

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit26).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#1@commit26)
```

Delivery order: send #0, send, kill the service after its commit 26, restart, resend #1, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-36/repro.sh`

### Finding 37

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit27).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#1@commit27)
```

Delivery order: send #0, send, kill the service after its commit 27, restart, resend #1, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-37/repro.sh`

### Finding 38

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit28).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#1@commit28)
```

Delivery order: send #0, send, kill the service after its commit 28, restart, resend #1, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-38/repro.sh`

### Finding 39

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit29).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#1@commit29)
```

Delivery order: send #0, send, kill the service after its commit 29, restart, resend #1, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-39/repro.sh`

### Finding 40

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit30).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#1@commit30)
```

Delivery order: send #0, send, kill the service after its commit 30, restart, resend #1, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-40/repro.sh`

### Finding 41

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit31).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#1@commit31)
```

Delivery order: send #0, send, kill the service after its commit 31, restart, resend #1, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-41/repro.sh`

### Finding 42

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit32).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#1@commit32)
```

Delivery order: send #0, send, kill the service after its commit 32, restart, resend #1, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-42/repro.sh`

### Finding 43

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit33).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#1@commit33)
```

Delivery order: send #0, send, kill the service after its commit 33, restart, resend #1, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-43/repro.sh`

### Finding 44

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit34).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#1@commit34)
```

Delivery order: send #0, send, kill the service after its commit 34, restart, resend #1, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-44/repro.sh`

### Finding 45

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit35).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#1@commit35)
```

Delivery order: send #0, send, kill the service after its commit 35, restart, resend #1, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-45/repro.sh`

### Finding 46

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit36).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#1@commit36)
```

Delivery order: send #0, send, kill the service after its commit 36, restart, resend #1, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-46/repro.sh`

### Finding 47

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit37).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#1@commit37)
```

Delivery order: send #0, send, kill the service after its commit 37, restart, resend #1, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-47/repro.sh`

### Finding 48

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit38).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#1@commit38)
```

Delivery order: send #0, send, kill the service after its commit 38, restart, resend #1, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-48/repro.sh`

### Finding 49

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit39).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#1@commit39)
```

Delivery order: send #0, send, kill the service after its commit 39, restart, resend #1, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-49/repro.sh`

### Finding 50

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit40).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#1@commit40)
```

Delivery order: send #0, send, kill the service after its commit 40, restart, resend #1, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-50/repro.sh`

### Finding 51

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit41).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#1@commit41)
```

Delivery order: send #0, send, kill the service after its commit 41, restart, resend #1, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-51/repro.sh`

### Finding 52

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit42).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#1@commit42)
```

Delivery order: send #0, send, kill the service after its commit 42, restart, resend #1, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-52/repro.sh`

### Finding 53

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit43).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#1@commit43)
```

Delivery order: send #0, send, kill the service after its commit 43, restart, resend #1, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-53/repro.sh`

### Finding 54

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit44).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#1@commit44)
```

Delivery order: send #0, send, kill the service after its commit 44, restart, resend #1, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-54/repro.sh`

### Finding 55

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit45).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#1@commit45)
```

Delivery order: send #0, send, kill the service after its commit 45, restart, resend #1, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-55/repro.sh`

### Finding 56

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit46).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#1@commit46)
```

Delivery order: send #0, send, kill the service after its commit 46, restart, resend #1, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-56/repro.sh`

### Finding 57

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#1@commit47).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#1@commit47)
```

Delivery order: send #0, send, kill the service after its commit 47, restart, resend #1, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-57/repro.sh`

### Finding 58

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#5@commit1).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#5@commit1)
```

Delivery order: send #0, send #1, send #4, send, kill the service after its commit 1, restart, resend #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-58/repro.sh`

### Finding 59

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#5@commit2).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#5@commit2)
```

Delivery order: send #0, send #1, send #4, send, kill the service after its commit 2, restart, resend #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-59/repro.sh`

### Finding 60

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#5@commit3).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#5@commit3)
```

Delivery order: send #0, send #1, send #4, send, kill the service after its commit 3, restart, resend #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-60/repro.sh`

### Finding 61

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit1).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit1)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 1, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-61/repro.sh`

### Finding 62

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit2).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit2)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 2, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-62/repro.sh`

### Finding 63

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit3).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit3)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 3, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-63/repro.sh`

### Finding 64

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit4).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit4)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 4, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-64/repro.sh`

### Finding 65

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit5).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit5)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 5, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-65/repro.sh`

### Finding 66

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit6).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit6)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 6, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-66/repro.sh`

### Finding 67

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit7).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit7)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 7, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-67/repro.sh`

### Finding 68

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit8).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit8)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 8, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-68/repro.sh`

### Finding 69

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit9).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit9)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 9, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-69/repro.sh`

### Finding 70

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit10).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit10)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 10, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-70/repro.sh`

### Finding 71

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit11).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit11)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 11, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-71/repro.sh`

### Finding 72

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit12).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit12)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 12, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-72/repro.sh`

### Finding 73

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit13).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit13)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 13, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-73/repro.sh`

### Finding 74

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit14).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit14)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 14, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-74/repro.sh`

### Finding 75

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit15).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit15)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 15, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-75/repro.sh`

### Finding 76

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit16).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit16)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 16, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-76/repro.sh`

### Finding 77

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit17).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit17)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 17, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-77/repro.sh`

### Finding 78

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit18).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit18)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 18, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-78/repro.sh`

### Finding 79

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit19).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit19)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 19, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-79/repro.sh`

### Finding 80

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit20).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit20)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 20, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-80/repro.sh`

### Finding 81

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit21).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit21)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 21, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-81/repro.sh`

### Finding 82

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit22).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit22)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 22, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-82/repro.sh`

### Finding 83

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit23).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit23)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 23, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-83/repro.sh`

### Finding 84

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit24).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit24)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 24, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-84/repro.sh`

### Finding 85

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit25).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit25)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 25, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-85/repro.sh`

### Finding 86

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit26).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit26)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 26, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-86/repro.sh`

### Finding 87

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit27).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit27)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 27, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-87/repro.sh`

### Finding 88

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit28).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit28)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 28, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-88/repro.sh`

### Finding 89

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit29).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit29)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 29, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-89/repro.sh`

### Finding 90

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit30).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit30)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 30, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-90/repro.sh`

### Finding 91

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit31).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit31)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 31, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-91/repro.sh`

### Finding 92

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true (1 single removals could not fire and count as passing). Reached from 1 single-fault schedules: crash(#3@commit32).

```
#0 POST /threads/t-google/messages {"text":"Lower the daily budget of the Brand Search - US campaign in demo-google to 150."}
#1 POST /proposals/763b6549-6101-46f1-b71f-e2eeac54bfa2/approve {}
#2 POST /threads/t-meta/messages {"text":"Pause the Prospecting - Lookalike 1% campaign in demo-meta."}
#3 POST /proposals/12009b9c-a632-4318-94f3-9d392b24a6a7/reject {"message":"Not this week."}
#4 POST /threads/t-reddit/messages {"text":"Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90."}
#5 POST /proposals/9ef5f743-51ce-457d-b619-f2b207db4658/approve {}
fault crash(#3@commit32)
```

Delivery order: send #0, send #1, send #2, send, kill the service after its commit 32, restart, resend #3, send #4, send #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
- pma_proposals x1 {"record":{"changeset":{"account_ref":"demo-reddit","after":[{"field":"daily_budget","unit":"account currency per day","value":90}],"before":[{"field":"daily_budget","unit":"account currency per day","value":75.0}],"canonical_args":{"account_id":"fixture-reddit-0001","campaign_id":"r-301","daily_budget":90},"catalog_revision":"5d9fa65fb3b21609c8bd","measurement_plan":"Compare spend and conversions over the next 7 days.","platform":"reddit_ads","policy_digest":"96c98bd2e184217a","proposal_id":"<self>","reason":"Dev tools community targeting converts well; the user asked for 90 a day.","requester_ref":"operator","reversal_plan":"Set the daily budget back to 75.","revision":1,"risk":"medium","risk_flags":["budget_delta","budget_increase"],"schema_hash":"e2772a0bfc428ad0","target_ref":"r-301","thread_id":"t-reddit","tool_name":"reddit_ads__update_campaign_budget"},"routing_id":"<self>","state":"verified"},"state":"verified","thread_id":"t-reddit"} (+4 referencing rows)
+ pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-1","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
- pma_receipts x1 {"proposal_id":"<ref:pma_proposals>","receipt":{"catalog_revision":"5d9fa65fb3b21609c8bd","checked_at":"<ts>","mutation_attempted":true,"proposal_id":"<ref:pma_proposals>","provider_acknowledged":true,"provider_operation_ref":"fixture-op-2","readback_attempts":1,"reason":"readback matched the approved change","revision":1,"status":"verified","verified_state":[{"field":"daily_budget","unit":"account currency per day","value":90.0}]}}
```

Repro: `finding-92/repro.sh`

Hunt time: 4713 s.
