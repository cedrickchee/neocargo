# import config.
# You can change the default config with `make cnf="config_special.env" build`
cnf ?= config.env
include $(cnf)
export $(shell sed 's/=.*//' $(cnf))

# HELP

# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

# DOCKER COMPOSE TASKS

# Build the containers
build: ## Build stack with Docker Compose
	docker-compose build --build-arg GITHUB_TOKEN=${GITHUB_TOKEN}

# Run the containers
run: ## Run stack with Docker Compose
	docker-compose up

run-cli: ## Run the CLI tool
	docker-compose run cli

# Stop the containers
stop: ## Teardown stack and stop all containers
	docker-compose down
