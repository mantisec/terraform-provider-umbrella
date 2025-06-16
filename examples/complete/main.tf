# Complete Terraform configuration demonstrating all Umbrella Provider resources
# This example shows how to create a comprehensive Umbrella deployment

terraform {
  required_version = ">= 1.0"
  required_providers {
    umbrella = {
      source = "mantisec/umbrella"
      # Version constraint removed for development/testing
      # version = ">= 1.0.0"
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

# Variables for resource configuration
variable "admin_password" {
  description = "Password for admin users"
  type        = string
  sensitive   = true
}

variable "readonly_password" {
  description = "Password for read-only users"
  type        = string
  sensitive   = true
}

variable "company_domain" {
  description = "Company domain name"
  type        = string
  default     = "example.com"
}

# Local values for consistent naming and configuration
locals {
  # Office locations
  office_locations = {
    "headquarters" = {
      name        = "Corporate Headquarters"
      is_default  = true
      networks    = {
        "users"    = "192.168.1.0/24"
        "servers"  = "192.168.2.0/24"
        "guest"    = "192.168.100.0/24"
      }
    }
    "branch_ny" = {
      name        = "New York Branch Office"
      is_default  = false
      networks    = {
        "users"    = "192.168.10.0/24"
        "servers"  = "192.168.11.0/24"
      }
    }
    "branch_la" = {
      name        = "Los Angeles Branch Office"
      is_default  = false
      networks    = {
        "users"    = "192.168.20.0/24"
        "servers"  = "192.168.21.0/24"
      }
    }
  }

  # Security destination lists
  security_lists = {
    "malware_domains" = {
      access = "block"
      destinations = [
        "malicious-site.com",
        "phishing-domain.net",
        "suspicious-website.org",
        "known-malware-c2.com"
      ]
    }
    "trusted_partners" = {
      access = "allow"
      destinations = [
        "partner1.com",
        "trusted-vendor.net",
        "business-partner.org"
      ]
    }
    "internal_resources" = {
      access = "allow"
      destinations = [
        "intranet.${var.company_domain}",
        "wiki.${var.company_domain}",
        "tools.${var.company_domain}"
      ]
    }
    "social_media" = {
      access = "block"
      destinations = [
        "facebook.com",
        "twitter.com",
        "instagram.com",
        "tiktok.com"
      ]
    }
  }

  # User accounts to create
  user_accounts = {
    "admin1" = {
      email     = "admin1@${var.company_domain}"
      firstname = "Primary"
      lastname  = "Administrator"
      role_id   = 1  # Full Admin
      timezone  = "America/New_York"
      password  = var.admin_password
    }
    "admin2" = {
      email     = "admin2@${var.company_domain}"
      firstname = "Secondary"
      lastname  = "Administrator"
      role_id   = 1  # Full Admin
      timezone  = "America/Los_Angeles"
      password  = var.admin_password
    }
    "security_analyst" = {
      email     = "security@${var.company_domain}"
      firstname = "Security"
      lastname  = "Analyst"
      role_id   = 2  # Read Only
      timezone  = "UTC"
      password  = var.readonly_password
    }
    "reports_user" = {
      email     = "reports@${var.company_domain}"
      firstname = "Reports"
      lastname  = "User"
      role_id   = 4  # Reporting Only
      timezone  = "America/New_York"
      password  = var.readonly_password
    }
  }
}

# Create sites for all office locations
resource "umbrella_sites" "offices" {
  for_each = local.office_locations

  name       = each.value.name
  is_default = each.value.is_default
}

# Create networks (Note: IP ranges must be verified by Cisco Support)
resource "umbrella_networks" "corporate_networks" {
  for_each = {
    "corporate_main" = {
      name          = "Corporate Main Network"
      ip_address    = "10.0.0.0"
      prefix_length = 16
      is_dynamic    = false
      status        = "OPEN"
    }
    "branch_networks" = {
      name          = "Branch Office Networks"
      ip_address    = "172.16.0.0"
      prefix_length = 12
      is_dynamic    = false
      status        = "OPEN"
    }
  }

  name          = each.value.name
  ip_address    = each.value.ip_address
  prefix_length = each.value.prefix_length
  is_dynamic    = each.value.is_dynamic
  status        = each.value.status
}

# Create internal networks for each office location
resource "umbrella_internalnetworks" "office_networks" {
  for_each = merge([
    for office_key, office in local.office_locations : {
      for network_key, network_cidr in office.networks :
      "${office_key}_${network_key}" => {
        office_key   = office_key
        network_key  = network_key
        network_cidr = network_cidr
        office_name  = office.name
      }
    }
  ]...)

  name          = "${each.value.network_key} Network - ${each.value.office_name}"
  ip_address    = split("/", each.value.network_cidr)[0]
  prefix_length = tonumber(split("/", each.value.network_cidr)[1])
  site_id       = umbrella_sites.offices[each.value.office_key].site_id
}

# Create destination lists for security policies
resource "umbrella_destinationlists" "security_lists" {
  for_each = local.security_lists

  name         = title(replace(each.key, "_", " "))
  access       = each.value.access
  destinations = each.value.destinations
}

# Create user accounts
resource "umbrella_users" "company_users" {
  for_each = local.user_accounts

  email     = each.value.email
  firstname = each.value.firstname
  lastname  = each.value.lastname
  password  = each.value.password
  role_id   = each.value.role_id
  timezone  = each.value.timezone
}


# Outputs for reference and monitoring
output "deployment_summary" {
  description = "Summary of the complete Umbrella deployment"
  value = {
    sites_created = {
      for k, v in umbrella_sites.offices : k => {
        id         = v.id
        site_id    = v.site_id
        name       = v.name
        is_default = v.is_default
      }
    }
    
    networks_created = {
      for k, v in umbrella_networks.corporate_networks : k => {
        id            = v.id
        name          = v.name
        ip_address    = v.ip_address
        prefix_length = v.prefix_length
        status        = v.status
      }
    }
    
    internal_networks_created = length(umbrella_internalnetworks.office_networks)
    
    destination_lists_created = {
      for k, v in umbrella_destinationlists.security_lists : k => {
        id               = v.id
        name             = v.name
        access           = v.access
        destination_count = length(v.destinations)
      }
    }
    
    users_created = {
      for k, v in umbrella_users.company_users : k => {
        id       = v.id
        email    = v.email
        role     = v.role
        role_id  = v.role_id
        status   = v.status
      }
    }
    
  }
}

output "security_configuration" {
  description = "Security-related configuration summary"
  value = {
    blocked_destinations = flatten([
      for k, v in umbrella_destinationlists.security_lists : v.destinations
      if v.access == "block"
    ])
    
    allowed_destinations = flatten([
      for k, v in umbrella_destinationlists.security_lists : v.destinations
      if v.access == "allow"
    ])
    
    admin_users = [
      for k, v in umbrella_users.company_users : v.email
      if v.role_id == 1
    ]
    
    readonly_users = [
      for k, v in umbrella_users.company_users : v.email
      if v.role_id == 2
    ]
  }
}

output "network_topology" {
  description = "Network topology information"
  value = {
    sites = {
      for k, v in umbrella_sites.offices : k => {
        name = v.name
        internal_networks = [
          for net_k, net_v in umbrella_internalnetworks.office_networks : {
            name       = net_v.name
            ip_address = net_v.ip_address
            prefix     = net_v.prefix_length
          }
          if startswith(net_k, k)
        ]
      }
    }
  }
}

# Data source examples for validation and reference
data "umbrella_sites" "all_sites" {
  depends_on = [umbrella_sites.offices]
}

data "umbrella_destinationlists" "all_lists" {
  depends_on = [umbrella_destinationlists.security_lists]
}

output "validation_info" {
  description = "Validation information from data sources"
  value = {
    total_sites_in_org = length(data.umbrella_sites.all_sites.sites)
    total_destination_lists_in_org = length(data.umbrella_destinationlists.all_lists.destination_lists)
    
    created_sites = [
      for site in data.umbrella_sites.all_sites.sites : site.name
      if contains([for k, v in umbrella_sites.offices : v.name], site.name)
    ]
  }
}