.DEFAULT_GOAL := help

_YELLOW=\033[0;33m
_NC=\033[0m

.PHONY: help setup # generic commands
help: ## prints this help
	@grep -hE '^[\.a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "${_YELLOW}%-16s${_NC} %s\n", $$1, $$2}'

setup: ## downloads dependencies
	go get -u golang.org/x/lint/golint
	go mod tidy

.PHONY: lint build # go commands
lint: ## runs the code linter
	golint ./...

build: ## builds the go binary
	go build -o stub ./cmd/stub

