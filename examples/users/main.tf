terraform {
  required_providers {
    umbrella = {
      source  = "mantisec/umbrella"
      version = "~> 1.0"
    }
  }
}

provider "umbrella" {
  api_key    = var.umbrella_api_key
  api_secret = var.umbrella_api_secret
  org_id     = var.umbrella_org_id
}

# Create an admin user
resource "umbrella_users" "admin_user" {
  email     = "admin@example.com"
  firstname = "John"
  lastname  = "Doe"
  password  = var.admin_password
  role_id   = 1 # Full Admin
  timezone  = "America/New_York"
}

# Create a read-only user
resource "umbrella_users" "readonly_user" {
  email     = "readonly@example.com"
  firstname = "Jane"
  lastname  = "Smith"
  password  = var.readonly_password
  role_id   = 2 # Read Only
  timezone  = "UTC"
}

# Create a reporting user
resource "umbrella_users" "reporting_user" {
  email     = "reports@example.com"
  firstname = "Report"
  lastname  = "User"
  password  = var.reporting_password
  role_id   = 4 # Reporting Only
  timezone  = "America/Los_Angeles"
}

# Create a user with block page bypass permissions
resource "umbrella_users" "bypass_user" {
  email     = "bypass@example.com"
  firstname = "Bypass"
  lastname  = "User"
  password  = var.bypass_password
  role_id   = 3 # Block Page Bypass
  timezone  = "Europe/London"
}

# Example with international characters
resource "umbrella_users" "international_user" {
  email     = "jose.garcia@example.com"
  firstname = "José"
  lastname  = "García-López"
  password  = var.international_password
  role_id   = 2 # Read Only
  timezone  = "Europe/Madrid"
}

# Output user information (excluding sensitive data)
output "admin_user_info" {
  description = "Information about the admin user"
  value = {
    id                   = umbrella_users.admin_user.id
    user_id              = umbrella_users.admin_user.user_id
    email                = umbrella_users.admin_user.email
    firstname            = umbrella_users.admin_user.firstname
    lastname             = umbrella_users.admin_user.lastname
    role                 = umbrella_users.admin_user.role
    role_id              = umbrella_users.admin_user.role_id
    timezone             = umbrella_users.admin_user.timezone
    status               = umbrella_users.admin_user.status
    two_factor_enabled   = umbrella_users.admin_user.two_factor_enabled
    last_login_time      = umbrella_users.admin_user.last_login_time
  }
}

output "readonly_user_info" {
  description = "Information about the read-only user"
  value = {
    id                   = umbrella_users.readonly_user.id
    user_id              = umbrella_users.readonly_user.user_id
    email                = umbrella_users.readonly_user.email
    firstname            = umbrella_users.readonly_user.firstname
    lastname             = umbrella_users.readonly_user.lastname
    role                 = umbrella_users.readonly_user.role
    role_id              = umbrella_users.readonly_user.role_id
    timezone             = umbrella_users.readonly_user.timezone
    status               = umbrella_users.readonly_user.status
    two_factor_enabled   = umbrella_users.readonly_user.two_factor_enabled
    last_login_time      = umbrella_users.readonly_user.last_login_time
  }
}

output "all_users_summary" {
  description = "Summary of all created users"
  value = {
    admin_user = {
      email = umbrella_users.admin_user.email
      role  = umbrella_users.admin_user.role
    }
    readonly_user = {
      email = umbrella_users.readonly_user.email
      role  = umbrella_users.readonly_user.role
    }
    reporting_user = {
      email = umbrella_users.reporting_user.email
      role  = umbrella_users.reporting_user.role
    }
    bypass_user = {
      email = umbrella_users.bypass_user.email
      role  = umbrella_users.bypass_user.role
    }
    international_user = {
      email = umbrella_users.international_user.email
      role  = umbrella_users.international_user.role
    }
  }
}