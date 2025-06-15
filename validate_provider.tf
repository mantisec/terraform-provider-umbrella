# Validation configuration to test the new provider resources
# This file can be used to validate the provider works correctly

terraform {
  required_providers {
    umbrella = {
      source  = "local/umbrella"
      version = "0.2.0"
    }
  }
}

# Provider configuration (use environment variables or terraform.tfvars)
provider "umbrella" {
  api_key    = var.umbrella_api_key
  api_secret = var.umbrella_api_secret
  org_id     = var.umbrella_org_id
}

# Variables
variable "umbrella_api_key" {
  description = "Umbrella API key"
  type        = string
  sensitive   = true
}

variable "umbrella_api_secret" {
  description = "Umbrella API secret"
  type        = string
  sensitive   = true
}

variable "umbrella_org_id" {
  description = "Umbrella organization ID"
  type        = string
}

variable "azure_metadata_url" {
  description = "Azure AD SAML metadata URL"
  type        = string
  default     = "https://login.microsoftonline.com/example-tenant-id/federationmetadata/2007-06/federationmetadata.xml"
}

# Test resources - uncomment to validate specific functionality

# Test SAML configuration
# resource "umbrella_saml" "test_saml" {
#   metadata_url = var.azure_metadata_url
#   auth_type    = "AzureAD"
# }

# Test destination list
# resource "umbrella_destination_list" "test_list" {
#   name = "Test-Validation-List"
#   type = "URL"
#   destinations = [
#     "example.com",
#     "test.org"
#   ]
# }

# Test ruleset
# resource "umbrella_ruleset" "test_ruleset" {
#   name                     = "Test Validation Ruleset"
#   description              = "Test ruleset for provider validation"
#   saml_enabled             = true
#   ssl_decryption_enabled   = false
# }

# Test rule (requires ruleset to exist)
# resource "umbrella_rule" "test_rule" {
#   ruleset_id        = umbrella_ruleset.test_ruleset.id
#   name              = "Test Validation Rule"
#   action            = "ALLOW"
#   rank              = 100
#   destination_lists = [umbrella_destination_list.test_list.name]
#   applications      = []
#   enabled           = true
# }

# Test tunnel
# resource "umbrella_tunnel" "test_tunnel" {
#   name            = "Test-Validation-Tunnel"
#   device_ip       = "203.0.113.100"
#   pre_shared_key  = "test-psk-validation-only"
# }

# Outputs for validation
# output "validation_results" {
#   description = "Validation test results"
#   value = {
#     saml_configured    = try(umbrella_saml.test_saml.enabled, false)
#     list_created       = try(umbrella_destination_list.test_list.id, "not_created")
#     ruleset_created    = try(umbrella_ruleset.test_ruleset.id, "not_created")
#     rule_created       = try(umbrella_rule.test_rule.id, "not_created")
#     tunnel_created     = try(umbrella_tunnel.test_tunnel.id, "not_created")
#   }
# }