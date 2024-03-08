GO := $(shell which go)

.DEFAULT_GOAL: bop.json

generate:
	@$(GO) generate ./...

run:
	@$(GO) run cmd/main.go

test:
	@$(eval PROFILE := $(shell mktemp))
	@$(GO) test -v -coverpkg=./... -coverprofile=$(PROFILE) ./...
	@$(GO) tool cover -func $(PROFILE)

bop.json:
	@$(GO) run cmd/main.go -o bop.json

.PHONY: generate run test
