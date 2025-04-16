# Variables
APP_NAME := paper-trail
API_SERVER := cmd/api-server/main.go
CLI := cmd/cli/main.go

# Default target
.PHONY: all
all: build

# Build the API server
.PHONY: build-api
build-api:
	@echo "Building API server..."
	go build -o bin/$(APP_NAME)-api-server $(API_SERVER)

# Build the CLI
.PHONY: build-cli
build-cli:
	@echo "Building CLI..."
	go build -o bin/$(APP_NAME)-cli $(CLI)

# Build all components
.PHONY: build
build: build-api build-cli

# Run the API server
.PHONY: run-api
run-api:
	@echo "Running API server..."
	go run $(API_SERVER)

# Run the CLI
.PHONY: run-cli
run-cli:
	@echo "Running CLI..."
	go run $(CLI)

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	go test ./...