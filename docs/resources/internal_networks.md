---
page_title: "umbrella_internalnetworks Resource - terraform-provider-umbrella"
subcategory: "Deployments"
description: |-
  Manages internal networks in Cisco Umbrella. Internal networks represent IP address ranges within sites and are foundational for network topology management.
---

# umbrella_internalnetworks (Resource)

Manages internal networks in Cisco Umbrella. Internal networks represent IP address ranges within sites and are foundational for network topology management.

Internal networks are used to define IP address ranges that are associated with sites, networks, or tunnels in your Umbrella deployment. They are essential for DNS and Web policy enforcement and network routing configuration.

## Example Usage

### Basic Internal Network with Site Association

```terraform
resource "umbrella_sites" "example" {
  name = "Main Office"
}

resource "umbrella_internalnetworks" "example" {
  name          = "Office Network"
  ip_address    = "192.168.1.0"
  prefix_length = 24
  site_id       = umbrella_sites.example.site_id
}
```

### Internal Network with Network Association

```terraform
resource "umbrella_networks" "example" {
  name          = "Corporate Network"
  ip_address    = "10.0.0.0"
  prefix_length = 16
  is_dynamic    = false
  status        = "OPEN"
}

resource "umbrella_internalnetworks" "example" {
  name          = "Corporate Internal Network"
  ip_address    = "10.0.1.0"
  prefix_length = 24
  network_id    = umbrella_networks.example.origin_id
}
```

### Internal Network with Tunnel Association

```terraform
resource "umbrella_internalnetworks" "example" {
  name          = "VPN Network"
  ip_address    = "172.16.0.0"
  prefix_length = 16
  tunnel_id     = 123
}
```

## Schema

### Required

- `name` (String) The name of the internal network. Must be between 1 and 50 characters.
- `ip_address` (String) The IPv4 address of the internal network. Must be between 7 and 15 characters.
- `prefix_length` (Number) The length of the prefix. Must be between 8 and 32.

### Optional

- `site_id` (Number) The site ID. For DNS policies, specify the ID of the site that is associated with internal network. Provide either `site_id`, `network_id`, or `tunnel_id`.
- `network_id` (Number) The network ID. For Web policies that use proxy chaining, specify the ID of the network associated with the internal network. Provide either `site_id`, `network_id`, or `tunnel_id`.
- `tunnel_id` (Number) The ID of the tunnel. For Web policies that use an IPsec tunnel, specify the ID of tunnel associated with the internal network. Provide either `site_id`, `network_id`, or `tunnel_id`.

### Read-Only

- `id` (String) The unique identifier for this internal network (same as `origin_id`).
- `origin_id` (Number) The origin ID of the internal network.
- `site_name` (String) The name of the site associated with the internal network.
- `network_name` (String) The name of the network associated with the internal network.
- `tunnel_name` (String) The name of the tunnel associated with the internal network.
- `created_at` (String) The date and time when the internal network was created.
- `modified_at` (String) The date and time when the internal network was last modified.

## Import

Internal networks can be imported using their origin ID:

```shell
terraform import umbrella_internalnetworks.example 12345
```

## Validation Rules

- **Name**: Must be between 1 and 50 characters
- **IP Address**: Must be between 7 and 15 characters (valid IPv4 format)
- **Prefix Length**: Must be between 8 and 32
- **Association**: Exactly one of `site_id`, `network_id`, or `tunnel_id` must be provided

## Dependencies

Internal networks depend on the existence of:
- Sites (when using `site_id`)
- Networks (when using `network_id`) 
- Tunnels (when using `tunnel_id`)

Ensure that the referenced site, network, or tunnel exists before creating the internal network.

## API Reference

This resource uses the Cisco Umbrella Internal Networks API:
- **Create**: `POST /deployments/v2/internalnetworks`
- **Read**: `GET /deployments/v2/internalnetworks/{internalNetworkId}`
- **Update**: `PUT /deployments/v2/internalnetworks/{internalNetworkId}`
- **Delete**: `DELETE /deployments/v2/internalnetworks/{internalNetworkId}`

## Notes

- Internal networks are foundational resources for network topology management in Umbrella
- They define IP address ranges that are used for policy enforcement and routing
- The association with sites, networks, or tunnels determines how policies are applied
- Changes to internal networks may affect policy enforcement for the associated IP ranges