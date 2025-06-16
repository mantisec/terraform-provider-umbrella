variable "umbrella_api_key" {
  description = "Cisco Umbrella API Key"
  type        = string
  sensitive   = true
}

variable "umbrella_api_secret" {
  description = "Cisco Umbrella API Secret"
  type        = string
  sensitive   = true
}

variable "umbrella_org_id" {
  description = "Cisco Umbrella Organization ID"
  type        = string
}

variable "admin_password" {
  description = "Password for the admin user"
  type        = string
  sensitive   = true
  validation {
    condition     = length(var.admin_password) >= 8
    error_message = "Admin password must be at least 8 characters long."
  }
}

variable "readonly_password" {
  description = "Password for the read-only user"
  type        = string
  sensitive   = true
  validation {
    condition     = length(var.readonly_password) >= 8
    error_message = "Read-only password must be at least 8 characters long."
  }
}

variable "reporting_password" {
  description = "Password for the reporting user"
  type        = string
  sensitive   = true
  validation {
    condition     = length(var.reporting_password) >= 8
    error_message = "Reporting password must be at least 8 characters long."
  }
}

variable "bypass_password" {
  description = "Password for the bypass user"
  type        = string
  sensitive   = true
  validation {
    condition     = length(var.bypass_password) >= 8
    error_message = "Bypass password must be at least 8 characters long."
  }
}

variable "international_password" {
  description = "Password for the international user"
  type        = string
  sensitive   = true
  validation {
    condition     = length(var.international_password) >= 8
    error_message = "International password must be at least 8 characters long."
  }
}

variable "default_timezone" {
  description = "Default timezone for users"
  type        = string
  default     = "UTC"
}

variable "organization_domain" {
  description = "Organization domain for email addresses"
  type        = string
  default     = "example.com"
  validation {
    condition     = can(regex("^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]\\.[a-zA-Z]{2,}$", var.organization_domain))
    error_message = "Organization domain must be a valid domain name."
  }
}
