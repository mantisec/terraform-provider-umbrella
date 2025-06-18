package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mantisec/terraform-provider-umbrella/tools/generator/config"
	"github.com/mantisec/terraform-provider-umbrella/tools/generator/parser"
)

// DocsGenerator handles generation of documentation for resources and data sources
type DocsGenerator struct {
	config         *config.Config
	templateEngine *TemplateEngine
}

// NewDocsGenerator creates a new documentation generator
func NewDocsGenerator(cfg *config.Config, templateEngine *TemplateEngine) *DocsGenerator {
	return &DocsGenerator{
		config:         cfg,
		templateEngine: templateEngine,
	}
}

// DocumentationData contains data for documentation template generation
type DocumentationData struct {
	ResourceName  string
	ResourceType  string // "resource" or "data_source"
	TypeName      string
	Description   string
	Attributes    []DocumentationAttribute
	Examples      []DocumentationExample
	ImportExample string
	Endpoints     []parser.Endpoint
}

// DocumentationAttribute represents an attribute in the documentation
type DocumentationAttribute struct {
	Name        string
	Type        string
	Required    bool
	Optional    bool
	Computed    bool
	Description string
	Example     string
}

// DocumentationExample represents a usage example
type DocumentationExample struct {
	Title       string
	Description string
	Code        string
}

// GenerateDocumentation generates documentation for a resource or data source
func (dg *DocsGenerator) GenerateDocumentation(resourceName string, endpoints []parser.Endpoint, resourceType string, outputDir string) error {
	// Prepare documentation data
	data := dg.prepareDocumentationData(resourceName, endpoints, resourceType)

	// Register the documentation template
	if err := dg.registerDocumentationTemplate(); err != nil {
		return fmt.Errorf("failed to register documentation template: %w", err)
	}

	// Generate the documentation
	content, err := dg.templateEngine.ExecuteTemplate("docs.md", data)
	if err != nil {
		return fmt.Errorf("failed to execute documentation template: %w", err)
	}

	// Write to file
	filename := fmt.Sprintf("%s.md", resourceName)
	if resourceType == "data_source" {
		filename = fmt.Sprintf("data-sources/%s.md", resourceName)
	} else {
		filename = fmt.Sprintf("resources/%s.md", resourceName)
	}

	outputPath := filepath.Join(outputDir, filename)

	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	if err := os.WriteFile(outputPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write documentation file: %w", err)
	}

	return nil
}

// prepareDocumentationData prepares the data structure for documentation template execution
func (dg *DocsGenerator) prepareDocumentationData(resourceName string, endpoints []parser.Endpoint, resourceType string) *DocumentationData {
	data := &DocumentationData{
		ResourceName:  resourceName,
		ResourceType:  resourceType,
		TypeName:      dg.config.Global.ProviderName + "_" + resourceName,
		Description:   dg.generateDescription(resourceName, endpoints),
		Attributes:    dg.generateAttributeDocumentation(endpoints),
		Examples:      dg.generateExamples(resourceName, endpoints, resourceType),
		ImportExample: dg.generateImportExample(resourceName),
		Endpoints:     endpoints,
	}

	return data
}

// generateDescription generates a description for the resource
func (dg *DocsGenerator) generateDescription(resourceName string, endpoints []parser.Endpoint) string {
	if len(endpoints) > 0 && endpoints[0].Operation.Description != "" {
		return endpoints[0].Operation.Description
	}

	if len(endpoints) > 0 && endpoints[0].Operation.Summary != "" {
		return endpoints[0].Operation.Summary
	}

	return fmt.Sprintf("Manages %s resources in Cisco Umbrella", strings.ReplaceAll(resourceName, "_", " "))
}

// generateAttributeDocumentation generates documentation for resource attributes
func (dg *DocsGenerator) generateAttributeDocumentation(endpoints []parser.Endpoint) []DocumentationAttribute {
	// Use the same enhanced schema generation logic as the resource generator
	schema := dg.generateEnhancedSchema(endpoints)

	var attributes []DocumentationAttribute
	for _, attr := range schema.Attributes {
		docAttr := DocumentationAttribute{
			Name:        attr.Name,
			Type:        dg.terraformTypeToDocType(attr.Type),
			Required:    attr.Required,
			Optional:    attr.Optional,
			Computed:    attr.Computed,
			Description: attr.Description,
			Example:     dg.generateExampleForType(attr.Type, attr.Name),
		}
		attributes = append(attributes, docAttr)
	}

	return attributes
}

