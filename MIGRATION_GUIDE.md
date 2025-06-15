# Terraform Provider Umbrella - Migration Guide

## Table of Contents

1. [Migration Overview](#migration-overview)
2. [Pre-Migration Assessment](#pre-migration-assessment)
3. [Migration Planning](#migration-planning)
4. [Step-by-Step Migration Process](#step-by-step-migration-process)
5. [Backward Compatibility](#backward-compatibility)
6. [Validation and Testing](#validation-and-testing)
7. [Rollback Procedures](#rollback-procedures)
8. [Post-Migration Optimization](#post-migration-optimization)
9. [Common Migration Scenarios](#common-migration-scenarios)
10. [Troubleshooting Migration Issues](#troubleshooting-migration-issues)

## Migration Overview

### What This Migration Covers

This guide helps you transition from manual Terraform provider development to the automated code generation system. The migration process ensures:

- **Zero Downtime**: Existing resources continue to function during migration
- **State Preservation**: Terraform state remains intact and compatible
- **Feature Parity**: All existing functionality is maintained or enhanced
- **Gradual Transition**: Migrate resources incrementally at your own pace

### Migration Benefits

#### Before: Manual Development Challenges
- **Time-Intensive**: 2-4 weeks per resource
- **Inconsistent Patterns**: Varying code quality and structure
- **Manual Testing**: Incomplete or inconsistent test coverage
- **Documentation Lag**: Often outdated or missing
- **Maintenance Overhead**: Manual updates for API changes

#### After: Automated Code Generation Benefits
- **Rapid Development**: Minutes instead of weeks
- **Consistent Quality**: Standardized patterns and best practices
- **Comprehensive Testing**: Automated test generation
- **Always Current**: Auto-generated documentation
- **Reduced Maintenance**: Automatic updates from API changes

### Migration Types

1. **Full Migration**: Replace all manual resources with generated ones
2. **Incremental Migration**: Migrate resources one at a time
3. **Hybrid Approach**: Keep some manual resources, generate others
4. **Selective Migration**: Migrate only specific resource types

## Pre-Migration Assessment

### Current State Analysis

#### 1. Inventory Existing Resources

Create a comprehensive inventory of your current manual resources:

```bash
# List all manual resource files
find internal/provider -name "resource_*.go" -not -name "generated_*" | sort

# List all manual data source files
find internal/provider -name "data_source_*.go" -not -name "generated_*" | sort

# Count lines of manual code
find internal/provider -name "*.go" -not -name "generated_*" -exec wc -l {} + | tail -1
```

#### 2. Analyze Resource Complexity

For each resource, assess:
- **CRUD Operations**: Which operations are implemented
- **Schema Complexity**: Number and types of attributes
- **Custom Logic**: Business rules and validations
- **API Dependencies**: External API interactions
- **Test Coverage**: Existing test quality and coverage

#### 3. Identify Dependencies

Map dependencies between resources:
```hcl
# Example dependency analysis
resource "umbrella_destination_list" "blocked" {
  # Independent resource - can migrate first
}

resource "umbrella_rule" "block_rule" {
  destination_lists = [umbrella_destination_list.blocked.name]
  # Depends on destination_list - migrate after
}
```

### Compatibility Assessment

#### API Specification Availability

Check if OpenAPI specifications exist for your resources:

```bash
# List available API specifications
ls -la *.yaml *.yml

# Validate OpenAPI specifications
go run tools/generator/cmd/generate/main.go -validate-specs
```

#### Custom Logic Identification

Identify custom logic that may need special handling:

```go
// Example: Custom validation logic
func (r *CustomResource) validateSpecialRules(ctx context.Context, data *CustomResourceModel) error {
    // This custom logic needs to be preserved or migrated
    if data.SpecialField.ValueString() == "restricted" {
        return fmt.Errorf("restricted values not allowed")
    }
    return nil
}
```

## Migration Planning

### Migration Strategy Selection

#### Strategy 1: Big Bang Migration
- **Pros**: Complete transition, immediate benefits
- **Cons**: Higher risk, requires extensive testing
- **Best For**: Small providers with few resources

#### Strategy 2: Incremental Migration
- **Pros**: Lower risk, gradual transition
- **Cons**: Longer timeline, temporary complexity
- **Best For**: Large providers with many resources

#### Strategy 3: Hybrid Approach
- **Pros**: Flexibility, preserve critical customizations
- **Cons**: Ongoing maintenance of both systems
- **Best For**: Providers with unique requirements

### Migration Timeline

#### Phase 1: Preparation (1-2 weeks)
- [ ] Complete pre-migration assessment
- [ ] Set up code generation system
- [ ] Create backup of existing code
- [ ] Establish testing environment

#### Phase 2: Initial Migration (2-4 weeks)
- [ ] Migrate simple resources first
- [ ] Validate generated code quality
- [ ] Update tests and documentation
- [ ] Perform integration testing

#### Phase 3: Complex Resources (3-6 weeks)
- [ ] Migrate resources with custom logic
- [ ] Implement schema overrides as needed
- [ ] Validate business rules preservation
- [ ] Update advanced configurations

#### Phase 4: Validation and Cleanup (1-2 weeks)
- [ ] Comprehensive testing
- [ ] Performance validation
- [ ] Documentation updates
- [ ] Remove deprecated manual code

## Step-by-Step Migration Process

### Step 1: Environment Setup

#### 1.1 Install Code Generation System

```bash
# Clone or update the repository
git clone https://github.com/mantisec/terraform-provider-umbrella.git
cd terraform-provider-umbrella

# Install dependencies
go mod tidy

# Build the generator
make build-generator
```

#### 1.2 Configure Generation Settings

Create or update configuration files:

```yaml
# tools/generator/config/generation.yaml
global:
  provider_name: "umbrella"
  package_name: "provider"
  go_module: "github.com/mantisec/terraform-provider-umbrella"

resources:
  defaults:
    generate_crud: true
    generate_import: true
    generate_docs: true
```

#### 1.3 Set Up Advanced Configuration

```yaml
# tools/generator/config/advanced_config.yaml
schema_overrides:
  global:
    type_mappings:
      "string_with_format_date": "string"
    
    field_mappings:
      "id": "ID"
      "url": "URL"
```

### Step 2: Backup and Preparation

#### 2.1 Create Backup

```bash
# Create backup of existing code
mkdir -p migration_backup/$(date +%Y%m%d)
cp -r internal/provider migration_backup/$(date +%Y%m%d)/
cp -r docs migration_backup/$(date +%Y%m%d)/
```

#### 2.2 Create Migration Branch

```bash
# Create dedicated migration branch
git checkout -b migration-to-generated-code
git add .
git commit -m "Pre-migration backup"
```

### Step 3: Resource-by-Resource Migration

#### 3.1 Start with Simple Resources

Begin with resources that have:
- Simple schemas
- Standard CRUD operations
- No custom business logic
- Available OpenAPI specifications

```bash
# Generate code for simple resources first
go run tools/generator/cmd/generate/main.go -specs destination_lists_api.yaml
```

#### 3.2 Compare Generated vs Manual Code

```bash
# Compare schemas
diff internal/provider/resource_destination_list.go internal/provider/generated_destination_list_resource.go

# Compare functionality
diff internal/provider/data_source_destination_list.go internal/provider/generated_destination_list_data_source.go
```

#### 3.3 Migrate Resource Registration

Update provider registration to use generated resources:

```go
// internal/provider/provider.go

func (p *umbrellaProvider) Resources(_ context.Context) []func() resource.Resource {
    // Initialize generated resources
    InitializeGeneratedResources()
    
    resources := []func() resource.Resource{
        // Keep existing manual resources temporarily
        NewDestinationListResource, // Manual - to be removed
        
        // Add generated resources
        // NewGeneratedDestinationListResource, // Generated - uncomment when ready
    }
    
    // Add all generated resources
    resources = append(resources, GetGeneratedRegistry().GetResources()...)
    
    return resources
}
```

### Step 4: Handle Custom Logic

#### 4.1 Identify Custom Validation

Extract custom validation logic:

```go
// Before: Manual resource with custom validation
func (r *DestinationListResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
    // Custom validation
    if len(data.Destinations) > 1000 {
        resp.Diagnostics.AddError("Too many destinations", "Maximum 1000 destinations allowed")
        return
    }
    
    // Standard creation logic...
}
```

#### 4.2 Configure Schema Overrides

Add custom validation through configuration:

```yaml
# tools/generator/config/advanced_config.yaml
schema_overrides:
  resources:
    destination_list:
      schema_transforms:
        - field: "destinations"
          type: "list"
          element_type: "string"
          validation: "min_items=1,max_items=1000"
      
      validation_rules:
        - rule: "destinations_not_empty"
          message: "Destination list must contain at least one destination"
```

#### 4.3 Implement Custom Validators

For complex business logic, create custom validators:

```go
// tools/generator/validator/custom_business_validator.go
package validator

type BusinessRuleValidator struct {
    rules []BusinessRule
}

func (v *BusinessRuleValidator) ValidateDestinationList(data *DestinationListModel) []ValidationError {
    var errors []ValidationError
    
    // Apply custom business rules
    if data.Type.ValueString() == "DOMAIN" {
        for _, dest := range data.Destinations {
            if !isValidDomain(dest.ValueString()) {
                errors = append(errors, ValidationError{
                    Field:   "destinations",
                    Message: fmt.Sprintf("Invalid domain: %s", dest.ValueString()),
                })
            }
        }
    }
    
    return errors
}
```

### Step 5: Update Tests

#### 5.1 Migrate Existing Tests

Update test files to use generated resources:

```go
// Before: Manual test
func TestDestinationListResource(t *testing.T) {
    resource.Test(t, resource.TestCase{
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: testAccDestinationListResourceConfig("test-list"),
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckResourceAttr("umbrella_destination_list.test", "name", "test-list"),
                ),
            },
        },
    })
}
```

```go
// After: Generated resource test
func TestGeneratedDestinationListResource(t *testing.T) {
    resource.Test(t, resource.TestCase{
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: testAccGeneratedDestinationListResourceConfig("test-list"),
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckResourceAttr("umbrella_destination_list.test", "name", "test-list"),
                    // Additional generated test checks
                    resource.TestCheckResourceAttrSet("umbrella_destination_list.test", "id"),
                    resource.TestCheckResourceAttrSet("umbrella_destination_list.test", "created_at"),
                ),
            },
        },
    })
}
```

#### 5.2 Add Generated Tests

The code generation system automatically creates comprehensive tests:

```bash
# Generate tests for all resources
go run tools/generator/cmd/generate/main.go -tests-only

# Run generated tests
go test ./internal/provider/tests/...
```

### Step 6: Update Documentation

#### 6.1 Generate New Documentation

```bash
# Generate documentation for all resources
go run tools/generator/cmd/generate/main.go -docs-only

# Verify documentation quality
ls -la docs/resources/
```

#### 6.2 Update Examples

Update example configurations to use generated resources:

```hcl
# examples/basic/main.tf - Updated for generated resources
resource "umbrella_destination_list" "blocked_domains" {
  name = "Blocked Domains List"
  type = "DOMAIN"
  destinations = [
    "malicious-site.com",
    "phishing-domain.net"
  ]
  
  # New attributes available in generated version
  tags = {
    "Environment" = "Production"
    "Team"        = "Security"
  }
}
```

## Backward Compatibility

### State Compatibility

The generated resources maintain full compatibility with existing Terraform state:

#### State Structure Preservation
```json
{
  "version": 4,
  "terraform_version": "1.5.0",
  "serial": 1,
  "lineage": "existing-lineage",
  "outputs": {},
  "resources": [
    {
      "mode": "managed",
      "type": "umbrella_destination_list",
      "name": "example",
      "provider": "provider[\"registry.terraform.io/mantisec/umbrella\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "id": "12345",
            "name": "Example List",
            "type": "DOMAIN",
            "destinations": ["example.com"]
          }
        }
      ]
    }
  ]
}
```

#### Schema Version Management

Generated resources maintain schema compatibility:

```go
// Generated resource maintains same schema version
func (r *GeneratedDestinationListResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
    return map[int64]resource.StateUpgrader{
        // Maintain compatibility with existing state
        0: {
            PriorSchema: &schema.Schema{
                // Previous schema definition
            },
            StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
                // State upgrade logic if needed
            },
        },
    }
}
```

### API Compatibility

#### Endpoint Compatibility
Generated resources use the same API endpoints as manual resources:

```go
// Manual resource API call
resp, err := client.do(ctx, "POST", "/policies/v2/organizations/"+orgID+"/destinationlists", body)

// Generated resource API call (same endpoint)
resp, err := client.CreateDestinationList(ctx, payload)
```

#### Request/Response Format
Generated client methods maintain the same request/response formats:

```go
// Manual implementation
func (c *apiClient) createDestinationList(ctx context.Context, data DestinationListData) (*DestinationList, error) {
    // Manual implementation
}

// Generated implementation
func (c *GeneratedClient) CreateDestinationList(ctx context.Context, payload map[string]interface{}) (map[string]interface{}, error) {
    // Generated implementation with same API contract
}
```

## Validation and Testing

### Pre-Migration Testing

#### 1. Baseline Testing

Establish baseline performance and functionality:

```bash
# Run existing tests
go test ./internal/provider/... -v

# Measure performance
go test ./internal/provider/... -bench=. -benchmem

# Check test coverage
go test ./internal/provider/... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

#### 2. Integration Testing

Test with real Umbrella API:

```bash
# Set up test environment
export UMBRELLA_API_KEY="test-key"
export UMBRELLA_API_SECRET="test-secret"
export UMBRELLA_ORG_ID="test-org"

# Run acceptance tests
TF_ACC=1 go test ./internal/provider/... -v -timeout 30m
```

### Migration Testing

#### 1. Generated Code Validation

```bash
# Validate generated code compiles
go build ./internal/provider/generated_*.go

# Run generated tests
go test ./internal/provider/tests/... -v

# Validate generated documentation
ls -la docs/resources/generated_*.md
```

#### 2. Compatibility Testing

Test that generated resources work with existing state:

```hcl
# Test configuration using existing state
terraform {
  required_providers {
    umbrella = {
      source = "local/mantisec/umbrella"
      version = "0.2.0"
    }
  }
}

# Import existing resource to test compatibility
resource "umbrella_destination_list" "existing" {
  # Configuration should match existing state
}
```

```bash
# Import existing resource
terraform import umbrella_destination_list.existing 12345

# Verify no changes needed
terraform plan
```

### Post-Migration Testing

#### 1. Comprehensive Testing

```bash
# Full test suite
make test-all

# Performance testing
make test-performance

# Integration testing
make test-integration
```

#### 2. User Acceptance Testing

Create test scenarios that mirror real usage:

```hcl
# Test scenario: Complete workflow
resource "umbrella_destination_list" "test_migration" {
  name = "Migration Test List"
  type = "DOMAIN"
  destinations = [
    "test-domain.com",
    "example-site.net"
  ]
}

resource "umbrella_rule" "test_migration_rule" {
  ruleset_id        = data.umbrella_ruleset.default.id
  name              = "Migration Test Rule"
  action            = "BLOCK"
  rank              = 100
  destination_lists = [umbrella_destination_list.test_migration.name]
}
```

## Rollback Procedures

### Rollback Planning

#### 1. Rollback Triggers

Define conditions that require rollback:
- Generated code doesn't compile
- Tests fail after migration
- Performance degradation > 20%
- API compatibility issues
- State corruption or loss

#### 2. Rollback Checkpoints

Create rollback points at each migration phase:

```bash
# Create rollback checkpoint
git tag migration-checkpoint-phase1
git push origin migration-checkpoint-phase1
```

### Rollback Execution

#### 1. Code Rollback

```bash
# Rollback to previous version
git checkout migration-checkpoint-phase1

# Restore manual resources
cp -r migration_backup/$(date +%Y%m%d)/internal/provider/* internal/provider/

# Remove generated files
rm -f internal/provider/generated_*.go
```

#### 2. Provider Registration Rollback

```go
// internal/provider/provider.go
func (p *umbrellaProvider) Resources(_ context.Context) []func() resource.Resource {
    resources := []func() resource.Resource{
        // Restore manual resources
        NewDestinationListResource,
        NewTunnelResource,
        // Remove generated resource registration
        // InitializeGeneratedResources() - commented out
    }
    
    // Don't append generated resources
    // resources = append(resources, GetGeneratedRegistry().GetResources()...)
    
    return resources
}
```

#### 3. State Recovery

If state issues occur:

```bash
# Restore state backup
cp terraform.tfstate.backup terraform.tfstate

# Or restore from remote state
terraform state pull > terraform.tfstate.recovered
```

### Rollback Validation

#### 1. Functionality Testing

```bash
# Verify manual resources work
go test ./internal/provider/... -v

# Test with real API
TF_ACC=1 go test ./internal/provider/... -v -timeout 30m
```

#### 2. State Validation

```bash
# Verify state is intact
terraform plan

# Should show no changes if rollback successful
```

## Post-Migration Optimization

### Performance Optimization

#### 1. Enable Caching

```yaml
# tools/generator/config/advanced_config.yaml
performance:
  caching:
    enabled: true
    cache_dir: ".umbrella_cache"
    ttl: "5m"
```

#### 2. Parallel Processing

```yaml
performance:
  parallel_processing:
    enabled: true
    max_workers: 4
    batch_size: 10
```

### Code Quality Improvements

#### 1. Enable Advanced Validation

```yaml
validation:
  enabled_validators:
    - "schema_consistency"
    - "api_compatibility"
    - "terraform_compliance"
    - "security_checks"
```

#### 2. Custom Quality Rules

```yaml
quality_assurance:
  code_quality:
    enabled_checks:
      - "gofmt"
      - "golint"
      - "govet"
      - "ineffassign"
      - "misspell"
```

### Monitoring and Observability

#### 1. Enable Metrics

```go
// Add metrics collection
type MigrationMetrics struct {
    ResourcesGenerated int
    GenerationTime     time.Duration
    TestCoverage       float64
    ErrorRate          float64
}
```

#### 2. Health Checks

```go
// Add health monitoring
func (p *Provider) HealthCheck() error {
    // Verify generated resources are working
    // Check API connectivity
    // Validate configuration
    return nil
}
```

## Common Migration Scenarios

### Scenario 1: Simple Resource Migration

**Situation**: Basic CRUD resource with standard schema

**Migration Steps**:
1. Verify OpenAPI specification exists
2. Generate resource code
3. Compare with manual implementation
4. Update provider registration
5. Run tests and validate

**Example**:
```bash
# Generate destination list resource
go run tools/generator/cmd/generate/main.go -specs destination_lists_api.yaml

# Compare implementations
diff internal/provider/resource_destination_list.go internal/provider/generated_destination_list_resource.go

# Update registration and test
```

### Scenario 2: Complex Resource with Custom Logic

**Situation**: Resource with business rules and custom validation

**Migration Steps**:
1. Identify custom logic components
2. Configure schema overrides
3. Implement custom validators
4. Generate with advanced configuration
5. Validate business rules preservation

**Example**:
```yaml
# Configure custom validation
schema_overrides:
  resources:
    complex_resource:
      validation_rules:
        - rule: "custom_business_rule"
          message: "Business rule violation"
```

### Scenario 3: Resource with Dependencies

**Situation**: Resources that depend on each other

**Migration Steps**:
1. Map dependency graph
2. Migrate in dependency order
3. Update references between resources
4. Test dependency chains
5. Validate state relationships

**Example**:
```hcl
# Migrate destination_list first (no dependencies)
resource "umbrella_destination_list" "example" { }

# Then migrate rule (depends on destination_list)
resource "umbrella_rule" "example" {
  destination_lists = [umbrella_destination_list.example.name]
}
```

## Troubleshooting Migration Issues

### Common Issues and Solutions

#### Issue 1: Generated Code Doesn't Compile

**Symptoms**:
```
./internal/provider/generated_resource.go:45:2: undefined: SomeType
```

**Solutions**:
1. Check import statements in generated code
2. Verify OpenAPI specification completeness
3. Update schema overrides for custom types
4. Regenerate with updated configuration

```yaml
# Fix missing imports
code_generation:
  imports:
    additional:
      - "custom/package/path"
```

#### Issue 2: Schema Validation Failures

**Symptoms**:
```
Error: Invalid attribute type
```

**Solutions**:
1. Review OpenAPI schema definitions
2. Add type mappings for custom formats
3. Configure field transformations
4. Update validation rules

```yaml
# Fix type mapping
schema_overrides:
  global:
    type_mappings:
      "custom_format": "string"
```

#### Issue 3: API Compatibility Issues

**Symptoms**:
```
Error: 400 Bad Request - Invalid field format
```

**Solutions**:
1. Compare API request formats
2. Check field name mappings
3. Verify request body structure
4. Update client method generation

```yaml
# Fix field mapping
schema_overrides:
  global:
    field_mappings:
      "api_field_name": "TerraformFieldName"
```

#### Issue 4: State Compatibility Problems

**Symptoms**:
```
Error: Provider produced inconsistent result after apply
```

**Solutions**:
1. Check schema version compatibility
2. Verify attribute names match
3. Implement state upgraders if needed
4. Test with existing state files

```go
// Add state upgrader if needed
func (r *Resource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
    return map[int64]resource.StateUpgrader{
        0: {
            StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
                // Upgrade logic
            },
        },
    }
}
```

### Debug Procedures

#### 1. Enable Debug Logging

```bash
export TF_LOG=DEBUG
export TF_LOG_PATH=migration-debug.log
terraform apply
```

#### 2. Compare Generated vs Manual

```bash
# Compare resource implementations
diff -u internal/provider/resource_manual.go internal/provider/generated_resource.go

# Compare API calls
grep -n "client\." internal/provider/resource_*.go
```

#### 3. Validate Configuration

```bash
# Check configuration syntax
go run tools/generator/cmd/generate/main.go -config-check

# Validate OpenAPI specs
go run tools/generator/cmd/generate/main.go -validate-specs
```

### Getting Help

#### 1. Documentation Resources
- Review [TECHNICAL_REFERENCE.md](TECHNICAL_REFERENCE.md) for implementation details
- Check [COMPLETE_USER_GUIDE.md](COMPLETE_USER_GUIDE.md) for usage examples
- Consult [docs/GENERATOR_GUIDE.md](docs/GENERATOR_GUIDE.md) for generation specifics

#### 2. Community Support
- Create GitHub issues for bugs and problems
- Use GitHub Discussions for questions and guidance
- Include debug logs and configuration when seeking help

#### 3. Professional Support
- Contact maintainers for complex migration scenarios
- Consider professional services for large-scale migrations
- Engage with the community for best practices sharing

This migration guide provides comprehensive guidance for transitioning from manual Terraform provider development to the automated code generation system, ensuring a smooth and successful migration process.