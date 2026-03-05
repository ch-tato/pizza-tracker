NAME=pizza-tracker
BIN_DIR=bin
CMD_DIR=./cmd

.PHONY: all build run clean test tidy run-dev

all: build

build:
	@echo "Building application..."
	@go build -o $(BIN_DIR)/$(NAME) $(CMD_DIR)

run: build
	@echo "Running application..."
	@./$(BIN_DIR)/$(NAME)

run-dev:
	@echo "Running in development mode..."
	@go run $(CMD_DIR)

clean:
	@echo "Cleaning binaries..."
	@rm -rf $(BIN_DIR)

test:
	@echo "Running tests..."
	@go test -v ./...

tidy:
	@echo "Tidying module dependencies..."
	@go mod tidy
	@go mod verify
