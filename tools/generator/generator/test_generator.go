package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mantisec/terraform-provider-umbrella/tools/generator/config"
	"github.com/mantisec/terraform-provider-umbrella/tools/generator/parser"
)

// TestGenerator handles generation of test files for resources and data sources
type TestGenerator struct {
	config         *config.Config
	templateEngine *TemplateEngine
}

// NewTestGenerator creates a new test generator
func NewTestGenerator(cfg *config.Config, templateEngine *TemplateEngine) *TestGenerator {
	return &TestGenerator{
		config:         cfg,
		templateEngine: templateEngine,
	}
}

// TestData contains data for test template generation
type TestData struct {
	Name           string // Add this field for template compatibility
	ResourceName   string
	ResourceType   string // "resource" or "data_source"
	TypeName       string
	PackageName    string
	StructName     string
	TestFunctions  []TestFunction
	MockResponses  []MockResponse
	TestAttributes []TestAttribute
	ImportID       string
}

// TestFunction represents a test function
type TestFunction struct {
	Name        string
	Description string
	TestType    string // "unit", "acceptance", "validation"
	Code        string
}

// MockResponse represents a mock API response
type MockResponse struct {
	Method     string
	Path       string
	StatusCode int
	Body       string
	Headers    map[string]string
}

// TestAttribute represents an attribute for testing
type TestAttribute struct {
	Name          string
	TestValue     string
	UpdateValue   string
	ValidValues   []string
	InvalidValues []string
}

// GenerateTests generates test files for a resource or data source
func (tg *TestGenerator) GenerateTests(resourceName string, endpoints []parser.Endpoint, resourceType string, outputDir string) error {
	// Prepare test data
	data := tg.prepareTestData(resourceName, endpoints, resourceType)

	// Register test templates
	if err := tg.registerTestTemplates(); err != nil {
		return fmt.Errorf("failed to register test templates: %w", err)
	}

	// Generate unit tests
	if err := tg.generateUnitTests(data, outputDir); err != nil {
		return fmt.Errorf("failed to generate unit tests: %w", err)
	}

	// Generate acceptance tests
	if err := tg.generateAcceptanceTests(data, outputDir); err != nil {
		return fmt.Errorf("failed to generate acceptance tests: %w", err)
	}

	// Generate validation tests
	if err := tg.generateValidationTests(data, outputDir); err != nil {
		return fmt.Errorf("failed to generate validation tests: %w", err)
	}

	return nil
}

// prepareTestData prepares the data structure for test generation
func (tg *TestGenerator) prepareTestData(resourceName string, endpoints []parser.Endpoint, resourceType string) *TestData {
	data := &TestData{
		Name:           resourceName, // Set the Name field
		ResourceName:   resourceName,
		ResourceType:   resourceType,
		TypeName:       tg.config.Global.ProviderName + "_" + resourceName,
		PackageName:    tg.config.Global.PackageName,
		StructName:     tg.templateEngine.toPascalCase(resourceName) + "Resource",
		TestAttributes: tg.generateTestAttributes(endpoints),
		MockResponses:  tg.generateMockResponses(endpoints),
		ImportID:       "test-id-12345",
	}

	// Generate test functions
	data.TestFunctions = tg.generateTestFunctions(data)

	return data
}

// generateTestAttributes generates test attributes from endpoints
func (tg *TestGenerator) generateTestAttributes(endpoints []parser.Endpoint) []TestAttribute {
	var attributes []TestAttribute

	// Add common test attributes
	attributes = append(attributes, TestAttribute{
		Name:          "name",
		TestValue:     "\"test-resource\"",
		UpdateValue:   "\"test-resource-updated\"",
		ValidValues:   []string{"\"valid-name\"", "\"another-valid-name\""},
		InvalidValues: []string{"\"\"", "\"invalid name with spaces\""},
	})

	// Extract attributes from endpoint schemas
	for _, endpoint := range endpoints {
		if endpoint.Operation.RequestBody != nil {
			attributes = append(attributes, tg.extractTestAttributesFromRequestBody(endpoint.Operation.RequestBody)...)
		}
	}

	return tg.removeDuplicateTestAttributes(attributes)
}

// extractTestAttributesFromRequestBody extracts test attributes from request body
func (tg *TestGenerator) extractTestAttributesFromRequestBody(requestBody *parser.RequestBody) []TestAttribute {
	var attributes []TestAttribute

	for _, mediaType := range requestBody.Content {
		if mediaType.Schema != nil && mediaType.Schema.Properties != nil {
			for propName, propSchema := range mediaType.Schema.Properties {
				if propName == "id" || propName == "name" {
					continue // Skip common attributes
				}

				attr := TestAttribute{
					Name:        propName,
					TestValue:   tg.generateTestValue(propSchema),
					UpdateValue: tg.generateUpdateValue(propSchema),
				}

				// Generate valid/invalid values based on schema
				attr.ValidValues, attr.InvalidValues = tg.generateValidationValues(propSchema)

				attributes = append(attributes, attr)
			}
		}
	}

	return attributes
}

