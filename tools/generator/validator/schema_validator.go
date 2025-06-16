package validator

import (
	"fmt"
	"strings"

	"github.com/mantisec/terraform-provider-umbrella/tools/generator/config"
	"github.com/mantisec/terraform-provider-umbrella/tools/generator/parser"
)

// SchemaValidator validates schema consistency and compliance
type SchemaValidator struct {
	config *config.AdvancedConfig
}

// SchemaValidationResult represents the result of schema validation
type SchemaValidationResult struct {
	Valid    bool              `json:"valid"`
	Errors   []SchemaError     `json:"errors"`
	Warnings []SchemaWarning   `json:"warnings"`
	Summary  ValidationSummary `json:"summary"`
}

// SchemaError represents a schema validation error
type SchemaError struct {
	Type       string `json:"type"`
	Message    string `json:"message"`
	Path       string `json:"path"`
	Field      string `json:"field,omitempty"`
	Suggestion string `json:"suggestion,omitempty"`
}

// SchemaWarning represents a schema validation warning
type SchemaWarning struct {
	Type       string `json:"type"`
	Message    string `json:"message"`
	Path       string `json:"path"`
	Field      string `json:"field,omitempty"`
	Suggestion string `json:"suggestion,omitempty"`
}

// ValidationSummary provides a summary of validation results
type ValidationSummary struct {
	TotalSchemas     int     `json:"total_schemas"`
	ValidSchemas     int     `json:"valid_schemas"`
	ErrorCount       int     `json:"error_count"`
	WarningCount     int     `json:"warning_count"`
	ConsistencyScore float64 `json:"consistency_score"`
}

// NewSchemaValidator creates a new schema validator
func NewSchemaValidator(config *config.AdvancedConfig) *SchemaValidator {
	return &SchemaValidator{
		config: config,
	}
}

// ValidateAPISpec validates an entire API specification for schema consistency
func (v *SchemaValidator) ValidateAPISpec(spec *parser.APISpec) (*SchemaValidationResult, error) {
	result := &SchemaValidationResult{
		Valid:    true,
		Errors:   make([]SchemaError, 0),
		Warnings: make([]SchemaWarning, 0),
		Summary: ValidationSummary{
			TotalSchemas: 0,
			ValidSchemas: 0,
		},
	}

	// Validate components schemas
	if spec.Components.Schemas != nil {
		for schemaName, schema := range spec.Components.Schemas {
			v.validateSchema(schema, fmt.Sprintf("components.schemas.%s", schemaName), result)
			result.Summary.TotalSchemas++
		}
	}

	// Validate operation schemas
	for path, pathItem := range spec.Paths {
		v.validatePathItem(&pathItem, path, result)
	}

	// Calculate summary
	result.Summary.ErrorCount = len(result.Errors)
	result.Summary.WarningCount = len(result.Warnings)
	result.Summary.ValidSchemas = result.Summary.TotalSchemas - result.Summary.ErrorCount

	if result.Summary.TotalSchemas > 0 {
		result.Summary.ConsistencyScore = float64(result.Summary.ValidSchemas) / float64(result.Summary.TotalSchemas) * 100
	}

	result.Valid = result.Summary.ErrorCount == 0

	return result, nil
}

// ValidateSchema validates a single schema
func (v *SchemaValidator) ValidateSchema(schema *parser.Schema, schemaPath string) (*SchemaValidationResult, error) {
	result := &SchemaValidationResult{
		Valid:    true,
		Errors:   make([]SchemaError, 0),
		Warnings: make([]SchemaWarning, 0),
		Summary: ValidationSummary{
			TotalSchemas: 1,
		},
	}

	v.validateSchema(schema, schemaPath, result)

	// Calculate summary
	result.Summary.ErrorCount = len(result.Errors)
	result.Summary.WarningCount = len(result.Warnings)
	result.Summary.ValidSchemas = 1 - result.Summary.ErrorCount
	result.Summary.ConsistencyScore = float64(result.Summary.ValidSchemas) * 100

	result.Valid = result.Summary.ErrorCount == 0

	return result, nil
}

