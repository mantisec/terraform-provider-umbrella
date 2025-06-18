package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mantisec/terraform-provider-umbrella/tools/generator/config"
	"github.com/mantisec/terraform-provider-umbrella/tools/generator/parser"
)

// ResourceGenerator handles generation of Terraform resources
type ResourceGenerator struct {
	config         *config.Config
	templateEngine *TemplateEngine
}

// NewResourceGenerator creates a new resource generator
func NewResourceGenerator(cfg *config.Config, templateEngine *TemplateEngine) *ResourceGenerator {
	return &ResourceGenerator{
		config:         cfg,
		templateEngine: templateEngine,
	}
}

// ResourceData contains data for resource template generation
type ResourceData struct {
	ResourceName   string
	StructName     string
	TypeName       string
	PackageName    string
	ProviderName   string
	Endpoints      []parser.Endpoint
	CreateEndpoint *parser.Endpoint
	ReadEndpoint   *parser.Endpoint
	UpdateEndpoint *parser.Endpoint
	DeleteEndpoint *parser.Endpoint
	Schema         *ResourceSchema
	Config         *config.Config
}

// ResourceSchema represents the Terraform schema for a resource
type ResourceSchema struct {
	Attributes []SchemaAttribute
}

// SchemaAttribute represents a single attribute in the schema
type SchemaAttribute struct {
	Name        string
	Type        string
	Required    bool
	Optional    bool
	Computed    bool
	Description string
	GoType      string
}

// GenerateResource generates a Terraform resource from endpoints
func (rg *ResourceGenerator) GenerateResource(resourceName string, endpoints []parser.Endpoint, outputDir string) error {
	// Prepare resource data
	data := rg.prepareResourceData(resourceName, endpoints)

	// Register the embedded template
	if err := rg.registerResourceTemplate(); err != nil {
		return fmt.Errorf("failed to register resource template: %w", err)
	}

	// Generate the resource code
	code, err := rg.templateEngine.ExecuteTemplate("resource.go", data)
	if err != nil {
		return fmt.Errorf("failed to execute resource template: %w", err)
	}

	// Write to file
	filename := fmt.Sprintf(rg.config.Output.ResourceFilePattern, resourceName)
	outputPath := filepath.Join(outputDir, filename)

	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	if err := os.WriteFile(outputPath, []byte(code), 0644); err != nil {
		return fmt.Errorf("failed to write resource file: %w", err)
	}

	return nil
}

// prepareResourceData prepares the data structure for template execution
func (rg *ResourceGenerator) prepareResourceData(resourceName string, endpoints []parser.Endpoint) *ResourceData {
	data := &ResourceData{
		ResourceName: resourceName,
		StructName:   rg.templateEngine.toPascalCase(resourceName) + "Resource",
		TypeName:     rg.config.Global.ProviderName + "_" + resourceName,
		PackageName:  rg.config.Global.PackageName,
		ProviderName: rg.config.Global.ProviderName,
		Endpoints:    endpoints,
		Schema:       rg.generateSchema(endpoints),
		Config:       rg.config,
	}

	// Identify CRUD endpoints
	for _, endpoint := range endpoints {
		switch endpoint.CRUDType {
		case "create":
			data.CreateEndpoint = &endpoint
		case "read":
			data.ReadEndpoint = &endpoint
		case "update":
			data.UpdateEndpoint = &endpoint
		case "delete":
			data.DeleteEndpoint = &endpoint
		}
	}

	return data
}

