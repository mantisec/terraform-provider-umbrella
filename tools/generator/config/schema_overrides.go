package config

import (
	"fmt"
	"os"
	"regexp"

	"github.com/mantisec/terraform-provider-umbrella/tools/generator/parser"
	"gopkg.in/yaml.v3"
)

// AdvancedConfig represents the advanced configuration system
type AdvancedConfig struct {
	SchemaOverrides  SchemaOverrides        `yaml:"schema_overrides"`
	Validation       ValidationConfig       `yaml:"validation"`
	PlanModifiers    PlanModifiersConfig    `yaml:"plan_modifiers"`
	APIConfigs       APIConfigs             `yaml:"api_configs"`
	CodeGeneration   CodeGenerationConfig   `yaml:"code_generation"`
	QualityAssurance QualityAssuranceConfig `yaml:"quality_assurance"`
	Performance      PerformanceConfig      `yaml:"performance"`
	Compatibility    CompatibilityConfig    `yaml:"compatibility"`
}

// SchemaOverrides contains schema transformation rules
type SchemaOverrides struct {
	Global    GlobalOverrides             `yaml:"global"`
	Resources map[string]ResourceOverride `yaml:"resources"`
}

// GlobalOverrides contains global schema transformations
type GlobalOverrides struct {
	TypeMappings  map[string]string `yaml:"type_mappings"`
	FieldMappings map[string]string `yaml:"field_mappings"`
	ForceRequired []string          `yaml:"force_required"`
	ForceOptional []string          `yaml:"force_optional"`
}

// ResourceOverride contains resource-specific overrides
type ResourceOverride struct {
	SchemaTransforms []SchemaTransform `yaml:"schema_transforms"`
	ValidationRules  []ValidationRule  `yaml:"validation_rules"`
}

// SchemaTransform represents a single schema transformation
type SchemaTransform struct {
	Field       string `yaml:"field"`
	Type        string `yaml:"type"`
	ElementType string `yaml:"element_type,omitempty"`
	Validation  string `yaml:"validation,omitempty"`
}

// ValidationRule represents a custom validation rule
type ValidationRule struct {
	Rule    string `yaml:"rule"`
	Message string `yaml:"message"`
}

// ValidationConfig contains validation settings
type ValidationConfig struct {
	EnabledValidators []string                   `yaml:"enabled_validators"`
	CustomValidators  map[string]CustomValidator `yaml:"custom_validators"`
	SecurityRules     []SecurityRule             `yaml:"security_rules"`
}

// CustomValidator represents a custom validation function
type CustomValidator struct {
	Pattern string `yaml:"pattern"`
	Message string `yaml:"message"`
}

// SecurityRule represents a security validation rule
type SecurityRule struct {
	Name     string `yaml:"name"`
	Pattern  string `yaml:"pattern"`
	Message  string `yaml:"message"`
	Severity string `yaml:"severity"`
}

// PlanModifiersConfig contains plan modifier settings
type PlanModifiersConfig struct {
	Defaults  map[string][]string            `yaml:"defaults"`
	Resources map[string]map[string][]string `yaml:"resources"`
}

// APIConfigs contains API-specific configurations
type APIConfigs map[string]APIConfig

// APIConfig represents configuration for a specific API
type APIConfig struct {
	BaseURL        string                    `yaml:"base_url"`
	Version        string                    `yaml:"version"`
	RateLimiting   RateLimitingConfig        `yaml:"rate_limiting"`
	Authentication AuthenticationConfig      `yaml:"authentication"`
	Endpoints      map[string]EndpointConfig `yaml:"endpoints"`
}

// RateLimitingConfig contains rate limiting settings
type RateLimitingConfig struct {
	RequestsPerSecond int    `yaml:"requests_per_second"`
	BurstSize         int    `yaml:"burst_size"`
	RetryAttempts     int    `yaml:"retry_attempts"`
	RetryDelay        string `yaml:"retry_delay"`
}

// AuthenticationConfig contains authentication settings
type AuthenticationConfig struct {
	Type   string   `yaml:"type"`
	Scopes []string `yaml:"scopes"`
}

// EndpointConfig contains endpoint-specific settings
type EndpointConfig struct {
	Timeout       string `yaml:"timeout"`
	RetryAttempts int    `yaml:"retry_attempts"`
	CacheTTL      string `yaml:"cache_ttl"`
}

// CodeGenerationConfig contains code generation customization
type CodeGenerationConfig struct {
	Imports       ImportsConfig       `yaml:"imports"`
	Naming        NamingConfig        `yaml:"naming"`
	ErrorHandling ErrorHandlingConfig `yaml:"error_handling"`
}

// ImportsConfig contains import customization
type ImportsConfig struct {
	Additional []string          `yaml:"additional"`
	Aliases    map[string]string `yaml:"aliases"`
}

