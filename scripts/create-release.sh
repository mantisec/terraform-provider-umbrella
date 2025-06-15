#!/bin/bash

# Create Release Script for terraform-provider-umbrella
# This script helps create new releases by tagging and pushing to trigger the release workflow

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if we're in a git repository
if ! git rev-parse --git-dir > /dev/null 2>&1; then
    print_error "Not in a git repository"
    exit 1
fi

# Check if we're on main branch
current_branch=$(git branch --show-current)
if [ "$current_branch" != "main" ]; then
    print_warning "You're on branch '$current_branch', not 'main'"
    read -p "Continue anyway? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        print_error "Aborted"
        exit 1
    fi
fi

# Check for uncommitted changes
if ! git diff-index --quiet HEAD --; then
    print_error "You have uncommitted changes. Please commit or stash them first."
    git status --porcelain
    exit 1
fi

# Get current version
current_version=$(git tag -l | grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$' | sort -V | tail -n1)
if [ -z "$current_version" ]; then
    current_version="v0.0.0"
    print_warning "No previous version found, starting from $current_version"
else
    print_status "Current version: $current_version"
fi

# Parse version components
if [[ $current_version =~ ^v([0-9]+)\.([0-9]+)\.([0-9]+)$ ]]; then
    major=${BASH_REMATCH[1]}
    minor=${BASH_REMATCH[2]}
    patch=${BASH_REMATCH[3]}
else
    major=0
    minor=0
    patch=0
fi

# Calculate next versions
next_patch="v$major.$minor.$((patch + 1))"
next_minor="v$major.$((minor + 1)).0"
next_major="v$((major + 1)).0.0"

echo
print_status "Available version options:"
echo "  1) Patch release: $next_patch (bug fixes)"
echo "  2) Minor release: $next_minor (new features, backward compatible)"
echo "  3) Major release: $next_major (breaking changes)"
echo "  4) Custom version"
echo

read -p "Select version type (1-4): " -n 1 -r version_choice
echo

case $version_choice in
    1)
        new_version=$next_patch
        ;;
    2)
        new_version=$next_minor
        ;;
    3)
        new_version=$next_major
        ;;
    4)
        read -p "Enter custom version (e.g., v1.2.3): " new_version
        if [[ ! $new_version =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
            print_error "Invalid version format. Use vX.Y.Z (e.g., v1.2.3)"
            exit 1
        fi
        ;;
    *)
        print_error "Invalid choice"
        exit 1
        ;;
esac

# Check if tag already exists
if git tag -l | grep -q "^$new_version$"; then
    print_error "Tag $new_version already exists"
    exit 1
fi

print_status "Creating release $new_version"

# Confirm before proceeding
echo
print_warning "This will:"
echo "  - Create and push tag: $new_version"
echo "  - Trigger the GitHub Actions release workflow"
echo "  - Create a GitHub release with binaries"
echo
read -p "Continue? (y/N): " -n 1 -r
echo

if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    print_error "Aborted"
    exit 1
fi

# Pull latest changes
print_status "Pulling latest changes..."
git pull origin main

# Create and push tag
print_status "Creating tag $new_version..."
git tag -a "$new_version" -m "Release $new_version"

print_status "Pushing tag to origin..."
git push origin "$new_version"

print_success "Tag $new_version created and pushed!"
print_status "GitHub Actions will now build and create the release."
print_status "Check the progress at: https://github.com/$(git config --get remote.origin.url | sed 's/.*github.com[:/]\([^.]*\).*/\1/')/actions"

echo
print_success "Release process initiated! 🎉"
print_status "The release will be available at: https://github.com/$(git config --get remote.origin.url | sed 's/.*github.com[:/]\([^.]*\).*/\1/')/releases"