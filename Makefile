NAME=geo-port

.PHONY: deps fmt test clean vet docker-build lint

deps:
	@echo "Installing dependencies ..."
	@go mod tidy && go mod download
	@go install github.com/onsi/ginkgo/v2/ginkgo
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
	@docker build -t $(NAME) .
	@echo "Building Docker image, done!"

clean:
	@echo "Cleaning ..."
	@go clean && rm -rf ./build
	@echo "Cleaning done"

vet:
	@echo "Running vet ..."
	@go vet ./...
	@echo "Running vet, done!"

lint:
	@golangci-lint run --tests=false