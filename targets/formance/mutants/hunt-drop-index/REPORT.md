# stutter hunt: Formance Ledger v2.4.18 mutant: unique index logs_idempotency_key dropped

Generated 2026-11-14T12:22:54Z. Trace: 43 requests. Pre-flight: every recorded webhook delivery got 2xx.

| Fault type | Schedules | Fired | Not fired | Changed the final state | No-op passes |
|---|---|---|---|---|---|
| concurrent | 43 | 43 | 0 | 38 | 0 |
| crash | 43 | 43 | 0 | 0 | 0 |
| lost | 43 | 43 | 0 | 0 | 0 |

No-op reorder pairs (dropping both events leaves the state unchanged, so their pass says nothing about ordering): 0.

## Effects: 38 distinct end-state changes

1. Findings [1]:

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

2. Findings [2]:

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

3. Findings [3]:

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

4. Findings [4]:

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

5. Findings [5]:

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

6. Findings [6]:

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

7. Findings [7]:

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

8. Findings [8]:

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

9. Findings [9]:

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

10. Findings [10]:

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

11. Findings [11]:

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

12. Findings [12]:

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

13. Findings [13]:

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

14. Findings [14]:

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

15. Findings [15]:

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

16. Findings [16]:

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

17. Findings [17]:

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

18. Findings [18]:

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

19. Findings [19]:

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

20. Findings [20]:

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

21. Findings [21]:

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

22. Findings [22]:

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

23. Findings [23]:

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

24. Findings [24]:

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

25. Findings [25]:

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

26. Findings [26]:

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

27. Findings [27]:

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

28. Findings [28]:

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

29. Findings [29]:

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

30. Findings [30]:

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

31. Findings [31]:

```
+ _default.logs {"data":{"metadata":{"note":"checked-0"},"targetType":"TRANSACTION"},"idempotency_key":"md-0","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

32. Findings [32]:

```
+ _default.logs {"data":{"metadata":{"note":"checked-3"},"targetType":"TRANSACTION"},"idempotency_key":"md-3","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

33. Findings [33]:

```
+ _default.logs {"data":{"metadata":{"note":"checked-4"},"targetType":"TRANSACTION"},"idempotency_key":"md-4","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

34. Findings [34]:

```
+ _default.logs {"data":{"metadata":{"tier":"gold-0"},"targetType":"ACCOUNT"},"idempotency_key":"am-0","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

35. Findings [35]:

```
+ _default.logs {"data":{"metadata":{"tier":"gold-2"},"targetType":"ACCOUNT"},"idempotency_key":"am-2","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

36. Findings [36]:

```
+ _default.logs {"data":{"metadata":{"tier":"gold-1"},"targetType":"ACCOUNT"},"idempotency_key":"am-1","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

37. Findings [37]:

```
+ _default.logs {"data":{"metadata":{"note":"checked-1"},"targetType":"TRANSACTION"},"idempotency_key":"md-1","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

38. Findings [38]:

```
+ _default.logs {"data":{"metadata":{"note":"checked-2"},"targetType":"TRANSACTION"},"idempotency_key":"md-2","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

## Findings: 38 distinct minimal reproductions

### Finding 1

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

Repro: `finding-1/repro.sh`

### Finding 2

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

Repro: `finding-2/repro.sh`

### Finding 3

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

Repro: `finding-3/repro.sh`

### Finding 4

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

Repro: `finding-4/repro.sh`

### Finding 5

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

Repro: `finding-5/repro.sh`

### Finding 6

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

Repro: `finding-6/repro.sh`

### Finding 7

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

Repro: `finding-7/repro.sh`

### Finding 8

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

Repro: `finding-8/repro.sh`

### Finding 9

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

Repro: `finding-9/repro.sh`

### Finding 10

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

Repro: `finding-10/repro.sh`

### Finding 11

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

Repro: `finding-11/repro.sh`

### Finding 12

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

Repro: `finding-12/repro.sh`

### Finding 13

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

Repro: `finding-13/repro.sh`

### Finding 14

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

Repro: `finding-14/repro.sh`

### Finding 15

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

Repro: `finding-15/repro.sh`

### Finding 16

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

Repro: `finding-16/repro.sh`

### Finding 17

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

Repro: `finding-17/repro.sh`

### Finding 18

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

Repro: `finding-18/repro.sh`

### Finding 19

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

Repro: `finding-19/repro.sh`

### Finding 20

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

Repro: `finding-20/repro.sh`

### Finding 21

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

Repro: `finding-21/repro.sh`

### Finding 22

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

Repro: `finding-22/repro.sh`

### Finding 23

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

Repro: `finding-23/repro.sh`

### Finding 24

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

Repro: `finding-24/repro.sh`

### Finding 25

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

Repro: `finding-25/repro.sh`

### Finding 26

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

Repro: `finding-26/repro.sh`

### Finding 27

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

Repro: `finding-27/repro.sh`

### Finding 28

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

Repro: `finding-28/repro.sh`

### Finding 29

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

Repro: `finding-29/repro.sh`

### Finding 30

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

Repro: `finding-30/repro.sh`

### Finding 31

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

Repro: `finding-31/repro.sh`

### Finding 32

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

Repro: `finding-32/repro.sh`

### Finding 33

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

Repro: `finding-33/repro.sh`

### Finding 34

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

Repro: `finding-34/repro.sh`

### Finding 35

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#42).

```
#27 POST /v2/default/transactions key=tx-27
#42 POST /v2/default/accounts/sellers:2/metadata key=am-2
fault concurrent(#42)
```

Delivery order: send #27, send-concurrently #42

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"tier":"gold-2"},"targetType":"ACCOUNT"},"idempotency_key":"am-2","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-35/repro.sh`

### Finding 36

Reproduced in 10 of 10 fresh replays and by repro.sh (curl, fresh databases; -1 = crash finding, not scriptable) 10 of 10 times; 1-minimal: true. Reached from 1 single-fault schedules: concurrent(#41).

```
#27 POST /v2/default/transactions key=tx-27
#28 POST /v2/default/transactions key=tx-28
#35 POST /v2/default/transactions/4000000000001/revert key=rv-0
#41 POST /v2/default/accounts/sellers:1/metadata key=am-1
fault concurrent(#41)
```

Delivery order: send #27, send #28, send #35, send-concurrently #41

State diff (+ only in the faulted run, - only in the in-order run):

```
+ _default.logs x1 {"data":{"metadata":{"tier":"gold-1"},"targetType":"ACCOUNT"},"idempotency_key":"am-1","ledger":"default","schema_version":null,"type":"SET_METADATA"}
```

Repro: `finding-36/repro.sh`

### Finding 37

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

Repro: `finding-37/repro.sh`

### Finding 38

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

Repro: `finding-38/repro.sh`

Hunt time: 399 s.
