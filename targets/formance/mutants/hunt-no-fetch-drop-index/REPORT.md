# stutter hunt: Formance Ledger v2.4.18 mutant: lookup removed and unique index dropped

Generated 2026-11-14T12:31:54Z. Trace: 43 requests. Pre-flight: every recorded webhook delivery got 2xx.

| Fault type | Schedules | Fired | Not fired | Changed the final state | No-op passes |
|---|---|---|---|---|---|
| concurrent | 43 | 43 | 0 | 38 | 0 |
| crash | 43 | 43 | 0 | 38 | 0 |
| lost | 43 | 43 | 0 | 38 | 0 |

No-op reorder pairs (dropping both events leaves the state unchanged, so their pass says nothing about ordering): 0.

## Effects: 38 distinct end-state changes

1. Findings [1 34 65]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"sellers:0","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"orders:0000","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"sellers:0","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0000","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0000":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0000":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0000","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0000"},{"amount":99,"asset":"USD/2","destination":"sellers:0","source":"orders:0000"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0000":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0000":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-0","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"orders:0000","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0000","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0000","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:0","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0000","fees","sellers:0"],"destinations_arrays":[{"0":"orders","1":"0000","2":null},{"0":"fees","1":null},{"0":"sellers","1":"0","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0000":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0000\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0000\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0000\",\"destination\":\"sellers:0\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0000","orders:0000"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0000","2":null},{"0":"orders","1":"0000","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

2. Findings [2 35 66]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"orders:0001","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"sellers:1","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0001","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"sellers:1","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0001":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0001":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0001","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0001"},{"amount":99,"asset":"USD/2","destination":"sellers:1","source":"orders:0001"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0001":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0001":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-1","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0001","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0001","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0001","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:1","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0001","fees","sellers:1"],"destinations_arrays":[{"0":"orders","1":"0001","2":null},{"0":"fees","1":null},{"0":"sellers","1":"1","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0001":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0001\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0001\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0001\",\"destination\":\"sellers:1\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0001","orders:0001"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0001","2":null},{"0":"orders","1":"0001","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

3. Findings [3 36 67]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"sellers:2","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"orders:0002","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"sellers:2","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0002","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0002":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0002":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0002","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0002"},{"amount":99,"asset":"USD/2","destination":"sellers:2","source":"orders:0002"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0002":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0002":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-2","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"orders:0002","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:2","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0002","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0002","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0002","fees","sellers:2"],"destinations_arrays":[{"0":"orders","1":"0002","2":null},{"0":"fees","1":null},{"0":"sellers","1":"2","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0002":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0002\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0002\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0002\",\"destination\":\"sellers:2\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0002","orders:0002"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0002","2":null},{"0":"orders","1":"0002","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

4. Findings [4 37 68]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"orders:0003","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes {"accounts_address":"sellers:3","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0003","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"sellers:3","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0003":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0003":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0003","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0003"},{"amount":99,"asset":"USD/2","destination":"sellers:3","source":"orders:0003"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0003":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0003":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-3","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"orders:0003","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0003","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0003","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:3","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0003","fees","sellers:3"],"destinations_arrays":[{"0":"orders","1":"0003","2":null},{"0":"fees","1":null},{"0":"sellers","1":"3","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0003":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0003\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0003\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0003\",\"destination\":\"sellers:3\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0003","orders:0003"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0003","2":null},{"0":"orders","1":"0003","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

5. Findings [5 38 69]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"orders:0004","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes {"accounts_address":"sellers:4","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"sellers:4","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0004","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0004":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0004":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0004","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0004"},{"amount":99,"asset":"USD/2","destination":"sellers:4","source":"orders:0004"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0004":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0004":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-4","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"orders:0004","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0004","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0004","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:4","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0004","fees","sellers:4"],"destinations_arrays":[{"0":"orders","1":"0004","2":null},{"0":"fees","1":null},{"0":"sellers","1":"4","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0004":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0004\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0004\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0004\",\"destination\":\"sellers:4\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0004","orders:0004"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0004","2":null},{"0":"orders","1":"0004","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

6. Findings [6 39 70]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"sellers:0","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"orders:0005","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"sellers:0","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0005","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0005":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0005":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0005","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0005"},{"amount":99,"asset":"USD/2","destination":"sellers:0","source":"orders:0005"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0005":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0005":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-5","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0005","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0005","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0005","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:0","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0005","fees","sellers:0"],"destinations_arrays":[{"0":"orders","1":"0005","2":null},{"0":"fees","1":null},{"0":"sellers","1":"0","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0005":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0005\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0005\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0005\",\"destination\":\"sellers:0\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0005","orders:0005"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0005","2":null},{"0":"orders","1":"0005","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

7. Findings [7 40 71]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"sellers:1","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"orders:0006","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes {"accounts_address":"orders:0006","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"sellers:1","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0006":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0006":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0006","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0006"},{"amount":99,"asset":"USD/2","destination":"sellers:1","source":"orders:0006"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0006":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0006":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-6","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"orders:0006","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0006","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:1","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0006","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0006","fees","sellers:1"],"destinations_arrays":[{"0":"orders","1":"0006","2":null},{"0":"fees","1":null},{"0":"sellers","1":"1","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0006":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0006\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0006\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0006\",\"destination\":\"sellers:1\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0006","orders:0006"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0006","2":null},{"0":"orders","1":"0006","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

8. Findings [8 41 72]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"orders:0007","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"sellers:2","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0007","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"sellers:2","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0007":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0007":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0007","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0007"},{"amount":99,"asset":"USD/2","destination":"sellers:2","source":"orders:0007"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0007":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0007":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-7","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"orders:0007","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:2","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0007","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0007","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0007","fees","sellers:2"],"destinations_arrays":[{"0":"orders","1":"0007","2":null},{"0":"fees","1":null},{"0":"sellers","1":"2","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0007":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0007\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0007\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0007\",\"destination\":\"sellers:2\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0007","orders:0007"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0007","2":null},{"0":"orders","1":"0007","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

9. Findings [9 42 73]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"orders:0008","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"sellers:3","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"sellers:3","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0008","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0008":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0008":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0008","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0008"},{"amount":99,"asset":"USD/2","destination":"sellers:3","source":"orders:0008"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0008":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0008":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-8","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"orders:0008","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:3","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0008","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0008","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0008","fees","sellers:3"],"destinations_arrays":[{"0":"orders","1":"0008","2":null},{"0":"fees","1":null},{"0":"sellers","1":"3","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0008":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0008\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0008\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0008\",\"destination\":\"sellers:3\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0008","orders:0008"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0008","2":null},{"0":"orders","1":"0008","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

10. Findings [10 43 74]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"sellers:4","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"orders:0009","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes {"accounts_address":"sellers:4","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0009","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0009":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0009":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0009","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0009"},{"amount":99,"asset":"USD/2","destination":"sellers:4","source":"orders:0009"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0009":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0009":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-9","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"orders:0009","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0009","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:4","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0009","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0009","fees","sellers:4"],"destinations_arrays":[{"0":"orders","1":"0009","2":null},{"0":"fees","1":null},{"0":"sellers","1":"4","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0009":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0009\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0009\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0009\",\"destination\":\"sellers:4\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0009","orders:0009"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0009","2":null},{"0":"orders","1":"0009","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

11. Findings [11 44 75]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"sellers:0","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"orders:0010","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes {"accounts_address":"sellers:0","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0010","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0010":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0010":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0010","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0010"},{"amount":99,"asset":"USD/2","destination":"sellers:0","source":"orders:0010"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0010":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0010":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-10","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0010","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0010","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0010","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:0","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0010","fees","sellers:0"],"destinations_arrays":[{"0":"orders","1":"0010","2":null},{"0":"fees","1":null},{"0":"sellers","1":"0","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0010":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0010\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0010\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0010\",\"destination\":\"sellers:0\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0010","orders:0010"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0010","2":null},{"0":"orders","1":"0010","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

12. Findings [12 45 76]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"orders:0011","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes {"accounts_address":"sellers:1","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0011","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"sellers:1","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0011":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0011":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0011","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0011"},{"amount":99,"asset":"USD/2","destination":"sellers:1","source":"orders:0011"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0011":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0011":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-11","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"orders:0011","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0011","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0011","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:1","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0011","fees","sellers:1"],"destinations_arrays":[{"0":"orders","1":"0011","2":null},{"0":"fees","1":null},{"0":"sellers","1":"1","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0011":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0011\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0011\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0011\",\"destination\":\"sellers:1\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0011","orders:0011"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0011","2":null},{"0":"orders","1":"0011","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

13. Findings [13 46 77]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"orders:0012","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"sellers:2","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0012","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"sellers:2","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0012":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0012":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0012","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0012"},{"amount":99,"asset":"USD/2","destination":"sellers:2","source":"orders:0012"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0012":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0012":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-12","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"sellers:2","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0012","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0012","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0012","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0012","fees","sellers:2"],"destinations_arrays":[{"0":"orders","1":"0012","2":null},{"0":"fees","1":null},{"0":"sellers","1":"2","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0012":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0012\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0012\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0012\",\"destination\":\"sellers:2\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0012","orders:0012"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0012","2":null},{"0":"orders","1":"0012","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

14. Findings [14 47 78]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"orders:0013","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes {"accounts_address":"sellers:3","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"sellers:3","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
- _default.accounts_volumes {"accounts_address":"orders:0013","asset":"USD/2","input":100,"ledger":"default","output":100}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0013":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0013":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0013","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0013"},{"amount":99,"asset":"USD/2","destination":"sellers:3","source":"orders:0013"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0013":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0013":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-13","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"orders:0013","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0013","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0013","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:3","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0013","fees","sellers:3"],"destinations_arrays":[{"0":"orders","1":"0013","2":null},{"0":"fees","1":null},{"0":"sellers","1":"3","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0013":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0013\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0013\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0013\",\"destination\":\"sellers:3\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0013","orders:0013"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0013","2":null},{"0":"orders","1":"0013","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

15. Findings [15 48 79]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"sellers:4","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"orders:0014","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"sellers:4","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0014","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0014":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0014":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0014","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0014"},{"amount":99,"asset":"USD/2","destination":"sellers:4","source":"orders:0014"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0014":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0014":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-14","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"orders:0014","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0014","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0014","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:4","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0014","fees","sellers:4"],"destinations_arrays":[{"0":"orders","1":"0014","2":null},{"0":"fees","1":null},{"0":"sellers","1":"4","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0014":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0014\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0014\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0014\",\"destination\":\"sellers:4\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0014","orders:0014"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0014","2":null},{"0":"orders","1":"0014","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

16. Findings [16 49 80]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"orders:0015","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes {"accounts_address":"sellers:0","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"sellers:0","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0015","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0015":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0015":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0015","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0015"},{"amount":99,"asset":"USD/2","destination":"sellers:0","source":"orders:0015"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0015":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0015":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-15","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0015","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0015","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0015","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:0","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0015","fees","sellers:0"],"destinations_arrays":[{"0":"orders","1":"0015","2":null},{"0":"fees","1":null},{"0":"sellers","1":"0","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0015":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0015\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0015\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0015\",\"destination\":\"sellers:0\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0015","orders:0015"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0015","2":null},{"0":"orders","1":"0015","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

17. Findings [17 50 81]:

```
+ _default.accounts_volumes {"accounts_address":"orders:0016","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"sellers:1","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"sellers:1","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0016","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0016":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0016":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0016","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0016"},{"amount":99,"asset":"USD/2","destination":"sellers:1","source":"orders:0016"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0016":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0016":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-16","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0016","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0016","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0016","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:1","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0016","fees","sellers:1"],"destinations_arrays":[{"0":"orders","1":"0016","2":null},{"0":"fees","1":null},{"0":"sellers","1":"1","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0016":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0016\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0016\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0016\",\"destination\":\"sellers:1\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0016","orders:0016"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0016","2":null},{"0":"orders","1":"0016","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

18. Findings [18 51 82]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"sellers:2","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"orders:0017","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes {"accounts_address":"orders:0017","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"sellers:2","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0017":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0017":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0017","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0017"},{"amount":99,"asset":"USD/2","destination":"sellers:2","source":"orders:0017"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0017":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0017":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-17","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"sellers:2","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0017","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0017","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0017","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0017","fees","sellers:2"],"destinations_arrays":[{"0":"orders","1":"0017","2":null},{"0":"fees","1":null},{"0":"sellers","1":"2","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0017":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0017\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0017\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0017\",\"destination\":\"sellers:2\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0017","orders:0017"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0017","2":null},{"0":"orders","1":"0017","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

19. Findings [19 52 83]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"orders:0018","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes {"accounts_address":"sellers:3","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0018","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"sellers:3","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0018":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0018":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0018","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0018"},{"amount":99,"asset":"USD/2","destination":"sellers:3","source":"orders:0018"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0018":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0018":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-18","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0018","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:3","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0018","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0018","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0018","fees","sellers:3"],"destinations_arrays":[{"0":"orders","1":"0018","2":null},{"0":"fees","1":null},{"0":"sellers","1":"3","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0018":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0018\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0018\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0018\",\"destination\":\"sellers:3\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0018","orders:0018"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0018","2":null},{"0":"orders","1":"0018","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

20. Findings [20 53 84]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"sellers:4","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"orders:0019","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0019","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"sellers:4","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0019":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0019":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0019","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0019"},{"amount":99,"asset":"USD/2","destination":"sellers:4","source":"orders:0019"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0019":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0019":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-19","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"orders:0019","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0019","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:4","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0019","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0019","fees","sellers:4"],"destinations_arrays":[{"0":"orders","1":"0019","2":null},{"0":"fees","1":null},{"0":"sellers","1":"4","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0019":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0019\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0019\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0019\",\"destination\":\"sellers:4\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0019","orders:0019"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0019","2":null},{"0":"orders","1":"0019","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

21. Findings [21 54 85]:

```
+ _default.accounts_volumes {"accounts_address":"orders:0020","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"sellers:0","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"sellers:0","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0020","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0020":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0020":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0020","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0020"},{"amount":99,"asset":"USD/2","destination":"sellers:0","source":"orders:0020"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0020":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0020":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-20","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"orders:0020","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0020","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0020","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:0","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0020","fees","sellers:0"],"destinations_arrays":[{"0":"orders","1":"0020","2":null},{"0":"fees","1":null},{"0":"sellers","1":"0","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0020":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0020\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0020\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0020\",\"destination\":\"sellers:0\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0020","orders:0020"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0020","2":null},{"0":"orders","1":"0020","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

22. Findings [22 55 86]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"orders:0021","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"sellers:1","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0021","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"sellers:1","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0021":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0021":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0021","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0021"},{"amount":99,"asset":"USD/2","destination":"sellers:1","source":"orders:0021"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0021":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0021":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-21","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"orders:0021","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0021","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:1","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0021","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0021","fees","sellers:1"],"destinations_arrays":[{"0":"orders","1":"0021","2":null},{"0":"fees","1":null},{"0":"sellers","1":"1","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0021":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0021\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0021\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0021\",\"destination\":\"sellers:1\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0021","orders:0021"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0021","2":null},{"0":"orders","1":"0021","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

23. Findings [23 56 87]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"orders:0022","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"sellers:2","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0022","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"sellers:2","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0022":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0022":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0022","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0022"},{"amount":99,"asset":"USD/2","destination":"sellers:2","source":"orders:0022"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0022":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0022":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-22","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"sellers:2","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0022","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0022","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0022","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0022","fees","sellers:2"],"destinations_arrays":[{"0":"orders","1":"0022","2":null},{"0":"fees","1":null},{"0":"sellers","1":"2","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0022":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0022\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0022\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0022\",\"destination\":\"sellers:2\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0022","orders:0022"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0022","2":null},{"0":"orders","1":"0022","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

24. Findings [24 57 88]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"orders:0023","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes {"accounts_address":"sellers:3","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0023","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"sellers:3","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0023":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0023":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0023","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0023"},{"amount":99,"asset":"USD/2","destination":"sellers:3","source":"orders:0023"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0023":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0023":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-23","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0023","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0023","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:3","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0023","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0023","fees","sellers:3"],"destinations_arrays":[{"0":"orders","1":"0023","2":null},{"0":"fees","1":null},{"0":"sellers","1":"3","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0023":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0023\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0023\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0023\",\"destination\":\"sellers:3\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0023","orders:0023"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0023","2":null},{"0":"orders","1":"0023","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

25. Findings [25 58 89]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"sellers:4","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"orders:0024","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes {"accounts_address":"sellers:4","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0024","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0024":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0024":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0024","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0024"},{"amount":99,"asset":"USD/2","destination":"sellers:4","source":"orders:0024"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0024":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0024":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-24","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0024","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0024","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:4","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0024","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0024","fees","sellers:4"],"destinations_arrays":[{"0":"orders","1":"0024","2":null},{"0":"fees","1":null},{"0":"sellers","1":"4","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0024":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0024\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0024\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0024\",\"destination\":\"sellers:4\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0024","orders:0024"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0024","2":null},{"0":"orders","1":"0024","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

26. Findings [26 59 90]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"sellers:0","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"orders:0025","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes {"accounts_address":"sellers:0","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0025","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0025":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0025":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0025","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0025"},{"amount":99,"asset":"USD/2","destination":"sellers:0","source":"orders:0025"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0025":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0025":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-25","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"orders:0025","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0025","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0025","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:0","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0025","fees","sellers:0"],"destinations_arrays":[{"0":"orders","1":"0025","2":null},{"0":"fees","1":null},{"0":"sellers","1":"0","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0025":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0025\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0025\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0025\",\"destination\":\"sellers:0\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0025","orders:0025"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0025","2":null},{"0":"orders","1":"0025","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

27. Findings [27 60 91]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"sellers:1","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"orders:0026","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"sellers:1","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0026","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0026":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0026":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0026","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0026"},{"amount":99,"asset":"USD/2","destination":"sellers:1","source":"orders:0026"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0026":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0026":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-26","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"orders:0026","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0026","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0026","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:1","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0026","fees","sellers:1"],"destinations_arrays":[{"0":"orders","1":"0026","2":null},{"0":"fees","1":null},{"0":"sellers","1":"1","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0026":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0026\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0026\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0026\",\"destination\":\"sellers:1\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0026","orders:0026"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0026","2":null},{"0":"orders","1":"0026","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

28. Findings [28 61 92]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"sellers:2","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"orders:0027","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"sellers:2","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0027","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0027":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0027":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0027","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0027"},{"amount":99,"asset":"USD/2","destination":"sellers:2","source":"orders:0027"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0027":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0027":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-27","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"sellers:2","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0027","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0027","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0027","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0027","fees","sellers:2"],"destinations_arrays":[{"0":"orders","1":"0027","2":null},{"0":"fees","1":null},{"0":"sellers","1":"2","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0027":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0027\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0027\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0027\",\"destination\":\"sellers:2\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0027","orders:0027"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0027","2":null},{"0":"orders","1":"0027","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

29. Findings [29 62 93]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"orders:0028","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"sellers:3","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0028","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"sellers:3","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0028":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0028":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0028","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0028"},{"amount":99,"asset":"USD/2","destination":"sellers:3","source":"orders:0028"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0028":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0028":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-28","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"orders:0028","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0028","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:3","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0028","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0028","fees","sellers:3"],"destinations_arrays":[{"0":"orders","1":"0028","2":null},{"0":"fees","1":null},{"0":"sellers","1":"3","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0028":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0028\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0028\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0028\",\"destination\":\"sellers:3\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0028","orders:0028"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0028","2":null},{"0":"orders","1":"0028","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

30. Findings [30 63 94]:

```
+ _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes {"accounts_address":"orders:0029","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes {"accounts_address":"sellers:4","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"sellers:4","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"orders:0029","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0029":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0029":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0029","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0029"},{"amount":99,"asset":"USD/2","destination":"sellers:4","source":"orders:0029"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0029":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0029":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-29","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves {"accounts_address":"orders:0029","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0029","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"orders:0029","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"sellers:4","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions {"destinations":["orders:0029","fees","sellers:4"],"destinations_arrays":[{"0":"orders","1":"0029","2":null},{"0":"fees","1":null},{"0":"sellers","1":"4","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0029":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0029\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0029\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0029\",\"destination\":\"sellers:4\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0029","orders:0029"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0029","2":null},{"0":"orders","1":"0029","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

31. Findings [31 95 104]:

```
+ _default.logs {"data":{"metadata":{"tier":"gold-0"},"targetType":"ACCOUNT"},"idempotency_key":"am-0","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

32. Findings [32 96 105]:

```
+ _default.logs {"data":{"metadata":{"tier":"gold-1"},"targetType":"ACCOUNT"},"idempotency_key":"am-1","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

33. Findings [33 64 97]:

```
+ _default.logs {"data":{"metadata":{"tier":"gold-2"},"targetType":"ACCOUNT"},"idempotency_key":"am-2","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

34. Findings [98 101 106]:

```
+ _default.logs {"data":{"metadata":{"note":"checked-0"},"targetType":"TRANSACTION"},"idempotency_key":"md-0","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

35. Findings [99 102 107]:

```
+ _default.logs {"data":{"metadata":{"note":"checked-3"},"targetType":"TRANSACTION"},"idempotency_key":"md-3","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

36. Findings [100 103 108]:

```
+ _default.logs {"data":{"metadata":{"note":"checked-4"},"targetType":"TRANSACTION"},"idempotency_key":"md-4","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

37. Findings [109 110 111]:

```
+ _default.logs {"data":{"metadata":{"note":"checked-1"},"targetType":"TRANSACTION"},"idempotency_key":"md-1","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

38. Findings [112 113 114]:

```
+ _default.logs {"data":{"metadata":{"note":"checked-2"},"targetType":"TRANSACTION"},"idempotency_key":"md-2","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

## Findings: 114 distinct minimal reproductions

### Finding 1

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#0).

```
#0 POST /v2/default/transactions key=tx-0
fault lost(#0)
```

Delivery order: send-twice #0

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0000","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0000","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0000":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0000":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0000","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0000"},{"amount":99,"asset":"USD/2","destination":"sellers:0","source":"orders:0000"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0000":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0000":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-0","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0000","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0000","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0000","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:0","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0000","fees","sellers:0"],"destinations_arrays":[{"0":"orders","1":"0000","2":null},{"0":"fees","1":null},{"0":"sellers","1":"0","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0000":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0000\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0000\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0000\",\"destination\":\"sellers:0\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0000","orders:0000"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0000","2":null},{"0":"orders","1":"0000","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-1/repro.sh`

### Finding 2

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#1).

```
#1 POST /v2/default/transactions key=tx-1
fault lost(#1)
```

Delivery order: send-twice #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0001","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0001","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0001":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0001":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0001","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0001"},{"amount":99,"asset":"USD/2","destination":"sellers:1","source":"orders:0001"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0001":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0001":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-1","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0001","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0001","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0001","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:1","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0001","fees","sellers:1"],"destinations_arrays":[{"0":"orders","1":"0001","2":null},{"0":"fees","1":null},{"0":"sellers","1":"1","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0001":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0001\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0001\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0001\",\"destination\":\"sellers:1\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0001","orders:0001"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0001","2":null},{"0":"orders","1":"0001","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-2/repro.sh`

### Finding 3

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#2).

```
#2 POST /v2/default/transactions key=tx-2
fault lost(#2)
```

Delivery order: send-twice #2

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0002","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0002","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0002":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0002":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0002","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0002"},{"amount":99,"asset":"USD/2","destination":"sellers:2","source":"orders:0002"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0002":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0002":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-2","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0002","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:2","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0002","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0002","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0002","fees","sellers:2"],"destinations_arrays":[{"0":"orders","1":"0002","2":null},{"0":"fees","1":null},{"0":"sellers","1":"2","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0002":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0002\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0002\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0002\",\"destination\":\"sellers:2\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0002","orders:0002"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0002","2":null},{"0":"orders","1":"0002","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-3/repro.sh`

### Finding 4

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#3).

```
#3 POST /v2/default/transactions key=tx-3
fault lost(#3)
```

Delivery order: send-twice #3

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0003","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0003","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0003":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0003":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0003","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0003"},{"amount":99,"asset":"USD/2","destination":"sellers:3","source":"orders:0003"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0003":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0003":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-3","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0003","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0003","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0003","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:3","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0003","fees","sellers:3"],"destinations_arrays":[{"0":"orders","1":"0003","2":null},{"0":"fees","1":null},{"0":"sellers","1":"3","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0003":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0003\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0003\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0003\",\"destination\":\"sellers:3\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0003","orders:0003"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0003","2":null},{"0":"orders","1":"0003","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-4/repro.sh`

### Finding 5

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#4).

```
#4 POST /v2/default/transactions key=tx-4
fault lost(#4)
```

Delivery order: send-twice #4

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0004","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0004","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0004":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0004":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0004","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0004"},{"amount":99,"asset":"USD/2","destination":"sellers:4","source":"orders:0004"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0004":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0004":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-4","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0004","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0004","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0004","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:4","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0004","fees","sellers:4"],"destinations_arrays":[{"0":"orders","1":"0004","2":null},{"0":"fees","1":null},{"0":"sellers","1":"4","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0004":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0004\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0004\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0004\",\"destination\":\"sellers:4\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0004","orders:0004"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0004","2":null},{"0":"orders","1":"0004","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-5/repro.sh`

### Finding 6

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#5).

```
#5 POST /v2/default/transactions key=tx-5
fault lost(#5)
```

Delivery order: send-twice #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0005","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0005","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0005":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0005":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0005","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0005"},{"amount":99,"asset":"USD/2","destination":"sellers:0","source":"orders:0005"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0005":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0005":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-5","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0005","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0005","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0005","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:0","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0005","fees","sellers:0"],"destinations_arrays":[{"0":"orders","1":"0005","2":null},{"0":"fees","1":null},{"0":"sellers","1":"0","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0005":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0005\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0005\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0005\",\"destination\":\"sellers:0\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0005","orders:0005"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0005","2":null},{"0":"orders","1":"0005","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-6/repro.sh`

### Finding 7

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#6).

```
#6 POST /v2/default/transactions key=tx-6
fault lost(#6)
```

Delivery order: send-twice #6

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0006","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"orders:0006","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0006":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0006":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0006","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0006"},{"amount":99,"asset":"USD/2","destination":"sellers:1","source":"orders:0006"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0006":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0006":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-6","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0006","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0006","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:1","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0006","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0006","fees","sellers:1"],"destinations_arrays":[{"0":"orders","1":"0006","2":null},{"0":"fees","1":null},{"0":"sellers","1":"1","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0006":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0006\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0006\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0006\",\"destination\":\"sellers:1\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0006","orders:0006"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0006","2":null},{"0":"orders","1":"0006","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-7/repro.sh`

### Finding 8

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#7).

```
#7 POST /v2/default/transactions key=tx-7
fault lost(#7)
```

Delivery order: send-twice #7

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0007","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0007","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0007":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0007":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0007","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0007"},{"amount":99,"asset":"USD/2","destination":"sellers:2","source":"orders:0007"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0007":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0007":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-7","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0007","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:2","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0007","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0007","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0007","fees","sellers:2"],"destinations_arrays":[{"0":"orders","1":"0007","2":null},{"0":"fees","1":null},{"0":"sellers","1":"2","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0007":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0007\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0007\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0007\",\"destination\":\"sellers:2\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0007","orders:0007"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0007","2":null},{"0":"orders","1":"0007","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-8/repro.sh`

### Finding 9

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#8).

```
#8 POST /v2/default/transactions key=tx-8
fault lost(#8)
```

Delivery order: send-twice #8

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0008","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0008","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0008":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0008":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0008","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0008"},{"amount":99,"asset":"USD/2","destination":"sellers:3","source":"orders:0008"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0008":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0008":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-8","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0008","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:3","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0008","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0008","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0008","fees","sellers:3"],"destinations_arrays":[{"0":"orders","1":"0008","2":null},{"0":"fees","1":null},{"0":"sellers","1":"3","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0008":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0008\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0008\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0008\",\"destination\":\"sellers:3\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0008","orders:0008"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0008","2":null},{"0":"orders","1":"0008","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-9/repro.sh`

### Finding 10

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#9).

```
#9 POST /v2/default/transactions key=tx-9
fault lost(#9)
```

Delivery order: send-twice #9

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0009","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0009","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0009":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0009":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0009","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0009"},{"amount":99,"asset":"USD/2","destination":"sellers:4","source":"orders:0009"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0009":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0009":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-9","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0009","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0009","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:4","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0009","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0009","fees","sellers:4"],"destinations_arrays":[{"0":"orders","1":"0009","2":null},{"0":"fees","1":null},{"0":"sellers","1":"4","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0009":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0009\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0009\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0009\",\"destination\":\"sellers:4\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0009","orders:0009"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0009","2":null},{"0":"orders","1":"0009","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-10/repro.sh`

### Finding 11

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#10).

```
#10 POST /v2/default/transactions key=tx-10
fault lost(#10)
```

Delivery order: send-twice #10

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0010","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0010","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0010":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0010":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0010","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0010"},{"amount":99,"asset":"USD/2","destination":"sellers:0","source":"orders:0010"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0010":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0010":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-10","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0010","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0010","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0010","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:0","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0010","fees","sellers:0"],"destinations_arrays":[{"0":"orders","1":"0010","2":null},{"0":"fees","1":null},{"0":"sellers","1":"0","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0010":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0010\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0010\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0010\",\"destination\":\"sellers:0\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0010","orders:0010"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0010","2":null},{"0":"orders","1":"0010","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-11/repro.sh`

### Finding 12

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#11).

```
#11 POST /v2/default/transactions key=tx-11
fault lost(#11)
```

Delivery order: send-twice #11

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0011","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0011","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0011":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0011":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0011","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0011"},{"amount":99,"asset":"USD/2","destination":"sellers:1","source":"orders:0011"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0011":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0011":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-11","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0011","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0011","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0011","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:1","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0011","fees","sellers:1"],"destinations_arrays":[{"0":"orders","1":"0011","2":null},{"0":"fees","1":null},{"0":"sellers","1":"1","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0011":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0011\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0011\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0011\",\"destination\":\"sellers:1\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0011","orders:0011"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0011","2":null},{"0":"orders","1":"0011","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-12/repro.sh`

### Finding 13

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#12).

```
#12 POST /v2/default/transactions key=tx-12
fault lost(#12)
```

Delivery order: send-twice #12

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0012","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0012","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0012":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0012":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0012","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0012"},{"amount":99,"asset":"USD/2","destination":"sellers:2","source":"orders:0012"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0012":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0012":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-12","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"sellers:2","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0012","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0012","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0012","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0012","fees","sellers:2"],"destinations_arrays":[{"0":"orders","1":"0012","2":null},{"0":"fees","1":null},{"0":"sellers","1":"2","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0012":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0012\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0012\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0012\",\"destination\":\"sellers:2\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0012","orders:0012"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0012","2":null},{"0":"orders","1":"0012","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-13/repro.sh`

### Finding 14

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#13).

```
#13 POST /v2/default/transactions key=tx-13
fault lost(#13)
```

Delivery order: send-twice #13

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0013","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
- _default.accounts_volumes x1 {"accounts_address":"orders:0013","asset":"USD/2","input":100,"ledger":"default","output":100}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0013":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0013":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0013","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0013"},{"amount":99,"asset":"USD/2","destination":"sellers:3","source":"orders:0013"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0013":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0013":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-13","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0013","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0013","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0013","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:3","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0013","fees","sellers:3"],"destinations_arrays":[{"0":"orders","1":"0013","2":null},{"0":"fees","1":null},{"0":"sellers","1":"3","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0013":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0013\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0013\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0013\",\"destination\":\"sellers:3\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0013","orders:0013"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0013","2":null},{"0":"orders","1":"0013","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-14/repro.sh`

### Finding 15

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#14).

```
#14 POST /v2/default/transactions key=tx-14
fault lost(#14)
```

Delivery order: send-twice #14

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0014","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0014","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0014":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0014":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0014","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0014"},{"amount":99,"asset":"USD/2","destination":"sellers:4","source":"orders:0014"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0014":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0014":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-14","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0014","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0014","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0014","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:4","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0014","fees","sellers:4"],"destinations_arrays":[{"0":"orders","1":"0014","2":null},{"0":"fees","1":null},{"0":"sellers","1":"4","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0014":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0014\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0014\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0014\",\"destination\":\"sellers:4\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0014","orders:0014"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0014","2":null},{"0":"orders","1":"0014","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-15/repro.sh`

### Finding 16

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#15).

```
#15 POST /v2/default/transactions key=tx-15
fault lost(#15)
```

Delivery order: send-twice #15

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0015","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0015","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0015":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0015":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0015","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0015"},{"amount":99,"asset":"USD/2","destination":"sellers:0","source":"orders:0015"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0015":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0015":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-15","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0015","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0015","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0015","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:0","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0015","fees","sellers:0"],"destinations_arrays":[{"0":"orders","1":"0015","2":null},{"0":"fees","1":null},{"0":"sellers","1":"0","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0015":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0015\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0015\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0015\",\"destination\":\"sellers:0\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0015","orders:0015"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0015","2":null},{"0":"orders","1":"0015","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-16/repro.sh`

### Finding 17

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#16).

```
#16 POST /v2/default/transactions key=tx-16
fault lost(#16)
```

Delivery order: send-twice #16

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"orders:0016","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0016","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0016":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0016":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0016","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0016"},{"amount":99,"asset":"USD/2","destination":"sellers:1","source":"orders:0016"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0016":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0016":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-16","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0016","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0016","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0016","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:1","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0016","fees","sellers:1"],"destinations_arrays":[{"0":"orders","1":"0016","2":null},{"0":"fees","1":null},{"0":"sellers","1":"1","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0016":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0016\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0016\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0016\",\"destination\":\"sellers:1\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0016","orders:0016"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0016","2":null},{"0":"orders","1":"0016","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-17/repro.sh`

### Finding 18

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#17).

```
#17 POST /v2/default/transactions key=tx-17
fault lost(#17)
```

Delivery order: send-twice #17

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0017","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"orders:0017","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0017":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0017":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0017","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0017"},{"amount":99,"asset":"USD/2","destination":"sellers:2","source":"orders:0017"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0017":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0017":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-17","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"sellers:2","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0017","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0017","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0017","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0017","fees","sellers:2"],"destinations_arrays":[{"0":"orders","1":"0017","2":null},{"0":"fees","1":null},{"0":"sellers","1":"2","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0017":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0017\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0017\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0017\",\"destination\":\"sellers:2\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0017","orders:0017"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0017","2":null},{"0":"orders","1":"0017","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-18/repro.sh`

### Finding 19

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#18).

```
#18 POST /v2/default/transactions key=tx-18
fault lost(#18)
```

Delivery order: send-twice #18

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0018","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0018","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0018":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0018":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0018","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0018"},{"amount":99,"asset":"USD/2","destination":"sellers:3","source":"orders:0018"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0018":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0018":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-18","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0018","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:3","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0018","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0018","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0018","fees","sellers:3"],"destinations_arrays":[{"0":"orders","1":"0018","2":null},{"0":"fees","1":null},{"0":"sellers","1":"3","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0018":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0018\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0018\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0018\",\"destination\":\"sellers:3\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0018","orders:0018"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0018","2":null},{"0":"orders","1":"0018","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-19/repro.sh`

### Finding 20

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#19).

```
#19 POST /v2/default/transactions key=tx-19
fault lost(#19)
```

Delivery order: send-twice #19

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0019","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0019","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0019":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0019":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0019","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0019"},{"amount":99,"asset":"USD/2","destination":"sellers:4","source":"orders:0019"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0019":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0019":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-19","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0019","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0019","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:4","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0019","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0019","fees","sellers:4"],"destinations_arrays":[{"0":"orders","1":"0019","2":null},{"0":"fees","1":null},{"0":"sellers","1":"4","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0019":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0019\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0019\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0019\",\"destination\":\"sellers:4\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0019","orders:0019"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0019","2":null},{"0":"orders","1":"0019","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-20/repro.sh`

### Finding 21

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#20).

```
#20 POST /v2/default/transactions key=tx-20
fault lost(#20)
```

Delivery order: send-twice #20

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"orders:0020","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0020","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0020":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0020":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0020","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0020"},{"amount":99,"asset":"USD/2","destination":"sellers:0","source":"orders:0020"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0020":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0020":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-20","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0020","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0020","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0020","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:0","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0020","fees","sellers:0"],"destinations_arrays":[{"0":"orders","1":"0020","2":null},{"0":"fees","1":null},{"0":"sellers","1":"0","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0020":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0020\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0020\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0020\",\"destination\":\"sellers:0\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0020","orders:0020"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0020","2":null},{"0":"orders","1":"0020","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-21/repro.sh`

### Finding 22

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#21).

```
#21 POST /v2/default/transactions key=tx-21
fault lost(#21)
```

Delivery order: send-twice #21

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0021","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0021","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0021":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0021":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0021","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0021"},{"amount":99,"asset":"USD/2","destination":"sellers:1","source":"orders:0021"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0021":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0021":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-21","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0021","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0021","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:1","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0021","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0021","fees","sellers:1"],"destinations_arrays":[{"0":"orders","1":"0021","2":null},{"0":"fees","1":null},{"0":"sellers","1":"1","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0021":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0021\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0021\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0021\",\"destination\":\"sellers:1\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0021","orders:0021"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0021","2":null},{"0":"orders","1":"0021","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-22/repro.sh`

### Finding 23

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#22).

```
#22 POST /v2/default/transactions key=tx-22
fault lost(#22)
```

Delivery order: send-twice #22

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0022","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0022","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0022":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0022":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0022","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0022"},{"amount":99,"asset":"USD/2","destination":"sellers:2","source":"orders:0022"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0022":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0022":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-22","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"sellers:2","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0022","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0022","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0022","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0022","fees","sellers:2"],"destinations_arrays":[{"0":"orders","1":"0022","2":null},{"0":"fees","1":null},{"0":"sellers","1":"2","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0022":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0022\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0022\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0022\",\"destination\":\"sellers:2\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0022","orders:0022"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0022","2":null},{"0":"orders","1":"0022","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-23/repro.sh`

### Finding 24

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#23).

```
#23 POST /v2/default/transactions key=tx-23
fault lost(#23)
```

Delivery order: send-twice #23

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0023","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0023","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0023":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0023":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0023","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0023"},{"amount":99,"asset":"USD/2","destination":"sellers:3","source":"orders:0023"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0023":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0023":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-23","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0023","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0023","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:3","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0023","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0023","fees","sellers:3"],"destinations_arrays":[{"0":"orders","1":"0023","2":null},{"0":"fees","1":null},{"0":"sellers","1":"3","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0023":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0023\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0023\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0023\",\"destination\":\"sellers:3\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0023","orders:0023"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0023","2":null},{"0":"orders","1":"0023","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-24/repro.sh`

### Finding 25

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#24).

```
#24 POST /v2/default/transactions key=tx-24
fault lost(#24)
```

Delivery order: send-twice #24

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0024","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0024","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0024":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0024":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0024","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0024"},{"amount":99,"asset":"USD/2","destination":"sellers:4","source":"orders:0024"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0024":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0024":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-24","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0024","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0024","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:4","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0024","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0024","fees","sellers:4"],"destinations_arrays":[{"0":"orders","1":"0024","2":null},{"0":"fees","1":null},{"0":"sellers","1":"4","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0024":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0024\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0024\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0024\",\"destination\":\"sellers:4\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0024","orders:0024"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0024","2":null},{"0":"orders","1":"0024","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-25/repro.sh`

### Finding 26

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#25).

```
#25 POST /v2/default/transactions key=tx-25
fault lost(#25)
```

Delivery order: send-twice #25

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0025","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0025","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0025":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0025":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0025","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0025"},{"amount":99,"asset":"USD/2","destination":"sellers:0","source":"orders:0025"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0025":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0025":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-25","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0025","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0025","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0025","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:0","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0025","fees","sellers:0"],"destinations_arrays":[{"0":"orders","1":"0025","2":null},{"0":"fees","1":null},{"0":"sellers","1":"0","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0025":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0025\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0025\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0025\",\"destination\":\"sellers:0\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0025","orders:0025"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0025","2":null},{"0":"orders","1":"0025","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-26/repro.sh`

### Finding 27

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#26).

```
#26 POST /v2/default/transactions key=tx-26
fault lost(#26)
```

Delivery order: send-twice #26

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0026","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0026","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0026":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0026":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0026","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0026"},{"amount":99,"asset":"USD/2","destination":"sellers:1","source":"orders:0026"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0026":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0026":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-26","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0026","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0026","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0026","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:1","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0026","fees","sellers:1"],"destinations_arrays":[{"0":"orders","1":"0026","2":null},{"0":"fees","1":null},{"0":"sellers","1":"1","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0026":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0026\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0026\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0026\",\"destination\":\"sellers:1\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0026","orders:0026"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0026","2":null},{"0":"orders","1":"0026","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-27/repro.sh`

### Finding 28

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#27).

```
#27 POST /v2/default/transactions key=tx-27
fault lost(#27)
```

Delivery order: send-twice #27

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0027","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0027","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0027":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0027":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0027","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0027"},{"amount":99,"asset":"USD/2","destination":"sellers:2","source":"orders:0027"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0027":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0027":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-27","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"sellers:2","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0027","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0027","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0027","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0027","fees","sellers:2"],"destinations_arrays":[{"0":"orders","1":"0027","2":null},{"0":"fees","1":null},{"0":"sellers","1":"2","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0027":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0027\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0027\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0027\",\"destination\":\"sellers:2\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0027","orders:0027"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0027","2":null},{"0":"orders","1":"0027","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-28/repro.sh`

### Finding 29

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#28).

```
#28 POST /v2/default/transactions key=tx-28
fault lost(#28)
```

Delivery order: send-twice #28

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0028","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0028","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0028":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0028":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0028","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0028"},{"amount":99,"asset":"USD/2","destination":"sellers:3","source":"orders:0028"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0028":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0028":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-28","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0028","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0028","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:3","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0028","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0028","fees","sellers:3"],"destinations_arrays":[{"0":"orders","1":"0028","2":null},{"0":"fees","1":null},{"0":"sellers","1":"3","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0028":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0028\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0028\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0028\",\"destination\":\"sellers:3\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0028","orders:0028"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0028","2":null},{"0":"orders","1":"0028","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-29/repro.sh`

### Finding 30

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#29).

```
#29 POST /v2/default/transactions key=tx-29
fault lost(#29)
```

Delivery order: send-twice #29

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0029","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0029","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0029":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0029":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0029","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0029"},{"amount":99,"asset":"USD/2","destination":"sellers:4","source":"orders:0029"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0029":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0029":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-29","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0029","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0029","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0029","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:4","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0029","fees","sellers:4"],"destinations_arrays":[{"0":"orders","1":"0029","2":null},{"0":"fees","1":null},{"0":"sellers","1":"4","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0029":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0029\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0029\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0029\",\"destination\":\"sellers:4\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0029","orders:0029"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0029","2":null},{"0":"orders","1":"0029","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-30/repro.sh`

### Finding 31

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#40).

```
#40 POST /v2/default/accounts/sellers:0/metadata key=am-0
fault lost(#40)
```

Delivery order: send-twice #40

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"tier":"gold-0"},"targetType":"ACCOUNT"},"idempotency_key":"am-0","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-31/repro.sh`

### Finding 32

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#41).

```
#41 POST /v2/default/accounts/sellers:1/metadata key=am-1
fault lost(#41)
```

Delivery order: send-twice #41

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"tier":"gold-1"},"targetType":"ACCOUNT"},"idempotency_key":"am-1","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-32/repro.sh`

### Finding 33

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#42).

```
#42 POST /v2/default/accounts/sellers:2/metadata key=am-2
fault lost(#42)
```

Delivery order: send-twice #42

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"tier":"gold-2"},"targetType":"ACCOUNT"},"idempotency_key":"am-2","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-33/repro.sh`

### Finding 34

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#0).

```
#0 POST /v2/default/transactions key=tx-0
fault concurrent(#0)
```

Delivery order: send-concurrently #0

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0000","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0000","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0000":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0000":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0000","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0000"},{"amount":99,"asset":"USD/2","destination":"sellers:0","source":"orders:0000"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0000":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0000":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-0","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0000","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0000","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0000","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:0","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0000","fees","sellers:0"],"destinations_arrays":[{"0":"orders","1":"0000","2":null},{"0":"fees","1":null},{"0":"sellers","1":"0","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0000":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0000\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0000\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0000\",\"destination\":\"sellers:0\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0000","orders:0000"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0000","2":null},{"0":"orders","1":"0000","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-34/repro.sh`

### Finding 35

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#1).

```
#1 POST /v2/default/transactions key=tx-1
fault concurrent(#1)
```

Delivery order: send-concurrently #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0001","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0001","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0001":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0001":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0001","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0001"},{"amount":99,"asset":"USD/2","destination":"sellers:1","source":"orders:0001"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0001":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0001":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-1","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0001","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0001","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0001","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:1","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0001","fees","sellers:1"],"destinations_arrays":[{"0":"orders","1":"0001","2":null},{"0":"fees","1":null},{"0":"sellers","1":"1","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0001":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0001\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0001\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0001\",\"destination\":\"sellers:1\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0001","orders:0001"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0001","2":null},{"0":"orders","1":"0001","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-35/repro.sh`

### Finding 36

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#2).

```
#2 POST /v2/default/transactions key=tx-2
fault concurrent(#2)
```

Delivery order: send-concurrently #2

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0002","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0002","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0002":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0002":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0002","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0002"},{"amount":99,"asset":"USD/2","destination":"sellers:2","source":"orders:0002"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0002":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0002":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-2","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0002","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:2","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0002","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0002","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0002","fees","sellers:2"],"destinations_arrays":[{"0":"orders","1":"0002","2":null},{"0":"fees","1":null},{"0":"sellers","1":"2","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0002":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0002\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0002\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0002\",\"destination\":\"sellers:2\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0002","orders:0002"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0002","2":null},{"0":"orders","1":"0002","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-36/repro.sh`

### Finding 37

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#3).

```
#3 POST /v2/default/transactions key=tx-3
fault concurrent(#3)
```

Delivery order: send-concurrently #3

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0003","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0003","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0003":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0003":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0003","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0003"},{"amount":99,"asset":"USD/2","destination":"sellers:3","source":"orders:0003"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0003":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0003":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-3","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0003","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0003","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0003","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:3","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0003","fees","sellers:3"],"destinations_arrays":[{"0":"orders","1":"0003","2":null},{"0":"fees","1":null},{"0":"sellers","1":"3","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0003":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0003\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0003\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0003\",\"destination\":\"sellers:3\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0003","orders:0003"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0003","2":null},{"0":"orders","1":"0003","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-37/repro.sh`

### Finding 38

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#4).

```
#4 POST /v2/default/transactions key=tx-4
fault concurrent(#4)
```

Delivery order: send-concurrently #4

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0004","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0004","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0004":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0004":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0004","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0004"},{"amount":99,"asset":"USD/2","destination":"sellers:4","source":"orders:0004"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0004":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0004":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-4","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0004","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0004","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0004","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:4","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0004","fees","sellers:4"],"destinations_arrays":[{"0":"orders","1":"0004","2":null},{"0":"fees","1":null},{"0":"sellers","1":"4","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0004":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0004\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0004\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0004\",\"destination\":\"sellers:4\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0004","orders:0004"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0004","2":null},{"0":"orders","1":"0004","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-38/repro.sh`

### Finding 39

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#5).

```
#5 POST /v2/default/transactions key=tx-5
fault concurrent(#5)
```

Delivery order: send-concurrently #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0005","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0005","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0005":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0005":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0005","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0005"},{"amount":99,"asset":"USD/2","destination":"sellers:0","source":"orders:0005"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0005":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0005":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-5","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0005","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0005","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0005","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:0","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0005","fees","sellers:0"],"destinations_arrays":[{"0":"orders","1":"0005","2":null},{"0":"fees","1":null},{"0":"sellers","1":"0","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0005":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0005\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0005\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0005\",\"destination\":\"sellers:0\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0005","orders:0005"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0005","2":null},{"0":"orders","1":"0005","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-39/repro.sh`

### Finding 40

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#6).

```
#6 POST /v2/default/transactions key=tx-6
fault concurrent(#6)
```

Delivery order: send-concurrently #6

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0006","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"orders:0006","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0006":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0006":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0006","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0006"},{"amount":99,"asset":"USD/2","destination":"sellers:1","source":"orders:0006"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0006":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0006":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-6","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0006","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0006","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:1","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0006","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0006","fees","sellers:1"],"destinations_arrays":[{"0":"orders","1":"0006","2":null},{"0":"fees","1":null},{"0":"sellers","1":"1","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0006":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0006\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0006\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0006\",\"destination\":\"sellers:1\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0006","orders:0006"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0006","2":null},{"0":"orders","1":"0006","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-40/repro.sh`

### Finding 41

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#7).

```
#7 POST /v2/default/transactions key=tx-7
fault concurrent(#7)
```

Delivery order: send-concurrently #7

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0007","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0007","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0007":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0007":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0007","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0007"},{"amount":99,"asset":"USD/2","destination":"sellers:2","source":"orders:0007"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0007":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0007":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-7","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0007","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:2","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0007","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0007","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0007","fees","sellers:2"],"destinations_arrays":[{"0":"orders","1":"0007","2":null},{"0":"fees","1":null},{"0":"sellers","1":"2","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0007":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0007\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0007\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0007\",\"destination\":\"sellers:2\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0007","orders:0007"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0007","2":null},{"0":"orders","1":"0007","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-41/repro.sh`

### Finding 42

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#8).

```
#8 POST /v2/default/transactions key=tx-8
fault concurrent(#8)
```

Delivery order: send-concurrently #8

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0008","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0008","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0008":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0008":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0008","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0008"},{"amount":99,"asset":"USD/2","destination":"sellers:3","source":"orders:0008"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0008":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0008":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-8","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0008","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:3","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0008","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0008","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0008","fees","sellers:3"],"destinations_arrays":[{"0":"orders","1":"0008","2":null},{"0":"fees","1":null},{"0":"sellers","1":"3","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0008":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0008\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0008\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0008\",\"destination\":\"sellers:3\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0008","orders:0008"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0008","2":null},{"0":"orders","1":"0008","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-42/repro.sh`

### Finding 43

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#9).

```
#9 POST /v2/default/transactions key=tx-9
fault concurrent(#9)
```

Delivery order: send-concurrently #9

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0009","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0009","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0009":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0009":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0009","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0009"},{"amount":99,"asset":"USD/2","destination":"sellers:4","source":"orders:0009"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0009":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0009":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-9","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0009","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0009","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:4","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0009","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0009","fees","sellers:4"],"destinations_arrays":[{"0":"orders","1":"0009","2":null},{"0":"fees","1":null},{"0":"sellers","1":"4","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0009":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0009\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0009\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0009\",\"destination\":\"sellers:4\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0009","orders:0009"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0009","2":null},{"0":"orders","1":"0009","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-43/repro.sh`

### Finding 44

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#10).

```
#10 POST /v2/default/transactions key=tx-10
fault concurrent(#10)
```

Delivery order: send-concurrently #10

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0010","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0010","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0010":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0010":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0010","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0010"},{"amount":99,"asset":"USD/2","destination":"sellers:0","source":"orders:0010"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0010":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0010":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-10","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0010","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0010","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0010","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:0","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0010","fees","sellers:0"],"destinations_arrays":[{"0":"orders","1":"0010","2":null},{"0":"fees","1":null},{"0":"sellers","1":"0","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0010":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0010\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0010\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0010\",\"destination\":\"sellers:0\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0010","orders:0010"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0010","2":null},{"0":"orders","1":"0010","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-44/repro.sh`

### Finding 45

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#11).

```
#11 POST /v2/default/transactions key=tx-11
fault concurrent(#11)
```

Delivery order: send-concurrently #11

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0011","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0011","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0011":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0011":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0011","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0011"},{"amount":99,"asset":"USD/2","destination":"sellers:1","source":"orders:0011"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0011":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0011":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-11","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0011","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0011","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0011","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:1","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0011","fees","sellers:1"],"destinations_arrays":[{"0":"orders","1":"0011","2":null},{"0":"fees","1":null},{"0":"sellers","1":"1","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0011":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0011\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0011\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0011\",\"destination\":\"sellers:1\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0011","orders:0011"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0011","2":null},{"0":"orders","1":"0011","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-45/repro.sh`

### Finding 46

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#12).

```
#12 POST /v2/default/transactions key=tx-12
fault concurrent(#12)
```

Delivery order: send-concurrently #12

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0012","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0012","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0012":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0012":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0012","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0012"},{"amount":99,"asset":"USD/2","destination":"sellers:2","source":"orders:0012"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0012":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0012":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-12","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"sellers:2","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0012","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0012","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0012","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0012","fees","sellers:2"],"destinations_arrays":[{"0":"orders","1":"0012","2":null},{"0":"fees","1":null},{"0":"sellers","1":"2","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0012":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0012\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0012\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0012\",\"destination\":\"sellers:2\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0012","orders:0012"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0012","2":null},{"0":"orders","1":"0012","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-46/repro.sh`

### Finding 47

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#13).

```
#13 POST /v2/default/transactions key=tx-13
fault concurrent(#13)
```

Delivery order: send-concurrently #13

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0013","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
- _default.accounts_volumes x1 {"accounts_address":"orders:0013","asset":"USD/2","input":100,"ledger":"default","output":100}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0013":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0013":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0013","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0013"},{"amount":99,"asset":"USD/2","destination":"sellers:3","source":"orders:0013"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0013":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0013":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-13","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0013","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0013","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0013","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:3","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0013","fees","sellers:3"],"destinations_arrays":[{"0":"orders","1":"0013","2":null},{"0":"fees","1":null},{"0":"sellers","1":"3","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0013":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0013\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0013\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0013\",\"destination\":\"sellers:3\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0013","orders:0013"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0013","2":null},{"0":"orders","1":"0013","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-47/repro.sh`

### Finding 48

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#14).

```
#14 POST /v2/default/transactions key=tx-14
fault concurrent(#14)
```

Delivery order: send-concurrently #14

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0014","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0014","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0014":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0014":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0014","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0014"},{"amount":99,"asset":"USD/2","destination":"sellers:4","source":"orders:0014"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0014":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0014":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-14","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0014","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0014","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0014","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:4","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0014","fees","sellers:4"],"destinations_arrays":[{"0":"orders","1":"0014","2":null},{"0":"fees","1":null},{"0":"sellers","1":"4","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0014":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0014\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0014\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0014\",\"destination\":\"sellers:4\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0014","orders:0014"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0014","2":null},{"0":"orders","1":"0014","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-48/repro.sh`

### Finding 49

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#15).

```
#15 POST /v2/default/transactions key=tx-15
fault concurrent(#15)
```

Delivery order: send-concurrently #15

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0015","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0015","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0015":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0015":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0015","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0015"},{"amount":99,"asset":"USD/2","destination":"sellers:0","source":"orders:0015"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0015":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0015":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-15","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0015","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0015","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0015","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:0","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0015","fees","sellers:0"],"destinations_arrays":[{"0":"orders","1":"0015","2":null},{"0":"fees","1":null},{"0":"sellers","1":"0","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0015":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0015\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0015\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0015\",\"destination\":\"sellers:0\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0015","orders:0015"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0015","2":null},{"0":"orders","1":"0015","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-49/repro.sh`

### Finding 50

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#16).

```
#16 POST /v2/default/transactions key=tx-16
fault concurrent(#16)
```

Delivery order: send-concurrently #16

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"orders:0016","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0016","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0016":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0016":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0016","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0016"},{"amount":99,"asset":"USD/2","destination":"sellers:1","source":"orders:0016"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0016":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0016":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-16","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0016","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0016","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0016","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:1","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0016","fees","sellers:1"],"destinations_arrays":[{"0":"orders","1":"0016","2":null},{"0":"fees","1":null},{"0":"sellers","1":"1","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0016":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0016\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0016\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0016\",\"destination\":\"sellers:1\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0016","orders:0016"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0016","2":null},{"0":"orders","1":"0016","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-50/repro.sh`

### Finding 51

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#17).

```
#17 POST /v2/default/transactions key=tx-17
fault concurrent(#17)
```

Delivery order: send-concurrently #17

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0017","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"orders:0017","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0017":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0017":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0017","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0017"},{"amount":99,"asset":"USD/2","destination":"sellers:2","source":"orders:0017"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0017":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0017":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-17","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"sellers:2","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0017","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0017","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0017","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0017","fees","sellers:2"],"destinations_arrays":[{"0":"orders","1":"0017","2":null},{"0":"fees","1":null},{"0":"sellers","1":"2","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0017":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0017\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0017\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0017\",\"destination\":\"sellers:2\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0017","orders:0017"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0017","2":null},{"0":"orders","1":"0017","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-51/repro.sh`

### Finding 52

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#18).

```
#18 POST /v2/default/transactions key=tx-18
fault concurrent(#18)
```

Delivery order: send-concurrently #18

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0018","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0018","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0018":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0018":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0018","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0018"},{"amount":99,"asset":"USD/2","destination":"sellers:3","source":"orders:0018"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0018":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0018":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-18","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0018","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:3","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0018","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0018","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0018","fees","sellers:3"],"destinations_arrays":[{"0":"orders","1":"0018","2":null},{"0":"fees","1":null},{"0":"sellers","1":"3","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0018":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0018\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0018\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0018\",\"destination\":\"sellers:3\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0018","orders:0018"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0018","2":null},{"0":"orders","1":"0018","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-52/repro.sh`

### Finding 53

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#19).

```
#19 POST /v2/default/transactions key=tx-19
fault concurrent(#19)
```

Delivery order: send-concurrently #19

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0019","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0019","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0019":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0019":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0019","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0019"},{"amount":99,"asset":"USD/2","destination":"sellers:4","source":"orders:0019"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0019":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0019":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-19","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0019","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0019","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:4","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0019","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0019","fees","sellers:4"],"destinations_arrays":[{"0":"orders","1":"0019","2":null},{"0":"fees","1":null},{"0":"sellers","1":"4","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0019":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0019\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0019\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0019\",\"destination\":\"sellers:4\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0019","orders:0019"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0019","2":null},{"0":"orders","1":"0019","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-53/repro.sh`

### Finding 54

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#20).

```
#20 POST /v2/default/transactions key=tx-20
fault concurrent(#20)
```

Delivery order: send-concurrently #20

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"orders:0020","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0020","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0020":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0020":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0020","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0020"},{"amount":99,"asset":"USD/2","destination":"sellers:0","source":"orders:0020"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0020":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0020":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-20","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0020","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0020","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0020","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:0","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0020","fees","sellers:0"],"destinations_arrays":[{"0":"orders","1":"0020","2":null},{"0":"fees","1":null},{"0":"sellers","1":"0","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0020":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0020\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0020\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0020\",\"destination\":\"sellers:0\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0020","orders:0020"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0020","2":null},{"0":"orders","1":"0020","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-54/repro.sh`

### Finding 55

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#21).

```
#21 POST /v2/default/transactions key=tx-21
fault concurrent(#21)
```

Delivery order: send-concurrently #21

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0021","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0021","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0021":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0021":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0021","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0021"},{"amount":99,"asset":"USD/2","destination":"sellers:1","source":"orders:0021"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0021":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0021":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-21","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0021","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0021","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:1","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0021","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0021","fees","sellers:1"],"destinations_arrays":[{"0":"orders","1":"0021","2":null},{"0":"fees","1":null},{"0":"sellers","1":"1","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0021":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0021\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0021\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0021\",\"destination\":\"sellers:1\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0021","orders:0021"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0021","2":null},{"0":"orders","1":"0021","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-55/repro.sh`

### Finding 56

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#22).

```
#22 POST /v2/default/transactions key=tx-22
fault concurrent(#22)
```

Delivery order: send-concurrently #22

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0022","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0022","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0022":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0022":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0022","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0022"},{"amount":99,"asset":"USD/2","destination":"sellers:2","source":"orders:0022"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0022":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0022":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-22","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"sellers:2","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0022","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0022","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0022","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0022","fees","sellers:2"],"destinations_arrays":[{"0":"orders","1":"0022","2":null},{"0":"fees","1":null},{"0":"sellers","1":"2","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0022":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0022\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0022\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0022\",\"destination\":\"sellers:2\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0022","orders:0022"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0022","2":null},{"0":"orders","1":"0022","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-56/repro.sh`

### Finding 57

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#23).

```
#23 POST /v2/default/transactions key=tx-23
fault concurrent(#23)
```

Delivery order: send-concurrently #23

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0023","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0023","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0023":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0023":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0023","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0023"},{"amount":99,"asset":"USD/2","destination":"sellers:3","source":"orders:0023"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0023":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0023":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-23","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0023","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0023","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:3","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0023","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0023","fees","sellers:3"],"destinations_arrays":[{"0":"orders","1":"0023","2":null},{"0":"fees","1":null},{"0":"sellers","1":"3","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0023":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0023\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0023\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0023\",\"destination\":\"sellers:3\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0023","orders:0023"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0023","2":null},{"0":"orders","1":"0023","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-57/repro.sh`

### Finding 58

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#24).

```
#24 POST /v2/default/transactions key=tx-24
fault concurrent(#24)
```

Delivery order: send-concurrently #24

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0024","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0024","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0024":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0024":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0024","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0024"},{"amount":99,"asset":"USD/2","destination":"sellers:4","source":"orders:0024"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0024":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0024":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-24","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0024","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0024","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:4","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0024","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0024","fees","sellers:4"],"destinations_arrays":[{"0":"orders","1":"0024","2":null},{"0":"fees","1":null},{"0":"sellers","1":"4","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0024":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0024\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0024\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0024\",\"destination\":\"sellers:4\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0024","orders:0024"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0024","2":null},{"0":"orders","1":"0024","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-58/repro.sh`

### Finding 59

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#25).

```
#25 POST /v2/default/transactions key=tx-25
fault concurrent(#25)
```

Delivery order: send-concurrently #25

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0025","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0025","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0025":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0025":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0025","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0025"},{"amount":99,"asset":"USD/2","destination":"sellers:0","source":"orders:0025"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0025":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0025":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-25","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0025","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0025","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0025","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:0","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0025","fees","sellers:0"],"destinations_arrays":[{"0":"orders","1":"0025","2":null},{"0":"fees","1":null},{"0":"sellers","1":"0","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0025":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0025\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0025\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0025\",\"destination\":\"sellers:0\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0025","orders:0025"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0025","2":null},{"0":"orders","1":"0025","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-59/repro.sh`

### Finding 60

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#26).

```
#26 POST /v2/default/transactions key=tx-26
fault concurrent(#26)
```

Delivery order: send-concurrently #26

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0026","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0026","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0026":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0026":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0026","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0026"},{"amount":99,"asset":"USD/2","destination":"sellers:1","source":"orders:0026"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0026":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0026":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-26","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0026","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0026","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0026","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:1","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0026","fees","sellers:1"],"destinations_arrays":[{"0":"orders","1":"0026","2":null},{"0":"fees","1":null},{"0":"sellers","1":"1","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0026":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0026\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0026\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0026\",\"destination\":\"sellers:1\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0026","orders:0026"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0026","2":null},{"0":"orders","1":"0026","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-60/repro.sh`

### Finding 61

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#27).

```
#27 POST /v2/default/transactions key=tx-27
fault concurrent(#27)
```

Delivery order: send-concurrently #27

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0027","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0027","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0027":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0027":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0027","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0027"},{"amount":99,"asset":"USD/2","destination":"sellers:2","source":"orders:0027"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0027":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0027":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-27","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"sellers:2","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0027","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0027","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0027","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0027","fees","sellers:2"],"destinations_arrays":[{"0":"orders","1":"0027","2":null},{"0":"fees","1":null},{"0":"sellers","1":"2","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0027":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0027\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0027\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0027\",\"destination\":\"sellers:2\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0027","orders:0027"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0027","2":null},{"0":"orders","1":"0027","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-61/repro.sh`

### Finding 62

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#28).

```
#28 POST /v2/default/transactions key=tx-28
fault concurrent(#28)
```

Delivery order: send-concurrently #28

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0028","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0028","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0028":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0028":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0028","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0028"},{"amount":99,"asset":"USD/2","destination":"sellers:3","source":"orders:0028"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0028":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0028":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-28","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0028","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0028","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:3","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0028","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0028","fees","sellers:3"],"destinations_arrays":[{"0":"orders","1":"0028","2":null},{"0":"fees","1":null},{"0":"sellers","1":"3","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0028":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0028\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0028\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0028\",\"destination\":\"sellers:3\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0028","orders:0028"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0028","2":null},{"0":"orders","1":"0028","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-62/repro.sh`

### Finding 63

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#29).

```
#29 POST /v2/default/transactions key=tx-29
fault concurrent(#29)
```

Delivery order: send-concurrently #29

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0029","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0029","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0029":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0029":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0029","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0029"},{"amount":99,"asset":"USD/2","destination":"sellers:4","source":"orders:0029"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0029":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0029":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-29","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0029","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0029","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0029","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:4","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0029","fees","sellers:4"],"destinations_arrays":[{"0":"orders","1":"0029","2":null},{"0":"fees","1":null},{"0":"sellers","1":"4","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0029":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0029\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0029\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0029\",\"destination\":\"sellers:4\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0029","orders:0029"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0029","2":null},{"0":"orders","1":"0029","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-63/repro.sh`

### Finding 64

Reproduced in 6 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: false. Reached from 1 single-fault schedules: concurrent(#42).

```
#42 POST /v2/default/accounts/sellers:2/metadata key=am-2
fault concurrent(#42)
```

Delivery order: send-concurrently #42

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"tier":"gold-2"},"targetType":"ACCOUNT"},"idempotency_key":"am-2","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-64/repro.sh`

### Finding 65

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#0@commit1).

```
#0 POST /v2/default/transactions key=tx-0
fault crash(#0@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #0

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0000","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0000","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0000":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0000":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0000","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0000"},{"amount":99,"asset":"USD/2","destination":"sellers:0","source":"orders:0000"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0000":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0000":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-0","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0000","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0000","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0000","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:0","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0000","fees","sellers:0"],"destinations_arrays":[{"0":"orders","1":"0000","2":null},{"0":"fees","1":null},{"0":"sellers","1":"0","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0000":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0000\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0000\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0000\",\"destination\":\"sellers:0\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0000","orders:0000"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0000","2":null},{"0":"orders","1":"0000","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-65/repro.sh`

### Finding 66

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#1@commit1).

```
#1 POST /v2/default/transactions key=tx-1
fault crash(#1@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #1

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0001","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0001","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0001":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0001":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0001","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0001"},{"amount":99,"asset":"USD/2","destination":"sellers:1","source":"orders:0001"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0001":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0001":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-1","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0001","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0001","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0001","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:1","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0001","fees","sellers:1"],"destinations_arrays":[{"0":"orders","1":"0001","2":null},{"0":"fees","1":null},{"0":"sellers","1":"1","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0001":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0001\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0001\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0001\",\"destination\":\"sellers:1\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0001","orders:0001"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0001","2":null},{"0":"orders","1":"0001","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-66/repro.sh`

### Finding 67

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#2@commit1).

```
#2 POST /v2/default/transactions key=tx-2
fault crash(#2@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #2

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0002","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0002","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0002":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0002":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0002","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0002"},{"amount":99,"asset":"USD/2","destination":"sellers:2","source":"orders:0002"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0002":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0002":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-2","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0002","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:2","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0002","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0002","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0002","fees","sellers:2"],"destinations_arrays":[{"0":"orders","1":"0002","2":null},{"0":"fees","1":null},{"0":"sellers","1":"2","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0002":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0002\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0002\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0002\",\"destination\":\"sellers:2\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0002","orders:0002"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0002","2":null},{"0":"orders","1":"0002","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-67/repro.sh`

### Finding 68

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#3@commit1).

```
#3 POST /v2/default/transactions key=tx-3
fault crash(#3@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #3

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0003","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0003","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0003":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0003":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0003","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0003"},{"amount":99,"asset":"USD/2","destination":"sellers:3","source":"orders:0003"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0003":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0003":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-3","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0003","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0003","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0003","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:3","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0003","fees","sellers:3"],"destinations_arrays":[{"0":"orders","1":"0003","2":null},{"0":"fees","1":null},{"0":"sellers","1":"3","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0003":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0003\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0003\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0003\",\"destination\":\"sellers:3\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0003","orders:0003"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0003","2":null},{"0":"orders","1":"0003","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-68/repro.sh`

### Finding 69

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#4@commit1).

```
#4 POST /v2/default/transactions key=tx-4
fault crash(#4@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #4

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0004","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0004","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0004":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0004":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0004","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0004"},{"amount":99,"asset":"USD/2","destination":"sellers:4","source":"orders:0004"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0004":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0004":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-4","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0004","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0004","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0004","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:4","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0004","fees","sellers:4"],"destinations_arrays":[{"0":"orders","1":"0004","2":null},{"0":"fees","1":null},{"0":"sellers","1":"4","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0004":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0004\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0004\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0004\",\"destination\":\"sellers:4\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0004","orders:0004"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0004","2":null},{"0":"orders","1":"0004","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-69/repro.sh`

### Finding 70

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#5@commit1).

```
#5 POST /v2/default/transactions key=tx-5
fault crash(#5@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #5

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0005","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0005","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0005":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0005":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0005","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0005"},{"amount":99,"asset":"USD/2","destination":"sellers:0","source":"orders:0005"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0005":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0005":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-5","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0005","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0005","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0005","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:0","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0005","fees","sellers:0"],"destinations_arrays":[{"0":"orders","1":"0005","2":null},{"0":"fees","1":null},{"0":"sellers","1":"0","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0005":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0005\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0005\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0005\",\"destination\":\"sellers:0\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0005","orders:0005"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0005","2":null},{"0":"orders","1":"0005","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-70/repro.sh`

### Finding 71

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#6@commit1).

```
#6 POST /v2/default/transactions key=tx-6
fault crash(#6@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #6

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0006","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"orders:0006","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0006":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0006":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0006","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0006"},{"amount":99,"asset":"USD/2","destination":"sellers:1","source":"orders:0006"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0006":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0006":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-6","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0006","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0006","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:1","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0006","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0006","fees","sellers:1"],"destinations_arrays":[{"0":"orders","1":"0006","2":null},{"0":"fees","1":null},{"0":"sellers","1":"1","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0006":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0006\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0006\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0006\",\"destination\":\"sellers:1\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0006","orders:0006"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0006","2":null},{"0":"orders","1":"0006","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-71/repro.sh`

### Finding 72

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#7@commit1).

```
#7 POST /v2/default/transactions key=tx-7
fault crash(#7@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #7

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0007","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0007","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0007":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0007":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0007","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0007"},{"amount":99,"asset":"USD/2","destination":"sellers:2","source":"orders:0007"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0007":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0007":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-7","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0007","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:2","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0007","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0007","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0007","fees","sellers:2"],"destinations_arrays":[{"0":"orders","1":"0007","2":null},{"0":"fees","1":null},{"0":"sellers","1":"2","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0007":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0007\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0007\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0007\",\"destination\":\"sellers:2\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0007","orders:0007"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0007","2":null},{"0":"orders","1":"0007","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-72/repro.sh`

### Finding 73

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#8@commit1).

```
#8 POST /v2/default/transactions key=tx-8
fault crash(#8@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #8

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0008","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0008","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0008":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0008":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0008","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0008"},{"amount":99,"asset":"USD/2","destination":"sellers:3","source":"orders:0008"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0008":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0008":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-8","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0008","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:3","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0008","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0008","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0008","fees","sellers:3"],"destinations_arrays":[{"0":"orders","1":"0008","2":null},{"0":"fees","1":null},{"0":"sellers","1":"3","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0008":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0008\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0008\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0008\",\"destination\":\"sellers:3\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0008","orders:0008"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0008","2":null},{"0":"orders","1":"0008","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-73/repro.sh`

### Finding 74

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#9@commit1).

```
#9 POST /v2/default/transactions key=tx-9
fault crash(#9@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #9

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0009","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0009","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0009":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0009":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0009","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0009"},{"amount":99,"asset":"USD/2","destination":"sellers:4","source":"orders:0009"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0009":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0009":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-9","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0009","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0009","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:4","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0009","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0009","fees","sellers:4"],"destinations_arrays":[{"0":"orders","1":"0009","2":null},{"0":"fees","1":null},{"0":"sellers","1":"4","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0009":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0009\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0009\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0009\",\"destination\":\"sellers:4\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0009","orders:0009"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0009","2":null},{"0":"orders","1":"0009","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-74/repro.sh`

### Finding 75

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#10@commit1).

```
#10 POST /v2/default/transactions key=tx-10
fault crash(#10@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #10

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0010","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0010","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0010":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0010":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0010","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0010"},{"amount":99,"asset":"USD/2","destination":"sellers:0","source":"orders:0010"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0010":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0010":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-10","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0010","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0010","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0010","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:0","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0010","fees","sellers:0"],"destinations_arrays":[{"0":"orders","1":"0010","2":null},{"0":"fees","1":null},{"0":"sellers","1":"0","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0010":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0010\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0010\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0010\",\"destination\":\"sellers:0\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0010","orders:0010"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0010","2":null},{"0":"orders","1":"0010","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-75/repro.sh`

### Finding 76

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#11@commit1).

```
#11 POST /v2/default/transactions key=tx-11
fault crash(#11@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #11

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0011","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0011","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0011":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0011":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0011","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0011"},{"amount":99,"asset":"USD/2","destination":"sellers:1","source":"orders:0011"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0011":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0011":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-11","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0011","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0011","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0011","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:1","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0011","fees","sellers:1"],"destinations_arrays":[{"0":"orders","1":"0011","2":null},{"0":"fees","1":null},{"0":"sellers","1":"1","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0011":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0011\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0011\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0011\",\"destination\":\"sellers:1\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0011","orders:0011"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0011","2":null},{"0":"orders","1":"0011","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-76/repro.sh`

### Finding 77

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#12@commit1).

```
#12 POST /v2/default/transactions key=tx-12
fault crash(#12@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #12

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0012","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0012","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0012":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0012":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0012","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0012"},{"amount":99,"asset":"USD/2","destination":"sellers:2","source":"orders:0012"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0012":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0012":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-12","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"sellers:2","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0012","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0012","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0012","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0012","fees","sellers:2"],"destinations_arrays":[{"0":"orders","1":"0012","2":null},{"0":"fees","1":null},{"0":"sellers","1":"2","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0012":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0012\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0012\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0012\",\"destination\":\"sellers:2\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0012","orders:0012"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0012","2":null},{"0":"orders","1":"0012","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-77/repro.sh`

### Finding 78

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#13@commit1).

```
#13 POST /v2/default/transactions key=tx-13
fault crash(#13@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #13

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0013","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
- _default.accounts_volumes x1 {"accounts_address":"orders:0013","asset":"USD/2","input":100,"ledger":"default","output":100}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0013":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0013":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0013","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0013"},{"amount":99,"asset":"USD/2","destination":"sellers:3","source":"orders:0013"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0013":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0013":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-13","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0013","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0013","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0013","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:3","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0013","fees","sellers:3"],"destinations_arrays":[{"0":"orders","1":"0013","2":null},{"0":"fees","1":null},{"0":"sellers","1":"3","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0013":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0013\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0013\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0013\",\"destination\":\"sellers:3\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0013","orders:0013"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0013","2":null},{"0":"orders","1":"0013","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-78/repro.sh`

### Finding 79

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#14@commit1).

```
#14 POST /v2/default/transactions key=tx-14
fault crash(#14@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #14

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0014","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0014","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0014":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0014":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0014","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0014"},{"amount":99,"asset":"USD/2","destination":"sellers:4","source":"orders:0014"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0014":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0014":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-14","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0014","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0014","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0014","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:4","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0014","fees","sellers:4"],"destinations_arrays":[{"0":"orders","1":"0014","2":null},{"0":"fees","1":null},{"0":"sellers","1":"4","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0014":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0014\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0014\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0014\",\"destination\":\"sellers:4\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0014","orders:0014"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0014","2":null},{"0":"orders","1":"0014","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-79/repro.sh`

### Finding 80

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#15@commit1).

```
#15 POST /v2/default/transactions key=tx-15
fault crash(#15@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #15

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0015","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0015","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0015":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0015":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0015","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0015"},{"amount":99,"asset":"USD/2","destination":"sellers:0","source":"orders:0015"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0015":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0015":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-15","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0015","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0015","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0015","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:0","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0015","fees","sellers:0"],"destinations_arrays":[{"0":"orders","1":"0015","2":null},{"0":"fees","1":null},{"0":"sellers","1":"0","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0015":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0015\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0015\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0015\",\"destination\":\"sellers:0\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0015","orders:0015"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0015","2":null},{"0":"orders","1":"0015","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-80/repro.sh`

### Finding 81

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#16@commit1).

```
#16 POST /v2/default/transactions key=tx-16
fault crash(#16@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #16

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"orders:0016","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0016","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0016":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0016":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0016","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0016"},{"amount":99,"asset":"USD/2","destination":"sellers:1","source":"orders:0016"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0016":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0016":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-16","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0016","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0016","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0016","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:1","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0016","fees","sellers:1"],"destinations_arrays":[{"0":"orders","1":"0016","2":null},{"0":"fees","1":null},{"0":"sellers","1":"1","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0016":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0016\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0016\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0016\",\"destination\":\"sellers:1\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0016","orders:0016"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0016","2":null},{"0":"orders","1":"0016","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-81/repro.sh`

### Finding 82

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#17@commit1).

```
#17 POST /v2/default/transactions key=tx-17
fault crash(#17@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #17

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0017","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"orders:0017","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0017":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0017":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0017","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0017"},{"amount":99,"asset":"USD/2","destination":"sellers:2","source":"orders:0017"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0017":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0017":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-17","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"sellers:2","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0017","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0017","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0017","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0017","fees","sellers:2"],"destinations_arrays":[{"0":"orders","1":"0017","2":null},{"0":"fees","1":null},{"0":"sellers","1":"2","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0017":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0017\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0017\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0017\",\"destination\":\"sellers:2\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0017","orders:0017"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0017","2":null},{"0":"orders","1":"0017","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-82/repro.sh`

### Finding 83

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#18@commit1).

```
#18 POST /v2/default/transactions key=tx-18
fault crash(#18@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #18

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0018","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0018","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0018":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0018":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0018","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0018"},{"amount":99,"asset":"USD/2","destination":"sellers:3","source":"orders:0018"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0018":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0018":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-18","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0018","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:3","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0018","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0018","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0018","fees","sellers:3"],"destinations_arrays":[{"0":"orders","1":"0018","2":null},{"0":"fees","1":null},{"0":"sellers","1":"3","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0018":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0018\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0018\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0018\",\"destination\":\"sellers:3\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0018","orders:0018"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0018","2":null},{"0":"orders","1":"0018","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-83/repro.sh`

### Finding 84

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#19@commit1).

```
#19 POST /v2/default/transactions key=tx-19
fault crash(#19@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #19

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0019","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0019","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0019":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0019":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0019","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0019"},{"amount":99,"asset":"USD/2","destination":"sellers:4","source":"orders:0019"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0019":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0019":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-19","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0019","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0019","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:4","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0019","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0019","fees","sellers:4"],"destinations_arrays":[{"0":"orders","1":"0019","2":null},{"0":"fees","1":null},{"0":"sellers","1":"4","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0019":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0019\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0019\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0019\",\"destination\":\"sellers:4\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0019","orders:0019"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0019","2":null},{"0":"orders","1":"0019","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-84/repro.sh`

### Finding 85

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#20@commit1).

```
#20 POST /v2/default/transactions key=tx-20
fault crash(#20@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #20

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"orders:0020","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0020","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0020":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0020":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0020","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0020"},{"amount":99,"asset":"USD/2","destination":"sellers:0","source":"orders:0020"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0020":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0020":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-20","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0020","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0020","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0020","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:0","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0020","fees","sellers:0"],"destinations_arrays":[{"0":"orders","1":"0020","2":null},{"0":"fees","1":null},{"0":"sellers","1":"0","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0020":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0020\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0020\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0020\",\"destination\":\"sellers:0\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0020","orders:0020"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0020","2":null},{"0":"orders","1":"0020","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-85/repro.sh`

### Finding 86

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#21@commit1).

```
#21 POST /v2/default/transactions key=tx-21
fault crash(#21@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #21

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0021","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0021","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0021":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0021":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0021","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0021"},{"amount":99,"asset":"USD/2","destination":"sellers:1","source":"orders:0021"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0021":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0021":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-21","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0021","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0021","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:1","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0021","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0021","fees","sellers:1"],"destinations_arrays":[{"0":"orders","1":"0021","2":null},{"0":"fees","1":null},{"0":"sellers","1":"1","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0021":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0021\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0021\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0021\",\"destination\":\"sellers:1\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0021","orders:0021"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0021","2":null},{"0":"orders","1":"0021","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-86/repro.sh`

### Finding 87

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#22@commit1).

```
#22 POST /v2/default/transactions key=tx-22
fault crash(#22@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #22

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0022","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0022","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0022":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0022":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0022","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0022"},{"amount":99,"asset":"USD/2","destination":"sellers:2","source":"orders:0022"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0022":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0022":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-22","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"sellers:2","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0022","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0022","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0022","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0022","fees","sellers:2"],"destinations_arrays":[{"0":"orders","1":"0022","2":null},{"0":"fees","1":null},{"0":"sellers","1":"2","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0022":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0022\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0022\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0022\",\"destination\":\"sellers:2\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0022","orders:0022"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0022","2":null},{"0":"orders","1":"0022","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-87/repro.sh`

### Finding 88

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#23@commit1).

```
#23 POST /v2/default/transactions key=tx-23
fault crash(#23@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #23

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0023","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0023","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0023":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0023":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0023","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0023"},{"amount":99,"asset":"USD/2","destination":"sellers:3","source":"orders:0023"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0023":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0023":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-23","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0023","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0023","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:3","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0023","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0023","fees","sellers:3"],"destinations_arrays":[{"0":"orders","1":"0023","2":null},{"0":"fees","1":null},{"0":"sellers","1":"3","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0023":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0023\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0023\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0023\",\"destination\":\"sellers:3\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0023","orders:0023"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0023","2":null},{"0":"orders","1":"0023","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-88/repro.sh`

### Finding 89

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#24@commit1).

```
#24 POST /v2/default/transactions key=tx-24
fault crash(#24@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #24

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0024","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0024","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0024":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0024":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0024","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0024"},{"amount":99,"asset":"USD/2","destination":"sellers:4","source":"orders:0024"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0024":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0024":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-24","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0024","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0024","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:4","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0024","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0024","fees","sellers:4"],"destinations_arrays":[{"0":"orders","1":"0024","2":null},{"0":"fees","1":null},{"0":"sellers","1":"4","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0024":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0024\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0024\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0024\",\"destination\":\"sellers:4\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0024","orders:0024"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0024","2":null},{"0":"orders","1":"0024","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-89/repro.sh`

### Finding 90

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#25@commit1).

```
#25 POST /v2/default/transactions key=tx-25
fault crash(#25@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #25

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0025","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"sellers:0","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0025","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0025":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0025":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0025","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0025"},{"amount":99,"asset":"USD/2","destination":"sellers:0","source":"orders:0025"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0025":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0025":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:0":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-25","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0025","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0025","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0025","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:0","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0025","fees","sellers:0"],"destinations_arrays":[{"0":"orders","1":"0025","2":null},{"0":"fees","1":null},{"0":"sellers","1":"0","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0025":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:0":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0025\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0025\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0025\",\"destination\":\"sellers:0\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0025","orders:0025"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0025","2":null},{"0":"orders","1":"0025","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-90/repro.sh`

### Finding 91

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#26@commit1).

```
#26 POST /v2/default/transactions key=tx-26
fault crash(#26@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #26

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0026","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:1","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0026","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0026":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0026":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0026","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0026"},{"amount":99,"asset":"USD/2","destination":"sellers:1","source":"orders:0026"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0026":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0026":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:1":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-26","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0026","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0026","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0026","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:1","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0026","fees","sellers:1"],"destinations_arrays":[{"0":"orders","1":"0026","2":null},{"0":"fees","1":null},{"0":"sellers","1":"1","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0026":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:1":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0026\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0026\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0026\",\"destination\":\"sellers:1\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0026","orders:0026"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0026","2":null},{"0":"orders","1":"0026","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-91/repro.sh`

### Finding 92

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#27@commit1).

```
#27 POST /v2/default/transactions key=tx-27
fault crash(#27@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #27

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0027","asset":"USD/2","input":200,"ledger":"default","output":200}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:2","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0027","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0027":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0027":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0027","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0027"},{"amount":99,"asset":"USD/2","destination":"sellers:2","source":"orders:0027"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0027":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0027":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:2":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-27","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"sellers:2","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0027","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0027","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0027","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0027","fees","sellers:2"],"destinations_arrays":[{"0":"orders","1":"0027","2":null},{"0":"fees","1":null},{"0":"sellers","1":"2","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0027":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:2":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0027\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0027\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0027\",\"destination\":\"sellers:2\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0027","orders:0027"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0027","2":null},{"0":"orders","1":"0027","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-92/repro.sh`

### Finding 93

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#28@commit1).

```
#28 POST /v2/default/transactions key=tx-28
fault crash(#28@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #28

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0028","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":198,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0028","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:3","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0028":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0028":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0028","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0028"},{"amount":99,"asset":"USD/2","destination":"sellers:3","source":"orders:0028"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0028":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0028":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:3":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-28","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0028","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0028","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:3","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0028","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0028","fees","sellers:3"],"destinations_arrays":[{"0":"orders","1":"0028","2":null},{"0":"fees","1":null},{"0":"sellers","1":"3","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0028":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:3":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0028\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0028\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0028\",\"destination\":\"sellers:3\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0028","orders:0028"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0028","2":null},{"0":"orders","1":"0028","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-93/repro.sh`

### Finding 94

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#29@commit1).

```
#29 POST /v2/default/transactions key=tx-29
fault crash(#29@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #29

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":201}
+ _default.accounts_volumes x1 {"accounts_address":"orders:0029","asset":"USD/2","input":200,"ledger":"default","output":200}
+ _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":198,"ledger":"default","output":0}
+ _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":2,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"sellers:4","asset":"USD/2","input":99,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"orders:0029","asset":"USD/2","input":100,"ledger":"default","output":100}
- _default.accounts_volumes x1 {"accounts_address":"fees","asset":"USD/2","input":1,"ledger":"default","output":0}
- _default.accounts_volumes x1 {"accounts_address":"world","asset":"USD/2","input":0,"ledger":"default","output":101}
+ _default.logs x1 {"data":{"accountMetadata":{},"transaction":{"id":"<ref:_default.transactions>","metadata":{},"postCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0029":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postCommitVolumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0029":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":[{"amount":100,"asset":"USD/2","destination":"orders:0029","source":"world"},{"amount":1,"asset":"USD/2","destination":"fees","source":"orders:0029"},{"amount":99,"asset":"USD/2","destination":"sellers:4","source":"orders:0029"}],"preCommitEffectiveVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0029":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"preCommitVolumes":{"fees":{"USD/2":{"balance":1,"input":1,"output":0}},"orders:0029":{"USD/2":{"balance":0,"input":100,"output":100}},"sellers:4":{"USD/2":{"balance":99,"input":99,"output":0}},"world":{"USD/2":{"balance":-101,"input":0,"output":101}}},"reverted":false}},"idempotency_key":"tx-29","ledger":"default","schema_version":null,"type":"NEW_TRANSACTION"}
+ _default.moves x1 {"accounts_address":"orders:0029","amount":100,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":100},"post_commit_volumes":{"inputs":200,"outputs":100},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"world","amount":100,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":0,"outputs":201},"post_commit_volumes":{"inputs":0,"outputs":201},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0029","amount":1,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":101},"post_commit_volumes":{"inputs":200,"outputs":101},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"orders:0029","amount":99,"asset":"USD/2","is_source":true,"ledger":"default","post_commit_effective_volumes":{"inputs":200,"outputs":200},"post_commit_volumes":{"inputs":200,"outputs":200},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"sellers:4","amount":99,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":198,"outputs":0},"post_commit_volumes":{"inputs":198,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.moves x1 {"accounts_address":"fees","amount":1,"asset":"USD/2","is_source":false,"ledger":"default","post_commit_effective_volumes":{"inputs":2,"outputs":0},"post_commit_volumes":{"inputs":2,"outputs":0},"transactions_id":"<ref:_default.transactions>"}
+ _default.transactions x1 {"destinations":["orders:0029","fees","sellers:4"],"destinations_arrays":[{"0":"orders","1":"0029","2":null},{"0":"fees","1":null},{"0":"sellers","1":"4","2":null}],"ledger":"default","metadata":{},"post_commit_volumes":{"fees":{"USD/2":{"balance":2,"input":2,"output":0}},"orders:0029":{"USD/2":{"balance":0,"input":200,"output":200}},"sellers:4":{"USD/2":{"balance":198,"input":198,"output":0}},"world":{"USD/2":{"balance":-201,"input":0,"output":201}}},"postings":"[{\"source\":\"world\",\"destination\":\"orders:0029\",\"amount\":100,\"asset\":\"USD/2\"},{\"source\":\"orders:0029\",\"destination\":\"fees\",\"amount\":1,\"asset\":\"USD/2\"},{\"source\":\"orders:0029\",\"destination\":\"sellers:4\",\"amount\":99,\"asset\":\"USD/2\"}]","reference":null,"sources":["world","orders:0029","orders:0029"],"sources_arrays":[{"0":"world","1":null},{"0":"orders","1":"0029","2":null},{"0":"orders","1":"0029","2":null}],"template":""} (+8 referencing rows)
+ _default.transactions_metadata x1 {"ledger":"default","metadata":{},"revision":1,"transactions_id":"<ref:_default.transactions>"}
```

Repro: `finding-94/repro.sh`

### Finding 95

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#40@commit1).

```
#40 POST /v2/default/accounts/sellers:0/metadata key=am-0
fault crash(#40@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #40

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"tier":"gold-0"},"targetType":"ACCOUNT"},"idempotency_key":"am-0","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-95/repro.sh`

### Finding 96

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#41@commit1).

```
#41 POST /v2/default/accounts/sellers:1/metadata key=am-1
fault crash(#41@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #41

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"tier":"gold-1"},"targetType":"ACCOUNT"},"idempotency_key":"am-1","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-96/repro.sh`

### Finding 97

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#42@commit1).

```
#42 POST /v2/default/accounts/sellers:2/metadata key=am-2
fault crash(#42@commit1)
```

Delivery order: send, kill the service after its commit 1, restart, resend #42

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"tier":"gold-2"},"targetType":"ACCOUNT"},"idempotency_key":"am-2","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-97/repro.sh`

### Finding 98

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#30).

```
#29 POST /v2/default/transactions key=tx-29
#30 POST /v2/default/transactions/4000000000000/metadata key=md-0
fault lost(#30)
```

Delivery order: send #29, send-twice #30

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"note":"checked-0"},"targetType":"TRANSACTION"},"idempotency_key":"md-0","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-98/repro.sh`

### Finding 99

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#33).

```
#18 POST /v2/default/transactions key=tx-18
#33 POST /v2/default/transactions/4000000000018/metadata key=md-3
fault lost(#33)
```

Delivery order: send #18, send-twice #33

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"note":"checked-3"},"targetType":"TRANSACTION"},"idempotency_key":"md-3","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-99/repro.sh`

### Finding 100

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#34).

```
#24 POST /v2/default/transactions key=tx-24
#34 POST /v2/default/transactions/4000000000024/metadata key=md-4
fault lost(#34)
```

Delivery order: send #24, send-twice #34

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"note":"checked-4"},"targetType":"TRANSACTION"},"idempotency_key":"md-4","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-100/repro.sh`

### Finding 101

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#30).

```
#29 POST /v2/default/transactions key=tx-29
#30 POST /v2/default/transactions/4000000000000/metadata key=md-0
fault concurrent(#30)
```

Delivery order: send #29, send-concurrently #30

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"note":"checked-0"},"targetType":"TRANSACTION"},"idempotency_key":"md-0","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-101/repro.sh`

### Finding 102

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#33).

```
#18 POST /v2/default/transactions key=tx-18
#33 POST /v2/default/transactions/4000000000018/metadata key=md-3
fault concurrent(#33)
```

Delivery order: send #18, send-concurrently #33

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"note":"checked-3"},"targetType":"TRANSACTION"},"idempotency_key":"md-3","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-102/repro.sh`

### Finding 103

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#34).

```
#24 POST /v2/default/transactions key=tx-24
#34 POST /v2/default/transactions/4000000000024/metadata key=md-4
fault concurrent(#34)
```

Delivery order: send #24, send-concurrently #34

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"note":"checked-4"},"targetType":"TRANSACTION"},"idempotency_key":"md-4","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-103/repro.sh`

### Finding 104

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#40).

```
#25 POST /v2/default/transactions key=tx-25
#40 POST /v2/default/accounts/sellers:0/metadata key=am-0
fault concurrent(#40)
```

Delivery order: send #25, send-concurrently #40

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"tier":"gold-0"},"targetType":"ACCOUNT"},"idempotency_key":"am-0","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-104/repro.sh`

### Finding 105

Reproduced in 4 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#41).

```
#41 POST /v2/default/accounts/sellers:1/metadata key=am-1
#42 POST /v2/default/accounts/sellers:2/metadata key=am-2
fault concurrent(#41)
```

Delivery order: send-concurrently #41, send #42

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"tier":"gold-1"},"targetType":"ACCOUNT"},"idempotency_key":"am-1","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-105/repro.sh`

### Finding 106

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#30@commit1).

```
#29 POST /v2/default/transactions key=tx-29
#30 POST /v2/default/transactions/4000000000000/metadata key=md-0
fault crash(#30@commit1)
```

Delivery order: send #29, send, kill the service after its commit 1, restart, resend #30

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"note":"checked-0"},"targetType":"TRANSACTION"},"idempotency_key":"md-0","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-106/repro.sh`

### Finding 107

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#33@commit1).

```
#18 POST /v2/default/transactions key=tx-18
#33 POST /v2/default/transactions/4000000000018/metadata key=md-3
fault crash(#33@commit1)
```

Delivery order: send #18, send, kill the service after its commit 1, restart, resend #33

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"note":"checked-3"},"targetType":"TRANSACTION"},"idempotency_key":"md-3","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-107/repro.sh`

### Finding 108

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#34@commit1).

```
#24 POST /v2/default/transactions key=tx-24
#34 POST /v2/default/transactions/4000000000024/metadata key=md-4
fault crash(#34@commit1)
```

Delivery order: send #24, send, kill the service after its commit 1, restart, resend #34

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"note":"checked-4"},"targetType":"TRANSACTION"},"idempotency_key":"md-4","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-108/repro.sh`

### Finding 109

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#31).

```
#23 POST /v2/default/transactions key=tx-23
#24 POST /v2/default/transactions key=tx-24
#25 POST /v2/default/transactions key=tx-25
#26 POST /v2/default/transactions key=tx-26
#27 POST /v2/default/transactions key=tx-27
#28 POST /v2/default/transactions key=tx-28
#29 POST /v2/default/transactions key=tx-29
#31 POST /v2/default/transactions/4000000000006/metadata key=md-1
fault lost(#31)
```

Delivery order: send #23, send #24, send #25, send #26, send #27, send #28, send #29, send-twice #31

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"note":"checked-1"},"targetType":"TRANSACTION"},"idempotency_key":"md-1","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-109/repro.sh`

### Finding 110

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#31).

```
#23 POST /v2/default/transactions key=tx-23
#24 POST /v2/default/transactions key=tx-24
#25 POST /v2/default/transactions key=tx-25
#26 POST /v2/default/transactions key=tx-26
#27 POST /v2/default/transactions key=tx-27
#28 POST /v2/default/transactions key=tx-28
#29 POST /v2/default/transactions key=tx-29
#31 POST /v2/default/transactions/4000000000006/metadata key=md-1
fault concurrent(#31)
```

Delivery order: send #23, send #24, send #25, send #26, send #27, send #28, send #29, send-concurrently #31

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"note":"checked-1"},"targetType":"TRANSACTION"},"idempotency_key":"md-1","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-110/repro.sh`

### Finding 111

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#31@commit1).

```
#23 POST /v2/default/transactions key=tx-23
#24 POST /v2/default/transactions key=tx-24
#25 POST /v2/default/transactions key=tx-25
#26 POST /v2/default/transactions key=tx-26
#27 POST /v2/default/transactions key=tx-27
#28 POST /v2/default/transactions key=tx-28
#29 POST /v2/default/transactions key=tx-29
#31 POST /v2/default/transactions/4000000000006/metadata key=md-1
fault crash(#31@commit1)
```

Delivery order: send #23, send #24, send #25, send #26, send #27, send #28, send #29, send, kill the service after its commit 1, restart, resend #31

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"note":"checked-1"},"targetType":"TRANSACTION"},"idempotency_key":"md-1","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-111/repro.sh`

### Finding 112

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: lost(#32).

```
#16 POST /v2/default/transactions key=tx-16
#17 POST /v2/default/transactions key=tx-17
#18 POST /v2/default/transactions key=tx-18
#19 POST /v2/default/transactions key=tx-19
#20 POST /v2/default/transactions key=tx-20
#21 POST /v2/default/transactions key=tx-21
#22 POST /v2/default/transactions key=tx-22
#23 POST /v2/default/transactions key=tx-23
#24 POST /v2/default/transactions key=tx-24
#25 POST /v2/default/transactions key=tx-25
#26 POST /v2/default/transactions key=tx-26
#27 POST /v2/default/transactions key=tx-27
#28 POST /v2/default/transactions key=tx-28
#32 POST /v2/default/transactions/4000000000012/metadata key=md-2
fault lost(#32)
```

Delivery order: send #16, send #17, send #18, send #19, send #20, send #21, send #22, send #23, send #24, send #25, send #26, send #27, send #28, send-twice #32

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"note":"checked-2"},"targetType":"TRANSACTION"},"idempotency_key":"md-2","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-112/repro.sh`

### Finding 113

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#32).

```
#16 POST /v2/default/transactions key=tx-16
#17 POST /v2/default/transactions key=tx-17
#18 POST /v2/default/transactions key=tx-18
#19 POST /v2/default/transactions key=tx-19
#20 POST /v2/default/transactions key=tx-20
#21 POST /v2/default/transactions key=tx-21
#22 POST /v2/default/transactions key=tx-22
#23 POST /v2/default/transactions key=tx-23
#24 POST /v2/default/transactions key=tx-24
#25 POST /v2/default/transactions key=tx-25
#26 POST /v2/default/transactions key=tx-26
#27 POST /v2/default/transactions key=tx-27
#28 POST /v2/default/transactions key=tx-28
#32 POST /v2/default/transactions/4000000000012/metadata key=md-2
fault concurrent(#32)
```

Delivery order: send #16, send #17, send #18, send #19, send #20, send #21, send #22, send #23, send #24, send #25, send #26, send #27, send #28, send-concurrently #32

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"note":"checked-2"},"targetType":"TRANSACTION"},"idempotency_key":"md-2","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-113/repro.sh`

### Finding 114

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) -1 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: crash(#32@commit1).

```
#16 POST /v2/default/transactions key=tx-16
#17 POST /v2/default/transactions key=tx-17
#18 POST /v2/default/transactions key=tx-18
#19 POST /v2/default/transactions key=tx-19
#20 POST /v2/default/transactions key=tx-20
#21 POST /v2/default/transactions key=tx-21
#22 POST /v2/default/transactions key=tx-22
#23 POST /v2/default/transactions key=tx-23
#24 POST /v2/default/transactions key=tx-24
#25 POST /v2/default/transactions key=tx-25
#26 POST /v2/default/transactions key=tx-26
#27 POST /v2/default/transactions key=tx-27
#28 POST /v2/default/transactions key=tx-28
#32 POST /v2/default/transactions/4000000000012/metadata key=md-2
fault crash(#32@commit1)
```

Delivery order: send #16, send #17, send #18, send #19, send #20, send #21, send #22, send #23, send #24, send #25, send #26, send #27, send #28, send, kill the service after its commit 1, restart, resend #32

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"note":"checked-2"},"targetType":"TRANSACTION"},"idempotency_key":"md-2","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-114/repro.sh`

Hunt time: 712 s.
