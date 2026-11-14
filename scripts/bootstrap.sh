#!/usr/bin/env bash
# Starts a project-local PostgreSQL 16 cluster on a unix socket in .pgsock.
# Binaries come from the pgserver Python package, installed into .venv with uv,
# so nothing is installed system-wide and no Docker is needed.
set -euo pipefail
cd "$(dirname "$0")/.."
ROOT="$PWD"

GO="${GO:-$(command -v go || echo "$HOME/sdk/go/bin/go")}"
if [ ! -x "$GO" ]; then
  echo "Go not found. Install it (https://go.dev/dl/) or set GO=/path/to/go." >&2
  exit 1
fi

if ! command -v uv >/dev/null 2>&1; then
  echo "uv not found. Install it: https://docs.astral.sh/uv/" >&2
  exit 1
fi

if [ ! -x .venv/bin/python ]; then
  uv venv -q --python 3.12 .venv
fi
if ! .venv/bin/python -c "import pgserver" 2>/dev/null; then
  uv pip install -q --python .venv/bin/python pgserver==0.1.4
fi
PGBIN="$(.venv/bin/python -c 'import pgserver,os;print(os.path.join(os.path.dirname(pgserver.__file__),"pginstall","bin"))')"
echo "$PGBIN" > .pgbin
# pgserver ships no contrib modules. Install stutter's SQL-only pgcrypto
# digest() shim (scripts/pgcrypto-shim) so targets that need it can migrate.
EXTDIR="$PGBIN/../share/postgresql/extension"
cp scripts/pgcrypto-shim/pgcrypto.control scripts/pgcrypto-shim/pgcrypto--1.3.sql "$EXTDIR/"

mkdir -p .pgsock
PGPORT="${STUTTER_PGPORT:-54329}"
echo "$PGPORT" > .pgport
if [ ! -f .pgdata/PG_VERSION ]; then
  "$PGBIN/initdb" -D .pgdata -U postgres --auth=trust --encoding=UTF8 --locale=C >/dev/null
fi

if ! "$PGBIN/pg_ctl" -D .pgdata status >/dev/null 2>&1; then
  # fsync, synchronous_commit and full_page_writes only affect durability after
  # an OS crash. They do not change isolation levels or lock behaviour, which is
  # all stutter's faults depend on.
  "$PGBIN/pg_ctl" -D .pgdata -l .pgdata/server.log -w start -o \
    "-k $ROOT/.pgsock -p $PGPORT -c listen_addresses=127.0.0.1 -c max_connections=300 \
     -c fsync=off -c synchronous_commit=off -c full_page_writes=off" >/dev/null
fi
echo "postgres ready on $ROOT/.pgsock and 127.0.0.1:$PGPORT"
