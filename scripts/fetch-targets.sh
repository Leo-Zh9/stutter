#!/usr/bin/env bash
# Sets up the real targets in .targets/ (gitignored), at pinned versions,
# without changing their code:
#   nextjs/saas-starter      commit 6e33e58 (pnpm, compile-only next build)
#   boxyhq/saas-starter-kit  commit abc9b68 (npm, prisma generate, compile-only next build)
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
