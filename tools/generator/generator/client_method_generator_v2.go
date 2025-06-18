package generator

import (
	"fmt"
	"strings"

	"github.com/mantisec/terraform-provider-umbrella/tools/generator/parser"
)

// ClientMethodGeneratorV2 generates HTTP client methods for API endpoints
type ClientMethodGeneratorV2 struct {
	templateEngine *TemplateEngine
}

// NewClientMethodGeneratorV2 creates a new client method generator
func NewClientMethodGeneratorV2(templateEngine *TemplateEngine) *ClientMethodGeneratorV2 {
	return &ClientMethodGeneratorV2{
		templateEngine: templateEngine,
	}
}

// ClientMethod represents a generated client method
type ClientMethod struct {
	Name           string
	HTTPMethod     string
	Path           string
	HasRequestBody bool
	HasPathParams  bool
	HasQueryParams bool
	RequestType    string
	ResponseType   string
	Description    string
	Parameters     []ClientMethodParameter
	PathParams     []ClientMethodParameter
	QueryParams    []ClientMethodParameter
}

// ClientMethodParameter represents a method parameter for v2 generator
type ClientMethodParameter struct {
	Name        string
	Type        string
	Description string
	Required    bool
	IsPathParam bool
	IsQuery     bool
}

// GenerateClientMethodsForGroup generates client methods for an endpoint group
func (g *ClientMethodGeneratorV2) GenerateClientMethodsForGroup(group parser.EndpointGroup) []ClientMethod {
	var methods []ClientMethod

	for _, endpoint := range group.Endpoints {
		method := g.generateClientMethod(endpoint, group.CanonicalName)
		methods = append(methods, method)
	}

	return methods
}

// generateClientMethod generates a client method for a single endpoint
func (g *ClientMethodGeneratorV2) generateClientMethod(endpoint parser.Endpoint, resourceName string) ClientMethod {
	method := ClientMethod{
		Name:        g.generateMethodName(endpoint, resourceName),
		HTTPMethod:  endpoint.Method,
		Path:        endpoint.Path,
		Description: g.generateMethodDescription(endpoint, resourceName),
		Parameters:  []ClientMethodParameter{},
		PathParams:  []ClientMethodParameter{},
		QueryParams: []ClientMethodParameter{},
	}

	// Extract parameters from the operation
	if endpoint.Operation != nil {
		method.Parameters = g.extractParameters(endpoint.Operation)
		method.PathParams = g.filterPathParams(method.Parameters)
		method.QueryParams = g.filterQueryParams(method.Parameters)
		method.HasPathParams = len(method.PathParams) > 0
		method.HasQueryParams = len(method.QueryParams) > 0
	}

	// Determine if method has request body
	method.HasRequestBody = endpoint.Method == "POST" || endpoint.Method == "PUT" || endpoint.Method == "PATCH"

	// Generate type names
	method.RequestType = g.generateTypeName(resourceName, "Request")
	method.ResponseType = g.generateTypeName(resourceName, "Response")

	return method
}

// generateMethodName generates a method name based on the endpoint
func (g *ClientMethodGeneratorV2) generateMethodName(endpoint parser.Endpoint, resourceName string) string {
	var prefix string

	switch endpoint.CRUDType {
	case "create":
		prefix = "Create"
	case "read":
		prefix = "Get"
	case "update":
		prefix = "Update"
	case "delete":
		prefix = "Delete"
	case "list":
		prefix = "List"
	default:
		prefix = strings.Title(strings.ToLower(endpoint.Method))
	}

	resourceTitle := g.templateEngine.toPascalCase(resourceName)

	// For list operations, use plural
	if endpoint.CRUDType == "list" {
		return prefix + g.pluralize(resourceTitle)
	}

	return prefix + resourceTitle
}

// generateMethodDescription generates a description for the method
func (g *ClientMethodGeneratorV2) generateMethodDescription(endpoint parser.Endpoint, resourceName string) string {
	if endpoint.Operation != nil && endpoint.Operation.Summary != "" {
		return endpoint.Operation.Summary
	}

	switch endpoint.CRUDType {
	case "create":
		return fmt.Sprintf("Creates a new %s", resourceName)
	case "read":
		return fmt.Sprintf("Retrieves a %s by ID", resourceName)
	case "update":
		return fmt.Sprintf("Updates an existing %s", resourceName)
	case "delete":
		return fmt.Sprintf("Deletes a %s by ID", resourceName)
	case "list":
		return fmt.Sprintf("Lists all %s", g.pluralize(resourceName))
	default:
		return fmt.Sprintf("Performs %s operation on %s", endpoint.Method, resourceName)
	}
}

