# docker compose
TARGET ?=
ENV_FILE ?= .env
COMPOSE_CMD = docker compose -f compose.yaml --env-file $(ENV_FILE)

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: mocks
mocks: mockery ## Generate mock implementations.
	$(MOCKERY) --config .mockery.yaml

##@ Docker Compose

.PHONY: compose-up
compose-up: ## Run components.
	$(COMPOSE_CMD) up -d $(TARGET)

.PHONY: compose-down
compose-down: ## Shutdown components.
	$(COMPOSE_CMD) down $(TARGET)

.PHONY: compose-ps
compose-ps: ## Print running components.
	$(COMPOSE_CMD) ps $(TARGET)

.PHONY: compose-logs
compose-logs: ## Tail logs of components.
	$(COMPOSE_CMD) logs -f $(TARGET)

##@ Tool Dependencies

## Location to install dependencies to.
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	@mkdir -p $(LOCALBIN)

## Tool Binaries
MOCKERY ?= $(LOCALBIN)/mockery

.PHONY: mockery
mockery: $(MOCKERY) ## Install mockery if necessary.
$(MOCKERY): $(LOCALBIN)
	@test -s $(MOCKERY) || \
	GOBIN=$(LOCALBIN) go install github.com/vektra/mockery/v2@latest
