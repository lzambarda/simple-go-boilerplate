help: ## Show help messages.
	@grep -E '^[0-9a-zA-Z_-]+:(.*?## .*)?$$' $(MAKEFILE_LIST) | sed 's/^Makefile://' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

name=BINARYNAME
main=./cmd/*.go

timeout=10m
dir="./..."
run=.
flags=

# Export any env variable which may be needed.
# This automatically picks up the override, if provided.
#
# Try to run both "go run cmd/main.go" and "make run" and compare the outputs.
ifneq ($(shell find ./config -name "dev.env"),) 
include ./config/dev.env
export $(shell sed 's/=.*//' ./config/dev.env)
else
include ./config/dev.env.example
export $(shell sed 's/=.*//' ./config/dev.env.example)
endif


.PHONY: dependencies
dependencies: ## Install dependencies needed to work on this repo
	@go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2


.PHONY: run
run: ## Example: make run env={env file} cmd={cmd} flags={flags}
	@go run $(main) $(cmd) $(flags)


# This is a great place to use e.g. golangci
# This can be quite taxing, you don't need to have any linting set up.
.PHONY: lint
lint: ## Run various linters
	golangci-lint run ./...


.PHONY: build
build: ## Builds the binary
	@CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags="-w -s" -o ./bin/$(name) $(main)


.PHONY: unit_test
unit_test: ## Run unit tests. You can set: [run, timeout, short, dir, flags]. Example: make unit_test flags="-race".
	@go mod tidy; go test -trimpath -failfast --timeout=$(timeout) $(short) $(dir) -run $(run) $(flags)


.PHONY: integration_test
integration_test: ## Run integration. You can set: [run, timeout, short, dir, flags]. Example: make integration_test flags="-race".
	@go mod tidy; go test -trimpath -failfast --timeout=$(timeout) -tags=integration $(short) $(dir) -run $(run) $(flags)


.PHONY: setup_test_dependencies
setup_test_dependencies: ## Set up dependencies needed by the integration test, such as a database.


.PHONY: reset_test_dependencies
reset_test_dependencies: ## Tear down dependencies created by setup_test_dependencies


.PHONY: mocks
mocks: ## Generate mocks in all packages.
	@go generate ./...
