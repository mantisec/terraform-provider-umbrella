---
page_title: "Troubleshooting Guide - Umbrella Provider"
subcategory: "Guides"
description: |-
  Comprehensive troubleshooting guide for the Terraform Provider for Cisco Umbrella.
---

# Troubleshooting Guide

This guide helps you diagnose and resolve common issues when using the Terraform Provider for Cisco Umbrella.

## Authentication Issues

### Error: Unable to authenticate

```
Error: Unable to authenticate
│ 
│   with provider["registry.terraform.io/mantisec/umbrella"],
│   on main.tf line 10, in provider "umbrella":
│   10: provider "umbrella" {
│ 
│ Failed to authenticate with Umbrella API: invalid credentials
```

**Causes and Solutions:**

1. **Incorrect API Credentials**
   ```bash
   # Verify your credentials
   export TF_VAR_umbrella_api_key="your-actual-api-key"
   export TF_VAR_umbrella_api_secret="your-actual-api-secret"
   export TF_VAR_umbrella_org_id="your-actual-org-id"
   ```

2. **Wrong Organization ID**
   - Check your Umbrella dashboard URL: `https://dashboard.umbrella.com/o/{org-id}/`
   - Or find it in **Admin** → **Organization Settings**

3. **API Key Permissions**
   - Ensure your API key has the required scopes:
     - `deployments.sites:read` and `deployments.sites:write`
     - `deployments.networks:read` and `deployments.networks:write`
     - `policies:read` and `policies:write`
     - `admin.users:read` and `admin.users:write`

4. **Expired or Revoked API Key**
   - Check the API key status in **Admin** → **API Keys**
   - Create a new API key if necessary

### Error: Token refresh failed

```
Error: Token refresh failed: 401 Unauthorized
```

**Solution:**
- This usually indicates expired or invalid API credentials
- Regenerate your API key and secret in the Umbrella dashboard
- Update your Terraform configuration with the new credentials

## Permission Issues

### Error: Insufficient permissions

```
Error: insufficient permissions: 403 Forbidden
│ 
│   with umbrella_sites.example,
│   on main.tf line 15, in resource "umbrella_sites" "example":
│   15: resource "umbrella_sites" "example" {
```

**Solutions:**

1. **Check API Key Scopes**
   ```bash
   # Required scopes for different operations:
   # Sites: deployments.sites:read, deployments.sites:write
   # Networks: deployments.networks:read, deployments.networks:write
   # Users: admin.users:read, admin.users:write
   # Policies: policies:read, policies:write
   ```

2. **Verify Organization Access**
   - Ensure your user account has admin privileges in the organization
   - Check that the organization ID is correct

3. **Resource-Specific Permissions**
   - Some operations require specific permissions beyond basic API access
   - Contact your Umbrella administrator to verify your access level

## Resource Creation Issues

### Error: IP range not verified

```
Error: IP range not verified: 403 Forbidden
│ 
│   with umbrella_networks.example,
│   on main.tf line 20, in resource "umbrella_networks" "example":
│   20: resource "umbrella_networks" "example" {
```

**Solution:**
1. **Contact Cisco Support** to verify your IP ranges
2. **Provide the following information:**
   - Organization ID
   - IP ranges you want to use
   - Business justification for the ranges
3. **Wait for verification** before creating network resources
4. **Use verified ranges only** in your Terraform configuration

### Error: Resource name already exists

```
Error: Failed to create site: site name already exists
│ 
│   with umbrella_sites.example,
│   on main.tf line 15, in resource "umbrella_sites" "example":
│   15: resource "umbrella_sites" "example" {
```

**Solutions:**

1. **Check Existing Resources**
   ```terraform
   # Use data sources to check existing resources
   data "umbrella_sites" "all" {}
   
   output "existing_sites" {
     value = [for site in data.umbrella_sites.all.sites : site.name]
   }
   ```

2. **Use Unique Names**
   ```terraform
   resource "umbrella_sites" "example" {
     name = "My Unique Site Name ${random_id.site_suffix.hex}"
   }
   
   resource "random_id" "site_suffix" {
     byte_length = 4
   }
   ```

