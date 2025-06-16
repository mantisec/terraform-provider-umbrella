# Sites Resource Example

This example demonstrates how to use the `umbrella_sites` resource to manage sites in Cisco Umbrella.

## Overview

Sites in Cisco Umbrella represent organizational locations and serve as foundational resources for other Umbrella components like internal networks and virtual appliances.

## Prerequisites

Before running this example, ensure you have:

1. **Terraform installed** (version 1.0 or later)
2. **Cisco Umbrella API credentials** with appropriate permissions
3. **Environment variables set** for authentication

## Authentication

Set the following environment variables:

```bash
export UMBRELLA_API_KEY="your-api-key"
export UMBRELLA_API_SECRET="your-api-secret"
export UMBRELLA_ORG_ID="your-organization-id"
```

Or on Windows:

```cmd
set UMBRELLA_API_KEY=your-api-key
set UMBRELLA_API_SECRET=your-api-secret
set UMBRELLA_ORG_ID=your-organization-id
```

## Required API Permissions

Your API credentials must have the following scopes:
- `deployments.sites:read` - For reading site information
- `deployments.sites:write` - For creating, updating, and deleting sites

## Usage

1. **Initialize Terraform:**
   ```bash
   terraform init
   ```

2. **Plan the deployment:**
   ```bash
   terraform plan
   ```

3. **Apply the configuration:**
   ```bash
   terraform apply
   ```

4. **View outputs:**
   ```bash
   terraform output
   ```

## What This Example Creates

This example creates:

1. **Main Office Site** - A primary site for the main office
2. **Branch Office Site** - A branch office site in New York
3. **Multiple Office Sites** - Three additional office sites using `for_each`

## Resources Created

- `umbrella_sites.main_office` - Main office site
- `umbrella_sites.branch_office` - Branch office site  
- `umbrella_sites.offices["San Francisco Office"]` - San Francisco office
- `umbrella_sites.offices["London Office"]` - London office
- `umbrella_sites.offices["Tokyo Office"]` - Tokyo office

## Outputs

The example provides several outputs:

- `main_office_id` - The unique identifier of the main office site
- `main_office_site_id` - The site ID of the main office
- `all_office_sites` - Information about all office sites created

## Customization

You can customize this example by:

1. **Modifying site names** in the `toset()` function
2. **Adding more sites** by extending the configuration
3. **Setting different default sites** using the `is_default` attribute
4. **Using variables** for dynamic site creation

### Example with Variables

Create a `variables.tf` file:

```hcl
variable "office_locations" {
  description = "List of office locations"
  type        = list(string)
  default     = ["San Francisco", "London", "Tokyo"]
}

variable "main_office_name" {
  description = "Name of the main office"
  type        = string
  default     = "Corporate Headquarters"
}
```

Then modify `main.tf` to use these variables:

```hcl
resource "umbrella_sites" "main_office" {
  name       = var.main_office_name
  is_default = true
}

resource "umbrella_sites" "offices" {
  for_each = toset(var.office_locations)
  name     = "${each.value} Office"
}
```

## Cleanup

To destroy the created resources:

```bash
terraform destroy
```

## Troubleshooting

### Common Issues

1. **Authentication Errors**
   - Verify your API credentials are correct
   - Ensure the organization ID is valid
   - Check that your API key has the required scopes

2. **Site Name Conflicts**
   - Site names must be unique within your organization
   - If you get a conflict error, choose different names

3. **Permission Errors**
   - Ensure your API credentials have `deployments.sites:write` permissions
   - Contact your Umbrella administrator if you need additional permissions

### Validation Errors

- **Site name too long**: Site names must be 255 characters or less
- **Empty site name**: Site names cannot be empty
- **Invalid characters**: While most characters are allowed, avoid special characters that might cause issues

## Next Steps

After creating sites, you can:

1. **Create internal networks** and associate them with sites
2. **Deploy virtual appliances** to sites
3. **Configure policies** for site-specific rules
4. **Set up monitoring** for site resources

## Related Resources

- `umbrella_networks` - For creating internal networks within sites
- `umbrella_destination_list` - For managing destination lists used in policies

## Support

For issues with this example:
1. Check the [provider documentation](../../docs/resources/sites.md)
2. Review the [troubleshooting guide](../../docs/resources/sites.md#troubleshooting)
3. Open an issue in the provider repository