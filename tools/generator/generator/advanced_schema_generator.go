package generator

import (
	"fmt"
	"strings"

	"github.com/mantisec/terraform-provider-umbrella/tools/generator/parser"
)

// AdvancedSchemaGenerator handles complex schema generation patterns
type AdvancedSchemaGenerator struct {
	templateEngine *TemplateEngine
}

// NewAdvancedSchemaGenerator creates a new advanced schema generator
func NewAdvancedSchemaGenerator(templateEngine *TemplateEngine) *AdvancedSchemaGenerator {
	return &AdvancedSchemaGenerator{
		templateEngine: templateEngine,
	}
}

// AdvancedSchemaAttribute represents an enhanced schema attribute with validation and modifiers
type AdvancedSchemaAttribute struct {
	SchemaAttribute
	Validators    []string
	PlanModifiers []string
	NestedType    *NestedObjectType
	IsNested      bool
	IsList        bool
	IsSet         bool
	ElementType   string
	MinItems      *int
	MaxItems      *int
	MinLength     *int
	MaxLength     *int
	Pattern       string
	Enum          []string
	Default       interface{}
}

// NestedObjectType represents a nested object schema
type NestedObjectType struct {
	Attributes []AdvancedSchemaAttribute
}

// GenerateAdvancedSchema generates an enhanced schema with proper validation and modifiers
func (asg *AdvancedSchemaGenerator) GenerateAdvancedSchema(endpoints []parser.Endpoint) *AdvancedResourceSchema {
	schema := &AdvancedResourceSchema{
		Attributes: []AdvancedSchemaAttribute{
			{
				SchemaAttribute: SchemaAttribute{
					Name:        "id",
					Type:        "types.String",
					Computed:    true,
					Description: "Resource identifier",
					GoType:      "string",
				},
				PlanModifiers: []string{"stringplanmodifier.UseStateForUnknown()"},
			},
		},
	}

	// Extract and enhance schema from endpoints
	for _, endpoint := range endpoints {
		if endpoint.Operation.RequestBody != nil {
			asg.extractAdvancedSchemaFromRequestBody(endpoint.Operation.RequestBody, schema)
		}

		for _, response := range endpoint.Operation.Responses {
			asg.extractAdvancedSchemaFromResponse(&response, schema)
		}
	}

	return schema
}

// AdvancedResourceSchema represents an enhanced resource schema
type AdvancedResourceSchema struct {
	Attributes []AdvancedSchemaAttribute
}

// extractAdvancedSchemaFromRequestBody extracts enhanced schema from request body
func (asg *AdvancedSchemaGenerator) extractAdvancedSchemaFromRequestBody(requestBody *parser.RequestBody, schema *AdvancedResourceSchema) {
	for _, mediaType := range requestBody.Content {
		if mediaType.Schema != nil {
			asg.extractAdvancedSchemaFromOpenAPISchema(mediaType.Schema, schema, true)
		}
	}
}

// extractAdvancedSchemaFromResponse extracts enhanced schema from response
func (asg *AdvancedSchemaGenerator) extractAdvancedSchemaFromResponse(response *parser.Response, schema *AdvancedResourceSchema) {
	for _, mediaType := range response.Content {
		if mediaType.Schema != nil {
			asg.extractAdvancedSchemaFromOpenAPISchema(mediaType.Schema, schema, false)
		}
	}
}

// extractAdvancedSchemaFromOpenAPISchema extracts enhanced attributes from OpenAPI schema
func (asg *AdvancedSchemaGenerator) extractAdvancedSchemaFromOpenAPISchema(apiSchema *parser.Schema, schema *AdvancedResourceSchema, isInput bool) {
	if apiSchema.Properties == nil {
		return
	}

	for propName, propSchema := range apiSchema.Properties {
		// Skip if attribute already exists
		if asg.attributeExists(schema, propName) {
			continue
		}

		attr := asg.createAdvancedAttribute(propName, propSchema, isInput, apiSchema.Required)
		schema.Attributes = append(schema.Attributes, attr)
	}
}

