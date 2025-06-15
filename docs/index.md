---
page_title: "Umbrella Provider"
subcategory: ""
description: |-
  The Umbrella provider is used to interact with Cisco Umbrella Secure Web Gateway resources through the REST API.
---

# Umbrella Provider

The Umbrella provider is used to interact with [Cisco Umbrella](https://umbrella.cisco.com/) Secure Web Gateway resources through the REST API. The provider allows you to manage destination lists, IPSec tunnels, SAML authentication, rulesets, and policy rules.

## Example Usage

```terraform
terraform {
  required_providers {
    umbrella = {
      source = "mantisec/umbrella"
      version = "~> 0.2"
    }
  }
}

provider "umbrella" {
  api_key    = var.umbrella_api_key
  api_secret = var.umbrella_api_secret
  org_id     = var.umbrella_org_id
}
```

## Authentication

The Umbrella provider requires OAuth2 client credentials for authentication:

- **API Key**: Your Umbrella API key (client ID)
- **API Secret**: Your Umbrella API secret (client secret)  
- **Organization ID**: Your Umbrella organization identifier

### Environment Variables

You can provide credentials via environment variables:

```bash
export TF_VAR_umbrella_api_key="your-api-key"
export TF_VAR_umbrella_api_secret="your-api-secret"
export TF_VAR_umbrella_org_id="your-org-id"
```

## Schema

### Required

- `api_key` (String, Sensitive) - Umbrella API key (client ID)
- `api_secret` (String, Sensitive) - Umbrella API secret (client secret)
- `org_id` (String) - Umbrella organization ID

## Resources

- [`umbrella_destination_list`](resources/destination_list.md) - Manages destination lists for policy enforcement
- [`umbrella_destination`](resources/destination.md) - Manages individual destinations within lists
- [`umbrella_tunnel`](resources/tunnel.md) - Manages IPSec tunnels for Secure Internet Gateway
- [`umbrella_saml`](resources/saml.md) - Manages SAML authentication configuration
- [`umbrella_ruleset`](resources/ruleset.md) - Manages SWG policy rulesets
- [`umbrella_rule`](resources/rule.md) - Manages individual policy rules within rulesets

## API Endpoints

The provider interacts with the following Umbrella API endpoints:

- **Authentication**: `POST /auth/v2/token`
- **Destination Lists**: `/policies/v2/organizations/{orgId}/destinationlists`
- **IPSec Tunnels**: `/v2/organizations/{orgId}/secureinternetgateway/ipsec/sites`
- **SAML Configuration**: `/v2/organizations/{orgId}/saml`
- **Rulesets**: `/policies/v2/organizations/{orgId}/rulesets`
- **Rules**: `/policies/v2/organizations/{orgId}/rulesets/{rulesetId}/rules`

## Security Best Practices

- Store credentials as environment variables or in Terraform Cloud workspace secrets
- Never commit plain credentials to source control
- Use variable files (`.tfvars`) that are excluded from version control
- Regularly rotate API credentials
- Use least-privilege access for API credentials