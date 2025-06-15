# Example Terraform configuration for Umbrella Provider
# This demonstrates both destination list and tunnel resources

terraform {
  required_providers {
    umbrella = {
      source = "local/umbrella"
      version = "0.1.0"
    }
  }
}

# Configure the Umbrella Provider
provider "umbrella" {
  api_key    = var.umbrella_api_key
  api_secret = var.umbrella_api_secret
  org_id     = var.umbrella_org_id
}

# Variables for provider configuration
variable "umbrella_api_key" {
  description = "Umbrella API key (client ID)"
  type        = string
  sensitive   = true
}

variable "umbrella_api_secret" {
  description = "Umbrella API secret (client secret)"
  type        = string
  sensitive   = true
}

variable "umbrella_org_id" {
  description = "Umbrella organization ID"
  type        = string
}

# Variables for tunnel configuration
variable "tunnel_device_ip" {
  description = "Device IP address for the tunnel"
  type        = string
  default     = "203.0.113.10"
}

variable "tunnel_pre_shared_key" {
  description = "Pre-shared key for IPSec tunnel"
  type        = string
  sensitive   = true
}

# Create a destination list for blocked domains
resource "umbrella_destination_list" "blocked_domains" {
  name = "Blocked Domains List"
  type = "DOMAIN"
  destinations = [
    "malicious-site.com",
    "phishing-domain.net",
    "suspicious-website.org"
  ]
}

# Create a destination list for allowed URLs
resource "umbrella_destination_list" "allowed_urls" {
  name = "Allowed URLs List"
  type = "URL"
  destinations = [
    "https://trusted-site.com/api",
    "https://corporate-portal.example.com",
    "https://secure-service.net/endpoint"
  ]
}

# Create a destination list for IP ranges
resource "umbrella_destination_list" "internal_networks" {
  name = "Internal Network Ranges"
  type = "CIDR"
  destinations = [
    "10.0.0.0/8",
    "172.16.0.0/12",
    "192.168.0.0/16"
  ]
}

# Create primary IPSec tunnel to Umbrella SIG
resource "umbrella_tunnel" "primary_tunnel" {
  name            = "Primary-SIG-Tunnel"
  device_ip       = var.tunnel_device_ip
  pre_shared_key  = var.tunnel_pre_shared_key
}

# Create secondary IPSec tunnel for redundancy
resource "umbrella_tunnel" "secondary_tunnel" {
  name            = "Secondary-SIG-Tunnel"
  device_ip       = var.tunnel_device_ip
  pre_shared_key  = var.tunnel_pre_shared_key
}

# Outputs for monitoring and reference
output "destination_lists" {
  description = "Created destination lists"
  value = {
    blocked_domains = {
      id   = umbrella_destination_list.blocked_domains.id
      name = umbrella_destination_list.blocked_domains.name
      type = umbrella_destination_list.blocked_domains.type
    }
    allowed_urls = {
      id   = umbrella_destination_list.allowed_urls.id
      name = umbrella_destination_list.allowed_urls.name
      type = umbrella_destination_list.allowed_urls.type
    }
    internal_networks = {
      id   = umbrella_destination_list.internal_networks.id
      name = umbrella_destination_list.internal_networks.name
      type = umbrella_destination_list.internal_networks.type
    }
  }
}

output "tunnels" {
  description = "Created IPSec tunnels"
  value = {
    primary = {
      id         = umbrella_tunnel.primary_tunnel.id
      name       = umbrella_tunnel.primary_tunnel.name
      device_ip  = umbrella_tunnel.primary_tunnel.device_ip
      status     = umbrella_tunnel.primary_tunnel.status
      created_at = umbrella_tunnel.primary_tunnel.created_at
      updated_at = umbrella_tunnel.primary_tunnel.updated_at
    }
    secondary = {
      id         = umbrella_tunnel.secondary_tunnel.id
      name       = umbrella_tunnel.secondary_tunnel.name
      device_ip  = umbrella_tunnel.secondary_tunnel.device_ip
      status     = umbrella_tunnel.secondary_tunnel.status
      created_at = umbrella_tunnel.secondary_tunnel.created_at
      updated_at = umbrella_tunnel.secondary_tunnel.updated_at
    }
  }
}

# Example terraform.tfvars file content (create separately):
# umbrella_api_key = "your-api-key-here"
# umbrella_api_secret = "your-api-secret-here"
# umbrella_org_id = "your-org-id-here"
# tunnel_device_ip = "203.0.113.10"
# tunnel_pre_shared_key = "your-secure-psk-here"