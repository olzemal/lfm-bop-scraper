GO := $(shell which go)
GOLINT := $(shell $(GO) env GOBIN)/golangci-lint

.DEFAULT_GOAL: bop.json

generate:
	@$(GO) generate ./...

run:
	@$(GO) run cmd/main.go

lint: $(GOLINT)
	@$(GOLINT) run

test:
	@$(eval PROFILE := $(shell mktemp))
	@$(GO) test -v -coverpkg=./... -coverprofile=$(PROFILE) ./...
	@$(GO) tool cover -func $(PROFILE)

$(GOLINT):
	@$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.56.2

bop.json:
	@$(GO) run cmd/main.go -o bop.json

.PHONY: generate run test