// ErrorHandlingConfig contains error handling customization
type ErrorHandlingConfig struct {
	CustomErrors []string `yaml:"custom_errors"`
	WrapPatterns []string `yaml:"wrap_patterns"`
}

// QualityAssuranceConfig contains quality assurance settings
type QualityAssuranceConfig struct {
	CodeQuality CodeQualityConfig `yaml:"code_quality"`
	Testing     TestingConfig     `yaml:"testing"`
}

// CodeQualityConfig contains code quality settings
type CodeQualityConfig struct {
	EnabledChecks []string      `yaml:"enabled_checks"`
	CustomRules   []QualityRule `yaml:"custom_rules"`
}

// QualityRule represents a custom quality rule
type QualityRule struct {
	Name          string `yaml:"name"`
	MaxLines      int    `yaml:"max_lines,omitempty"`
	MaxComplexity int    `yaml:"max_complexity,omitempty"`
	Message       string `yaml:"message"`
}

// TestingConfig contains testing settings
type TestingConfig struct {
	TestTypes []string       `yaml:"test_types"`
	Coverage  CoverageConfig `yaml:"coverage"`
	Mocking   MockingConfig  `yaml:"mocking"`
}

// CoverageConfig contains test coverage settings
type CoverageConfig struct {
	Minimum int `yaml:"minimum"`
	Target  int `yaml:"target"`
}

// MockingConfig contains mocking settings
type MockingConfig struct {
	Enabled          bool   `yaml:"enabled"`
	MockExternalAPIs bool   `yaml:"mock_external_apis"`
	MockFilePattern  string `yaml:"mock_file_pattern"`
}

// PerformanceConfig contains performance optimization settings
type PerformanceConfig struct {
	ParallelProcessing ParallelProcessingConfig `yaml:"parallel_processing"`
	Caching            CachingConfig            `yaml:"caching"`
	Incremental        IncrementalConfig        `yaml:"incremental"`
}

// ParallelProcessingConfig contains parallel processing settings
type ParallelProcessingConfig struct {
	Enabled    bool `yaml:"enabled"`
	MaxWorkers int  `yaml:"max_workers"`
	BatchSize  int  `yaml:"batch_size"`
}

// CachingConfig contains caching settings
type CachingConfig struct {
	Enabled   bool              `yaml:"enabled"`
	CacheDir  string            `yaml:"cache_dir"`
	TTL       string            `yaml:"ttl"`
	CacheKeys map[string]string `yaml:"cache_keys"`
}

// IncrementalConfig contains incremental generation settings
type IncrementalConfig struct {
	Enabled            bool   `yaml:"enabled"`
	ChangeDetection    string `yaml:"change_detection"`
	DependencyTracking bool   `yaml:"dependency_tracking"`
}

// CompatibilityConfig contains backward compatibility settings
type CompatibilityConfig struct {
	VersionDetection VersionDetectionConfig `yaml:"version_detection"`
	Migration        MigrationConfig        `yaml:"migration"`
	Deprecation      DeprecationConfig      `yaml:"deprecation"`
}

// VersionDetectionConfig contains version detection settings
type VersionDetectionConfig struct {
	Enabled     bool   `yaml:"enabled"`
	VersionFile string `yaml:"version_file"`
}

// MigrationConfig contains migration settings
type MigrationConfig struct {
	Enabled        bool   `yaml:"enabled"`
	MigrationDir   string `yaml:"migration_dir"`
	BackupExisting bool   `yaml:"backup_existing"`
}

// DeprecationConfig contains deprecation settings
type DeprecationConfig struct {
	Enabled            bool                `yaml:"enabled"`
	WarningFormat      string              `yaml:"warning_format"`
	DeprecatedPatterns []DeprecatedPattern `yaml:"deprecated_patterns"`
}

// DeprecatedPattern represents a deprecated pattern
type DeprecatedPattern struct {
	Pattern     string `yaml:"pattern"`
	Replacement string `yaml:"replacement"`
	Version     string `yaml:"version"`
}

// SchemaOverrideEngine handles schema transformations
type SchemaOverrideEngine struct {
	config *AdvancedConfig
}

// NewSchemaOverrideEngine creates a new schema override engine
func NewSchemaOverrideEngine(config *AdvancedConfig) *SchemaOverrideEngine {
	return &SchemaOverrideEngine{
		config: config,
	}
}

