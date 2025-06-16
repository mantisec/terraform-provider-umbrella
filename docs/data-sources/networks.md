---
page_title: "umbrella_networks Data Source - terraform-provider-umbrella"
subcategory: "Deployments"
description: |-
  Use this data source to retrieve information about networks in your Cisco Umbrella organization.
---

# umbrella_networks (Data Source)

Use this data source to retrieve information about networks in your Cisco Umbrella organization. Networks represent IP address ranges that are associated with your organization and are used to apply policies and track network activity.

## Example Usage

### Get a Specific Network by Name

```terraform
data "umbrella_networks" "office_network" {
  name = "Office Network"
}

# Use the network data in other resources
resource "umbrella_internalnetworks" "office_internal" {
  name          = "Office Internal Network"
  ip_address    = "192.168.1.0"
  prefix_length = 24
  network_id    = data.umbrella_networks.office_network.origin_id
}
```

### Get a Network by ID

```terraform
data "umbrella_networks" "example" {
  id = "12345"
}

output "network_info" {
  value = {
    name          = data.umbrella_networks.example.name
    ip_address    = data.umbrella_networks.example.ip_address
    prefix_length = data.umbrella_networks.example.prefix_length
    status        = data.umbrella_networks.example.status
    is_dynamic    = data.umbrella_networks.example.is_dynamic
    is_verified   = data.umbrella_networks.example.is_verified
  }
}
```

### Get All Networks

```terraform
data "umbrella_networks" "all" {}

# Filter verified networks
locals {
  verified_networks = [
    for network in data.umbrella_networks.all.networks : network
    if network.is_verified == true
  ]
}

output "verified_networks" {
  value = local.verified_networks
}
```

## Schema

### Optional

- `id` (String) The unique identifier for a specific network. If provided, returns information for this network only.
- `name` (String) The name of a specific network. If provided, returns information for the network with this name.

### Read-Only

When querying a specific network (by `id` or `name`):

- `origin_id` (Number) The origin ID of the network
- `ip_address` (String) The IP address of the network
- `prefix_length` (Number) The length of the prefix (29-32)
- `status` (String) The status of the network (OPEN or CLOSED)
- `is_dynamic` (Boolean) Whether the IP address is dynamic
- `is_verified` (Boolean) Whether the network is verified by Cisco Support
- `created_at` (String) The date and time when the network was created

When querying all networks (no filters):

- `networks` (List of Object) List of all networks with the following attributes:
  - `id` (String) The unique identifier for the network
  - `origin_id` (Number) The origin ID of the network
  - `name` (String) The name of the network
  - `ip_address` (String) The IP address of the network
  - `prefix_length` (Number) The length of the prefix
  - `status` (String) The status of the network
  - `is_dynamic` (Boolean) Whether the IP address is dynamic
  - `is_verified` (Boolean) Whether the network is verified
  - `created_at` (String) Creation timestamp

## Usage Examples

### Network Validation and Filtering

```terraform
# Get all networks
data "umbrella_networks" "all" {}

# Filter networks by status and verification
locals {
  active_verified_networks = [
    for network in data.umbrella_networks.all.networks : network
    if network.status == "OPEN" && network.is_verified == true
  ]
  
  dynamic_networks = [
    for network in data.umbrella_networks.all.networks : network
    if network.is_dynamic == true
  ]
}

output "active_verified_networks" {
  description = "Networks that are open and verified"
  value       = local.active_verified_networks
}

output "dynamic_networks" {
  description = "Networks with dynamic IP addresses"
  value       = local.dynamic_networks
}
```

### Conditional Internal Network Creation

```terraform
data "umbrella_networks" "corporate" {
  name = "Corporate Network"
}

# Only create internal networks if the parent network is verified
resource "umbrella_internalnetworks" "subnets" {
  for_each = data.umbrella_networks.corporate.is_verified ? {
    "users"    = "10.0.1.0/24"
    "servers"  = "10.0.2.0/24"
    "devices"  = "10.0.3.0/24"
  } : {}

  name          = "${each.key} subnet"
  ip_address    = split("/", each.value)[0]
  prefix_length = tonumber(split("/", each.value)[1])
  network_id    = data.umbrella_networks.corporate.origin_id
}
```

### Network Discovery for Policy Assignment

```terraform
# Get all verified networks for policy assignment
data "umbrella_networks" "all" {}

locals {
  verified_network_ids = [
    for network in data.umbrella_networks.all.networks : network.origin_id
    if network.is_verified == true && network.status == "OPEN"
  ]
}

# Use verified networks in policy configurations
# (This would be used with policy resources when available)
output "policy_applicable_networks" {
  description = "Network IDs that can be used in policies"
  value       = local.verified_network_ids
}
```

### Network Status Monitoring

```terraform
data "umbrella_networks" "all" {}

# Create monitoring outputs
locals {
  network_status_summary = {
    total_networks     = length(data.umbrella_networks.all.networks)
    open_networks      = length([for n in data.umbrella_networks.all.networks : n if n.status == "OPEN"])
    closed_networks    = length([for n in data.umbrella_networks.all.networks : n if n.status == "CLOSED"])
    verified_networks  = length([for n in data.umbrella_networks.all.networks : n if n.is_verified == true])
    dynamic_networks   = length([for n in data.umbrella_networks.all.networks : n if n.is_dynamic == true])
  }
}

output "network_status_summary" {
  description = "Summary of network status across the organization"
  value       = local.network_status_summary
}
```

## API Reference

This data source uses the following Cisco Umbrella API endpoints:

- **Get Network**: `GET /deployments/v2/networks/{networkId}` (when using `id`)
- **List Networks**: `GET /deployments/v2/networks` (when using `name` or no filters)

## Validation and Constraints

### Network Status Values
- `OPEN` - Network is active and allows traffic
- `CLOSED` - Network is inactive and blocks traffic

### Prefix Length Range
- Must be between 29 and 32 (as per Cisco Umbrella API requirements)

### IP Address Format
- Must be valid IPv4 address format
- Required for static networks (`is_dynamic = false`)
- Optional for dynamic networks (`is_dynamic = true`)

## Notes

- **Verification Requirement**: Networks must be verified by Cisco Support before they can be used effectively
- **Dynamic vs Static**: Dynamic networks can have changing IP addresses, while static networks have fixed IPs
- **Status Impact**: Only `OPEN` networks actively process traffic
- **Filtering**: When no filters are provided, all networks in the organization are returned
- **Case Sensitivity**: Network names are case-sensitive when filtering

## Troubleshooting

### Common Errors

**Network not found**
```
Error: network not found with name "Office Network"
```
Verify the network name exists and matches exactly (case-sensitive).

**Unverified network usage**
```
Warning: Network is not verified by Cisco Support
```
Contact Cisco Support to verify your IP range before using the network in production.

**Permission denied**
```
Error: insufficient permissions: 403 Forbidden
```
Ensure your API credentials have the `deployments.networks:read` scope.

### Best Practices

1. **Verification Check**: Always check `is_verified` before using networks in production configurations
2. **Status Validation**: Verify network status is `OPEN` before creating dependent resources
3. **Dynamic Network Handling**: Account for IP address changes in dynamic networks
4. **Monitoring**: Regularly check network status and verification state
5. **Documentation**: Document the purpose and scope of each network for team collaboration

## Related Resources

- [`umbrella_internalnetworks`](../resources/internal_networks.md) - Create internal networks associated with this network
- [`umbrella_sites`](sites.md) - Sites that may be associated with networks
- Network policies and rules (when available) for traffic management