3. **Import Existing Resources**
   ```bash
   # Import existing resource instead of creating new
   terraform import umbrella_sites.example 12345
   ```

### Error: Invalid email format

```
Error: validation failed: email must be a valid email address
│ 
│   with umbrella_users.example,
│   on main.tf line 25, in resource "umbrella_users" "example":
│   25: resource "umbrella_users" "example" {
```

**Solution:**
```terraform
resource "umbrella_users" "example" {
  email     = "user@example.com"  # Must contain @ symbol
  firstname = "John"
  lastname  = "Doe"
  password  = var.user_password
  role_id   = 1
  timezone  = "UTC"
}
```

## State Management Issues

### Error: Resource not found in state

```
Error: resource not found in state
```

**Solutions:**

1. **Import Missing Resources**
   ```bash
   # Find the resource ID in Umbrella dashboard
   terraform import umbrella_sites.example 12345
   ```

2. **Refresh State**
   ```bash
   terraform refresh
   ```

3. **Remove from State (if resource was deleted externally)**
   ```bash
   terraform state rm umbrella_sites.example
   ```

### Error: State lock

```
Error: Error acquiring the state lock
```

**Solutions:**

1. **Wait for Lock Release**
   - Another Terraform operation may be in progress
   - Wait for it to complete

2. **Force Unlock (use with caution)**
   ```bash
   terraform force-unlock LOCK_ID
   ```

3. **Check Backend Configuration**
   - Ensure your state backend is properly configured
   - Verify network connectivity to remote state storage

## Import Issues

### Error: Resource not found during import

```
Error: Cannot import non-existent remote object
```

**Solutions:**

1. **Verify Resource ID**
   ```bash
   # Get correct resource ID from Umbrella dashboard or API
   curl -H "Authorization: Bearer $TOKEN" \
     "https://api.umbrella.com/deployments/v2/sites" | jq '.[] | {id: .siteId, name: .name}'
   ```

2. **Check Resource Type**
   ```bash
   # Ensure you're using the correct resource type
   terraform import umbrella_sites.example 12345  # For sites
   terraform import umbrella_users.example 67890  # For users
   ```

3. **Verify Permissions**
   - Ensure you have read access to the resource
   - Check that the resource exists in your organization

## Network Connectivity Issues

### Error: Connection timeout

```
Error: context deadline exceeded (Client.Timeout exceeded while awaiting headers)
```

**Solutions:**

1. **Check Network Connectivity**
   ```bash
   # Test connectivity to Umbrella API
   curl -I https://api.umbrella.com
   ```

2. **Configure Proxy (if needed)**
   ```bash
   export HTTP_PROXY=http://proxy.company.com:8080
   export HTTPS_PROXY=http://proxy.company.com:8080
   ```

3. **Firewall Rules**
   - Ensure outbound HTTPS (443) access to `api.umbrella.com`
   - Check corporate firewall rules

### Error: DNS resolution failed

```
Error: no such host: api.umbrella.com
```

**Solutions:**

1. **Check DNS Configuration**
   ```bash
   nslookup api.umbrella.com
   dig api.umbrella.com
   ```

2. **Use Alternative DNS**
   ```bash
   # Temporarily use public DNS
   export DNS_SERVER=8.8.8.8
   ```

3. **Verify Network Configuration**
   - Check `/etc/resolv.conf` (Linux/macOS)
   - Verify DNS settings in network configuration

## Rate Limiting Issues

### Error: Rate limit exceeded

```
Error: rate limit exceeded: 429 Too Many Requests
```

**Solutions:**

1. **Reduce Parallelism**
   ```bash
   terraform apply -parallelism=1
   ```

2. **Add Delays Between Operations**
   ```terraform
   resource "time_sleep" "wait" {
     create_duration = "30s"
   }
   
   resource "umbrella_sites" "example" {
     depends_on = [time_sleep.wait]
     name = "Example Site"
   }
   ```

3. **Contact Cisco Support**
   - Request rate limit increase if needed
   - Discuss your use case and requirements

## Validation Errors

### Error: Invalid prefix length

