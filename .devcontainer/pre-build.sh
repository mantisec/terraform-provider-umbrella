#!/bin/bash
# Pre-build script - runs on the host before container creation

echo "🏗️ Pre-build: Preparing for dev container creation..."

# Ensure the .devcontainer directory exists and scripts are executable
chmod +x .devcontainer/*.sh

echo "✅ Pre-build complete!"