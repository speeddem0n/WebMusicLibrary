.PHONY: build

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build web music app.
	docker compose build

test: ## Launch all unit tests.
	go test -v ./...

run: ## Start web music app and postgres.
	docker compose up -d

install: test build run ## Test, build and start all services.

down: ## Stop all services.
	docker compose down

restart: down run ## Restart all services.

fmt: ## Run go formatter on all project's files
	gofmt -s -w .