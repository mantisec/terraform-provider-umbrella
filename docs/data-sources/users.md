---
page_title: "umbrella_users Data Source - terraform-provider-umbrella"
subcategory: "Admin"
description: |-
  Use this data source to retrieve information about user accounts in your Cisco Umbrella organization.
---

# umbrella_users (Data Source)

Use this data source to retrieve information about user accounts in your Cisco Umbrella organization. Users represent individual accounts with specific roles and permissions within your organization.

## Example Usage

### Get a Specific User by Email

```terraform
data "umbrella_users" "admin_user" {
  email = "admin@example.com"
}

output "admin_user_info" {
  value = {
    user_id             = data.umbrella_users.admin_user.user_id
    firstname           = data.umbrella_users.admin_user.firstname
    lastname            = data.umbrella_users.admin_user.lastname
    role                = data.umbrella_users.admin_user.role
    role_id             = data.umbrella_users.admin_user.role_id
    status              = data.umbrella_users.admin_user.status
    two_factor_enabled  = data.umbrella_users.admin_user.two_factor_enabled
    last_login_time     = data.umbrella_users.admin_user.last_login_time
  }
}
```

### Get a User by ID

```terraform
data "umbrella_users" "example" {
  id = "12345"
}

# Use user data for conditional logic
locals {
  user_is_admin = data.umbrella_users.example.role_id == 1
}
```

### Get All Users

```terraform
data "umbrella_users" "all" {}

# Filter users by role
locals {
  admin_users = [
    for user in data.umbrella_users.all.users : user
    if user.role_id == 1
  ]
  
  readonly_users = [
    for user in data.umbrella_users.all.users : user
    if user.role_id == 2
  ]
}

output "user_summary" {
  value = {
    total_users    = length(data.umbrella_users.all.users)
    admin_users    = length(local.admin_users)
    readonly_users = length(local.readonly_users)
  }
}
```

## Schema

### Optional

- `id` (String) The unique identifier for a specific user. If provided, returns information for this user only.
- `email` (String) The email address of a specific user. If provided, returns information for the user with this email.

### Read-Only

When querying a specific user (by `id` or `email`):

- `user_id` (Number) The numeric user ID
- `firstname` (String) The user's first name
- `lastname` (String) The user's last name
- `timezone` (String) The user's timezone
- `role` (String) The user's role name
- `role_id` (Number) The role ID assigned to the user
- `status` (String) The user's status (e.g., 'on', 'off')
- `two_factor_enabled` (Boolean) Whether two-factor authentication is enabled
- `last_login_time` (String) The user's last login date and time (ISO8601 timestamp)

When querying all users (no filters):

- `users` (List of Object) List of all users with the following attributes:
  - `id` (String) The unique identifier for the user
  - `user_id` (Number) The numeric user ID
  - `email` (String) The user's email address
  - `firstname` (String) The user's first name
  - `lastname` (String) The user's last name
  - `timezone` (String) The user's timezone
  - `role` (String) The user's role name
  - `role_id` (Number) The role ID assigned to the user
  - `status` (String) The user's status
  - `two_factor_enabled` (Boolean) Whether 2FA is enabled
  - `last_login_time` (String) Last login timestamp

## Usage Examples

### User Role Analysis

```terraform
data "umbrella_users" "all" {}

locals {
  # Group users by role
  users_by_role = {
    for user in data.umbrella_users.all.users : user.role => user...
  }
  
  # Count users by status
  active_users = [
    for user in data.umbrella_users.all.users : user
    if user.status == "on"
  ]
  
  # Users with 2FA enabled
  secure_users = [
    for user in data.umbrella_users.all.users : user
    if user.two_factor_enabled == true
  ]
}

output "user_analytics" {
  value = {
    users_by_role     = { for role, users in local.users_by_role : role => length(users) }
    active_users      = length(local.active_users)
    secure_users      = length(local.secure_users)
    total_users       = length(data.umbrella_users.all.users)
  }
}
```

### Security Compliance Monitoring

