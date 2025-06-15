package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"strings"
	"text/template"

	"github.com/mantisec/terraform-provider-umbrella/tools/generator/config"
	"github.com/mantisec/terraform-provider-umbrella/tools/generator/parser"
)

// TemplateEngine handles template processing and code generation
type TemplateEngine struct {
	config    *config.Config
	templates map[string]*template.Template
	funcMap   template.FuncMap
}

// NewTemplateEngine creates a new template engine
func NewTemplateEngine(cfg *config.Config) *TemplateEngine {
	engine := &TemplateEngine{
		config:    cfg,
		templates: make(map[string]*template.Template),
	}

	engine.setupFuncMap()
	engine.loadTemplates()

	return engine
}

// setupFuncMap sets up template helper functions
func (te *TemplateEngine) setupFuncMap() {
	te.funcMap = template.FuncMap{
		"title":      strings.Title,
		"lower":      strings.ToLower,
		"upper":      strings.ToUpper,
		"camelCase":  te.toCamelCase,
		"snakeCase":  te.toSnakeCase,
		"pascalCase": te.toPascalCase,
		"pluralize":  te.pluralize,
		"goType":     te.schemaToGoType,
		"tfType":     te.schemaToTerraformType,
		"join":       strings.Join,
		"contains":   strings.Contains,
		"hasPrefix":  strings.HasPrefix,
		"hasSuffix":  strings.HasSuffix,
		"replace":    strings.ReplaceAll,
		"cleanDesc":  te.cleanDescription,
	}
}

// loadTemplates loads all template files
func (te *TemplateEngine) loadTemplates() {
	// For now, we'll use embedded templates
	// In a full implementation, these would be loaded from files
}

// ExecuteTemplate executes a template with the given data
func (te *TemplateEngine) ExecuteTemplate(templateName string, data interface{}) (string, error) {
	tmpl, exists := te.templates[templateName]
	if !exists {
		return "", fmt.Errorf("template %s not found", templateName)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template %s: %w", templateName, err)
	}

	result := buf.String()

	// Format Go code if requested
	if te.config.Output.FormatCode && strings.HasSuffix(templateName, ".go") {
		formatted, err := format.Source([]byte(result))
		if err != nil {
			// If formatting fails, return unformatted code with a warning
			fmt.Printf("Warning: Failed to format generated code: %v\n", err)
			return result, nil
		}
		result = string(formatted)
	}

	return result, nil
}

// Helper functions for templates

// toCamelCase converts a string to camelCase
func (te *TemplateEngine) toCamelCase(s string) string {
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-' || r == ' '
	})

	if len(parts) == 0 {
		return s
	}

	result := strings.ToLower(parts[0])
	for i := 1; i < len(parts); i++ {
		result += strings.Title(strings.ToLower(parts[i]))
	}

	return result
}

// toPascalCase converts a string to PascalCase
func (te *TemplateEngine) toPascalCase(s string) string {
	// Remove file extensions and invalid characters
	s = strings.ReplaceAll(s, ".json", "")
	s = strings.ReplaceAll(s, ".xml", "")
	s = strings.ReplaceAll(s, ".", "_")

	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-' || r == ' '
	})

	var result string
	for _, part := range parts {
		if part != "" {
			result += strings.Title(strings.ToLower(part))
		}
	}

	// Ensure it starts with a letter
	if result == "" || !((result[0] >= 'A' && result[0] <= 'Z') || (result[0] >= 'a' && result[0] <= 'z')) {
		result = "Resource" + result
	}

	return result
}

// toSnakeCase converts a string to snake_case
func (te *TemplateEngine) toSnakeCase(s string) string {
	// Remove file extensions and invalid characters
	s = strings.ReplaceAll(s, ".json", "")
	s = strings.ReplaceAll(s, ".xml", "")
	s = strings.ReplaceAll(s, ".", "_")

	result := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(s, "-", "_"), " ", "_"))

	// Ensure it's a valid Go identifier
	if result == "" || (result[0] >= '0' && result[0] <= '9') {
		result = "resource_" + result
	}

	return result
}

// pluralize adds 's' to make a word plural (simple implementation)
func (te *TemplateEngine) pluralize(s string) string {
	if strings.HasSuffix(s, "s") || strings.HasSuffix(s, "x") || strings.HasSuffix(s, "z") {
		return s + "es"
	}
	if strings.HasSuffix(s, "y") {
		return s[:len(s)-1] + "ies"
	}
	return s + "s"
}

// schemaToGoType converts an OpenAPI schema to a Go type
func (te *TemplateEngine) schemaToGoType(schema *parser.Schema) string {
	if schema == nil {
		return "interface{}"
	}

	switch schema.Type {
	case "string":
		return "string"
	case "integer":
		if schema.Format == "int64" {
			return "int64"
		}
		return "int"
	case "number":
		if schema.Format == "float" {
			return "float32"
		}
		return "float64"
	case "boolean":
		return "bool"
	case "array":
		if schema.Items != nil {
			itemType := te.schemaToGoType(schema.Items)
			return "[]" + itemType
		}
		return "[]interface{}"
	case "object":
		return "map[string]interface{}"
	default:
		return "interface{}"
	}
}

// schemaToTerraformType converts an OpenAPI schema to a Terraform type
func (te *TemplateEngine) schemaToTerraformType(schema *parser.Schema) string {
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
		if schema.Items != nil {
			itemType := te.schemaToTerraformType(schema.Items)
			if itemType == "types.String" {
				return "types.Set" // or types.List depending on requirements
			}
		}
		return "types.Set"
	case "object":
		return "types.Object"
	default:
		return "types.String"
	}
}

// RegisterTemplate registers a new template
func (te *TemplateEngine) RegisterTemplate(name, content string) error {
	tmpl, err := template.New(name).Funcs(te.funcMap).Parse(content)
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", name, err)
	}

	te.templates[name] = tmpl
	return nil
}

// LoadTemplateFromFile loads a template from a file
func (te *TemplateEngine) LoadTemplateFromFile(name, path string) error {
	tmpl, err := template.New(name).Funcs(te.funcMap).ParseFiles(path)
	if err != nil {
		return fmt.Errorf("failed to load template from %s: %w", path, err)
	}

	te.templates[name] = tmpl
	return nil
}

// cleanDescription cleans a description string for use in Go code
func (te *TemplateEngine) cleanDescription(desc string) string {
	if desc == "" {
		return ""
	}

	// Replace newlines and multiple spaces with single spaces
	desc = strings.ReplaceAll(desc, "\n", " ")
	desc = strings.ReplaceAll(desc, "\r", " ")
	desc = strings.ReplaceAll(desc, "\t", " ")

	// Replace multiple spaces with single space
	for strings.Contains(desc, "  ") {
		desc = strings.ReplaceAll(desc, "  ", " ")
	}

	// Escape quotes
	desc = strings.ReplaceAll(desc, `"`, `\"`)

	// Remove markdown table syntax and clean up
	desc = strings.ReplaceAll(desc, "|", "")
	desc = strings.ReplaceAll(desc, "-----", "")

	return strings.TrimSpace(desc)
}