// generateTestValue generates a test value for a schema
func (tg *TestGenerator) generateTestValue(schema *parser.Schema) string {
	if schema == nil {
		return "\"test-value\""
	}

	switch schema.Type {
	case "string":
		if len(schema.Enum) > 0 {
			return fmt.Sprintf("\"%v\"", schema.Enum[0])
		}
		return "\"test-value\""
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
		return "\"test-value\""
	}
}

// generateUpdateValue generates an update test value for a schema
func (tg *TestGenerator) generateUpdateValue(schema *parser.Schema) string {
	if schema == nil {
		return "\"updated-value\""
	}

	switch schema.Type {
	case "string":
		if len(schema.Enum) > 1 {
			return fmt.Sprintf("\"%v\"", schema.Enum[1])
		}
		return "\"updated-value\""
	case "integer":
		return "456"
	case "number":
		return "456.78"
	case "boolean":
		return "false"
	case "array":
		return "[\"updated1\", \"updated2\"]"
	case "object":
		return "{}"
	default:
		return "\"updated-value\""
	}
}

// generateValidationValues generates valid and invalid values for testing
func (tg *TestGenerator) generateValidationValues(schema *parser.Schema) ([]string, []string) {
	var validValues, invalidValues []string

	if schema == nil {
		return []string{"\"valid\""}, []string{"\"\""}
	}

	switch schema.Type {
	case "string":
		validValues = []string{"\"valid-string\"", "\"another-valid\""}
		invalidValues = []string{"\"\""}

		if len(schema.Enum) > 0 {
			validValues = nil
			for _, enum := range schema.Enum {
				validValues = append(validValues, fmt.Sprintf("\"%v\"", enum))
			}
			invalidValues = []string{"\"invalid-enum-value\""}
		}
	case "integer":
		validValues = []string{"1", "100", "999"}
		invalidValues = []string{"-1", "\"not-a-number\""}
	case "boolean":
		validValues = []string{"true", "false"}
		invalidValues = []string{"\"not-a-boolean\"", "1"}
	}

	return validValues, invalidValues
}

// removeDuplicateTestAttributes removes duplicate test attributes
func (tg *TestGenerator) removeDuplicateTestAttributes(attributes []TestAttribute) []TestAttribute {
	seen := make(map[string]bool)
	var result []TestAttribute

	for _, attr := range attributes {
		if !seen[attr.Name] {
			seen[attr.Name] = true
			result = append(result, attr)
		}
	}

	return result
}

// generateMockResponses generates mock API responses for testing
func (tg *TestGenerator) generateMockResponses(endpoints []parser.Endpoint) []MockResponse {
	var responses []MockResponse

	for _, endpoint := range endpoints {
		response := MockResponse{
			Method:     endpoint.Method,
			Path:       endpoint.Path,
			StatusCode: tg.getSuccessStatusCode(endpoint.Method),
			Headers:    map[string]string{"Content-Type": "application/json"},
		}

		// Generate response body based on operation
		response.Body = tg.generateMockResponseBody(endpoint)

		responses = append(responses, response)
	}

	return responses
}

// getSuccessStatusCode returns the expected success status code for a method
func (tg *TestGenerator) getSuccessStatusCode(method string) int {
	switch method {
	case "GET":
		return 200
	case "POST":
		return 201
	case "PUT", "PATCH":
		return 200
	case "DELETE":
		return 204
	default:
		return 200
	}
}

// generateMockResponseBody generates a mock response body
func (tg *TestGenerator) generateMockResponseBody(endpoint parser.Endpoint) string {
	switch endpoint.Method {
	case "GET":
		if strings.Contains(endpoint.Path, "{") {
			// Single resource
			return `{"id": "test-id-12345", "name": "test-resource"}`
		} else {
			// List of resources
			return `[{"id": "test-id-12345", "name": "test-resource"}]`
		}
	case "POST", "PUT", "PATCH":
		return `{"id": "test-id-12345", "name": "test-resource"}`
	case "DELETE":
		return ""
	default:
		return `{"id": "test-id-12345"}`
	}
}

