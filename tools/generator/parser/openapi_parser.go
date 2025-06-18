package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// OpenAPIParser handles parsing of OpenAPI 3.0 specifications
type OpenAPIParser struct {
	normalizer *SchemaNormalizer
	extractor  *EndpointExtractor
	analyzer   *AuthAnalyzer
}

// NewOpenAPIParser creates a new OpenAPI parser
func NewOpenAPIParser() *OpenAPIParser {
	return &OpenAPIParser{
		normalizer: NewSchemaNormalizer(),
		extractor:  NewEndpointExtractor(),
		analyzer:   NewAuthAnalyzer(),
	}
}

// ParseFile parses an OpenAPI specification from a file
func (p *OpenAPIParser) ParseFile(path string) (*APISpec, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", path, err)
	}

	return p.ParseBytes(data)
}

// ParseBytes parses an OpenAPI specification from byte data
func (p *OpenAPIParser) ParseBytes(data []byte) (*APISpec, error) {
	var spec APISpec

	// Try YAML first, then JSON
	if err := yaml.Unmarshal(data, &spec); err != nil {
		if err := json.Unmarshal(data, &spec); err != nil {
			return nil, fmt.Errorf("failed to parse as YAML or JSON: %w", err)
		}
	}

	// Validate that this is an OpenAPI 3.x spec
	if err := p.validateSpec(&spec); err != nil {
		return nil, fmt.Errorf("invalid OpenAPI spec: %w", err)
	}

	// Schema extraction is now working correctly with YAML tags

	// Normalize schemas and resolve references
	if err := p.normalizer.NormalizeSpec(&spec); err != nil {
		return nil, fmt.Errorf("failed to normalize spec: %w", err)
	}

	// Normalization complete

	return &spec, nil
}

// validateSpec performs basic validation of the OpenAPI spec
func (p *OpenAPIParser) validateSpec(spec *APISpec) error {
	if spec.Info.Title == "" {
		return fmt.Errorf("missing required field: info.title")
	}

	if spec.Info.Version == "" {
		return fmt.Errorf("missing required field: info.version")
	}

	if len(spec.Paths) == 0 {
		return fmt.Errorf("no paths defined in specification")
	}

	return nil
}

// ExtractEndpoints extracts and classifies endpoints from the parsed spec
func (p *OpenAPIParser) ExtractEndpoints(spec *APISpec) ([]Endpoint, error) {
	return p.extractor.ExtractEndpoints(spec)
}

// AnalyzeAuth analyzes authentication and authorization requirements
func (p *OpenAPIParser) AnalyzeAuth(spec *APISpec) (*AuthInfo, error) {
	return p.analyzer.AnalyzeAuth(spec)
}

// GetResourceName generates a resource name from an operation
func (p *OpenAPIParser) GetResourceName(operation *Operation, path string) string {
	// Use operation ID if available
	if operation.OperationID != "" {
		return p.normalizeResourceName(operation.OperationID)
	}

	// Extract from path
	pathParts := strings.Split(strings.Trim(path, "/"), "/")
	if len(pathParts) > 0 {
		// Use the last non-parameter part
		for i := len(pathParts) - 1; i >= 0; i-- {
			part := pathParts[i]
			if !strings.HasPrefix(part, "{") && !strings.HasSuffix(part, "}") {
				return p.normalizeResourceName(part)
			}
		}
	}

	// Fallback to summary or description
	if operation.Summary != "" {
		return p.normalizeResourceName(operation.Summary)
	}

	return "unknown_resource"
}

// normalizeResourceName converts a string to a valid resource name
func (p *OpenAPIParser) normalizeResourceName(name string) string {
	// Convert to lowercase and replace non-alphanumeric with underscores
	result := strings.ToLower(name)
	result = strings.ReplaceAll(result, "-", "_")
	result = strings.ReplaceAll(result, " ", "_")

	// Remove common prefixes/suffixes
	prefixes := []string{"create", "get", "update", "delete", "list", "fetch"}
	suffixes := []string{"api", "endpoint", "resource"}

	for _, prefix := range prefixes {
		if strings.HasPrefix(result, prefix+"_") {
			result = strings.TrimPrefix(result, prefix+"_")
			break
		}
	}

	for _, suffix := range suffixes {
		if strings.HasSuffix(result, "_"+suffix) {
			result = strings.TrimSuffix(result, "_"+suffix)
			break
		}
	}

	// Ensure it's not empty
	if result == "" {
		result = "resource"
	}

	return result
}

// ClassifyEndpoint determines if an endpoint should be a resource or data source
func (p *OpenAPIParser) ClassifyEndpoint(method, path string, operation *Operation) ResourceClassification {
	classification := ResourceClassification{}

	// Determine CRUD type based on HTTP method
	switch strings.ToUpper(method) {
	case "GET":
		if p.isListEndpoint(path, operation) {
			classification.CRUDType = "list"
			classification.IsDataSource = true
		} else {
			classification.CRUDType = "read"
			classification.IsResource = true
			classification.IsDataSource = true
		}
	case "POST":
		classification.CRUDType = "create"
		classification.IsResource = true
	case "PUT", "PATCH":
		classification.CRUDType = "update"
		classification.IsResource = true
	case "DELETE":
		classification.CRUDType = "delete"
		classification.IsResource = true
	}

	classification.ResourceName = p.GetResourceName(operation, path)
	fmt.Println(classification)
	return classification
}

// isListEndpoint determines if a GET endpoint returns a list of items
func (p *OpenAPIParser) isListEndpoint(path string, operation *Operation) bool {
	// Check if path doesn't contain ID parameters
	if !strings.Contains(path, "{") {
		return true
	}

	// Check operation ID for list indicators
	if operation.OperationID != "" {
		listIndicators := []string{"list", "get_all", "fetch_all", "search"}
		operationLower := strings.ToLower(operation.OperationID)
		for _, indicator := range listIndicators {
			if strings.Contains(operationLower, indicator) {
				return true
			}
		}
	}

	// Check summary for list indicators
	if operation.Summary != "" {
		listIndicators := []string{"list", "get all", "fetch all", "search"}
		summaryLower := strings.ToLower(operation.Summary)
		for _, indicator := range listIndicators {
			if strings.Contains(summaryLower, indicator) {
				return true
			}
		}
	}

	return false
}
