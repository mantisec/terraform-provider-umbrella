---
page_title: "umbrella_sites Data Source - terraform-provider-umbrella"
subcategory: "Deployments"
description: |-
  Use this data source to retrieve information about sites in your Cisco Umbrella organization.
---

# umbrella_sites (Data Source)

Use this data source to retrieve information about sites in your Cisco Umbrella organization. Sites represent organizational locations and are foundational for other Umbrella resources.

## Example Usage

### Get a Specific Site by Name

```terraform
data "umbrella_sites" "main_office" {
  name = "Main Office"
}

# Use the site data in other resources
resource "umbrella_internalnetworks" "office_network" {
  name          = "Office Internal Network"
  ip_address    = "192.168.1.0"
  prefix_length = 24
  site_id       = data.umbrella_sites.main_office.site_id
}
```

### Get a Site by ID

```terraform
data "umbrella_sites" "example" {
  id = "12345"
}

output "site_info" {
  value = {
    name                    = data.umbrella_sites.example.name
    site_id                 = data.umbrella_sites.example.site_id
    internal_network_count  = data.umbrella_sites.example.internal_network_count
    va_count               = data.umbrella_sites.example.va_count
  }
}
```

### Get All Sites

```terraform
data "umbrella_sites" "all" {}

output "all_sites" {
  value = data.umbrella_sites.all.sites
}
```

## Schema

### Optional

- `id` (String) The unique identifier for a specific site. If provided, returns information for this site only.
- `name` (String) The name of a specific site. If provided, returns information for the site with this name.

### Read-Only

When querying a specific site (by `id` or `name`):

- `site_id` (Number) The ID of the site
- `origin_id` (Number) The origin ID of the site
- `type` (String) The type of the site (SITE or MOBILE_DEVICE)
- `is_default` (Boolean) Whether the site is the default site
- `internal_network_count` (Number) The number of internal networks attached to the site
- `va_count` (Number) The number of virtual appliances attached to the site
- `created_at` (String) The date and time when the site was created
- `modified_at` (String) The date and time when the site was last modified

When querying all sites (no filters):

- `sites` (List of Object) List of all sites with the following attributes:
  - `id` (String) The unique identifier for the site
  - `site_id` (Number) The ID of the site
  - `name` (String) The name of the site
  - `origin_id` (Number) The origin ID of the site
  - `type` (String) The type of the site
  - `is_default` (Boolean) Whether the site is the default site
  - `internal_network_count` (Number) The number of internal networks attached
  - `va_count` (Number) The number of virtual appliances attached
  - `created_at` (String) Creation timestamp
  - `modified_at` (String) Last modification timestamp

## Usage Examples

### Reference Site in Multiple Resources

```terraform
# Get the main office site
data "umbrella_sites" "main_office" {
  name = "Main Office"
}

# Create multiple internal networks for the site
resource "umbrella_internalnetworks" "office_networks" {
  for_each = {
    "user_network"  = "192.168.1.0/24"
    "server_network" = "192.168.2.0/24"
    "guest_network"  = "192.168.100.0/24"
  }

  name          = "${each.key} - ${data.umbrella_sites.main_office.name}"
  ip_address    = split("/", each.value)[0]
  prefix_length = tonumber(split("/", each.value)[1])
  site_id       = data.umbrella_sites.main_office.site_id
}
```

### Conditional Resource Creation

```terraform
data "umbrella_sites" "headquarters" {
  name = "Corporate Headquarters"
}

# Only create VPN network if headquarters exists and has virtual appliances
resource "umbrella_internalnetworks" "vpn_network" {
  count = data.umbrella_sites.headquarters.va_count > 0 ? 1 : 0

  name          = "VPN Network - ${data.umbrella_sites.headquarters.name}"
  ip_address    = "10.10.0.0"
  prefix_length = 16
  site_id       = data.umbrella_sites.headquarters.site_id
}
```

### Site Discovery and Validation

```terraform
# Get all sites for validation
data "umbrella_sites" "all" {}

# Validate that required sites exist
locals {
  required_sites = ["Main Office", "Branch Office", "Data Center"]
  existing_sites = [for site in data.umbrella_sites.all.sites : site.name]
  missing_sites  = setsubtract(toset(local.required_sites), toset(local.existing_sites))
}

# Output missing sites for validation
output "missing_sites" {
  value = local.missing_sites
}

# Create missing sites
resource "umbrella_sites" "missing" {
  for_each = local.missing_sites
  name     = each.value
}
```

## API Reference

This data source uses the following Cisco Umbrella API endpoints:

- **Get Site**: `GET /deployments/v2/sites/{siteId}` (when using `id`)
- **List Sites**: `GET /deployments/v2/sites` (when using `name` or no filters)

## Notes

- If both `id` and `name` are provided, `id` takes precedence
- When no filters are provided, all sites in the organization are returned
- Site names are case-sensitive when filtering
- The `type` field is automatically determined by the Umbrella system
- Default sites are marked with `is_default = true`

## Troubleshooting

### Common Errors

**Site not found**
```
Error: site not found with name "Office"
```
Verify the site name exists and matches exactly (case-sensitive).

**Permission denied**
```
Error: insufficient permissions: 403 Forbidden
```
Ensure your API credentials have the `deployments.sites:read` scope.

### Best Practices

1. **Use Descriptive Names**: Reference sites by descriptive names rather than IDs when possible
2. **Validate Existence**: Use data sources to validate that required sites exist before creating dependent resources
3. **Cache Results**: Store frequently used site information in local values to avoid repeated API calls
4. **Error Handling**: Use conditional logic to handle cases where sites might not exist