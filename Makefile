# Set the name of your binary
BINARY_NAME=nhsnumber_bin

# Set the Go compiler
GO=go

# Set the flags for the Go compiler
GOFLAGS=

# Set the source files for your project
SOURCES=$(wildcard *.go)


# Set the output directory
OUTPUT_DIR=bin

# Set the output binary
OUTPUT_BINARY=$(OUTPUT_DIR)/$(BINARY_NAME)

# Set the build flags
BUILD_FLAGS=

# Set the clean flags
CLEAN_FLAGS=-r

# Set the default target
.DEFAULT_GOAL: all

# Build the binary
$(OUTPUT_BINARY): $(SOURCES)
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(OUTPUT_DIR)
	$(GO) build $(GOFLAGS) $(BUILD_FLAGS) -o $(OUTPUT_BINARY) $(SOURCES)

# Clean the build directory
clean:
	@echo "Cleaning..."
	@rm $(CLEAN_FLAGS) $(OUTPUT_DIR)

# Build the binary and clean the build directory
all: clean $(OUTPUT_BINARY)
