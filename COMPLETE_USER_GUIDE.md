# Terraform Provider Umbrella - Complete User Guide

## Table of Contents

1. [Getting Started](#getting-started)
2. [Installation and Setup](#installation-and-setup)
3. [Basic Usage](#basic-usage)
4. [Advanced Configuration](#advanced-configuration)
5. [Resource Management](#resource-management)
6. [Data Sources](#data-sources)
7. [Authentication and Security](#authentication-and-security)
8. [Best Practices](#best-practices)
9. [Common Workflows](#common-workflows)
10. [Troubleshooting](#troubleshooting)
11. [FAQ](#faq)

## Getting Started

### What is the Terraform Provider Umbrella?

The Terraform Provider Umbrella is an advanced, code generation-driven Terraform provider for managing Cisco Umbrella Secure Web Gateway resources. It automatically generates provider code from OpenAPI specifications, ensuring consistency, quality, and comprehensive coverage of the Umbrella API.

### Key Features

- **Automated Code Generation**: Resources and data sources generated from OpenAPI specs
- **Comprehensive Coverage**: Support for destination lists, tunnels, SAML, policies, and more
- **Production-Ready**: Built-in caching, error handling, and validation
- **Complete Documentation**: Auto-generated documentation with examples
- **Testing Framework**: Comprehensive test coverage for all resources

### System Requirements

- **Terraform**: Version 1.0 or later
- **Go**: Version 1.21 or later (for development)
- **Cisco Umbrella**: Valid API credentials with appropriate permissions
- **Operating System**: Windows, macOS, or Linux

## Installation and Setup

### Option 1: From Terraform Registry (Recommended)

Add the provider to your Terraform configuration:

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

### Option 2: Local Development Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/mantisec/terraform-provider-umbrella.git
   cd terraform-provider-umbrella
   ```

2. **Build the provider**:
   ```bash
   make build
   ```

3. **Install locally**:
   ```bash
   make install-local
   ```

4. **Configure for local use**:
   ```hcl
   terraform {
     required_providers {
       umbrella = {
         source  = "local/mantisec/umbrella"
         version = "0.2.0"
       }
     }
   }
   ```

### Initial Configuration

Configure the provider with your Umbrella API credentials:

```hcl
provider "umbrella" {
  api_key    = var.umbrella_api_key     # Your API key (client ID)
  api_secret = var.umbrella_api_secret  # Your API secret (client secret)
  org_id     = var.umbrella_org_id      # Your organization ID
}
```

### Environment Variables

You can also configure the provider using environment variables:

```bash
export UMBRELLA_API_KEY="your-api-key"
export UMBRELLA_API_SECRET="your-api-secret"
export UMBRELLA_ORG_ID="your-org-id"
```

## Basic Usage

### Creating Your First Resource

Start with a simple destination list:

```hcl
resource "umbrella_destination_list" "blocked_sites" {
  name = "Blocked Websites"
  type = "DOMAIN"
  destinations = [
    "malicious-site.com",
    "phishing-domain.net"
  ]
}
```

### Using Data Sources

Retrieve information about existing resources:

```hcl
data "umbrella_destination_list" "existing" {
  id = "12345"
}

output "destination_count" {
  value = length(data.umbrella_destination_list.existing.destinations)
}
```

### Complete Basic Example

```hcl
terraform {
  required_providers {
    umbrella = {
      source  = "mantisec/umbrella"
      version = "~> 0.2.0"
    }
  }
}

provider "umbrella" {
  api_key    = var.umbrella_api_key
  api_secret = var.umbrella_api_secret
  org_id     = var.umbrella_org_id
}

# Create destination list
resource "umbrella_destination_list" "corporate_blocked" {
  name = "Corporate Blocked Sites"
  type = "DOMAIN"
  destinations = [
    "social-media.com",
    "gaming-site.net",
    "streaming-service.org"
  ]
}

# Create IPSec tunnel
resource "umbrella_tunnel" "main_office" {
  name            = "Main-Office-Tunnel"
  device_ip       = "203.0.113.10"
  pre_shared_key  = var.tunnel_psk
}

# Output tunnel information
output "tunnel_endpoint" {
  value = umbrella_tunnel.main_office.tunnel_endpoint
}
```

## Advanced Configuration

### Schema Overrides and Customization

The provider supports advanced configuration through the code generation system:

```yaml
# tools/generator/config/advanced_config.yaml
schema_overrides:
  resources:
    destination_list:
      schema_transforms:
        - field: "destinations"
          type: "list"
          element_type: "string"
          validation: "min_items=1,max_items=1000"
      
      validation_rules:
        - rule: "destinations_not_empty"
          message: "Destination list must contain at least one destination"
```

### Custom Validation Rules

Configure custom validation for your resources:

```yaml
validation:
  custom_validators:
    ip_address:
      pattern: "^(?:[0-9]{1,3}\\.){3}[0-9]{1,3}$"
      message: "Must be a valid IPv4 address"
    
    domain_name:
      pattern: "^[a-zA-Z0-9]([a-zA-Z0-9\\-]{0,61}[a-zA-Z0-9])?\\.[a-zA-Z]{2,}$"
      message: "Must be a valid domain name"
```

### Performance Optimization

Enable caching and parallel processing:

```yaml
performance:
  caching:
    enabled: true
    cache_dir: ".umbrella_cache"
    ttl: "5m"
  
  parallel_processing:
    enabled: true
    max_workers: 4
```

## Resource Management

### Destination Lists

Destination lists are collections of URLs, domains, or IP addresses used in policies.

#### Domain Destination List

```hcl
resource "umbrella_destination_list" "blocked_domains" {
  name = "Blocked Domains"
  type = "DOMAIN"
  destinations = [
    "malicious-site.com",
    "phishing-domain.net",
    "suspicious-website.org"
  ]
}
```

#### URL Destination List

```hcl
resource "umbrella_destination_list" "allowed_urls" {
  name = "Allowed URLs"
  type = "URL"
  destinations = [
    "https://trusted-api.com/endpoint",
    "https://corporate-portal.example.com",
    "https://secure-service.net/api"
  ]
}
```

#### CIDR Destination List

```hcl
resource "umbrella_destination_list" "internal_networks" {
  name = "Internal Networks"
  type = "CIDR"
  destinations = [
    "10.0.0.0/8",
    "172.16.0.0/12",
    "192.168.0.0/16"
  ]
}
```

### IPSec Tunnels

Configure secure tunnels to Umbrella's Secure Internet Gateway:

```hcl
resource "umbrella_tunnel" "primary_tunnel" {
  name            = "Primary-SIG-Tunnel"
  site_origin_id  = var.site_origin_id
  device_ip       = "203.0.113.10"
  pre_shared_key  = var.tunnel_psk
  local_networks  = ["10.0.0.0/8", "192.168.0.0/16"]
  tunnel_type     = "IPSEC"
}
```

### SAML Configuration

Configure SAML authentication for SSO integration:

```hcl
resource "umbrella_saml" "azure_ad" {
  metadata_url = "https://login.microsoftonline.com/tenant-id/federationmetadata/2007-06/federationmetadata.xml"
  auth_type    = "AzureAD"
}
```

### Policy Management

#### Rulesets

```hcl
resource "umbrella_ruleset" "corporate_policy" {
  name                     = "Corporate Web Policy"
  description              = "Main web filtering policy"
  saml_enabled             = true
  ssl_decryption_enabled   = true
}
```

#### Rules

```hcl
resource "umbrella_rule" "bypass_azure_auth" {
  ruleset_id        = umbrella_ruleset.corporate_policy.id
  name              = "Bypass Azure AD Authentication"
  action            = "DO_NOT_DECRYPT"
  rank              = 1
  destination_lists = [umbrella_destination_list.azure_bypass.name]
  enabled           = true
}
```

## Data Sources

### Retrieving Existing Resources

Use data sources to reference existing Umbrella resources:

```hcl
# Get existing destination list
data "umbrella_destination_list" "existing_blocked" {
  id = "12345"
}

# Use in policy rule
resource "umbrella_rule" "block_existing_list" {
  ruleset_id        = umbrella_ruleset.main.id
  name              = "Block Existing List"
  action            = "BLOCK"
  rank              = 10
  destination_lists = [data.umbrella_destination_list.existing_blocked.name]
}
```

### Threat Intelligence Data Sources

Access Umbrella's threat intelligence data:

```hcl
# Domain categorization
data "umbrella_domain_categorization" "example" {
  domain = "example.com"
}

# Domain risk score
data "umbrella_domain_risk_score" "suspicious" {
  domain = "suspicious-site.com"
}

# Passive DNS lookup
data "umbrella_passive_dns" "investigation" {
  domain = "malware-domain.com"
}
```

## Authentication and Security

### API Credentials

Obtain API credentials from the Umbrella dashboard:

1. Log in to your Umbrella dashboard
2. Navigate to Admin → API Keys
3. Create a new API key with appropriate scopes
4. Note the Client ID (API Key) and Client Secret (API Secret)

### Required Scopes

Ensure your API key has the necessary scopes:

- `policies:read` - Read policy information
- `policies:write` - Create and modify policies
- `deployments:read` - Read deployment information
- `deployments:write` - Manage deployments
- `investigate:read` - Access threat intelligence data

### Secure Credential Management

#### Using Environment Variables

```bash
export TF_VAR_umbrella_api_key="your-api-key"
export TF_VAR_umbrella_api_secret="your-api-secret"
export TF_VAR_umbrella_org_id="your-org-id"
```

#### Using Terraform Cloud/Enterprise

Store credentials as sensitive workspace variables:

- `umbrella_api_key` (sensitive)
- `umbrella_api_secret` (sensitive)
- `umbrella_org_id`

#### Using Variable Files

Create a `terraform.tfvars` file (exclude from version control):

```hcl
umbrella_api_key    = "your-api-key"
umbrella_api_secret = "your-api-secret"
umbrella_org_id     = "your-org-id"
```

## Best Practices

### Resource Organization

#### Use Consistent Naming

```hcl
resource "umbrella_destination_list" "corp_blocked_social" {
  name = "Corporate - Blocked Social Media"
  type = "DOMAIN"
  # ...
}

resource "umbrella_destination_list" "corp_allowed_business" {
  name = "Corporate - Allowed Business Sites"
  type = "URL"
  # ...
}
```

#### Group Related Resources

```hcl
# Security-focused destination lists
resource "umbrella_destination_list" "security_malware" {
  name = "Security - Known Malware Domains"
  type = "DOMAIN"
  destinations = var.malware_domains
}

resource "umbrella_destination_list" "security_phishing" {
  name = "Security - Phishing Sites"
  type = "URL"
  destinations = var.phishing_urls
}

# Create security ruleset
resource "umbrella_ruleset" "security_policy" {
  name        = "Security Policy"
  description = "High-priority security rules"
}

# Block malware domains
resource "umbrella_rule" "block_malware" {
  ruleset_id        = umbrella_ruleset.security_policy.id
  name              = "Block Malware Domains"
  action            = "BLOCK"
  rank              = 1
  destination_lists = [umbrella_destination_list.security_malware.name]
}
```

### State Management

#### Use Remote State

```hcl
terraform {
  backend "s3" {
    bucket = "your-terraform-state"
    key    = "umbrella/terraform.tfstate"
    region = "us-west-2"
  }
}
```

#### Import Existing Resources

```bash
# Import existing destination list
terraform import umbrella_destination_list.existing 12345

# Import existing tunnel
terraform import umbrella_tunnel.existing 67890
```

### Variable Management

#### Use Structured Variables

```hcl
variable "destination_lists" {
  description = "Destination lists configuration"
  type = map(object({
    name         = string
    type         = string
    destinations = list(string)
  }))
  default = {
    blocked_social = {
      name = "Blocked Social Media"
      type = "DOMAIN"
      destinations = [
        "facebook.com",
        "twitter.com",
        "instagram.com"
      ]
    }
  }
}

# Use in resources
resource "umbrella_destination_list" "configured" {
  for_each = var.destination_lists
  
  name         = each.value.name
  type         = each.value.type
  destinations = each.value.destinations
}
```

### Error Handling

#### Use Validation

```hcl
variable "tunnel_device_ip" {
  description = "Device IP address for tunnel"
  type        = string
  
  validation {
    condition = can(regex("^(?:[0-9]{1,3}\\.){3}[0-9]{1,3}$", var.tunnel_device_ip))
    error_message = "Device IP must be a valid IPv4 address."
  }
}
```

#### Handle Dependencies

```hcl
# Ensure destination list exists before creating rule
resource "umbrella_rule" "block_list" {
  depends_on = [umbrella_destination_list.blocked_sites]
  
  ruleset_id        = umbrella_ruleset.main.id
  name              = "Block Sites"
  action            = "BLOCK"
  rank              = 5
  destination_lists = [umbrella_destination_list.blocked_sites.name]
}
```

## Common Workflows

### Workflow 1: Setting Up Basic Web Filtering

1. **Create destination lists**:
   ```hcl
   resource "umbrella_destination_list" "blocked_categories" {
     name = "Blocked Categories"
     type = "DOMAIN"
     destinations = [
       "gambling-site.com",
       "adult-content.net",
       "malware-domain.org"
     ]
   }
   ```

2. **Create policy ruleset**:
   ```hcl
   resource "umbrella_ruleset" "web_filtering" {
     name        = "Web Filtering Policy"
     description = "Basic web filtering rules"
   }
   ```

3. **Add blocking rule**:
   ```hcl
   resource "umbrella_rule" "block_categories" {
     ruleset_id        = umbrella_ruleset.web_filtering.id
     name              = "Block Unwanted Categories"
     action            = "BLOCK"
     rank              = 1
     destination_lists = [umbrella_destination_list.blocked_categories.name]
   }
   ```

### Workflow 2: Configuring Site-to-Site VPN

1. **Create IPSec tunnel**:
   ```hcl
   resource "umbrella_tunnel" "site_tunnel" {
     name            = "Site-to-Site-Tunnel"
     site_origin_id  = var.site_id
     device_ip       = var.firewall_ip
     pre_shared_key  = var.psk
     local_networks  = var.local_subnets
   }
   ```

2. **Configure routing**:
   ```hcl
   output "tunnel_config" {
     value = {
       endpoint    = umbrella_tunnel.site_tunnel.tunnel_endpoint
       psk         = var.psk
       local_nets  = umbrella_tunnel.site_tunnel.local_networks
     }
     sensitive = true
   }
   ```

### Workflow 3: SAML SSO Integration

1. **Configure SAML**:
   ```hcl
   resource "umbrella_saml" "corporate_sso" {
     metadata_url = var.saml_metadata_url
     auth_type    = "AzureAD"
   }
   ```

2. **Create bypass rules for authentication**:
   ```hcl
   resource "umbrella_destination_list" "auth_bypass" {
     name = "Authentication Bypass"
     type = "URL"
     destinations = [
       "https://login.microsoftonline.com",
       "https://msauth.net"
     ]
   }
   
   resource "umbrella_rule" "bypass_auth" {
     ruleset_id        = umbrella_ruleset.main.id
     name              = "Bypass Authentication"
     action            = "DO_NOT_DECRYPT"
     rank              = 1
     destination_lists = [umbrella_destination_list.auth_bypass.name]
   }
   ```

### Workflow 4: Threat Intelligence Integration

1. **Query domain reputation**:
   ```hcl
   data "umbrella_domain_risk_score" "suspicious" {
     domain = "suspicious-domain.com"
   }
   ```

2. **Create dynamic blocking based on risk**:
   ```hcl
   locals {
     high_risk_domains = [
       for domain in var.domains_to_check :
       domain if data.umbrella_domain_risk_score[domain].risk_score > 70
     ]
   }
   
   resource "umbrella_destination_list" "high_risk" {
     name         = "High Risk Domains"
     type         = "DOMAIN"
     destinations = local.high_risk_domains
   }
   ```

## Troubleshooting

### Common Issues and Solutions

#### Issue: Authentication Failures

**Symptoms**: `401 Unauthorized` errors

**Solutions**:
1. Verify API credentials are correct
2. Check that API key has required scopes
3. Ensure organization ID is accurate
4. Verify credentials aren't expired

```bash
# Test authentication
curl -X POST https://api.umbrella.com/auth/v2/token \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "grant_type=client_credentials" \
  -u "your-api-key:your-api-secret"
```

#### Issue: Resource Not Found

**Symptoms**: `404 Not Found` errors during read operations

**Solutions**:
1. Verify resource ID is correct
2. Check if resource was deleted outside Terraform
3. Refresh Terraform state
4. Import existing resource if needed

```bash
# Refresh state
terraform refresh

# Import existing resource
terraform import umbrella_destination_list.example 12345
```

#### Issue: Rate Limiting

**Symptoms**: `429 Too Many Requests` errors

**Solutions**:
1. Reduce parallel operations
2. Add delays between operations
3. Enable caching to reduce API calls
4. Contact Umbrella support for rate limit increase

```hcl
# Configure rate limiting
provider "umbrella" {
  api_key    = var.umbrella_api_key
  api_secret = var.umbrella_api_secret
  org_id     = var.umbrella_org_id
  
  # Add rate limiting configuration
  rate_limit = {
    requests_per_second = 5
    burst_size         = 10
  }
}
```

#### Issue: Validation Errors

**Symptoms**: Schema validation failures

**Solutions**:
1. Check field types and formats
2. Verify required fields are provided
3. Validate field constraints (min/max values)
4. Review API documentation for field requirements

```hcl
# Example validation fix
resource "umbrella_destination_list" "fixed" {
  name = "Valid Name"  # Must not be empty
  type = "DOMAIN"      # Must be DOMAIN, URL, or CIDR
  destinations = [     # Must contain at least one item
    "example.com"
  ]
}
```

### Debug Mode

Enable debug logging for detailed troubleshooting:

```bash
export TF_LOG=DEBUG
export TF_LOG_PATH=terraform.log
terraform apply
```

### Getting Help

1. **Check Documentation**: Review resource documentation in `docs/resources/`
2. **GitHub Issues**: Search existing issues or create new ones
3. **Community Support**: Use GitHub Discussions for questions
4. **Debug Logs**: Include debug logs when reporting issues

## FAQ

### General Questions

**Q: What APIs does the provider support?**
A: The provider supports Umbrella's Policies API, Deployments API, Authentication API, and Investigate API for threat intelligence.

**Q: Can I use this provider with Terraform Cloud?**
A: Yes, the provider is fully compatible with Terraform Cloud and Enterprise.

**Q: How often is the provider updated?**
A: The provider is automatically updated when Umbrella's APIs change, thanks to the code generation system.

### Authentication Questions

**Q: What permissions do I need for my API key?**
A: Your API key needs appropriate scopes based on the resources you want to manage. See the [Authentication and Security](#authentication-and-security) section for details.

**Q: Can I use the same API key for multiple Terraform configurations?**
A: Yes, but ensure the API key has sufficient permissions for all resources across configurations.

### Resource Management Questions

**Q: Can I import existing Umbrella resources?**
A: Yes, most resources support Terraform import. Use `terraform import resource_type.name resource_id`.

**Q: How do I handle resource dependencies?**
A: Use Terraform's `depends_on` argument or reference attributes to create implicit dependencies.

**Q: Can I manage multiple organizations?**
A: You can use provider aliases to manage multiple organizations:

```hcl
provider "umbrella" {
  alias      = "org1"
  api_key    = var.org1_api_key
  api_secret = var.org1_api_secret
  org_id     = var.org1_id
}

provider "umbrella" {
  alias      = "org2"
  api_key    = var.org2_api_key
  api_secret = var.org2_api_secret
  org_id     = var.org2_id
}
```

### Performance Questions

**Q: How can I improve performance for large configurations?**
A: Enable caching, use parallel processing, and consider breaking large configurations into smaller modules.

**Q: Does the provider cache API responses?**
A: Yes, the provider includes intelligent caching for read operations with configurable TTL.

### Troubleshooting Questions

**Q: Why am I getting rate limit errors?**
A: Umbrella has API rate limits. Reduce parallelism, enable caching, or contact support for higher limits.

**Q: How do I debug API issues?**
A: Enable Terraform debug logging with `TF_LOG=DEBUG` and review the detailed API request/response logs.

**Q: What should I do if a resource gets out of sync?**
A: Run `terraform refresh` to update state, or use `terraform import` to re-import the resource.

This comprehensive user guide provides everything needed to effectively use the Terraform Provider Umbrella, from basic setup to advanced configurations and troubleshooting.