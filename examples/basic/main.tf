# Example Terraform configuration for Umbrella Provider
# This demonstrates both destination list and tunnel resources

terraform {
  required_providers {
    umbrella = {
      source = "mantisec/umbrella"
      # Version constraint removed for development/testing
      # version = ">= 0.1.0"
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


# Create a destination list for blocked domains
resource "umbrella_destinationlists" "blocked_domains" {
  name = "Blocked Domains List"
  access = "block"
  destinations = [
    "malicious-site.com",
    "phishing-domain.net",
    "suspicious-website.org"
  ]
}

# Create a destination list for allowed URLs
resource "umbrella_destinationlists" "allowed_urls" {
  name = "Allowed URLs List"
  access = "allow"
  destinations = [
    "https://trusted-site.com/api",
    "https://corporate-portal.example.com",
    "https://secure-service.net/endpoint"
  ]
}

# Create a network
resource "umbrella_networks" "corporate_network" {
  name          = "Corporate Network"
  ip_address    = "10.0.0.0"
  prefix_length = 16
  is_dynamic    = false
  status        = "OPEN"
}

# Create a site
resource "umbrella_sites" "main_office" {
  name       = "Main Office"
  is_default = true
}

# Outputs for monitoring and reference
output "destination_lists" {
  description = "Created destination lists"
  value = {
    blocked_domains = {
      id     = umbrella_destinationlists.blocked_domains.id
      name   = umbrella_destinationlists.blocked_domains.name
      access = umbrella_destinationlists.blocked_domains.access
    }
    allowed_urls = {
      id     = umbrella_destinationlists.allowed_urls.id
      name   = umbrella_destinationlists.allowed_urls.name
      access = umbrella_destinationlists.allowed_urls.access
    }
  }
}

output "network" {
  description = "Created network"
  value = {
    id            = umbrella_networks.corporate_network.id
    name          = umbrella_networks.corporate_network.name
    ip_address    = umbrella_networks.corporate_network.ip_address
    prefix_length = umbrella_networks.corporate_network.prefix_length
    status        = umbrella_networks.corporate_network.status
  }
}

output "site" {
  description = "Created site"
  value = {
    id         = umbrella_sites.main_office.id
    name       = umbrella_sites.main_office.name
    is_default = umbrella_sites.main_office.is_default
  }
}

# Example terraform.tfvars file content (create separately):
# umbrella_api_key = "your-api-key-here"
# umbrella_api_secret = "your-api-secret-here"
# umbrella_org_id = "your-org-id-here"