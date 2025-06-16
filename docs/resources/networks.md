---
page_title: "umbrella_networks Resource - terraform-provider-umbrella"
subcategory: "Deployments"
description: |-
  Manages networks in Cisco Umbrella. Networks represent IP address ranges that are associated with your organization.
---

# umbrella_networks (Resource)

Manages networks in Cisco Umbrella. Networks represent IP address ranges that are associated with your organization and are used to apply policies and track network activity.

**Important:** Before you can create a network, you must contact Cisco Support to get your IP range verified.

## Example Usage

### Static Network

```terraform
resource "umbrella_networks" "example_static" {
  name          = "Office Network"
  ip_address    = "192.168.1.0"
  prefix_length = 30
  is_dynamic    = false
  status        = "OPEN"
}
```

### Dynamic Network

```terraform
resource "umbrella_networks" "example_dynamic" {
  name          = "Remote Workers"
  prefix_length = 32
  is_dynamic    = true
  status        = "OPEN"
}
```

## Schema

### Required

- `name` (String) The name of the network. Must be between 1 and 50 characters.
- `prefix_length` (Number) The length of the prefix. Must be between 29 and 32.
- `status` (String) The status of the network. Must be either `OPEN` or `CLOSED`.

### Optional

- `ip_address` (String) The IP address of the network. Required for static networks (when `is_dynamic` is `false`), optional for dynamic networks. Must be at least 7 characters long.
- `is_dynamic` (Boolean) Specifies whether the IP address is dynamic. Defaults to `false`.

### Read-Only

- `id` (String) The unique identifier for this network (same as `origin_id`).
- `origin_id` (Number) The origin ID of the network.
- `is_verified` (Boolean) Specifies whether the network is verified by Cisco Support.
- `created_at` (String) The date and time when the network was created.

## Import

Networks can be imported using their ID:

```shell
terraform import umbrella_networks.example 12345
```

## API Reference

This resource uses the following Cisco Umbrella API endpoints:

- **Create Network**: `POST /deployments/v2/networks`
- **Get Network**: `GET /deployments/v2/networks/{networkId}`
- **Update Network**: `PUT /deployments/v2/networks/{networkId}`
- **Delete Network**: `DELETE /deployments/v2/networks/{networkId}`

## Notes

1. **IP Range Verification**: Before creating a network, you must contact Cisco Support to verify your IP range. Attempting to create a network with an unverified IP range will result in a 403 Forbidden error.

2. **Static vs Dynamic Networks**: 
   - Static networks require an `ip_address` to be specified
   - Dynamic networks do not require an `ip_address` and can have their IP address change over time

3. **Prefix Length Constraints**: The prefix length must be between 29 and 32, as required by the Cisco Umbrella API.

4. **Network Status**: Networks can be either `OPEN` (allowing traffic) or `CLOSED` (blocking traffic).

5. **Updates**: When updating a network, you can modify the name, IP address (if verified), and status. The `is_dynamic` and `prefix_length` fields may have restrictions on updates depending on the current network configuration.