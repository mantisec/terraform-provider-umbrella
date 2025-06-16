# Terraform Provider for Cisco Umbrella

[![Terraform Registry](https://img.shields.io/badge/terraform-registry-blue.svg)](https://registry.terraform.io/providers/mantisec/umbrella/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/mantisec/terraform-provider-umbrella)](https://goreportcard.com/report/github.com/mantisec/terraform-provider-umbrella)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

The Terraform Provider for Cisco Umbrella enables Infrastructure as Code management of your Cisco Umbrella resources through the REST API. Manage sites, networks, users, destination lists, and internal networks with full Terraform lifecycle support.

> **Status**: Production Ready - Stable API with comprehensive resource coverage
>
> **Maintainer**: Grant Lawton | Mantisec | [grant.lawton@mantisec.com.au](mailto:grant.lawton@mantisec.com.au)

## Features

- **Complete Resource Management**: Full CRUD operations for all supported resources
- **Sites Management**: Create and manage organizational sites and locations
- **Network Management**: Define and manage network configurations and IP ranges
- **User Management**: Manage user accounts, roles, and permissions
- **Destination Lists**: Create and manage URL, domain, and CIDR destination lists for policy enforcement
- **Internal Networks**: Configure internal network topologies and associations
- **IPSec Tunnels**: Create and manage IPSec tunnels for Secure Internet Gateway (SIG)
- **OAuth2 Authentication**: Automatic token management with refresh capabilities
- **State Management**: Full Terraform state tracking and drift detection
- **Import Support**: Import existing Umbrella resources into Terraform management

## Supported Resources

### Core Resources

| Resource | Description | Documentation |
|----------|-------------|---------------|
| [`umbrella_sites`](docs/resources/sites.md) | Manages organizational sites and locations | [View Docs](docs/resources/sites.md) |
| [`umbrella_networks`](docs/resources/networks.md) | Manages network definitions and IP ranges | [View Docs](docs/resources/networks.md) |
| [`umbrella_internalnetworks`](docs/resources/internal_networks.md) | Manages internal network configurations | [View Docs](docs/resources/internal_networks.md) |
| [`umbrella_users`](docs/resources/users.md) | Manages user accounts and permissions | [View Docs](docs/resources/users.md) |
| [`umbrella_destination_list`](docs/resources/destination_list.md) | Manages destination lists for policy enforcement | [View Docs](docs/resources/destination_list.md) |

### Additional Resources

| Resource | Description | Documentation |
|----------|-------------|---------------|
| [`umbrella_tunnel`](docs/resources/tunnel.md) | Manages IPSec tunnels for Secure Internet Gateway | [View Docs](docs/resources/tunnel.md) |

### Data Sources

| Data Source | Description | Documentation |
|-------------|-------------|---------------|
| [`umbrella_sites`](docs/data-sources/sites.md) | Retrieve information about sites | [View Docs](docs/data-sources/sites.md) |
| [`umbrella_networks`](docs/data-sources/networks.md) | Retrieve information about networks | [View Docs](docs/data-sources/networks.md) |
| [`umbrella_users`](docs/data-sources/users.md) | Retrieve information about users | [View Docs](docs/data-sources/users.md) |
| [`umbrella_destination_list`](docs/data-sources/destination_list.md) | Retrieve information about destination lists | [View Docs](docs/data-sources/destination_list.md) |

## Quick Start

### 1. Install the Provider

```hcl
terraform {
  required_providers {
    umbrella = {
      source  = "mantisec/umbrella"
      version = "~> 1.0"
    }
  }
}
```

### 2. Configure Authentication

```hcl
provider "umbrella" {
  api_key    = var.umbrella_api_key
  api_secret = var.umbrella_api_secret
  org_id     = var.umbrella_org_id
}
```

### 3. Create Your First Resources

```hcl
# Create a site
resource "umbrella_sites" "main_office" {
  name = "Main Office"
}

# Create a destination list
resource "umbrella_destination_list" "blocked_domains" {
  name   = "Blocked Domains"
  access = "block"
  destinations = [
    "malicious-site.com",
    "phishing-domain.net"
  ]
}

# Create an internal network
resource "umbrella_internalnetworks" "office_network" {
  name          = "Office Network"
  ip_address    = "192.168.1.0"
  prefix_length = 24
  site_id       = umbrella_sites.main_office.site_id
}
```

## Documentation

### Getting Started
- [Getting Started Guide](docs/guides/getting-started.md) - Complete setup and first steps
- [Migration Guide](docs/guides/migration.md) - Migrate from manual or script-based management
- [Provider Documentation](docs/index.md) - Complete provider reference

### Examples
- [Basic Examples](examples/basic/) - Simple resource configurations
- [Complete Examples](examples/complete/) - Advanced configurations
- [User Management](examples/users/) - User account management
- [Site Management](examples/sites/) - Site and location management
- [Internal Networks](examples/internal_networks/) - Network topology examples

## Installation

### From Terraform Registry (Recommended)

```hcl
terraform {
  required_providers {
    umbrella = {
      source  = "mantisec/umbrella"
      version = "~> 1.0"
    }
  }
}
```

### Authentication Setup

1. **Obtain API Credentials** from your Umbrella dashboard:
   - Navigate to **Admin** → **API Keys**
   - Create a new API key with appropriate scopes
   - Note the API Key, API Secret, and Organization ID

2. **Configure Provider**:
   ```hcl
   provider "umbrella" {
     api_key    = var.umbrella_api_key
     api_secret = var.umbrella_api_secret
     org_id     = var.umbrella_org_id
   }
   ```

3. **Set Environment Variables** (recommended):
   ```bash
   export TF_VAR_umbrella_api_key="your-api-key"
   export TF_VAR_umbrella_api_secret="your-api-secret"
   export TF_VAR_umbrella_org_id="your-org-id"
   ```

## API Coverage

The provider interacts with the following Umbrella API endpoints:

### Authentication
- `POST /auth/v2/token` - OAuth2 token management

### Deployments
- `/deployments/v2/sites` - Site management
- `/deployments/v2/networks` - Network management
- `/deployments/v2/internalnetworks` - Internal network management

### Policies
- `/policies/v2/organizations/{orgId}/destinationlists` - Destination list management

### Admin
- `/admin/v2/organizations/{orgId}/users` - User management

### Secure Internet Gateway
- `/v2/organizations/{orgId}/secureinternetgateway/ipsec/sites` - IPSec tunnel management

## Development

### Requirements
- Go 1.21+
- Terraform 1.0+
- Valid Umbrella API credentials

### Building from Source

```bash
# Clone the repository
git clone https://github.com/mantisec/terraform-provider-umbrella.git
cd terraform-provider-umbrella

# Build the provider
go mod tidy
go build -o terraform-provider-umbrella

# Install locally for testing
make install-local
```

### Testing

```bash
# Run unit tests
go test ./...

# Run acceptance tests (requires valid API credentials)
TF_ACC=1 go test ./internal/provider/tests/
```

### Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## Migration from Other Tools

### From Manual Management
If you're currently managing Umbrella resources through the web dashboard, see our [Migration Guide](docs/guides/migration.md) for step-by-step instructions.

### From curl/API Scripts
Replace your curl-based scripts with proper Terraform resources:

**Before:**
```bash
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  "https://api.umbrella.com/policies/v2/organizations/$ORG_ID/destinationlists" \
  -d '{"name":"Blocked Sites","access":"block","destinations":["malicious.com"]}'
```

**After:**
```hcl
resource "umbrella_destination_list" "blocked_sites" {
  name   = "Blocked Sites"
  access = "block"
  destinations = ["malicious.com"]
}
```

Benefits of Terraform approach:
- ✅ Proper state management
- ✅ Resource drift detection
- ✅ Dependency management
- ✅ Plan/apply workflow
- ✅ Import existing resources

## Support and Community

### Getting Help
- 📖 [Documentation](docs/) - Comprehensive guides and references
- 💬 [GitHub Issues](https://github.com/mantisec/terraform-provider-umbrella/issues) - Bug reports and feature requests
- 📧 [Email Support](mailto:grant.lawton@mantisec.com.au) - Direct support from maintainer

### Contributing
We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details on:
- Code style and standards
- Testing requirements
- Pull request process
- Issue reporting guidelines

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Cisco Umbrella team for providing comprehensive API documentation
- HashiCorp for the excellent Terraform Plugin Framework
- The Terraform community for best practices and guidance

---

**Maintained by**: [Mantisec](https://mantisec.com.au) | **Author**: Grant Lawton