// createAdvancedAttribute creates an enhanced attribute with validation and modifiers
func (asg *AdvancedSchemaGenerator) createAdvancedAttribute(name string, schema *parser.Schema, isInput bool, required []string) AdvancedSchemaAttribute {
	attr := AdvancedSchemaAttribute{
		SchemaAttribute: SchemaAttribute{
			Name:        name,
			Type:        asg.determineAdvancedType(schema),
			GoType:      asg.templateEngine.schemaToGoType(schema),
			Description: schema.Description,
		},
		Validators:    asg.generateValidators(schema),
		PlanModifiers: asg.generatePlanModifiers(schema, isInput),
	}

	// Set required/optional/computed
	if isInput {
		attr.Required = asg.isRequired(name, required)
		attr.Optional = !attr.Required
	} else {
		attr.Computed = true
	}

	// Handle nested objects
	if schema.Type == "object" && schema.Properties != nil {
		attr.IsNested = true
		attr.NestedType = asg.createNestedObjectType(schema)
	}

	// Handle arrays
	if schema.Type == "array" {
		attr.IsList = true
		if schema.Items != nil {
			attr.ElementType = asg.determineAdvancedType(schema.Items)
			if schema.Items.Type == "object" {
				attr.IsNested = true
				attr.NestedType = asg.createNestedObjectType(schema.Items)
			}
		}
	}

	// Set constraints
	asg.setConstraints(&attr, schema)

	return attr
}

// determineAdvancedType determines the appropriate Terraform type for complex schemas
func (asg *AdvancedSchemaGenerator) determineAdvancedType(schema *parser.Schema) string {
	if schema == nil {
		return "types.String"
	}

	switch schema.Type {
	case "string":
		return "types.String"
	case "integer":
		return "types.Int64"
	case "number":
		return "types.Float64"
	case "boolean":
		return "types.Bool"
	case "array":
		if schema.Items != nil {
			switch schema.Items.Type {
			case "string":
				return "types.List" // or types.Set depending on uniqueness
			case "object":
				return "types.List" // List of nested objects
			default:
				return "types.List"
			}
		}
		return "types.List"
	case "object":
		return "types.Object"
	default:
		return "types.String"
	}
}

// generateValidators generates appropriate validators for a schema
func (asg *AdvancedSchemaGenerator) generateValidators(schema *parser.Schema) []string {
	var validators []string

	if schema == nil {
		return validators
	}

	switch schema.Type {
	case "string":
		if len(schema.Enum) > 0 {
			enumValues := make([]string, len(schema.Enum))
			for i, v := range schema.Enum {
				enumValues[i] = fmt.Sprintf("\"%v\"", v)
			}
			validators = append(validators, fmt.Sprintf("stringvalidator.OneOf(%s)", strings.Join(enumValues, ", ")))
		}

		// Add length validators if specified
		// Note: These would need to be extracted from OpenAPI constraints

	case "integer", "number":
		// Add range validators if specified
		// Note: These would need to be extracted from OpenAPI constraints

	case "array":
		// Add list validators if specified
		// Note: These would need to be extracted from OpenAPI constraints
	}

	return validators
}

// generatePlanModifiers generates appropriate plan modifiers for a schema
func (asg *AdvancedSchemaGenerator) generatePlanModifiers(schema *parser.Schema, isInput bool) []string {
	var modifiers []string

	if schema == nil {
		return modifiers
	}

	// Add UseStateForUnknown for computed attributes
	if !isInput {
		switch schema.Type {
		case "string":
			modifiers = append(modifiers, "stringplanmodifier.UseStateForUnknown()")
		case "integer", "number":
			modifiers = append(modifiers, "int64planmodifier.UseStateForUnknown()")
		case "boolean":
			modifiers = append(modifiers, "boolplanmodifier.UseStateForUnknown()")
		case "array":
			modifiers = append(modifiers, "listplanmodifier.UseStateForUnknown()")
		case "object":
			modifiers = append(modifiers, "objectplanmodifier.UseStateForUnknown()")
		}
	}

	// Add RequiresReplace for immutable attributes
	// This would be determined from API documentation or conventions

	return modifiers
}

