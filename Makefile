# Nombre del ejecutable de tu aplicación
APP_NAME := $(shell basename $(PWD))

# Directorio de salida para el ejecutable
BUILD_DIR := bin

# Go source files
GO_FILES := $(shell find . -name '*.go' -print)


# Comando para construir la aplicación
# Ahora go build busca el paquete main dentro del directorio ./cmd
build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd
	@echo "Build complete. Executable: $(BUILD_DIR)/$(APP_NAME)"



# Comando para ejecutar la aplicación
run: build
	@echo "Running $(APP_NAME)..."
	@$(BUILD_DIR)/$(APP_NAME)



# Comando para ejecutar las pruebas unitarias
test:
	@echo "Running tests..."
	@go test -v ./...



# Comando para limpiar los artefactos de la construcción
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete."