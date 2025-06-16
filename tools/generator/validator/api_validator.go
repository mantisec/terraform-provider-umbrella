package validator

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/mantisec/terraform-provider-umbrella/tools/generator/config"
	"github.com/mantisec/terraform-provider-umbrella/tools/generator/parser"
)

// APIValidator validates API compatibility and compliance
type APIValidator struct {
	config *config.AdvancedConfig
}

// APIValidationResult represents the result of API validation
type APIValidationResult struct {
	Valid         bool                 `json:"valid"`
	Errors        []APIError           `json:"errors"`
	Warnings      []APIWarning         `json:"warnings"`
	Summary       APIValidationSummary `json:"summary"`
	Compatibility CompatibilityReport  `json:"compatibility"`
}

// APIError represents an API validation error
type APIError struct {
	Type       string `json:"type"`
	Message    string `json:"message"`
	Path       string `json:"path"`
	Method     string `json:"method,omitempty"`
	Suggestion string `json:"suggestion,omitempty"`
	Severity   string `json:"severity"`
}

// APIWarning represents an API validation warning
type APIWarning struct {
	Type       string `json:"type"`
	Message    string `json:"message"`
	Path       string `json:"path"`
	Method     string `json:"method,omitempty"`
	Suggestion string `json:"suggestion,omitempty"`
}

// APIValidationSummary provides a summary of API validation results
type APIValidationSummary struct {
	TotalEndpoints     int     `json:"total_endpoints"`
	ValidEndpoints     int     `json:"valid_endpoints"`
	ErrorCount         int     `json:"error_count"`
	WarningCount       int     `json:"warning_count"`
	CompatibilityScore float64 `json:"compatibility_score"`
	RESTCompliance     float64 `json:"rest_compliance"`
}

// CompatibilityReport provides detailed compatibility information
type CompatibilityReport struct {
	TerraformCompatible bool                 `json:"terraform_compatible"`
	RESTCompliant       bool                 `json:"rest_compliant"`
	SecurityCompliant   bool                 `json:"security_compliant"`
	Issues              []CompatibilityIssue `json:"issues"`
	Recommendations     []string             `json:"recommendations"`
}

// CompatibilityIssue represents a compatibility issue
type CompatibilityIssue struct {
	Type        string `json:"type"`
	Severity    string `json:"severity"`
	Description string `json:"description"`
	Path        string `json:"path"`
	Solution    string `json:"solution"`
}

// NewAPIValidator creates a new API validator
func NewAPIValidator(config *config.AdvancedConfig) *APIValidator {
	return &APIValidator{
		config: config,
	}
}

// ValidateAPISpec validates an entire API specification for compatibility
func (v *APIValidator) ValidateAPISpec(spec *parser.APISpec) (*APIValidationResult, error) {
	result := &APIValidationResult{
		Valid:    true,
		Errors:   make([]APIError, 0),
		Warnings: make([]APIWarning, 0),
		Summary: APIValidationSummary{
			TotalEndpoints: 0,
			ValidEndpoints: 0,
		},
		Compatibility: CompatibilityReport{
			Issues:          make([]CompatibilityIssue, 0),
			Recommendations: make([]string, 0),
		},
	}

	// Validate API info
	v.validateAPIInfo(&spec.Info, result)

	// Validate servers
	v.validateServers(spec.Servers, result)

	// Validate security schemes
	v.validateSecuritySchemes(spec.Components.SecuritySchemes, result)

	// Validate paths and operations
	for path, pathItem := range spec.Paths {
		v.validatePath(path, &pathItem, result)
	}

	// Validate overall API design
	v.validateAPIDesign(spec, result)

	// Calculate summary
	v.calculateSummary(result)

	// Generate compatibility report
	v.generateCompatibilityReport(spec, result)

	result.Valid = result.Summary.ErrorCount == 0

	return result, nil
}

// validateAPIInfo validates API information
func (v *APIValidator) validateAPIInfo(info *parser.APIInfo, result *APIValidationResult) {
	if info.Title == "" {
		v.addError(result, "missing_title", "API title is required", "info.title", "",
			"Add a descriptive title for the API", "error")
	}

	if info.Version == "" {
		v.addError(result, "missing_version", "API version is required", "info.version", "",
			"Add version information (e.g., '1.0.0')", "error")
	} else if !v.isValidVersion(info.Version) {
		v.addWarning(result, "invalid_version_format", "API version should follow semantic versioning",
			"info.version", "", "Use semantic versioning format (e.g., '1.0.0')")
	}

	if info.Description == "" {
		v.addWarning(result, "missing_description", "API description is recommended",
			"info.description", "", "Add a description explaining the API purpose")
	}
}