// ApplyOverrides applies schema overrides to a parsed schema
func (e *SchemaOverrideEngine) ApplyOverrides(schema *parser.Schema, resourceName string) (*parser.Schema, error) {
	if schema == nil {
		return nil, fmt.Errorf("schema cannot be nil")
	}

	// Create a copy of the schema to avoid modifying the original
	overriddenSchema := e.copySchema(schema)

	// Apply global overrides
	if err := e.applyGlobalOverrides(overriddenSchema); err != nil {
		return nil, fmt.Errorf("failed to apply global overrides: %w", err)
	}

	// Apply resource-specific overrides
	if resourceOverride, exists := e.config.SchemaOverrides.Resources[resourceName]; exists {
		if err := e.applyResourceOverrides(overriddenSchema, resourceOverride); err != nil {
			return nil, fmt.Errorf("failed to apply resource overrides for %s: %w", resourceName, err)
		}
	}

	return overriddenSchema, nil
}

// applyGlobalOverrides applies global schema transformations
func (e *SchemaOverrideEngine) applyGlobalOverrides(schema *parser.Schema) error {
	global := e.config.SchemaOverrides.Global

	// Apply type mappings
	if mappedType, exists := global.TypeMappings[e.getTypeKey(schema)]; exists {
		schema.Type = mappedType
	}

	// Apply field mappings to properties
	if schema.Properties != nil {
		newProperties := make(map[string]*parser.Schema)
		for fieldName, fieldSchema := range schema.Properties {
			newFieldName := fieldName
			if mappedName, exists := global.FieldMappings[fieldName]; exists {
				newFieldName = mappedName
			}
			newProperties[newFieldName] = fieldSchema
		}
		schema.Properties = newProperties
	}

	// Apply force required/optional
	e.applyRequiredOverrides(schema, global.ForceRequired, global.ForceOptional)

	return nil
}

// applyResourceOverrides applies resource-specific schema transformations
func (e *SchemaOverrideEngine) applyResourceOverrides(schema *parser.Schema, override ResourceOverride) error {
	// Apply schema transforms
	for _, transform := range override.SchemaTransforms {
		if err := e.applySchemaTransform(schema, transform); err != nil {
			return fmt.Errorf("failed to apply schema transform for field %s: %w", transform.Field, err)
		}
	}

	return nil
}

// applySchemaTransform applies a single schema transformation
func (e *SchemaOverrideEngine) applySchemaTransform(schema *parser.Schema, transform SchemaTransform) error {
	if schema.Properties == nil {
		return nil
	}

	fieldSchema, exists := schema.Properties[transform.Field]
	if !exists {
		return nil // Field doesn't exist, skip transformation
	}

	// Apply type transformation
	if transform.Type != "" {
		fieldSchema.Type = transform.Type
	}

	// Apply element type for arrays
	if transform.ElementType != "" && fieldSchema.Items != nil {
		fieldSchema.Items.Type = transform.ElementType
	}

	// Apply validation (this would be used during code generation)
	if transform.Validation != "" {
		// Store validation rules in description or custom field
		if fieldSchema.Description != "" {
			fieldSchema.Description += fmt.Sprintf(" [Validation: %s]", transform.Validation)
		} else {
			fieldSchema.Description = fmt.Sprintf("[Validation: %s]", transform.Validation)
		}
	}

	return nil
}

// applyRequiredOverrides applies required/optional field overrides
func (e *SchemaOverrideEngine) applyRequiredOverrides(schema *parser.Schema, forceRequired, forceOptional []string) {
	if schema.Properties == nil {
		return
	}

	// Convert required slice to map for easier manipulation
	requiredMap := make(map[string]bool)
	for _, field := range schema.Required {
		requiredMap[field] = true
	}

	// Apply force required
	for _, field := range forceRequired {
		if _, exists := schema.Properties[field]; exists {
			requiredMap[field] = true
		}
	}

	// Apply force optional
	for _, field := range forceOptional {
		delete(requiredMap, field)
	}

	// Convert back to slice
	schema.Required = make([]string, 0, len(requiredMap))
	for field := range requiredMap {
		schema.Required = append(schema.Required, field)
	}
}

// getTypeKey generates a key for type mapping lookup
func (e *SchemaOverrideEngine) getTypeKey(schema *parser.Schema) string {
	if schema.Format != "" {
		return fmt.Sprintf("%s_with_format_%s", schema.Type, schema.Format)
	}
	return schema.Type
}

