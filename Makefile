dep: ## Install dependencies
	@echo "Installing dependencies..."
	go mod download && go mod tidy

build: ## Build the binary file
	@echo "Building..."
	go build -o badgy

install: ## Install the binary file
	@echo "Installing..."
	go install
# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
