package generator

import (
	"fmt"
	"strings"

	"github.com/mantisec/terraform-provider-umbrella/tools/generator/config"
	"github.com/mantisec/terraform-provider-umbrella/tools/generator/parser"
)

// ClientMethodGenerator handles generation of actual API client methods
type ClientMethodGenerator struct {
	config         *config.Config
	templateEngine *TemplateEngine
}

// NewClientMethodGenerator creates a new client method generator
func NewClientMethodGenerator(cfg *config.Config, templateEngine *TemplateEngine) *ClientMethodGenerator {
	return &ClientMethodGenerator{
		config:         cfg,
		templateEngine: templateEngine,
	}
}

// ClientMethodData contains data for client method template generation
type ClientMethodData struct {
	MethodName      string
	Description     string
	Path            string
	Method          string
	Parameters      []ClientParameter
	PathParameters  []ClientParameter
	QueryParameters []ClientParameter
	RequestBody     string
	ResponseType    string
	ReturnType      string
	ZeroValue       string
	HasRequestBody  bool
	HasResponse     bool
	CacheTTL        string
	StatusCodes     []int
	ErrorHandling   []ErrorCase
}

// ClientParameter represents a method parameter
type ClientParameter struct {
	Name     string
	GoType   string
	Required bool
	IsLast   bool
	In       string // "path", "query", "header"
}

// ErrorCase represents an error handling case
type ErrorCase struct {
	StatusCode  int
	Description string
	ErrorType   string
}

// GenerateClientMethods generates client methods for all endpoints
func (cmg *ClientMethodGenerator) GenerateClientMethods(endpoints []parser.Endpoint) (string, error) {
	var methods strings.Builder

	// Register the client method template
	if err := cmg.registerClientMethodTemplate(); err != nil {
		return "", fmt.Errorf("failed to register client method template: %w", err)
	}

	for _, endpoint := range endpoints {
		methodData := cmg.prepareClientMethodData(endpoint)

		methodCode, err := cmg.templateEngine.ExecuteTemplate("client_method.go", methodData)
		if err != nil {
			return "", fmt.Errorf("failed to execute client method template for %s: %w", endpoint.Operation.OperationID, err)
		}

		methods.WriteString(methodCode)
		methods.WriteString("\n\n")
	}

	return methods.String(), nil
}

// prepareClientMethodData prepares data for client method generation
func (cmg *ClientMethodGenerator) prepareClientMethodData(endpoint parser.Endpoint) *ClientMethodData {
	data := &ClientMethodData{
		MethodName:  cmg.generateMethodName(endpoint),
		Description: cmg.getDescription(endpoint),
		Path:        endpoint.Path,
		Method:      endpoint.Method,
		StatusCodes: cmg.getExpectedStatusCodes(endpoint),
	}

	// Extract parameters
	data.Parameters = cmg.extractParameters(endpoint.Operation.Parameters)
	data.PathParameters = cmg.filterParametersByLocation(data.Parameters, "path")
	data.QueryParameters = cmg.filterParametersByLocation(data.Parameters, "query")

	// Handle request body
	if endpoint.Operation.RequestBody != nil {
		data.HasRequestBody = true
		data.RequestBody = cmg.generateRequestBodyCode(endpoint.Operation.RequestBody)
	}

	// Handle response
	data.ResponseType, data.ReturnType, data.ZeroValue = cmg.determineResponseType(endpoint.Operation.Responses)
	data.HasResponse = data.ReturnType != "error"

	// Set cache TTL for read operations
	if endpoint.Method == "GET" {
		data.CacheTTL = "5 * time.Minute" // Default cache TTL for GET requests
	}

	// Generate error handling
	data.ErrorHandling = cmg.generateErrorHandling(endpoint.Operation.Responses)

	return data
}