// createNestedObjectType creates a nested object type from a schema
func (asg *AdvancedSchemaGenerator) createNestedObjectType(schema *parser.Schema) *NestedObjectType {
	if schema.Properties == nil {
		return nil
	}

	nestedType := &NestedObjectType{
		Attributes: make([]AdvancedSchemaAttribute, 0),
	}

	for propName, propSchema := range schema.Properties {
		attr := asg.createAdvancedAttribute(propName, propSchema, true, schema.Required)
		nestedType.Attributes = append(nestedType.Attributes, attr)
	}

	return nestedType
}

// setConstraints sets validation constraints from OpenAPI schema
func (asg *AdvancedSchemaGenerator) setConstraints(attr *AdvancedSchemaAttribute, schema *parser.Schema) {
	if schema == nil {
		return
	}

	// Set enum values
	if len(schema.Enum) > 0 {
		attr.Enum = make([]string, len(schema.Enum))
		for i, v := range schema.Enum {
			attr.Enum[i] = fmt.Sprintf("%v", v)
		}
	}

	// Set default value
	if schema.Default != nil {
		attr.Default = schema.Default
	}

	// Note: OpenAPI constraints like minLength, maxLength, minimum, maximum, pattern
	// would need to be extracted from the schema and converted to appropriate validators
}

// attributeExists checks if an attribute already exists in the schema
func (asg *AdvancedSchemaGenerator) attributeExists(schema *AdvancedResourceSchema, name string) bool {
	for _, attr := range schema.Attributes {
		if attr.Name == name {
			return true
		}
	}
	return false
}

// isRequired checks if a property is in the required list
func (asg *AdvancedSchemaGenerator) isRequired(propName string, required []string) bool {
	for _, req := range required {
		if req == propName {
			return true
		}
	}
	return false
}

// GenerateSchemaCode generates the Go code for the advanced schema
func (asg *AdvancedSchemaGenerator) GenerateSchemaCode(schema *AdvancedResourceSchema) string {
	var code strings.Builder

	code.WriteString("schema.Schema{\n")
	code.WriteString("\tAttributes: map[string]schema.Attribute{\n")

	for _, attr := range schema.Attributes {
		code.WriteString(fmt.Sprintf("\t\t\"%s\": ", attr.Name))
		code.WriteString(asg.generateAttributeCode(attr))
		code.WriteString(",\n")
	}

	code.WriteString("\t},\n")
	code.WriteString("}")

	return code.String()
}

// generateAttributeCode generates Go code for a single attribute
func (asg *AdvancedSchemaGenerator) generateAttributeCode(attr AdvancedSchemaAttribute) string {
	var code strings.Builder

	// Determine attribute type
	switch {
	case attr.IsNested && attr.IsList:
		code.WriteString("schema.ListNestedAttribute{\n")
		code.WriteString(asg.generateNestedAttributeCode(attr))
	case attr.IsNested:
		code.WriteString("schema.SingleNestedAttribute{\n")
		code.WriteString(asg.generateNestedAttributeCode(attr))
	case attr.IsList:
		code.WriteString("schema.ListAttribute{\n")
		code.WriteString(asg.generateListAttributeCode(attr))
	case attr.IsSet:
		code.WriteString("schema.SetAttribute{\n")
		code.WriteString(asg.generateSetAttributeCode(attr))
	default:
		code.WriteString(asg.generateSimpleAttributeCode(attr))
	}

	code.WriteString("\t}")

	return code.String()
}

