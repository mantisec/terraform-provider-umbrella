# Example Terraform configuration replacing the curl commands from sso.tf
# This demonstrates proper Terraform resource management for SAML, rulesets, and rules

terraform {
  required_providers {
    umbrella = {
      source  = "local/umbrella"
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

variable "azure_metadata_url" {
  description = "Azure AD SAML metadata URL"
  type        = string
}

# -------------------------------------------------------------
# 1) Configure SAML Authentication (replaces curl SAML push)
# -------------------------------------------------------------
resource "umbrella_saml" "azure_ad" {
  metadata_url = var.azure_metadata_url
  auth_type    = "AzureAD"
}

# -------------------------------------------------------------
# 2) Create destination list for Microsoft login domains
#    (replaces the curl destination list creation)
# -------------------------------------------------------------
resource "umbrella_destination_list" "azure_ad_bypass" {
  name = "AzureAD-Bypass"
  type = "URL"
  destinations = [
    "login.microsoftonline.com",
    "msauth.net",
    "msftauth.net",
    "login.live.com"
  ]
}

# -------------------------------------------------------------
# 3) Configure the Default Web Policy ruleset
#    (replaces the curl ruleset patch)
# -------------------------------------------------------------
resource "umbrella_ruleset" "default_web_policy" {
  name                     = "Default Web Policy"
  description              = "Default SWG policy with SAML and SSL decryption enabled"
  saml_enabled             = true
  ssl_decryption_enabled   = true
}

# -------------------------------------------------------------
# 4) Create the bypass rule for Azure AD domains
#    (replaces the curl rule creation)
# -------------------------------------------------------------
resource "umbrella_rule" "bypass_azure_ad" {
  ruleset_id       = umbrella_ruleset.default_web_policy.id
  name             = "Bypass AzureAD"
  action           = "DO_NOT_DECRYPT"
  rank             = 1
  destination_lists = [umbrella_destination_list.azure_ad_bypass.name]
  applications     = []
  enabled          = true

  depends_on = [
    umbrella_destination_list.azure_ad_bypass,
    umbrella_ruleset.default_web_policy
  ]
}

# -------------------------------------------------------------
# Additional example rules for comprehensive SAML setup
# -------------------------------------------------------------

# Allow rule for corporate applications
resource "umbrella_rule" "allow_corporate_apps" {
  ruleset_id       = umbrella_ruleset.default_web_policy.id
  name             = "Allow Corporate Applications"
  action           = "ALLOW"
  rank             = 2
  destination_lists = []
  applications     = ["Office365", "SharePoint", "Teams"]
  enabled          = true
}

# Block rule for high-risk categories
resource "umbrella_rule" "block_high_risk" {
  ruleset_id       = umbrella_ruleset.default_web_policy.id
  name             = "Block High Risk Categories"
  action           = "BLOCK"
  rank             = 10
  destination_lists = []
  applications     = []
  enabled          = true
}

# -------------------------------------------------------------
# Outputs for monitoring and reference
# -------------------------------------------------------------
output "saml_configuration" {
  description = "SAML configuration details"
  value = {
    id           = umbrella_saml.azure_ad.id
    metadata_url = umbrella_saml.azure_ad.metadata_url
    auth_type    = umbrella_saml.azure_ad.auth_type
    enabled      = umbrella_saml.azure_ad.enabled
  }
}

output "azure_ad_bypass_list" {
  description = "Azure AD bypass destination list"
  value = {
    id           = umbrella_destination_list.azure_ad_bypass.id
    name         = umbrella_destination_list.azure_ad_bypass.name
    type         = umbrella_destination_list.azure_ad_bypass.type
    destinations = umbrella_destination_list.azure_ad_bypass.destinations
  }
}

output "default_web_policy" {
  description = "Default Web Policy ruleset configuration"
  value = {
    id                       = umbrella_ruleset.default_web_policy.id
    name                     = umbrella_ruleset.default_web_policy.name
    saml_enabled             = umbrella_ruleset.default_web_policy.saml_enabled
    ssl_decryption_enabled   = umbrella_ruleset.default_web_policy.ssl_decryption_enabled
    created_at               = umbrella_ruleset.default_web_policy.created_at
    updated_at               = umbrella_ruleset.default_web_policy.updated_at
  }
}

output "bypass_rule" {
  description = "Azure AD bypass rule configuration"
  value = {
    id               = umbrella_rule.bypass_azure_ad.id
    name             = umbrella_rule.bypass_azure_ad.name
    action           = umbrella_rule.bypass_azure_ad.action
    rank             = umbrella_rule.bypass_azure_ad.rank
    destination_lists = umbrella_rule.bypass_azure_ad.destination_lists
    enabled          = umbrella_rule.bypass_azure_ad.enabled
  }
}

# Example terraform.tfvars file content (create separately):
# umbrella_api_key = "your-api-key-here"
# umbrella_api_secret = "your-api-secret-here"
# umbrella_org_id = "your-org-id-here"
# azure_metadata_url = "https://login.microsoftonline.com/your-tenant-id/federationmetadata/2007-06/federationmetadata.xml"