// generateMethodName generates a Go method name from the endpoint
func (cmg *ClientMethodGenerator) generateMethodName(endpoint parser.Endpoint) string {
	if endpoint.Operation.OperationID != "" {
		return cmg.templateEngine.toPascalCase(endpoint.Operation.OperationID)
	}

	// Generate from path and method
	pathParts := strings.Split(strings.Trim(endpoint.Path, "/"), "/")
	var nameParts []string

	// Add method prefix
	switch endpoint.Method {
	case "GET":
		if strings.Contains(endpoint.Path, "{") {
			nameParts = append(nameParts, "Get")
		} else {
			nameParts = append(nameParts, "List")
		}
	case "POST":
		nameParts = append(nameParts, "Create")
	case "PUT", "PATCH":
		nameParts = append(nameParts, "Update")
	case "DELETE":
		nameParts = append(nameParts, "Delete")
	}

	// Add path parts (skip parameters and common words)
	for _, part := range pathParts {
		if !strings.Contains(part, "{") && part != "v1" && part != "v2" && part != "api" {
			nameParts = append(nameParts, cmg.templateEngine.toPascalCase(part))
		}
	}

	return strings.Join(nameParts, "")
}

// getDescription gets the operation description
func (cmg *ClientMethodGenerator) getDescription(endpoint parser.Endpoint) string {
	if endpoint.Operation.Description != "" {
		return endpoint.Operation.Description
	}
	if endpoint.Operation.Summary != "" {
		return endpoint.Operation.Summary
	}
	return fmt.Sprintf("%s %s", endpoint.Method, endpoint.Path)
}

// extractParameters extracts and converts OpenAPI parameters to client parameters
func (cmg *ClientMethodGenerator) extractParameters(params []parser.Parameter) []ClientParameter {
	var clientParams []ClientParameter

	for i, param := range params {
		clientParam := ClientParameter{
			Name:     cmg.templateEngine.toCamelCase(param.Name),
			GoType:   cmg.parameterToGoType(param),
			Required: param.Required,
			IsLast:   i == len(params)-1,
			In:       param.In,
		}
		clientParams = append(clientParams, clientParam)
	}

	return clientParams
}

// filterParametersByLocation filters parameters by their location (path, query, etc.)
func (cmg *ClientMethodGenerator) filterParametersByLocation(params []ClientParameter, location string) []ClientParameter {
	var filtered []ClientParameter
	for _, param := range params {
		if param.In == location {
			filtered = append(filtered, param)
		}
	}
	return filtered
}

// parameterToGoType converts an OpenAPI parameter to a Go type
func (cmg *ClientMethodGenerator) parameterToGoType(param parser.Parameter) string {
	if param.Schema == nil {
		return "string"
	}
	return cmg.templateEngine.schemaToGoType(param.Schema)
}

// generateRequestBodyCode generates code for handling request body
func (cmg *ClientMethodGenerator) generateRequestBodyCode(requestBody *parser.RequestBody) string {
	// For now, assume JSON request body
	return "payload"
}

// determineResponseType determines the response type and return type
func (cmg *ClientMethodGenerator) determineResponseType(responses map[string]parser.Response) (string, string, string) {
	// Look for successful response (200, 201, etc.)
	for statusCode, response := range responses {
		if strings.HasPrefix(statusCode, "2") {
			if len(response.Content) > 0 {
				for contentType, mediaType := range response.Content {
					if strings.Contains(contentType, "json") && mediaType.Schema != nil {
						goType := cmg.templateEngine.schemaToGoType(mediaType.Schema)
						return goType, goType, cmg.getZeroValue(goType)
					}
				}
			}
		}
	}

	// Default to error return type for operations without response body
	return "", "error", "nil"
}

// getZeroValue returns the zero value for a Go type
func (cmg *ClientMethodGenerator) getZeroValue(goType string) string {
	switch goType {
	case "string":
		return "\"\""
	case "int", "int32", "int64":
		return "0"
	case "float32", "float64":
		return "0.0"
	case "bool":
		return "false"
	default:
		if strings.HasPrefix(goType, "[]") {
			return "nil"
		}
		if strings.HasPrefix(goType, "map[") {
			return "nil"
		}
		return "nil"
	}
}

// getExpectedStatusCodes returns the expected HTTP status codes for the operation
func (cmg *ClientMethodGenerator) getExpectedStatusCodes(endpoint parser.Endpoint) []int {
	var codes []int

	for statusCode := range endpoint.Operation.Responses {
		if strings.HasPrefix(statusCode, "2") {
			switch statusCode {
			case "200":
				codes = append(codes, 200)
			case "201":
				codes = append(codes, 201)
			case "202":
				codes = append(codes, 202)
			case "204":
				codes = append(codes, 204)
			}
		}
	}

	// Default status codes based on method
	if len(codes) == 0 {
		switch endpoint.Method {
		case "GET":
			codes = []int{200}
		case "POST":
			codes = []int{201}
		case "PUT", "PATCH":
			codes = []int{200}
		case "DELETE":
			codes = []int{204}
		default:
			codes = []int{200}
		}
	}

	return codes
}