// generateEnhancedSchema generates the enhanced schema using the same logic as resource generator
func (dg *DocsGenerator) generateEnhancedSchema(endpoints []parser.Endpoint) *ResourceSchema {
	schema := &ResourceSchema{
		Attributes: []SchemaAttribute{
			{
				Name:        "id",
				Type:        "types.String",
				Computed:    true,
				Description: "Resource identifier",
				GoType:      "string",
			},
		},
	}

	// Track processed schemas to avoid duplicates
	processedFields := make(map[string]bool)
	processedFields["id"] = true

	// First pass: Extract from request bodies (for input fields)
	hasCreateEndpoint := false
	for _, endpoint := range endpoints {
		if endpoint.CRUDType == "create" && endpoint.Operation.RequestBody != nil {
			hasCreateEndpoint = true
			dg.extractSchemaFromRequestBodyEnhanced(endpoint.Operation.RequestBody, schema, processedFields, true)
		}
	}

	// Second pass: Extract from successful responses (for computed fields and missing input fields)
	for _, endpoint := range endpoints {
		for statusCode, response := range endpoint.Operation.Responses {
			isSuccessResponse := statusCode == "200" || statusCode == "201" || statusCode == "202"
			if isSuccessResponse {
				// If we don't have a create endpoint, treat response fields as potential inputs
				isInput := !hasCreateEndpoint && endpoint.CRUDType == "create"
				dg.extractSchemaFromResponseEnhanced(&response, schema, processedFields, isInput)
			}
		}
	}

	// If we still have minimal schema, add common attributes
	if len(schema.Attributes) <= 1 {
		dg.addCommonSchemaAttributesForDocs("", schema, processedFields, false)
	}

	return schema
}

// extractSchemaFromRequestBodyEnhanced extracts schema attributes from request body
func (dg *DocsGenerator) extractSchemaFromRequestBodyEnhanced(requestBody *parser.RequestBody, schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
	for _, mediaType := range requestBody.Content {
		if mediaType.Schema != nil {
			dg.extractSchemaFromOpenAPISchemaEnhanced(mediaType.Schema, schema, processedFields, isInput)
		}
	}
}

// extractSchemaFromResponseEnhanced extracts schema attributes from response
func (dg *DocsGenerator) extractSchemaFromResponseEnhanced(response *parser.Response, schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
	for _, mediaType := range response.Content {
		if mediaType.Schema != nil {
			dg.extractSchemaFromOpenAPISchemaEnhanced(mediaType.Schema, schema, processedFields, isInput)
		}
	}
}

// extractSchemaFromOpenAPISchemaEnhanced extracts attributes from an OpenAPI schema
func (dg *DocsGenerator) extractSchemaFromOpenAPISchemaEnhanced(apiSchema *parser.Schema, schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
	if apiSchema == nil {
		return
	}

	// Handle schema references
	if apiSchema.Ref != "" {
		dg.handleSchemaReferenceForDocs(apiSchema.Ref, schema, processedFields, isInput)
		return
	}

	// Handle direct properties
	if apiSchema.Properties != nil {
		for propName, propSchema := range apiSchema.Properties {
			// Skip if already processed
			if processedFields[propName] {
				continue
			}

			attr := dg.createSchemaAttributeForDocs(propName, propSchema, apiSchema.Required, isInput)
			if attr != nil {
				schema.Attributes = append(schema.Attributes, *attr)
				processedFields[propName] = true
			}
		}
	} else {
		// If we have no properties and no reference, add common attributes
		if apiSchema.Type == "" && apiSchema.Ref == "" {
			dg.addCommonSchemaAttributesForDocs("", schema, processedFields, isInput)
		}
	}

	// Handle nested schemas in data wrappers
	if apiSchema.Type == "object" && len(apiSchema.Properties) > 0 {
		if dataSchema, exists := apiSchema.Properties["data"]; exists {
			dg.extractSchemaFromOpenAPISchemaEnhanced(dataSchema, schema, processedFields, isInput)
		}
	}

	// Handle array items
	if apiSchema.Type == "array" && apiSchema.Items != nil {
		dg.extractSchemaFromOpenAPISchemaEnhanced(apiSchema.Items, schema, processedFields, isInput)
	}

	// Handle allOf, oneOf, anyOf
	for _, subSchema := range apiSchema.AllOf {
		dg.extractSchemaFromOpenAPISchemaEnhanced(subSchema, schema, processedFields, isInput)
	}
	for _, subSchema := range apiSchema.OneOf {
		dg.extractSchemaFromOpenAPISchemaEnhanced(subSchema, schema, processedFields, isInput)
	}
	for _, subSchema := range apiSchema.AnyOf {
		dg.extractSchemaFromOpenAPISchemaEnhanced(subSchema, schema, processedFields, isInput)
	}
}

