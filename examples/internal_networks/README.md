# Internal Networks Example

This example demonstrates how to create and manage internal networks in Cisco Umbrella using the Terraform provider.

## Overview

Internal networks represent IP address ranges within sites and are foundational for network topology management in Umbrella. They are used for:

- DNS policy enforcement (when associated with sites)
- Web policy enforcement with proxy chaining (when associated with networks)
- Web policy enforcement with IPsec tunnels (when associated with tunnels)

## Resources Created

This example creates:

1. **Site**: A main office site that serves as the foundation
2. **Network**: A corporate network for proxy chaining scenarios
3. **Internal Networks**:
   - Office network associated with the site (for DNS policies)
   - Corporate internal network associated with the network (for Web policies)
   - Guest network with smaller subnet (/28)
   - Datacenter network with larger subnet (/8)
   - VPN network associated with a tunnel (optional)

## Prerequisites

- Cisco Umbrella account with API access
- API key and secret with appropriate permissions
- Organization ID
- (Optional) Existing tunnel ID for tunnel-based internal network

## Usage

1. **Set up variables**: Create a `terraform.tfvars` file or set environment variables:

```hcl
umbrella_api_key    = "your-api-key"
umbrella_api_secret = "your-api-secret"
umbrella_org_id     = "your-org-id"
tunnel_id           = 123  # Optional: ID of existing tunnel
```

2. **Initialize Terraform**:
```bash
terraform init
```

3. **Plan the deployment**:
```bash
terraform plan
```

4. **Apply the configuration**:
```bash
terraform apply
```

## Configuration Details

### Site Association (DNS Policies)
```hcl
resource "umbrella_internalnetworks" "office_network" {
  name          = "Office Internal Network"
  ip_address    = "192.168.1.0"
  prefix_length = 24
  site_id       = umbrella_sites.main_office.site_id
}
```

### Network Association (Web Policies with Proxy Chaining)
```hcl
resource "umbrella_internalnetworks" "corporate_internal" {
  name          = "Corporate Internal Network"
  ip_address    = "10.0.1.0"
  prefix_length = 24
  network_id    = umbrella_networks.corporate_network.origin_id
}
```

### Tunnel Association (Web Policies with IPsec)
```hcl
resource "umbrella_internalnetworks" "vpn_network" {
  name          = "VPN Internal Network"
  ip_address    = "172.16.0.0"
  prefix_length = 16
  tunnel_id     = var.tunnel_id
}
```

## Validation Rules

- **Name**: 1-50 characters
- **IP Address**: 7-15 characters (valid IPv4 format)
- **Prefix Length**: 8-32 (CIDR notation)
- **Association**: Exactly one of `site_id`, `network_id`, or `tunnel_id` must be provided

## Outputs

The example provides detailed outputs including:
- Site and network information
- All internal network details (ID, name, IP configuration, associations)
- Creation timestamps
- Associated resource names

## Network Planning

### Subnet Sizing Examples

- `/8` (255.0.0.0): ~16.7 million addresses - Large enterprise networks
- `/16` (255.255.0.0): ~65,000 addresses - Medium enterprise networks  
- `/24` (255.255.255.0): 254 addresses - Standard office networks
- `/28` (255.255.255.240): 14 addresses - Small guest networks
- `/30` (255.255.255.252): 2 addresses - Point-to-point links
- `/32` (255.255.255.255): 1 address - Single host

### IP Address Ranges

Common private IP ranges used in this example:
- `192.168.x.x` - Class C private networks
- `10.x.x.x` - Class A private networks
- `172.16.x.x` - Class B private networks

## Dependencies

Internal networks depend on:
- **Sites** (when using `site_id`)
- **Networks** (when using `network_id`)
- **Tunnels** (when using `tunnel_id`)

Ensure the referenced resources exist before creating internal networks.

## Cleanup

To destroy all resources:
```bash
terraform destroy
```

## Notes

- Internal networks are foundational for Umbrella's network topology
- They define IP ranges for policy enforcement and routing
- Changes may affect policy enforcement for associated IP ranges
- Consider network overlap and routing implications when planning IP ranges