# Makefile for HRM Backend

.PHONY: swag-init swag-fmt run build clean

# Generate swagger documentation
swag-init:
	swag init -g cmd/main.go -o docs/

# Format swagger comments
swag-fmt:
	swag fmt

# Run the application
run:
	go run cmd/main.go

# Build the application
build:
	go build -o bin/hrm-backend cmd/main.go

# Clean build artifacts
clean:
	rm -rf bin/
	rm -rf docs/

# Install swag if not installed
install-swag:
	go install github.com/swaggo/swag/cmd/swag@latest

# Generate docs and run
dev: swag-init run
