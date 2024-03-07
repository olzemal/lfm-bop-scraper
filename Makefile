GO := $(shell which go)

generate:
	@$(GO) generate ./...

build:
	@$(GO) build -o lfmbopscraper cmd/main.go

run:
	@$(GO) run cmd/main.go

.PHONY: generate build run
