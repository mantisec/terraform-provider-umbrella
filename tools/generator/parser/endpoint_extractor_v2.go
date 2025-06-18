package parser

import (
	"fmt"
	"regexp"
	"strings"
)

// EndpointExtractorV2 handles extraction and consolidation of API endpoints
type EndpointExtractorV2 struct {
	verbPrefixes []string
	verbSuffixes []string
	pathVerbMap  map[string]string
}

// NewEndpointExtractorV2 creates a new enhanced endpoint extractor
func NewEndpointExtractorV2() *EndpointExtractorV2 {
	return &EndpointExtractorV2{
		verbPrefixes: []string{
			"create", "add", "new", "post",
			"get", "fetch", "retrieve", "read", "list", "find",
			"update", "modify", "edit", "put", "patch",
			"delete", "remove", "destroy",
		},
		verbSuffixes: []string{
			"api", "endpoint", "resource", "service",
		},
		pathVerbMap: make(map[string]string),
	}
}

// EndpointGroup represents a consolidated group of endpoints for the same resource
type EndpointGroup struct {
	CanonicalName string
	ResourceName  string
	Endpoints     []Endpoint
	CRUDOps       map[string]Endpoint // "create", "read", "update", "delete", "list"
	HasResource   bool                // true if has create/update/delete operations
	HasDataSource bool                // true if has read/list operations
}

// ExtractAndConsolidate extracts endpoints and consolidates them by canonical name
func (e *EndpointExtractorV2) ExtractAndConsolidate(spec *APISpec) (map[string]EndpointGroup, error) {
	// First extract all endpoints
	endpoints, err := e.extractEndpoints(spec)
	if err != nil {
		return nil, fmt.Errorf("failed to extract endpoints: %w", err)
	}

	// Then consolidate by canonical name
	return e.consolidateEndpoints(endpoints), nil
}

// extractEndpoints extracts all endpoints from the spec
func (e *EndpointExtractorV2) extractEndpoints(spec *APISpec) ([]Endpoint, error) {
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

// classifyEndpoint classifies a single endpoint with enhanced logic
func (e *EndpointExtractorV2) classifyEndpoint(path, method string, operation *Operation) Endpoint {
	endpoint := Endpoint{
		Path:      path,
		Method:    method,
		Operation: operation,
	}

	// Generate canonical name (this is the key improvement)
	endpoint.ResourceName = e.getCanonicalName(operation, path, method)

	// Classify as resource or data source based on method
	classification := e.classifyResourceType(method, path, operation)
	endpoint.ResourceType = classification.ResourceType
	endpoint.CRUDType = classification.CRUDType

	return endpoint
}

// getCanonicalName generates a canonical resource name by removing verbs and normalizing
func (e *EndpointExtractorV2) getCanonicalName(operation *Operation, path, method string) string {
	var candidates []string

	// Candidate 1: Operation ID
	if operation.OperationID != "" {
		candidates = append(candidates, operation.OperationID)
	}

	// Candidate 2: Path-based name
	pathName := e.extractNameFromPath(path)
	if pathName != "" {
		candidates = append(candidates, pathName)
	}

	// Candidate 3: Summary
	if operation.Summary != "" {
		candidates = append(candidates, operation.Summary)
	}

	// Candidate 4: First tag
	if len(operation.Tags) > 0 {
		candidates = append(candidates, operation.Tags[0])
	}

	// Process each candidate and pick the best one
	var bestCandidate string
	var bestScore int

	for _, candidate := range candidates {
		normalized := e.normalizeToCanonical(candidate)
		score := e.scoreCandidate(normalized, path, method)

		if score > bestScore {
			bestScore = score
			bestCandidate = normalized
		}
	}

	if bestCandidate == "" {
		bestCandidate = "unknown_resource"
	}

	return bestCandidate
}

// extractNameFromPath extracts a resource name from the URL path
func (e *EndpointExtractorV2) extractNameFromPath(path string) string {
	// Remove leading/trailing slashes and split
	parts := strings.Split(strings.Trim(path, "/"), "/")

	// Find the best non-parameter part
	for i := len(parts) - 1; i >= 0; i-- {
		part := parts[i]

		// Skip path parameters
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			continue
		}

		// Skip common API prefixes
		if part == "api" || part == "v1" || part == "v2" || part == "v3" {
			continue
		}

		// This looks like a resource name
		if len(part) > 0 {
			return part
		}
	}

	return ""
}