// generateTestFunctions generates test functions
func (tg *TestGenerator) generateTestFunctions(data *TestData) []TestFunction {
	var functions []TestFunction

	if data.ResourceType == "resource" {
		// Resource tests
		functions = append(functions, TestFunction{
			Name:        fmt.Sprintf("Test%sResource_basic", tg.templateEngine.toPascalCase(data.ResourceName)),
			Description: fmt.Sprintf("Test basic %s resource functionality", data.ResourceName),
			TestType:    "acceptance",
		})

		functions = append(functions, TestFunction{
			Name:        fmt.Sprintf("Test%sResource_update", tg.templateEngine.toPascalCase(data.ResourceName)),
			Description: fmt.Sprintf("Test %s resource update functionality", data.ResourceName),
			TestType:    "acceptance",
		})

		functions = append(functions, TestFunction{
			Name:        fmt.Sprintf("Test%sResource_import", tg.templateEngine.toPascalCase(data.ResourceName)),
			Description: fmt.Sprintf("Test %s resource import functionality", data.ResourceName),
			TestType:    "acceptance",
		})
	} else {
		// Data source tests
		functions = append(functions, TestFunction{
			Name:        fmt.Sprintf("Test%sDataSource_basic", tg.templateEngine.toPascalCase(data.ResourceName)),
			Description: fmt.Sprintf("Test basic %s data source functionality", data.ResourceName),
			TestType:    "acceptance",
		})
	}

	// Schema validation tests
	functions = append(functions, TestFunction{
		Name:        fmt.Sprintf("Test%sSchema_validation", tg.templateEngine.toPascalCase(data.ResourceName)),
		Description: fmt.Sprintf("Test %s schema validation", data.ResourceName),
		TestType:    "validation",
	})

	return functions
}

// generateUnitTests generates unit test files
func (tg *TestGenerator) generateUnitTests(data *TestData, outputDir string) error {
	content, err := tg.templateEngine.ExecuteTemplate("unit_test.go", data)
	if err != nil {
		return fmt.Errorf("failed to execute unit test template: %w", err)
	}

	filename := fmt.Sprintf("%s_test.go", data.ResourceName)
	outputPath := filepath.Join(outputDir, filename)

	return tg.writeTestFile(outputPath, content)
}

// generateAcceptanceTests generates acceptance test files
func (tg *TestGenerator) generateAcceptanceTests(data *TestData, outputDir string) error {
	content, err := tg.templateEngine.ExecuteTemplate("acceptance_test.go", data)
	if err != nil {
		return fmt.Errorf("failed to execute acceptance test template: %w", err)
	}

	filename := fmt.Sprintf("%s_acceptance_test.go", data.ResourceName)
	outputPath := filepath.Join(outputDir, filename)

	return tg.writeTestFile(outputPath, content)
}

// generateValidationTests generates validation test files
func (tg *TestGenerator) generateValidationTests(data *TestData, outputDir string) error {
	content, err := tg.templateEngine.ExecuteTemplate("validation_test.go", data)
	if err != nil {
		return fmt.Errorf("failed to execute validation test template: %w", err)
	}

	filename := fmt.Sprintf("%s_validation_test.go", data.ResourceName)
	outputPath := filepath.Join(outputDir, filename)

	return tg.writeTestFile(outputPath, content)
}

// writeTestFile writes a test file to disk
func (tg *TestGenerator) writeTestFile(outputPath, content string) error {
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("failed to create test output directory: %w", err)
	}

	if err := os.WriteFile(outputPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write test file: %w", err)
	}

	return nil
}

// registerTestTemplates registers all test templates
func (tg *TestGenerator) registerTestTemplates() error {
	templates := map[string]string{
		"unit_test.go":       tg.getUnitTestTemplate(),
		"acceptance_test.go": tg.getAcceptanceTestTemplate(),
		"validation_test.go": tg.getValidationTestTemplate(),
	}

	for name, template := range templates {
		if err := tg.templateEngine.RegisterTemplate(name, template); err != nil {
			return fmt.Errorf("failed to register template %s: %w", name, err)
		}
	}

	return nil
}

// getUnitTestTemplate returns the unit test template
func (tg *TestGenerator) getUnitTestTemplate() string {
	return `package {{.PackageName}}

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

func Test{{pascalCase .ResourceName}}{{if eq .ResourceType "resource"}}Resource{{else}}DataSource{{end}}_Schema(t *testing.T) {
	{{if eq .ResourceType "resource"}}
	r := New{{.StructName}}()
	{{else}}
	r := New{{.StructName}}()
	{{end}}

	// Test that the resource implements the correct interface
	{{if eq .ResourceType "resource"}}
	var _ resource.Resource = r
	{{else}}
	var _ datasource.DataSource = r
	{{end}}
}

{{range .TestFunctions}}
{{if eq .TestType "unit"}}
func {{.Name}}(t *testing.T) {
	// {{.Description}}
	// Test basic functionality
	assert.NotNil(t, r, "Resource should not be nil")

	// Test schema validation
	schemaReq := resource.SchemaRequest{}
	schemaResp := &resource.SchemaResponse{}
	r.Schema(context.Background(), schemaReq, schemaResp)

	assert.False(t, schemaResp.Diagnostics.HasError(), "Schema should not have errors")
	assert.NotNil(t, schemaResp.Schema, "Schema should not be nil")
}
{{end}}
{{end}}
`
}

