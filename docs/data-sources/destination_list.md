---
page_title: "umbrella_destination_list Data Source - terraform-provider-umbrella"
subcategory: "Policies"
description: |-
  Use this data source to retrieve information about destination lists in your Cisco Umbrella organization.
---

# umbrella_destination_list (Data Source)

Use this data source to retrieve information about destination lists in your Cisco Umbrella organization. Destination lists contain collections of domains, IP addresses, CIDR blocks, or URLs that are used in policies to allow or block access to specific destinations.

## Example Usage

### Get a Specific Destination List by Name

```terraform
data "umbrella_destination_list" "blocked_domains" {
  name = "Blocked Malicious Domains"
}

# Use the destination list data in policy configurations
output "blocklist_info" {
  value = {
    id           = data.umbrella_destination_list.blocked_domains.id
    access       = data.umbrella_destination_list.blocked_domains.access
    destinations = data.umbrella_destination_list.blocked_domains.destinations
    is_global    = data.umbrella_destination_list.blocked_domains.is_global
  }
}
```

### Get a Destination List by ID

```terraform
data "umbrella_destination_list" "example" {
  id = "12345"
}

# Check if list contains specific destinations
locals {
  contains_malicious_domain = contains(
    data.umbrella_destination_list.example.destinations,
    "malicious-site.com"
  )
}
```

### Get All Destination Lists

```terraform
data "umbrella_destination_list" "all" {}

# Filter lists by access type
locals {
  block_lists = [
    for list in data.umbrella_destination_list.all.destination_lists : list
    if list.access == "block"
  ]
  
  allow_lists = [
    for list in data.umbrella_destination_list.all.destination_lists : list
    if list.access == "allow"
  ]
  
  global_lists = [
    for list in data.umbrella_destination_list.all.destination_lists : list
    if list.is_global == true
  ]
}

output "destination_list_summary" {
  value = {
    total_lists  = length(data.umbrella_destination_list.all.destination_lists)
    block_lists  = length(local.block_lists)
    allow_lists  = length(local.allow_lists)
    global_lists = length(local.global_lists)
  }
}
```

## Schema

### Optional

- `id` (String) The unique identifier for a specific destination list. If provided, returns information for this list only.
- `name` (String) The name of a specific destination list. If provided, returns information for the list with this name.

### Read-Only

When querying a specific destination list (by `id` or `name`):

- `access` (String) Access type for the destination list (allow/block)
- `destinations` (Set of String) Set of destination entries in the list
- `is_global` (Boolean) Whether this is a global destination list
- `created_at` (String) Timestamp when the destination list was created (ISO 8601 format)
- `modified_at` (String) Timestamp when the destination list was last modified (ISO 8601 format)

When querying all destination lists (no filters):

- `destination_lists` (List of Object) List of all destination lists with the following attributes:
  - `id` (String) The unique identifier for the destination list
  - `name` (String) The name of the destination list
  - `access` (String) Access type (allow/block)
  - `destinations` (Set of String) Set of destination entries
  - `is_global` (Boolean) Whether this is a global destination list
  - `created_at` (String) Creation timestamp
  - `modified_at` (String) Last modification timestamp

## Usage Examples

### Policy Configuration Reference

```terraform
# Get existing destination lists for policy reference
data "umbrella_destination_list" "security_blocklist" {
  name = "Security Block List"
}

data "umbrella_destination_list" "trusted_domains" {
  name = "Trusted Business Domains"
}

# Use in policy rules (example for when policy resources are available)
locals {
  security_policy_config = {
    block_destinations = data.umbrella_destination_list.security_blocklist.destinations
    allow_destinations = data.umbrella_destination_list.trusted_domains.destinations
  }
}
```

### Destination List Validation and Auditing

```terraform
data "umbrella_destination_list" "all" {}

locals {
  # Audit destination list configurations
  list_audit = {
    for list in data.umbrella_destination_list.all.destination_lists : list.name => {
      access_type        = list.access
      destination_count  = length(list.destinations)
      is_global         = list.is_global
      last_modified     = list.modified_at
      
      # Check for common security domains
      has_security_domains = length([
        for dest in list.destinations : dest
        if can(regex("(malware|phishing|suspicious)", dest))
      ]) > 0
      
      # Check for internal domains
      has_internal_domains = length([
        for dest in list.destinations : dest
        if can(regex("\\.(local|internal|corp)$", dest))
      ]) > 0
    }
  }
}

output "destination_list_audit" {
  value = local.list_audit
}
```

### Destination Overlap Analysis

