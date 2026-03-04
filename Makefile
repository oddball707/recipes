export TEST_CONTAINER_NAME=test-pg-docker
export TEST_CONTAINER_REPOSITORY=postgres
export TEST_PG_HOST=localhost
export TEST_PG_PORT=6542
export TEST_PG_USER=postgres
export TEST_PG_PASS=docker
export TEST_PG_NAME=recipes


GOFILES = $(shell find . -type f -name '*.go')

LOCAL_BIN:=$(CURDIR)/bin
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
SWAG_BIN:=$(LOCAL_BIN)/swag
GOOSE_BIN:=$(LOCAL_BIN)/goose
GOTESTSUM_BIN:=$(LOCAL_BIN)/gotestsum

install-lint:
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	$(info Downloading golangci-lint latest)
	GOFLAGS=-mod=mod GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.7.2
endif

install-swag:
ifeq ($(wildcard $(SWAG_BIN)),)
	$(info Downloading swag latest)
	GOBIN=$(LOCAL_BIN) go install github.com/swaggo/swag/cmd/swag@v1.16.2
endif

install-goose:
ifeq ($(wildcard $(GOOSE_BIN)),)
	$(info Downloading goose latest)
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest
endif

install-gotestsum:
ifeq ($(wildcard $(GOTESTSUM_BIN)),)
	$(info Downloading gotestsum latest)
	GOBIN=$(LOCAL_BIN) go install gotest.tools/gotestsum@latest
endif


fmt: # Format code
	$(info Formatting...)
	@gofmt -s -w ${GOFILES}

lint: install-lint # Run lint
	$(info Running lint...)
	$(GOLANGCI_BIN) run -v --config=.golangci.yaml ./... --fix

test: install-gotestsum # Run all tests (unit+DAL)
	$(info Cleaning test cache...)
	go clean -testcache
	$(info Running tests...)
	$(GOTESTSUM_BIN) ./... -tags=dal

build:
	GOPROXY=direct CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/main ./main.go

start: build start-db
	./bin/main

start-db: # Start only db dependency
	docker compose up db --wait -d

new-migration: install-goose # Create new sql migration. Usage: SQL={MIGRATION-NAME} make new-sql-migration
	$(info Creating new sql migration...)
	$(GOOSE_BIN) -dir migrations/sql create $(SQL) sql

new-seed: install-goose # Create new sql migration. Usage: SQL={MIGRATION-NAME} make new-seed-migration
	$(info Creating new seed migration...)
	$(GOOSE_BIN) -dir migrations/seeds create $(SQL) sql


migrations: install-goose start-db # Run postgres migrations
	$(GOOSE_BIN) -dir migrations/sql -allow-missing -table goose_db_version postgres "user=postgres host=localhost password=postgres dbname=postgres sslmode=disable" up

migrations-down: install-goose start-db # Rollback one migration
	$(GOOSE_BIN) -dir migrations/sql -allow-missing -table goose_db_version postgres "user=postgres host=localhost password=postgres dbname=postgres sslmode=disable" down

seeds: migrations # Run seeds
	$(info Running seeds...)
	$(GOOSE_BIN) -dir migrations/seeds -table goose_seed_version postgres "user=postgres host=localhost password=postgres dbname=postgres sslmode=disable" up


clean: clean-docker clean-binaries

clean-docker: # Stop and remove all containers and volumes,
	-docker compose down -v

clean-binaries: # Remove all binaries from bin and build folders
	rm -rf ${LOCAL_BIN}/*
	rm -rf $(CURDIR)/build/*
