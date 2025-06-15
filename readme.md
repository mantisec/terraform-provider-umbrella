# Terraform Provider for Cisco Umbrella

A VERY opinionated bootleg Terraform provider for managing Cisco Umbrella Secure Web Gateway resources through the REST API.
features will be added as I need them... Deal with it :P 

> **Status**  : ⚠️ **Early preview / PoC** – Expect breaking changes until v1.0
>
> **Author** : Grant Lawton  |  Mantisec  |  [grant.lawton@mantisec.com.au](mailto\:grant.lawton@mantisec.com.au)


## Features

- **Destination Lists**: Create and manage URL, domain, and CIDR destination lists
- **IPSec Tunnels**: Create and manage IPSec tunnels for Secure Internet Gateway (SIG)
- **SAML Authentication**: Configure SAML SSO integration with identity providers
- **Rulesets**: Manage SWG policy rulesets with SAML and SSL decryption settings
- **Rules**: Create and manage individual policy rules within rulesets
- **OAuth2 Authentication**: Automatic token management with refresh capabilities

## Supported Resources

### `umbrella_destination_list`

Manages Umbrella destination lists for policy enforcement.

**Arguments:**
- `name` (Required) - Name of the destination list
- `type` (Required) - Type of destinations: `URL`, `DOMAIN`, or `CIDR`
- `destinations` (Optional) - Set of destination entries

**Attributes:**
- `id` - Unique identifier of the destination list

### `umbrella_tunnel`

Manages IPSec tunnels for Umbrella Secure Internet Gateway.

**Arguments:**
- `name` (Required) - Name of the tunnel
- `device_ip` (Required) - Device IP address for the tunnel endpoint
- `pre_shared_key` (Required, Sensitive) - Pre-shared key for IPSec authentication

**Attributes:**
- `id` - Unique identifier of the tunnel
- `status` - Current status of the tunnel
- `created_at` - Creation timestamp
- `updated_at` - Last update timestamp

### `umbrella_saml`

Manages SAML authentication configuration for SSO integration.

**Arguments:**
- `metadata_url` (Required) - SAML metadata URL from your identity provider
- `auth_type` (Required) - Authentication type (e.g., "AzureAD", "ADFS")

**Attributes:**
- `id` - SAML configuration identifier
- `enabled` - Whether SAML authentication is enabled

### `umbrella_ruleset`

Manages SWG policy rulesets with SAML and SSL decryption capabilities.

**Arguments:**
- `name` (Required) - Name of the ruleset
- `description` (Optional) - Description of the ruleset
- `saml_enabled` (Optional) - Enable SAML authentication for this ruleset
- `ssl_decryption_enabled` (Optional) - Enable SSL decryption for this ruleset

**Attributes:**
- `id` - Unique identifier of the ruleset
- `created_at` - Creation timestamp
- `updated_at` - Last update timestamp

### `umbrella_rule`

Manages individual policy rules within a ruleset.

**Arguments:**
- `ruleset_id` (Required) - ID of the ruleset this rule belongs to
- `name` (Required) - Name of the rule
- `action` (Required) - Rule action: `ALLOW`, `BLOCK`, `DO_NOT_DECRYPT`, etc.
- `rank` (Required) - Rule priority (lower numbers = higher priority)
- `destination_lists` (Optional) - Set of destination list names to apply this rule to
- `applications` (Optional) - Set of applications to apply this rule to
- `enabled` (Optional) - Whether the rule is enabled

**Attributes:**
- `id` - Unique identifier of the rule
- `created_at` - Creation timestamp
- `updated_at` - Last update timestamp

## Provider Configuration

```hcl
provider "umbrella" {
  api_key    = var.umbrella_api_key     # Umbrella API key (client ID)
  api_secret = var.umbrella_api_secret  # Umbrella API secret (client secret)
  org_id     = var.umbrella_org_id      # Umbrella organization ID
}
```

## Usage Examples

### Basic Destination List

```hcl
resource "umbrella_destination_list" "blocked_domains" {
  name = "Blocked Domains"
  type = "DOMAIN"
  destinations = [
    "malicious-site.com",
    "phishing-domain.net"
  ]
}
```

### IPSec Tunnel Configuration

```hcl
resource "umbrella_tunnel" "primary_tunnel" {
  name            = "Primary-SIG-Tunnel"
  device_ip       = "203.0.113.10"
  pre_shared_key  = var.tunnel_psk
}
```

### SAML Authentication Setup

```hcl
resource "umbrella_saml" "azure_ad" {
  metadata_url = "https://login.microsoftonline.com/your-tenant-id/federationmetadata/2007-06/federationmetadata.xml"
  auth_type    = "AzureAD"
}
```

### Ruleset with SAML and SSL Decryption

```hcl
resource "umbrella_ruleset" "default_web_policy" {
  name                     = "Default Web Policy"
  description              = "Main SWG policy with SAML enabled"
  saml_enabled             = true
  ssl_decryption_enabled   = true
}
```

### Policy Rules

