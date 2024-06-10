PROJECT ?= farmacia-tech-hub


DOCKER_COMPOSE_FILE_BUILD=build/docker-compose.yml

DOCKER_COMPOSE_FILE_LOCAL=docker-compose.yml

GOLANGCI_LINT_PATH=$$(go env GOPATH)/bin/golangci-lint
GOLANGCI_LINT_VERSION=1.59.0

build:
	docker compose -f $(DOCKER_COMPOSE_FILE_BUILD) -p $(PROJECT) down --remove-orphans
	docker compose -f $(DOCKER_COMPOSE_FILE_BUILD) -p $(PROJECT) up --remove-orphans

start:
	docker compose -f $(DOCKER_COMPOSE_FILE_LOCAL) -p $(PROJECT) down --remove-orphans
	docker compose -f $(DOCKER_COMPOSE_FILE_LOCAL) -p $(PROJECT) up --remove-orphans

mock:

lint:
	@echo "==> Installing golangci-lint"
ifeq (,$(findstring $(GOLANGCI_LINT_VERSION),$(shell which $(GOLANGCI_LINT_PATH) && eval $(GOLANGCI_LINT_PATH) version)))
	@echo "installing golangci-lint v$(GOLANGCI_LINT_VERSION)"
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v$(GOLANGCI_LINT_VERSION)
else
	@echo "already installed: $(shell eval $(GOLANGCI_LINT_PATH) version)"
endif
	@echo "==> Running golangci-lint"
	@$(GOLANGCI_LINT_PATH) run -c ./.golangci.yml --fix

encrypt:
	go run ./tools/encrypt.go --word=$(filter-out $@,$(MAKECMDGOALS))
