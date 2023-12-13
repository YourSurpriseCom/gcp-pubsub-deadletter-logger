.DEFAULT_GOAL := help

.PHONY: help
help:
	@printf "\033[33mUsage:\033[0m\n  make [target] [arg=\"val\"...]\n\n\033[33mTargets:\033[0m\n"
	@grep -E '^[-a-zA-Z0-9_\.\/]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[32m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: init
init: .env  install ## Initialise and install all dependencies

.PHONY: install 
install: ## install dependencies 
	@go mod tidy
	@go mod download

.PHONY: run
run: .env ## Start the application
	@bash -c "env `cat .env | xargs` go run ."

.env: | .env.dist
	@echo "Copying .env.dist to .env"
	@cp -nv .env.dist .env

.PHONY: check
check: golangci-lint go-vet unit-tests security-code-scan security-vulnerability-scan ## Run application checks
.PHONY: golangci-lint
golangci-lint:
	@golangci-lint run

.PHONY: go-vat
go-vet:
	@go vet

.PHONY: unit-tests
unit-tests:
	@go test ./...

.PHONY: security-code-scan
security-code-scan:
	@gosec ./...

.PHONY: security-vulnerability-scan
security-vulnerability-scan:
	@govulncheck ./...
