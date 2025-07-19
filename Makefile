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

# Run benchmarks
benchmark:
	@echo "Running Go benchmarks..."
	@go test -bench=. -benchmem ./benchmarks

# Run comprehensive benchmark suite
benchmark-full:
	@echo "Running comprehensive benchmark suite..."
	@chmod +x scripts/benchmark.sh
	@./scripts/benchmark.sh

# Run memory profile benchmark
benchmark-mem:
	@echo "Running memory profile benchmark..."
	@go test -bench=. -memprofile=mem.prof ./benchmarks
	@go tool pprof -http=:8081 mem.prof

# Run CPU profile benchmark  
benchmark-cpu:
	@echo "Running CPU profile benchmark..."
	@go test -bench=. -cpuprofile=cpu.prof ./benchmarks
	@go tool pprof -http=:8081 cpu.prof

# Load testing with hey (if available)
load-test:
	@echo "Running load test..."
	@if command -v hey >/dev/null 2>&1; then \
		echo "Starting application in background..."; \
		go run $(MAIN_PATH) & \
		APP_PID=$$!; \
		sleep 3; \
		echo "Running load test with hey..."; \
		hey -n 10000 -c 100 http://localhost:8080/health; \
		kill $$APP_PID; \
	else \
		echo "hey not found. Install with: go install github.com/rakyll/hey@latest"; \
	fi

# Performance analysis
perf-analysis: benchmark load-test
	@echo "Performance analysis complete"

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
	@echo "  benchmark     Run Go benchmarks"
	@echo "  benchmark-full Run comprehensive benchmark suite"
	@echo "  benchmark-mem Run memory profile benchmark"
	@echo "  benchmark-cpu Run CPU profile benchmark"
	@echo "  load-test     Run load test with hey"
	@echo "  perf-analysis Run complete performance analysis"
	@echo "  clean         Clean build artifacts"
	@echo "  deps          Install dependencies"
	@echo "  lint          Lint the code"
	@echo "  fmt           Format the code"
	@echo "  mocks         Generate mocks"
	@echo "  migrate-up    Run database migrations"
	@echo "  migrate-down  Rollback database migrations"
	@echo "  help          Show this help message"
