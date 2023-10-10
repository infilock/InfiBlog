.PHONY: help
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.DEFAULT_GOAL := help

.PHONY: new
new: ## Create a new migration
	sql-migrate new

.PHONY: up
up: ## Migrates the database to the most recent version available
	sql-migrate up

.PHONY: down
down: ## Undo a database migration
	sql-migrate down