.PHONY: build run test clean help

# Variables
APP_NAME=golang-domain-driven-design
BUILD_DIR=build
MAIN_PATH=cmd/api/main.go

# Build the application
build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)

# Run the application
run:
	@echo "Running $(APP_NAME)..."
	@go run $(MAIN_PATH)

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

# Lint the code
lint:
	@echo "Linting code..."
	@golangci-lint run

# Format the code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Generate mocks
mocks:
	@echo "Generating mocks..."
	@mockgen -source=internal/domain/repositories/user_repository.go -destination=mocks/user_repository_mock.go

# Database migration up
migrate-up:
	@echo "Running database migrations..."
	@migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" up

# Database migration down
migrate-down:
	@echo "Rolling back database migrations..."
	@migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" down

# Help
help:
	@echo "Available commands:"
	@echo "  build         Build the application"
	@echo "  run           Run the application"
	@echo "  test          Run tests"
	@echo "  test-coverage Run tests with coverage"
	@echo "  clean         Clean build artifacts"
	@echo "  deps          Install dependencies"
	@echo "  lint          Lint the code"
	@echo "  fmt           Format the code"
	@echo "  mocks         Generate mocks"
	@echo "  migrate-up    Run database migrations"
	@echo "  migrate-down  Rollback database migrations"
	@echo "  help          Show this help message"
