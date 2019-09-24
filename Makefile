SHELL := /bin/bash

all: test build

test:
	@echo " === Running: $@ === "
	@go vet ./...
	@go test -v -race -cover ./...

build:
	@echo " === Running: $@ === "
	@go build ./cmd/...
	@echo
