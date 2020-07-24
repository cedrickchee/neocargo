build: ## Build proto file and then build the container
	protoc --proto_path=. --go_out=. --micro_out=. \
		--go_opt=paths=source_relative \
		proto/vessel/vessel.proto
	docker build -t neocargo-service-vessel .

run: ## Run the container
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 \
		-e DB_HOST=mongodb://datastore:27017 \
		neocargo-service-vessel