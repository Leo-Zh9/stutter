# Blind screen (Lane 2)

Screener: Agent B, on 2026-11-14 (UTC), under [blind-protocol.md](blind-protocol.md) as committed in c7ee40d. Search output: [blind-search.json](blind-search.json), 31 repositories (sha256 `ec13ef4a36e0e9c1fa870d0733e502dcf10f7db9b271c3e622c2e25ae3454078`).

Rank is the position in `blind-search.json`. Ranks 3 (nextjs/saas-starter) and 29 (stripe/sync-engine) are excluded by the protocol and were not judged. The pinned SHA is the HEAD of each default branch on the screening day, read with `gh api repos/OWNER/REPO/commits/DEFAULT_BRANCH --jq .sha`.

The list was exhausted with 3 passes, fewer than the 4 the protocol seeks.

| Rank | Repository | Stars | Pinned SHA | Result |
|---|---|---|---|---|
| 1 | better-auth/better-auth | 30748 | `8382ce613f04bf590d3cb6c1363dfd0e41110944` | fail |
| 2 | kriasoft/react-starter-kit | 23717 | `dbdb0d4991bbaae1671ff62180089f2733dd406a` | pass |
| 4 | vercel/next-forge | 7780 | `f189de79ceef7c1ef69f61f12e272f99b4cdb699` | fail |
| 5 | yournextstore/yournextstore | 5538 | `b0caa2dd5f0bf33f1fabd4bdce904a56de1fcaf4` | fail |
| 6 | sadmann7/skateshop | 5429 | `e954d54379a7a6f798d3f49ffc6664449220adfc` | screened out: the workload cannot reach it |
| 7 | stripe/stripe-node | 4570 | `6237106099542b8e0a3b52ef78308b3f678e93a6` | fail |
| 8 | async-labs/builderbook | 3790 | `a71c72e21d8aeb9f00da33d5199b47dd90b7aef0` | fail |
| 9 | mickasmt/next-saas-stripe-starter | 3018 | `c6df229ae73c207e87356d5e592779c511a9b56d` | fail |
| 10 | nextify-limited/saasfly | 2946 | `49f7e28f69eae9cd9eed84221e13a3dbae87da67` | fail |
| 11 | juicycleff/ultimate-backend | 2907 | `0ae62a4b894c0e46baf75228ec88fae98f6abe06` | fail |
| 12 | adrianhajdin/ecommerce_sanity_stripe | 2313 | `c4d1c0593769004d912bedbfe269ed9180f74939` | fail |
| 13 | talivia-group/talivia | 2290 | `9b7ad2a6039bce79e20364dbcea44c66a765f398` | pass |
| 14 | mrvautin/expressCart | 2250 | `b31302f4e99c3293bd742c6d076a721e168118b0` | fail |
| 15 | stripe/ai | 2063 | `74f7f39b1d7f3b56e49218a43b0dc87a268c03c7` | fail |
| 16 | stripe/react-stripe-js | 2046 | `d6670a276b1271b934b3fd58359ec554a6fe30ec` | fail |
| 17 | ParabolInc/parabol | 2018 | `9edd33ac1de635b1ae9406236825ee3bc9d484e2` | fail |
| 18 | agentic-commerce-protocol/agentic-commerce-protocol | 1602 | `7fdd78df677a94dce04c770644b0fbbb1401272b` | fail |
| 19 | reliverse/relivator | 1567 | `a1871b006ab09df647b99fc71d3c080acd797e24` | fail |
| 20 | dev-xo/remix-saas | 1452 | `cdedef9402664aae3d451217103bf2443acdf2f3` | fail |
| 21 | adrianhajdin/ecommerce | 1443 | `7c70b5ccb755c04be4148af542e725a43d4717ba` | fail |
| 22 | stripe/stripe-react-native | 1432 | `13f0e79823920f0ca60870d35b20c43622b150d8` | fail |
| 23 | nextacular/nextacular | 1398 | `06bee752a0423faa2ba0217ea2fedd719a52bda9` | screened out: the workload cannot reach it |
| 24 | adrianhajdin/ai_saas_app | 1367 | `f82bfdd03c97000eea0d7fada94a307e61b58b5f` | fail |
| 25 | moinulmoin/chadnext | 1319 | `1a1f25a6de88453bd3c80efa8c3f61f9c159661c` | screened out: the workload cannot reach it |
| 26 | RevenueCat/react-native-purchases | 1223 | `fbebad47e865a5a7ae91663d124f2eeca5b35ba9` | fail |
| 27 | vue-stripe/vue-stripe | 1126 | `9a8b0754d6d7951a8d6508502886df6b9cce9060` | fail |
| 28 | ShenSeanChen/launch-mvp-stripe-nextjs-supabase | 1108 | `513a8f0ca6e405725dcf98eab1fc5cd6468b5f10` | fail |
| 30 | getpaykit/paykit | 1077 | `c7572ff7e0116428a43a2f98ad25dc2c26c8927e` | pass |
| 31 | Skolaczk/next-starter | 1003 | `a4d25f24c1761a9476c81c2ef15f34b5897a0261` | fail |

Passed 3, failed 23, screened out 3 (29 judged).

The screener's reasons are in a file kept outside the repository until every hunt is pushed. Its sha256:

```
fce0217c659c8c9bafa536a43ba4dd4b1ad505916421cfbea87dc83df5243264  blind-reasons.md
```
