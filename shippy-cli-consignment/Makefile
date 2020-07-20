# Simple Makefile to build, run, and tag a Docker container locally
# Inspired by: https://gist.github.com/mpneuried/0594963ad38e68917ef189b4e6a269db

# import config.
# You can change the default config with `make cnf="config_special.env" build`
cnf ?= config.env
include $(cnf)
export $(shell sed 's/=.*//' $(cnf))

build: ## Build the container
	docker build --build-arg GITHUB_TOKEN=${GITHUB_TOKEN} -t shippy-cli-consignment .

run: ## Run the container
	docker run shippy-cli-consignment