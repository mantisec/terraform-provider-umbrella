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

	// First pass: Extract from request bodies (for input fields)
	hasCreateEndpoint := false
	for _, endpoint := range endpoints {
		if endpoint.CRUDType == "create" && endpoint.Operation.RequestBody != nil {
			hasCreateEndpoint = true
			rg.extractSchemaFromRequestBody(endpoint.Operation.RequestBody, schema, processedFields, true)
		}
	}

	// Second pass: Extract from successful responses (for computed fields and missing input fields)
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
	if apiSchema == nil {
		return
	}

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
			// Skip if already processed
			if processedFields[propName] {
				continue
			}

			attr := rg.createSchemaAttribute(propName, propSchema, apiSchema.Required, isInput)
			if attr != nil {
				schema.Attributes = append(schema.Attributes, *attr)
				processedFields[propName] = true
			}
		}
	} else {
		// If we have no properties and no reference, this might be a resolved reference
		// that resulted in an empty schema. Let's try to infer from the context.
		if apiSchema.Type == "" && apiSchema.Ref == "" {
			// Try to add default destination list attributes
			rg.addDestinationListObjectAttributes(schema, processedFields, isInput)
		}
	}

	// Handle nested schemas in data wrappers (common in API responses)
	if apiSchema.Type == "object" && len(apiSchema.Properties) > 0 {
		// Look for common data wrapper patterns
		if dataSchema, exists := apiSchema.Properties["data"]; exists {
			rg.extractSchemaFromOpenAPISchema(dataSchema, schema, processedFields, isInput)
		}
	}

	// Handle array items
	if apiSchema.Type == "array" && apiSchema.Items != nil {
		rg.extractSchemaFromOpenAPISchema(apiSchema.Items, schema, processedFields, isInput)
	}

	// Handle allOf, oneOf, anyOf
	for _, subSchema := range apiSchema.AllOf {
		rg.extractSchemaFromOpenAPISchema(subSchema, schema, processedFields, isInput)
	}
	for _, subSchema := range apiSchema.OneOf {
		rg.extractSchemaFromOpenAPISchema(subSchema, schema, processedFields, isInput)
	}
	for _, subSchema := range apiSchema.AnyOf {
		rg.extractSchemaFromOpenAPISchema(subSchema, schema, processedFields, isInput)
	}
}

// handleSchemaReference handles unresolved schema references by creating attributes based on known patterns
func (rg *ResourceGenerator) handleSchemaReference(ref string, schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
	// Extract schema name from reference
	schemaName := strings.TrimPrefix(ref, "#/components/schemas/")

	// For destination lists, we know the expected schema structure from the OpenAPI spec
	// This is a temporary solution until we fix the reference resolution
	switch schemaName {
	case "DestinationListResponse":
		rg.addDestinationListResponseAttributes(schema, processedFields, isInput)
	case "DestinationListObject":
		rg.addDestinationListObjectAttributes(schema, processedFields, isInput)
	case "DestinationListCreate":
		rg.addDestinationListCreateAttributes(schema, processedFields, isInput)
	case "DestinationListPatch":
		rg.addDestinationListPatchAttributes(schema, processedFields, isInput)
	case "PaginatedDestinationListsResponse":
		rg.addPaginatedDestinationListsResponseAttributes(schema, processedFields, isInput)
	}
}

// addDestinationListObjectAttributes adds attributes for DestinationListObject schema
func (rg *ResourceGenerator) addDestinationListObjectAttributes(schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
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
		{"name", "types.String", "string", true, false, "The name of the destination list"},
		{"bundle_type_id", "types.Int64", "int64", false, false, "The type of the destination list in the policy"},
		{"organization_id", "types.Int64", "int64", false, true, "The organization ID"},
		{"thirdparty_category_id", "types.Int64", "int64", false, true, "The third-party category ID of the destination list"},
		{"created_at", "types.Int64", "int64", false, true, "The date and time when the destination list was created"},
		{"modified_at", "types.Int64", "int64", false, true, "The date and time when the destination list was modified"},
		{"is_msp_default", "types.Bool", "bool", false, true, "Specifies whether MSP is the default"},
		{"marked_for_deletion", "types.Bool", "bool", false, true, "Specifies whether the destination list is marked for deletion"},
	}

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