// generateSchema generates the Terraform schema from endpoints
func (rg *ResourceGenerator) generateSchema(endpoints []parser.Endpoint) *ResourceSchema {
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

	// Enhanced schema extraction: Analyze POST and PUT operations first
	rg.extractSchemaFromCRUDOperations(endpoints, schema, processedFields)

	// Fallback: Extract from request bodies (for input fields)
	hasCreateEndpoint := false
	for _, endpoint := range endpoints {
		if endpoint.CRUDType == "create" && endpoint.Operation.RequestBody != nil {
			hasCreateEndpoint = true
			rg.extractSchemaFromRequestBody(endpoint.Operation.RequestBody, schema, processedFields, true)
		}
	}

	// Extract from successful responses (for computed fields and missing input fields)
	for _, endpoint := range endpoints {
		for statusCode, response := range endpoint.Operation.Responses {
			isSuccessResponse := statusCode == "200" || statusCode == "201" || statusCode == "202"
			if isSuccessResponse {
				// If we don't have a create endpoint, treat response fields as potential inputs
				isInput := !hasCreateEndpoint && endpoint.CRUDType == "create"
				rg.extractSchemaFromResponse(&response, schema, processedFields, isInput)
			}
		}
	}

	return schema
}

// extractSchemaFromCRUDOperations analyzes POST and PUT operations to extract complete resource schemas
func (rg *ResourceGenerator) extractSchemaFromCRUDOperations(endpoints []parser.Endpoint, schema *ResourceSchema, processedFields map[string]bool) {
	// Extract schema from CRUD operations

	// First, analyze POST operations (create) for required fields and input schema
	for _, endpoint := range endpoints {
		// Check endpoint for request body
		if endpoint.Method == "POST" && endpoint.Operation.RequestBody != nil {
			// Process POST endpoint request body
			rg.extractSchemaFromRequestBody(endpoint.Operation.RequestBody, schema, processedFields, true)
		}
	}

	// Then, analyze PUT operations (update) for additional optional fields
	for _, endpoint := range endpoints {
		if endpoint.Method == "PUT" && endpoint.Operation.RequestBody != nil {
			// Process PUT endpoint request body
			rg.extractSchemaFromRequestBody(endpoint.Operation.RequestBody, schema, processedFields, true)
		}
	}

	// Finally, analyze PATCH operations for partial updates
	for _, endpoint := range endpoints {
		if endpoint.Method == "PATCH" && endpoint.Operation.RequestBody != nil {
			rg.extractSchemaFromRequestBody(endpoint.Operation.RequestBody, schema, processedFields, true)
		}
	}
}

// extractSchemaFromRequestBody extracts schema attributes from request body
func (rg *ResourceGenerator) extractSchemaFromRequestBody(requestBody *parser.RequestBody, schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
	for _, mediaType := range requestBody.Content {
		if mediaType.Schema != nil {
			rg.extractSchemaFromOpenAPISchema(mediaType.Schema, schema, processedFields, isInput)
		}
	}
}

// extractSchemaFromResponse extracts schema attributes from response
func (rg *ResourceGenerator) extractSchemaFromResponse(response *parser.Response, schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
	for _, mediaType := range response.Content {
		if mediaType.Schema != nil {
			rg.extractSchemaFromOpenAPISchema(mediaType.Schema, schema, processedFields, isInput)
		}
	}
}

// extractSchemaFromOpenAPISchema extracts attributes from an OpenAPI schema
func (rg *ResourceGenerator) extractSchemaFromOpenAPISchema(apiSchema *parser.Schema, schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
	rg.extractSchemaFromOpenAPISchemaWithPrefix(apiSchema, schema, processedFields, isInput, "")
}

