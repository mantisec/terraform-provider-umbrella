# Complete Umbrella Deployment Example

This example demonstrates a comprehensive Cisco Umbrella deployment using all available resources in the Terraform Provider. It creates a realistic multi-site organization with security policies, user management, and network configurations.

## What This Example Creates

### Sites and Locations
- **Corporate Headquarters** (default site)
- **New York Branch Office**
- **Los Angeles Branch Office**

### Network Infrastructure
- **Corporate Networks** with verified IP ranges
- **Internal Networks** for each office location:
  - User networks (192.168.x.0/24)
  - Server networks (192.168.x.0/24)
  - Guest networks (where applicable)

### Security Policies
- **Malware Domains** block list
- **Trusted Partners** allow list
- **Internal Resources** allow list
- **Social Media** block list

### User Management
- **Primary Administrator** (Full Admin role)
- **Secondary Administrator** (Full Admin role)
- **Security Analyst** (Read Only role)
- **Reports User** (Reporting Only role)

### Secure Connectivity
- **Primary IPSec Tunnel** for headquarters
- **Secondary IPSec Tunnel** for redundancy

## Prerequisites

1. **Cisco Umbrella Account** with administrative access
2. **API Credentials** with the following scopes:
   - `deployments.sites:read` and `deployments.sites:write`
   - `deployments.networks:read` and `deployments.networks:write`
   - `deployments.internalnetworks:read` and `deployments.internalnetworks:write`
   - `policies:read` and `policies:write`
   - `admin.users:read` and `admin.users:write`
3. **IP Range Verification** - Contact Cisco Support to verify your IP ranges before creating networks
4. **Terraform** version 1.0 or later

## Quick Start

### 1. Clone and Navigate

```bash
git clone https://github.com/mantisec/terraform-provider-umbrella.git
cd terraform-provider-umbrella/examples/complete
```

### 2. Configure Variables

```bash
# Copy the example variables file
cp terraform.tfvars.example terraform.tfvars

# Edit with your actual values
nano terraform.tfvars
```

Update the following values in `terraform.tfvars`:

```hcl
umbrella_api_key    = "your-api-key-here"
umbrella_api_secret = "your-api-secret-here"
umbrella_org_id     = "your-org-id-here"
admin_password      = "SecureAdminPassword123!"
readonly_password   = "SecureReadOnlyPassword456!"
company_domain      = "yourcompany.com"
```

### 3. Initialize and Deploy

```bash
# Initialize Terraform
terraform init

# Review the planned changes
terraform plan

# Apply the configuration
terraform apply
```

## Configuration Details

### Site Configuration

The example creates three sites representing different office locations:

```hcl
local.office_locations = {
  "headquarters" = {
    name        = "Corporate Headquarters"
    is_default  = true
    networks    = {
      "users"    = "192.168.1.0/24"
      "servers"  = "192.168.2.0/24"
      "guest"    = "192.168.100.0/24"
    }
  }
  # ... additional sites
}
```

### Security Lists

Four types of destination lists are created:

1. **Malware Domains** - Blocks known malicious sites
2. **Trusted Partners** - Allows access to business partners
3. **Internal Resources** - Allows access to company resources
4. **Social Media** - Blocks social media platforms

### User Roles

The example creates users with different permission levels:

- **Role ID 1** - Full Admin (complete access)
- **Role ID 2** - Read Only (view-only access)
- **Role ID 4** - Reporting Only (reports access only)

### Network Topology

Internal networks are created for each site with logical segmentation:

- **User Networks** - End-user devices
- **Server Networks** - Infrastructure servers
- **Guest Networks** - Visitor access (headquarters only)

## Customization

### Adding More Sites

To add additional sites, extend the `local.office_locations` map:

```hcl
locals {
  office_locations = {
    # ... existing sites
    "branch_chicago" = {
      name        = "Chicago Branch Office"
      is_default  = false
      networks    = {
        "users"    = "192.168.30.0/24"
        "servers"  = "192.168.31.0/24"
      }
    }
  }
}
```

