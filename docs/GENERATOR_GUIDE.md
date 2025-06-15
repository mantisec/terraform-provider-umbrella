# Terraform Provider Code Generator - Complete Guide

This guide provides comprehensive documentation for the advanced code generation system for the terraform-provider-umbrella project.

## Table of Contents

1. [Overview](#overview)
2. [Installation & Setup](#installation--setup)
3. [Basic Usage](#basic-usage)
4. [Advanced Configuration](#advanced-configuration)
5. [Quality Assurance](#quality-assurance)
6. [CI/CD Integration](#cicd-integration)
7. [Performance Optimization](#performance-optimization)
8. [Error Handling & Debugging](#error-handling--debugging)
9. [Customization](#customization)
10. [Best Practices](#best-practices)
11. [Troubleshooting](#troubleshooting)

## Overview

The terraform-provider-umbrella code generator is a sophisticated system that automatically generates Terraform provider code from OpenAPI specifications. It includes advanced features for quality assurance, performance optimization, and CI/CD integration.

### Key Features

- **Automated Code Generation**: Generate resources, data sources, and client methods from OpenAPI specs
- **Advanced Configuration**: Schema overrides, validation rules, and custom mappings
- **Quality Assurance**: Comprehensive validation and testing framework
- **CI/CD Integration**: Automated workflows for continuous integration
- **Performance Optimization**: Parallel processing and caching for large-scale generation
- **Error Handling**: Advanced error recovery and structured logging
- **Backward Compatibility**: Migration tools and version detection

### Architecture

```
tools/generator/
├── cmd/generate/           # Command-line interface
├── config/                 # Configuration management
│   ├── config.go          # Basic configuration
│   ├── schema_overrides.go # Advanced schema transformations
│   ├── generation.yaml    # Basic configuration file
│   └── advanced_config.yaml # Advanced configuration file
├── parser/                 # OpenAPI specification parsing
├── generator/              # Code generation engines
├── templates/              # Code templates
├── validator/              # Quality assurance validators
├── performance/            # Performance optimization
└── logger/                 # Structured logging
```

## Installation & Setup

### Prerequisites

- Go 1.21 or later
- Git
- Access to OpenAPI specification files

### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/mantisec/terraform-provider-umbrella.git
   cd terraform-provider-umbrella
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Build the generator**:
   ```bash
   go build -o bin/generator ./tools/generator/cmd/generate
   ```

4. **Verify installation**:
   ```bash
   ./bin/generator --help
   ```

### Initial Configuration

1. **Copy configuration templates**:
   ```bash
   cp tools/generator/config/generation.yaml.example tools/generator/config/generation.yaml
   cp tools/generator/config/advanced_config.yaml.example tools/generator/config/advanced_config.yaml
   ```

2. **Edit configuration files** to match your environment and requirements.

## Basic Usage

### Command Line Interface

The generator provides a command-line interface with the following options:

```bash
./bin/generator [options]

Options:
  -config string
        Path to generation configuration file (default "tools/generator/config/generation.yaml")
  -specs string
        Directory containing OpenAPI specification files (default ".")
  -output string
        Output directory for generated files (default "internal/provider")
  -verbose
        Enable verbose logging
  -debug
        Enable debug mode with detailed logging
  -parallel
        Enable parallel processing (default: true)
  -cache
        Enable caching (default: true)
  -validate
        Run validation after generation (default: true)
```

### Basic Generation

1. **Generate from all specs**:
   ```bash
   ./bin/generator -config tools/generator/config/generation.yaml -specs . -output internal/provider
   ```

2. **Generate with verbose output**:
   ```bash
   ./bin/generator -verbose -debug
   ```

3. **Generate specific specs**:
   ```bash
   ./bin/generator -specs policies_destination_lists_api.yaml
   ```

### Generated Files

The generator creates the following types of files:

- **Resources**: `resource_*.go` - Terraform resource implementations
- **Data Sources**: `data_source_*.go` - Terraform data source implementations
- **Client Methods**: `generated_*_methods.go` - API client methods
- **Registry**: `generated_registry.go` - Resource and data source registry
- **Tests**: `*_test.go` - Unit and acceptance tests
- **Documentation**: `docs/resources/*.md` - Resource documentation

## Advanced Configuration

### Schema Overrides

The advanced configuration system allows you to customize schema generation:

```yaml
# tools/generator/config/advanced_config.yaml
schema_overrides:
  global:
    type_mappings:
      "string_with_format_date": "string"
      "integer_with_format_int64": "int64"
    
    field_mappings:
      "id": "ID"
      "url": "URL"
    
    force_required:
      - "name"
      - "id"
  
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

### Validation Configuration

Configure custom validation rules:

```yaml
validation:
  enabled_validators:
    - "schema_consistency"
    - "api_compatibility"
    - "terraform_compliance"
    - "security_checks"
  
  custom_validators:
    ip_address:
      pattern: "^(?:[0-9]{1,3}\\.){3}[0-9]{1,3}$"
      message: "Must be a valid IPv4 address"
  
  security_rules:
    - name: "no_hardcoded_secrets"
      pattern: "(password|secret|key|token)\\s*=\\s*[\"'][^\"']+[\"']"
      message: "Hardcoded secrets detected"
      severity: "error"
```

### Plan Modifiers

Configure Terraform plan modifiers:

```yaml
plan_modifiers:
  defaults:
    string:
      - "stringplanmodifier.RequiresReplace()"
      - "stringplanmodifier.UseStateForUnknown()"
  
  resources:
    destination_list:
      name:
        - "stringplanmodifier.RequiresReplace()"
      destinations:
        - "listplanmodifier.UseStateForUnknown()"
```

## Quality Assurance

### Code Validation

The generator includes comprehensive code validation:

1. **Syntax Validation**: Ensures generated code compiles
2. **Quality Checks**: Code formatting, naming conventions, complexity
3. **Security Scanning**: Detects security vulnerabilities
4. **Terraform Compliance**: Validates Terraform provider patterns

### Running Validation

```bash
# Run all validations
./bin/generator -validate

# Run specific validation types
go run ./tools/generator/validator/code_validator.go -dir internal/provider
go run ./tools/generator/validator/schema_validator.go -spec api.yaml
go run ./tools/generator/validator/api_validator.go -spec api.yaml
```

### Validation Reports

Validation generates detailed reports:

- **Code Quality Report**: Formatting, linting, complexity issues
- **Schema Validation Report**: Schema consistency and compliance
- **API Compatibility Report**: REST compliance and Terraform compatibility
- **Security Report**: Security vulnerabilities and recommendations

## CI/CD Integration

### GitHub Actions Workflows

The generator includes pre-configured GitHub Actions workflows:

1. **Code Generation Workflow** (`.github/workflows/code-generation.yml`):
   - Triggered on OpenAPI spec changes
   - Validates specifications
   - Generates code in parallel
   - Creates pull requests with generated code

2. **Validation Workflow** (`.github/workflows/validate-generated.yml`):
   - Validates generated code quality
   - Runs unit and integration tests
   - Performs security scanning
   - Generates comprehensive reports

### Workflow Configuration

Configure workflows by setting environment variables:

```yaml
env:
  GO_VERSION: '1.21'
  GENERATOR_CONFIG: 'tools/generator/config/generation.yaml'
  ADVANCED_CONFIG: 'tools/generator/config/advanced_config.yaml'
```

### Manual Triggers

Workflows can be triggered manually with custom parameters:

- **Force Regeneration**: Regenerate all code regardless of changes
- **Specific Files**: Process only specified OpenAPI files
- **Validation Level**: Choose validation depth (basic, standard, comprehensive)

## Performance Optimization

### Parallel Processing

Enable parallel processing for large-scale generation:

```yaml
performance:
  parallel_processing:
    enabled: true
    max_workers: 4
    batch_size: 10
```

### Caching

Configure caching to improve performance:

```yaml
performance:
  caching:
    enabled: true
    cache_dir: ".generator_cache"
    ttl: "1h"
    
    cache_keys:
      parsed_specs: "spec_%s_%s"
      generated_schemas: "schema_%s_%s"
      template_renders: "template_%s_%s"
```

### Incremental Generation

Enable incremental generation to process only changed files:

```yaml
performance:
  incremental:
    enabled: true
    change_detection: "file_hash"
    dependency_tracking: true
```

### Performance Monitoring

Monitor generation performance:

```bash
# Enable performance logging
./bin/generator -verbose -debug

# View cache statistics
./bin/generator -cache-stats

# Monitor parallel processing
./bin/generator -parallel -verbose
```

## Error Handling & Debugging

### Structured Logging

The generator uses structured logging for better debugging:

```go
// Enable debug logging
logger := logger.NewGenerationLogger(config, "generator")
logger.SetDebugMode(true)
logger.SetLevel(logger.DEBUG)

// Log with context
logger.WithOperation("generate_resource").
       WithFile("api.yaml").
       Info("Starting resource generation")
```

### Error Recovery

The system includes automatic error recovery:

- **Panic Recovery**: Catches and logs panics during generation
- **Retry Logic**: Retries failed operations with exponential backoff
- **Partial Failure Handling**: Continues processing other files when one fails

### Debug Mode

Enable debug mode for detailed troubleshooting:

```bash
./bin/generator -debug -verbose
```

Debug mode provides:
- Detailed stack traces
- Function call information
- File and line numbers
- Performance timing
- Cache hit/miss statistics

## Customization

### Custom Templates

Create custom templates for specific needs:

1. **Create template file**:
   ```go
   // tools/generator/templates/custom_resource.go.tmpl
   package provider
   
   import (
       "context"
       "github.com/hashicorp/terraform-plugin-framework/resource"
   )
   
   // {{.ResourceName}}Resource implements the resource interface
   type {{.ResourceName}}Resource struct {
       // Custom implementation
   }
   ```

2. **Register template**:
   ```yaml
   templates:
     custom_resource_template: "templates/custom_resource.go.tmpl"
   ```

### Custom Validators

Implement custom validation logic:

```go
// tools/generator/validator/custom_validator.go
package validator

type CustomValidator struct {
    config *config.AdvancedConfig
}

func (v *CustomValidator) ValidateCustomRule(schema *parser.Schema) error {
    // Custom validation logic
    return nil
}
```

### Custom Generators

Create specialized generators:

```go
// tools/generator/generator/custom_generator.go
package generator

type CustomGenerator struct {
    config *config.Config
    engine *TemplateEngine
}

func (g *CustomGenerator) GenerateCustomCode(spec *parser.APISpec) error {
    // Custom generation logic
    return nil
}
```

## Best Practices

### OpenAPI Specification Guidelines

1. **Use Descriptive Names**: Clear, consistent naming for operations and schemas
2. **Include Examples**: Provide examples for better code generation
3. **Define Proper Types**: Use appropriate data types and formats
4. **Add Descriptions**: Document all fields and operations
5. **Use Tags**: Organize operations with meaningful tags

### Configuration Management

1. **Version Control**: Keep configuration files in version control
2. **Environment-Specific**: Use different configs for different environments
3. **Validation**: Validate configuration files before use
4. **Documentation**: Document custom configuration options

### Code Generation Workflow

1. **Incremental Updates**: Use incremental generation for efficiency
2. **Validation First**: Always validate before generating
3. **Review Generated Code**: Review generated code before merging
4. **Test Thoroughly**: Run comprehensive tests on generated code
5. **Monitor Performance**: Track generation performance and optimize

### Quality Assurance

1. **Automated Validation**: Use CI/CD for automatic validation
2. **Code Reviews**: Review generated code changes
3. **Security Scanning**: Regular security scans of generated code
4. **Performance Testing**: Test performance of generated provider
5. **Documentation**: Keep documentation up to date

## Troubleshooting

### Common Issues

#### Generation Fails

**Problem**: Code generation fails with errors

**Solutions**:
1. Check OpenAPI specification validity
2. Verify configuration file syntax
3. Ensure all dependencies are installed
4. Check file permissions
5. Enable debug mode for detailed logs

```bash
# Debug generation issues
./bin/generator -debug -verbose
```

#### Invalid Generated Code

**Problem**: Generated code doesn't compile

**Solutions**:
1. Run code validation
2. Check template syntax
3. Verify schema overrides
4. Review custom mappings
5. Check for naming conflicts

```bash
# Validate generated code
go build ./internal/provider/...
go vet ./internal/provider/...
```

#### Performance Issues

**Problem**: Generation is slow

**Solutions**:
1. Enable parallel processing
2. Configure caching
3. Use incremental generation
4. Optimize OpenAPI specifications
5. Monitor resource usage

```bash
# Enable performance optimizations
./bin/generator -parallel -cache -incremental
```

#### Cache Issues

**Problem**: Stale or corrupted cache

**Solutions**:
1. Clear cache directory
2. Disable caching temporarily
3. Check cache configuration
4. Verify file permissions

```bash
# Clear cache
rm -rf .generator_cache

# Disable cache
./bin/generator -cache=false
```

### Debug Commands

```bash
# Check configuration
./bin/generator -config-check

# Validate OpenAPI specs
./bin/generator -validate-specs

# Test templates
./bin/generator -test-templates

# Cache statistics
./bin/generator -cache-stats

# Performance profiling
./bin/generator -profile
```

### Log Analysis

Analyze logs for troubleshooting:

```bash
# Filter error logs
grep "ERROR" generator.log

# Check performance metrics
grep "duration" generator.log

# View cache statistics
grep "cache" generator.log

# Monitor parallel processing
grep "worker" generator.log
```

### Getting Help

1. **Documentation**: Check this guide and inline documentation
2. **Issues**: Create GitHub issues for bugs and feature requests
3. **Discussions**: Use GitHub discussions for questions
4. **Logs**: Include debug logs when reporting issues
5. **Configuration**: Share configuration files (sanitized) when needed

## Conclusion

The terraform-provider-umbrella code generator provides a comprehensive solution for automated Terraform provider development. With its advanced features for quality assurance, performance optimization, and CI/CD integration, it enables efficient and reliable provider development at scale.

For additional help and support, please refer to the project's GitHub repository and documentation.