PROJECT ?= farmacia-tech-hub


DOCKER_COMPOSE_FILE_BUILD=build/docker-compose.yml

DOCKER_COMPOSE_FILE_LOCAL=docker-compose.yml

build:
	docker compose -f $(DOCKER_COMPOSE_FILE_BUILD) -p $(PROJECT) down --remove-orphans
	docker compose -f $(DOCKER_COMPOSE_FILE_BUILD) -p $(PROJECT) up --remove-orphans

start:
	docker compose -f $(DOCKER_COMPOSE_FILE_LOCAL) -p $(PROJECT) down --remove-orphans
	docker compose -f $(DOCKER_COMPOSE_FILE_LOCAL) -p $(PROJECT) up --remove-orphans

mock:

lint:

encrypt:
	go run ./tools/encrypt.go --word=$(filter-out $@,$(MAKECMDGOALS))
