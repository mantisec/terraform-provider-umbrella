package generator

import (
	"fmt"
	"os"
	"path/filepath"

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

	// Extract schema from request/response bodies
	for _, endpoint := range endpoints {
		if endpoint.Operation.RequestBody != nil {
			rg.extractSchemaFromRequestBody(endpoint.Operation.RequestBody, schema)
		}

		for _, response := range endpoint.Operation.Responses {
			rg.extractSchemaFromResponse(&response, schema)
		}
	}

	return schema
}

// extractSchemaFromRequestBody extracts schema attributes from request body
func (rg *ResourceGenerator) extractSchemaFromRequestBody(requestBody *parser.RequestBody, schema *ResourceSchema) {
	for _, mediaType := range requestBody.Content {
		if mediaType.Schema != nil {
			rg.extractSchemaFromOpenAPISchema(mediaType.Schema, schema, true)
		}
	}
}

// extractSchemaFromResponse extracts schema attributes from response
func (rg *ResourceGenerator) extractSchemaFromResponse(response *parser.Response, schema *ResourceSchema) {
	for _, mediaType := range response.Content {
		if mediaType.Schema != nil {
			rg.extractSchemaFromOpenAPISchema(mediaType.Schema, schema, false)
		}
	}
}

// extractSchemaFromOpenAPISchema extracts attributes from an OpenAPI schema
func (rg *ResourceGenerator) extractSchemaFromOpenAPISchema(apiSchema *parser.Schema, schema *ResourceSchema, isInput bool) {
	if apiSchema.Properties == nil {
		return
	}

	for propName, propSchema := range apiSchema.Properties {
		// Skip if attribute already exists
		if rg.attributeExists(schema, propName) {
			continue
		}

		attr := SchemaAttribute{
			Name:        propName,
			Type:        rg.templateEngine.schemaToTerraformType(propSchema),
			GoType:      rg.templateEngine.schemaToGoType(propSchema),
			Description: propSchema.Description,
		}

		// Determine if required/optional/computed
		if isInput {
			attr.Required = rg.isRequired(propName, apiSchema.Required)
			attr.Optional = !attr.Required
		} else {
			attr.Computed = true
		}

		schema.Attributes = append(schema.Attributes, attr)
	}
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
	template := `package {{.PackageName}}

import (
	"context"


	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

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

{{- if .CreateEndpoint}}
// Create creates a new {{.ResourceName}}
func (r *{{.StructName}}) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan {{.ResourceName}}Model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement create logic using {{.CreateEndpoint.Method}} {{.CreateEndpoint.Path}}
	
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}
{{- end}}

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

{{- if .UpdateEndpoint}}
// Update updates the {{.ResourceName}}
func (r *{{.StructName}}) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan {{.ResourceName}}Model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement update logic using {{.UpdateEndpoint.Method}} {{.UpdateEndpoint.Path}}
	
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}
{{- end}}

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
{{- end}}
`

	return rg.templateEngine.RegisterTemplate("resource.go", template)
}