// handleSchemaReferenceForDocs handles unresolved schema references
func (dg *DocsGenerator) handleSchemaReferenceForDocs(ref string, schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
	schemaName := strings.TrimPrefix(ref, "#/components/schemas/")
	dg.addCommonSchemaAttributesForDocs(schemaName, schema, processedFields, isInput)
}

// addCommonSchemaAttributesForDocs adds attributes based on common schema patterns
func (dg *DocsGenerator) addCommonSchemaAttributesForDocs(schemaName string, schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
	// Define common attributes that appear in most API resources
	commonAttributes := []struct {
		name        string
		tfType      string
		goType      string
		required    bool
		computed    bool
		description string
	}{
		{"name", "types.String", "string", true, false, "The name of the resource"},
		{"description", "types.String", "string", false, false, "The description of the resource"},
		{"enabled", "types.Bool", "bool", false, false, "Whether the resource is enabled"},
		{"active", "types.Bool", "bool", false, false, "Whether the resource is active"},
		{"status", "types.String", "string", false, true, "The status of the resource"},
		{"organization_id", "types.Int64", "int64", false, true, "The organization ID"},
		{"created_at", "types.Int64", "int64", false, true, "The date and time when the resource was created"},
		{"modified_at", "types.Int64", "int64", false, true, "The date and time when the resource was modified"},
		{"updated_at", "types.Int64", "int64", false, true, "The date and time when the resource was updated"},
		{"created_by", "types.String", "string", false, true, "The user who created the resource"},
		{"modified_by", "types.String", "string", false, true, "The user who modified the resource"},
	}

	// Add schema-specific attributes based on naming patterns
	schemaLower := strings.ToLower(schemaName)

	// For destination list schemas, add specific attributes
	if strings.Contains(schemaLower, "destinationlist") {
		dg.addDestinationListSpecificAttributesForDocs(schema, processedFields, isInput)
		return
	}

	// For tunnel/VPN schemas
	if strings.Contains(schemaLower, "tunnel") || strings.Contains(schemaLower, "vpn") || strings.Contains(schemaLower, "ipsec") {
		dg.addTunnelSpecificAttributesForDocs(schema, processedFields, isInput)
	}

	// For rule/policy schemas
	if strings.Contains(schemaLower, "rule") || strings.Contains(schemaLower, "policy") {
		dg.addRuleSpecificAttributesForDocs(schema, processedFields, isInput)
	}

	// For SAML schemas
	if strings.Contains(schemaLower, "saml") {
		dg.addSAMLSpecificAttributesForDocs(schema, processedFields, isInput)
	}

	// Add common attributes that haven't been processed yet
	for _, attr := range commonAttributes {
		if !processedFields[attr.name] {
			schemaAttr := &SchemaAttribute{
				Name:        attr.name,
				Type:        attr.tfType,
				GoType:      attr.goType,
				Description: attr.description,
			}

			// Proper field classification logic
			if attr.computed {
				schemaAttr.Computed = true
				schemaAttr.Required = false
				schemaAttr.Optional = false
			} else if isInput {
				schemaAttr.Required = attr.required
				schemaAttr.Optional = !attr.required
				schemaAttr.Computed = false
			} else {
				schemaAttr.Computed = true
				schemaAttr.Required = false
				schemaAttr.Optional = false
			}

			schema.Attributes = append(schema.Attributes, *schemaAttr)
			processedFields[attr.name] = true
		}
	}
}

