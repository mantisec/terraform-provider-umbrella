terraform {
  required_providers {
    umbrella = {
      source  = "cisco-open/umbrella"
      version = "~> 1.0"
    }
  }
}

provider "umbrella" {
  api_key    = var.umbrella_api_key
  api_secret = var.umbrella_api_secret
  org_id     = var.umbrella_org_id
}

# Create a site first (required for internal network association)
resource "umbrella_sites" "main_office" {
  name = "Main Office Site"
}

# Create a network (alternative association option)
resource "umbrella_networks" "corporate_network" {
  name          = "Corporate Network"
  ip_address    = "10.0.0.0"
  prefix_length = 16
  is_dynamic    = false
  status        = "OPEN"
}

# Internal Network associated with a Site (for DNS policies)
resource "umbrella_internalnetworks" "office_network" {
  name          = "Office Internal Network"
  ip_address    = "192.168.1.0"
  prefix_length = 24
  site_id       = umbrella_sites.main_office.site_id
}

# Internal Network associated with a Network (for Web policies with proxy chaining)
resource "umbrella_internalnetworks" "corporate_internal" {
  name          = "Corporate Internal Network"
  ip_address    = "10.0.1.0"
  prefix_length = 24
  network_id    = umbrella_networks.corporate_network.origin_id
}

# Internal Network associated with a Tunnel (for Web policies with IPsec tunnel)
# Note: tunnel_id would typically reference an existing tunnel resource
resource "umbrella_internalnetworks" "vpn_network" {
  name          = "VPN Internal Network"
  ip_address    = "172.16.0.0"
  prefix_length = 16
  tunnel_id     = var.tunnel_id # This should reference an existing tunnel
}

# Example with smaller subnet
resource "umbrella_internalnetworks" "guest_network" {
  name          = "Guest Network"
  ip_address    = "192.168.100.0"
  prefix_length = 28 # /28 subnet (16 addresses)
  site_id       = umbrella_sites.main_office.site_id
}

# Example with larger subnet
resource "umbrella_internalnetworks" "datacenter_network" {
  name          = "Datacenter Network"
  ip_address    = "10.10.0.0"
  prefix_length = 8 # /8 subnet (large network)
  network_id    = umbrella_networks.corporate_network.origin_id
}