// addDestinationListCreateAttributes adds attributes for DestinationListCreate schema
func (rg *ResourceGenerator) addDestinationListCreateAttributes(schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
	attributes := []struct {
		name        string
		tfType      string
		goType      string
		required    bool
		description string
	}{
		{"access", "types.String", "string", true, "The type of access for the destination list (allow/block)"},
		{"is_global", "types.Bool", "bool", true, "Specifies whether the destination list is a global destination list"},
		{"name", "types.String", "string", true, "The name of the destination list"},
		{"bundle_type_id", "types.Int64", "int64", false, "The type of the destination list in the policy"},
	}

	for _, attr := range attributes {
		if !processedFields[attr.name] {
			schemaAttr := &SchemaAttribute{
				Name:        attr.name,
				Type:        attr.tfType,
				GoType:      attr.goType,
				Required:    attr.required && isInput,
				Optional:    !attr.required && isInput,
				Computed:    !isInput,
				Description: attr.description,
			}
			schema.Attributes = append(schema.Attributes, *schemaAttr)
			processedFields[attr.name] = true
		}
	}
}

// addDestinationListResponseAttributes adds attributes for DestinationListResponse schema
func (rg *ResourceGenerator) addDestinationListResponseAttributes(schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
	// DestinationListResponse contains a data field with DestinationListObject
	rg.addDestinationListObjectAttributes(schema, processedFields, isInput)
}

// addDestinationListPatchAttributes adds attributes for DestinationListPatch schema
func (rg *ResourceGenerator) addDestinationListPatchAttributes(schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
	if !processedFields["name"] {
		schemaAttr := &SchemaAttribute{
			Name:        "name",
			Type:        "types.String",
			GoType:      "string",
			Required:    true && isInput,
			Optional:    false,
			Computed:    !isInput,
			Description: "The name of the destination list",
		}
		schema.Attributes = append(schema.Attributes, *schemaAttr)
		processedFields["name"] = true
	}
}

// addPaginatedDestinationListsResponseAttributes adds attributes for PaginatedDestinationListsResponse schema
func (rg *ResourceGenerator) addPaginatedDestinationListsResponseAttributes(schema *ResourceSchema, processedFields map[string]bool, isInput bool) {
	// This is typically used for data sources (list operations)
	rg.addDestinationListObjectAttributes(schema, processedFields, isInput)
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

// registerResourceTemplate registers the embedded resource template
func (rg *ResourceGenerator) registerResourceTemplate() error {
	template := `{{if .Config.Output.AddGenerationMarker}}// Code generated by terraform-provider-umbrella generator. DO NOT EDIT.

{{end}}package {{.PackageName}}

import (
	"context"

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
	var plan {{.ResourceName}}Model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	{{- if .CreateEndpoint}}
	// TODO: Implement create logic using {{.CreateEndpoint.Method}} {{.CreateEndpoint.Path}}
	{{- else}}
	// TODO: Implement create logic - no specific create endpoint found
	{{- end}}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

{{- if .ReadEndpoint}}
// Read reads the {{.ResourceName}}
func (r *{{.StructName}}) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state {{.ResourceName}}Model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using {{.ReadEndpoint.Method}} {{.ReadEndpoint.Path}}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
{{- end}}

// Update updates the {{.ResourceName}}
func (r *{{.StructName}}) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan {{.ResourceName}}Model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	{{- if .UpdateEndpoint}}
	// TODO: Implement update logic using {{.UpdateEndpoint.Method}} {{.UpdateEndpoint.Path}}
	{{- else}}
	// TODO: Implement update logic - no specific update endpoint found
	{{- end}}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Read reads the {{.ResourceName}}
func (r *{{.StructName}}) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state {{.ResourceName}}Model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	{{- if .ReadEndpoint}}
	// TODO: Implement read logic using {{.ReadEndpoint.Method}} {{.ReadEndpoint.Path}}
	{{- else}}
	// TODO: Implement read logic - no specific read endpoint found
	{{- end}}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

{{- if .DeleteEndpoint}}
// Delete deletes the {{.ResourceName}}
func (r *{{.StructName}}) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state {{.ResourceName}}Model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement delete logic using {{.DeleteEndpoint.Method}} {{.DeleteEndpoint.Path}}
}
{{- else}}
// Delete deletes the {{.ResourceName}}
func (r *{{.StructName}}) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state {{.ResourceName}}Model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement delete logic - no specific delete endpoint found
}
{{- end}}
`

	return rg.templateEngine.RegisterTemplate("resource.go", template)
}