// normalizeToCanonical converts a string to canonical form by removing verbs
func (e *EndpointExtractorV2) normalizeToCanonical(name string) string {
	// Convert to lowercase and handle camelCase
	result := e.camelCaseToSnakeCase(name)
	result = strings.ToLower(result)
	result = strings.ReplaceAll(result, "-", "_")
	result = strings.ReplaceAll(result, " ", "_")

	// Remove non-alphanumeric except underscores
	reg := regexp.MustCompile(`[^a-z0-9_]`)
	result = reg.ReplaceAllString(result, "_")

	// Remove multiple underscores
	reg = regexp.MustCompile(`_+`)
	result = reg.ReplaceAllString(result, "_")

	// Remove verb prefixes
	for _, prefix := range e.verbPrefixes {
		patterns := []string{
			prefix + "_",
			prefix,
		}
		for _, pattern := range patterns {
			if strings.HasPrefix(result, pattern) {
				result = strings.TrimPrefix(result, pattern)
				result = strings.TrimPrefix(result, "_")
				break
			}
		}
	}

	// Remove verb suffixes
	for _, suffix := range e.verbSuffixes {
		patterns := []string{
			"_" + suffix,
			suffix,
		}
		for _, pattern := range patterns {
			if strings.HasSuffix(result, pattern) {
				result = strings.TrimSuffix(result, pattern)
				result = strings.TrimSuffix(result, "_")
				break
			}
		}
	}

	// Remove common suffixes like "by_id", "by_name", etc.
	commonSuffixes := []string{"by_id", "by_name", "by_uuid", "byid", "byname", "byuuid"}
	for _, suffix := range commonSuffixes {
		if strings.HasSuffix(result, suffix) {
			result = strings.TrimSuffix(result, suffix)
			result = strings.TrimSuffix(result, "_")
			break
		}
	}

	// Convert to singular form (basic pluralization rules)
	result = e.singularize(result)

	// Clean up
	result = strings.Trim(result, "_")
	if result == "" {
		result = "resource"
	}

	return result
}

// singularize converts plural nouns to singular (improved implementation)
func (e *EndpointExtractorV2) singularize(word string) string {
	if len(word) < 2 {
		return word
	}

	// Special cases first
	specialCases := map[string]string{
		"policies":  "policy",
		"addresses": "address",
		"lives":     "life",
		"knives":    "knife",
		"wives":     "wife",
		"leaves":    "leaf",
		"thieves":   "thief",
		"children":  "child",
		"people":    "person",
		"men":       "man",
		"women":     "woman",
		"feet":      "foot",
		"teeth":     "tooth",
		"geese":     "goose",
		"mice":      "mouse",
	}

	if singular, exists := specialCases[word]; exists {
		return singular
	}

	// Pattern-based rules (order matters)
	patterns := []struct {
		pattern     string
		replacement string
	}{
		{"ies$", "y"},   // policies -> policy, companies -> company
		{"ves$", "f"},   // lives -> life, knives -> knife
		{"oes$", "o"},   // heroes -> hero, potatoes -> potato
		{"ses$", "s"},   // addresses -> address, classes -> class
		{"ches$", "ch"}, // matches -> match, churches -> church
		{"shes$", "sh"}, // wishes -> wish, dishes -> dish
		{"xes$", "x"},   // boxes -> box, fixes -> fix
		{"zes$", "z"},   // prizes -> prize, sizes -> size
		{"men$", "man"}, // workmen -> workman
		{"s$", ""},      // networks -> network, users -> user
	}

	for _, p := range patterns {
		reg := regexp.MustCompile(p.pattern)
		if reg.MatchString(word) {
			return reg.ReplaceAllString(word, p.replacement)
		}
	}

	return word
}

// scoreCandidate scores how good a candidate name is
func (e *EndpointExtractorV2) scoreCandidate(candidate, path, method string) int {
	score := 0

	// Prefer longer names (more descriptive)
	score += len(candidate)

	// Prefer names that appear in the path
	if strings.Contains(path, candidate) {
		score += 10
	}

	// Penalize generic names
	genericNames := []string{"resource", "item", "object", "data", "info"}
	for _, generic := range genericNames {
		if candidate == generic {
			score -= 5
		}
	}

	// Prefer names that don't contain method verbs
	methodWords := []string{"get", "post", "put", "delete", "patch"}
	for _, methodWord := range methodWords {
		if strings.Contains(candidate, methodWord) {
			score -= 3
		}
	}

	return score
}

