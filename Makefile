.PHONY: help init build-local db-up up down logs ps migrate all-clean
.DEFAULT_GOAL := help

init: ## Initialize environment
	docker compose build

build-local: ## Build docker image to local development
	docker compose build --no-cache

up: ## Do docker compose up with hot reload
	docker compose up -d

db-up: ## Run docker compose up db
	docker compose up db -d

down: ## Do docker compose down
	docker compose down

logs: ## Tail docker compose logs
	docker compose logs -f

ps: ## Check container status
	docker compose ps

migrate: ## db migrate
	docker compose run --rm migrate

all-clean:
	docker compose down --rmi all --volumes --remove-orphans

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