// extractParameters extracts parameters from an operation
func (g *ClientMethodGeneratorV2) extractParameters(operation *parser.Operation) []ClientMethodParameter {
	var params []ClientMethodParameter

	for _, param := range operation.Parameters {
		clientParam := ClientMethodParameter{
			Name:        param.Name,
			Type:        g.convertSchemaTypeToGo(param.Schema),
			Description: param.Description,
			Required:    param.Required,
			IsPathParam: param.In == "path",
			IsQuery:     param.In == "query",
		}
		params = append(params, clientParam)
	}

	return params
}

// filterPathParams filters parameters to only path parameters
func (g *ClientMethodGeneratorV2) filterPathParams(params []ClientMethodParameter) []ClientMethodParameter {
	var pathParams []ClientMethodParameter
	for _, param := range params {
		if param.IsPathParam {
			pathParams = append(pathParams, param)
		}
	}
	return pathParams
}

// filterQueryParams filters parameters to only query parameters
func (g *ClientMethodGeneratorV2) filterQueryParams(params []ClientMethodParameter) []ClientMethodParameter {
	var queryParams []ClientMethodParameter
	for _, param := range params {
		if param.IsQuery {
			queryParams = append(queryParams, param)
		}
	}
	return queryParams
}

// convertSchemaTypeToGo converts OpenAPI schema type to Go type
func (g *ClientMethodGeneratorV2) convertSchemaTypeToGo(schema *parser.Schema) string {
	if schema == nil {
		return "string"
	}

	switch schema.Type {
	case "string":
		return "string"
	case "integer":
		return "int64"
	case "number":
		return "float64"
	case "boolean":
		return "bool"
	case "array":
		return "[]string" // Simplified for now
	case "object":
		return "map[string]interface{}"
	default:
		return "string"
	}
}

// generateTypeName generates a type name for requests/responses
func (g *ClientMethodGeneratorV2) generateTypeName(resourceName, suffix string) string {
	return g.templateEngine.toPascalCase(resourceName) + suffix
}

// pluralize converts a singular noun to plural (basic implementation)
func (g *ClientMethodGeneratorV2) pluralize(word string) string {
	if strings.HasSuffix(word, "y") {
		return strings.TrimSuffix(word, "y") + "ies"
	}
	if strings.HasSuffix(word, "s") || strings.HasSuffix(word, "sh") || strings.HasSuffix(word, "ch") || strings.HasSuffix(word, "x") || strings.HasSuffix(word, "z") {
		return word + "es"
	}
	return word + "s"
}

// GenerateClientMethodCode generates the actual Go code for a client method
func (g *ClientMethodGeneratorV2) GenerateClientMethodCode(method ClientMethod) string {
	var code strings.Builder

	// Method signature
	code.WriteString(fmt.Sprintf("// %s %s\n", method.Name, method.Description))
	code.WriteString(fmt.Sprintf("func (c *apiClient) %s(ctx context.Context", method.Name))

	// Add parameters
	for _, param := range method.Parameters {
		code.WriteString(fmt.Sprintf(", %s %s", param.Name, param.Type))
	}

	// Add request body parameter if needed
	if method.HasRequestBody {
		code.WriteString(fmt.Sprintf(", req *%s", method.RequestType))
	}

	// Return type
	if method.HTTPMethod == "DELETE" {
		code.WriteString(") error {\n")
	} else {
		code.WriteString(fmt.Sprintf(") (*%s, error) {\n", method.ResponseType))
	}

	// Method body
	code.WriteString(g.generateMethodBody(method))

	code.WriteString("}\n\n")

	return code.String()
}

