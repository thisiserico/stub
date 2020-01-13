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


.PHONY: pack push # docker commands
pack: ## builds the docker image
	docker build -f ./cmd/stub/Dockerfile -t stub .

push: ## pushes the docker image to the registry
	docker tag stub docker.pkg.github.com/thisiserico/stub:1.0.0
	docker push docker.pkg.github.com/thisiserico/stub:1.0.0

