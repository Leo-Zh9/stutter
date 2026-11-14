# Blind hunt protocol (Lane 2)

These rules were written and committed before any blind target was chosen. They do not change after that commit. If a rule turns out to be wrong, the run is reported under the rule as written, and the problem is stated in the write-up.

## Roles

- **Screener.** Agent B. Runs the search, judges each repository against the criteria, and publishes pass or fail only.
- **Hunter.** A different agent. It receives only repository names and pinned commits. It never sees the screener's reasons until every hunt is pushed.

## The workload

The workload is `targets/stripe/lifecycles.mjs`, unchanged. Its sha256 is recorded in [workload.sha256](workload.sha256):

```
bc121c1f1bee7f2f8a4bb624c4a5767e2c775e068d8dde42c09474cfb345cdb6  targets/stripe/lifecycles.mjs
```

It sends 16 signed `customer.subscription.created`, `customer.subscription.updated` and `customer.subscription.deleted` events across five customers, 60 seconds apart, each with `metadata: {}`.

Before every `stutter record` and every `stutter hunt`, the hunter runs this from the repo root and saves the output in the target's results directory:

```bash
shasum -a 256 -c docs/evidence/workload.sha256
```

A run without a passing check does not count.

## The search

The screener runs this from the repo root, verbatim, on the screening day:

```bash
{ gh search repos stripe --language=TypeScript --stars=">=1000" --archived=false --include-forks=false --sort=stars --order=desc --limit=200 --json fullName,url,stargazersCount,language,defaultBranch,isFork,isArchived,pushedAt;
  gh search repos stripe --language=JavaScript --stars=">=1000" --archived=false --include-forks=false --sort=stars --order=desc --limit=200 --json fullName,url,stargazersCount,language,defaultBranch,isFork,isArchived,pushedAt; } \
  | jq -s 'add | unique_by(.fullName) | sort_by(-.stargazersCount)' > docs/evidence/blind-search.json
```

The search order is the order of `blind-search.json`: most stars first.

## Exclusions

These are removed from the list before screening. They have been read, or they are forks of a target that has been read:

- `nextjs/saas-starter` and any fork of it;
- `boxyhq/saas-starter-kit` and any fork of it;
- `stripe/sync-engine` (also known as `supabase/stripe-sync-engine`);
- `lukevella/rallly`.

## Pinned commit

Each repository is judged at the HEAD of its default branch on the screening day. The screener records the SHA from:

```bash
gh api repos/OWNER/REPO/commits/DEFAULT_BRANCH --jq .sha
```

The hunter uses that SHA and no other.

## Pass criteria

A repository passes only if every test below holds at the pinned commit.

- **(P) It can run here.** PostgreSQL is its store. Its migrations need no extension beyond `plpgsql`, `vector` and the project's `pgcrypto` `digest()` shim. Running the webhook path needs no Docker, no Redis and no other service.
- **(a) A Stripe webhook route exists.** The app serves an HTTP route that verifies a Stripe signature and handles Stripe events.
- **(b) The path makes no outside calls.** Handling a `customer.subscription.*` event on that route makes no call to the Stripe API, Redis, a queue, or any other outside service.
- **(c) The route writes PostgreSQL on a subscription event.** At least one of `customer.subscription.created`, `customer.subscription.updated` or `customer.subscription.deleted` causes a write to PostgreSQL.
- **(d) It finds the app's row by Stripe id.** The handler finds the row it changes by the Stripe customer id or subscription id in the event payload. Fixtures can supply such a row. A handler that finds the row by `metadata` fails, because the workload sends empty metadata.

A repository that fails (c) or (d) is listed as **screened out: the workload cannot reach it**. It is never reported as clean.

## Refill

The screener judges repositories in search order until 4 pass or the list is exhausted. If a passing target is later reported as not run, the screener continues from where it stopped, until 4 targets have run or the list is exhausted.

## What the screener publishes

1. `docs/evidence/blind-search.json`: the raw search output.
2. `docs/evidence/blind-screen.md`: one row per repository judged, in search order. Columns: rank, repository, stars, pinned SHA, pass or fail. It gives no reasons.
3. The sha256 of the screener's reasons file, in `blind-screen.md`. The reasons file stays outside the repo. It is committed only after every hunt is pushed, and its hash must match.

## Allowed per-target inputs

The hunter may change only these:

- `targets/<name>/stutter.json`;
- `targets/<name>/fixtures.sql`. Its rows are derived from the target's database schema (migrations or schema file) and its env example;
- three generator arguments, recorded in the target's results:
  - `--api basil|acacia`: `basil` if the target pins `stripe` 18 or later in `package.json`, `acacia` otherwise;
  - `--sdk`: the target's own `node_modules/stripe` if it installs one. Otherwise the fixed copy, `stripe@18.1.0` at `.targets/saas-starter/node_modules/stripe`;
  - `--customers`: the Stripe customer ids the fixtures hold.

`--url` and `--secret` are mechanical. They come from the recording proxy and from `stutter.json`.

The hunter may read the README, the docs, the env example and the database schema. It may not read the webhook handler or any other application source until step 4 below.

If the fixtures cannot satisfy the workload, the target is reported as **not run**. The workload, the generator and the target's code are never patched to make it run.

## Order of work, per target

1. Check out the pinned SHA, install and build.
2. Write `stutter.json` and `fixtures.sql`.
3. Check the workload hash. Record, check the hash again, and hunt. Commit config, trace and results, and push them.
4. Only after that push, read the handler and write the root causes.

## Reporting

- **M** is the number of distinct root causes per target. A root cause is one defect in the code that needs its own fix. Failing pairs and their effects are reported separately, with their counts.
- A hunt where no `reorder` pair fires is reported as **not reached**, not as clean.
- Clean, not reached, not run and screened out are each counted and listed.

## The claim

If the hunts find defects:

> Against a fixed Stripe webhook workload (sha256 bc121c1f), used unmodified, stutter found M distinct webhook-ordering root causes in K open-source apps screened by a fixed, published repo search and pass/fail criteria.

If they find fewer or none, the claim is the same sentence with the outcome counts:

> Against a fixed Stripe webhook workload (sha256 bc121c1f), used unmodified, stutter ran on K open-source apps screened by a fixed, published repo search and pass/fail criteria. It found M distinct root causes. C apps were clean, R were not reached and N were not run.
