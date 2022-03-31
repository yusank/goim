SHELL:=/usr/bin/env bash

.DEFAULT_GOAL:=help

Srv ?= push
BinPath ?= bin/$(Srv)
CmdPath ?= apps/$(Srv)/cmd/main.go
CfgPath ?= apps/$(Srv)/config
ProtoFile ?= api/config/v1/config.proto
IMAGE ?= goim/$(Srv)
VERSION ?= $(shell git describe --exact-match --tags 2> /dev/null || git rev-parse --abbrev-ref HEAD)

## env
export ROCKETMQ_GO_LOG_LEVEL=warn

##################################################
# Development                                    #
##################################################

##@ Development

.PHONY: vet
vet: ## Run go vet against code.
	go vet -v ./...

.PHONY: lint
lint: ## Run go lint against code.
	golangci-lint run ./... -v

.PHONEY: test
test: ## Run test against code.
	go test -v ./...

##################################################
# Generate                                          #
##################################################

##@ Generate

.PHONY: protoc
protoc: ## Run protoc command to generate pb code.
	protoc -I. --go_out=. --go-grpc_out=. $(ProtoFile)

##################################################
# Build                                          #
##################################################

##@ Build

.PHONY: build
build: ## build provided server
	go build -o $(BinPath) $(CmdPath)

.PHONY: build-all
build-all: ## build all apps
	make build Srv=push
	make build Srv=gateway
	make build Srv=msg
##################################################
# Docker                                         #
##################################################

##@ Docker

.PHONY: docker-build
docker-build: ## build docker image
	## build binary for docker image
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BinPath) $(CmdPath)
	docker build -f build/$(Srv).Dockerfile -t $(IMAGE):$(VERSION) .

##################################################
# Run                                            #
##################################################

##@ Run

.PHONY: run
run: build ## run provided server
	./$(BinPath) --conf $(CfgPath)

##################################################
# General                                        #
##################################################

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