```hcl
# Create destination list for bypass domains
resource "umbrella_destination_list" "azure_bypass" {
  name = "AzureAD-Bypass"
  type = "URL"
  destinations = [
    "login.microsoftonline.com",
    "msauth.net"
  ]
}

# Create bypass rule for Azure AD authentication
resource "umbrella_rule" "bypass_azure_ad" {
  ruleset_id        = umbrella_ruleset.default_web_policy.id
  name              = "Bypass AzureAD"
  action            = "DO_NOT_DECRYPT"
  rank              = 1
  destination_lists = [umbrella_destination_list.azure_bypass.name]
  enabled           = true
}
```

### Complete Examples

- See [`examples/`](./examples/) directory for various usage examples
- See [`examples/basic/main.tf`](./examples/basic/main.tf) for basic destination lists and tunnels

## Installation

### From Terraform Registry (Recommended)

Once published, you can use the provider directly from the Terraform Registry:

```hcl
terraform {
  required_providers {
    umbrella = {
      source  = "mantisec/umbrella"
      version = "~> 0.2.0"
    }
  }
}
```

### Local Development

1. Build the provider:
   ```bash
   make build
   ```

2. Install locally for testing:
   ```bash
   make install-local
   ```

3. Create a local provider configuration in your Terraform directory:
   ```hcl
   terraform {
     required_providers {
       umbrella = {
         source = "local/mantisec/umbrella"
         version = "0.2.0"
       }
     }
   }
   ```

### Authentication

The provider requires Umbrella API credentials with appropriate permissions:

1. **API Key & Secret**: OAuth2 client credentials for API authentication
2. **Organization ID**: Your Umbrella organization identifier

**Security Best Practices:**
- Store credentials as environment variables or in Terraform Cloud workspace secrets
- Never commit plain credentials to source control
- Use variable files (`.tfvars`) that are excluded from version control

```bash
export TF_VAR_umbrella_api_key="your-api-key"
export TF_VAR_umbrella_api_secret="your-api-secret"
export TF_VAR_umbrella_org_id="your-org-id"
```

## API Endpoints

The provider interacts with the following Umbrella API endpoints:

- **Authentication**: `POST /auth/v2/token`
- **Destination Lists**: `/policies/v2/organizations/{orgId}/destinationlists`
- **IPSec Tunnels**: `/v2/organizations/{orgId}/secureinternetgateway/ipsec/sites`
- **SAML Configuration**: `/v2/organizations/{orgId}/saml`
- **Rulesets**: `/policies/v2/organizations/{orgId}/rulesets`
- **Rules**: `/policies/v2/organizations/{orgId}/rulesets/{rulesetId}/rules`

## Development

### Requirements

- Go 1.21+
- Terraform 1.0+
- Valid Umbrella API credentials

### Building

```bash
go mod tidy
go build -o terraform-provider-umbrella.exe
```

### Testing

```bash
go test ./...
```


## Migration from curl Commands

If you're currently using curl commands in `null_resource` blocks (like in `sso.tf`), you can migrate to proper Terraform resources:

**Before (curl approach):**
```hcl
resource "null_resource" "umbrella_push" {
  provisioner "local-exec" {
    command = <<-BASH
      curl -X PUT \
        -H "Authorization: Bearer ${token}" \
        -H "Content-Type: application/json" \
        "https://api.umbrella.com/v2/organizations/${org_id}/saml" \
        -d '{"metadataUrl":"${metadata_url}","authType":"AzureAD"}'
    BASH
  }
}
```

**After (Terraform resource):**
```hcl
resource "umbrella_saml" "azure_ad" {
  metadata_url = var.azure_metadata_url
  auth_type    = "AzureAD"
}
```

This approach provides:
- ✅ Proper state management
- ✅ Resource drift detection
- ✅ Dependency management
- ✅ Plan/apply workflow
- ✅ Import existing resources

## Publishing & Releases

This provider is configured for automated publishing to the Terraform Registry using GitHub Actions. See [`PUBLISHING.md`](./PUBLISHING.md) for detailed setup instructions.

### Release Process

1. **Create a release** using the provided script:
   ```bash
   # On Linux/macOS
   ./scripts/release.sh 1.0.0
   
   # On Windows
   bash scripts/release.sh 1.0.0
   ```

2. **Manual release** (alternative):
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```

3. **GitHub Actions will automatically**:
   - Build binaries for multiple platforms
   - Sign releases with GPG
   - Publish to GitHub Releases
   - Update Terraform Registry

## Changelog

### v0.2.0
- Added SAML authentication configuration support
- Added SWG ruleset management
- Added policy rule management within rulesets
- Enhanced support for SAML-enabled environments
- Comprehensive examples for replacing curl-based configurations
- **Added automated publishing to Terraform Registry**
- **Added GitHub Actions workflows for CI/CD**
- **Added GoReleaser configuration for multi-platform builds**

### v0.1.0
- Initial release
- Support for destination lists (URL, DOMAIN, CIDR)
- Support for IPSec tunnels
- OAuth2 authentication with token refresh
- Comprehensive CRUD operations for supported resources