// validateServers validates server configurations
func (v *APIValidator) validateServers(servers []parser.Server, result *APIValidationResult) {
	if len(servers) == 0 {
		v.addWarning(result, "no_servers", "No servers defined", "servers", "",
			"Add at least one server URL")
		return
	}

	for i, server := range servers {
		serverPath := fmt.Sprintf("servers[%d]", i)

		if server.URL == "" {
			v.addError(result, "missing_server_url", "Server URL is required",
				fmt.Sprintf("%s.url", serverPath), "", "Add a valid server URL", "error")
			continue
		}

		// Validate URL format
		if !v.isValidURL(server.URL) {
			v.addError(result, "invalid_server_url", "Invalid server URL format",
				fmt.Sprintf("%s.url", serverPath), "", "Use a valid URL format", "error")
		}

		// Check for HTTPS in production
		if !strings.HasPrefix(server.URL, "https://") && !strings.Contains(server.URL, "localhost") {
			v.addWarning(result, "insecure_server", "Server should use HTTPS in production",
				fmt.Sprintf("%s.url", serverPath), "", "Use HTTPS for production servers")
		}
	}
}

// validateSecuritySchemes validates security scheme configurations
func (v *APIValidator) validateSecuritySchemes(schemes map[string]parser.SecurityScheme, result *APIValidationResult) {
	if len(schemes) == 0 {
		v.addWarning(result, "no_security", "No security schemes defined", "components.securitySchemes", "",
			"Add appropriate security schemes for API protection")
		return
	}

	for schemeName, scheme := range schemes {
		schemePath := fmt.Sprintf("components.securitySchemes.%s", schemeName)

		switch scheme.Type {
		case "oauth2":
			v.validateOAuth2Scheme(&scheme, schemePath, result)
		case "apiKey":
			v.validateAPIKeyScheme(&scheme, schemePath, result)
		case "http":
			v.validateHTTPScheme(&scheme, schemePath, result)
		case "openIdConnect":
			v.validateOpenIDConnectScheme(&scheme, schemePath, result)
		default:
			v.addError(result, "invalid_security_type",
				fmt.Sprintf("Invalid security scheme type: %s", scheme.Type),
				fmt.Sprintf("%s.type", schemePath), "",
				"Use a valid security scheme type (oauth2, apiKey, http, openIdConnect)", "error")
		}
	}
}

// validateOAuth2Scheme validates OAuth2 security schemes
func (v *APIValidator) validateOAuth2Scheme(scheme *parser.SecurityScheme, path string, result *APIValidationResult) {
	if scheme.Flows == nil {
		v.addError(result, "missing_oauth2_flows", "OAuth2 scheme missing flows",
			fmt.Sprintf("%s.flows", path), "", "Add OAuth2 flows configuration", "error")
		return
	}

	// Validate each flow type
	if scheme.Flows.AuthorizationCode != nil {
		v.validateOAuth2Flow(scheme.Flows.AuthorizationCode, fmt.Sprintf("%s.flows.authorizationCode", path), "authorization_code", result)
	}
	if scheme.Flows.ClientCredentials != nil {
		v.validateOAuth2Flow(scheme.Flows.ClientCredentials, fmt.Sprintf("%s.flows.clientCredentials", path), "client_credentials", result)
	}
	if scheme.Flows.Implicit != nil {
		v.validateOAuth2Flow(scheme.Flows.Implicit, fmt.Sprintf("%s.flows.implicit", path), "implicit", result)
	}
	if scheme.Flows.Password != nil {
		v.validateOAuth2Flow(scheme.Flows.Password, fmt.Sprintf("%s.flows.password", path), "password", result)
	}
}