// terraformTypeToDocType converts Terraform types to documentation types
func (dg *DocsGenerator) terraformTypeToDocType(tfType string) string {
	switch tfType {
	case "types.String":
		return "String"
	case "types.Int64":
		return "Number"
	case "types.Bool":
		return "Boolean"
	case "types.Set":
		return "Set of String"
	case "types.List":
		return "List"
	default:
		return "String"
	}
}

// generateExampleForType generates an example value for a Terraform type
func (dg *DocsGenerator) generateExampleForType(tfType, fieldName string) string {
	switch tfType {
	case "types.String":
		if strings.Contains(strings.ToLower(fieldName), "email") {
			return "\"user@example.com\""
		}
		if strings.Contains(strings.ToLower(fieldName), "url") {
			return "\"https://example.com\""
		}
		if strings.Contains(strings.ToLower(fieldName), "name") {
			return "\"example-name\""
		}
		return "\"example\""
	case "types.Int64":
		return "123"
	case "types.Bool":
		return "true"
	case "types.Set":
		return "[\"item1\", \"item2\"]"
	case "types.List":
		return "[\"item1\", \"item2\"]"
	default:
		return "\"example\""
	}
}

// createSchemaAttributeForDocs creates a schema attribute from an OpenAPI property for documentation
func (dg *DocsGenerator) createSchemaAttributeForDocs(propName string, propSchema *parser.Schema, requiredFields []string, isInput bool) *SchemaAttribute {
	if propSchema == nil {
		return nil
	}

	// Skip certain system fields that shouldn't be in Terraform schema
	skipFields := map[string]bool{
		"status":     true, // API status responses
		"meta":       true, // Pagination metadata
		"txId":       true, // Transaction IDs
		"statusCode": true, // HTTP status codes
		"error":      true, // Error fields
	}

	if skipFields[propName] {
		return nil
	}

	attr := &SchemaAttribute{
		Name:        propName,
		Type:        dg.schemaToTerraformTypeForDocs(propSchema),
		GoType:      dg.schemaToGoTypeForDocs(propSchema),
		Description: dg.cleanDescriptionForDocs(propSchema.Description),
	}

	// Determine if required/optional/computed
	if isInput {
		attr.Required = dg.isRequiredForDocs(propName, requiredFields)
		attr.Optional = !attr.Required
	} else {
		attr.Computed = true
	}

	// Handle special cases
	switch propName {
	case "id":
		attr.Computed = true
		attr.Required = false
		attr.Optional = false
	case "organizationId", "organization_id":
		attr.Computed = true
		attr.Required = false
		attr.Optional = false
	case "createdAt", "created_at", "modifiedAt", "modified_at":
		attr.Computed = true
		attr.Required = false
		attr.Optional = false
	}

	return attr
}

// schemaToTerraformTypeForDocs converts OpenAPI schema to Terraform type for docs
func (dg *DocsGenerator) schemaToTerraformTypeForDocs(schema *parser.Schema) string {
	if schema == nil {
		return "types.String"
	}

	switch schema.Type {
	case "string":
		return "types.String"
	case "integer", "number":
		return "types.Int64"
	case "boolean":
		return "types.Bool"
	case "array":
		return "types.Set"
	case "object":
		return "types.String" // JSON string representation
	default:
		return "types.String"
	}
}

// schemaToGoTypeForDocs converts OpenAPI schema to Go type for docs
func (dg *DocsGenerator) schemaToGoTypeForDocs(schema *parser.Schema) string {
	if schema == nil {
		return "string"
	}

	switch schema.Type {
	case "string":
		return "string"
	case "integer", "number":
		return "int64"
	case "boolean":
		return "bool"
	case "array":
		return "[]string"
	case "object":
		return "string"
	default:
		return "string"
	}
}

// cleanDescriptionForDocs cleans up description text for documentation
func (dg *DocsGenerator) cleanDescriptionForDocs(desc string) string {
	if desc == "" {
		return ""
	}

	// Basic cleanup
	desc = strings.TrimSpace(desc)
	desc = strings.ReplaceAll(desc, "\n", " ")
	desc = strings.ReplaceAll(desc, "\r", " ")
	desc = strings.ReplaceAll(desc, "\t", " ")

	// Remove multiple spaces
	for strings.Contains(desc, "  ") {
		desc = strings.ReplaceAll(desc, "  ", " ")
	}

	return desc
}