// extractSchemaFromOpenAPISchemaWithPrefix extracts attributes from an OpenAPI schema with field prefix support
func (rg *ResourceGenerator) extractSchemaFromOpenAPISchemaWithPrefix(apiSchema *parser.Schema, schema *ResourceSchema, processedFields map[string]bool, isInput bool, prefix string) {
	if apiSchema == nil {
		return
	}

	// Process schema properties

	// Handle schema references - this is the key issue we need to fix
	if apiSchema.Ref != "" {
		// The schema reference should have been resolved by the normalizer
		// If we still have a reference here, it means normalization failed
		rg.handleSchemaReference(apiSchema.Ref, schema, processedFields, isInput)
		return
	}

	// Handle direct properties
	if apiSchema.Properties != nil {
		for propName, propSchema := range apiSchema.Properties {
			// Create flattened field name
			fieldName := propName
			if prefix != "" {
				fieldName = prefix + "_" + propName
			}

			// Skip if already processed
			if processedFields[fieldName] {
				continue
			}

			// Handle nested objects by flattening them
			if propSchema.Type == "object" && propSchema.Properties != nil {
				// Recursively process nested object properties
				rg.extractSchemaFromOpenAPISchemaWithPrefix(propSchema, schema, processedFields, isInput, fieldName)
			} else {
				// Create attribute for primitive types
				attr := rg.createSchemaAttributeWithPrefix(fieldName, propName, propSchema, apiSchema.Required, isInput)
				if attr != nil {
					schema.Attributes = append(schema.Attributes, *attr)
					processedFields[fieldName] = true
				} else {
				}
			}
		}
	} else {
		// If we have no properties and no reference, this might be a resolved reference
		// that resulted in an empty schema. Let's try to infer from the context.
		if apiSchema.Type == "" && apiSchema.Ref == "" {
			// Try to add common attributes based on context
			rg.addCommonSchemaAttributes("", schema, processedFields, isInput)
		}
	}

	// Handle nested schemas in data wrappers (common in API responses)
	if apiSchema.Type == "object" && len(apiSchema.Properties) > 0 {
		// Look for common data wrapper patterns
		if dataSchema, exists := apiSchema.Properties["data"]; exists {
			rg.extractSchemaFromOpenAPISchemaWithPrefix(dataSchema, schema, processedFields, isInput, prefix)
		}
	}

	// Handle array items
	if apiSchema.Type == "array" && apiSchema.Items != nil {
		rg.extractSchemaFromOpenAPISchemaWithPrefix(apiSchema.Items, schema, processedFields, isInput, prefix)
	}

	// Handle allOf, oneOf, anyOf
	for _, subSchema := range apiSchema.AllOf {
		rg.extractSchemaFromOpenAPISchemaWithPrefix(subSchema, schema, processedFields, isInput, prefix)
	}
	for _, subSchema := range apiSchema.OneOf {
		rg.extractSchemaFromOpenAPISchemaWithPrefix(subSchema, schema, processedFields, isInput, prefix)
	}
	for _, subSchema := range apiSchema.AnyOf {
		rg.extractSchemaFromOpenAPISchemaWithPrefix(subSchema, schema, processedFields, isInput, prefix)
	}
}

// handleSchemaReference handles unresolved schema references by creating attributes based on common patterns
func (rg *ResourceGenerator) handleSchemaReference(ref string, schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
	// Extract schema name from reference
	schemaName := strings.TrimPrefix(ref, "#/components/schemas/")

	// Try to infer common schema patterns dynamically
	rg.addCommonSchemaAttributes(schemaName, schema, processedFields, isInput)
}

// addCommonSchemaAttributes adds attributes based on common schema patterns
func (rg *ResourceGenerator) addCommonSchemaAttributes(schemaName string, schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
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
		rg.addDestinationListSpecificAttributes(schema, processedFields, isInput)
		return
	}

	// For tunnel/VPN schemas
	if strings.Contains(schemaLower, "tunnel") || strings.Contains(schemaLower, "vpn") || strings.Contains(schemaLower, "ipsec") {
		rg.addTunnelSpecificAttributes(schema, processedFields, isInput)
	}

	// For rule/policy schemas
	if strings.Contains(schemaLower, "rule") || strings.Contains(schemaLower, "policy") {
		rg.addRuleSpecificAttributes(schema, processedFields, isInput)
	}

	// For SAML schemas
	if strings.Contains(schemaLower, "saml") {
		rg.addSAMLSpecificAttributes(schema, processedFields, isInput)
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
				// Always computed fields (timestamps, IDs, etc.)
				schemaAttr.Computed = true
				schemaAttr.Required = false
				schemaAttr.Optional = false
			} else if isInput {
				// Input context: required or optional fields
				schemaAttr.Required = attr.required
				schemaAttr.Optional = !attr.required
				schemaAttr.Computed = false
			} else {
				// Response context: all fields are computed
				schemaAttr.Computed = true
				schemaAttr.Required = false
				schemaAttr.Optional = false
			}

			schema.Attributes = append(schema.Attributes, *schemaAttr)
			processedFields[attr.name] = true
		}
	}
}

