#!/bin/bash

# Release script for terraform-provider-umbrella
# This script helps create and push a new release tag

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if version is provided
if [ $# -eq 0 ]; then
    print_error "Please provide a version number"
    echo "Usage: $0 <version>"
    echo "Example: $0 1.0.0"
    exit 1
fi

VERSION=$1

# Validate version format (basic semver check)
if [[ ! $VERSION =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    print_error "Version must be in format X.Y.Z (e.g., 1.0.0)"
    exit 1
fi

# Check if we're on main/master branch
CURRENT_BRANCH=$(git branch --show-current)
if [[ "$CURRENT_BRANCH" != "main" && "$CURRENT_BRANCH" != "master" ]]; then
    print_warning "You're not on main/master branch (current: $CURRENT_BRANCH)"
    read -p "Continue anyway? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        print_status "Aborted"
        exit 1
    fi
fi

# Check if working directory is clean
if [[ -n $(git status --porcelain) ]]; then
    print_error "Working directory is not clean. Please commit or stash changes."
    git status --short
    exit 1
fi

# Check if tag already exists
if git tag -l | grep -q "^v$VERSION$"; then
    print_error "Tag v$VERSION already exists"
    exit 1
fi

print_status "Preparing release v$VERSION"

# Update version in Makefile if needed
if grep -q "VERSION?=" Makefile; then
    print_status "Updating version in Makefile"
    sed -i.bak "s/VERSION?=.*/VERSION?=$VERSION/" Makefile
    rm -f Makefile.bak
    git add Makefile
fi

# Run tests
print_status "Running tests..."
if ! make test; then
    print_error "Tests failed. Please fix before releasing."
    exit 1
fi

# Build to ensure everything compiles
print_status "Building provider..."
if ! make build; then
    print_error "Build failed. Please fix before releasing."
    exit 1
fi

# Commit version changes if any
if [[ -n $(git status --porcelain) ]]; then
    print_status "Committing version updates"
    git commit -m "Bump version to v$VERSION"
fi

# Create and push tag
print_status "Creating tag v$VERSION"
git tag -a "v$VERSION" -m "Release v$VERSION"

print_status "Pushing tag to origin"
git push origin "v$VERSION"

print_status "Release v$VERSION has been tagged and pushed!"
print_status "GitHub Actions will now build and publish the release."
print_status "Check the Actions tab in your GitHub repository for progress."

# Provide next steps
echo
print_status "Next steps:"
echo "1. Monitor GitHub Actions: https://github.com/mantisec/terraform-provider-umbrella/actions"
echo "2. Verify release: https://github.com/mantisec/terraform-provider-umbrella/releases"
echo "3. Check Terraform Registry: https://registry.terraform.io/providers/mantisec/umbrella"