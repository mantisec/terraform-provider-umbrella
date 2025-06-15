# Terraform Provider for Cisco Umbrella - Build Automation

# Variables
BINARY_NAME=terraform-provider-umbrella
VERSION?=0.2.0
GOOS?=windows
GOARCH?=amd64

# Default target
.PHONY: all
all: build

# Build the provider
.PHONY: build
build:
	go build -o $(BINARY_NAME).exe

# Build for specific OS/Architecture
.PHONY: build-cross
build-cross:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BINARY_NAME)-$(GOOS)-$(GOARCH)

# Run tests
.PHONY: test
test:
	go test ./...

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Format code
.PHONY: fmt
fmt:
	go fmt ./...

# Lint code
.PHONY: lint
lint:
	golangci-lint run

# Clean build artifacts
.PHONY: clean
clean:
	rm -f $(BINARY_NAME).exe
	rm -f $(BINARY_NAME)-*
	rm -f coverage.out coverage.html

# Install dependencies
.PHONY: deps
deps:
	go mod tidy
	go mod download

# Generate documentation
.PHONY: docs
docs:
	tfplugindocs generate

# Install the provider locally for development
.PHONY: install-local
install-local: build
	mkdir -p ~/.terraform.d/plugins/local/mantisec/umbrella/$(VERSION)/windows_amd64/
	cp $(BINARY_NAME).exe ~/.terraform.d/plugins/local/mantisec/umbrella/$(VERSION)/windows_amd64/

# Validate Terraform examples
.PHONY: validate-examples
validate-examples:
	@for dir in examples/*/; do \
		echo "Validating $$dir"; \
		cd $$dir && terraform init && terraform validate && cd ../..; \
	done

# Help target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build           - Build the provider binary"
	@echo "  build-cross     - Cross-compile for specific OS/Architecture"
	@echo "  test            - Run tests"
	@echo "  test-coverage   - Run tests with coverage report"
	@echo "  fmt             - Format Go code"
	@echo "  lint            - Run linter"
	@echo "  clean           - Clean build artifacts"
	@echo "  deps            - Install/update dependencies"
	@echo "  docs            - Generate documentation"
	@echo "  install-local   - Install provider locally for development"
	@echo "  validate-examples - Validate all Terraform examples"
	@echo "  help            - Show this help message"