// validateSchema performs validation on a single schema
func (v *SchemaValidator) validateSchema(schema *parser.Schema, path string, result *SchemaValidationResult) {
	if schema == nil {
		v.addError(result, "null_schema", "Schema is null", path, "", "Ensure schema is properly defined")
		return
	}

	// Validate basic schema structure
	v.validateBasicStructure(schema, path, result)

	// Validate type consistency
	v.validateTypeConsistency(schema, path, result)

	// Validate required fields
	v.validateRequiredFields(schema, path, result)

	// Validate property schemas recursively
	if schema.Properties != nil {
		for propName, propSchema := range schema.Properties {
			propPath := fmt.Sprintf("%s.properties.%s", path, propName)
			v.validateSchema(propSchema, propPath, result)
		}
	}

	// Validate array item schemas
	if schema.Items != nil {
		itemPath := fmt.Sprintf("%s.items", path)
		v.validateSchema(schema.Items, itemPath, result)
	}

	// Validate composition schemas (allOf, oneOf, anyOf)
	v.validateCompositionSchemas(schema, path, result)

	// Validate custom rules
	v.validateCustomRules(schema, path, result)
}

// validateBasicStructure validates basic schema structure
func (v *SchemaValidator) validateBasicStructure(schema *parser.Schema, path string, result *SchemaValidationResult) {
	// Check for missing type
	if schema.Type == "" && schema.Ref == "" && len(schema.AllOf) == 0 && len(schema.OneOf) == 0 && len(schema.AnyOf) == 0 {
		v.addWarning(result, "missing_type", "Schema missing type definition", path, "", "Add explicit type or use $ref")
	}

	// Check for invalid type combinations
	if schema.Type != "" && schema.Ref != "" {
		v.addError(result, "invalid_combination", "Schema cannot have both type and $ref", path, "", "Use either type or $ref, not both")
	}

	// Validate enum values match type
	if len(schema.Enum) > 0 && schema.Type != "" {
		v.validateEnumValues(schema, path, result)
	}

	// Validate format for string types
	if schema.Type == "string" && schema.Format != "" {
		v.validateStringFormat(schema, path, result)
	}
}

// validateTypeConsistency validates type consistency across the schema
func (v *SchemaValidator) validateTypeConsistency(schema *parser.Schema, path string, result *SchemaValidationResult) {
	switch schema.Type {
	case "object":
		if schema.Properties == nil && schema.AdditionalProperties == nil {
			v.addWarning(result, "empty_object", "Object type with no properties or additionalProperties", path, "", "Add properties or allow additionalProperties")
		}
	case "array":
		if schema.Items == nil {
			v.addError(result, "missing_items", "Array type missing items definition", path, "", "Add items schema for array type")
		}
	case "string":
		if schema.Properties != nil {
			v.addError(result, "type_mismatch", "String type cannot have properties", path, "", "Remove properties or change type to object")
		}
	case "number", "integer":
		if schema.Properties != nil {
			v.addError(result, "type_mismatch", "Numeric type cannot have properties", path, "", "Remove properties or change type to object")
		}
	case "boolean":
		if schema.Properties != nil {
			v.addError(result, "type_mismatch", "Boolean type cannot have properties", path, "", "Remove properties or change type to object")
		}
	}
}

// validateRequiredFields validates required field definitions
func (v *SchemaValidator) validateRequiredFields(schema *parser.Schema, path string, result *SchemaValidationResult) {
	if len(schema.Required) == 0 {
		return
	}

	if schema.Properties == nil {
		v.addError(result, "required_without_properties", "Required fields defined but no properties exist", path, "", "Add properties or remove required fields")
		return
	}

	// Check that all required fields exist in properties
	for _, requiredField := range schema.Required {
		if _, exists := schema.Properties[requiredField]; !exists {
			v.addError(result, "missing_required_property",
				fmt.Sprintf("Required field '%s' not found in properties", requiredField),
				path, requiredField,
				fmt.Sprintf("Add property '%s' or remove from required list", requiredField))
		}
	}

	// Check for duplicate required fields
	seen := make(map[string]bool)
	for _, requiredField := range schema.Required {
		if seen[requiredField] {
			v.addWarning(result, "duplicate_required",
				fmt.Sprintf("Duplicate required field '%s'", requiredField),
				path, requiredField,
				"Remove duplicate entry from required list")
		}
		seen[requiredField] = true
	}
}

