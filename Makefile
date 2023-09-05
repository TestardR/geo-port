NAME=geo-port

.PHONY: deps fmt test clean vet docker-build

deps:
	@echo "Installing dependencies ..."
	@go mod tidy && go mod download
	@echo "Installing dependencies, done!"

fmt:
	@echo "Formatting code ..."
	@go fmt ./...
	@echo "Formatting code, done!"

test:
	@echo "Running tests ..."
	@ginkgo -r -cover ./internal
	@echo "Running tests, done!"

build: deps clean
	@echo "Building ...."
	@mkdir -p ./build
	@go build -o ./build/$(NAME)
	@echo "Build done"

docker-build: deps clean
	@echo "Building Docker image ..."
	@env GOOS=linux CGO_ENABLED=0 go build -o ./bin/$(NAME)
	@docker build -t $(NAME) .
	@rm -rf ./bin
	@echo "Building Docker image, done!"

clean:
	@echo "Cleaning ..."
	@go clean && rm -rf ./build
	@echo "Cleaning done"

vet:
	@echo "Running vet ..."
	@go vet ./...
	@echo "Running vet, done!"
