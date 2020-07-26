##@ Infra

up-postgresql:  ## postgresql
	docker-compose -f ./infra/docker-compose-postgres.yaml up -d

client-postgres:  ## run pgcli client
	pgcli -h localhost -p 5432 -U postgres

.PHONY: up-postgresql client-postgresql

##@ Help

.PHONY: help

help:  ## Display this help
		@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z\_\-0-9]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