// generateErrorHandling generates error handling cases
func (cmg *ClientMethodGenerator) generateErrorHandling(responses map[string]parser.Response) []ErrorCase {
	var errorCases []ErrorCase

	for statusCode, response := range responses {
		if strings.HasPrefix(statusCode, "4") || strings.HasPrefix(statusCode, "5") {
			errorCase := ErrorCase{
				Description: response.Description,
				ErrorType:   "APIError",
			}

			switch statusCode {
			case "400":
				errorCase.StatusCode = 400
			case "401":
				errorCase.StatusCode = 401
			case "403":
				errorCase.StatusCode = 403
			case "404":
				errorCase.StatusCode = 404
			case "429":
				errorCase.StatusCode = 429
			case "500":
				errorCase.StatusCode = 500
			}

			if errorCase.StatusCode > 0 {
				errorCases = append(errorCases, errorCase)
			}
		}
	}

	return errorCases
}

// registerClientMethodTemplate registers the enhanced client method template
func (cmg *ClientMethodGenerator) registerClientMethodTemplate() error {
	template := `// {{.MethodName}} {{.Description}}
func (c *GeneratedClient) {{.MethodName}}(ctx context.Context{{if .Parameters}}{{range .Parameters}}, {{.Name}} {{.GoType}}{{end}}{{end}}) ({{.ReturnType}}, error) {
	{{- if .PathParameters}}
	path := fmt.Sprintf("{{.Path}}", {{range .PathParameters}}{{.Name}}{{if not .IsLast}}, {{end}}{{end}})
	{{- else}}
	path := "{{.Path}}"
	{{- end}}

	{{- if .QueryParameters}}
	// Add query parameters
	queryParams := make(map[string]string)
	{{- range .QueryParameters}}
	{{- if .Required}}
	queryParams["{{.Name}}"] = fmt.Sprintf("%v", {{.Name}})
	{{- else}}
	if {{.Name}} != {{if eq .GoType "string"}}""{{else if eq .GoType "int"}}0{{else}}nil{{end}} {
		queryParams["{{.Name}}"] = fmt.Sprintf("%v", {{.Name}})
	}
	{{- end}}
	{{- end}}

	if len(queryParams) > 0 {
		var params []string
		for k, v := range queryParams {
			params = append(params, fmt.Sprintf("%s=%s", k, v))
		}
		path += "?" + strings.Join(params, "&")
	}
	{{- end}}

	{{- if .HasRequestBody}}
	payload := {{.RequestBody}}
	body, err := json.Marshal(payload)
	if err != nil {
		return {{.ZeroValue}}, fmt.Errorf("failed to marshal request body: %w", err)
	}
	{{- else}}
	var body []byte
	{{- end}}

	// Log the request
	c.logRequest("{{.Method}}", path, body)

	{{- if eq .Method "GET" }}
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "{{.Method}}", path, body, {{.CacheTTL}})
	{{- else}}
	// Clear cache for write operations
	c.clearCacheForPath(path)
	resp, err := c.do(ctx, "{{.Method}}", path, body)
	{{- end}}
	if err != nil {
		return {{.ZeroValue}}, err
	}
	defer resp.Body.Close()

	// Log the response
	c.logResponse(resp)

	// Validate response status
	if err := c.validateResponse(resp, {{range $i, $code := .StatusCodes}}{{if $i}}, {{end}}{{$code}}{{end}}); err != nil {
		return {{.ZeroValue}}, err
	}

	{{- if .HasResponse}}
	var result {{.ResponseType}}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return {{.ZeroValue}}, fmt.Errorf("failed to decode response: %w", err)
	}

	return result, nil
	{{- else}}
	return nil
	{{- end}}
}`

	return cmg.templateEngine.RegisterTemplate("client_method.go", template)
}