// addDestinationListSpecificAttributes adds destination list specific attributes
func (rg *ResourceGenerator) addDestinationListSpecificAttributes(schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
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

	rg.addAttributesToSchema(attributes, schema, processedFields, isInput)
}

// addTunnelSpecificAttributes adds tunnel/VPN specific attributes
func (rg *ResourceGenerator) addTunnelSpecificAttributes(schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
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

	rg.addAttributesToSchema(attributes, schema, processedFields, isInput)
}

// addRuleSpecificAttributes adds rule/policy specific attributes
func (rg *ResourceGenerator) addRuleSpecificAttributes(schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
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

	rg.addAttributesToSchema(attributes, schema, processedFields, isInput)
}

// addSAMLSpecificAttributes adds SAML specific attributes
func (rg *ResourceGenerator) addSAMLSpecificAttributes(schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
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

	rg.addAttributesToSchema(attributes, schema, processedFields, isInput)
}

// addAttributesToSchema is a helper method to add attributes to schema
func (rg *ResourceGenerator) addAttributesToSchema(attributes []struct {
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
				// Always computed fields (timestamps, IDs, etc.)
				schemaAttr.Computed = true
				schemaAttr.Required = false
				schemaAttr.Optional = false
			} else if isInput {
				// Input context: required or optional fields
				schemaAttr.Required = attr.required
				schemaAttr.Optional = !attr.required
				schemaAttr.Computed = false
			} else {
				// Response context: all fields are computed
				schemaAttr.Computed = true
				schemaAttr.Required = false
				schemaAttr.Optional = false
			}

			schema.Attributes = append(schema.Attributes, *schemaAttr)
			processedFields[attr.name] = true
		}
	}
}

// createSchemaAttribute creates a schema attribute from an OpenAPI property
func (rg *ResourceGenerator) createSchemaAttribute(propName string, propSchema *parser.Schema, requiredFields []string, isInput bool) *SchemaAttribute {
	if propSchema == nil {
		return nil
	}

	// Skip certain system fields that shouldn't be in Terraform schema
	skipFields := map[string]bool{
		"status":     true, // API status responses
		"meta":       true, // Pagination metadata (handle separately if needed)
		"txId":       true, // Transaction IDs
		"statusCode": true, // HTTP status codes
		"error":      true, // Error fields
	}

	if skipFields[propName] {
		return nil
	}

	attr := &SchemaAttribute{
		Name:        propName,
		Type:        rg.templateEngine.schemaToTerraformType(propSchema),
		GoType:      rg.templateEngine.schemaToGoType(propSchema),
		Description: rg.cleanDescription(propSchema.Description),
	}

	// Determine if required/optional/computed
	if isInput {
		attr.Required = rg.isRequired(propName, requiredFields)
		attr.Optional = !attr.Required
	} else {
		attr.Computed = true
	}

	// Handle special cases
	switch propName {
	case "id":
		// ID is always computed
		attr.Computed = true
		attr.Required = false
		attr.Optional = false
	case "organizationId", "organization_id":
		// Organization ID is typically computed
		attr.Computed = true
		attr.Required = false
		attr.Optional = false
	case "createdAt", "created_at", "modifiedAt", "modified_at":
		// Timestamps are always computed
		attr.Computed = true
		attr.Required = false
		attr.Optional = false
	}

	return attr
}

