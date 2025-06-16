terraform {
  required_providers {
    umbrella = {
      source  = "mantisec/umbrella"
      version = "~> 1.0"
    }
  }
}

provider "umbrella" {
  # Configuration will be provided via environment variables:
  # UMBRELLA_API_KEY
  # UMBRELLA_API_SECRET
  # UMBRELLA_ORG_ID
}

# Basic site configuration
resource "umbrella_sites" "main_office" {
  name = "Main Office"
}

# Site with explicit configuration
resource "umbrella_sites" "branch_office" {
  name       = "Branch Office - New York"
  is_default = false
}

# Multiple sites for different locations
resource "umbrella_sites" "offices" {
  for_each = toset([
    "San Francisco Office",
    "London Office", 
    "Tokyo Office"
  ])
  
  name = each.value
}

# Output the site information
output "main_office_id" {
  description = "The ID of the main office site"
  value       = umbrella_sites.main_office.id
}

output "main_office_site_id" {
  description = "The site ID of the main office"
  value       = umbrella_sites.main_office.site_id
}

output "all_office_sites" {
  description = "All office sites created"
  value = {
    for k, v in umbrella_sites.offices : k => {
      id      = v.id
      site_id = v.site_id
      name    = v.name
    }
  }
}