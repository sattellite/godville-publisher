BINARY     := godville-publisher
BUILDFLAGS := "-s -w"

.SUFFIXES:
.PHONY: build run lint test help
.DEFAULT_GOAL := help

### MAIN COMMANDS ###
build: lint test ## Build binary
	@go mod tidy -e
	@mkdir -p build
	go build -ldflags $(BUILDFLAGS) -o build/$(BINARY) $(MAINFILE)

run: tidy ## Run binary
	go run cmd/godville/main.go

### ADDITIONAL COMMANDS ###

lint: ## Run linter
	golangci-lint run

test: tidy ## Run tests
	go test -count 1 -race -v ./...

tidy: ## Run go mod tidy
	go mod tidy -e

# Auto documented Makefile https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
