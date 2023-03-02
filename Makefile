##@ Test

test: ## Run tests.
ifeq (, $(shell which gotest))
	$(shell go install github.com/rakyll/gotest)
endif
	GO_ENV=test godotenv -f ${PWD}/.env.test gotest `${NON_GENERATED_PACKAGES}`

test-coverage: ## Run tests and build coverage.
	GO_ENV=test godotenv -f ${PWD}/.env.test go test -coverprofile=c.out ./...

coverage: test-coverage ## Build coverage.
	GO_ENV=test godotenv -f ${PWD}/.env.test go tool cover -html=c.out -o coverage.html

##@ Development

lint: check-golangci-lint ## Lints code.
	golangci-lint run

setup: check-golangci-lint husky ## Setup environment.
	go get github.com/princjef/gomarkdoc/cmd/gomarkdoc

husky: ## Initialize husky hooks.
	@type husky >/dev/null 2>&1 || ( \
		echo -n "husky not found. Installing... "; \
		go install github.com/automation-co/husky@latest; \
		echo " Done." \
		)
	husky install

check-golangci-lint: ## Check that golangci-lint is installed, and installs it otherwise.
	@type golangci-lint >/dev/null 2>&1 || ( \
		echo -n "golangci-lint not found. Installing... "; \
		go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.1; \
		echo " Done." \
		)

help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[ a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

clean: ## Cleanup output and coverage files.
	rm coverage.html
