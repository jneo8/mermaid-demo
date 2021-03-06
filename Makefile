GO ?= go
GOFMT ?= gofmt "-s"
PACKAGES ?= $(shell $(GO) list ./...)
VETPACKAGES ?= $(shell $(GO) list ./... | grep -v /examples/)
GOFILES := $(shell find . -name "*.go")
TESTFOLDER := $(shell $(GO) list ./... | grep -E 'mermaid-demo$$|cmd' | grep -v examples)
TESTTAGS ?= ""
PROJ = mermaid-demo

##@ Run

run-with-mermaid-example:  ## Run with mermaid example
	go run ./cmd/with-mermaid/main.go

run-without-mermaid-example:  ## Run without mermaid example
	go run ./cmd/without-mermaid/main.go

.PHONY: run-with-mermaid-example run-without-mermaid-example

##@ Test

install-richgo:  ## Install richgo
	go get -u github.com/kyoh86/richgo

test:  ## Run test
	echo "mode: count" > coverage.out
	for d in $(TESTFOLDER); do \
		$(GO) test -tags $(TESTTAGS) -v -covermode=count -coverprofile=profile.out $$d | richgo testfilter > tmp.out; \
		cat tmp.out; \
		if grep -q "^--- FAIL" tmp.out; then \
			rm tmp.out; \
			exit 1; \
		elif grep -q "build failed" tmp.out; then \
			rm tmp.out; \
			exit 1; \
		elif grep -q "setup failed" tmp.out; then \
			rm tmp.out; \
			exit 1; \
		fi; \
		if [ -f profile.out ]; then \
			cat profile.out | grep -v "mode:" >> coverage.out; \
			rm profile.out; \
		fi; \
	done

.PHONY: test install-richgo

##@ Infra

up-postgresql:  ## postgresql
	docker-compose -f ./infra/docker-compose-postgres.yaml up -d

client-postgres:  ## run pgcli client
	pgcli -h localhost -p 5432 -U postgres

create-db:  ## Create database in postgres
	docker exec -it mermaid-db psql -c 'create database mermaid_demo' -U postgres

.PHONY: up-postgresql client-postgres create-db

##@ lint

install-lint:  ## Install golangci-lint to ./bin
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.27.0

linter-run:  ## Run linter for all
		./bin/golangci-lint run ./...

.PHONY: linter-run install-lint create-db

##@ Help

help:  ## Display this help
		@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: help

.DEFAULT_GOAL := help

