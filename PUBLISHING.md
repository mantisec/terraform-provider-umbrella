# Publishing to Terraform Registry

This document explains how to set up and publish the Terraform Provider for Cisco Umbrella to the Terraform Registry.

## Prerequisites

1. **GitHub Repository**: Your provider must be hosted on GitHub
2. **GPG Key**: Required for signing releases
3. **Terraform Registry Account**: Sign up at [registry.terraform.io](https://registry.terraform.io)

## Setup Steps

### 1. GPG Key Setup

Generate a GPG key for signing releases:

```bash
# Generate a new GPG key
gpg --full-generate-key

# Export the private key (keep this secure!)
gpg --armor --export-secret-keys YOUR_KEY_ID > private-key.asc

# Export the public key
gpg --armor --export YOUR_KEY_ID > public-key.asc
```

### 2. GitHub Secrets Configuration

Add the following secrets to your GitHub repository (`Settings > Secrets and variables > Actions`):

- `GPG_PRIVATE_KEY`: Content of your private GPG key (private-key.asc)
- `PASSPHRASE`: Passphrase for your GPG key (if you set one)

### 3. Repository Configuration

Ensure your repository follows the naming convention:
- Repository name: `terraform-provider-{NAME}`
- For this provider: `terraform-provider-umbrella`

### 4. Terraform Registry Setup

1. Go to [registry.terraform.io](https://registry.terraform.io)
2. Sign in with your GitHub account
3. Click "Publish" > "Provider"
4. Select your repository (`terraform-provider-umbrella`)
5. Follow the verification steps

## Release Process

### Creating a Release

1. **Tag your release** following semantic versioning:
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```

2. **GitHub Actions will automatically**:
   - Build binaries for multiple platforms
   - Sign the release with your GPG key
   - Create checksums
   - Publish to GitHub Releases

3. **Terraform Registry will automatically**:
   - Detect the new release
   - Import the provider
   - Make it available for public use

### Version Numbering

Follow [Semantic Versioning](https://semver.org/):
- `v1.0.0` - Major release
- `v1.1.0` - Minor release (new features)
- `v1.0.1` - Patch release (bug fixes)

## Files Created for Publishing

The following files have been added to support publishing:

### `.github/workflows/release.yml`
GitHub Actions workflow that:
- Triggers on version tags (`v*`)
- Uses GoReleaser to build and package
- Signs releases with GPG
- Publishes to GitHub Releases

### `.github/workflows/test.yml`
GitHub Actions workflow that:
- Runs on pull requests and pushes
- Builds the provider
- Runs linting
- Executes tests

### `.goreleaser.yml`
GoReleaser configuration that:
- Builds for multiple platforms (Windows, Linux, macOS, FreeBSD)
- Creates proper archive naming
- Generates checksums
- Signs releases

### `terraform-registry-manifest.json`
Terraform Registry manifest that:
- Specifies protocol version compatibility
- Required for registry publishing

## Testing Before Release

Before creating a release, test locally:

```bash
# Build the provider
make build

# Run tests
make test

# Test with GoReleaser (dry run)
goreleaser release --snapshot --clean
```

## Provider Usage After Publishing

Once published, users can use your provider like this:

```hcl
terraform {
  required_providers {
    umbrella = {
      source  = "mantisec/umbrella"
      version = "~> 1.0"
    }
  }
}

provider "umbrella" {
  api_key    = var.umbrella_api_key
  api_secret = var.umbrella_api_secret
  org_id     = var.umbrella_org_id
}
```

## Troubleshooting

### Common Issues

1. **GPG Signing Fails**
   - Verify GPG_PRIVATE_KEY secret is correctly formatted
   - Ensure PASSPHRASE secret matches your key

2. **Build Fails**
   - Check Go version compatibility in workflows
   - Verify all dependencies are properly declared

3. **Registry Import Fails**
   - Ensure repository name follows convention
   - Verify terraform-registry-manifest.json is present
   - Check that releases are properly signed

### Getting Help

- [Terraform Registry Documentation](https://www.terraform.io/docs/registry/providers/publishing.html)
- [GoReleaser Documentation](https://goreleaser.com/)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)

## Security Considerations

- Never commit GPG private keys to the repository
- Use GitHub Secrets for sensitive information
- Regularly rotate GPG keys
- Monitor release signatures for integrity