// createSchemaAttributeWithPrefix creates a schema attribute with support for flattened field names
func (rg *ResourceGenerator) createSchemaAttributeWithPrefix(fieldName, originalName string, propSchema *parser.Schema, requiredFields []string, isInput bool) *SchemaAttribute {
	if propSchema == nil {
		return nil
	}

	// Use fieldName if provided, otherwise use originalName
	name := fieldName
	if name == "" {
		name = originalName
	}

	// Skip certain system fields that shouldn't be in Terraform schema
	skipFields := map[string]bool{
		"status":     true, // API status responses
		"meta":       true, // Pagination metadata (handle separately if needed)
		"txId":       true, // Transaction IDs
		"statusCode": true, // HTTP status codes
		"error":      true, // Error fields
	}

	if skipFields[originalName] {
		return nil
	}

	attr := &SchemaAttribute{
		Name:        rg.toTerraformName(name),
		Type:        rg.templateEngine.schemaToTerraformType(propSchema),
		GoType:      rg.templateEngine.schemaToGoType(propSchema),
		Description: rg.cleanDescription(propSchema.Description),
	}

	// Determine if required/optional/computed (check against original name in OpenAPI spec)
	if isInput {
		attr.Required = rg.isRequired(originalName, requiredFields)
		attr.Optional = !attr.Required
	} else {
		attr.Computed = true
	}

	// Handle special cases based on original field name
	switch originalName {
	case "id":
		// ID is always computed
		attr.Computed = true
		attr.Required = false
		attr.Optional = false
	case "organizationId", "organization_id":
		// Organization ID is typically computed
		attr.Computed = true
		attr.Required = false
		attr.Optional = false
	case "createdAt", "created_at", "modifiedAt", "modified_at":
		// Timestamps are always computed
		attr.Computed = true
		attr.Required = false
		attr.Optional = false
	}

	return attr
}

// cleanDescription cleans up description text for Terraform schema
func (rg *ResourceGenerator) cleanDescription(desc string) string {
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

	// Escape quotes for Go strings
	desc = strings.ReplaceAll(desc, `"`, `\"`)

	return desc
}

// attributeExists checks if an attribute already exists in the schema
func (rg *ResourceGenerator) attributeExists(schema *ResourceSchema, name string) bool {
	for _, attr := range schema.Attributes {
		if attr.Name == name {
			return true
		}
	}
	return false
}

// isRequired checks if a property is in the required list
func (rg *ResourceGenerator) isRequired(propName string, required []string) bool {
	for _, req := range required {
		if req == propName {
			return true
		}
	}
	return false
}

// toTerraformName converts a field name to Terraform-friendly snake_case format
func (rg *ResourceGenerator) toTerraformName(name string) string {
	// Convert camelCase to snake_case
	result := ""
	for i, r := range name {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result += "_"
		}
		result += strings.ToLower(string(r))
	}
	return result
}

