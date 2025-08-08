APP_NAME := $(shell basename $(PWD))
BUILD_DIR := bin
GO_FILES := $(shell find . -name '*.go' -print)


build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd
	@echo "Build complete. Executable: $(BUILD_DIR)/$(APP_NAME)"



run: build
	@echo "Running $(APP_NAME)..."
	@$(BUILD_DIR)/$(APP_NAME)



test:
	@echo "Running tests..."
	@go test -v ./...



clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete."