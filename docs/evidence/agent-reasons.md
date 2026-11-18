# Screener reasons (private)

Date: 2026-11-18. Protocol: docs/evidence/agent-protocol.md at commit 2bb512e. Input: docs/evidence/agent-prefilter.md (41 repositories).
All code read through the GitHub API (`gh api repos/O/R/git/trees/SHA`, `gh api repos/O/R/contents/PATH?ref=SHA`) at the pinned SHA, plus `gh search code` to locate files. Nothing cloned, downloaded or run.
A repository fails at the first criterion that clearly fails, in the order: exclusions, (P), (a), (b), (c), (d), (e), (f).

## 1. bytedance/deer-flow @ ec3f3d4d22cf87b39ef180c7b2ab782359bca0d4 — PASS

- Exclusions: README calls it a "super agent harness" application with a web UI, not a library/framework/SDK/template. Judged as an app.
- (P): `config.example.yaml` `database:` section supports `backend: postgres` (`postgres_url: $DATABASE_URL`) for the checkpointer, store and app tables. 55 alembic migrations under `backend/packages/harness/deerflow/persistence/migrations/versions/` contain no `CREATE EXTENSION`; `pg_trgm` appears only in `examples/deerflow-extension-workspace-search`. Sandbox default `use: deerflow.sandbox.local:LocalSandboxProvider` (no Docker). Redis only for multi-instance stream bridge/ownership (optional, config comments ~line 1899).
- (a): `deerflow/tools/builtins/update_agent_tool.py` persists through `get_agent_store()` (db store writes the `agents` table, migration 0006_agents); `deerflow/tools/scheduled_tasks.py` `schedule_task(action="create"...)` writes scheduled tasks (migration 0003), enabled by `scheduler.tool_enabled` in config.
- (b): models are configured in config.yaml with `use: langchain_openai:ChatOpenAI` and `base_url:` (e.g. config.example.yaml lines ~374-398); no code change needed.
- (c): `backend/app/gateway/routers/runs.py` `/api/runs/wait` calls `start_run` then `wait_for_run_completion` and returns the final state; `/stream` streams SSE from the in-process run (`services.py` `run_agent(...)` ~line 2938).
- (d): `routers/thread_runs.py` line ~152 `Idempotency-Key` header; `services.py` `_run_idempotency_request` and `command.resume` handling (~lines 1028, 2535) resume a pending HITL interrupt.
- (e): `backend/app/gateway/auth_disabled.py`: `DEER_FLOW_AUTH_DISABLED=1` bypasses auth outside production (documented local/E2E mode).
- (f): `deerflow/config/checkpointer_config.py` `CheckpointerType = Literal["memory","sqlite","postgres"]`; database backend postgres drives the LangGraph checkpointer.

## 2. langchain-ai/langgraph @ 84ced603d3a201ab0f9379936bb35868e243f3f3 — fail (excluded)

- Named exclusion in the protocol.

## 3. langchain-ai/deepagents @ 16966e76e433465fcd8feb3a01534f31095c52a3 — fail (excluded)

- README: "available as a JavaScript/TypeScript library", links "release notes for the Python SDK" (libs/deepagents). A pip-installed agent harness library, not an application.

## 4. xerrors/Yuxi @ 47e4e5a6d905584b3ed47c993c8a7aeb47c15f39 — fail at (P)

- `backend/yuxi/workers/settings.py`: ARQ worker (`redis_settings = get_arq_redis_settings()`) registers `process_agent_run`; agent runs execute in a Redis-queued worker. `docker-compose.yml` `REDIS_URL`, Milvus, Neo4j, MinIO, sandbox containers. Redis queue on the agent path.

## 5. JoshuaC215/agent-service-toolkit @ 3d771312f01d0397da3bc0cd2c2db8d4e1786128 — fail at (a) (screened out)