// registerResourceTemplate registers the embedded resource template
func (rg *ResourceGenerator) registerResourceTemplate() error {
	template := `{{if .Config.Output.AddGenerationMarker}}// Code generated by terraform-provider-umbrella generator. DO NOT EDIT.

{{end}}package {{.PackageName}}

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// init registers this resource with the generated resource registry
func init() {
	RegisterGeneratedResource(New{{.StructName}})
}

// {{.StructName}} implements the {{.ResourceName}} resource
type {{.StructName}} struct {
	client *apiClient
}

// {{snakeCase .ResourceName}}Model represents the resource data model
type {{snakeCase .ResourceName}}Model struct {
{{- range .Schema.Attributes}}
	{{pascalCase .Name}} {{.Type}} ` + "`tfsdk:\"{{.Name}}\"`" + `
{{- end}}
}

// New{{.StructName}} creates a new {{.ResourceName}} resource
func New{{.StructName}}() resource.Resource {
	return &{{.StructName}}{}
}

// Metadata returns the resource type name
func (r *{{.StructName}}) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "{{.TypeName}}"
}

// Configure configures the resource with the provider client
func (r *{{.StructName}}) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*apiClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", "Expected *apiClient")
		return
	}

	r.client = client
}

// Schema defines the resource schema
func (r *{{.StructName}}) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "{{.ResourceName}} resource",
		Attributes: map[string]schema.Attribute{
{{- range .Schema.Attributes}}
			"{{.Name}}": schema.{{if eq .Type "types.String"}}String{{else if eq .Type "types.Int64"}}Int64{{else if eq .Type "types.Bool"}}Bool{{else if eq .Type "types.Set"}}Set{{else}}String{{end}}Attribute{
				{{- if .Required}} Required: true,{{end}}
				{{- if .Optional}} Optional: true,{{end}}
				{{- if .Computed}} Computed: true,{{end}}
				{{- if .Description}} Description: "{{cleanDesc .Description}}",{{end}}
				{{- if eq .Type "types.Set"}} ElementType: types.StringType,{{end}}
			},
{{- end}}
		},
	}
}

// Create creates a new {{.ResourceName}}
func (r *{{.StructName}}) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan {{snakeCase .ResourceName}}Model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	{{- if .CreateEndpoint}}
	// Create request body from plan
	requestBody := make(map[string]interface{})
	{{- range .Schema.Attributes}}
	{{- if not .Computed}}
	if !plan.{{pascalCase .Name}}.IsNull() {
		{{- if eq .Type "types.String"}}
		requestBody["{{.Name}}"] = plan.{{pascalCase .Name}}.ValueString()
		{{- else if eq .Type "types.Bool"}}
		requestBody["{{.Name}}"] = plan.{{pascalCase .Name}}.ValueBool()
		{{- else if eq .Type "types.Int64"}}
		requestBody["{{.Name}}"] = plan.{{pascalCase .Name}}.ValueInt64()
		{{- else}}
		requestBody["{{.Name}}"] = plan.{{pascalCase .Name}}.ValueString()
		{{- end}}
	}
	{{- end}}
	{{- end}}

	// Make API call
	result, err := r.client.CreateResource(ctx, "{{.CreateEndpoint.Path}}", requestBody)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create {{.ResourceName}}, got error: %s", err))
		return
	}

	// Update state with response data
	if result.Data != nil {
		if dataMap, ok := result.Data.(map[string]interface{}); ok {
			{{- range .Schema.Attributes}}
			if val, exists := dataMap["{{.Name}}"]; exists && val != nil {
				{{- if eq .Type "types.String"}}
				if strVal, ok := val.(string); ok {
					plan.{{pascalCase .Name}} = types.StringValue(strVal)
				}
				{{- else if eq .Type "types.Bool"}}
				if boolVal, ok := val.(bool); ok {
					plan.{{pascalCase .Name}} = types.BoolValue(boolVal)
				}
				{{- else if eq .Type "types.Int64"}}
				if floatVal, ok := val.(float64); ok {
					plan.{{pascalCase .Name}} = types.Int64Value(int64(floatVal))
				}
				{{- end}}
			}
			{{- end}}
		}
	}
	{{- else}}
	// No specific create endpoint found
	resp.Diagnostics.AddError("Configuration Error", "No create endpoint configured for {{.ResourceName}}")
	return
	{{- end}}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Read reads the {{.ResourceName}}
func (r *{{.StructName}}) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state {{snakeCase .ResourceName}}Model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	{{- if .ReadEndpoint}}
	// Build path with ID
	path := fmt.Sprintf("{{.ReadEndpoint.Path}}", state.Id.ValueString())

	// Make API call
	result, err := r.client.GetResource(ctx, path)
	if err != nil {
		if err.Error() == "resource not found" {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read {{.ResourceName}}, got error: %s", err))
		return
	}

	// Update state with response data
	if result.Data != nil {
		if dataMap, ok := result.Data.(map[string]interface{}); ok {
			{{- range .Schema.Attributes}}
			if val, exists := dataMap["{{.Name}}"]; exists && val != nil {
				{{- if eq .Type "types.String"}}
				if strVal, ok := val.(string); ok {
					state.{{pascalCase .Name}} = types.StringValue(strVal)
				}
				{{- else if eq .Type "types.Bool"}}
				if boolVal, ok := val.(bool); ok {
					state.{{pascalCase .Name}} = types.BoolValue(boolVal)
				}
				{{- else if eq .Type "types.Int64"}}
				if floatVal, ok := val.(float64); ok {
					state.{{pascalCase .Name}} = types.Int64Value(int64(floatVal))
				}
				{{- end}}
			}
			{{- end}}
		}
	}
	{{- else}}
	// No specific read endpoint found - return current state
	// This is a no-op read that just returns the current state
	{{- end}}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update updates the {{.ResourceName}}
func (r *{{.StructName}}) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan {{snakeCase .ResourceName}}Model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	{{- if .UpdateEndpoint}}
	// Create request body from plan
	requestBody := make(map[string]interface{})
	{{- range .Schema.Attributes}}
	{{- if and (not .Computed) (ne .Name "id")}}
	if !plan.{{pascalCase .Name}}.IsNull() {
		{{- if eq .Type "types.String"}}
		requestBody["{{.Name}}"] = plan.{{pascalCase .Name}}.ValueString()
		{{- else if eq .Type "types.Bool"}}
		requestBody["{{.Name}}"] = plan.{{pascalCase .Name}}.ValueBool()
		{{- else if eq .Type "types.Int64"}}
		requestBody["{{.Name}}"] = plan.{{pascalCase .Name}}.ValueInt64()
		{{- else}}
		requestBody["{{.Name}}"] = plan.{{pascalCase .Name}}.ValueString()
		{{- end}}
	}
	{{- end}}
	{{- end}}

	// Build path with ID
	path := fmt.Sprintf("{{.UpdateEndpoint.Path}}", plan.Id.ValueString())

	// Make API call
	result, err := r.client.UpdateResource(ctx, path, requestBody)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update {{.ResourceName}}, got error: %s", err))
		return
	}

	// Update state with response data
	if result.Data != nil {
		if dataMap, ok := result.Data.(map[string]interface{}); ok {
			{{- range .Schema.Attributes}}
			if val, exists := dataMap["{{.Name}}"]; exists && val != nil {
				{{- if eq .Type "types.String"}}
				if strVal, ok := val.(string); ok {
					plan.{{pascalCase .Name}} = types.StringValue(strVal)
				}
				{{- else if eq .Type "types.Bool"}}
				if boolVal, ok := val.(bool); ok {
					plan.{{pascalCase .Name}} = types.BoolValue(boolVal)
				}
				{{- else if eq .Type "types.Int64"}}
				if floatVal, ok := val.(float64); ok {
					plan.{{pascalCase .Name}} = types.Int64Value(int64(floatVal))
				}
				{{- end}}
			}
			{{- end}}
		}
	}
	{{- else}}
	// No specific update endpoint found
	resp.Diagnostics.AddError("Configuration Error", "No update endpoint configured for {{.ResourceName}}")
	return
	{{- end}}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

{{- if .DeleteEndpoint}}
// Delete deletes the {{.ResourceName}}
func (r *{{.StructName}}) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state {{snakeCase .ResourceName}}Model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Build path with ID
	path := fmt.Sprintf("{{.DeleteEndpoint.Path}}", state.Id.ValueString())

	// Make API call
	err := r.client.DeleteResource(ctx, path)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete {{.ResourceName}}, got error: %s", err))
		return
	}
}
{{- else}}
// Delete deletes the {{.ResourceName}}
func (r *{{.StructName}}) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state {{snakeCase .ResourceName}}Model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// No specific delete endpoint found
	resp.Diagnostics.AddError("Configuration Error", "No delete endpoint configured for {{.ResourceName}}")
}
{{- end}}
`

	return rg.templateEngine.RegisterTemplate("resource.go", template)
}
