---
page_title: "Migration Guide - Umbrella Provider"
subcategory: "Guides"
description: |-
  Guide for migrating to the Terraform Provider for Cisco Umbrella from other management methods.
---

# Migration Guide

This guide helps you migrate to the Terraform Provider for Cisco Umbrella from various existing management approaches, including manual configurations, scripts, and other automation tools.

## Migration Scenarios

### 1. From Manual Dashboard Management

If you've been managing Umbrella resources through the web dashboard, this section will help you transition to Infrastructure as Code.

#### Step 1: Inventory Existing Resources

First, document your existing resources:

```terraform
# Use data sources to discover existing resources
data "umbrella_sites" "all" {}
data "umbrella_networks" "all" {}
data "umbrella_users" "all" {}
data "umbrella_destination_list" "all" {}

# Output current state for review
output "current_inventory" {
  value = {
    sites = [
      for site in data.umbrella_sites.all.sites : {
        id   = site.id
        name = site.name
        type = site.type
      }
    ]
    networks = [
      for network in data.umbrella_networks.all.networks : {
        id         = network.id
        name       = network.name
        ip_address = network.ip_address
        status     = network.status
      }
    ]
    users = [
      for user in data.umbrella_users.all.users : {
        id    = user.id
        email = user.email
        role  = user.role
      }
    ]
    destination_lists = [
      for list in data.umbrella_destination_list.all.destination_lists : {
        id     = list.id
        name   = list.name
        access = list.access
      }
    ]
  }
}
```

#### Step 2: Import Existing Resources

Import your existing resources into Terraform state:

```bash
# Import sites
terraform import umbrella_sites.main_office 12345
terraform import umbrella_sites.branch_office 12346

# Import networks (if any)
terraform import umbrella_networks.corporate_network 67890

# Import destination lists
terraform import umbrella_destination_list.blocked_domains 11111
terraform import umbrella_destination_list.allowed_sites 11112

# Import users
terraform import umbrella_users.admin_user 22222
terraform import umbrella_users.readonly_user 22223
```

#### Step 3: Create Terraform Configuration

After importing, create corresponding resource configurations:

```terraform
# Sites
resource "umbrella_sites" "main_office" {
  name       = "Main Office"
  is_default = true
}

resource "umbrella_sites" "branch_office" {
  name = "Branch Office"
}

# Destination Lists
resource "umbrella_destination_list" "blocked_domains" {
  name   = "Blocked Domains"
  access = "block"
  destinations = [
    "malicious-site.com",
    "phishing-domain.net"
  ]
}

# Users
resource "umbrella_users" "admin_user" {
  email     = "admin@company.com"
  firstname = "Admin"
  lastname  = "User"
  password  = var.admin_password
  role_id   = 1
  timezone  = "America/New_York"
}
```

### 2. From curl/API Scripts

If you've been using curl commands or custom scripts to manage Umbrella resources, here's how to migrate.

#### Before: curl-based Management

```bash
# Old curl-based approach
curl -X POST \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  "https://api.umbrella.com/policies/v2/organizations/$ORG_ID/destinationlists" \
  -d '{
    "name": "Blocked Sites",
    "access": "block",
    "destinations": ["malicious.com", "phishing.net"]
  }'
```

#### After: Terraform Resource

```terraform
# New Terraform approach
resource "umbrella_destination_list" "blocked_sites" {
  name   = "Blocked Sites"
  access = "block"
  destinations = [
    "malicious.com",
    "phishing.net"
  ]
}
```

#### Migration Script Example

Create a migration script to convert your existing curl commands:

```bash
#!/bin/bash
# migration-helper.sh

# Extract existing destination lists
curl -H "Authorization: Bearer $TOKEN" \
  "https://api.umbrella.com/policies/v2/organizations/$ORG_ID/destinationlists" \
  | jq -r '.[] | "terraform import umbrella_destination_list.\(.name | gsub("[^a-zA-Z0-9_]"; "_")) \(.id)"' \
  > import_commands.sh

# Execute imports
chmod +x import_commands.sh
./import_commands.sh
```

