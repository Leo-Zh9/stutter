GO ?= $(shell command -v go 2>/dev/null || echo $(HOME)/sdk/go/bin/go)
export GO

.PHONY: demo pg build trace test bench explore hunt stop clean miniagent agent-bench

# One command: start Postgres, build, run the full measurement suite.
demo: pg build trace bench

pg:
	@./scripts/bootstrap.sh

build:
	$(GO) build -o bin/stutter ./cmd/stutter
	$(GO) build -o bin/minipay ./targets/minipay
	$(GO) build -o bin/minipay-load ./targets/minipay/load

# The committed trace is reused; delete it to record a fresh one.
trace: traces/minipay.jsonl
traces/minipay.jsonl: | pg build
	@mkdir -p traces
	./bin/stutter record -config targets/minipay/stutter.json -out $@ -- ./bin/minipay-load -target {proxy}

bench: pg build trace
	./bin/stutter bench -config targets/minipay/stutter.json -out results

# Point stutter at one variant by hand, e.g. make explore BUG=key-after-commit
explore: pg build trace
	./bin/stutter explore -config targets/minipay/stutter.json -env BUG=$(BUG)

# Real targets (needs scripts/fetch-targets.sh first): about 70 minutes.
hunt: pg build
	./scripts/hunt-all.sh

# Agent mode: miniagent is a FastAPI + LangGraph + PostgresSaver app with one
# seeded bug per variant. agent-bench hunts every variant (about 15 min).
miniagent: targets/miniagent/.venv/bin/python
targets/miniagent/.venv/bin/python: targets/miniagent/requirements.txt
	uv venv -q --allow-existing --python 3.12 targets/miniagent/.venv
	uv pip install -q --python targets/miniagent/.venv/bin/python -r targets/miniagent/requirements.txt

agent-bench: pg build miniagent
	./scripts/agent-bench.sh

test: pg build
	STUTTER_REQUIRE_PG=1 $(GO) test ./... -count=1

stop:
	@$$(cat .pgbin)/pg_ctl -D .pgdata stop -m fast

clean: stop
	rm -rf bin .pgdata .pgsock .pgbin .pgport .venv .schema.sql
