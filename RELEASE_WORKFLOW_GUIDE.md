# GitHub Release Workflow Guide

## Current Issue Analysis

The release workflow is **working correctly** but only triggers on **new tag pushes**, not on regular commits or successful builds. 

**Current Status:**
- ✅ Release workflow exists and is properly configured
- ✅ Tag `v0.0.1` exists and should have triggered a release
- ❌ No new tags have been created since v0.0.1

## How the Current Release Process Works

### 1. Tag-Based Releases (Current Setup)
The release workflow triggers only when you push a new tag starting with 'v':

```bash
# Create and push a new tag to trigger a release
git tag v0.0.2
git push origin v0.0.2
```

### 2. Required Secrets
The workflow requires these GitHub secrets to be configured:
- `GPG_PRIVATE_KEY` - Your GPG private key for signing releases
- `PASSPHRASE` - Passphrase for your GPG key
- `GITHUB_TOKEN` - Automatically provided by GitHub

## Solutions

### Option 1: Create a New Release (Recommended)
Create and push a new tag to trigger the release workflow:

```bash
# Create a new tag (increment version as appropriate)
git tag v0.0.2
git push origin v0.0.2
```

### Option 2: Add Automatic Release on Main Branch
If you want releases to be created automatically on successful builds to main:

```yaml
# Add this to .github/workflows/release.yml
on:
  push:
    tags:
      - 'v*'
  # Add automatic release on main branch
  push:
    branches:
      - main
    paths-ignore:
      - 'README.md'
      - 'docs/**'
```

### Option 3: Manual Release Trigger
Add a workflow_dispatch trigger for manual releases:

```yaml
on:
  push:
    tags:
      - 'v*'
  # Add manual trigger
  workflow_dispatch:
    inputs:
      version:
        description: 'Release version (e.g., v0.0.2)'
        required: true
        type: string
```

### Option 4: Automated Semantic Versioning
Use semantic-release for automatic version bumping:

```yaml
# New workflow: .github/workflows/semantic-release.yml
name: Semantic Release
on:
  push:
    branches: [main]
jobs:
  release:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0
      - uses: actions/setup-node@v4
        with:
          node-version: 18
      - run: npm install -g semantic-release @semantic-release/github
      - run: semantic-release
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

## Checking Release Status

### 1. Verify GitHub Secrets
Check if required secrets are configured in your repository:
- Go to Settings → Secrets and variables → Actions
- Ensure `GPG_PRIVATE_KEY` and `PASSPHRASE` are set

### 2. Check Workflow Runs
- Go to Actions tab in your GitHub repository
- Look for "Release" workflow runs
- Check if any runs failed due to missing secrets

### 3. Verify GPG Key Setup
```bash
# Check if GPG key is properly configured
gpg --list-secret-keys --keyid-format LONG
```

## Recommended Next Steps

### Immediate Action (Create New Release)
```bash
# 1. Ensure you're on the latest main branch
git checkout main
git pull origin main

# 2. Create and push a new tag
git tag v0.0.2
git push origin v0.0.2

# 3. Check GitHub Actions for the release workflow
```

### Long-term Improvements
1. **Set up GPG signing** if not already configured
2. **Add workflow_dispatch** for manual releases
3. **Consider semantic versioning** for automated releases
4. **Add release notes automation**

## Troubleshooting

### Common Issues
1. **Missing GPG secrets** - Release will fail during signing
2. **Insufficient permissions** - Ensure GITHUB_TOKEN has write access
3. **GoReleaser configuration** - Check .goreleaser.yml syntax
4. **Tag format** - Must start with 'v' (e.g., v1.0.0, not 1.0.0)

### Debug Commands
```bash
# Check existing tags
git tag -l

# Check remote tags
git ls-remote --tags origin

# Check workflow file syntax
yamllint .github/workflows/release.yml
```

## Current Workflow Analysis

Your current setup is **production-ready** and follows best practices:
- ✅ Uses GoReleaser for multi-platform builds
- ✅ Includes GPG signing for security
- ✅ Proper permissions and token usage
- ✅ Fetches full git history for changelog generation
- ✅ Creates checksums and signatures

The workflow will create a GitHub release with:
- Multi-platform binaries (Linux, macOS, Windows, FreeBSD)
- SHA256 checksums
- GPG signatures
- Terraform registry manifest