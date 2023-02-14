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
	@go build -o bin/${NAME}

.PHONY: run
run: build ## Build and run Go binary
	@./bin/${NAME}

.PHONY: test
test: ## Run tests
	@go test -v -race -count=1 ./...

.PHONY: proto
proto: ## Generate gRPC code from proto files
	@protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/*.proto
