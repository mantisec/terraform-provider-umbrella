---
page_title: "Getting Started with the Umbrella Provider"
subcategory: "Guides"
description: |-
  A comprehensive guide to getting started with the Terraform Provider for Cisco Umbrella.
---

# Getting Started with the Umbrella Provider

This guide will walk you through setting up and using the Terraform Provider for Cisco Umbrella to manage your Umbrella resources with Infrastructure as Code.

## Prerequisites

Before you begin, ensure you have:

1. **Terraform installed** (version 1.0 or later)
2. **Cisco Umbrella account** with administrative access
3. **API credentials** from your Umbrella dashboard
4. **Basic understanding** of Terraform concepts

## Step 1: Obtain API Credentials

### 1.1 Access the Umbrella Dashboard

1. Log in to your [Cisco Umbrella Dashboard](https://dashboard.umbrella.com)
2. Navigate to **Admin** → **API Keys**

### 1.2 Create API Key

1. Click **Create** to generate a new API key
2. Provide a descriptive name (e.g., "Terraform Provider")
3. Select the required scopes:
   - `deployments.sites:read` and `deployments.sites:write` - For site management
   - `deployments.networks:read` and `deployments.networks:write` - For network management
   - `deployments.internalnetworks:read` and `deployments.internalnetworks:write` - For internal networks
   - `policies:read` and `policies:write` - For destination lists and policies
   - `admin.users:read` and `admin.users:write` - For user management

### 1.3 Save Credentials

After creating the API key, save the following information securely:
- **API Key** (Client ID)
- **API Secret** (Client Secret)
- **Organization ID** (found in your dashboard URL or organization settings)

## Step 2: Set Up Your Terraform Configuration

### 2.1 Create a New Directory

```bash
mkdir umbrella-terraform
cd umbrella-terraform
```

### 2.2 Create Provider Configuration

Create a `main.tf` file:

```terraform
terraform {
  required_version = ">= 1.0"
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
```

### 2.3 Create Variables File

Create a `variables.tf` file:

```terraform
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
```

### 2.4 Create Terraform Variables File

Create a `terraform.tfvars` file (never commit this to version control):

```hcl
umbrella_api_key    = "your-api-key-here"
umbrella_api_secret = "your-api-secret-here"
umbrella_org_id     = "your-org-id-here"
```

**Important**: Add `terraform.tfvars` to your `.gitignore` file to prevent committing credentials.

## Step 3: Your First Resources

### 3.1 Create a Site

Add to your `main.tf`:

```terraform
# Create your first site
resource "umbrella_sites" "main_office" {
  name = "Main Office"
}

output "main_office_id" {
  description = "The ID of the main office site"
  value       = umbrella_sites.main_office.id
}
```

### 3.2 Create a Destination List

Add to your `main.tf`:

```terraform
# Create a destination list for blocked domains
resource "umbrella_destination_list" "blocked_domains" {
  name   = "Blocked Domains"
  access = "block"
  destinations = [
    "malicious-site.com",
    "phishing-domain.net",
    "suspicious-website.org"
  ]
}

output "blocked_domains_id" {
  description = "The ID of the blocked domains list"
  value       = umbrella_destination_list.blocked_domains.id
}
```

## Step 4: Initialize and Apply

### 4.1 Initialize Terraform

```bash
terraform init
```

This command will:
- Download the Umbrella provider
- Initialize the working directory
- Create a `.terraform` directory

### 4.2 Plan Your Changes

```bash
terraform plan
```

Review the planned changes to ensure they match your expectations.

### 4.3 Apply Configuration

```bash
terraform apply
```

Type `yes` when prompted to confirm the changes.

## Step 5: Verify Your Resources

### 5.1 Check Terraform State

```bash
terraform show
```

### 5.2 Verify in Umbrella Dashboard

1. Log in to your Umbrella dashboard
2. Navigate to **Deployments** → **Sites** to see your created site
3. Navigate to **Policies** → **Destination Lists** to see your destination list

## Step 6: Advanced Configuration

### 6.1 Create Internal Networks

```terraform
# Create an internal network associated with the site
resource "umbrella_internalnetworks" "office_network" {
  name          = "Office Internal Network"
  ip_address    = "192.168.1.0"
  prefix_length = 24
  site_id       = umbrella_sites.main_office.site_id
}
```

### 6.2 Create Users

```terraform
# Create an admin user
resource "umbrella_users" "admin_user" {
  email     = "admin@yourcompany.com"
  firstname = "Admin"
  lastname  = "User"
  password  = var.admin_password
  role_id   = 1  # Full Admin
  timezone  = "America/New_York"
}

# Add the admin password variable
variable "admin_password" {
  description = "Password for the admin user"
  type        = string
  sensitive   = true
}
```

### 6.3 Create Networks

```terraform
# Create a network (requires IP verification from Cisco)
resource "umbrella_networks" "corporate_network" {
  name          = "Corporate Network"
  ip_address    = "10.0.0.0"
  prefix_length = 16
  is_dynamic    = false
  status        = "OPEN"
}
```

**Note**: Before creating networks, contact Cisco Support to verify your IP ranges.

## Step 7: Best Practices

### 7.1 Environment Variables

Instead of using `terraform.tfvars`, you can use environment variables:

```bash
export TF_VAR_umbrella_api_key="your-api-key"
export TF_VAR_umbrella_api_secret="your-api-secret"
export TF_VAR_umbrella_org_id="your-org-id"
```

### 7.2 Remote State

For team environments, use remote state:

```terraform
terraform {
  backend "s3" {
    bucket = "your-terraform-state-bucket"
    key    = "umbrella/terraform.tfstate"
    region = "us-west-2"
  }
}
```

### 7.3 Organize Your Code

Structure your configuration files:

```
umbrella-terraform/
├── main.tf              # Provider and main resources
├── variables.tf         # Variable definitions
├── outputs.tf          # Output definitions
├── terraform.tfvars    # Variable values (not in git)
├── .gitignore          # Git ignore file
└── modules/            # Custom modules
    ├── sites/
    ├── networks/
    └── users/
```

### 7.4 Use Data Sources

Leverage data sources to reference existing resources:

```terraform
# Get existing site information
data "umbrella_sites" "existing_site" {
  name = "Existing Site Name"
}

# Use in other resources
resource "umbrella_internalnetworks" "new_network" {
  name          = "New Internal Network"
  ip_address    = "192.168.2.0"
  prefix_length = 24
  site_id       = data.umbrella_sites.existing_site.site_id
}
```

## Step 8: Import Existing Resources

If you have existing Umbrella resources, you can import them:

### 8.1 Import a Site

```bash
terraform import umbrella_sites.existing_site 12345
```

### 8.2 Import a Destination List

```bash
terraform import umbrella_destination_list.existing_list 67890
```

### 8.3 Write Configuration for Imported Resources

After importing, add the corresponding resource configuration to your `.tf` files.

## Troubleshooting

### Common Issues

#### Authentication Errors
```
Error: Unable to authenticate
```
**Solution**: Verify your API credentials and organization ID are correct.

#### Permission Errors
```
Error: insufficient permissions: 403 Forbidden
```
**Solution**: Ensure your API key has the required scopes for the operations you're performing.

#### Resource Not Found
```
Error: resource not found: 404
```
**Solution**: Verify the resource exists in your Umbrella organization and you have access to it.

#### Network Verification Required
```
Error: IP range not verified: 403 Forbidden
```
**Solution**: Contact Cisco Support to verify your IP ranges before creating network resources.

### Getting Help

- Review the [provider documentation](../index.md)
- Check the [examples directory](https://github.com/mantisec/terraform-provider-umbrella/tree/main/examples)
- Open an issue on the [provider repository](https://github.com/mantisec/terraform-provider-umbrella)

## Next Steps

Now that you have the basics working:

1. **Explore Examples**: Check out the [examples directory](https://github.com/mantisec/terraform-provider-umbrella/tree/main/examples) for more complex configurations
2. **Read Resource Documentation**: Review individual resource documentation for detailed configuration options
3. **Implement CI/CD**: Set up automated Terraform workflows for your infrastructure
4. **Monitor and Maintain**: Regularly review and update your Umbrella infrastructure

## Security Reminders

- Never commit API credentials to version control
- Use environment variables or secure secret management systems
- Regularly rotate API credentials
- Follow the principle of least privilege for API key permissions
- Monitor API key usage through the Umbrella dashboard

Congratulations! You now have a working Terraform configuration for managing your Cisco Umbrella resources.