// validateOAuth2Flow validates a single OAuth2 flow
func (v *APIValidator) validateOAuth2Flow(flow *parser.OAuthFlow, path, flowType string, result *APIValidationResult) {
	switch flowType {
	case "authorization_code":
		if flow.AuthorizationURL == "" {
			v.addError(result, "missing_auth_url", "Authorization code flow missing authorizationUrl",
				fmt.Sprintf("%s.authorizationUrl", path), "", "Add authorization URL", "error")
		}
		if flow.TokenURL == "" {
			v.addError(result, "missing_token_url", "Authorization code flow missing tokenUrl",
				fmt.Sprintf("%s.tokenUrl", path), "", "Add token URL", "error")
		}
	case "client_credentials", "password":
		if flow.TokenURL == "" {
			v.addError(result, "missing_token_url", fmt.Sprintf("%s flow missing tokenUrl", flowType),
				fmt.Sprintf("%s.tokenUrl", path), "", "Add token URL", "error")
		}
	case "implicit":
		if flow.AuthorizationURL == "" {
			v.addError(result, "missing_auth_url", "Implicit flow missing authorizationUrl",
				fmt.Sprintf("%s.authorizationUrl", path), "", "Add authorization URL", "error")
		}
	}

	// Validate URLs
	if flow.AuthorizationURL != "" && !v.isValidURL(flow.AuthorizationURL) {
		v.addError(result, "invalid_auth_url", "Invalid authorization URL format",
			fmt.Sprintf("%s.authorizationUrl", path), "", "Use a valid URL format", "error")
	}
	if flow.TokenURL != "" && !v.isValidURL(flow.TokenURL) {
		v.addError(result, "invalid_token_url", "Invalid token URL format",
			fmt.Sprintf("%s.tokenUrl", path), "", "Use a valid URL format", "error")
	}

	// Validate scopes
	if len(flow.Scopes) == 0 {
		v.addWarning(result, "no_scopes", "OAuth2 flow has no scopes defined",
			fmt.Sprintf("%s.scopes", path), "", "Define appropriate scopes for the flow")
	}
}

// validateAPIKeyScheme validates API key security schemes
func (v *APIValidator) validateAPIKeyScheme(scheme *parser.SecurityScheme, path string, result *APIValidationResult) {
	if scheme.Name == "" {
		v.addError(result, "missing_api_key_name", "API key scheme missing name",
			fmt.Sprintf("%s.name", path), "", "Add API key parameter name", "error")
	}

	if scheme.In == "" {
		v.addError(result, "missing_api_key_location", "API key scheme missing 'in' location",
			fmt.Sprintf("%s.in", path), "", "Specify where the API key should be sent (header, query, cookie)", "error")
	} else if scheme.In != "header" && scheme.In != "query" && scheme.In != "cookie" {
		v.addError(result, "invalid_api_key_location", "Invalid API key location",
			fmt.Sprintf("%s.in", path), "", "Use 'header', 'query', or 'cookie'", "error")
	}

	// Security recommendation
	if scheme.In == "query" {
		v.addWarning(result, "insecure_api_key", "API key in query parameter is less secure",
			fmt.Sprintf("%s.in", path), "", "Consider using header for better security")
	}
}

// validateHTTPScheme validates HTTP security schemes
func (v *APIValidator) validateHTTPScheme(scheme *parser.SecurityScheme, path string, result *APIValidationResult) {
	if scheme.Scheme == "" {
		v.addError(result, "missing_http_scheme", "HTTP scheme missing 'scheme' field",
			fmt.Sprintf("%s.scheme", path), "", "Add HTTP authentication scheme (basic, bearer, etc.)", "error")
	}

	validSchemes := []string{"basic", "bearer", "digest", "hoba", "mutual", "negotiate", "oauth", "scram-sha-1", "scram-sha-256", "vapid"}
	isValid := false
	for _, validScheme := range validSchemes {
		if scheme.Scheme == validScheme {
			isValid = true
			break
		}
	}

	if !isValid {
		v.addWarning(result, "unknown_http_scheme", fmt.Sprintf("Unknown HTTP scheme: %s", scheme.Scheme),
			fmt.Sprintf("%s.scheme", path), "", "Use a standard HTTP authentication scheme")
	}

	// Validate bearer format
	if scheme.Scheme == "bearer" && scheme.BearerFormat == "" {
		v.addWarning(result, "missing_bearer_format", "Bearer scheme missing bearerFormat",
			fmt.Sprintf("%s.bearerFormat", path), "", "Add bearer token format (e.g., 'JWT')")
	}
}

// validateOpenIDConnectScheme validates OpenID Connect security schemes
func (v *APIValidator) validateOpenIDConnectScheme(scheme *parser.SecurityScheme, path string, result *APIValidationResult) {
	if scheme.OpenIDConnectURL == "" {
		v.addError(result, "missing_openid_url", "OpenID Connect scheme missing openIdConnectUrl",
			fmt.Sprintf("%s.openIdConnectUrl", path), "", "Add OpenID Connect discovery URL", "error")
	} else if !v.isValidURL(scheme.OpenIDConnectURL) {
		v.addError(result, "invalid_openid_url", "Invalid OpenID Connect URL format",
			fmt.Sprintf("%s.openIdConnectUrl", path), "", "Use a valid URL format", "error")
	}
}

