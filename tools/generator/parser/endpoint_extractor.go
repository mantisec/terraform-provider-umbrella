package parser

import (
	"strings"
)

// EndpointExtractor handles extraction and classification of API endpoints
type EndpointExtractor struct{}

// NewEndpointExtractor creates a new endpoint extractor
func NewEndpointExtractor() *EndpointExtractor {
	return &EndpointExtractor{}
}

// ExtractEndpoints extracts and classifies all endpoints from an API spec
func (e *EndpointExtractor) ExtractEndpoints(spec *APISpec) ([]Endpoint, error) {
	var endpoints []Endpoint

	for path, pathItem := range spec.Paths {
		operations := map[string]*Operation{
			"GET":    pathItem.Get,
			"POST":   pathItem.Post,
			"PUT":    pathItem.Put,
			"DELETE": pathItem.Delete,
			"PATCH":  pathItem.Patch,
		}

		for method, operation := range operations {
			if operation == nil {
				continue
			}

			endpoint := e.classifyEndpoint(path, method, operation)
			endpoints = append(endpoints, endpoint)
		}
	}

	return endpoints, nil
}

// classifyEndpoint classifies a single endpoint
func (e *EndpointExtractor) classifyEndpoint(path, method string, operation *Operation) Endpoint {
	endpoint := Endpoint{
		Path:      path,
		Method:    method,
		Operation: operation,
	}

	// Determine resource name
	endpoint.ResourceName = e.getResourceName(operation, path)

	// Classify as resource or data source based on method and path
	classification := e.classifyResourceType(method, path, operation)
	endpoint.ResourceType = classification.ResourceType
	endpoint.CRUDType = classification.CRUDType

	return endpoint
}

// ResourceTypeClassification represents the classification result
type ResourceTypeClassification struct {
	ResourceType string // "resource" or "data_source"
	CRUDType     string // "create", "read", "update", "delete", "list"
}

// classifyResourceType determines if an endpoint should be a resource or data source
func (e *EndpointExtractor) classifyResourceType(method, path string, operation *Operation) ResourceTypeClassification {
	classification := ResourceTypeClassification{}

	switch strings.ToUpper(method) {
	case "GET":
		if e.isListEndpoint(path, operation) {
			classification.ResourceType = "data_source"
			classification.CRUDType = "list"
		} else {
			// Single item GET can be both resource (for import) and data source
			classification.ResourceType = "data_source"
			classification.CRUDType = "read"
		}
	case "POST":
		classification.ResourceType = "resource"
		classification.CRUDType = "create"
	case "PUT", "PATCH":
		classification.ResourceType = "resource"
		classification.CRUDType = "update"
	case "DELETE":
		classification.ResourceType = "resource"
		classification.CRUDType = "delete"
	default:
		classification.ResourceType = "resource"
		classification.CRUDType = "unknown"
	}

	return classification
}

// getResourceName generates a resource name from operation and path
func (e *EndpointExtractor) getResourceName(operation *Operation, path string) string {
	// Use operation ID if available
	if operation.OperationID != "" {
		return e.normalizeResourceName(operation.OperationID)
	}

	// Extract from path
	pathParts := strings.Split(strings.Trim(path, "/"), "/")
	if len(pathParts) > 0 {
		// Use the last non-parameter part
		for i := len(pathParts) - 1; i >= 0; i-- {
			part := pathParts[i]
			if !strings.HasPrefix(part, "{") && !strings.HasSuffix(part, "}") {
				return e.normalizeResourceName(part)
			}
		}
	}

	// Fallback to summary
	if operation.Summary != "" {
		return e.normalizeResourceName(operation.Summary)
	}

	return "unknown_resource"
}

// normalizeResourceName converts a string to a valid resource name
func (e *EndpointExtractor) normalizeResourceName(name string) string {
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

// isListEndpoint determines if a GET endpoint returns a list of items
func (e *EndpointExtractor) isListEndpoint(path string, operation *Operation) bool {
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

// GroupEndpointsByResource groups endpoints by their resource name
func (e *EndpointExtractor) GroupEndpointsByResource(endpoints []Endpoint) map[string][]Endpoint {
	groups := make(map[string][]Endpoint)

	for _, endpoint := range endpoints {
		resourceName := endpoint.ResourceName
		groups[resourceName] = append(groups[resourceName], endpoint)
	}

	return groups
}

// FilterEndpoints filters endpoints based on criteria
func (e *EndpointExtractor) FilterEndpoints(endpoints []Endpoint, skipPaths []string) []Endpoint {
	var filtered []Endpoint

	for _, endpoint := range endpoints {
		skip := false
		for _, skipPath := range skipPaths {
			if strings.Contains(endpoint.Path, skipPath) {
				skip = true
				break
			}
		}

		if !skip {
			filtered = append(filtered, endpoint)
		}
	}

	return filtered
}
