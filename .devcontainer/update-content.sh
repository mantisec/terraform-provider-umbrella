#!/bin/bash
# Update-content script - runs when container content is updated

echo "🔄 Update-content: Refreshing development environment..."

# Update Go dependencies
cd /workspaces/terraform-provider-umbrella
go mod download
go mod tidy

# Update Go tools if needed
echo "🔧 Updating Go tools..."
go install -a github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install -a golang.org/x/tools/cmd/goimports@latest
go install -a github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@latest

# Rebuild the provider to ensure everything still works
echo "🔨 Rebuilding provider..."
go build -o terraform-provider-umbrella

echo "✅ Update-content complete!"