// validatePath validates a single API path
func (v *APIValidator) validatePath(path string, pathItem *parser.PathItem, result *APIValidationResult) {
	// Validate path format
	if !v.isValidAPIPath(path) {
		v.addWarning(result, "invalid_path_format", "Path doesn't follow REST conventions",
			path, "", "Use lowercase, hyphens for word separation, and proper resource hierarchy")
	}

	// Validate operations
	operations := map[string]*parser.Operation{
		"GET":    pathItem.Get,
		"POST":   pathItem.Post,
		"PUT":    pathItem.Put,
		"DELETE": pathItem.Delete,
		"PATCH":  pathItem.Patch,
	}

	for method, operation := range operations {
		if operation != nil {
			result.Summary.TotalEndpoints++
			v.validateOperation(operation, path, method, result)
		}
	}

	// Validate REST compliance for the path
	v.validateRESTCompliance(path, pathItem, result)
}

// validateOperation validates a single API operation
func (v *APIValidator) validateOperation(operation *parser.Operation, path, method string, result *APIValidationResult) {
	opPath := fmt.Sprintf("%s.%s", path, strings.ToLower(method))

	// Validate operation ID
	if operation.OperationID == "" {
		v.addWarning(result, "missing_operation_id", "Operation missing operationId",
			opPath, method, "Add unique operationId for code generation")
	} else if !v.isValidOperationID(operation.OperationID) {
		v.addWarning(result, "invalid_operation_id", "Operation ID doesn't follow naming conventions",
			opPath, method, "Use camelCase for operation IDs")
	}

	// Validate summary and description
	if operation.Summary == "" {
		v.addWarning(result, "missing_summary", "Operation missing summary",
			opPath, method, "Add brief summary describing the operation")
	}

	if operation.Description == "" {
		v.addWarning(result, "missing_description", "Operation missing description",
			opPath, method, "Add detailed description of the operation")
	}

	// Validate parameters
	v.validateParameters(operation.Parameters, opPath, method, result)

	// Validate request body
	if operation.RequestBody != nil {
		v.validateRequestBody(operation.RequestBody, opPath, method, result)
	}

	// Validate responses
	v.validateResponses(operation.Responses, opPath, method, result)

	// Validate tags
	if len(operation.Tags) == 0 {
		v.addWarning(result, "missing_tags", "Operation missing tags",
			opPath, method, "Add tags for better organization")
	}

	// Validate Terraform compatibility
	v.validateTerraformCompatibility(operation, path, method, result)
}

// validateParameters validates operation parameters
func (v *APIValidator) validateParameters(parameters []parser.Parameter, opPath, method string, result *APIValidationResult) {
	paramNames := make(map[string]bool)

	for i, param := range parameters {
		paramPath := fmt.Sprintf("%s.parameters[%d]", opPath, i)

		// Check for duplicate parameter names
		paramKey := fmt.Sprintf("%s:%s", param.In, param.Name)
		if paramNames[paramKey] {
			v.addError(result, "duplicate_parameter",
				fmt.Sprintf("Duplicate parameter '%s' in '%s'", param.Name, param.In),
				paramPath, method, "Remove duplicate parameter", "error")
		}
		paramNames[paramKey] = true

		// Validate parameter name
		if param.Name == "" {
			v.addError(result, "missing_parameter_name", "Parameter missing name",
				paramPath, method, "Add parameter name", "error")
		}

		// Validate parameter location
		validLocations := []string{"query", "header", "path", "cookie"}
		isValidLocation := false
		for _, loc := range validLocations {
			if param.In == loc {
				isValidLocation = true
				break
			}
		}
		if !isValidLocation {
			v.addError(result, "invalid_parameter_location",
				fmt.Sprintf("Invalid parameter location: %s", param.In),
				paramPath, method, "Use 'query', 'header', 'path', or 'cookie'", "error")
		}

		// Path parameters must be required
		if param.In == "path" && !param.Required {
			v.addError(result, "path_param_not_required", "Path parameters must be required",
				paramPath, method, "Set required: true for path parameters", "error")
		}

		// Validate schema
		if param.Schema == nil {
			v.addError(result, "missing_parameter_schema", "Parameter missing schema",
				paramPath, method, "Add schema definition for parameter", "error")
		}
	}
}

