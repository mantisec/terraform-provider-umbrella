---
page_title: "Umbrella Provider"
subcategory: ""
description: |-
  The Umbrella provider is used to interact with Cisco Umbrella resources through the REST API. Manage networks, sites, users, destination lists, and internal networks with full Terraform lifecycle support.
---

# Umbrella Provider

The Umbrella provider is used to interact with [Cisco Umbrella](https://umbrella.cisco.com/) resources through the REST API. The provider enables Infrastructure as Code management of your Umbrella deployment including networks, sites, users, destination lists, and internal network configurations.

## Features

- **Complete Resource Management**: Create, read, update, and delete Umbrella resources
- **OAuth2 Authentication**: Automatic token management with refresh capabilities
- **State Management**: Full Terraform state tracking and drift detection
- **Import Support**: Import existing Umbrella resources into Terraform
- **Validation**: Comprehensive input validation and error handling

## Quick Start

```terraform
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
```

## Authentication

The Umbrella provider requires OAuth2 client credentials for authentication. You can obtain these credentials from the Umbrella dashboard under **Admin > API Keys**.

### Required Credentials

- **API Key**: Your Umbrella API key (client ID)
- **API Secret**: Your Umbrella API secret (client secret)
- **Organization ID**: Your Umbrella organization identifier

### Configuration Methods

#### Method 1: Provider Configuration (Recommended)

```terraform
provider "umbrella" {
  api_key    = var.umbrella_api_key
  api_secret = var.umbrella_api_secret
  org_id     = var.umbrella_org_id
}
```

#### Method 2: Environment Variables

```bash
export TF_VAR_umbrella_api_key="your-api-key"
export TF_VAR_umbrella_api_secret="your-api-secret"
export TF_VAR_umbrella_org_id="your-org-id"
```

#### Method 3: Terraform Variables File

Create a `terraform.tfvars` file (excluded from version control):

```hcl
umbrella_api_key    = "your-api-key"
umbrella_api_secret = "your-api-secret"
umbrella_org_id     = "your-org-id"
```

## Schema

### Required

- `api_key` (String, Sensitive) - Umbrella API key (client ID). Can also be set via `UMBRELLA_API_KEY` environment variable.
- `api_secret` (String, Sensitive) - Umbrella API secret (client secret). Can also be set via `UMBRELLA_API_SECRET` environment variable.
- `org_id` (String) - Umbrella organization ID. Can also be set via `UMBRELLA_ORG_ID` environment variable.

## Resources

### Core Resources

- [`umbrella_sites`](resources/sites.md) - Manages organizational sites and locations
- [`umbrella_networks`](resources/networks.md) - Manages network definitions and IP ranges
- [`umbrella_internalnetworks`](resources/internal_networks.md) - Manages internal network configurations
- [`umbrella_users`](resources/users.md) - Manages user accounts and permissions
- [`umbrella_destination_list`](resources/destination_list.md) - Manages destination lists for policy enforcement

### Additional Resources

- [`umbrella_tunnel`](resources/tunnel.md) - Manages IPSec tunnels for Secure Internet Gateway

## Data Sources

### Core Data Sources

- [`umbrella_sites`](data-sources/sites.md) - Retrieve information about sites
- [`umbrella_networks`](data-sources/networks.md) - Retrieve information about networks
- [`umbrella_internalnetworks`](data-sources/internal_networks.md) - Retrieve information about internal networks
- [`umbrella_users`](data-sources/users.md) - Retrieve information about users
- [`umbrella_destination_list`](data-sources/destination_list.md) - Retrieve information about destination lists

## Guides

### Getting Started
- [Getting Started Guide](guides/getting-started.md) - Complete setup and first steps with the provider
- [Migration Guide](guides/migration.md) - Migrate from manual or script-based management
- [Troubleshooting Guide](guides/troubleshooting.md) - Comprehensive troubleshooting and debugging
- [API Reference](guides/api-reference.md) - Complete API endpoint documentation and reference

### Quick Start

1. **Install the Provider**
   ```terraform
   terraform {
     required_providers {
       umbrella = {
         source  = "mantisec/umbrella"
         version = "~> 1.0"
       }
     }
   }
   ```

2. **Configure Authentication**
   ```terraform
   provider "umbrella" {
     api_key    = var.umbrella_api_key
     api_secret = var.umbrella_api_secret
     org_id     = var.umbrella_org_id
   }
   ```

3. **Create Your First Resource**
   ```terraform
   resource "umbrella_sites" "example" {
     name = "My First Site"
   }
   ```

4. **Initialize and Apply**
   ```bash
   terraform init
   terraform plan
   terraform apply
   ```

For detailed setup instructions, see the [Getting Started Guide](guides/getting-started.md).

## Examples

### Basic Usage
- [Basic Examples](https://github.com/mantisec/terraform-provider-umbrella/tree/main/examples/basic) - Simple resource configurations
- [Complete Examples](https://github.com/mantisec/terraform-provider-umbrella/tree/main/examples/complete) - Comprehensive multi-resource deployment

### Specific Use Cases
- [User Management](https://github.com/mantisec/terraform-provider-umbrella/tree/main/examples/users) - User account management
- [Site Management](https://github.com/mantisec/terraform-provider-umbrella/tree/main/examples/sites) - Site and location management
- [Internal Networks](https://github.com/mantisec/terraform-provider-umbrella/tree/main/examples/internal_networks) - Network topology examples

## API Endpoints

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

## Security Best Practices

### Credential Management
- **Never commit credentials to source control**
- Store credentials as environment variables or in secure secret management systems
- Use Terraform Cloud/Enterprise workspace variables for team environments
- Regularly rotate API credentials

### Access Control
- Use least-privilege access for API credentials
- Create separate API keys for different environments (dev, staging, prod)
- Monitor API key usage through Umbrella dashboard

### Network Security
- Ensure your Terraform execution environment can reach `api.umbrella.com`
- Configure appropriate firewall rules for API access
- Use secure networks for Terraform operations

## Troubleshooting

### Common Issues

#### Authentication Errors
```
Error: Unable to authenticate
```
- Verify your API credentials are correct
- Ensure the organization ID matches your Umbrella account
- Check that your API key has the necessary permissions

#### Resource Not Found
```
Error: resource not found: 404
```
- Verify the resource ID exists in your Umbrella organization
- Check that you have permissions to access the resource

#### Rate Limiting
```
Error: rate limit exceeded: 429
```
- The provider automatically handles rate limiting with exponential backoff
- If issues persist, contact Cisco support to review rate limits

### Getting Help

- Review the [examples](https://github.com/mantisec/terraform-provider-umbrella/tree/main/examples) directory
- Check the individual resource documentation for detailed usage
- Open an issue on the [provider repository](https://github.com/mantisec/terraform-provider-umbrella) for bugs or feature requests

## Migration from Manual Configuration

If you're currently managing Umbrella resources manually or through scripts, you can import existing resources:

```bash
# Import an existing site
terraform import umbrella_sites.example 12345

# Import an existing destination list
terraform import umbrella_destination_list.example 67890
```

See individual resource documentation for specific import instructions.