// validateEnumValues validates enum values match the schema type
func (v *SchemaValidator) validateEnumValues(schema *parser.Schema, path string, result *SchemaValidationResult) {
	for i, enumValue := range schema.Enum {
		if !v.isValidEnumValue(enumValue, schema.Type) {
			v.addError(result, "invalid_enum_value",
				fmt.Sprintf("Enum value at index %d does not match type %s", i, schema.Type),
				path, "",
				fmt.Sprintf("Ensure enum value matches type %s", schema.Type))
		}
	}
}

// validateStringFormat validates string format specifications
func (v *SchemaValidator) validateStringFormat(schema *parser.Schema, path string, result *SchemaValidationResult) {
	validFormats := []string{
		"date", "date-time", "time", "email", "hostname", "ipv4", "ipv6",
		"uri", "uri-reference", "uuid", "password", "byte", "binary",
	}

	isValid := false
	for _, validFormat := range validFormats {
		if schema.Format == validFormat {
			isValid = true
			break
		}
	}

	if !isValid {
		v.addWarning(result, "unknown_format",
			fmt.Sprintf("Unknown string format '%s'", schema.Format),
			path, "",
			"Use a standard string format or remove format specification")
	}
}

// validateCompositionSchemas validates allOf, oneOf, anyOf schemas
func (v *SchemaValidator) validateCompositionSchemas(schema *parser.Schema, path string, result *SchemaValidationResult) {
	// Validate allOf
	if len(schema.AllOf) > 0 {
		for i, subSchema := range schema.AllOf {
			subPath := fmt.Sprintf("%s.allOf[%d]", path, i)
			v.validateSchema(subSchema, subPath, result)
		}
	}

	// Validate oneOf
	if len(schema.OneOf) > 0 {
		for i, subSchema := range schema.OneOf {
			subPath := fmt.Sprintf("%s.oneOf[%d]", path, i)
			v.validateSchema(subSchema, subPath, result)
		}
	}

	// Validate anyOf
	if len(schema.AnyOf) > 0 {
		for i, subSchema := range schema.AnyOf {
			subPath := fmt.Sprintf("%s.anyOf[%d]", path, i)
			v.validateSchema(subSchema, subPath, result)
		}
	}

	// Check for conflicting composition keywords
	compositionCount := 0
	if len(schema.AllOf) > 0 {
		compositionCount++
	}
	if len(schema.OneOf) > 0 {
		compositionCount++
	}
	if len(schema.AnyOf) > 0 {
		compositionCount++
	}

	if compositionCount > 1 {
		v.addWarning(result, "multiple_composition",
			"Schema uses multiple composition keywords (allOf, oneOf, anyOf)",
			path, "",
			"Consider using only one composition keyword for clarity")
	}
}

// validateCustomRules validates custom validation rules from configuration
func (v *SchemaValidator) validateCustomRules(schema *parser.Schema, path string, result *SchemaValidationResult) {
	// Apply custom validators from configuration
	for name, validator := range v.config.Validation.CustomValidators {
		if err := v.applyCustomValidator(schema, name, validator, path, result); err != nil {
			v.addError(result, "validator_error",
				fmt.Sprintf("Error applying custom validator %s: %v", name, err),
				path, "",
				"Check custom validator configuration")
		}
	}
}

// validatePathItem validates schemas in a path item
func (v *SchemaValidator) validatePathItem(pathItem *parser.PathItem, path string, result *SchemaValidationResult) {
	operations := map[string]*parser.Operation{
		"get":    pathItem.Get,
		"post":   pathItem.Post,
		"put":    pathItem.Put,
		"delete": pathItem.Delete,
		"patch":  pathItem.Patch,
	}

	for method, operation := range operations {
		if operation != nil {
			opPath := fmt.Sprintf("%s.%s", path, method)
			v.validateOperation(operation, opPath, result)
		}
	}
}

// validateOperation validates schemas in an operation
func (v *SchemaValidator) validateOperation(operation *parser.Operation, path string, result *SchemaValidationResult) {
	// Validate parameter schemas
	for i, param := range operation.Parameters {
		if param.Schema != nil {
			paramPath := fmt.Sprintf("%s.parameters[%d].schema", path, i)
			v.validateSchema(param.Schema, paramPath, result)
			result.Summary.TotalSchemas++
		}
	}

	// Validate request body schemas
	if operation.RequestBody != nil {
		for mediaType, media := range operation.RequestBody.Content {
			if media.Schema != nil {
				reqPath := fmt.Sprintf("%s.requestBody.content.%s.schema", path, mediaType)
				v.validateSchema(media.Schema, reqPath, result)
				result.Summary.TotalSchemas++
			}
		}
	}

	// Validate response schemas
	for statusCode, response := range operation.Responses {
		for mediaType, media := range response.Content {
			if media.Schema != nil {
				respPath := fmt.Sprintf("%s.responses.%s.content.%s.schema", path, statusCode, mediaType)
				v.validateSchema(media.Schema, respPath, result)
				result.Summary.TotalSchemas++
			}
		}
	}
}