// validateRequestBody validates operation request body
func (v *APIValidator) validateRequestBody(requestBody *parser.RequestBody, opPath, method string, result *APIValidationResult) {
	reqPath := fmt.Sprintf("%s.requestBody", opPath)

	if len(requestBody.Content) == 0 {
		v.addError(result, "empty_request_body", "Request body has no content",
			reqPath, method, "Add content types for request body", "error")
		return
	}

	// Check for common content types
	hasJSON := false
	for contentType := range requestBody.Content {
		if strings.Contains(contentType, "application/json") {
			hasJSON = true
			break
		}
	}

	if !hasJSON && (method == "POST" || method == "PUT" || method == "PATCH") {
		v.addWarning(result, "missing_json_content", "Request body missing JSON content type",
			reqPath, method, "Add 'application/json' content type")
	}

	// Validate content schemas
	for contentType, media := range requestBody.Content {
		if media.Schema == nil {
			v.addError(result, "missing_content_schema",
				fmt.Sprintf("Content type '%s' missing schema", contentType),
				fmt.Sprintf("%s.content.%s", reqPath, contentType), method,
				"Add schema definition for content type", "error")
		}
	}
}

// validateResponses validates operation responses
func (v *APIValidator) validateResponses(responses map[string]parser.Response, opPath, method string, result *APIValidationResult) {
	respPath := fmt.Sprintf("%s.responses", opPath)

	if len(responses) == 0 {
		v.addError(result, "no_responses", "Operation has no responses defined",
			respPath, method, "Add response definitions", "error")
		return
	}

	// Check for required response codes
	hasSuccess := false
	hasError := false

	for statusCode := range responses {
		if strings.HasPrefix(statusCode, "2") {
			hasSuccess = true
		}
		if strings.HasPrefix(statusCode, "4") || strings.HasPrefix(statusCode, "5") {
			hasError = true
		}
	}

	if !hasSuccess {
		v.addWarning(result, "missing_success_response", "Operation missing success response",
			respPath, method, "Add 2xx response definition")
	}

	if !hasError {
		v.addWarning(result, "missing_error_response", "Operation missing error response",
			respPath, method, "Add 4xx/5xx error response definitions")
	}

	// Validate specific responses
	for statusCode, response := range responses {
		responseKey := fmt.Sprintf("%s.%s", respPath, statusCode)

		if response.Description == "" {
			v.addWarning(result, "missing_response_description",
				fmt.Sprintf("Response %s missing description", statusCode),
				responseKey, method, "Add response description")
		}

		// Validate content for success responses
		if strings.HasPrefix(statusCode, "2") && len(response.Content) == 0 && method == "GET" {
			v.addWarning(result, "missing_response_content",
				fmt.Sprintf("Success response %s missing content", statusCode),
				responseKey, method, "Add response content definition")
		}
	}
}

// validateRESTCompliance validates REST API compliance
func (v *APIValidator) validateRESTCompliance(path string, pathItem *parser.PathItem, result *APIValidationResult) {
	// Check for proper HTTP method usage
	if pathItem.Get == nil && pathItem.Post == nil && pathItem.Put == nil &&
		pathItem.Delete == nil && pathItem.Patch == nil {
		v.addWarning(result, "no_operations", "Path has no operations defined",
			path, "", "Add appropriate HTTP operations")
		return
	}

	// Validate resource-oriented design
	if !v.isResourceOriented(path) {
		v.addWarning(result, "non_resource_oriented", "Path doesn't follow resource-oriented design",
			path, "", "Use noun-based resource paths")
	}

	// Check for proper CRUD operations
	v.validateCRUDOperations(path, pathItem, result)
}

// validateTerraformCompatibility validates Terraform provider compatibility
func (v *APIValidator) validateTerraformCompatibility(operation *parser.Operation, path, method string, result *APIValidationResult) {
	// Check for idempotency requirements
	if method == "PUT" || method == "DELETE" {
		// These should be idempotent - check for proper design
		if !v.isIdempotentOperation(operation, method) {
			v.addWarning(result, "non_idempotent",
				fmt.Sprintf("%s operation may not be idempotent", method),
				path, method, "Ensure operation is idempotent for Terraform compatibility")
		}
	}

	// Check for proper resource identification
	if method == "GET" || method == "PUT" || method == "DELETE" {
		if !v.hasResourceIdentifier(path) {
			v.addError(result, "missing_resource_id", "Resource operation missing identifier in path",
				path, method, "Add resource identifier to path (e.g., /resources/{id})", "error")
		}
	}

	// Check for pagination support in list operations
	if method == "GET" && v.isListOperation(path) {
		if !v.supportsPagination(operation) {
			v.addWarning(result, "missing_pagination", "List operation should support pagination",
				path, method, "Add pagination parameters (limit, offset, cursor)")
		}
	}
}

