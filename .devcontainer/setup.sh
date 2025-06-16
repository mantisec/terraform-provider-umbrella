#!/bin/bash
set -e

echo "🚀 Setting up Terraform Provider Umbrella development environment..."

# Update package lists
sudo apt-get update

# Install essential development tools
echo "📦 Installing essential development tools..."
sudo apt-get install -y \
    curl \
    wget \
    git \
    make \
    build-essential \
    unzip \
    jq \
    tree \
    htop \
    vim \
    nano \
    ca-certificates \
    gnupg \
    lsb-release

# Install Go tools
echo "🔧 Installing Go development tools..."
go install -a github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install -a golang.org/x/tools/cmd/goimports@latest
go install -a golang.org/x/tools/cmd/godoc@latest
go install -a github.com/go-delve/delve/cmd/dlv@latest
go install -a honnef.co/go/tools/cmd/staticcheck@latest
go install -a github.com/fatih/gomodifytags@latest
go install -a github.com/josharian/impl@latest
go install -a github.com/cweill/gotests/gotests@latest

# Install GoReleaser for building releases
echo "📦 Installing GoReleaser..."
curl -sfL https://goreleaser.com/static/run | bash -s -- --version
echo 'deb [trusted=yes] https://repo.goreleaser.com/apt/ /' | sudo tee /etc/apt/sources.list.d/goreleaser.list
sudo apt-get update
sudo apt-get install -y goreleaser

# Install Terraform Plugin Framework tools
echo "🔌 Installing Terraform Plugin Framework tools..."
go install -a github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@latest

# Install additional Terraform tools
echo "🏗️ Installing additional Terraform tools..."
# Install terraform-docs
curl -sSLo ./terraform-docs.tar.gz https://terraform-docs.io/dl/v0.16.0/terraform-docs-v0.16.0-$(uname)-amd64.tar.gz
tar -xzf terraform-docs.tar.gz
chmod +x terraform-docs
sudo mv terraform-docs /usr/local/bin/terraform-docs
rm terraform-docs.tar.gz

# Install tfsec for security scanning
curl -s https://raw.githubusercontent.com/aquasecurity/tfsec/master/scripts/install_linux.sh | bash
sudo mv tfsec /usr/local/bin/

# Install checkov for additional security scanning
pip3 install checkov

# Install pre-commit for git hooks
pip3 install pre-commit

# Set up Go workspace
echo "🏠 Setting up Go workspace..."
mkdir -p /go/src /go/bin /go/pkg
export GOPATH=/go
export PATH=$GOPATH/bin:$PATH

# Install project dependencies
echo "📚 Installing project dependencies..."
cd /workspaces/terraform-provider-umbrella
go mod download
go mod tidy

# Build the project to ensure everything works
echo "🔨 Building project to verify setup..."
go build -o terraform-provider-umbrella

# Set up git configuration (if not already set)
echo "🔧 Configuring Git..."
if [ -z "$(git config --global user.name)" ]; then
    git config --global user.name "Dev Container User"
fi
if [ -z "$(git config --global user.email)" ]; then
    git config --global user.email "dev@example.com"
fi

# Set up shell aliases and functions
echo "🐚 Setting up shell aliases..."
cat >> ~/.zshrc << 'EOF'

# Terraform Provider Development Aliases
alias tf='terraform'
alias tfi='terraform init'
alias tfp='terraform plan'
alias tfa='terraform apply'
alias tfd='terraform destroy'
alias tfv='terraform validate'
alias tff='terraform fmt'

# Go Development Aliases
alias gob='go build'
alias gor='go run'
alias got='go test'
alias gom='go mod'
alias gof='go fmt'
alias gov='go vet'
alias gol='golangci-lint run'

# Project Specific Aliases
alias gen='go run tools/generator/cmd/generate/main.go'
alias gen-full='make generate-full'
alias clean-gen='make clean-generated'
alias build-provider='go build -o terraform-provider-umbrella'
alias test-provider='go test ./...'

