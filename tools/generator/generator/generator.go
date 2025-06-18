package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mantisec/terraform-provider-umbrella/tools/generator/config"
	"github.com/mantisec/terraform-provider-umbrella/tools/generator/parser"
)

// Generator handles code generation from OpenAPI specifications
type Generator struct {
	config            *config.Config
	templateEngine    *TemplateEngine
	resourceGen       *ResourceGenerator
	dataSourceGen     *DataSourceGenerator
	docsGen           *DocsGenerator
	advancedSchemaGen *AdvancedSchemaGenerator
	clientMethodGen   *ClientMethodGenerator
	testGen           *TestGenerator
}

// NewGenerator creates a new code generator
func NewGenerator(cfg *config.Config) *Generator {
	templateEngine := NewTemplateEngine(cfg)

	return &Generator{
		config:            cfg,
		templateEngine:    templateEngine,
		resourceGen:       NewResourceGenerator(cfg, templateEngine),
		dataSourceGen:     NewDataSourceGenerator(cfg, templateEngine),
		docsGen:           NewDocsGenerator(cfg, templateEngine),
		advancedSchemaGen: NewAdvancedSchemaGenerator(templateEngine),
		clientMethodGen:   NewClientMethodGenerator(cfg, templateEngine),
		testGen:           NewTestGenerator(cfg, templateEngine),
	}
}

// GenerateFromSpec generates Terraform provider code from an OpenAPI specification
func (g *Generator) GenerateFromSpec(spec *parser.APISpec, outputDir string) error {
	// Extract endpoints from the spec
	extractor := parser.NewEndpointExtractor()
	endpoints, err := extractor.ExtractEndpoints(spec)
	if err != nil {
		return fmt.Errorf("failed to extract endpoints: %w", err)
	}

	// Filter endpoints based on configuration
	endpoints = extractor.FilterEndpoints(endpoints, g.config.Parsing.SkipEndpoints)

	// Group endpoints by resource
	resourceGroups := extractor.GroupEndpointsByResource(endpoints)

	// Generate resources and data sources
	for resourceName, resourceEndpoints := range resourceGroups {
		if err := g.generateResource(resourceName, resourceEndpoints, outputDir); err != nil {
			return fmt.Errorf("failed to generate resource %s: %w", resourceName, err)
		}
	}

	// Generate client methods using the new v2 generator
	if err := g.generateClientMethodsV2(resourceGroups, outputDir); err != nil {
		return fmt.Errorf("failed to generate client methods: %w", err)
	}

	// Generate documentation
	if err := g.generateDocumentation(resourceGroups, outputDir); err != nil {
		return fmt.Errorf("failed to generate documentation: %w", err)
	}

	// Generate tests
	if err := g.generateTests(resourceGroups, outputDir); err != nil {
		return fmt.Errorf("failed to generate tests: %w", err)
	}

	return nil
}

// generateResource generates code for a single resource
func (g *Generator) generateResource(resourceName string, endpoints []parser.Endpoint, outputDir string) error {
	// Separate resource and data source endpoints
	var resourceEndpoints, dataSourceEndpoints []parser.Endpoint

	for _, endpoint := range endpoints {
		if endpoint.ResourceType == "resource" {
			resourceEndpoints = append(resourceEndpoints, endpoint)
		} else if endpoint.ResourceType == "data_source" {
			dataSourceEndpoints = append(dataSourceEndpoints, endpoint)
		}
	}

	// Generate resource file if we have resource endpoints
	if len(resourceEndpoints) > 0 {
		if err := g.resourceGen.GenerateResource(resourceName, resourceEndpoints, outputDir); err != nil {
			return fmt.Errorf("failed to generate resource: %w", err)
		}
	}

	// Generate data source file if we have data source endpoints
	if len(dataSourceEndpoints) > 0 {
		if err := g.dataSourceGen.GenerateDataSource(resourceName, dataSourceEndpoints, outputDir); err != nil {
			return fmt.Errorf("failed to generate data source: %w", err)
		}
	}

	return nil
}

// ensureOutputDir ensures the output directory exists
func (g *Generator) ensureOutputDir(outputDir string) error {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory %s: %w", outputDir, err)
	}
	return nil
}

// getOutputPath generates the output file path for a resource
func (g *Generator) getOutputPath(outputDir, resourceName, fileType string) string {
	var pattern string

	switch fileType {
	case "resource":
		pattern = g.config.Output.ResourceFilePattern
	case "data_source":
		pattern = g.config.Output.DataSourceFilePattern
	default:
		pattern = "%s.go"
	}

	filename := fmt.Sprintf(pattern, resourceName)
	return filepath.Join(outputDir, filename)
}

// generateClientMethods generates client methods for all endpoints
func (g *Generator) generateClientMethods(endpoints []parser.Endpoint, outputDir string) error {
	clientMethods, err := g.clientMethodGen.GenerateClientMethods(endpoints)
	if err != nil {
		return fmt.Errorf("failed to generate client methods: %w", err)
	}

	// Write client methods to generated_client_methods.go
	outputPath := filepath.Join(outputDir, "generated_client_methods.go")

	// Create the complete file content
	fileContent := fmt.Sprintf(`package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

%s`, clientMethods)

	if err := os.WriteFile(outputPath, []byte(fileContent), 0644); err != nil {
		return fmt.Errorf("failed to write client methods file: %w", err)
	}

	return nil
}

