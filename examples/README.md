# Terraform Provider Umbrella - Examples

This directory contains example Terraform configurations demonstrating how to use the Umbrella provider.

## Directory Structure

- [`basic/`](basic/) - Basic usage examples with destination lists and tunnels
- [`complete/`](complete/) - Complete configuration with individual destination management
- [`sso/`](sso/) - SAML SSO configuration examples

## Prerequisites

1. **Umbrella API Credentials**: You need valid Umbrella API credentials with appropriate permissions
2. **Terraform**: Install Terraform 1.0 or later
3. **Provider**: The examples use the published provider from the Terraform Registry

## Quick Start

1. **Choose an example directory** based on your use case
2. **Copy the terraform.tfvars.example** file to `terraform.tfvars`
3. **Update the variables** with your actual Umbrella credentials and configuration
4. **Initialize and apply**:
   ```bash
   cd examples/basic
   terraform init
   terraform plan
   terraform apply
   ```

## Configuration Variables

All examples require these basic variables:

```hcl
# terraform.tfvars
umbrella_api_key    = "your-api-key-here"
umbrella_api_secret = "your-api-secret-here"
umbrella_org_id     = "your-org-id-here"
```

Additional variables may be required depending on the example.

## Security Best Practices

- **Never commit credentials** to version control
- **Use environment variables** for sensitive values:
  ```bash
  export TF_VAR_umbrella_api_key="your-api-key"
  export TF_VAR_umbrella_api_secret="your-api-secret"
  export TF_VAR_umbrella_org_id="your-org-id"
  ```
- **Use Terraform Cloud/Enterprise** workspace variables for team environments
- **Rotate API credentials** regularly

## Example Descriptions

### Basic Example
- Creates multiple destination lists (domains, URLs, CIDR blocks)
- Sets up primary and secondary IPSec tunnels
- Demonstrates basic resource configuration and outputs

### Complete Example  
- Shows individual destination management within lists
- More granular control over destination entries
- Advanced configuration patterns

### SSO Example
- SAML authentication configuration
- Integration with identity providers (Azure AD, ADFS)
- Policy rules for SSO bypass scenarios

## Validation

You can validate all examples using the Makefile in the root directory:

```bash
make validate-examples
```

Or validate individually:

```bash
cd examples/basic
terraform init
terraform validate
```

## Troubleshooting

### Common Issues

1. **Authentication Errors**
   - Verify your API credentials are correct
   - Ensure the organization ID matches your Umbrella account
   - Check that your API key has the necessary permissions

2. **Resource Conflicts**
   - Resource names must be unique within your organization
   - Use `terraform import` for existing resources

3. **Network Connectivity**
   - Ensure your environment can reach `api.umbrella.com`
   - Check firewall rules if running from restricted networks

### Getting Help

- Review the [provider documentation](../docs/)
- Check the [migration guide](../docs/guides/migration.md) if migrating from curl-based configurations
- Open an issue on the provider repository for bugs or feature requests

## Contributing

When adding new examples:

1. Create a new directory with a descriptive name
2. Include a complete `main.tf` with provider configuration
3. Add a `terraform.tfvars.example` file with sample values
4. Document any special requirements or considerations
5. Test the example thoroughly before submitting