- No app tables at all. Tools in `src/agents/tools.py` are `calculator` and `Database_Search` (Chroma read). `src/memory/postgres.py` only builds the saver and LangGraph store. No tool writes app-owned PostgreSQL rows.

## 6. GanymedeNil/document.ai @ 265cb9f82eacfc30f31dfbb0aaa009805846fa35 — fail at (P)

- README architecture and `eaer/retrieval/opensearch.py`: retrieval on the query path uses OpenSearch (`docker-compose.yml` service `opensearch`). A vector/search store other than Postgres on the agent path.

## 7. langchain-ai/langgraphjs @ 7257a24f7106ca8b8f5e14da6d1d72197077d6a7 — fail (excluded)

- Named exclusion in the protocol.

## 8. wassim249/fastapi-langgraph-agent-production-ready-template @ fe052a7c6acf692c05ca0e0a50c0ff2ffba448b7 — fail at (a) (screened out)

- Tools: `app/core/langgraph/tools/ask_human.py`, `duckduckgo_search.py`. Neither writes app tables; long-term memory (mem0) is written outside any tool.

## 9. aegra/aegra @ b907dab9740a06e6032da00dd66699a5db2bb23e — fail (excluded)

- README: "drop-in replacement for LangSmith Deployments"; `pip install aegra-cli`, `aegra init` "prompts for location, template, and name". A deployment server plus project template generator that hosts other people's graphs, not an application.

## 10. cuga-project/cuga-agent @ b3a68d5a12b78d980856ce1227bc5e5fbc64d0c7 — fail at (a) (screened out)

- Tools come from external apps through the registry (`src/cuga/backend/tools_env/registry/...`, MCP/OpenAPI). `src/cuga/backend/storage/relational/prod.py` stores CUGA config/policies, written by UI/admin, not by an agent tool. No tool writes CUGA's own tables.

## 11. StreetLamb/tribe @ e1c2020a7dd7377d1ec14292f113c53ddf50d917 — fail at (a) (screened out)

- `backend/app/core/graph/skills/__init__.py` managed skills: duckduckgo-search, wikipedia, yahoo-finance, ask-human; plus user API tools (`api_tool.py`, HTTP) and Qdrant retriever. None writes app tables.

## 12. FareedKhan-dev/production-grade-agentic-system @ d778f208bd6187b8278e4c0a5d4c716757bcfda5 — fail at (a) (screened out)

- Only tool: `src/agent/tools/web_search.py`. App tables (`schema.sql`: users, sessions) are written by API handlers, not tools.

## 13. icey1287/SuperMew @ e2ebabe81e066e0498bdb58000942ef5199d338c — fail at (P)

- `backend/events/bus.py` uses `RedisEventTransport`; `backend/runs/service.py` "HTTP 与执行 worker 共享", `claim_run(worker_id)`; `.env.example` requires `REDIS_URL` and Milvus (`MILVUS_HOST`). Redis and Milvus on the agent path.

## 14. judahpaul16/gpt-home @ 9e00a15acbba924def267f162cb1c49749d60626 — fail at (a) (screened out)

- `src/agent/core.py`: tools are action tools (alarm, calendar, lights, spotify, weather) plus langmem `create_manage_memory_tool` writing the LangGraph `AsyncPostgresStore` (framework store, not app tables). No tool writes app-owned tables.

## 15. wassim249/YT-Navigator @ 61e3ddfbb8c52ed23c01a03cc73498fb48bc7a5a — fail at (a) (screened out)

- `app/services/agent/react_graph.py` tools: `VectorDatabaseTools`, `SQLTools`; `sql_tool.py` line ~86 rejects anything not starting with SELECT. Read-only.

## 16. NicholasGoh/fastapi-mcp-langgraph-template @ 2bd004a51e5d741e8eaf5bb01716b6a673684a53 — fail at (a) (screened out)

- Tools come from the MCP server `backend/shared_mcp/tools.py` (`add`) and a community YouTube MCP. No tool writes app tables.

