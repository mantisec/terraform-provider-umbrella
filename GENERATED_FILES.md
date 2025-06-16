# Generated Files Documentation

This document explains which files are automatically generated from the OpenAPI specifications and how to manage them.

## Overview

The Terraform Provider Umbrella uses a sophisticated code generation system that automatically creates Terraform resources, data sources, client methods, tests, and documentation from OpenAPI specifications located in the `api-specs/` directory.

## Generated File Categories

### 1. Generated Go Files (`internal/provider/`)

These files are automatically generated from the API specifications:

#### **Core Generated Files**
- `generated_client.go` - Enhanced API client with caching and error handling
- `generated_registry.go` - Automatic resource registration system

#### **Resource-Specific Generated Files**
Pattern: `generated_*_methods.go` and `generated_*_resource.go`

Currently generated:
- `generated_destination_list_methods.go` - Destination list API client methods
- `generated_destination_list_resource.go` - Destination list Terraform resource
- `generated_internal_network_methods.go` - Internal networks API client methods
- `generated_network_methods.go` - Networks API client methods
- `generated_site_methods.go` - Sites API client methods
- `generated_user_methods.go` - Users API client methods

### 2. Generated Test Files (`internal/provider/tests/`)

Pattern: `*_test.go` (except `provider_test.go` which is manually maintained)

Currently generated:
- `destination_list_test.go` - Destination list resource tests
- `internal_networks_test.go` - Internal networks resource tests
- `networks_test.go` - Networks resource tests
- `sites_test.go` - Sites resource tests
- `users_test.go` - Users resource tests

**Note**: `provider_test.go` is manually maintained and should NOT be deleted.

### 3. Generated Documentation (`docs/resources/`)

Pattern: `*.md` files in `docs/resources/`

Currently generated:
- `destination_list.md` - Destination list resource documentation
- `internalnetworks.md` - Internal networks resource documentation
- `networks.md` - Networks resource documentation
- `sites.md` - Sites resource documentation
- `users.md` - Users resource documentation

**Note**: `docs/index.md` is manually maintained and should NOT be deleted.

## Manual vs Generated Files

### ✅ **Safe to Delete (Generated)**
These files are automatically regenerated when you run `make generate-full`:

```
internal/provider/generated_*.go
internal/provider/tests/*_test.go (except provider_test.go)
docs/resources/*.md
```

### ⚠️ **DO NOT DELETE (Manual)**
These files are manually created and maintained:

```
internal/provider/provider.go
internal/provider/client.go
internal/provider/utils.go
internal/provider/resource_*.go (non-generated resources)
internal/provider/data_source_*.go
internal/provider/tests/provider_test.go
docs/index.md
```

## Makefile Targets for Generated Files

### Clean Generated Files
```bash
# Clean only generated files from API specs
make clean-generated
```

This target removes:
- All `internal/provider/generated_*.go` files
- All test files except `provider_test.go`
- All documentation files except `docs/index.md`

### Clean Everything
```bash
# Clean build artifacts + generated files
make clean-all
```

### Regenerate After Cleaning
```bash
# Clean and regenerate everything
make clean-generated
make generate-full
```

## Generation Process

### Input Sources
The generator processes these OpenAPI specification files from `api-specs/`:

- `admin_users_and_roles_api.yaml` - User management API
- `auth_api.yaml` - Authentication API
- `deployments_internal_domains_api.yaml` - Internal domains API
- `deployments_internal_networks_api.yaml` - Internal networks API
- `deployments_networks_api.yaml` - Networks API
- `deployments_sites_api.yaml` - Sites API
- `investigate_api.yaml` - Investigation API
- `policies_destination_lists_api.yaml` - Destination lists API
- `reports_reporting_api.yaml` - Reporting API

### Generation Configuration
The generation process is controlled by:
- `tools/generator/config/generation.yaml` - Main configuration
- `tools/generator/config/advanced_config.yaml` - Advanced settings

### File Naming Patterns
Based on the configuration in `generation.yaml`:

- **Resources**: `generated_{resource_name}_resource.go`
- **Client Methods**: `generated_{resource_name}_methods.go`
- **Tests**: `{resource_name}_test.go`
- **Documentation**: `{resource_name}.md`

## Best Practices

### 1. Always Clean Before Regenerating
```bash
make clean-generated && make generate-full
```

### 2. Don't Edit Generated Files
Generated files contain a header comment indicating they are auto-generated. Never edit these files manually as changes will be lost.

### 3. Use Version Control
Generated files are tracked in Git to ensure consistency across environments. After regeneration, commit the changes:

```bash
make clean-generated
make generate-full
git add .
git commit -m "Regenerate provider code from updated API specs"
```

### 4. Validate After Generation
```bash
make generate-full
make test
make build
```

## Troubleshooting

### Generated Files Not Updating
1. Clean generated files: `make clean-generated`
2. Regenerate: `make generate-full`
3. Check for errors in the generation process

### Build Errors After Generation
1. Ensure all API specs are valid OpenAPI 3.0+ format
2. Check the generation configuration in `tools/generator/config/`
3. Run `go fmt ./internal/provider/generated_*.go`

### Missing Generated Files
1. Verify API spec files exist in `api-specs/`
2. Check that specs contain valid resource definitions
3. Review generation logs for parsing errors

## Development Workflow

### Adding New API Specifications
1. Add new `.yaml` file to `api-specs/`
2. Run `make clean-generated && make generate-full`
3. Test the generated resources
4. Commit the changes

### Modifying Existing API Specifications
1. Update the `.yaml` file in `api-specs/`
2. Run `make clean-generated && make generate-full`
3. Review the generated changes
4. Update any manual code that depends on the changes
5. Test and commit

### Custom Resource Development
For resources that require custom logic beyond what can be generated:
1. Create manual resource files (e.g., `resource_custom.go`)
2. Add them to the provider registration in `provider.go`
3. Do NOT use the `generated_` prefix to avoid conflicts

This documentation ensures clear understanding of the generated file system and proper maintenance procedures.