// copySchema creates a deep copy of a schema
func (e *SchemaOverrideEngine) copySchema(schema *parser.Schema) *parser.Schema {
	if schema == nil {
		return nil
	}

	schemaCopy := &parser.Schema{
		Type:        schema.Type,
		Format:      schema.Format,
		Description: schema.Description,
		Required:    make([]string, len(schema.Required)),
		Enum:        make([]interface{}, len(schema.Enum)),
		Default:     schema.Default,
		Example:     schema.Example,
		Ref:         schema.Ref,
	}

	// Copy required fields
	copy(schemaCopy.Required, schema.Required)

	// Copy enum values
	copy(schemaCopy.Enum, schema.Enum)

	// Copy properties
	if schema.Properties != nil {
		schemaCopy.Properties = make(map[string]*parser.Schema)
		for key, value := range schema.Properties {
			schemaCopy.Properties[key] = e.copySchema(value)
		}
	}

	// Copy items
	if schema.Items != nil {
		schemaCopy.Items = e.copySchema(schema.Items)
	}

	// Copy AllOf, OneOf, AnyOf
	if schema.AllOf != nil {
		schemaCopy.AllOf = make([]*parser.Schema, len(schema.AllOf))
		for i, s := range schema.AllOf {
			schemaCopy.AllOf[i] = e.copySchema(s)
		}
	}

	if schema.OneOf != nil {
		schemaCopy.OneOf = make([]*parser.Schema, len(schema.OneOf))
		for i, s := range schema.OneOf {
			schemaCopy.OneOf[i] = e.copySchema(s)
		}
	}

	if schema.AnyOf != nil {
		schemaCopy.AnyOf = make([]*parser.Schema, len(schema.AnyOf))
		for i, s := range schema.AnyOf {
			schemaCopy.AnyOf[i] = e.copySchema(s)
		}
	}

	return schemaCopy
}

// ValidateSchema validates a schema against custom validation rules
func (e *SchemaOverrideEngine) ValidateSchema(schema *parser.Schema, resourceName string) []ValidationError {
	var errors []ValidationError

	// Apply custom validators
	for name, validator := range e.config.Validation.CustomValidators {
		if err := e.validateWithCustomValidator(schema, name, validator); err != nil {
			errors = append(errors, *err)
		}
	}

	// Apply resource-specific validation rules
	if resourceOverride, exists := e.config.SchemaOverrides.Resources[resourceName]; exists {
		for _, rule := range resourceOverride.ValidationRules {
			if err := e.validateWithRule(schema, rule); err != nil {
				errors = append(errors, *err)
			}
		}
	}

	return errors
}

// validateWithCustomValidator validates schema with a custom validator
func (e *SchemaOverrideEngine) validateWithCustomValidator(schema *parser.Schema, name string, validator CustomValidator) *ValidationError {
	pattern, err := regexp.Compile(validator.Pattern)
	if err != nil {
		return &ValidationError{
			Type:    "validator_error",
			Message: fmt.Sprintf("Invalid regex pattern in validator %s: %v", name, err),
		}
	}

	// Check if schema description or example matches the pattern
	if schema.Description != "" && pattern.MatchString(schema.Description) {
		return nil // Valid
	}

	if schema.Example != nil {
		if exampleStr, ok := schema.Example.(string); ok && pattern.MatchString(exampleStr) {
			return nil // Valid
		}
	}

	return &ValidationError{
		Type:    "validation_error",
		Message: validator.Message,
		Field:   name,
	}
}

// validateWithRule validates schema with a validation rule
func (e *SchemaOverrideEngine) validateWithRule(schema *parser.Schema, rule ValidationRule) *ValidationError {
	// This is a simplified implementation
	// In a real implementation, you would have specific validation logic for each rule
	switch rule.Rule {
	case "destinations_not_empty":
		if schema.Properties != nil {
			if destProp, exists := schema.Properties["destinations"]; exists {
				if destProp.Type == "array" && len(destProp.Required) == 0 {
					return &ValidationError{
						Type:    "validation_error",
						Message: rule.Message,
						Field:   "destinations",
					}
				}
			}
		}
	case "valid_access_type":
		if schema.Properties != nil {
			if accessProp, exists := schema.Properties["access"]; exists {
				if accessProp.Type == "string" && len(accessProp.Enum) == 0 {
					return &ValidationError{
						Type:    "validation_error",
						Message: rule.Message,
						Field:   "access",
					}
				}
			}
		}
	}

	return nil
}

// ValidationError represents a validation error
type ValidationError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

// Error implements the error interface
func (e ValidationError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("%s in field %s: %s", e.Type, e.Field, e.Message)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

// GetPlanModifiers returns plan modifiers for a specific resource and field
func (e *SchemaOverrideEngine) GetPlanModifiers(resourceName, fieldName, fieldType string) []string {
	// Check resource-specific modifiers first
	if resourceModifiers, exists := e.config.PlanModifiers.Resources[resourceName]; exists {
		if fieldModifiers, exists := resourceModifiers[fieldName]; exists {
			return fieldModifiers
		}
	}

	// Fall back to default modifiers for the field type
	if defaultModifiers, exists := e.config.PlanModifiers.Defaults[fieldType]; exists {
		return defaultModifiers
	}

	return nil
}

// LoadAdvancedConfig loads advanced configuration from a YAML file
func LoadAdvancedConfig(path string) (*AdvancedConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read advanced config file %s: %w", path, err)
	}

	var config AdvancedConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse advanced config file %s: %w", path, err)
	}

	return &config, nil
}
