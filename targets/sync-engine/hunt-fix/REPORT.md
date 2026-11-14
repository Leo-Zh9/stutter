# stutter hunt: stripe/sync-engine @ d486b41

Generated 2026-11-14T23:19:04Z. Trace: 16 requests. Pre-flight: every recorded webhook delivery got 2xx.

| Fault type | Schedules | Fired | Not fired | Changed the final state | No-op passes |
|---|---|---|---|---|---|
| concurrent | 16 | 16 | 0 | 0 | 0 |
| crash | 32 | 32 | 0 | 0 | 0 |
| lost | 16 | 16 | 0 | 0 | 0 |
| reorder | 23 | 23 | 0 | 0 | 11 |

No-op reorder pairs (dropping both events leaves the state unchanged, so their pass says nothing about ordering): 11.

## Effects: 0 distinct end-state changes

## Findings: 0 distinct minimal reproductions

Hunt time: 16 s.
