GO := $(shell which go)

.DEFAULT_GOAL: bop.json

generate:
	@$(GO) generate ./...

run:
	@$(GO) run cmd/main.go

bop.json:
	@$(GO) run cmd/main.go -o bop.json

.PHONY: generate run
