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
	docker build -t shippy-service-user .

build: ## Generate protobuf code and then build the container
	protoc --proto_path=. --go_out=. --micro_out=. \
		--go_opt=paths=source_relative \
		proto/user/user.proto
	docker build -t shippy-service-user .

run: ## Run the container
	docker run -p 50051 -e MICRO_SERVER_ADDRESS=:50051 \
		-e DB_HOST=${DB_HOST} \
		-e DB_USER=${DB_USER} \
		-e DB_PASSWORD=${DB_PASSWORD} \
		-e DB_NAME=${DB_NAME} \
		shippy-service-user
