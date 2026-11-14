#!/usr/bin/env bash
# Runs stutter against every real target and every Formance mutant, one at a
# time (each target gets its own template database). Needs the targets set
# up first: see scripts/fetch-targets.sh. Exit status: 0 if every run did
# what it should, so a finding counts as success here.
set -uo pipefail
cd "$(dirname "$0")/.."
run() { # name, expected exit (0 clean, 1 findings), args...
  local name=$1 want=$2; shift 2
  echo "== $name"
  ./bin/stutter "$@" > "results/logs/$name.log" 2>&1
  local got=$?
  echo "   exit $got (expected $want)"
  [ "$got" = "$want" ] || FAILED=1
}
mkdir -p results/logs
FAILED=0
run saas-starter 1 hunt -config targets/saas-starter/stutter.json
run boxyhq 1 hunt -config targets/boxyhq/stutter.json
run formance 0 hunt -config targets/formance/stutter.json
run formance-random 0 explore -config targets/formance/stutter.json -random 200 -lanes 6
for m in source-build drop-index no-fetch no-fetch-drop-index split; do
  want=1; [ "$m" = source-build ] && want=0; [ "$m" = no-fetch ] && want=0
  run "formance-$m" "$want" hunt -config "targets/formance/mutants/$m.json" -out "targets/formance/mutants/hunt-$m"
done
exit $FAILED
