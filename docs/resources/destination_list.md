---
page_title: "umbrella_destination_list Resource - terraform-provider-umbrella"
subcategory: "Policies"
description: |-
  Manages destination lists in Cisco Umbrella. Destination lists allow you to create custom lists of domains, IPs, or URLs that can be used in policies to allow or block access.
---

# umbrella_destination_list (Resource)

Manages destination lists in Cisco Umbrella. Destination lists allow you to create custom lists of domains, IP addresses, CIDR blocks, or URLs that can be used in policies to allow or block access to specific destinations.

Destination lists are fundamental building blocks for Umbrella policies, enabling you to group related destinations together for consistent policy enforcement across your organization.

## Example Usage

### Block List for Malicious Domains

```terraform
resource "umbrella_destination_list" "blocked_domains" {
  name   = "Blocked Malicious Domains"
  access = "block"
  destinations = [
    "malicious-site.com",
    "phishing-domain.net",
    "suspicious-website.org"
  ]
}
```

### Allow List for Trusted Services

```terraform
resource "umbrella_destination_list" "trusted_services" {
  name   = "Trusted Business Services"
  access = "allow"
  destinations = [
    "trusted-partner.com",
    "business-app.example.com",
    "secure-api.company.net"
  ]
}
```

### IP Address and CIDR Block List

```terraform
resource "umbrella_destination_list" "internal_networks" {
  name   = "Internal Network Ranges"
  access = "allow"
  destinations = [
    "10.0.0.0/8",
    "172.16.0.0/12",
    "192.168.1.100",
    "203.0.113.0/24"
  ]
}
```

### URL-Specific Destinations

```terraform
resource "umbrella_destination_list" "specific_urls" {
  name   = "Specific URL Endpoints"
  access = "allow"
  destinations = [
    "https://api.example.com/v1/webhook",
    "https://secure.payment-processor.com/callback",
    "https://cdn.trusted-site.com/assets"
  ]
}
```

### Global Destination List

```terraform
resource "umbrella_destination_list" "global_blocklist" {
  name      = "Global Security Blocklist"
  access    = "block"
  is_global = true
  destinations = [
    "known-malware-c2.com",
    "botnet-command.net",
    "ransomware-payment.onion"
  ]
}
```

## Schema

### Required

- `name` (String) The name of the destination list. Must be unique within your organization and between 1-255 characters.
- `access` (String) Access type for the destination list. Valid values: `allow`, `block`.

### Optional

- `destinations` (Set of String) Set of destination entries. Can include domains, IP addresses, CIDR blocks, or URLs. Each destination must be valid according to its type.
- `is_global` (Boolean) Whether this is a global destination list shared across all policies in the organization. Defaults to `false`.

### Read-Only

- `id` (String) The unique identifier for this destination list.
- `created_at` (String) Timestamp when the destination list was created (ISO 8601 format).
- `modified_at` (String) Timestamp when the destination list was last modified (ISO 8601 format).

## Import

Destination lists can be imported using their ID:

```shell
terraform import umbrella_destination_list.example 12345
```

## Validation Rules

### Name Validation
- Must be between 1 and 255 characters
- Must be unique within your organization
- Cannot contain special characters that interfere with policy processing

### Access Type Validation
- Must be either `allow` or `block`
- Case-sensitive

### Destination Validation
- **Domains**: Must be valid domain names (e.g., `example.com`, `sub.domain.org`)
- **IP Addresses**: Must be valid IPv4 addresses (e.g., `192.168.1.1`)
- **CIDR Blocks**: Must be valid CIDR notation (e.g., `10.0.0.0/8`, `192.168.1.0/24`)
- **URLs**: Must be valid URLs with protocol (e.g., `https://example.com/path`)

## API Reference

This resource uses the following Cisco Umbrella API endpoints:

- **Create**: `POST /policies/v2/organizations/{orgId}/destinationlists`
- **Read**: `GET /policies/v2/organizations/{orgId}/destinationlists/{id}`
- **Update**: `PUT /policies/v2/organizations/{orgId}/destinationlists/{id}`
- **Delete**: `DELETE /policies/v2/organizations/{orgId}/destinationlists/{id}`

## Usage Patterns

### Dynamic Destination Management

```terraform
# Use local values for easier management
locals {
  blocked_domains = [
    "malicious-site.com",
    "phishing-domain.net",
    "suspicious-website.org"
  ]
  
  trusted_domains = [
    "corporate-app.company.com",
    "partner-api.trusted.net"
  ]
}

resource "umbrella_destination_list" "dynamic_blocklist" {
  name         = "Dynamic Block List"
  access       = "block"
  destinations = toset(local.blocked_domains)
}

resource "umbrella_destination_list" "dynamic_allowlist" {
  name         = "Dynamic Allow List"
  access       = "allow"
  destinations = toset(local.trusted_domains)
}
```

### Environment-Specific Lists

```terraform
variable "environment" {
  description = "Environment name"
  type        = string
  default     = "production"
}

resource "umbrella_destination_list" "env_specific" {
  name   = "${title(var.environment)} Environment Destinations"
  access = "allow"
  destinations = var.environment == "production" ? [
    "prod-api.company.com",
    "prod-cdn.company.com"
  ] : [
    "dev-api.company.com",
    "staging-api.company.com"
  ]
}
```

## Best Practices

### Organization and Naming
- Use descriptive, consistent naming conventions
- Group related destinations logically
- Consider using prefixes to indicate purpose (e.g., `BLOCK-`, `ALLOW-`)

### Destination Management
- Regularly review and update destination lists
- Use CIDR blocks for IP ranges rather than individual IPs when possible
- Document the purpose of each destination list

### Security Considerations
- Regularly audit allow lists to prevent security gaps
- Monitor block list effectiveness through Umbrella reporting
- Consider using global lists for organization-wide policies

## Troubleshooting

### Common Errors

**Invalid destination format**
```
Error: destination "invalid-domain" is not a valid domain, IP, CIDR, or URL
```
Ensure all destinations follow the correct format for their type.

**Duplicate destination list name**
```
Error: destination list name already exists
```
Destination list names must be unique within your organization.

**Permission denied**
```
Error: insufficient permissions: 403 Forbidden
```
Ensure your API credentials have the `policies:write` scope for create/update/delete operations.

### Performance Considerations

- Large destination lists (>1000 entries) may impact policy processing performance
- Consider splitting very large lists into smaller, more focused lists
- Monitor policy evaluation times through Umbrella dashboard

## Related Resources

- Use with policy rules to enforce access controls
- Combine with [`umbrella_sites`](sites.md) for location-specific policies
- Reference in security policies for comprehensive protection

## Notes

- Changes to destination lists may take a few minutes to propagate across the Umbrella network
- Global destination lists are shared across all policies in the organization
- Destination lists support mixed content types (domains, IPs, URLs) in a single list
- The order of destinations in the list does not affect policy evaluation