// getAcceptanceTestTemplate returns the acceptance test template
func (tg *TestGenerator) getAcceptanceTestTemplate() string {
	return `package {{.PackageName}}

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

{{range .TestFunctions}}
{{if eq .TestType "acceptance"}}
func {{.Name}}(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAcc{{pascalCase $.ResourceName}}{{if eq $.ResourceType "resource"}}Resource{{else}}DataSource{{end}}Config_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					{{if eq $.ResourceType "resource"}}
					resource.TestCheckResourceAttr("{{$.TypeName}}.test", "name", "test-{{$.ResourceName}}"),
					{{else}}
					resource.TestCheckDataSourceAttr("{{$.TypeName}}.test", "name", "test-{{$.ResourceName}}"),
					{{end}}
				),
			},
			{{if eq $.ResourceType "resource"}}
			{
				Config: testAcc{{pascalCase $.ResourceName}}ResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("{{$.TypeName}}.test", "name", "test-{{$.ResourceName}}-updated"),
				),
			},
			{{end}}
		},
	})
}
{{end}}
{{end}}

func testAcc{{pascalCase .ResourceName}}{{if eq .ResourceType "resource"}}Resource{{else}}DataSource{{end}}Config_basic() string {
	return ` + "`" + `
{{if eq .ResourceType "resource"}}
resource "{{.TypeName}}" "test" {
  name = "test-{{.ResourceName}}"
  {{range .TestAttributes}}
  {{.Name}} = {{.TestValue}}
  {{end}}
}
{{else}}
data "{{.TypeName}}" "test" {
  id = "{{.ImportID}}"
}
{{end}}
` + "`" + `
}

{{if eq .ResourceType "resource"}}
func testAcc{{pascalCase .ResourceName}}ResourceConfig_update() string {
	return ` + "`" + `
resource "{{.TypeName}}" "test" {
  name = "test-{{.ResourceName}}-updated"
  {{range .TestAttributes}}
  {{.Name}} = {{.UpdateValue}}
  {{end}}
}
` + "`" + `
}
{{end}}
`
}

// getValidationTestTemplate returns the validation test template
func (tg *TestGenerator) getValidationTestTemplate() string {
	return `package {{.PackageName}}

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

{{range .TestFunctions}}
{{if eq .TestType "validation"}}
func {{.Name}}(t *testing.T) {
	// Test valid configurations
	{{range $.TestAttributes}}
	{{range .ValidValues}}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAcc{{pascalCase $.ResourceName}}Config_{{$.Name}}_valid({{.}}),
				Check: resource.ComposeAggregateTestCheckFunc(
					{{if eq $.ResourceType "resource"}}
					resource.TestCheckResourceAttr("{{$.TypeName}}.test", "{{$.Name}}", {{.}}),
					{{else}}
					resource.TestCheckDataSourceAttr("{{$.TypeName}}.test", "{{$.Name}}", {{.}}),
					{{end}}
				),
			},
		},
	})
	{{end}}
	{{end}}

	// Test invalid configurations
	{{range $.TestAttributes}}
	{{range .InvalidValues}}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAcc{{pascalCase $.ResourceName}}Config_{{$.Name}}_invalid({{.}}),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})
	{{end}}
	{{end}}
}
{{end}}
{{end}}

{{range .TestAttributes}}
func testAcc{{pascalCase $.ResourceName}}Config_{{.Name}}_valid(value string) string {
	return fmt.Sprintf(` + "`" + `
{{if eq $.ResourceType "resource"}}
resource "{{$.TypeName}}" "test" {
  {{.Name}} = %s
}
{{else}}
data "{{$.TypeName}}" "test" {
  {{.Name}} = %s
}
{{end}}
` + "`" + `, value)
}

func testAcc{{pascalCase $.ResourceName}}Config_{{.Name}}_invalid(value string) string {
	return fmt.Sprintf(` + "`" + `
{{if eq $.ResourceType "resource"}}
resource "{{$.TypeName}}" "test" {
  {{.Name}} = %s
}
{{else}}
data "{{$.TypeName}}" "test" {
  {{.Name}} = %s
}
{{end}}
` + "`" + `, value)
}
{{end}}
`
}