// consolidateEndpoints groups endpoints by canonical name
func (e *EndpointExtractorV2) consolidateEndpoints(endpoints []Endpoint) map[string]EndpointGroup {
	groups := make(map[string]EndpointGroup)

	for _, endpoint := range endpoints {
		canonicalName := endpoint.ResourceName

		group, exists := groups[canonicalName]
		if !exists {
			group = EndpointGroup{
				CanonicalName: canonicalName,
				ResourceName:  canonicalName,
				Endpoints:     []Endpoint{},
				CRUDOps:       make(map[string]Endpoint),
				HasResource:   false,
				HasDataSource: false,
			}
		}

		// Add endpoint to group
		group.Endpoints = append(group.Endpoints, endpoint)

		// Map CRUD operation
		if endpoint.CRUDType != "" {
			group.CRUDOps[endpoint.CRUDType] = endpoint
		}

		// Update flags
		if endpoint.ResourceType == "resource" {
			group.HasResource = true
		}
		if endpoint.ResourceType == "data_source" {
			group.HasDataSource = true
		}

		groups[canonicalName] = group
	}

	return groups
}

// classifyResourceType determines if an endpoint should be a resource or data source
func (e *EndpointExtractorV2) classifyResourceType(method, path string, operation *Operation) ResourceTypeClassification {
	classification := ResourceTypeClassification{}

	switch strings.ToUpper(method) {
	case "GET":
		if e.isListEndpoint(path, operation) {
			classification.ResourceType = "data_source"
			classification.CRUDType = "list"
		} else {
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

// isListEndpoint determines if a GET endpoint returns a list of items
func (e *EndpointExtractorV2) isListEndpoint(path string, operation *Operation) bool {
	// Check if path doesn't contain ID parameters
	if !strings.Contains(path, "{") {
		return true
	}

	// Check operation ID for list indicators
	if operation.OperationID != "" {
		listIndicators := []string{"list", "get_all", "fetch_all", "search", "find_all"}
		operationLower := strings.ToLower(operation.OperationID)
		for _, indicator := range listIndicators {
			if strings.Contains(operationLower, indicator) {
				return true
			}
		}
	}

	// Check summary for list indicators
	if operation.Summary != "" {
		listIndicators := []string{"list", "get all", "fetch all", "search", "find all"}
		summaryLower := strings.ToLower(operation.Summary)
		for _, indicator := range listIndicators {
			if strings.Contains(summaryLower, indicator) {
				return true
			}
		}
	}

	return false
}

// GetConsolidationReport generates a report of the consolidation process
func (e *EndpointExtractorV2) GetConsolidationReport(groups map[string]EndpointGroup) ConsolidationReport {
	report := ConsolidationReport{
		TotalGroups:      len(groups),
		ResourceGroups:   0,
		DataSourceGroups: 0,
		MixedGroups:      0,
		Groups:           make([]GroupSummary, 0, len(groups)),
	}

	for _, group := range groups {
		summary := GroupSummary{
			CanonicalName: group.CanonicalName,
			EndpointCount: len(group.Endpoints),
			HasResource:   group.HasResource,
			HasDataSource: group.HasDataSource,
			CRUDOps:       make([]string, 0, len(group.CRUDOps)),
		}

		for crudType := range group.CRUDOps {
			summary.CRUDOps = append(summary.CRUDOps, crudType)
		}

		if group.HasResource && group.HasDataSource {
			report.MixedGroups++
		} else if group.HasResource {
			report.ResourceGroups++
		} else if group.HasDataSource {
			report.DataSourceGroups++
		}

		report.Groups = append(report.Groups, summary)
	}

	return report
}

// ConsolidationReport provides insights into the consolidation process
type ConsolidationReport struct {
	TotalGroups      int
	ResourceGroups   int
	DataSourceGroups int
	MixedGroups      int
	Groups           []GroupSummary
}

// GroupSummary summarizes a consolidated endpoint group
type GroupSummary struct {
	CanonicalName string
	EndpointCount int
	HasResource   bool
	HasDataSource bool
	CRUDOps       []string
}

// String returns a string representation of the consolidation report
func (r ConsolidationReport) String() string {
	return fmt.Sprintf("Consolidation Report: %d total groups (%d resource, %d data-source, %d mixed)",
		r.TotalGroups, r.ResourceGroups, r.DataSourceGroups, r.MixedGroups)
}

// camelCaseToSnakeCase converts camelCase to snake_case
func (e *EndpointExtractorV2) camelCaseToSnakeCase(str string) string {
	// Insert underscore before uppercase letters (except the first character)
	reg := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	result := reg.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(result)
}
