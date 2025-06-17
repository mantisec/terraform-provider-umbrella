package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mantisec/terraform-provider-umbrella/tools/generator/config"
	"github.com/mantisec/terraform-provider-umbrella/tools/generator/parser"
)

// DataSourceGenerator handles generation of Terraform data sources
type DataSourceGenerator struct {
	config         *config.Config
	templateEngine *TemplateEngine
}

// NewDataSourceGenerator creates a new data source generator
func NewDataSourceGenerator(cfg *config.Config, templateEngine *TemplateEngine) *DataSourceGenerator {
	return &DataSourceGenerator{
		config:         cfg,
		templateEngine: templateEngine,
	}
}

// DataSourceData contains data for data source template generation
type DataSourceData struct {
	ResourceName string
	StructName   string
	TypeName     string
	PackageName  string
	ProviderName string
	Endpoints    []parser.Endpoint
	ReadEndpoint *parser.Endpoint
	Schema       *ResourceSchema
	Config       *config.Config
}

// GenerateDataSource generates a Terraform data source from endpoints
func (dsg *DataSourceGenerator) GenerateDataSource(resourceName string, endpoints []parser.Endpoint, outputDir string) error {
	// Prepare data source data
	data := dsg.prepareDataSourceData(resourceName, endpoints)

	// Register the embedded template
	if err := dsg.registerDataSourceTemplate(); err != nil {
		return fmt.Errorf("failed to register data source template: %w", err)
	}

	// Generate the data source code
	code, err := dsg.templateEngine.ExecuteTemplate("data_source.go", data)
	if err != nil {
		return fmt.Errorf("failed to execute data source template: %w", err)
	}

	// Write to file
	filename := fmt.Sprintf(dsg.config.Output.DataSourceFilePattern, resourceName)
	outputPath := filepath.Join(outputDir, filename)

	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	if err := os.WriteFile(outputPath, []byte(code), 0644); err != nil {
		return fmt.Errorf("failed to write data source file: %w", err)
	}

	return nil
}

// prepareDataSourceData prepares the data structure for template execution
func (dsg *DataSourceGenerator) prepareDataSourceData(resourceName string, endpoints []parser.Endpoint) *DataSourceData {
	data := &DataSourceData{
		ResourceName: resourceName,
		StructName:   dsg.templateEngine.toPascalCase(resourceName) + "DataSource",
		TypeName:     dsg.config.Global.ProviderName + "_" + resourceName,
		PackageName:  dsg.config.Global.PackageName,
		ProviderName: dsg.config.Global.ProviderName,
		Endpoints:    endpoints,
		Schema:       dsg.generateSchema(endpoints),
		Config:       dsg.config,
	}

	// Find the read endpoint
	for _, endpoint := range endpoints {
		if endpoint.CRUDType == "read" || endpoint.CRUDType == "list" {
			data.ReadEndpoint = &endpoint
			break
		}
	}

	return data
}

// generateSchema generates the Terraform schema from endpoints
func (dsg *DataSourceGenerator) generateSchema(endpoints []parser.Endpoint) *ResourceSchema {
	schema := &ResourceSchema{
		Attributes: []SchemaAttribute{
			{
				Name:        "id",
				Type:        "types.String",
				Computed:    true,
				Description: "Data source identifier",
				GoType:      "string",
			},
		},
	}

	// Extract schema from response bodies (data sources are read-only)
	for _, endpoint := range endpoints {
		for _, response := range endpoint.Operation.Responses {
			dsg.extractSchemaFromResponse(&response, schema)
		}
	}

	return schema
}

// extractSchemaFromResponse extracts schema attributes from response
func (dsg *DataSourceGenerator) extractSchemaFromResponse(response *parser.Response, schema *ResourceSchema) {
	for _, mediaType := range response.Content {
		if mediaType.Schema != nil {
			dsg.extractSchemaFromOpenAPISchema(mediaType.Schema, schema)
		}
	}
}

// extractSchemaFromOpenAPISchema extracts attributes from an OpenAPI schema
func (dsg *DataSourceGenerator) extractSchemaFromOpenAPISchema(apiSchema *parser.Schema, schema *ResourceSchema) {
	if apiSchema.Properties == nil {
		return
	}

	for propName, propSchema := range apiSchema.Properties {
		// Skip if attribute already exists
		if dsg.attributeExists(schema, propName) {
			continue
		}

		attr := SchemaAttribute{
			Name:        propName,
			Type:        dsg.templateEngine.schemaToTerraformType(propSchema),
			GoType:      dsg.templateEngine.schemaToGoType(propSchema),
			Description: propSchema.Description,
			Computed:    true, // Data source attributes are always computed
		}

		schema.Attributes = append(schema.Attributes, attr)
	}
}

// attributeExists checks if an attribute already exists in the schema
func (dsg *DataSourceGenerator) attributeExists(schema *ResourceSchema, name string) bool {
	for _, attr := range schema.Attributes {
		if attr.Name == name {
			return true
		}
	}
	return false
}

// registerDataSourceTemplate registers the embedded data source template
func (dsg *DataSourceGenerator) registerDataSourceTemplate() error {
	template := `{{if .Config.Output.AddGenerationMarker}}// Code generated by terraform-provider-umbrella generator. DO NOT EDIT.

{{end}}package {{.PackageName}}

import (
	"context"


	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// {{.StructName}} implements the {{.ResourceName}} data source
type {{.StructName}} struct {
	client *apiClient
}

// {{snakeCase .ResourceName}}DataModel represents the data source data model
type {{snakeCase .ResourceName}}DataModel struct {
{{- range .Schema.Attributes}}
	{{pascalCase .Name}} {{.Type}} ` + "`tfsdk:\"{{.Name}}\"`" + `
{{- end}}
}

// New{{.StructName}} creates a new {{.ResourceName}} data source
func New{{.StructName}}() datasource.DataSource {
	return &{{.StructName}}{}
}

// Metadata returns the data source type name
func (d *{{.StructName}}) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "{{.TypeName}}"
}

// Configure configures the data source with the provider client
func (d *{{.StructName}}) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*apiClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected DataSource Configure Type", "Expected *apiClient")
		return
	}

	d.client = client
}

// Schema defines the data source schema
func (d *{{.StructName}}) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "{{.ResourceName}} data source",
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

{{- if .ReadEndpoint}}
// Read reads the {{.ResourceName}} data
func (d *{{.StructName}}) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config {{snakeCase .ResourceName}}DataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using {{.ReadEndpoint.Method}} {{.ReadEndpoint.Path}}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
{{- end}}
`

	return dsg.templateEngine.RegisterTemplate("data_source.go", template)
}