// isRequiredForDocs checks if a property is in the required list for docs
func (dg *DocsGenerator) isRequiredForDocs(propName string, required []string) bool {
	for _, req := range required {
		if req == propName {
			return true
		}
	}
	return false
}

// addDestinationListSpecificAttributesForDocs adds destination list specific attributes for docs
func (dg *DocsGenerator) addDestinationListSpecificAttributesForDocs(schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
	attributes := []struct {
		name        string
		tfType      string
		goType      string
		required    bool
		computed    bool
		description string
	}{
		{"access", "types.String", "string", true, false, "The type of access for the destination list (allow/block)"},
		{"is_global", "types.Bool", "bool", true, false, "Specifies whether the destination list is a global destination list"},
		{"bundle_type_id", "types.Int64", "int64", false, false, "The type of the destination list in the policy"},
		{"thirdparty_category_id", "types.Int64", "int64", false, true, "The third-party category ID of the destination list"},
		{"is_msp_default", "types.Bool", "bool", false, true, "Specifies whether MSP is the default"},
		{"marked_for_deletion", "types.Bool", "bool", false, true, "Specifies whether the destination list is marked for deletion"},
	}

	dg.addAttributesToSchemaForDocs(attributes, schema, processedFields, isInput)
}

// addTunnelSpecificAttributesForDocs adds tunnel/VPN specific attributes for docs
func (dg *DocsGenerator) addTunnelSpecificAttributesForDocs(schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
	attributes := []struct {
		name        string
		tfType      string
		goType      string
		required    bool
		computed    bool
		description string
	}{
		{"tunnel_name", "types.String", "string", true, false, "The name of the tunnel"},
		{"remote_gateway", "types.String", "string", true, false, "The remote gateway IP address"},
		{"local_gateway", "types.String", "string", false, false, "The local gateway IP address"},
		{"preshared_key", "types.String", "string", true, false, "The pre-shared key for authentication"},
		{"encryption", "types.String", "string", false, false, "The encryption algorithm"},
		{"authentication", "types.String", "string", false, false, "The authentication algorithm"},
	}

	dg.addAttributesToSchemaForDocs(attributes, schema, processedFields, isInput)
}

// addRuleSpecificAttributesForDocs adds rule/policy specific attributes for docs
func (dg *DocsGenerator) addRuleSpecificAttributesForDocs(schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
	attributes := []struct {
		name        string
		tfType      string
		goType      string
		required    bool
		computed    bool
		description string
	}{
		{"action", "types.String", "string", true, false, "The action to take (allow/block)"},
		{"priority", "types.Int64", "int64", false, false, "The priority of the rule"},
		{"source", "types.String", "string", false, false, "The source of the rule"},
		{"destination", "types.String", "string", false, false, "The destination of the rule"},
		{"protocol", "types.String", "string", false, false, "The protocol for the rule"},
		{"port", "types.String", "string", false, false, "The port for the rule"},
	}

	dg.addAttributesToSchemaForDocs(attributes, schema, processedFields, isInput)
}

// addSAMLSpecificAttributesForDocs adds SAML specific attributes for docs
func (dg *DocsGenerator) addSAMLSpecificAttributesForDocs(schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
	attributes := []struct {
		name        string
		tfType      string
		goType      string
		required    bool
		computed    bool
		description string
	}{
		{"entity_id", "types.String", "string", true, false, "The SAML entity ID"},
		{"sso_url", "types.String", "string", true, false, "The SSO URL"},
		{"certificate", "types.String", "string", false, false, "The SAML certificate"},
		{"attribute_mapping", "types.String", "string", false, false, "The attribute mapping configuration"},
	}

	dg.addAttributesToSchemaForDocs(attributes, schema, processedFields, isInput)
}

// addAttributesToSchemaForDocs is a helper method to add attributes to schema for docs
func (dg *DocsGenerator) addAttributesToSchemaForDocs(attributes []struct {
	name        string
	tfType      string
	goType      string
	required    bool
	computed    bool
	description string
}, schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
	for _, attr := range attributes {
		if !processedFields[attr.name] {
			schemaAttr := &SchemaAttribute{
				Name:        attr.name,
				Type:        attr.tfType,
				GoType:      attr.goType,
				Description: attr.description,
			}

			// Proper field classification logic
			if attr.computed {
				schemaAttr.Computed = true
				schemaAttr.Required = false
				schemaAttr.Optional = false
			} else if isInput {
				schemaAttr.Required = attr.required
				schemaAttr.Optional = !attr.required
				schemaAttr.Computed = false
			} else {
				schemaAttr.Computed = true
				schemaAttr.Required = false
				schemaAttr.Optional = false
			}

			schema.Attributes = append(schema.Attributes, *schemaAttr)
			processedFields[attr.name] = true
		}
	}
}