// Helper functions

// addError adds an error to the validation result
func (v *APIValidator) addError(result *APIValidationResult, errorType, message, path, method, suggestion, severity string) {
	result.Errors = append(result.Errors, APIError{
		Type:       errorType,
		Message:    message,
		Path:       path,
		Method:     method,
		Suggestion: suggestion,
		Severity:   severity,
	})
}

// addWarning adds a warning to the validation result
func (v *APIValidator) addWarning(result *APIValidationResult, warningType, message, path, method, suggestion string) {
	result.Warnings = append(result.Warnings, APIWarning{
		Type:       warningType,
		Message:    message,
		Path:       path,
		Method:     method,
		Suggestion: suggestion,
	})
}

// isValidVersion checks if version follows semantic versioning
func (v *APIValidator) isValidVersion(version string) bool {
	semverPattern := regexp.MustCompile(`^v?(\d+)\.(\d+)\.(\d+)(?:-([0-9A-Za-z-]+(?:\.[0-9A-Za-z-]+)*))?(?:\+([0-9A-Za-z-]+(?:\.[0-9A-Za-z-]+)*))?$`)
	return semverPattern.MatchString(version)
}

// isValidURL checks if URL is valid
func (v *APIValidator) isValidURL(urlStr string) bool {
	_, err := url.Parse(urlStr)
	return err == nil && (strings.HasPrefix(urlStr, "http://") || strings.HasPrefix(urlStr, "https://"))
}

// isValidAPIPath checks if API path follows conventions
func (v *APIValidator) isValidAPIPath(path string) bool {
	// Basic validation - starts with /, uses lowercase, proper structure
	if !strings.HasPrefix(path, "/") {
		return false
	}

	// Check for proper resource hierarchy
	pathPattern := regexp.MustCompile(`^/[a-z0-9\-_/{}]+$`)
	return pathPattern.MatchString(path)
}

// isValidOperationID checks if operation ID follows naming conventions
func (v *APIValidator) isValidOperationID(operationID string) bool {
	// Should be camelCase
	camelCasePattern := regexp.MustCompile(`^[a-z][a-zA-Z0-9]*$`)
	return camelCasePattern.MatchString(operationID)
}

// isResourceOriented checks if path follows resource-oriented design
func (v *APIValidator) isResourceOriented(path string) bool {
	// Should use nouns, not verbs
	verbPattern := regexp.MustCompile(`/(create|get|update|delete|list|add|remove|set)`)
	return !verbPattern.MatchString(strings.ToLower(path))
}

// validateCRUDOperations validates CRUD operation consistency
func (v *APIValidator) validateCRUDOperations(path string, pathItem *parser.PathItem, result *APIValidationResult) {
	// Collection vs item resource validation
	isCollection := !strings.Contains(path, "{") || strings.HasSuffix(path, "}")

	if isCollection {
		// Collection should have GET (list) and POST (create)
		if pathItem.Get == nil {
			v.addWarning(result, "missing_list_operation", "Collection missing GET operation",
				path, "", "Add GET operation for listing resources")
		}
		if pathItem.Post == nil {
			v.addWarning(result, "missing_create_operation", "Collection missing POST operation",
				path, "", "Add POST operation for creating resources")
		}
		// Collection shouldn't have PUT/DELETE
		if pathItem.Put != nil {
			v.addWarning(result, "invalid_collection_operation", "Collection shouldn't have PUT operation",
				path, "PUT", "Use POST for creating resources in collections")
		}
		if pathItem.Delete != nil {
			v.addWarning(result, "invalid_collection_operation", "Collection shouldn't have DELETE operation",
				path, "DELETE", "Use DELETE on individual resources")
		}
	} else {
		// Item resource should have GET, PUT, DELETE
		if pathItem.Get == nil {
			v.addWarning(result, "missing_read_operation", "Resource missing GET operation",
				path, "", "Add GET operation for reading resource")
		}
		if pathItem.Put == nil && pathItem.Patch == nil {
			v.addWarning(result, "missing_update_operation", "Resource missing PUT/PATCH operation",
				path, "", "Add PUT or PATCH operation for updating resource")
		}
		if pathItem.Delete == nil {
			v.addWarning(result, "missing_delete_operation", "Resource missing DELETE operation",
				path, "", "Add DELETE operation for removing resource")
		}
	}
}

