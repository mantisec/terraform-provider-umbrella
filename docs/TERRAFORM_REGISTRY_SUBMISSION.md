# Terraform Registry Submission Guide

This guide provides step-by-step instructions for submitting the Terraform Provider for Cisco Umbrella to the official Terraform Registry.

## Overview

The Terraform Registry is the official repository for Terraform providers. Publishing to the registry makes your provider easily discoverable and installable by users worldwide.

## Prerequisites

Before submitting to the Terraform Registry, ensure you have:

- ✅ **GitHub Repository**: Provider hosted on GitHub with proper naming (`terraform-provider-umbrella`)
- ✅ **GPG Signing**: GPG key configured for signing releases (see [GPG Setup Guide](GPG_SETUP_GUIDE.md))
- ✅ **GitHub Actions**: Automated release workflow configured
- ✅ **Documentation**: Complete provider documentation following Terraform standards
- ✅ **Testing**: Comprehensive test suite with passing tests
- ✅ **Examples**: Working examples for all resources and data sources
- ✅ **Versioning**: Proper semantic versioning with tagged releases

## Step 1: Repository Requirements

### Repository Structure
Ensure your repository follows the required structure:

```
terraform-provider-umbrella/
├── .github/
│   └── workflows/
│       ├── release.yml          # Release automation
│       └── test.yml             # Testing workflow
├── docs/                        # Provider documentation
│   ├── index.md                 # Provider documentation
│   ├── data-sources/            # Data source docs
│   ├── resources/               # Resource docs
│   └── guides/                  # Usage guides
├── examples/                    # Usage examples
├── internal/                    # Provider implementation
├── .goreleaser.yml             # Release configuration
├── terraform-registry-manifest.json  # Registry manifest
├── go.mod                      # Go module definition
├── main.go                     # Provider entry point
├── README.md                   # Repository README
└── CHANGELOG.md                # Version history
```

### Repository Settings
1. **Repository Name**: Must be `terraform-provider-umbrella`
2. **Visibility**: Public repository required
3. **Topics**: Add relevant topics (e.g., `terraform`, `terraform-provider`, `cisco`, `umbrella`)
4. **Description**: Clear description of the provider's purpose
5. **License**: Include appropriate license (e.g., MIT, Apache 2.0)

### Required Files

#### terraform-registry-manifest.json
```json
{
  "version": 1,
  "metadata": {
    "protocol_versions": ["5.0"]
  }
}
```

#### .goreleaser.yml
Must include proper configuration for:
- Multi-platform builds
- GPG signing
- Checksum generation
- Registry manifest inclusion

## Step 2: Documentation Requirements

### Provider Documentation (docs/index.md)
Must include:
- Provider description and purpose
- Authentication configuration
- Required and optional arguments
- Usage examples
- Version compatibility information

### Resource Documentation (docs/resources/*.md)
Each resource must have:
- Resource description
- Argument reference (required/optional)
- Attribute reference (computed values)
- Import instructions
- Usage examples

### Data Source Documentation (docs/data-sources/*.md)
Each data source must have:
- Data source description
- Argument reference
- Attribute reference
- Usage examples

### Guides (docs/guides/*.md)
Recommended guides:
- Getting started guide
- Migration guide (if applicable)
- Troubleshooting guide
- Best practices

## Step 3: Release Requirements

