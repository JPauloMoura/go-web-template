include .env
export
 
run: build
	@echo "Start..."
	@./bin/main

docker-run: generation-image
	@echo "Run app in container..."
	@docker run -it -p 3002:3002 jp0user0docker/validator-uuid-api:latest

build:
	@echo "Build project..."
	@go mod tidy
	@go build -o ./bin/main 

generation-image: build
	@echo "Generation docker image..."
	@docker build --no-cache -f Dockerfile -t jp0user0docker/validator-uuid-api:latest .