// extractAttributesFromRequestBody extracts attributes from request body schema
func (dg *DocsGenerator) extractAttributesFromRequestBody(requestBody *parser.RequestBody) []DocumentationAttribute {
	var attributes []DocumentationAttribute

	for _, mediaType := range requestBody.Content {
		if mediaType.Schema != nil && mediaType.Schema.Properties != nil {
			for propName, propSchema := range mediaType.Schema.Properties {
				attr := DocumentationAttribute{
					Name:        propName,
					Type:        dg.schemaTypeToDocType(propSchema),
					Required:    dg.isPropertyRequired(propName, mediaType.Schema.Required),
					Optional:    !dg.isPropertyRequired(propName, mediaType.Schema.Required),
					Description: propSchema.Description,
					Example:     dg.generateExampleValue(propSchema),
				}
				attributes = append(attributes, attr)
			}
		}
	}

	return attributes
}

// extractAttributesFromResponse extracts attributes from response schema
func (dg *DocsGenerator) extractAttributesFromResponse(response *parser.Response) []DocumentationAttribute {
	var attributes []DocumentationAttribute

	for _, mediaType := range response.Content {
		if mediaType.Schema != nil && mediaType.Schema.Properties != nil {
			for propName, propSchema := range mediaType.Schema.Properties {
				attr := DocumentationAttribute{
					Name:        propName,
					Type:        dg.schemaTypeToDocType(propSchema),
					Computed:    true,
					Description: propSchema.Description,
					Example:     dg.generateExampleValue(propSchema),
				}
				attributes = append(attributes, attr)
			}
		}
	}

	return attributes
}

// schemaTypeToDocType converts OpenAPI schema type to documentation type
func (dg *DocsGenerator) schemaTypeToDocType(schema *parser.Schema) string {
	if schema == nil {
		return "String"
	}

	switch schema.Type {
	case "string":
		return "String"
	case "integer", "number":
		return "Number"
	case "boolean":
		return "Boolean"
	case "array":
		if schema.Items != nil {
			itemType := dg.schemaTypeToDocType(schema.Items)
			return fmt.Sprintf("List of %s", itemType)
		}
		return "List"
	case "object":
		return "Object"
	default:
		return "String"
	}
}

// generateExampleValue generates an example value for a schema
func (dg *DocsGenerator) generateExampleValue(schema *parser.Schema) string {
	if schema == nil {
		return "\"\""
	}

	if schema.Example != nil {
		return fmt.Sprintf("\"%v\"", schema.Example)
	}

	switch schema.Type {
	case "string":
		if schema.Format == "email" {
			return "\"user@example.com\""
		}
		if schema.Format == "uri" {
			return "\"https://example.com\""
		}
		return "\"example\""
	case "integer":
		return "123"
	case "number":
		return "123.45"
	case "boolean":
		return "true"
	case "array":
		return "[\"item1\", \"item2\"]"
	case "object":
		return "{}"
	default:
		return "\"\""
	}
}

// isPropertyRequired checks if a property is required
func (dg *DocsGenerator) isPropertyRequired(propName string, required []string) bool {
	for _, req := range required {
		if req == propName {
			return true
		}
	}
	return false
}

// removeDuplicateAttributes removes duplicate attributes from the list
func (dg *DocsGenerator) removeDuplicateAttributes(attributes []DocumentationAttribute) []DocumentationAttribute {
	seen := make(map[string]bool)
	var result []DocumentationAttribute

	for _, attr := range attributes {
		if !seen[attr.Name] {
			seen[attr.Name] = true
			result = append(result, attr)
		}
	}

	return result
}