// Helper functions

// addError adds an error to the validation result
func (v *SchemaValidator) addError(result *SchemaValidationResult, errorType, message, path, field, suggestion string) {
	result.Errors = append(result.Errors, SchemaError{
		Type:       errorType,
		Message:    message,
		Path:       path,
		Field:      field,
		Suggestion: suggestion,
	})
}

// addWarning adds a warning to the validation result
func (v *SchemaValidator) addWarning(result *SchemaValidationResult, warningType, message, path, field, suggestion string) {
	result.Warnings = append(result.Warnings, SchemaWarning{
		Type:       warningType,
		Message:    message,
		Path:       path,
		Field:      field,
		Suggestion: suggestion,
	})
}

// isValidEnumValue checks if an enum value matches the expected type
func (v *SchemaValidator) isValidEnumValue(value interface{}, schemaType string) bool {
	switch schemaType {
	case "string":
		_, ok := value.(string)
		return ok
	case "integer":
		switch value.(type) {
		case int, int32, int64:
			return true
		case float64:
			// JSON numbers are parsed as float64, check if it's actually an integer
			if f, ok := value.(float64); ok {
				return f == float64(int64(f))
			}
		}
		return false
	case "number":
		switch value.(type) {
		case int, int32, int64, float32, float64:
			return true
		}
		return false
	case "boolean":
		_, ok := value.(bool)
		return ok
	default:
		return true // Unknown type, assume valid
	}
}

// applyCustomValidator applies a custom validator to a schema
func (v *SchemaValidator) applyCustomValidator(schema *parser.Schema, name string, validator config.CustomValidator, path string, result *SchemaValidationResult) error {
	// This is a simplified implementation
	// In a real implementation, you would have more sophisticated validation logic

	// For now, just check if the schema description contains validation hints
	if schema.Description != "" && strings.Contains(schema.Description, fmt.Sprintf("[Validation: %s]", name)) {
		// Validation rule is already applied, consider it valid
		return nil
	}

	// If no validation hints found, add a warning
	v.addWarning(result, "missing_validation",
		fmt.Sprintf("Schema may need validation rule '%s'", name),
		path, "",
		fmt.Sprintf("Consider applying validation rule '%s'", name))

	return nil
}

// ValidateSchemaConsistency validates consistency across multiple schemas
func (v *SchemaValidator) ValidateSchemaConsistency(schemas map[string]*parser.Schema) (*SchemaValidationResult, error) {
	result := &SchemaValidationResult{
		Valid:    true,
		Errors:   make([]SchemaError, 0),
		Warnings: make([]SchemaWarning, 0),
		Summary: ValidationSummary{
			TotalSchemas: len(schemas),
		},
	}

	// Check for naming consistency
	v.validateNamingConsistency(schemas, result)

	// Check for type consistency
	v.validateTypeConsistencyAcrossSchemas(schemas, result)

	// Check for field consistency
	v.validateFieldConsistency(schemas, result)

	// Calculate summary
	result.Summary.ErrorCount = len(result.Errors)
	result.Summary.WarningCount = len(result.Warnings)
	result.Summary.ValidSchemas = result.Summary.TotalSchemas - result.Summary.ErrorCount

	if result.Summary.TotalSchemas > 0 {
		result.Summary.ConsistencyScore = float64(result.Summary.ValidSchemas) / float64(result.Summary.TotalSchemas) * 100
	}

	result.Valid = result.Summary.ErrorCount == 0

	return result, nil
}

// validateNamingConsistency validates naming consistency across schemas
func (v *SchemaValidator) validateNamingConsistency(schemas map[string]*parser.Schema, result *SchemaValidationResult) {
	// Check for consistent naming patterns
	for schemaName := range schemas {
		if !v.isValidSchemaName(schemaName) {
			v.addWarning(result, "inconsistent_naming",
				fmt.Sprintf("Schema name '%s' doesn't follow naming conventions", schemaName),
				fmt.Sprintf("schemas.%s", schemaName), "",
				"Use PascalCase for schema names")
		}
	}
}

