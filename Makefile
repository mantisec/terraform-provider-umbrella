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

# Clean generated files from API specs
.PHONY: clean-generated
clean-generated:
	@echo "Cleaning generated files from API specifications..."
	# Generated Go files in internal/provider/
	rm -f internal/provider/generated_*.go
	# Generated test files (keeping manually created ones)
	rm -f internal/provider/tests/*_test.go
	# Keep provider_test.go as it's manually created
	git checkout HEAD -- internal/provider/tests/provider_test.go 2>/dev/null || true
	# Generated documentation files
	rm -f docs/resources/*.md
	# Keep index.md as it's manually maintained
	git checkout HEAD -- docs/index.md 2>/dev/null || true
	@echo "Generated files cleaned successfully!"

# Clean all (build artifacts + generated files)
.PHONY: clean-all
clean-all: clean clean-generated
	@echo "All build artifacts and generated files cleaned!"

# Install dependencies
.PHONY: deps
deps:
	go mod tidy
	go mod download

# Generate documentation
.PHONY: docs
docs:
	tfplugindocs generate

# Generate provider code from OpenAPI specs
.PHONY: generate
generate:
	go run tools/generator/cmd/generate/main.go

# Generate with full Phase 2 features (client methods, docs, tests)
.PHONY: generate-full
generate-full:
	@echo "Generating Mantisec Umbrella provider code with advanced features..."
	go run tools/generator/cmd/generate/main.go
	@echo "Formatting generated code..."
	go fmt ./internal/provider/generated_*.go
	@echo "Validating generated code..."
	go vet ./internal/provider/generated_*.go
	@echo "Generation complete!"

# Generate only client methods
.PHONY: generate-client
generate-client:
	@echo "Generating client methods..."
	go run tools/generator/cmd/generate/main.go -client-only
	go fmt ./internal/provider/generated_client_methods.go

# Generate only documentation
.PHONY: generate-docs
generate-docs:
	@echo "Generating documentation..."
	go run tools/generator/cmd/generate/main.go -docs-only
	@echo "Documentation generated in docs/ directory"

# Generate only tests
.PHONY: generate-tests
generate-tests:
	@echo "Generating test files..."
	go run tools/generator/cmd/generate/main.go -tests-only
	go fmt ./internal/provider/tests/*_test.go

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
	@echo "  build             - Build the provider binary"
	@echo "  build-cross       - Cross-compile for specific OS/Architecture"
	@echo "  test              - Run tests"
	@echo "  test-coverage     - Run tests with coverage report"
	@echo "  fmt               - Format Go code"
	@echo "  lint              - Run linter"
	@echo "  clean             - Clean build artifacts"
	@echo "  clean-generated   - Clean generated files from API specs"
	@echo "  clean-all         - Clean build artifacts + generated files"
	@echo "  deps              - Install/update dependencies"
	@echo "  docs              - Generate documentation"
	@echo "  generate          - Generate provider code from OpenAPI specs"
	@echo "  generate-full     - Generate Phase 2 provider with all features"
	@echo "  generate-client   - Generate only client methods"
	@echo "  generate-docs     - Generate only documentation"
	@echo "  generate-tests    - Generate only test files"
	@echo "  install-local     - Install provider locally for development"
	@echo "  validate-examples - Validate all Terraform examples"
	@echo "  release-test      - Test release build with GoReleaser"
	@echo "  help              - Show this help message"

# Test release build with GoReleaser (requires GoReleaser to be installed)
.PHONY: release-test
release-test:
	goreleaser release --snapshot --clean