### 3. From Other Terraform Providers

If you're migrating from a different Umbrella provider or custom modules:

#### Step 1: State Migration

```bash
# Remove old provider resources from state
terraform state rm old_provider_resource.example

# Import with new provider
terraform import umbrella_sites.example 12345
```

#### Step 2: Update Resource Names

Update your configuration to use the new provider's resource names:

```terraform
# Old provider (example)
resource "old_umbrella_site" "example" {
  name = "Example Site"
}

# New provider
resource "umbrella_sites" "example" {
  name = "Example Site"
}
```

### 4. From Configuration Management Tools

If you're using Ansible, Puppet, or other configuration management tools:

#### Ansible to Terraform Migration

**Before (Ansible):**
```yaml
- name: Create Umbrella destination list
  uri:
    url: "https://api.umbrella.com/policies/v2/organizations/{{ org_id }}/destinationlists"
    method: POST
    headers:
      Authorization: "Bearer {{ token }}"
    body_format: json
    body:
      name: "Blocked Domains"
      access: "block"
      destinations: ["malicious.com"]
```

**After (Terraform):**
```terraform
resource "umbrella_destination_list" "blocked_domains" {
  name   = "Blocked Domains"
  access = "block"
  destinations = ["malicious.com"]
}
```

## Migration Strategies

### 1. Big Bang Migration

Migrate all resources at once:

**Pros:**
- Complete migration in one operation
- Consistent state across all resources
- Immediate benefits of IaC

**Cons:**
- Higher risk
- Requires significant planning
- Potential for extended downtime

**Best for:** Small to medium deployments with simple configurations

### 2. Gradual Migration

Migrate resources incrementally:

**Pros:**
- Lower risk
- Easier to troubleshoot issues
- Can validate each step

**Cons:**
- Longer migration timeline
- Mixed management approaches during transition
- More complex coordination

**Best for:** Large deployments with complex interdependencies

### 3. Parallel Migration

Run both systems in parallel:

**Pros:**
- Zero downtime
- Easy rollback
- Thorough testing possible

**Cons:**
- Resource duplication
- Complex synchronization
- Higher costs during transition

**Best for:** Critical production environments

## Migration Checklist

### Pre-Migration

- [ ] **Inventory existing resources** using data sources
- [ ] **Document current configurations** and dependencies
- [ ] **Set up Terraform environment** with proper backend
- [ ] **Test API credentials** and permissions
- [ ] **Create backup** of current configurations
- [ ] **Plan migration order** based on dependencies

### During Migration

- [ ] **Import resources** in dependency order
- [ ] **Create Terraform configurations** matching existing resources
- [ ] **Validate configurations** with `terraform plan`
- [ ] **Test in non-production** environment first
- [ ] **Monitor for drift** between old and new systems
- [ ] **Update documentation** and runbooks

### Post-Migration

- [ ] **Verify all resources** are under Terraform management
- [ ] **Remove old management tools** and scripts
- [ ] **Set up CI/CD pipelines** for Terraform
- [ ] **Train team members** on new workflows
- [ ] **Establish monitoring** and alerting
- [ ] **Document new processes** and procedures

## Common Migration Challenges

### 1. Resource Dependencies

**Challenge:** Umbrella resources often have dependencies (e.g., internal networks depend on sites).

**Solution:** Import resources in dependency order:

```bash
# Import in correct order
terraform import umbrella_sites.main_office 12345
terraform import umbrella_internalnetworks.office_network 67890
```

### 2. Sensitive Data

**Challenge:** Passwords and API keys in existing configurations.

**Solution:** Use Terraform variables and secure storage:

```terraform
variable "user_passwords" {
  description = "User passwords"
  type        = map(string)
  sensitive   = true
}

resource "umbrella_users" "users" {
  for_each = var.user_passwords
  
  email     = each.key
  password  = each.value
  # ... other attributes
}
```

### 3. Resource Naming

**Challenge:** Existing resource names may not follow Terraform naming conventions.

**Solution:** Use consistent naming patterns:

