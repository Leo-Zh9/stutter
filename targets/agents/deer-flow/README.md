# bytedance/deer-flow @ ec3f3d4: not run

Hunted under [the agent protocol](../../../docs/evidence/agent-protocol.md) on 2026-11-18. Reported **not run**: the allowed inputs could not make a workload in which the agent calls a tool that writes the app's tables, on a path that needs no login.

## What was tried

Only the README, `docs/`, `backend/docs/`, `config.example.yaml`, the served `/openapi.json` and the tool schemas the app sent to the model were read.

1. Installed `backend` with `uv sync --locked --extra postgres` and ran the Gateway (`uvicorn app.gateway.app:app`) on PostgreSQL with [config.yaml](config.yaml): `config.example.yaml` with one OpenAI-compatible model pointed at the scripted model, the scheduler and its conversation tool on, and `hitl_approval` asking a human before `schedule_task`.
2. The tool the docs name for writing durable app state from a conversation is `schedule_task` (`scheduled_tasks`, `scheduled_task_runs`). The docs offer it to "authorized interactive" turns.
3. The scripted model logged the tools offered on a fresh turn of `POST /api/threads/{id}/runs/wait`:
   - **Internal auth** (`X-DeerFlow-Internal-Token`, the static deployment token the auth design documents for server-to-server use): 15 tools, no `schedule_task`. The Gateway also refused it without a CSRF double-submit pair, though the auth design says internal auth takes no part in CSRF.
   - **Personal access token** (scopes `threads:*`, `runs:*`): the same 15 tools, no `schedule_task`.
   - **Browser session** (cookie from `POST /api/v1/auth/initialize`): 16 tools, including `schedule_task`.

None of the 15 tools offered to a static token writes the app's PostgreSQL tables. They work on files, the web, uploads and run status, and `present_shared_artifact` needs an artifact from another conversation. The table-writing path needs a login, which criterion (e) excludes.

## Under the protocol

The screener passed this repository, so the hunter reports what the run found: **not run**, not clean and not screened out. The refill rule cannot apply, because the screen judged all 41 pre-filtered repositories.
