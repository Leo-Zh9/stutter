GO ?= $(shell command -v go 2>/dev/null || echo $(HOME)/sdk/go/bin/go)
export GO

.PHONY: demo pg build trace test bench explore hunt stop clean

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

# Real targets (needs scripts/fetch-targets.sh first): about 30 minutes.
hunt: pg build
	./scripts/hunt-all.sh

test: pg build
	STUTTER_REQUIRE_PG=1 $(GO) test ./... -count=1

stop:
	@$$(cat .pgbin)/pg_ctl -D .pgdata stop -m fast

clean: stop
	rm -rf bin .pgdata .pgsock .pgbin .pgport .venv .schema.sql
