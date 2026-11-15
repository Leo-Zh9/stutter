#!/usr/bin/env bash
# Sets up the real targets in .targets/ (gitignored), at pinned versions,
# without changing their code:
#   nextjs/saas-starter      commit 6e33e58 (pnpm, compile-only next build)
#   boxyhq/saas-starter-kit  commit abc9b68 (npm, prisma generate, compile-only next build)
#   stripe/sync-engine       commits 96e6b39 (pre-fix) and d486b41 (fix, PR #177) (pnpm build)
#   talivia-group/talivia    commit 9b7ad2a (pnpm@10.10.0, prisma generate, compile-only next build)
#   getpaykit/paykit         commit c7572ff (pnpm@12.4.2, compile-only next build of apps/demo)
#   kriasoft/react-starter-kit commit dbdb0d4 (npx bun@1.4.2 install, @repo/email build)
#   Formance Ledger          v2.4.18 release binary, checksum-verified, plus the
#                            v2.4.18 source and the mutant builds in targets/formance/mutants
# Needs git, node (with pnpm), Go, and network access. About 2 GB on disk.
set -euo pipefail
cd "$(dirname "$0")/.."
ROOT="$PWD"
GO="${GO:-$(command -v go || echo "$HOME/sdk/go/bin/go")}"
mkdir -p .targets
cd .targets

clone() { # dir url commit
  if [ ! -d "$1/.git" ]; then git clone -q "$2" "$1"; fi
  git -C "$1" checkout -q "$3"
}

# Compile-only builds: a full `next build` prerenders pages that call the live
# Stripe API, which these runs never do.
export NEXT_TELEMETRY_DISABLED=1 STRIPE_SECRET_KEY=sk_test_offline_never_used STRIPE_WEBHOOK_SECRET=whsec_stutter_local

clone saas-starter https://github.com/nextjs/saas-starter 6e33e58b1e553a41fe22e6b941a7229a002de361
(cd saas-starter && pnpm install --frozen-lockfile &&
  POSTGRES_URL=postgres://unused BASE_URL=http://unused AUTH_SECRET=unused \
  node_modules/.bin/next build --experimental-build-mode=compile)

clone saas-starter-kit https://github.com/boxyhq/saas-starter-kit abc9b686823cbfb4973c79bc36fea37a3244be6c
(cd saas-starter-kit && npm ci --no-audit --no-fund &&
  DATABASE_URL=postgres://unused node_modules/.bin/prisma generate &&
  DATABASE_URL=postgres://unused NEXTAUTH_URL=http://unused NEXTAUTH_SECRET=unused APP_URL=http://unused FEATURE_TEAM_PAYMENTS=true \
  node_modules/.bin/next build --experimental-build-mode=compile)

# stripe/sync-engine before and after PR #177 (an older webhook could overwrite a
# newer row): 96e6b39 is the first parent of the merge d486b41. `pnpm build` runs
# tsup for the engine and tsc for the fastify app; neither calls Stripe.
clone sync-engine-pre https://github.com/stripe/sync-engine 96e6b398d45d0c57ff2193fa4123fa988b1c4e45
(cd sync-engine-pre && pnpm install --frozen-lockfile && pnpm build)
clone sync-engine-fix https://github.com/stripe/sync-engine d486b41dc1b18626ef798890b5bbc7f11f476517
(cd sync-engine-fix && pnpm install --frozen-lockfile && pnpm build)

# talivia-group/talivia (blind target): pnpm at the version its packageManager
# field pins. prisma generate needs DATABASE_URL set but does not connect;
# build-prisma-client.js is the repo's own esbuild bundle of the generated
# client. The compile-only next build skips prerender. None of this calls Stripe.
clone talivia https://github.com/talivia-group/talivia 9b7ad2a6039bce79e20364dbcea44c66a765f398
(cd talivia && npx -y pnpm@10.10.0 install --frozen-lockfile &&
  DATABASE_URL=postgres://unused CHECKPOINT_DISABLE=1 PRISMA_HIDE_UPDATE_MESSAGE=1 node_modules/.bin/prisma generate &&
  node scripts/build-prisma-client.js &&
  DATABASE_URL=postgres://unused APP_SECRET=unused node_modules/.bin/next build --webpack --experimental-build-mode=compile)

# getpaykit/paykit (embedded Stripe billing framework). paykitjs is the library;
# apps/demo is the Next.js app that mounts the paykit handler at /paykit/webhook.
# The library installs its own stripe@19 under packages/paykit. The compile-only
# demo build skips prerender, which would otherwise touch Stripe; env validation
# is skipped the same way. Neither step calls Stripe.
clone paykit https://github.com/getpaykit/paykit c7572ff7e0116428a43a2f98ad25dc2c26c8927e
(cd paykit && npx -y pnpm@12.4.2 install --frozen-lockfile &&
  cd apps/demo && SKIP_ENV_VALIDATION=1 STRIPE_SECRET_KEY=sk_test_offline_never_used \
  STRIPE_WEBHOOK_SECRET=whsec_stutter_local PAYKIT_DATABASE_URL=postgres://unused \
  AUTH_DATABASE_URL=postgres://unused APP_URL=http://unused BETTER_AUTH_SECRET=unused \
  node_modules/.bin/next build --experimental-build-mode=compile)

# kriasoft/react-starter-kit (blind target): Bun only, taken from the npm
# registry at the version its packageManager field pins. The API's dev server
# imports the built @repo/email package, so that is built too (tsc only).
clone react-starter-kit https://github.com/kriasoft/react-starter-kit dbdb0d4991bbaae1671ff62180089f2733dd406a
(cd react-starter-kit && npx -y bun@1.4.2 install --frozen-lockfile &&
  npx -y bun@1.4.2 --filter @repo/email build)

mkdir -p formance
(cd formance &&
  if [ ! -x ledger ]; then
    gh release download v2.4.18 -R formancehq/ledger -p 'ledger_darwin-arm64.tar.gz' -p 'ledger_checksums.txt' --clobber
    grep darwin-arm64.tar.gz ledger_checksums.txt | shasum -a 256 -c
    tar xzf ledger_darwin-arm64.tar.gz
  fi)

if [ ! -d formance-src ]; then
  git clone -q --depth 1 --branch v2.4.18 https://github.com/formancehq/ledger formance-src
fi
(cd formance-src && "$GO" build -o ../formance/ledger-src .)
for m in no-fetch split; do
  rm -rf "formance-mut-$m"
  cp -R formance-src "formance-mut-$m"
  (cd "formance-mut-$m" && patch -s -p1 < "$ROOT/targets/formance/mutants/$m.patch" && "$GO" build -o "../formance/ledger-$m" .)
done
echo "targets ready in $ROOT/.targets"
