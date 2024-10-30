# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=pruner
BUILD_DIR=build

# Build the binary
build:
	mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v ./cmd/main.go

# Clean the binary
clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

# Run the binary
run: build
	./$(BUILD_DIR)/$(BINARY_NAME) 
# Default target
all: build

.PHONY: build clean run all

# Debug build for debugging with dlv or any other debugger
debug:
	mkdir -p $(BUILD_DIR)
	$(GOBUILD) -gcflags="all=-N -l" -o $(BUILD_DIR)/$(BINARY_NAME) -v ./cmd/main.go


# Install the binary
install:
	@echo "Installing $(BINARY_NAME)..."
	@if [ "$$(uname)" = "Darwin" ]; then \
		$(GOBUILD) -o /usr/local/bin/$(BINARY_NAME) -v ./cmd/main.go; \
	elif [ "$$(uname)" = "Linux" ]; then \
		$(GOBUILD) -o /usr/local/bin/$(BINARY_NAME) -v ./cmd/main.go; \
	else \
		echo "Unsupported operating system. Please install manually."; \
		exit 1; \
	fi
	@echo "$(BINARY_NAME) has been installed to /usr/local/bin/$(BINARY_NAME)"

.PHONY: install