### Modifying Security Lists

Update the `local.security_lists` map to add or modify destination lists:

```hcl
locals {
  security_lists = {
    # ... existing lists
    "streaming_services" = {
      access = "block"
      destinations = [
        "netflix.com",
        "hulu.com",
        "disney.com"
      ]
    }
  }
}
```

### Adding Users

Extend the `local.user_accounts` map to create additional users:

```hcl
locals {
  user_accounts = {
    # ... existing users
    "helpdesk" = {
      email     = "helpdesk@${var.company_domain}"
      firstname = "Help"
      lastname  = "Desk"
      role_id   = 3  # Block Page Bypass
      timezone  = "America/New_York"
      password  = var.readonly_password
    }
  }
}
```

## Important Notes

### IP Range Verification

**Critical**: Before creating network resources, you must contact Cisco Support to verify your IP ranges. Attempting to create networks with unverified IP ranges will result in API errors.

1. Contact Cisco Support
2. Provide your IP ranges for verification
3. Wait for confirmation before running `terraform apply`

### Security Considerations

1. **Passwords**: Use strong, unique passwords for all user accounts
2. **API Credentials**: Store credentials securely and never commit to version control
3. **Tunnels**: Use strong pre-shared keys for IPSec tunnels
4. **Access Control**: Regularly review user roles and permissions

### Network Planning

1. **IP Addressing**: Ensure IP ranges don't conflict with existing networks
2. **Subnetting**: Plan subnet sizes based on expected device counts
3. **Growth**: Leave room for future expansion in your IP addressing scheme

## Outputs

The configuration provides comprehensive outputs for monitoring and reference:

### Deployment Summary
- Sites created with IDs and configuration
- Networks created with IP ranges and status
- Internal networks count and details
- Destination lists with access types and counts
- Users created with roles and status
- Tunnels created with endpoints and status

### Security Configuration
- All blocked destinations across lists
- All allowed destinations across lists
- Admin users list
- Read-only users list

### Network Topology
- Sites with associated internal networks
- Tunnel configurations and status

### Validation Information
- Total resources in organization
- Confirmation of created resources

## Troubleshooting

### Common Issues

#### IP Range Not Verified
```
Error: IP range not verified: 403 Forbidden
```
**Solution**: Contact Cisco Support to verify your IP ranges before creating network resources.

#### Authentication Errors
```
Error: Unable to authenticate
```
**Solution**: Verify your API credentials and organization ID are correct.

#### User Creation Failures
```
Error: Email already in use
```
**Solution**: Ensure email addresses are unique across your organization.

#### Tunnel Configuration Issues
```
Error: Invalid device IP
```
**Solution**: Ensure device IP addresses are public IPs that Umbrella can reach.

### Validation Steps

1. **Check API Credentials**: Verify credentials work with a simple data source query
2. **Validate IP Ranges**: Confirm IP ranges are verified with Cisco Support
3. **Test Incrementally**: Apply resources in stages to isolate issues
4. **Review Outputs**: Use outputs to verify resource creation

## Cleanup

To remove all resources created by this example:

```bash
terraform destroy
```

**Warning**: This will permanently delete all created resources. Ensure you have backups of any important configurations.

## Next Steps

After deploying this example:

1. **Monitor Resources**: Use the Umbrella dashboard to monitor resource status
2. **Configure Policies**: Set up additional policies using the created destination lists
3. **Test Connectivity**: Verify tunnel connectivity and network routing
4. **User Training**: Train users on new security policies and procedures
5. **Documentation**: Document your deployment for operational teams

## Support

For issues with this example:

1. Check the [provider documentation](../../docs/)
2. Review [troubleshooting guides](../../docs/guides/migration.md#troubleshooting)
3. Open an issue on the [provider repository](https://github.com/mantisec/terraform-provider-umbrella/issues)

## Related Examples

- [Basic Example](../basic/) - Simple resource configurations
- [User Management](../users/) - Focused user account management
- [Site Management](../sites/) - Site and location management
- [Internal Networks](../internal_networks/) - Network topology examples