// generateMethodBody generates the method body
func (g *ClientMethodGeneratorV2) generateMethodBody(method ClientMethod) string {
	var body strings.Builder

	// Build URL with path parameters
	body.WriteString(fmt.Sprintf("\turl := c.baseURL + \"%s\"\n", method.Path))

	if method.HasPathParams {
		for _, param := range method.PathParams {
			body.WriteString(fmt.Sprintf("\turl = strings.ReplaceAll(url, \"{%s}\", fmt.Sprintf(\"%%v\", %s))\n", param.Name, param.Name))
		}
	}

	// Add query parameters
	if method.HasQueryParams {
		body.WriteString("\n\tparams := url.Values{}\n")
		for _, param := range method.QueryParams {
			if param.Required {
				body.WriteString(fmt.Sprintf("\tparams.Set(\"%s\", fmt.Sprintf(\"%%v\", %s))\n", param.Name, param.Name))
			} else {
				body.WriteString(fmt.Sprintf("\tif %s != \"\" {\n", param.Name))
				body.WriteString(fmt.Sprintf("\t\tparams.Set(\"%s\", fmt.Sprintf(\"%%v\", %s))\n", param.Name, param.Name))
				body.WriteString("\t}\n")
			}
		}
		body.WriteString("\tif len(params) > 0 {\n")
		body.WriteString("\t\turl += \"?\" + params.Encode()\n")
		body.WriteString("\t}\n")
	}

	// Create HTTP request
	body.WriteString("\n")
	if method.HasRequestBody {
		body.WriteString("\treqBody, err := json.Marshal(req)\n")
		body.WriteString("\tif err != nil {\n")
		if method.HTTPMethod == "DELETE" {
			body.WriteString("\t\treturn fmt.Errorf(\"failed to marshal request: %w\", err)\n")
		} else {
			body.WriteString("\t\treturn nil, fmt.Errorf(\"failed to marshal request: %w\", err)\n")
		}
		body.WriteString("\t}\n")
		body.WriteString(fmt.Sprintf("\treq, err := http.NewRequestWithContext(ctx, \"%s\", url, bytes.NewBuffer(reqBody))\n", method.HTTPMethod))
	} else {
		body.WriteString(fmt.Sprintf("\treq, err := http.NewRequestWithContext(ctx, \"%s\", url, nil)\n", method.HTTPMethod))
	}

	body.WriteString("\tif err != nil {\n")
	if method.HTTPMethod == "DELETE" {
		body.WriteString("\t\treturn fmt.Errorf(\"failed to create request: %w\", err)\n")
	} else {
		body.WriteString("\t\treturn nil, fmt.Errorf(\"failed to create request: %w\", err)\n")
	}
	body.WriteString("\t}\n")

	// Set headers
	if method.HasRequestBody {
		body.WriteString("\treq.Header.Set(\"Content-Type\", \"application/json\")\n")
	}
	body.WriteString("\treq.Header.Set(\"Accept\", \"application/json\")\n")

	// Execute request
	body.WriteString("\n\tresp, err := c.httpClient.Do(req)\n")
	body.WriteString("\tif err != nil {\n")
	if method.HTTPMethod == "DELETE" {
		body.WriteString("\t\treturn fmt.Errorf(\"failed to execute request: %w\", err)\n")
	} else {
		body.WriteString("\t\treturn nil, fmt.Errorf(\"failed to execute request: %w\", err)\n")
	}
	body.WriteString("\t}\n")
	body.WriteString("\tdefer resp.Body.Close()\n")

	// Handle response
	body.WriteString("\n\tif resp.StatusCode < 200 || resp.StatusCode >= 300 {\n")
	body.WriteString("\t\tbody, _ := io.ReadAll(resp.Body)\n")
	if method.HTTPMethod == "DELETE" {
		body.WriteString("\t\treturn fmt.Errorf(\"API error %d: %s\", resp.StatusCode, string(body))\n")
	} else {
		body.WriteString("\t\treturn nil, fmt.Errorf(\"API error %d: %s\", resp.StatusCode, string(body))\n")
	}
	body.WriteString("\t}\n")

	if method.HTTPMethod == "DELETE" {
		body.WriteString("\n\treturn nil\n")
	} else {
		body.WriteString(fmt.Sprintf("\n\tvar result %s\n", method.ResponseType))
		body.WriteString("\tif err := json.NewDecoder(resp.Body).Decode(&result); err != nil {\n")
		body.WriteString("\t\treturn nil, fmt.Errorf(\"failed to decode response: %w\", err)\n")
		body.WriteString("\t}\n")
		body.WriteString("\n\treturn &result, nil\n")
	}

	return body.String()
}
