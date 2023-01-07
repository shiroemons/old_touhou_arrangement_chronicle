.PHONY: help test generate lint
.DEFAULT_GOAL := help

test: ## Execute tests
	go test -race -shuffle=on ./...

generate: ## Run go generate ./...
	go generate ./...

lint:
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:latest golangci-lint run -v --timeout 5m

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
