# Variables
GOA = ./bin/goa
GO = go
MODULE=$(shell go list -m)

# Target to generate Goa code from the design
generate: clean
	@echo "Generating Goa service code..."
	$(GOA) gen $(MODULE)/design
	$(GOA) example $(MODULE)/design

# Target to clean up generated and compiled files
clean:
	@echo "Cleaning up generated and build files..."
	rm -rf ./gen todo.go ./cmd

# Format the Go code using go fmt
fmt:
	@echo "Formatting Go code..."
	$(GO) fmt ./...

# Check for linting errors
lint:
	@echo "Linting Go code..."
	$(GO) vet ./...

# Target to run the application in development (without rebuilding)
dev:
	@echo "Running the application in development mode..."
	$(GO) run ./cmd/todo

# Help command to show available targets
help:
	@echo "Available make targets:"
	@echo "  make generate   - Generate Goa service code from design"
	@echo "  make run        - Run the application"
	@echo "  make clean      - Clean generated and build files"
	@echo "  make fmt        - Format Go code"
	@echo "  make lint       - Lint Go code"
	@echo "  make dev        - Run the app in development mode"
	@echo "  make help       - Show this help message"

.PHONY: generate run clean fmt lint dev help