// generateDocumentation generates documentation for all resources and data sources
func (g *Generator) generateDocumentation(resourceGroups map[string][]parser.Endpoint, outputDir string) error {
	docsDir := filepath.Join(outputDir, "docs")

	for resourceName, endpoints := range resourceGroups {
		// Separate resource and data source endpoints
		var resourceEndpoints, dataSourceEndpoints []parser.Endpoint

		for _, endpoint := range endpoints {
			if endpoint.ResourceType == "resource" {
				resourceEndpoints = append(resourceEndpoints, endpoint)
			} else if endpoint.ResourceType == "data_source" {
				dataSourceEndpoints = append(dataSourceEndpoints, endpoint)
			}
		}

		// Generate resource documentation
		if len(resourceEndpoints) > 0 {
			if err := g.docsGen.GenerateDocumentation(resourceName, resourceEndpoints, "resource", docsDir); err != nil {
				return fmt.Errorf("failed to generate resource documentation for %s: %w", resourceName, err)
			}
		}

		// Generate data source documentation
		if len(dataSourceEndpoints) > 0 {
			if err := g.docsGen.GenerateDocumentation(resourceName, dataSourceEndpoints, "data_source", docsDir); err != nil {
				return fmt.Errorf("failed to generate data source documentation for %s: %w", resourceName, err)
			}
		}
	}

	return nil
}

// generateTests generates test files for all resources and data sources
func (g *Generator) generateTests(resourceGroups map[string][]parser.Endpoint, outputDir string) error {
	testsDir := filepath.Join(outputDir, "tests")

	for resourceName, endpoints := range resourceGroups {
		// Separate resource and data source endpoints
		var resourceEndpoints, dataSourceEndpoints []parser.Endpoint

		for _, endpoint := range endpoints {
			if endpoint.ResourceType == "resource" {
				resourceEndpoints = append(resourceEndpoints, endpoint)
			} else if endpoint.ResourceType == "data_source" {
				dataSourceEndpoints = append(dataSourceEndpoints, endpoint)
			}
		}

		// Generate resource tests
		if len(resourceEndpoints) > 0 {
			if err := g.testGen.GenerateTests(resourceName, resourceEndpoints, "resource", testsDir); err != nil {
				return fmt.Errorf("failed to generate resource tests for %s: %w", resourceName, err)
			}
		}

		// Generate data source tests
		if len(dataSourceEndpoints) > 0 {
			if err := g.testGen.GenerateTests(resourceName, dataSourceEndpoints, "data_source", testsDir); err != nil {
				return fmt.Errorf("failed to generate data source tests for %s: %w", resourceName, err)
			}
		}
	}

	return nil
}

// ValidateGeneration validates that the generated code compiles
func (g *Generator) ValidateGeneration(outputDir string) error {
	// This would run go build or similar to validate the generated code
	// For now, we'll skip this validation step
	return nil
}

// generateClientMethodsV2 generates client methods using the new v2 generator
func (g *Generator) generateClientMethodsV2(resourceGroups map[string][]parser.Endpoint, outputDir string) error {
	// Convert old endpoint groups to new format
	endpointGroups := g.convertToEndpointGroups(resourceGroups)

	clientMethodGen := NewClientMethodGeneratorV2(g.templateEngine)

	var allMethods []ClientMethod
	for _, group := range endpointGroups {
		methods := clientMethodGen.GenerateClientMethodsForGroup(group)
		allMethods = append(allMethods, methods...)
	}

	// Generate the client methods file
	var methodsCode strings.Builder
	methodsCode.WriteString("package provider\n\n")
	methodsCode.WriteString("import (\n")
	methodsCode.WriteString("\t\"bytes\"\n")
	methodsCode.WriteString("\t\"context\"\n")
	methodsCode.WriteString("\t\"encoding/json\"\n")
	methodsCode.WriteString("\t\"fmt\"\n")
	methodsCode.WriteString("\t\"io\"\n")
	methodsCode.WriteString("\t\"net/http\"\n")
	methodsCode.WriteString("\t\"net/url\"\n")
	methodsCode.WriteString("\t\"strings\"\n")
	methodsCode.WriteString(")\n\n")

	for _, method := range allMethods {
		methodCode := clientMethodGen.GenerateClientMethodCode(method)
		methodsCode.WriteString(methodCode)
	}

	// Write to file
	outputPath := filepath.Join(outputDir, "generated_client_methods.go")
	if err := os.WriteFile(outputPath, []byte(methodsCode.String()), 0644); err != nil {
		return fmt.Errorf("failed to write client methods file: %w", err)
	}

	return nil
}

// convertToEndpointGroups converts old endpoint groups to new EndpointGroup format
func (g *Generator) convertToEndpointGroups(resourceGroups map[string][]parser.Endpoint) map[string]parser.EndpointGroup {
	endpointGroups := make(map[string]parser.EndpointGroup)

	for resourceName, endpoints := range resourceGroups {
		group := parser.EndpointGroup{
			CanonicalName: resourceName,
			ResourceName:  resourceName,
			Endpoints:     endpoints,
			CRUDOps:       make(map[string]parser.Endpoint),
			HasResource:   false,
			HasDataSource: false,
		}

		// Map CRUD operations and determine capabilities
		for _, endpoint := range endpoints {
			if endpoint.CRUDType != "" {
				group.CRUDOps[endpoint.CRUDType] = endpoint
			}

			if endpoint.ResourceType == "resource" {
				group.HasResource = true
			}
			if endpoint.ResourceType == "data_source" {
				group.HasDataSource = true
			}
		}

		endpointGroups[resourceName] = group
	}

	return endpointGroups
}