// generateSimpleAttributeCode generates code for simple attributes
func (asg *AdvancedSchemaGenerator) generateSimpleAttributeCode(attr AdvancedSchemaAttribute) string {
	var code strings.Builder

	switch attr.Type {
	case "types.String":
		code.WriteString("schema.StringAttribute{\n")
	case "types.Int64":
		code.WriteString("schema.Int64Attribute{\n")
	case "types.Float64":
		code.WriteString("schema.Float64Attribute{\n")
	case "types.Bool":
		code.WriteString("schema.BoolAttribute{\n")
	default:
		code.WriteString("schema.StringAttribute{\n")
	}

	// Add common properties
	if attr.Required {
		code.WriteString("\t\tRequired: true,\n")
	}
	if attr.Optional {
		code.WriteString("\t\tOptional: true,\n")
	}
	if attr.Computed {
		code.WriteString("\t\tComputed: true,\n")
	}
	if attr.Description != "" {
		code.WriteString(fmt.Sprintf("\t\tDescription: \"%s\",\n", attr.Description))
	}

	// Add validators
	if len(attr.Validators) > 0 {
		code.WriteString("\t\tValidators: []validator.String{\n")
		for _, validator := range attr.Validators {
			code.WriteString(fmt.Sprintf("\t\t\t%s,\n", validator))
		}
		code.WriteString("\t\t},\n")
	}

	// Add plan modifiers
	if len(attr.PlanModifiers) > 0 {
		code.WriteString("\t\tPlanModifiers: []planmodifier.String{\n")
		for _, modifier := range attr.PlanModifiers {
			code.WriteString(fmt.Sprintf("\t\t\t%s,\n", modifier))
		}
		code.WriteString("\t\t},\n")
	}

	return code.String()
}

// generateListAttributeCode generates code for list attributes
func (asg *AdvancedSchemaGenerator) generateListAttributeCode(attr AdvancedSchemaAttribute) string {
	var code strings.Builder

	// Add common properties
	if attr.Required {
		code.WriteString("\t\tRequired: true,\n")
	}
	if attr.Optional {
		code.WriteString("\t\tOptional: true,\n")
	}
	if attr.Computed {
		code.WriteString("\t\tComputed: true,\n")
	}
	if attr.Description != "" {
		code.WriteString(fmt.Sprintf("\t\tDescription: \"%s\",\n", attr.Description))
	}

	// Add element type
	if attr.ElementType != "" {
		code.WriteString(fmt.Sprintf("\t\tElementType: %s,\n", attr.ElementType))
	}

	return code.String()
}

// generateSetAttributeCode generates code for set attributes
func (asg *AdvancedSchemaGenerator) generateSetAttributeCode(attr AdvancedSchemaAttribute) string {
	// Similar to list but for sets
	return asg.generateListAttributeCode(attr)
}

// generateNestedAttributeCode generates code for nested attributes
func (asg *AdvancedSchemaGenerator) generateNestedAttributeCode(attr AdvancedSchemaAttribute) string {
	var code strings.Builder

	// Add common properties
	if attr.Required {
		code.WriteString("\t\tRequired: true,\n")
	}
	if attr.Optional {
		code.WriteString("\t\tOptional: true,\n")
	}
	if attr.Computed {
		code.WriteString("\t\tComputed: true,\n")
	}
	if attr.Description != "" {
		code.WriteString(fmt.Sprintf("\t\tDescription: \"%s\",\n", attr.Description))
	}

	// Add nested attributes
	if attr.NestedType != nil {
		code.WriteString("\t\tNestedObject: schema.NestedAttributeObject{\n")
		code.WriteString("\t\t\tAttributes: map[string]schema.Attribute{\n")

		for _, nestedAttr := range attr.NestedType.Attributes {
			code.WriteString(fmt.Sprintf("\t\t\t\t\"%s\": ", nestedAttr.Name))
			code.WriteString(asg.generateAttributeCode(nestedAttr))
			code.WriteString(",\n")
		}

		code.WriteString("\t\t\t},\n")
		code.WriteString("\t\t},\n")
	}

	return code.String()
}