// validateTypeConsistencyAcrossSchemas validates type consistency across schemas
func (v *SchemaValidator) validateTypeConsistencyAcrossSchemas(schemas map[string]*parser.Schema, result *SchemaValidationResult) {
	// Track field types across schemas to detect inconsistencies
	fieldTypes := make(map[string]map[string]string) // fieldName -> schemaName -> type

	for schemaName, schema := range schemas {
		if schema.Properties != nil {
			for fieldName, fieldSchema := range schema.Properties {
				if fieldTypes[fieldName] == nil {
					fieldTypes[fieldName] = make(map[string]string)
				}
				fieldTypes[fieldName][schemaName] = fieldSchema.Type
			}
		}
	}

	// Check for inconsistent field types
	for fieldName, schemaTypes := range fieldTypes {
		if len(schemaTypes) > 1 {
			types := make([]string, 0, len(schemaTypes))
			for _, fieldType := range schemaTypes {
				types = append(types, fieldType)
			}

			// Check if all types are the same
			firstType := types[0]
			isConsistent := true
			for _, t := range types[1:] {
				if t != firstType {
					isConsistent = false
					break
				}
			}

			if !isConsistent {
				v.addWarning(result, "inconsistent_field_types",
					fmt.Sprintf("Field '%s' has inconsistent types across schemas", fieldName),
					"", fieldName,
					"Ensure field types are consistent across all schemas")
			}
		}
	}
}

// validateFieldConsistency validates field consistency across schemas
func (v *SchemaValidator) validateFieldConsistency(schemas map[string]*parser.Schema, result *SchemaValidationResult) {
	// Check for common fields that should be consistent
	commonFields := []string{"id", "name", "description", "created_at", "updated_at"}

	for _, fieldName := range commonFields {
		v.validateCommonField(fieldName, schemas, result)
	}
}

// validateCommonField validates a common field across schemas
func (v *SchemaValidator) validateCommonField(fieldName string, schemas map[string]*parser.Schema, result *SchemaValidationResult) {
	fieldDefinitions := make(map[string]*parser.Schema)

	// Collect all definitions of this field
	for schemaName, schema := range schemas {
		if schema.Properties != nil {
			if fieldSchema, exists := schema.Properties[fieldName]; exists {
				fieldDefinitions[schemaName] = fieldSchema
			}
		}
	}

	if len(fieldDefinitions) <= 1 {
		return // Not enough instances to compare
	}

	// Check for consistency in type, format, and other properties
	var referenceType, referenceFormat string
	var referenceSchema *parser.Schema

	for schemaName, fieldSchema := range fieldDefinitions {
		if referenceSchema == nil {
			referenceSchema = fieldSchema
			referenceType = fieldSchema.Type
			referenceFormat = fieldSchema.Format
			continue
		}

		if fieldSchema.Type != referenceType {
			v.addWarning(result, "inconsistent_common_field",
				fmt.Sprintf("Common field '%s' has inconsistent type in schema '%s'", fieldName, schemaName),
				fmt.Sprintf("schemas.%s.properties.%s", schemaName, fieldName), fieldName,
				fmt.Sprintf("Use consistent type '%s' for field '%s'", referenceType, fieldName))
		}

		if fieldSchema.Format != referenceFormat {
			v.addWarning(result, "inconsistent_common_field",
				fmt.Sprintf("Common field '%s' has inconsistent format in schema '%s'", fieldName, schemaName),
				fmt.Sprintf("schemas.%s.properties.%s", schemaName, fieldName), fieldName,
				fmt.Sprintf("Use consistent format '%s' for field '%s'", referenceFormat, fieldName))
		}
	}
}

// isValidSchemaName checks if a schema name follows naming conventions
func (v *SchemaValidator) isValidSchemaName(name string) bool {
	// Check for PascalCase
	if len(name) == 0 {
		return false
	}

	// First character should be uppercase
	if name[0] < 'A' || name[0] > 'Z' {
		return false
	}

	// Should not contain underscores or hyphens
	if strings.Contains(name, "_") || strings.Contains(name, "-") {
		return false
	}

	return true
}
