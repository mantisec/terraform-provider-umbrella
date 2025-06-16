output "site_id" {
  description = "ID of the created site"
  value       = umbrella_sites.main_office.site_id
}

output "site_name" {
  description = "Name of the created site"
  value       = umbrella_sites.main_office.name
}

output "network_id" {
  description = "ID of the created network"
  value       = umbrella_networks.corporate_network.origin_id
}

output "network_name" {
  description = "Name of the created network"
  value       = umbrella_networks.corporate_network.name
}

output "internal_networks" {
  description = "Details of all created internal networks"
  value = {
    office_network = {
      id            = umbrella_internalnetworks.office_network.id
      origin_id     = umbrella_internalnetworks.office_network.origin_id
      name          = umbrella_internalnetworks.office_network.name
      ip_address    = umbrella_internalnetworks.office_network.ip_address
      prefix_length = umbrella_internalnetworks.office_network.prefix_length
      site_id       = umbrella_internalnetworks.office_network.site_id
      site_name     = umbrella_internalnetworks.office_network.site_name
      created_at    = umbrella_internalnetworks.office_network.created_at
    }
    corporate_internal = {
      id            = umbrella_internalnetworks.corporate_internal.id
      origin_id     = umbrella_internalnetworks.corporate_internal.origin_id
      name          = umbrella_internalnetworks.corporate_internal.name
      ip_address    = umbrella_internalnetworks.corporate_internal.ip_address
      prefix_length = umbrella_internalnetworks.corporate_internal.prefix_length
      network_id    = umbrella_internalnetworks.corporate_internal.network_id
      network_name  = umbrella_internalnetworks.corporate_internal.network_name
      created_at    = umbrella_internalnetworks.corporate_internal.created_at
    }
    guest_network = {
      id            = umbrella_internalnetworks.guest_network.id
      origin_id     = umbrella_internalnetworks.guest_network.origin_id
      name          = umbrella_internalnetworks.guest_network.name
      ip_address    = umbrella_internalnetworks.guest_network.ip_address
      prefix_length = umbrella_internalnetworks.guest_network.prefix_length
      site_id       = umbrella_internalnetworks.guest_network.site_id
      site_name     = umbrella_internalnetworks.guest_network.site_name
      created_at    = umbrella_internalnetworks.guest_network.created_at
    }
    datacenter_network = {
      id            = umbrella_internalnetworks.datacenter_network.id
      origin_id     = umbrella_internalnetworks.datacenter_network.origin_id
      name          = umbrella_internalnetworks.datacenter_network.name
      ip_address    = umbrella_internalnetworks.datacenter_network.ip_address
      prefix_length = umbrella_internalnetworks.datacenter_network.prefix_length
      network_id    = umbrella_internalnetworks.datacenter_network.network_id
      network_name  = umbrella_internalnetworks.datacenter_network.network_name
      created_at    = umbrella_internalnetworks.datacenter_network.created_at
    }
  }
}

output "vpn_network" {
  description = "Details of VPN internal network (if tunnel_id is provided)"
  value = var.tunnel_id != null ? {
    id            = umbrella_internalnetworks.vpn_network.id
    origin_id     = umbrella_internalnetworks.vpn_network.origin_id
    name          = umbrella_internalnetworks.vpn_network.name
    ip_address    = umbrella_internalnetworks.vpn_network.ip_address
    prefix_length = umbrella_internalnetworks.vpn_network.prefix_length
    tunnel_id     = umbrella_internalnetworks.vpn_network.tunnel_id
    tunnel_name   = umbrella_internalnetworks.vpn_network.tunnel_name
    created_at    = umbrella_internalnetworks.vpn_network.created_at
  } : null
}