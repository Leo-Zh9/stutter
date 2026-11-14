# Guarantees and the tests that hold them

`make test` runs all of these. The database-backed tests use the project-local Postgres, and under `make test` they fail rather than skip when it is missing.

| Guarantee | Test |
|---|---|
| Equivalent states compare equal under any insert order, uuids, sequence gaps and timestamps | `TestEquivalentStatesAreEqual`, `TestBurnedIdsAreNotADiff` |
| Isomorphic parents never swap labels; moving a child between them is a diff | `TestIsomorphicParents` |
| One changed value, one extra row, or a re-pointed reference is a diff | `TestRealChangesAreDetected`, `TestBurnedIdsAreNotADiff` |
| Reference cycles are reported, not looped on | `TestReferenceCycleIsAnError` |
| Unstable columns are named and can be ignored | `TestUnstableColumnsAreNamed` |
| The generated-value rule matches the real catalog (uuid, serial and identity are generated; constants are not) | `TestCatalogGeneratedRule` |
| Lost sends twice | `TestLostFires` |
| Concurrent puts two backends behind the barrier | `TestConcurrentBarrierFires` |
| Reorder swaps the delivery order | `TestReorderSwapsDelivery` |
| A barrier that never fills is not fired, never a pass | `TestBarrierTimeoutIsNotFired` |
| Unkeyed API calls are never faulted | `TestUnkeyedRequestsAreNeverFaulted` |
| Fault-free replay is deterministic | `TestFaultFreeReplayIsDeterministic` |
| ddmin output fails and is 1-minimal, checked by brute force over 300 random failure predicates | `TestDDMinOneMinimalProperty`, `TestDDMinFindsExactCause` |
| No false alarms on the correct variant | `TestCorrectVariantHasNoFalseAlarms` |
| Every in-model bug is found by its own fault type and shrinks to a 1-minimal schedule | `TestEveryInModelBugIsDetectedAndShrunk` |
| The two race bugs survive every `lost` fault; only concurrency finds them | `TestRaceBugsNeedConcurrency` |
| The control stays undetected | `TestControlIsOutsideTheOracle` |

## Mutation checks

A passing test proves little until it is shown to fail when the mechanism it guards breaks. Each mechanism below was broken on purpose and its test was run. All 11 mutants were killed:

| Mutant | Killed by |
|---|---|
| child signatures dropped from the row hash | `TestIsomorphicParents` |
| string values never treated as edges | `TestBurnedIdsAreNotADiff`, `TestEquivalentStatesAreEqual` |
| generated integer ids compared raw | `TestBurnedIdsAreNotADiff` |
| identity columns not treated as generated | `TestCatalogGeneratedRule` |
| volatile defaults not treated as generated | `TestCatalogGeneratedRule` |
| barrier timeout counted as fired | `TestBarrierTimeoutIsNotFired` |
| lost fault skips the discarded first attempt | `TestLostFires` |
| unkeyed calls treated as retryable | `TestUnkeyedRequestsAreNeverFaulted` |
| reorder does not swap | `TestReorderSwapsDelivery` |
| ddmin skips the complement step | `TestDDMinOneMinimalProperty` |
| ddmin stops before granularity reaches len(c) | `TestDDMinOneMinimalProperty` (hangs, then the test times out) |

Two first attempts were replaced because they did not actually break what their test guards:
- **Comparing generated string ids raw** changes nothing, because a row's own id already becomes `<self>`. It was replaced by the integer identity-id mutant above.
- **Capping retries** at the lost fault did not remove the discarded first attempt the test counts. It was replaced by the "skips the discarded first attempt" mutant.
