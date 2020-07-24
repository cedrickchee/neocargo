# Simple Makefile to build, run, and tag a Docker container locally
# Inspired by: https://gist.github.com/mpneuried/0594963ad38e68917ef189b4e6a269db

# import config.
# You can change the default config with `make cnf="config_special.env" build`
cnf ?= config.env
include $(cnf)
export $(shell sed 's/=.*//' $(cnf))

build-proto: ## Generate protobuf code only
	protoc --proto_path=. --go_out=. --micro_out=. \
		--go_opt=paths=source_relative \
		proto/user/user.proto

build-image: ## Build the container only
	docker build -t neocargo-service-user .

build: ## Generate protobuf code and then build the container
	protoc --proto_path=. --go_out=. --micro_out=. \
		--go_opt=paths=source_relative \
		proto/user/user.proto
	docker build -t neocargo-service-user .

run: ## Run the container
	docker run -p 50051 -e MICRO_SERVER_ADDRESS=:50051 \
		-e POSTGRES_HOST=${POSTGRES_HOST} \
		-e POSTGRES_USER=${POSTGRES_USER} \
		-e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
		-e POSTGRES_DBNAME=${POSTGRES_DBNAME} \
		neocargo-service-user

test: ## Run test suite
	go test -v -timeout 30s github.com/haxorbit/neocargo/neocargo-service-user