```terraform
data "umbrella_users" "all" {}

locals {
  # Identify security compliance issues
  inactive_users = [
    for user in data.umbrella_users.all.users : user
    if user.status != "on"
  ]
  
  users_without_2fa = [
    for user in data.umbrella_users.all.users : user
    if user.two_factor_enabled == false && user.role_id == 1  # Admin users
  ]
  
  # Users who haven't logged in recently (example: 90 days)
  stale_users = [
    for user in data.umbrella_users.all.users : user
    if user.last_login_time != null && 
       timecmp(timeadd(timestamp(), "-2160h"), user.last_login_time) > 0  # 90 days
  ]
}

output "security_compliance" {
  value = {
    inactive_users_count    = length(local.inactive_users)
    admins_without_2fa     = length(local.users_without_2fa)
    stale_users_count      = length(local.stale_users)
    
    # Detailed lists for remediation
    inactive_users         = [for u in local.inactive_users : u.email]
    admins_without_2fa     = [for u in local.users_without_2fa : u.email]
    stale_users           = [for u in local.stale_users : u.email]
  }
}
```

### User Provisioning Validation

```terraform
# Check if required users exist
data "umbrella_users" "all" {}

locals {
  required_admin_emails = [
    "admin1@company.com",
    "admin2@company.com",
    "security@company.com"
  ]
  
  existing_user_emails = [for user in data.umbrella_users.all.users : user.email]
  missing_admin_users  = setsubtract(toset(local.required_admin_emails), toset(local.existing_user_emails))
}

# Create missing admin users
resource "umbrella_users" "missing_admins" {
  for_each = local.missing_admin_users
  
  email     = each.value
  firstname = split("@", each.value)[0]  # Use email prefix as firstname
  lastname  = "Admin"
  password  = var.default_admin_password
  role_id   = 1  # Full Admin
  timezone  = "UTC"
}

output "user_provisioning_status" {
  value = {
    required_users = local.required_admin_emails
    missing_users  = local.missing_admin_users
    created_users  = keys(umbrella_users.missing_admins)
  }
}
```

### Role-Based Access Control Audit

```terraform
data "umbrella_users" "all" {}

locals {
  # Define role mappings for clarity
  role_names = {
    1 = "Full Admin"
    2 = "Read Only"
    3 = "Block Page Bypass"
    4 = "Reporting Only"
  }
  
  # Audit user roles
  role_distribution = {
    for role_id, role_name in local.role_names :
    role_name => [
      for user in data.umbrella_users.all.users : user
      if user.role_id == role_id
    ]
  }
}

output "rbac_audit" {
  value = {
    for role_name, users in local.role_distribution :
    role_name => {
      count = length(users)
      users = [for u in users : "${u.firstname} ${u.lastname} (${u.email})"]
    }
  }
}
```

## API Reference

This data source uses the following Cisco Umbrella API endpoints:

- **Get User**: `GET /admin/v2/organizations/{orgId}/users/{userId}` (when using `id`)
- **List Users**: `GET /admin/v2/organizations/{orgId}/users` (when using `email` or no filters)

## Common Role IDs

The following are common role IDs in Umbrella organizations:

- `1` - Full Admin
- `2` - Read Only
- `3` - Block Page Bypass
- `4` - Reporting Only

**Note**: Role IDs may vary between organizations. Use the Umbrella console or API to verify the correct role IDs for your organization.

## Notes

- **Email Uniqueness**: Email addresses are unique across the organization
- **Case Sensitivity**: Email addresses are case-insensitive when filtering
- **Password Security**: User passwords are never returned by the API for security reasons
- **Status Values**: Common status values include 'on' (active) and 'off' (inactive)
- **Timezone Format**: Timezones follow standard timezone identifiers (e.g., "America/New_York", "UTC")

## Troubleshooting

### Common Errors

**User not found**
```
Error: user not found with email "user@example.com"
```
Verify the email address exists and is spelled correctly.

**Permission denied**
```
Error: insufficient permissions: 403 Forbidden
```
Ensure your API credentials have the `admin.users:read` scope.

**Invalid role ID reference**
```
Error: role ID 5 does not exist
```
Verify role IDs exist in your organization through the Umbrella console.

### Best Practices

1. **Security Monitoring**: Regularly audit user roles and permissions
2. **2FA Enforcement**: Monitor and encourage two-factor authentication adoption
3. **Access Reviews**: Periodically review user access and remove inactive accounts
4. **Role Validation**: Verify role IDs are appropriate for user responsibilities
5. **Compliance Tracking**: Use data sources to maintain security compliance records

## Security Considerations

- **Sensitive Data**: User information should be handled according to your organization's privacy policies
- **Access Logging**: Monitor who accesses user data through Terraform operations
- **Role Segregation**: Ensure users have appropriate roles based on the principle of least privilege
- **Regular Audits**: Use this data source for regular security and compliance audits

## Related Resources

- [`umbrella_users`](../resources/users.md) - Create and manage user accounts
- Role management resources (when available) for role definitions
- Audit logging and compliance reporting tools