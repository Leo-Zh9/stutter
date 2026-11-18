# Blind hunt protocol: LLM agent apps

These rules were written and committed before any agent target was chosen. They do not change after that commit. If a rule turns out to be wrong, the run is reported under the rule as written, and the problem is stated in the write-up. The structure follows the Stripe protocol, [blind-protocol.md](blind-protocol.md).

## What is tested

An agent app's own side effects under the faults a real deployment produces:

- a retried chat message (`lost`, `concurrent`), when the request carries a per-action id;
- a double-submitted or retried approval of a pending `interrupt()` (`concurrent`, `lost`);
- a process crash right after any database commit, including the framework's checkpoint writes, then restart and the client's retry (`crash`).

The rule is Stutter's: a retried, duplicated or crashed request must leave the app's tables exactly as a fault-free run of the same workload does. The framework's checkpoint tables are not compared.

## The model

There is no real model and no recording. Each target gets a frozen model script (see [agent mode](../agents.md)): for each user message, the tool calls the model makes, step by step, then a plain reply. The scripted model is a pure function of the conversation it is sent, so a node that LangGraph re-runs gets the same prompt and makes the same tool call, with fresh tool-call ids as a real model would. A request the script does not cover is a script miss: in a fault-free run the target is reported **not run**; in a faulted run that schedule is **diverged**, neither a pass nor a failure.

The model's choices are scripted, not sampled. This is the same position as the Stripe workload: the workload is fixed, and the app code is what is tested.

## Roles

- **Screener.** One agent. Runs the search, judges each repository against the criteria by reading its code, and publishes pass or fail only, with the sha256 of its private reasons file.
- **Hunter.** A different agent. It receives only repository names and pinned commits. It never sees the screener's reasons until every hunt is pushed.

## The search

The screener runs these from the repo root, verbatim, on the screening day:

```bash
{ gh search repos langgraph --language=Python --stars=">=100" --archived=false --include-forks=false --sort=stars --order=desc --limit=200 --json fullName,url,stargazersCount,language,defaultBranch,isFork,isArchived,pushedAt;
  gh search repos langgraph --language=TypeScript --stars=">=100" --archived=false --include-forks=false --sort=stars --order=desc --limit=200 --json fullName,url,stargazersCount,language,defaultBranch,isFork,isArchived,pushedAt; } \
  | jq -s 'add | unique_by(.fullName) | sort_by(-.stargazersCount)' > docs/evidence/agent-search.json
```

The search order is the order of `agent-search.json`: most stars first.

**Pre-filter (mechanical).** Before judging, the screener keeps only repositories whose default branch, at the pinned commit, declares a PostgreSQL checkpointer as a dependency: `langgraph-checkpoint-postgres` in any `pyproject.toml`, `requirements*.txt`, `setup.py`, `setup.cfg` or `uv.lock`/`poetry.lock`; or `@langchain/langgraph-checkpoint-postgres` in any `package.json`. It publishes the kept list, in search order, as `agent-prefilter.md`, with the file and line that matched.

## Exclusions

- `langchain-ai/langgraph`, `langchain-ai/langgraphjs` and any repository whose README says it is a library, framework, SDK, course, tutorial collection or template generator rather than an application. These cannot be judged against an app's side effects.

## Pinned commit

Each repository is judged at the HEAD of its default branch on the screening day:

```bash
gh api repos/OWNER/REPO/commits/DEFAULT_BRANCH --jq .sha
```

## Pass criteria

A repository passes only if every test holds at the pinned commit.

