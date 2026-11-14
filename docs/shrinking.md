# Shrinking

A failing schedule is reduced with Zeller's ddmin, run over one combined list of trace requests and faults. Code: [internal/shrink/shrink.go](../internal/shrink/shrink.go) and `Shrink` in [internal/explore/explore.go](../internal/explore/explore.go).

## How a candidate is tested

- Dropping a request also drops every fault that names it.
- Each candidate is compared against the fault-free baseline of **its own reduced trace**. Baselines are cached by request subset. This keeps the result sound when the removed request is the one that created an id.
- A candidate with a `concurrent` fault is run 3 times. It counts as failing only if all 3 runs fired and failed.
- A candidate whose fault did not fire counts as passing. That can only keep extra elements, never drop a needed one.
- Candidates are evaluated in parallel, and ddmin always takes the lowest-index failing candidate, so the result matches a sequential run.

## Guarantee

The result is **1-minimal over fired candidates**: removing any single remaining request or fault makes the failure go away. This holds because ddmin only stops at granularity n = |c|, where every complement is "c minus one element" and all of them passed.

The bench re-checks this by brute force for every bug, removing each element in turn and rerunning.

## Measured

Each failing 67-element schedule shrank to 3 to 5 elements in 2 to 18 s on 8 lanes ([results/results.json](../results/results.json)).

The slowest case is `webhook-check-then-insert`. Its candidates carry a `concurrent` fault, so each one runs 3 times, and 48 of its candidates could not fire because a request they depend on had been dropped.
