#!/bin/bash
# On-create script - runs once when the container is first created

echo "🎯 On-create: Setting up container for first use..."

# Set proper permissions for scripts
chmod +x .devcontainer/*.sh
chmod +x scripts/*.sh 2>/dev/null || true
chmod +x scripts/dev/*.sh 2>/dev/null || true

# Create necessary directories
mkdir -p ~/.terraform.d/plugins
mkdir -p ~/.config/go
mkdir -p /tmp/terraform-provider-dev

# Set up Go environment
export GOPATH=/go
export PATH=$GOPATH/bin:$PATH

echo "✅ On-create setup complete!"