// isIdempotentOperation checks if operation is designed to be idempotent
func (v *APIValidator) isIdempotentOperation(operation *parser.Operation, method string) bool {
	// This is a simplified check - in practice, you'd analyze the operation more deeply
	return method == "GET" || method == "PUT" || method == "DELETE"
}

// hasResourceIdentifier checks if path has resource identifier
func (v *APIValidator) hasResourceIdentifier(path string) bool {
	return strings.Contains(path, "{") && strings.Contains(path, "}")
}

// isListOperation checks if this is a list operation
func (v *APIValidator) isListOperation(path string) bool {
	// List operations typically don't have resource identifiers
	return !strings.Contains(path, "{")
}

// supportsPagination checks if operation supports pagination
func (v *APIValidator) supportsPagination(operation *parser.Operation) bool {
	for _, param := range operation.Parameters {
		if param.Name == "limit" || param.Name == "offset" || param.Name == "cursor" || param.Name == "page" {
			return true
		}
	}
	return false
}

// validateAPIDesign validates overall API design patterns
func (v *APIValidator) validateAPIDesign(spec *parser.APISpec, result *APIValidationResult) {
	// Check for consistent naming patterns
	v.validateNamingConsistency(spec, result)

	// Check for proper versioning strategy
	v.validateVersioningStrategy(spec, result)

	// Check for proper error handling
	v.validateErrorHandling(spec, result)
}

// validateNamingConsistency validates naming consistency across the API
func (v *APIValidator) validateNamingConsistency(spec *parser.APISpec, result *APIValidationResult) {
	// Check for consistent path naming
	pathPatterns := make(map[string]int)

	for path := range spec.Paths {
		// Extract naming pattern (camelCase, snake_case, kebab-case)
		pattern := v.detectNamingPattern(path)
		pathPatterns[pattern]++
	}

	// If multiple patterns are used, warn about inconsistency
	if len(pathPatterns) > 1 {
		v.addWarning(result, "inconsistent_naming", "API uses inconsistent naming patterns",
			"paths", "", "Use consistent naming pattern throughout the API")
	}
}

// validateVersioningStrategy validates API versioning approach
func (v *APIValidator) validateVersioningStrategy(spec *parser.APISpec, result *APIValidationResult) {
	hasVersionInPath := false
	hasVersionInHeader := false

	// Check for version in paths
	for path := range spec.Paths {
		if strings.Contains(path, "/v1/") || strings.Contains(path, "/v2/") {
			hasVersionInPath = true
			break
		}
	}

	// Check for version in parameters (simplified check)
	for _, pathItem := range spec.Paths {
		operations := []*parser.Operation{pathItem.Get, pathItem.Post, pathItem.Put, pathItem.Delete, pathItem.Patch}
		for _, op := range operations {
			if op != nil {
				for _, param := range op.Parameters {
					if param.Name == "version" || param.Name == "api-version" {
						hasVersionInHeader = true
						break
					}
				}
			}
		}
	}

	if !hasVersionInPath && !hasVersionInHeader {
		v.addWarning(result, "missing_versioning", "API lacks clear versioning strategy",
			"", "", "Add version information in path or headers")
	}
}

// validateErrorHandling validates error handling patterns
func (v *APIValidator) validateErrorHandling(spec *parser.APISpec, result *APIValidationResult) {
	hasStandardErrorSchema := false

	// Check for standard error response schema
	if spec.Components.Schemas != nil {
		for schemaName := range spec.Components.Schemas {
			if strings.ToLower(schemaName) == "error" || strings.ToLower(schemaName) == "errorresponse" {
				hasStandardErrorSchema = true
				break
			}
		}
	}

	if !hasStandardErrorSchema {
		v.addWarning(result, "missing_error_schema", "API lacks standard error response schema",
			"components.schemas", "", "Define standard error response schema")
	}
}

// detectNamingPattern detects the naming pattern used in a path
func (v *APIValidator) detectNamingPattern(path string) string {
	if strings.Contains(path, "_") {
		return "snake_case"
	}
	if strings.Contains(path, "-") {
		return "kebab-case"
	}
	// Check for camelCase (simplified)
	camelPattern := regexp.MustCompile(`[a-z][A-Z]`)
	if camelPattern.MatchString(path) {
		return "camelCase"
	}
	return "lowercase"
}

