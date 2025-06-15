---
page_title: "umbrella_tunnel Resource - terraform-provider-umbrella"
subcategory: ""
description: |-
  Manages IPSec tunnels for Umbrella Secure Internet Gateway.
---

# umbrella_tunnel (Resource)

Manages IPSec tunnels for Umbrella Secure Internet Gateway (SIG). These tunnels provide secure connectivity between your network infrastructure and Umbrella's cloud security services.

## Example Usage

### Basic Tunnel Configuration

```terraform
resource "umbrella_tunnel" "primary_tunnel" {
  name            = "Primary-SIG-Tunnel"
  site_origin_id  = 12345
  device_ip       = "203.0.113.10"
  pre_shared_key  = var.tunnel_psk
  local_networks  = ["10.0.0.0/8", "192.168.1.0/24"]
  tunnel_type     = "IPSEC"
}
```

### Multiple Tunnels for Redundancy

```terraform
resource "umbrella_tunnel" "primary_tunnel" {
  name            = "Primary-SIG-Tunnel"
  site_origin_id  = 12345
  device_ip       = "203.0.113.10"
  pre_shared_key  = var.primary_tunnel_psk
  local_networks  = ["10.0.0.0/8", "192.168.1.0/24"]
  tunnel_type     = "IPSEC"
}

resource "umbrella_tunnel" "secondary_tunnel" {
  name            = "Secondary-SIG-Tunnel"
  site_origin_id  = 12345
  device_ip       = "203.0.113.11"
  pre_shared_key  = var.secondary_tunnel_psk
  local_networks  = ["10.0.0.0/8", "192.168.1.0/24"]
  tunnel_type     = "IPSEC"
}
```

## Schema

### Required

- `name` (String) - Name of the tunnel. Must be unique within your organization
- `site_origin_id` (Number) - Site origin ID to associate with the tunnel. This identifies the site/location for the tunnel
- `device_ip` (String) - Public IP address of your network device that will establish the IPSec tunnel
- `pre_shared_key` (String, Sensitive) - Pre-shared key for IPSec authentication. This should be a strong, randomly generated key
- `local_networks` (List of String) - List of local network CIDR blocks that will use this tunnel (e.g., ["10.0.0.0/8", "192.168.1.0/24"])

### Optional

- `tunnel_type` (String) - Type of tunnel. Defaults to "IPSEC" if not specified

### Read-Only

- `id` (String) - Unique identifier of the tunnel
- `status` (String) - Current status of the tunnel (e.g., "ACTIVE", "INACTIVE", "PENDING")
- `tunnel_endpoint` (String) - Umbrella tunnel endpoint IP address for configuring your network device
- `created_at` (String) - Creation timestamp in ISO 8601 format
- `updated_at` (String) - Last update timestamp in ISO 8601 format

## Import

Tunnels can be imported using their ID:

```bash
terraform import umbrella_tunnel.example 12345678-1234-1234-1234-123456789012
```

## Notes

### Security Considerations

- **Pre-shared Key**: Use a strong, randomly generated pre-shared key. Consider using Terraform's `random_password` resource:
  ```terraform
  resource "random_password" "tunnel_psk" {
    length  = 32
    special = true
  }
  
  resource "umbrella_tunnel" "example" {
    name            = "Example-Tunnel"
    site_origin_id  = 12345
    device_ip       = "203.0.113.10"
    pre_shared_key  = random_password.tunnel_psk.result
    local_networks  = ["10.0.0.0/8", "192.168.1.0/24"]
    tunnel_type     = "IPSEC"
  }
  ```

- **IP Address**: The device IP must be a public IP address that Umbrella can reach
- **Firewall Rules**: Ensure your firewall allows IPSec traffic (UDP 500, UDP 4500, ESP protocol)

### Tunnel Configuration

After creating the tunnel resource, you'll need to configure your network device (firewall, router) with:
- The tunnel endpoints provided by Umbrella
- The pre-shared key
- Appropriate IPSec parameters (encryption, authentication, etc.)

### Monitoring

- Use the `status` attribute to monitor tunnel health
- Check `updated_at` to see when the tunnel configuration was last modified
- Monitor tunnel connectivity through your network device and Umbrella dashboard

### Limitations

- Tunnel names must be unique within your organization
- The device IP cannot be changed after tunnel creation (requires recreation)
- Pre-shared key changes will trigger tunnel reconfiguration