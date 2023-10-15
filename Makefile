LOGFILE=$(LOGPATH) `date +'%A-%b-%d-%Y'`
branch := $(shell git branch --show-current)
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export GO111MODULE=on

.PHONY: help
help: ## Shows help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.DEFAULT_GOAL := help

.PHONY: api
api: ## run api
	CGO_ENABLED=1 go build -v -o $(ROOT)/bin/api -ldflags="-s -w" $(ROOT)/cmd/api/*.go

.PHONY: build-job-gpt
build-job-gpt: ## run job crawler from chat gpt
	CGO_ENABLED=1 go build -v -o $(ROOT)/bin/chatgpt -ldflags="-s -w" $(ROOT)/cmd/job/chatgpt/*.go

.PHONY: build-job-wp
build-job-wp: ## run job send article to blog post
	CGO_ENABLED=1 go build -v -o $(ROOT)/bin/wordpress -ldflags="-s -w" $(ROOT)/cmd/job/wordpress/*.go

.PHONY: build-job-sm
build-job-sm: ## run job send post to social media
	CGO_ENABLED=1 go build -v -o $(ROOT)/bin/socialmedia -ldflags="-s -w" $(ROOT)/cmd/job/socialmedia/*.go

.which-go:
	@which go > /dev/null || (echo "install go from https://golang.org/dl/" & exit 1)

.which-lint:
	@which golangci-lint > /dev/null || (echo "install golangci-lint from https://github.com/golangci/golangci-lint" & exit 1)

lint: .which-lint
	golangci-lint run -v

clean: # run make format and make lint
	fieldalignment -fix ./...
	gosec ./...
	golangci-lint run -v

.PHONY: alignment
alignment: ## Analyzer that detects structs
	fieldalignment -fix ./...

.PHONY: gosec
gosec: ## check security golang code
	gosec ./...

.PHONY: run
run: ## run application
	godotenv -f .env go run ./cmd/api/*.go