// calculateSummary calculates validation summary
func (v *APIValidator) calculateSummary(result *APIValidationResult) {
	result.Summary.ErrorCount = len(result.Errors)
	result.Summary.WarningCount = len(result.Warnings)
	result.Summary.ValidEndpoints = result.Summary.TotalEndpoints - result.Summary.ErrorCount

	if result.Summary.TotalEndpoints > 0 {
		result.Summary.CompatibilityScore = float64(result.Summary.ValidEndpoints) / float64(result.Summary.TotalEndpoints) * 100
	}

	// Calculate REST compliance score
	restViolations := 0
	for _, warning := range result.Warnings {
		if strings.Contains(warning.Type, "rest") || strings.Contains(warning.Type, "crud") {
			restViolations++
		}
	}

	if result.Summary.TotalEndpoints > 0 {
		result.Summary.RESTCompliance = float64(result.Summary.TotalEndpoints-restViolations) / float64(result.Summary.TotalEndpoints) * 100
	}
}

// generateCompatibilityReport generates detailed compatibility report
func (v *APIValidator) generateCompatibilityReport(spec *parser.APISpec, result *APIValidationResult) {
	// Determine overall compatibility
	result.Compatibility.TerraformCompatible = v.isTerraformCompatible(result)
	result.Compatibility.RESTCompliant = result.Summary.RESTCompliance >= 80.0
	result.Compatibility.SecurityCompliant = v.isSecurityCompliant(result)

	// Generate issues from errors and warnings
	for _, err := range result.Errors {
		result.Compatibility.Issues = append(result.Compatibility.Issues, CompatibilityIssue{
			Type:        err.Type,
			Severity:    err.Severity,
			Description: err.Message,
			Path:        err.Path,
			Solution:    err.Suggestion,
		})
	}

	for _, warning := range result.Warnings {
		result.Compatibility.Issues = append(result.Compatibility.Issues, CompatibilityIssue{
			Type:        warning.Type,
			Severity:    "warning",
			Description: warning.Message,
			Path:        warning.Path,
			Solution:    warning.Suggestion,
		})
	}

	// Generate recommendations
	v.generateRecommendations(spec, result)
}

// isTerraformCompatible checks if API is compatible with Terraform
func (v *APIValidator) isTerraformCompatible(result *APIValidationResult) bool {
	// Check for critical Terraform compatibility issues
	for _, err := range result.Errors {
		if err.Type == "missing_resource_id" || err.Type == "non_idempotent" {
			return false
		}
	}
	return true
}

// isSecurityCompliant checks if API meets security requirements
func (v *APIValidator) isSecurityCompliant(result *APIValidationResult) bool {
	// Check for security-related errors
	for _, err := range result.Errors {
		if strings.Contains(err.Type, "security") || err.Type == "insecure_server" {
			return false
		}
	}
	return true
}

// generateRecommendations generates improvement recommendations
func (v *APIValidator) generateRecommendations(spec *parser.APISpec, result *APIValidationResult) {
	recommendations := []string{}

	// Security recommendations
	if len(spec.Components.SecuritySchemes) == 0 {
		recommendations = append(recommendations, "Implement proper authentication and authorization")
	}

	// Documentation recommendations
	if spec.Info.Description == "" {
		recommendations = append(recommendations, "Add comprehensive API documentation")
	}

	// Versioning recommendations
	hasVersioning := false
	for path := range spec.Paths {
		if strings.Contains(path, "/v") {
			hasVersioning = true
			break
		}
	}
	if !hasVersioning {
		recommendations = append(recommendations, "Implement API versioning strategy")
	}

	// Error handling recommendations
	errorSchemaExists := false
	if spec.Components.Schemas != nil {
		for schemaName := range spec.Components.Schemas {
			if strings.ToLower(schemaName) == "error" {
				errorSchemaExists = true
				break
			}
		}
	}
	if !errorSchemaExists {
		recommendations = append(recommendations, "Define standard error response format")
	}

	// Pagination recommendations
	hasPagination := false
	for _, pathItem := range spec.Paths {
		if pathItem.Get != nil {
			for _, param := range pathItem.Get.Parameters {
				if param.Name == "limit" || param.Name == "offset" {
					hasPagination = true
					break
				}
			}
		}
	}
	if !hasPagination {
		recommendations = append(recommendations, "Implement pagination for list operations")
	}

	result.Compatibility.Recommendations = recommendations
}