## 17. martian56/redcell @ 238212cc92fdde729c70f612b195b66693e834d4 — fail at (P)

- README: Redis "bus"; tools run inside a Kali container over `docker exec`. Redis and Docker on the agent path.

## 18. goruck/home-generative-agent @ 7426df6591140590b7b704dbb090007811ce938f — fail at (P)

- README: a Home Assistant integration (custom_components/home_generative_agent). Runs only inside a Home Assistant core process; tools act on HA entities/services. No standalone HTTP request path; HA cannot be replaced by a stub.

## 19. fancyboi999/Loomic @ 21206c30054572291b9df000d84568aeca97aeda — fail at (P)

- Server uses `@supabase/supabase-js` (`apps/server/src/supabase/*`) for data and auth, `apps/server/src/queue/pgmq-client.ts` + `src/worker.ts` (pgmq extension), Supabase storage buckets in migrations. Supabase services and pgmq extension on the path.

## 20. rokbenko/ai-playground @ c825ff0ad704a2b2c1d251119cd0dd19cd85271a — fail (excluded)

- README: "Code from YouTube channel" with a table of contents of tutorials; matched file under `langgraph-tutorials/`. Tutorial collection.

## 21. yuyuanweb/ai-picture-editor @ 49f6ad8c96740a68b7e6228a47e8e2be6876eb34 — fail at (P)

- `backend/app/queue.py`: arq `create_pool(RedisSettings.from_dsn(redis_url))`, `enqueue(task, run_id)`; `services/agent.py` runs are `RunStatus.QUEUED`. Redis queue on the agent path.

## 22. emma-sue/AegisOps @ 211e0fddf193218ebcb31eb0f518591a2d60f127 — fail at (P)

- `services/agent-service/ops_agent/worker.py`: Redis stream `ops:runs` consumer group; graph runs in the worker. Business actions go through a separate Java `ops-api` service. Redis on the agent path.

## 23. KodyKendall/LlamaBot @ 884cd2a491a26e6a01488e7757908429c20dddf1 — fail (excluded)

- README: "LlamaBot is a production-ready AI coding agent framework".

## 24. Idun-Group/idun-agent-platform @ e643895106b1d6195c88749c796a98bc80d1baa3 — fail (excluded)

- README: `idun_agent_engine` — "the engine SDK (FastAPI app factory, ...)"; you `pip install idun-agent-engine` and bring your own agent. SDK/platform.

## 25. xxPcy/PharmaLLM-Agent @ 392e7a7cea5562ad90387c3b80144242d3775bda — fail at (P)

- `PharmaLLM_V1.1/agent_app/src/pharma_agent/jobs.py`: Celery with Redis broker; `compose.yaml` services redis, jaeger; `api.py` `POST /conversations/{id}/runs` returns 202. Redis/Celery on the path.

## 26. chatchat-space/LangGraph-Chatchat @ 0ae53b7ef7f23b04fe37062a8ad031b39e8f11a3 — fail at (a) (screened out)

