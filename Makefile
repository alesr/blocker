.DEFAULT_GOAL := help

NAME = $(shell basename $(PWD))
AWAIT_DB_SCRIPT = wait-for-it.sh

.PHONY: help
help:
	@echo "------------------------------------------------------------------------"
	@echo "${NAME}"
	@echo "------------------------------------------------------------------------"
	@grep -E '^[a-zA-Z0-9_/%\-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Build Go binary to ./bin
	@echo "Building ${NAME}..."
	@go build -o bin/${NAME}

.PHONY: run
run: build ## Build and run Go binary
	@echo "Running ${NAME}..."
	@./bin/${NAME}

.PHONY: test
test: ## Run tests
	@echo "Running tests..."
	@go test -v -race -count=1 ./...
