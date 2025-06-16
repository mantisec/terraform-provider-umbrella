# Umbrella Users Example

This example demonstrates how to create and manage user accounts in Cisco Umbrella using the Terraform provider.

## Overview

This example creates five different users with various roles:

1. **Admin User** - Full administrative access
2. **Read-Only User** - Read-only access to the console
3. **Reporting User** - Access to reporting features only
4. **Bypass User** - Can bypass block pages
5. **International User** - Demonstrates support for international characters

## Prerequisites

- Cisco Umbrella account with API access
- Terraform >= 1.0
- Valid API credentials with user management permissions

## Required Permissions

Your API credentials must have the following scopes:
- `admin.users:write` - To create and delete users
- `admin.users:read` - To read user information
- `admin.roles:read` - To validate role assignments

## Usage

1. **Clone or download this example**

2. **Set up your variables**

   Create a `terraform.tfvars` file:
   ```hcl
   umbrella_api_key    = "your-api-key"
   umbrella_api_secret = "your-api-secret"
   umbrella_org_id     = "your-org-id"
   
   admin_password         = "SecureAdminPassword123!"
   readonly_password      = "SecureReadOnlyPassword123!"
   reporting_password     = "SecureReportingPassword123!"
   bypass_password        = "SecureBypassPassword123!"
   international_password = "SecureInternationalPassword123!"
   ```

   **Security Note**: Never commit passwords to version control. Consider using environment variables or a secure secret management system.

3. **Initialize Terraform**
   ```bash
   terraform init
   ```

4. **Plan the deployment**
   ```bash
   terraform plan
   ```

5. **Apply the configuration**
   ```bash
   terraform apply
   ```

6. **View the outputs**
   ```bash
   terraform output
   ```

## Important Notes

### API Limitations

- **No Updates**: The Umbrella Users API does not support updating existing users. Any changes to user attributes will force recreation of the user resource.
- **Email Uniqueness**: Email addresses must be unique across the organization.
- **Password Security**: Passwords are not returned by the API and cannot be verified through drift detection.

### Role IDs

The example uses common role IDs:
- `1` - Full Admin
- `2` - Read Only  
- `3` - Block Page Bypass
- `4` - Reporting Only

**Note**: Role IDs may vary between organizations. Verify the correct role IDs for your organization using the Umbrella console or API.

### Security Best Practices

1. **Strong Passwords**: Use complex passwords with a mix of characters
2. **Unique Passwords**: Each user should have a unique password
3. **Regular Rotation**: Implement a password rotation policy
4. **Least Privilege**: Assign users the minimum role required for their function
5. **Two-Factor Authentication**: Enable 2FA through the Umbrella console

## Customization

### Adding More Users

To add additional users, create new `umbrella_users` resources:

```hcl
resource "umbrella_users" "custom_user" {
  email     = "custom@example.com"
  firstname = "Custom"
  lastname  = "User"
  password  = var.custom_password
  role_id   = 2
  timezone  = "America/Chicago"
}
```

### Using Dynamic Configuration

For managing many users, consider using Terraform's `for_each`:

```hcl
variable "users" {
  description = "Map of users to create"
  type = map(object({
    firstname = string
    lastname  = string
    email     = string
    role_id   = number
    timezone  = string
  }))
}

resource "umbrella_users" "dynamic_users" {
  for_each = var.users
  
  email     = each.value.email
  firstname = each.value.firstname
  lastname  = each.value.lastname
  password  = var.default_password
  role_id   = each.value.role_id
  timezone  = each.value.timezone
}
```

## Outputs

The example provides several outputs:

- **Individual User Info**: Detailed information for each user
- **Summary**: Overview of all created users
- **User Count**: Total number of users created
- **Users by Role**: Users grouped by their assigned roles

## Troubleshooting

### Common Issues

1. **"Email already in use"**
   - The email address is already associated with another user
   - Use a different email address or delete the existing user first

2. **"Invalid role ID"**
   - The specified role ID doesn't exist in your organization
   - Check available roles using the Umbrella console or API

3. **"Insufficient permissions"**
   - Your API credentials lack the required permissions
   - Ensure your API key has `admin.users:write` scope

4. **"Update Not Supported"**
   - Attempted to modify an existing user
   - Delete and recreate the user resource instead

### Validation Errors

The example includes validation for:
- Password length (minimum 8 characters)
- Domain format for organization domain
- Required fields are not empty

## Cleanup

To remove all created users:

```bash
terraform destroy
```

**Warning**: This will permanently delete all user accounts created by this configuration.

## Related Resources

- [Umbrella Users Resource Documentation](../../docs/resources/users.md)
- [Cisco Umbrella API Documentation](https://docs.umbrella.com/umbrella-api/)
- [Terraform Umbrella Provider](https://registry.terraform.io/providers/mantisec/umbrella/latest)

## Support

For issues related to:
- **Terraform Provider**: Open an issue in the provider repository
- **Umbrella API**: Contact Cisco Umbrella support
- **Terraform**: Refer to Terraform documentation