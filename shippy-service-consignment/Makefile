# Simple Makefile to build, run, and tag a Docker container locally
# Inspired by: https://gist.github.com/mpneuried/0594963ad38e68917ef189b4e6a269db

# import config.
# You can change the default config with `make cnf="config_special.env" build`
cnf ?= config.env
include $(cnf)
export $(shell sed 's/=.*//' $(cnf))

build: ## Build proto file and then build the container
	protoc --proto_path=. --go_out=. --micro_out=. \
		--go_opt=paths=source_relative \
		proto/consignment/consignment.proto
	docker build --build-arg GITHUB_TOKEN=${GITHUB_TOKEN} -t shippy-service-consignment .

run: ## Run the container
	docker run -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e DB_HOST=mongodb://datastore:27017 \
		shippy-service-consignment