// generateExamples generates usage examples for the resource
func (dg *DocsGenerator) generateExamples(resourceName string, endpoints []parser.Endpoint, resourceType string) []DocumentationExample {
	examples := []DocumentationExample{}

	if resourceType == "resource" {
		examples = append(examples, DocumentationExample{
			Title:       "Basic Usage",
			Description: fmt.Sprintf("Basic usage of the %s resource", resourceName),
			Code:        dg.generateBasicResourceExample(resourceName),
		})
	} else {
		examples = append(examples, DocumentationExample{
			Title:       "Basic Usage",
			Description: fmt.Sprintf("Basic usage of the %s data source", resourceName),
			Code:        dg.generateBasicDataSourceExample(resourceName),
		})
	}

	return examples
}

// generateBasicResourceExample generates a basic resource usage example
func (dg *DocsGenerator) generateBasicResourceExample(resourceName string) string {
	typeName := dg.config.Global.ProviderName + "_" + resourceName

	// Generate example based on resource type
	switch {
	case strings.Contains(strings.ToLower(resourceName), "destinationlist"):
		return fmt.Sprintf(`resource "%s" "example" {
  name      = "example-destination-list"
  access    = "allow"
  is_global = false
}`, typeName)
	case strings.Contains(strings.ToLower(resourceName), "tunnel"):
		return fmt.Sprintf(`resource "%s" "example" {
  name           = "example-tunnel"
  remote_gateway = "192.168.1.1"
  preshared_key  = "your-preshared-key"
}`, typeName)
	case strings.Contains(strings.ToLower(resourceName), "user"):
		return fmt.Sprintf(`resource "%s" "example" {
  name        = "example-user"
  description = "Example user account"
  enabled     = true
}`, typeName)
	case strings.Contains(strings.ToLower(resourceName), "network"):
		return fmt.Sprintf(`resource "%s" "example" {
  name        = "example-network"
  description = "Example network configuration"
  enabled     = true
}`, typeName)
	default:
		return fmt.Sprintf(`resource "%s" "example" {
  name        = "example-%s"
  description = "Example %s resource"
}`, typeName, resourceName, strings.ReplaceAll(resourceName, "_", " "))
	}
}

// generateBasicDataSourceExample generates a basic data source usage example
func (dg *DocsGenerator) generateBasicDataSourceExample(resourceName string) string {
	typeName := dg.config.Global.ProviderName + "_" + resourceName
	return fmt.Sprintf(`data "%s" "example" {
  # Add filter attributes here
  id = "12345"
}`, typeName)
}

// generateImportExample generates an import example for resources
func (dg *DocsGenerator) generateImportExample(resourceName string) string {
	typeName := dg.config.Global.ProviderName + "_" + resourceName
	return fmt.Sprintf(`terraform import %s.example 12345`, typeName)
}

// registerDocumentationTemplate registers the embedded documentation template
func (dg *DocsGenerator) registerDocumentationTemplate() error {
	template := `---
page_title: "{{.TypeName}} {{if eq .ResourceType "resource"}}Resource{{else}}Data Source{{end}} - terraform-provider-umbrella"
description: |-
  {{.Description}}
---

# {{.TypeName}} ({{if eq .ResourceType "resource"}}Resource{{else}}Data Source{{end}})

{{.Description}}

## Example Usage

{{range .Examples}}
### {{.Title}}

{{.Description}}

` + "```terraform" + `
{{.Code}}
` + "```" + `

{{end}}

## Argument Reference

The following arguments are supported:

### Required

{{range .Attributes}}{{if .Required}}- **` + "`{{.Name}}`" + `** ({{.Type}}) - {{.Description}}{{if .Example}} Example: ` + "`{{.Example}}`" + `{{end}}
{{end}}{{end}}

### Optional

{{range .Attributes}}{{if .Optional}}- **` + "`{{.Name}}`" + `** ({{.Type}}) - {{.Description}}{{if .Example}} Example: ` + "`{{.Example}}`" + `{{end}}
{{end}}{{end}}

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

{{range .Attributes}}{{if .Computed}}- **` + "`{{.Name}}`" + `** ({{.Type}}) - {{.Description}}
{{end}}{{end}}

{{if eq .ResourceType "resource"}}
## Import

{{.TypeName}} can be imported using the resource ID:

` + "```shell" + `
{{.ImportExample}}
` + "```" + `

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.
{{end}}
`

	return dg.templateEngine.RegisterTemplate("docs.md", template)
}