# Utility Functions
function tf-init-local() {
    terraform init -plugin-dir=/workspaces/terraform-provider-umbrella
}

function provider-test() {
    echo "Running provider tests..."
    TF_ACC=1 go test ./internal/provider/tests/ -v -timeout 30m
}

function provider-build() {
    echo "Building provider..."
    go build -o terraform-provider-umbrella
    echo "Provider built successfully!"
}

function provider-install-local() {
    echo "Installing provider locally..."
    make install-local
    echo "Provider installed locally!"
}

# Environment setup
export GOPATH=/go
export PATH=$GOPATH/bin:$PATH
export GO111MODULE=on
export CGO_ENABLED=0

# Terraform environment
export TF_LOG=INFO
export TF_CLI_CONFIG_FILE=/workspaces/terraform-provider-umbrella/.terraformrc

EOF

# Create Terraform CLI configuration for local development
echo "🔧 Setting up Terraform CLI configuration..."
cat > /workspaces/terraform-provider-umbrella/.terraformrc << 'EOF'
provider_installation {
  dev_overrides {
    "local/mantisec/umbrella" = "/workspaces/terraform-provider-umbrella"
  }
  
  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
EOF

# Set up pre-commit hooks
echo "🪝 Setting up pre-commit hooks..."
cat > /workspaces/terraform-provider-umbrella/.pre-commit-config.yaml << 'EOF'
repos:
  - repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v4.4.0
    hooks:
      - id: trailing-whitespace
      - id: end-of-file-fixer
      - id: check-yaml
      - id: check-added-large-files
      - id: check-merge-conflict
  
  - repo: https://github.com/golangci/golangci-lint
    rev: v1.54.2
    hooks:
      - id: golangci-lint
  
  - repo: https://github.com/terraform-docs/terraform-docs
    rev: v0.16.0
    hooks:
      - id: terraform-docs-go
        args: ["markdown", "table", "--output-file", "README.md"]
EOF

# Initialize pre-commit
cd /workspaces/terraform-provider-umbrella
pre-commit install

# Create development scripts directory
mkdir -p /workspaces/terraform-provider-umbrella/scripts/dev

# Create a development helper script
cat > /workspaces/terraform-provider-umbrella/scripts/dev/dev-setup.sh << 'EOF'
#!/bin/bash
# Development setup helper script

echo "🚀 Terraform Provider Umbrella - Development Setup"
echo ""
echo "Available commands:"
echo "  gen-full      - Generate all provider code from API specs"
echo "  clean-gen     - Clean generated files"
echo "  build         - Build the provider"
echo "  test          - Run all tests"
echo "  test-acc      - Run acceptance tests"
echo "  lint          - Run linter"
echo "  fmt           - Format code"
echo "  docs          - Generate documentation"
echo "  install       - Install provider locally"
echo ""
echo "Environment:"
echo "  Go version: $(go version)"
echo "  Terraform version: $(terraform version)"
echo "  GoReleaser version: $(goreleaser --version)"
echo ""
echo "Project structure:"
tree -L 2 -I 'node_modules|.git'
EOF

chmod +x /workspaces/terraform-provider-umbrella/scripts/dev/dev-setup.sh

# Final verification
echo "✅ Development environment setup complete!"
echo ""
echo "🔍 Verification:"
echo "  Go version: $(go version)"
echo "  Terraform version: $(terraform version)"
echo "  Make version: $(make --version | head -n1)"
echo "  Git version: $(git --version)"
echo ""
echo "🎉 Ready for Terraform Provider development!"
echo ""
echo "💡 Useful commands:"
echo "  make help                    - Show all available make targets"
echo "  ./scripts/dev/dev-setup.sh   - Show development helper info"
echo "  gen-full                     - Generate provider code (alias)"
echo "  build-provider               - Build the provider (alias)"
echo "  test-provider                - Run tests (alias)"
echo ""
echo "📁 Project is located at: /workspaces/terraform-provider-umbrella"