```terraform
# Use descriptive, consistent names
resource "umbrella_sites" "main_office_site" {
  name = "Main Office"  # Display name can be different
}

resource "umbrella_sites" "branch_office_site" {
  name = "Branch Office"
}
```

### 4. State Drift

**Challenge:** Resources modified outside Terraform after import.

**Solution:** Regular drift detection:

```bash
# Check for drift
terraform plan

# Refresh state
terraform refresh

# Import changes if needed
terraform import umbrella_sites.example 12345
```

## Migration Tools and Scripts

### Resource Discovery Script

```bash
#!/bin/bash
# discover-resources.sh

ORG_ID="your-org-id"
TOKEN="your-token"

echo "Discovering Umbrella resources..."

# Get sites
echo "Sites:"
curl -s -H "Authorization: Bearer $TOKEN" \
  "https://api.umbrella.com/deployments/v2/sites" \
  | jq -r '.[] | "ID: \(.siteId), Name: \(.name)"'

# Get destination lists
echo "Destination Lists:"
curl -s -H "Authorization: Bearer $TOKEN" \
  "https://api.umbrella.com/policies/v2/organizations/$ORG_ID/destinationlists" \
  | jq -r '.[] | "ID: \(.id), Name: \(.name), Access: \(.access)"'
```

### Import Generator Script

```bash
#!/bin/bash
# generate-imports.sh

# Generate import commands for sites
echo "# Site imports"
curl -s -H "Authorization: Bearer $TOKEN" \
  "https://api.umbrella.com/deployments/v2/sites" \
  | jq -r '.[] | "terraform import umbrella_sites.\(.name | gsub("[^a-zA-Z0-9_]"; "_") | ascii_downcase) \(.siteId)"'

# Generate import commands for destination lists
echo "# Destination list imports"
curl -s -H "Authorization: Bearer $TOKEN" \
  "https://api.umbrella.com/policies/v2/organizations/$ORG_ID/destinationlists" \
  | jq -r '.[] | "terraform import umbrella_destination_list.\(.name | gsub("[^a-zA-Z0-9_]"; "_") | ascii_downcase) \(.id)"'
```

## Best Practices for Migration

### 1. Start Small

Begin with non-critical resources to gain experience:

```terraform
# Start with simple resources
resource "umbrella_sites" "test_site" {
  name = "Test Site"
}
```

### 2. Use Version Control

Track all changes during migration:

```bash
git add .
git commit -m "Initial Terraform configuration for Umbrella resources"
```

### 3. Test Thoroughly

Validate configurations before applying:

```bash
terraform validate
terraform plan
terraform apply -auto-approve=false
```

### 4. Document Everything

Keep detailed records of the migration process:

```markdown
# Migration Log

## 2024-01-15
- Imported main office site (ID: 12345)
- Created Terraform configuration for site
- Validated with terraform plan

## 2024-01-16
- Imported destination lists
- Updated configurations to match existing settings
```

## Rollback Procedures

If you need to rollback during migration:

### 1. Remove from Terraform State

```bash
terraform state rm umbrella_sites.example
```

### 2. Restore Original Management

Resume using your previous management method until issues are resolved.

### 3. Investigate and Fix

Identify and resolve the issues that caused the rollback need.

## Getting Help

If you encounter issues during migration:

1. **Check the documentation** for specific resource requirements
2. **Review examples** in the provider repository
3. **Open an issue** on the provider repository with migration-specific details
4. **Consult Cisco support** for API-related questions

## Conclusion

Migrating to the Terraform Provider for Cisco Umbrella provides significant benefits in terms of consistency, repeatability, and infrastructure management. While the migration process requires careful planning and execution, the long-term benefits of Infrastructure as Code make it a worthwhile investment.

Remember to:
- Plan thoroughly before starting
- Test in non-production environments
- Migrate incrementally when possible
- Keep detailed documentation
- Have rollback procedures ready

With proper planning and execution, your migration to Terraform-managed Umbrella resources will provide a solid foundation for scalable, maintainable infrastructure management.