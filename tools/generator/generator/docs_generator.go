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
	attributes := []DocumentationAttribute{
		{
			Name:        "id",
			Type:        "String",
			Computed:    true,
			Description: "The unique identifier for this resource",
			Example:     "\"12345\"",
		},
	}

	// Extract attributes from endpoint schemas
	for _, endpoint := range endpoints {
		if endpoint.Operation.RequestBody != nil {
			attributes = append(attributes, dg.extractAttributesFromRequestBody(endpoint.Operation.RequestBody)...)
		}

		for _, response := range endpoint.Operation.Responses {
			attributes = append(attributes, dg.extractAttributesFromResponse(&response)...)
		}
	}

	// Remove duplicates
	return dg.removeDuplicateAttributes(attributes)
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
	return fmt.Sprintf(`resource "%s" "example" {
  # Add required attributes here
  name = "example-%s"
}`, typeName, resourceName)
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
page_title: "{{.TypeName}} {{if eq .ResourceType "resource"}}Resource{{else}}Data Source{{end}} - {{.ResourceName}}"
description: |-
  {{.Description}}
---

# {{.TypeName}} ({{if eq .ResourceType "resource"}}Resource{{else}}Data Source{{end}})

{{.Description}}

## Example Usage

{{range .Examples}}
### {{.Title}}

{{.Description}}

` + "```hcl" + `
{{.Code}}
` + "```" + `

{{end}}

## Schema

### {{if eq .ResourceType "resource"}}Required{{else}}Optional{{end}}

{{range .Attributes}}{{if .Required}}- ` + "`{{.Name}}`" + ` ({{.Type}}) {{.Description}}
{{end}}{{end}}

### Optional

{{range .Attributes}}{{if .Optional}}- ` + "`{{.Name}}`" + ` ({{.Type}}) {{.Description}}
{{end}}{{end}}

### Read-Only

{{range .Attributes}}{{if .Computed}}- ` + "`{{.Name}}`" + ` ({{.Type}}) {{.Description}}
{{end}}{{end}}

{{if eq .ResourceType "resource"}}
## Import

Import is supported using the following syntax:

` + "```shell" + `
{{.ImportExample}}
` + "```" + `
{{end}}
`

	return dg.templateEngine.RegisterTemplate("docs.md", template)
}