- **(P) It can run here.** PostgreSQL is its store. Its migrations need no extension beyond `plpgsql`, `vector` and the project's `pgcrypto` `digest()` shim. The request path that runs the agent needs no Docker, no Redis, no message queue and no other service that cannot be pointed at a local stub.
- **(a) A tool writes the app's tables.** At least one tool the agent can call inserts, updates or deletes rows in the app's own PostgreSQL tables (not only the checkpointer's).
- **(b) Every outbound host is configurable.** The model provider is OpenAI-compatible and its base URL is settable by environment or config without patching code. Every other outside host the workload path calls has a base URL settable the same way. No MITM proxy is used.
- **(c) The agent runs inside the request.** The graph runs inside the HTTP request (blocking or streamed), not in a background worker the request only enqueues.
- **(d) A retry is meaningful.** At least one request on the workload path either carries a per-action client id (a message id, an interrupt id, a tool-call id) in its body or headers, or resumes a pending dynamic `interrupt()` (a human approval). A thread id alone does not key a chat message. Static breakpoints (`interrupt_before`, `interrupt_after`) are out of scope.
- **(e) No login.** The workload path needs no authentication, or accepts a static token settable by environment.
- **(f) The checkpointer is PostgreSQL** on the workload path (not the in-memory saver), because crash points are commits.

A repository that fails (a) or (f) is listed as **screened out: no durable side effect to test**. It is never reported as clean.

## Refill

The screener judges repositories in pre-filter order until 4 pass or the list is exhausted. If a passing target is later reported as not run, the screener continues from where it stopped, until 4 targets have run or the list is exhausted.

## What the screener publishes

1. `docs/evidence/agent-search.json`: the raw search output.
2. `docs/evidence/agent-prefilter.md`: the kept list with the matching dependency line.
3. `docs/evidence/agent-screen.md`: one row per repository judged, in order: rank, repository, stars, pinned SHA, pass or fail. No reasons.
4. The sha256 of the screener's reasons file, in `agent-screen.md`. The reasons file stays outside the repo and is committed only after every hunt is pushed.

## Allowed per-target inputs

The hunter may read the README, the docs, the env example, the database schema or migrations, any API description the app serves itself (for example FastAPI's `/openapi.json`), and the tool names and JSON schemas the app sends to the model (the stub logs them). It may not read handler, graph, tool or other application source until step 4.

The hunter writes only:

- `targets/agents/<name>/stutter.json`, including `retry_rules`, `ignore_tables` and `headers`;
- `targets/agents/<name>/fixtures.sql`, derived from the schema;
- `targets/agents/<name>/model.json`, the model script: one turn per workload message, each using only tools the app offers, with arguments that satisfy their schemas;
- `targets/agents/<name>/workload.*`, the requests a user would make: at least one conversation that calls a tool that writes the app's tables, and, if the app has approvals, one approval.

If these inputs cannot make the workload run, the target is reported **not run**. The app's code is never patched to make it run, except to point it at the stub through its own configuration.

## Order of work, per target

1. Check out the pinned SHA, install and build.
2. Write the inputs above. Record the workload with `stutter record`.
3. Run `stutter hunt`. Commit config, script, trace and results, and push them.
4. Only after that push, read the code and write root causes.

## What counts

**A finding** is a schedule that, against the fault-free run of the same workload, duplicates or drops a row in the app's business tables, is 1-minimal or reported as not, reproduces in 10 of 10 fresh replays, and has an app-level fix (a claim or lock, a transaction around claim and action, an idempotency key the app owns, or an outbox).

**Attribution.** LangGraph documents that a node re-runs from its start on resume. An unguarded side effect in app code is therefore an app root cause, even when it resembles a filed LangGraph issue (#8039, #9185). "Known upstream" applies only to a framework defect with no app-level fix.

**Not counted:**

- diffs only in checkpoint tables or in chat-transcript tables (listed as cosmetic);
- diverged runs, and crash faults whose two runs disagree (nondetermined);
- a re-run with no app-visible effect.

**M** is the number of distinct root causes: one defect that needs its own fix. Failing schedules and effects are reported separately.

**Not reached:** no fault fires deterministically (fired, not diverged, not nondetermined) on any request whose fault-free run writes the app's business tables. A not-reached target is never reported as clean.

## The claim

> Against frozen per-target model scripts, stutter ran on K open-source LangGraph apps chosen by a fixed, published repo search and pass/fail criteria. It found M distinct root causes where a tool's side effect ran twice or never. C apps were clean, R were not reached and N were not run.

No issue is filed upstream by this project. Issue drafts are written into each target's doc.
