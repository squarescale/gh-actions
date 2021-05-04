help: ## This help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

coverage: ## Run test suites in all packages with code coverage
	go test ./... -cover -coverprofile=coverage.out

coverage_html: coverage ## Show code coverage html report
	go tool cover -html=coverage.out -o coverage.html