```terraform
data "umbrella_destination_list" "all" {}

locals {
  # Find overlapping destinations between lists
  all_destinations = flatten([
    for list in data.umbrella_destination_list.all.destination_lists : [
      for dest in list.destinations : {
        destination = dest
        list_name   = list.name
        access_type = list.access
      }
    ]
  ])
  
  # Group destinations by value to find duplicates
  destination_groups = {
    for dest in local.all_destinations : dest.destination => dest...
  }
  
  # Find destinations that appear in multiple lists
  overlapping_destinations = {
    for dest, occurrences in local.destination_groups : dest => occurrences
    if length(occurrences) > 1
  }
}

output "destination_overlap_analysis" {
  value = {
    total_unique_destinations = length(keys(local.destination_groups))
    overlapping_destinations  = local.overlapping_destinations
    overlap_count            = length(local.overlapping_destinations)
  }
}
```

### Compliance and Security Monitoring

```terraform
data "umbrella_destination_list" "all" {}

locals {
  # Security compliance checks
  security_analysis = {
    # Lists without destinations (potentially misconfigured)
    empty_lists = [
      for list in data.umbrella_destination_list.all.destination_lists : list.name
      if length(list.destinations) == 0
    ]
    
    # Block lists (should be monitored for effectiveness)
    block_lists_summary = {
      for list in data.umbrella_destination_list.all.destination_lists : list.name => {
        destination_count = length(list.destinations)
        is_global        = list.is_global
        last_modified    = list.modified_at
      }
      if list.access == "block"
    }
    
    # Global lists (organization-wide impact)
    global_lists = [
      for list in data.umbrella_destination_list.all.destination_lists : {
        name             = list.name
        access           = list.access
        destination_count = length(list.destinations)
      }
      if list.is_global == true
    ]
  }
}

output "security_compliance_report" {
  value = local.security_analysis
}
```

### Dynamic List Management

```terraform
# Get current destination lists
data "umbrella_destination_list" "all" {}

locals {
  # Required destination lists for compliance
  required_lists = {
    "Security Block List" = "block"
    "Trusted Partners"    = "allow"
    "Internal Resources"  = "allow"
  }
  
  existing_list_names = [
    for list in data.umbrella_destination_list.all.destination_lists : list.name
  ]
  
  missing_lists = {
    for name, access in local.required_lists : name => access
    if !contains(local.existing_list_names, name)
  }
}

# Create missing required lists
resource "umbrella_destination_list" "required" {
  for_each = local.missing_lists
  
  name         = each.key
  access       = each.value
  destinations = []  # Start with empty list
}

output "list_management_status" {
  value = {
    required_lists = keys(local.required_lists)
    existing_lists = local.existing_list_names
    missing_lists  = keys(local.missing_lists)
    created_lists  = keys(umbrella_destination_list.required)
  }
}
```

## API Reference

This data source uses the following Cisco Umbrella API endpoints:

- **Get Destination List**: `GET /policies/v2/organizations/{orgId}/destinationlists/{id}` (when using `id`)
- **List Destination Lists**: `GET /policies/v2/organizations/{orgId}/destinationlists` (when using `name` or no filters)

## Destination Types

Destination lists can contain various types of destinations:

### Domain Names
- `example.com`
- `subdomain.example.org`
- `*.wildcard-domain.net`

### IP Addresses
- `192.168.1.1`
- `203.0.113.10`

### CIDR Blocks
- `10.0.0.0/8`
- `192.168.1.0/24`
- `172.16.0.0/12`

### URLs
- `https://example.com/specific/path`
- `http://api.service.com/endpoint`

## Notes

- **Access Types**: Lists can be either `allow` (permit access) or `block` (deny access)
- **Global Lists**: Global destination lists are shared across all policies in the organization
- **Case Sensitivity**: Destination list names are case-sensitive when filtering
- **Propagation**: Changes to destination lists may take a few minutes to propagate across the Umbrella network
- **Mixed Content**: A single list can contain different types of destinations (domains, IPs, URLs)

## Troubleshooting

### Common Errors

**Destination list not found**
```
Error: destination list not found with name "My List"
```
Verify the destination list name exists and matches exactly (case-sensitive).

**Permission denied**
```
Error: insufficient permissions: 403 Forbidden
```
Ensure your API credentials have the `policies:read` scope.

**Empty destination list**
```
Warning: destination list contains no destinations
```
This may indicate a misconfigured list that should be reviewed.

### Best Practices

1. **Regular Audits**: Periodically review destination lists for accuracy and relevance
2. **Naming Conventions**: Use consistent, descriptive naming for destination lists
3. **Access Type Validation**: Ensure access types align with security policies
4. **Global List Management**: Carefully manage global lists as they affect the entire organization
5. **Overlap Monitoring**: Monitor for destination overlaps between lists that might cause conflicts

## Security Considerations

- **Block List Effectiveness**: Regularly update block lists with new threat intelligence
- **Allow List Validation**: Ensure allow lists don't inadvertently permit malicious destinations
- **Global List Impact**: Changes to global lists affect all policies organization-wide
- **Access Control**: Limit who can modify destination lists, especially global ones

## Related Resources

- [`umbrella_destination_list`](../resources/destination_list.md) - Create and manage destination lists
- Policy and rule resources (when available) that reference destination lists
- Security monitoring and threat intelligence integration tools