```
Error: prefix_length must be between 29 and 32
│ 
│   with umbrella_networks.example,
│   on main.tf line 20, in resource "umbrella_networks" "example":
│   20:   prefix_length = 24
```

**Solution:**
```terraform
resource "umbrella_networks" "example" {
  name          = "Example Network"
  ip_address    = "192.168.1.0"
  prefix_length = 30  # Must be between 29-32
  is_dynamic    = false
  status        = "OPEN"
}
```

### Error: Invalid role ID

```
Error: role ID 5 does not exist in organization
```

**Solution:**
```terraform
# Use valid role IDs (may vary by organization)
resource "umbrella_users" "example" {
  email     = "user@example.com"
  firstname = "John"
  lastname  = "Doe"
  password  = var.user_password
  role_id   = 1  # 1=Full Admin, 2=Read Only, 3=Block Page Bypass, 4=Reporting Only
  timezone  = "UTC"
}
```

## Debugging Techniques

### Enable Debug Logging

```bash
# Enable Terraform debug logging
export TF_LOG=DEBUG
export TF_LOG_PATH=terraform.log

# Run Terraform with debug output
terraform apply
```

### Validate Configuration

```bash
# Check configuration syntax
terraform validate

# Format configuration files
terraform fmt

# Check for potential issues
terraform plan -detailed-exitcode
```

### Test API Connectivity

```bash
# Test API authentication
curl -X POST \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "grant_type=client_credentials" \
  -u "$API_KEY:$API_SECRET" \
  "https://api.umbrella.com/auth/v2/token"
```

### Use Data Sources for Validation

```terraform
# Validate existing resources
data "umbrella_sites" "all" {}

output "debug_sites" {
  value = data.umbrella_sites.all.sites
}
```

## Getting Help

### Before Seeking Help

1. **Check this troubleshooting guide**
2. **Review the [provider documentation](../index.md)**
3. **Search [existing issues](https://github.com/mantisec/terraform-provider-umbrella/issues)**
4. **Enable debug logging** and review the output

### When Opening an Issue

Include the following information:

1. **Provider version**
   ```bash
   terraform version
   ```

2. **Terraform configuration** (sanitized)
   ```hcl
   # Remove sensitive information
   provider "umbrella" {
     api_key    = "***"
     api_secret = "***"
     org_id     = "***"
   }
   ```

3. **Error messages** (complete output)

4. **Debug logs** (if relevant, sanitized)

5. **Steps to reproduce** the issue

6. **Expected vs actual behavior**

### Support Channels

- **GitHub Issues**: [Provider Repository](https://github.com/mantisec/terraform-provider-umbrella/issues)
- **Email Support**: [grant.lawton@mantisec.com.au](mailto:grant.lawton@mantisec.com.au)
- **Cisco Support**: For API-related questions and IP verification

## Prevention Tips

### Best Practices

1. **Use Version Constraints**
   ```hcl
   terraform {
     required_providers {
       umbrella = {
         source  = "mantisec/umbrella"
         version = "~> 1.0"
       }
     }
   }
   ```

2. **Validate Before Apply**
   ```bash
   terraform validate
   terraform plan
   # Review plan output before applying
   terraform apply
   ```

3. **Use Remote State**
   ```hcl
   terraform {
     backend "s3" {
       bucket = "my-terraform-state"
       key    = "umbrella/terraform.tfstate"
       region = "us-west-2"
     }
   }
   ```

4. **Implement Proper Error Handling**
   ```terraform
   resource "umbrella_sites" "example" {
     name = var.site_name
     
     lifecycle {
       prevent_destroy = true
     }
   }
   ```

5. **Regular State Maintenance**
   ```bash
   # Regularly refresh state
   terraform refresh
   
   # Check for drift
   terraform plan
   ```

### Monitoring and Alerting

1. **Set up monitoring** for Terraform operations
2. **Monitor API usage** through Umbrella dashboard
3. **Track resource changes** and state modifications
4. **Implement automated testing** for configurations

By following this troubleshooting guide and implementing the suggested best practices, you should be able to resolve most common issues with the Umbrella provider and maintain a stable Infrastructure as Code deployment.