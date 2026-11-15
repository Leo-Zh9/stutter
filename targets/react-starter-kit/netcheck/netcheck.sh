#!/bin/bash
# Network check: fresh DB, the config's migrate and seed, the config's start
# command plus a logging preload, then the frozen workload delivered once.
# Polls lsof over the server's whole process tree for non-loopback sockets.
set -u
SP=/private/tmp/claude-501/-Users-leozhang/efb647dd-3381-4a0b-8ea6-6ab29c3e5190/scratchpad/rsk
R=/Users/leozhang/stutter
T=$R/.targets/react-starter-kit
PSQL="$(cat $R/.pgbin)/psql -h $R/.pgsock -p 54329 -U postgres -q"
DB=rsk_netcheck2; PORT=18795
OUT=$SP/netcheck; rm -rf $OUT; mkdir -p $OUT
$PSQL -c "drop database if exists $DB" -c "create database $DB"
URL="postgres://postgres@127.0.0.1:54329/$DB?sslmode=disable"
(cd $T/db && DATABASE_URL=$URL npx --offline -y bun@1.4.2 --bun node_modules/.bin/drizzle-kit migrate) > $OUT/migrate.log 2>&1
$PSQL -d $DB -f $R/targets/react-starter-kit/fixtures.sql
cd $T/apps/api
env NETWATCH_LOG=$OUT/netwatch.log STUTTER_DB=$URL PORT=$PORT APP_NAME=Example APP_ORIGIN=http://127.0.0.1:$PORT ENVIRONMENT=development \
  BETTER_AUTH_SECRET=stutter_local_auth_secret_0123456789 RESEND_API_KEY=re_offline_never_used RESEND_EMAIL_FROM=onboarding@example.test \
  STRIPE_SECRET_KEY=sk_test_offline_never_used STRIPE_WEBHOOK_SECRET=whsec_stutter_local STRIPE_STARTER_PRICE_ID=price_prod_basic \
  STRIPE_PRO_PRICE_ID=price_prod_plus WRANGLER_SEND_METRICS=false \
  sh -c 'U="postgres://postgres:postgres@${STUTTER_DB#postgres://postgres@}"; export CLOUDFLARE_HYPERDRIVE_LOCAL_CONNECTION_STRING_HYPERDRIVE_CACHED="$U" CLOUDFLARE_HYPERDRIVE_LOCAL_CONNECTION_STRING_HYPERDRIVE_UNCACHED="$U" DATABASE_URL="$U"; exec npx --offline -y bun@1.4.2 run --preload '$SP'/netwatch-bun.cjs ./dev.ts' > $OUT/server.log 2>&1 &
SPID=$!
tree() { local p; for p in "$@"; do echo $p; tree $(pgrep -P $p); done; }
( while kill -0 $SPID 2>/dev/null; do P=$(tree $SPID | paste -sd, -); [ -n "$P" ] && lsof -nP -i -a -p $P 2>/dev/null | awk 'NR>1' >> $OUT/lsof-all.log; ps -o pid,ppid,pgid,comm -p $P >> $OUT/ps.log 2>/dev/null; sleep 0.2; done ) &
MON=$!
for i in $(seq 1 60); do curl -s -o /dev/null -m 2 http://127.0.0.1:$PORT/health && break; sleep 0.5; done
node $R/targets/stripe/lifecycles.mjs --sdk $T/apps/api/node_modules/stripe --url http://127.0.0.1:$PORT/api/auth/stripe/webhook --secret whsec_stutter_local --api basil > $OUT/workload.log 2>&1
sleep 2
ALL=$(tree $SPID); kill $ALL 2>/dev/null
sleep 1; kill $MON 2>/dev/null
$PSQL -d $DB -c 'select plan, reference_id, stripe_customer_id, stripe_subscription_id, status, cancel_at_period_end from subscription order by stripe_subscription_id' > $OUT/subscriptions.txt
grep -E '^-->' $OUT/server.log | sed 's/\x1b\[[0-9;]*m//g' | awk '{print $2, $3, $4}' | sort | uniq -c > $OUT/requests.txt
awk '{print $9, $10, $11}' $OUT/lsof-all.log | sort | uniq -c > $OUT/lsof-summary.txt
cut -d' ' -f3- $OUT/netwatch.log | sort | uniq -c > $OUT/netwatch-summary.txt
echo done