- `chatchat/server/agent/tools_factory/*`: search, weather, arxiv, calculate, shell, text2image, url_reader, wikipedia, wolfram, text2promql, and `text2sql.py`/`querysql.py` against an externally configured `sqlalchemy_connect_str` (read_only mode, not the app's tables). No tool writes app tables.

## 27. skygazer42/Weaver @ 745693afbf1e8a5b68dad5f616b2681cb2042ee9 — fail at (a) (screened out)

- No migrations or app SQL tables; `tools/core/memory_client.py` uses mem0 or a JSON file (`data/memory_store.json`); other tools are browser/sandbox/search/crawl. No tool writes app PostgreSQL tables.

## 28. nihalashetty/Forge @ 25074df07b9872912902b323706a01dd47cd8477 — PASS

- Exclusions: README "self-hosted platform for visually building, testing, and shipping AI agents"; an application (UI + API), not a library/SDK/template generator.
- (P): `apps/api/forge/config.py`: `database_url`, `vector_backend` "chroma | pgvector", `redis_url: None => in-process fakes`, `checkpoint_backend` "postgres"; README "Swap to Postgres + pgvector, Redis ... with config only". Migrations `apps/api/migrations/versions/0001_baseline.py`, `0002_*` have no `CREATE EXTENSION`. `forge/queue.py`: arq only for webhook/schedule triggers when Redis configured; interactive paths run inline.
- (a): `forge/tools/builtin.py` built-in `remember` tool (line ~155) calls `MemoryService.remember` which `session.add(Memory(...))` + commit (`services/memory.py` line ~27). Built-ins are auto-provisioned into every project.
- (b): `forge/engine/models.py` `init_chat_model("openai:...")` → langchain-openai honours `OPENAI_BASE_URL`; embedder can be `openai:text-embedding-*` per project (`knowledge/embeddings.py` header); outbound tool hosts are user-defined.
- (c): `forge/routers/project_run.py` `POST /run` with `stream=false` calls `run_service.run_to_completion` and returns the reply; `stream=true` streams SSE from the request.
- (d): `project_run.py` `Idempotency-Key` header (line ~103); `body.resume` resumes a HITL interrupt (`services/runs.py` `resume`, `_expand_hitl_decisions`).
- (e): `config.py` `service_api_token` (`Authorization: Bearer <token>`, static, via env) and `auth_required` toggle.
- (f): `config.py` `checkpoint_backend: "postgres"` + `checkpoint_postgres_url`; `forge/main.py` uses PostgresSaver (gh search hit).

## 29. ljxpython/ai-agent-platform @ 84d41a460e40dc630673532a94d4fec8bf48bfc2 — fail at (P)

- `apps/runtime-service/.env.example`: GraphHarbor persistence with `REDIS_URI=redis://...`, API + Worker split ("API/Worker share one execution domain"); README lists Runtime Service tech as "LangGraph + GraphHarbor + Redis" and a Docker sandbox. Redis/worker on the path.

## 30. colossus-lab/openarg_backend @ 58b9a3c038c87c0cde9093813c92bde798087be7 — fail at (P)

- README: "PostgreSQL 16 + pgvector, Redis 7, Celery ... Claude models on AWS Bedrock". `src/app/infrastructure/adapters/cache/redis_cache_adapter.py` backs the answer pipeline cache node (`application/pipeline/nodes/cache.py`). Redis on the path (model is Bedrock as well).

## 31. HsienW/chat-gun @ eb7410ac2b736cbae3006354e7cb42f0d6560105 — fail at (f) (screened out)

- `backend/package.json` `"start": "langgraphjs dev"`; `backend/langgraph.json` serves graphs through the LangGraph JS dev server, whose persistence is in-memory. `PostgresSaver` appears only in `*.integration.test.ts`, `order-branch.test.ts` and docs/specs. No PostgreSQL checkpointer on the workload path.

## 32. agentailor/fullstack-langgraph-nextjs-agent @ 40414f81fb8df752cfc9f41dc2eb2bc34a70b050 — fail at (a) (screened out)

- `src/lib/agent/index.ts`: tools = `configTools` + `getMCPTools()` (user-configured MCP servers). No built-in tool writes the app's Prisma tables.

## 33. luizomf/react_agent_langgraph_course @ 1d6ccf0c16b22725b869e82ab6ccc98efa4b85f6 — fail (excluded)

- README: "Curso LangChain e LangGraph", support material for a YouTube playlist. Course.

## 34. langchain-ai/paid-media-agent @ d4ee460b729ae6a9a7c0fb0e84e7ab07756fd5ee — PASS

- Exclusions: README "A paid-media analyst for your team's Slack". An application; not one of the two named langchain-ai exclusions.
- (P): `src/paid_media_agent/persistence/postgres.py` creates plain tables `pma_proposals`, `pma_approvals`, `pma_receipts`, `pma_dedupe`, `pma_thread_owners` (no extensions). `runtime/sandbox.py`: sandbox is only for Managed Deep Agents; self-hosted runtime uses a local workspace. Sample data mode uses fixtures (`config.py` `paid_media_data_mode` "sample always uses fixtures").
- (a): `tools/write_tools.py` `propose_change` → `ProposalService.propose` → `PostgresProposalRepository.save` INSERT INTO `pma_proposals`; `execute_change` writes receipts/approval use. Write tools are always assembled (`assembly.py` line ~241).
- (b): `.env.example` `PAID_MEDIA_MODEL=provider:model`, `PAID_MEDIA_MODEL_BASE_URL`, "OpenAI-compatible endpoints (set PAID_MEDIA_MODEL_BASE_URL and PAID_MEDIA_MODEL_API_KEY_ENV)". Ad platform hosts unused in sample mode.
- (c): `surfaces/api/app.py` `POST /threads/{thread_id}/messages` awaits `runner.send(...)` (graph `ainvoke`) and returns the outcome.
- (d): `POST /proposals/{proposal_id}/approve` → `runner.approve` → `resume` → `Command(resume={"decisions": [...]})` on a pending HITL interrupt (`surfaces/runner.py` lines ~241-276; `write_tools.py` `InterruptOnConfig`).
- (e): `app.py` `caller()`: static bearer tokens from `PAID_MEDIA_API_TOKENS` (env).
- (f): `runtime/self_hosted.py` `AsyncPostgresSaver(pool)` when `DATABASE_URL` is set (lines ~45-85).

## 35. spacesky-cell/TwinFix @ 27ab18097f46f674b05c1058737b3641b972d77e — fail at (P)

- `server/twinfix/api/app.py` lines ~35-37: `Redis.from_url(settings.redis_url)`, `RedisPublisher`, `RabbitDispatcher.connect(settings.amqp_url)`; incident routes return 202 and the engine runs in the RabbitMQ worker. Redis and a message queue on the path.

## 36. lucasboscatti/Whatsapp-Langgraph-Agent-Integration @ 10d307c96201e4a0fc88fbcedf5d9406fd7ac582 — fail at (a) (screened out)

- `app/agent.py`: a chat graph with AsyncPostgresSaver and no tools. No app tables.

## 37. 01-ai/langcrew @ 008c3407b2adb3f0d7058bc8b565a393920d81f0 — fail (excluded)

- README: "LangCrew is a high-level multi-agent development framework built on LangGraph".

## 38. BharathxD/ClaimeAI @ 8b13492d1a5f0e91b21149e9f27c586e1d12f996 — fail at (P)

- `apps/agent/utils/redis.py` + `apps/agent/security/auth.py` / `api_keys.py`: request auth reads API keys from Redis. Redis on the request path (and no tool writes app tables either).

## 39. Kyle-liu888/langgraph-trip-planner @ 04508887af99760be7b6ef759c7ca17021f61804 — fail at (a) (screened out)

- `backend/app/graph/nodes.py`, `builder.py`: no tools bound; trips are written by `services/run_manager.py` outside the graph (and in an `asyncio.create_task` background run).

## 40. realyinchen/AgentHub @ 92998975d760fd8ba4672733d604e37dbada87b8 — fail at (a) (screened out)

- `backend/app/agents/tools/`: `time.py`, `vectorstore_retriever.py`, `web.py`. All read-only.

## 41. gotohuman/gotohuman-langgraph-lead-example @ 7f5c4dc70b9c04ec8e1d9da715ad0a0f4100839b — fail at (a) (screened out)

- Single route `src/app/api/agent/route.ts`; the agent drafts an email and sends it for review to the external gotoHuman service. No app tables.

## Totals

Passed: 3 (bytedance/deer-flow, nihalashetty/Forge, langchain-ai/paid-media-agent). Judged: 41 of 41. List exhausted before 4 passes.