### Version Tags
- Must use semantic versioning (e.g., `v1.0.0`)
- Tags must start with `v`
- Follow [Semantic Versioning](https://semver.org/) principles

### Release Assets
Each release must include:
- Multi-platform binaries (Linux, macOS, Windows, FreeBSD)
- SHA256 checksums
- GPG signatures
- terraform-registry-manifest.json

### Supported Platforms
Minimum required platforms:
- `linux_amd64`
- `linux_arm64`
- `darwin_amd64`
- `darwin_arm64`
- `windows_amd64`

## Step 4: Testing Requirements

### Test Coverage
- Unit tests for all resources and data sources
- Integration tests with actual API (where possible)
- Acceptance tests following Terraform testing standards
- Test coverage reporting

### Continuous Integration
- Automated testing on pull requests
- Multi-version Go testing
- Linting and code quality checks
- Security scanning

## Step 5: Submission Process

### 1. Prepare Repository
Ensure all requirements are met:

```powershell
# Run the preparation script
.\scripts\prepare-release.ps1 -Version "1.0.0" -DryRun

# Create the release
.\scripts\create-release.ps1 -Version "1.0.0"
```

### 2. Verify Release
After creating a release, verify:
- ✅ Release appears on GitHub
- ✅ All platform binaries are present
- ✅ Checksums and signatures are valid
- ✅ terraform-registry-manifest.json is included

### 3. Submit to Registry

#### Access Terraform Registry
1. Go to [registry.terraform.io](https://registry.terraform.io)
2. Sign in with your GitHub account
3. Ensure your GitHub account has access to the repository

#### Publish Provider
1. Click **"Publish"** in the top navigation
2. Select **"Provider"**
3. Choose your repository: `mantisec/terraform-provider-umbrella`
4. Click **"Publish Provider"**

#### Verification Process
The registry will:
1. Verify repository structure
2. Check for required files
3. Validate documentation
4. Verify GPG signatures
5. Import the latest release

### 4. Upload GPG Public Key
1. In the registry, go to your provider page
2. Navigate to **Settings** > **Signing Keys**
3. Upload your GPG public key
4. Verify the key fingerprint matches

## Step 6: Post-Submission

### Verification
After submission, verify:
- ✅ Provider appears in search results
- ✅ Documentation renders correctly
- ✅ Examples work as expected
- ✅ Installation works via Terraform

### Test Installation
Create a test configuration:

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

Test commands:
```bash
terraform init
terraform plan
```

## Step 7: Maintenance

### Regular Updates
- Keep documentation up to date
- Release bug fixes and new features
- Maintain compatibility with Terraform versions
- Update dependencies regularly

### Community Engagement
- Respond to issues and pull requests
- Maintain changelog
- Provide migration guides for breaking changes
- Engage with user feedback

## Troubleshooting

### Common Submission Issues

#### "Repository not found or access denied"
**Solution**: 
- Verify repository is public
- Ensure GitHub account has access
- Check repository name matches exactly

#### "Invalid terraform-registry-manifest.json"
**Solution**:
- Verify JSON syntax is correct
- Ensure protocol_versions includes supported versions
- Check file is in repository root

#### "Missing required documentation"
**Solution**:
- Ensure docs/index.md exists and is complete
- Verify all resources have documentation
- Check documentation follows required format

#### "GPG signature verification failed"
**Solution**:
- Verify GPG key is properly configured
- Check signatures are present in release
- Upload correct public key to registry

#### "Unsupported platform"
**Solution**:
- Ensure all required platforms are built
- Check .goreleaser.yml configuration
- Verify release assets include all platforms

### Debug Steps

1. **Check Repository Structure**
   ```bash
   # Verify required files exist
   ls -la terraform-registry-manifest.json
   ls -la docs/index.md
   ls -la .goreleaser.yml
   ```

2. **Validate Documentation**
   ```bash
   # Generate and check docs
   tfplugindocs generate --provider-name umbrella
   ```

3. **Test Release Process**
   ```bash
   # Test GoReleaser configuration
   goreleaser check
   goreleaser release --snapshot --clean
   ```

4. **Verify GPG Setup**
   ```bash
   # Check GPG key
   gpg --list-secret-keys
   gpg --armor --export YOUR_KEY_ID
   ```

## Registry Policies

### Content Guidelines
- Providers must be functional and well-documented
- No malicious or harmful code
- Respect intellectual property rights
- Follow community standards

### Naming Conventions
- Repository: `terraform-provider-{name}`
- Provider address: `{namespace}/{name}`
- Resources: `{provider}_{resource_type}`

### Version Management
- Use semantic versioning
- Maintain backward compatibility when possible
- Provide migration guides for breaking changes
- Tag releases properly

## Support and Resources

### Official Documentation
- [Terraform Registry Documentation](https://www.terraform.io/docs/registry/providers/publishing.html)
- [Provider Development](https://www.terraform.io/docs/extend/writing-custom-providers.html)
- [GoReleaser Documentation](https://goreleaser.com/)

### Community Resources
- [Terraform Provider Development Program](https://www.terraform.io/docs/partnerships/index.html)
- [HashiCorp Community Forum](https://discuss.hashicorp.com/)
- [Terraform GitHub Discussions](https://github.com/hashicorp/terraform/discussions)

### Getting Help
If you encounter issues during submission:
1. Check the troubleshooting section above
2. Review official documentation
3. Search existing issues in the Terraform Registry repository
4. Contact HashiCorp support if needed

---

## Submission Checklist

Before submitting to the Terraform Registry, ensure:

### Repository Requirements
- ✅ Repository name: `terraform-provider-umbrella`
- ✅ Public repository on GitHub
- ✅ Proper repository structure
- ✅ All required files present
- ✅ Clear README and description

### Documentation
- ✅ Complete provider documentation (docs/index.md)
- ✅ All resources documented (docs/resources/*.md)
- ✅ All data sources documented (docs/data-sources/*.md)
- ✅ Usage guides and examples
- ✅ Documentation follows Terraform standards

### Release Process
- ✅ GPG signing configured
- ✅ GitHub Actions workflows working
- ✅ GoReleaser configuration valid
- ✅ Multi-platform builds
- ✅ Proper semantic versioning

### Testing
- ✅ Comprehensive test suite
- ✅ All tests passing
- ✅ Continuous integration configured
- ✅ Code quality checks

### Registry Submission
- ✅ Terraform Registry account created
- ✅ Repository connected to registry
- ✅ GPG public key uploaded
- ✅ Provider published successfully
- ✅ Installation tested

**Ready for Production**: Once all items are checked, your provider is ready for public use! 🚀