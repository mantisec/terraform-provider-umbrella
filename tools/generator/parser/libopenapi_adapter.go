package parser

import (
	"fmt"

	"github.com/mantisec/terraform-provider-umbrella/tools/generator/openapi"
	"github.com/pb33f/libopenapi/datamodel/high/base"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
	"gopkg.in/yaml.v3"
)

// LibOpenAPIAdapter converts libopenapi v3.Document to our internal APISpec
type LibOpenAPIAdapter struct {
	loader *openapi.Loader
}

// NewLibOpenAPIAdapter creates a new adapter
func NewLibOpenAPIAdapter() *LibOpenAPIAdapter {
	return &LibOpenAPIAdapter{
		loader: openapi.NewLoader(),
	}
}

// ParseFile parses an OpenAPI file using libopenapi and converts to APISpec
func (a *LibOpenAPIAdapter) ParseFile(path string) (*APISpec, error) {
	// Load using libopenapi
	doc, err := a.loader.LoadFromFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load document: %w", err)
	}

	// Convert to our APISpec
	return a.ConvertDocument(doc)
}

// ParseBytes parses OpenAPI bytes using libopenapi and converts to APISpec
func (a *LibOpenAPIAdapter) ParseBytes(data []byte) (*APISpec, error) {
	// Load using libopenapi
	doc, err := a.loader.LoadFromBytes(data)
	if err != nil {
		return nil, fmt.Errorf("failed to load document: %w", err)
	}

	// Convert to our APISpec
	return a.ConvertDocument(doc)
}

// ConvertDocument converts a libopenapi v3.Document to our APISpec
func (a *LibOpenAPIAdapter) ConvertDocument(doc *v3.Document) (*APISpec, error) {
	spec := &APISpec{
		Info:       a.convertInfo(doc.Info),
		Servers:    a.convertServers(doc.Servers),
		Paths:      make(map[string]PathItem),
		Components: a.convertComponents(doc.Components),
		Security:   a.convertSecurity(doc.Security),
	}

	// Convert paths
	if doc.Paths != nil && doc.Paths.PathItems != nil {
		for pair := doc.Paths.PathItems.First(); pair != nil; pair = pair.Next() {
			pathKey := pair.Key()
			pathItem := pair.Value()
			spec.Paths[pathKey] = a.convertPathItem(pathItem)
		}
	}

	return spec, nil
}

// convertInfo converts base.Info to APIInfo
func (a *LibOpenAPIAdapter) convertInfo(info *base.Info) APIInfo {
	if info == nil {
		return APIInfo{}
	}

	return APIInfo{
		Title:       info.Title,
		Version:     info.Version,
		Description: info.Description,
	}
}

// convertServers converts v3.Server slice to Server slice
func (a *LibOpenAPIAdapter) convertServers(servers []*v3.Server) []Server {
	var result []Server
	for _, server := range servers {
		if server == nil {
			continue
		}

		s := Server{
			URL:         server.URL,
			Description: server.Description,
			Variables:   make(map[string]ServerVariable),
		}

		// Convert server variables
		if server.Variables != nil {
			for pair := server.Variables.First(); pair != nil; pair = pair.Next() {
				key := pair.Key()
				variable := pair.Value()
				if variable != nil {
					s.Variables[key] = ServerVariable{
						Default:     variable.Default,
						Description: variable.Description,
						Enum:        variable.Enum,
					}
				}
			}
		}

		result = append(result, s)
	}
	return result
}

// convertPathItem converts v3.PathItem to PathItem
func (a *LibOpenAPIAdapter) convertPathItem(pathItem *v3.PathItem) PathItem {
	if pathItem == nil {
		return PathItem{}
	}

	return PathItem{
		Get:    a.convertOperation(pathItem.Get),
		Post:   a.convertOperation(pathItem.Post),
		Put:    a.convertOperation(pathItem.Put),
		Delete: a.convertOperation(pathItem.Delete),
		Patch:  a.convertOperation(pathItem.Patch),
	}
}

// convertOperation converts v3.Operation to Operation
func (a *LibOpenAPIAdapter) convertOperation(op *v3.Operation) *Operation {
	if op == nil {
		return nil
	}

	operation := &Operation{
		OperationID: op.OperationId,
		Summary:     op.Summary,
		Description: op.Description,
		Tags:        op.Tags,
		Parameters:  a.convertParameters(op.Parameters),
		RequestBody: a.convertRequestBody(op.RequestBody),
		Responses:   a.convertResponses(op.Responses),
		Security:    a.convertSecurity(op.Security),
	}

	return operation
}

// convertParameters converts v3.Parameter slice to Parameter slice
func (a *LibOpenAPIAdapter) convertParameters(params []*v3.Parameter) []Parameter {
	var result []Parameter
	for _, param := range params {
		if param == nil {
			continue
		}

		p := Parameter{
			Name:        param.Name,
			In:          param.In,
			Description: param.Description,
			Required:    param.Required != nil && *param.Required,
			Schema:      a.convertSchemaProxy(param.Schema),
		}

		result = append(result, p)
	}
	return result
}

// convertRequestBody converts v3.RequestBody to RequestBody
func (a *LibOpenAPIAdapter) convertRequestBody(rb *v3.RequestBody) *RequestBody {
	if rb == nil {
		return nil
	}

	requestBody := &RequestBody{
		Description: rb.Description,
		Required:    rb.Required != nil && *rb.Required,
		Content:     make(map[string]MediaType),
	}

	// Convert content
	if rb.Content != nil {
		for pair := rb.Content.First(); pair != nil; pair = pair.Next() {
			contentType := pair.Key()
			mediaType := pair.Value()
			if mediaType != nil {
				requestBody.Content[contentType] = a.convertMediaType(mediaType)
			}
		}
	}

	return requestBody
}

// convertResponses converts v3.Responses to map[string]Response
func (a *LibOpenAPIAdapter) convertResponses(responses *v3.Responses) map[string]Response {
	result := make(map[string]Response)
	if responses == nil || responses.Codes == nil {
		return result
	}

	for pair := responses.Codes.First(); pair != nil; pair = pair.Next() {
		code := pair.Key()
		response := pair.Value()
		if response != nil {
			result[code] = a.convertResponse(response)
		}
	}

	return result
}

// convertResponse converts v3.Response to Response
func (a *LibOpenAPIAdapter) convertResponse(resp *v3.Response) Response {
	if resp == nil {
		return Response{}
	}

	response := Response{
		Description: resp.Description,
		Content:     make(map[string]MediaType),
		Headers:     make(map[string]Header),
	}

	// Convert content
	if resp.Content != nil {
		for pair := resp.Content.First(); pair != nil; pair = pair.Next() {
			contentType := pair.Key()
			mediaType := pair.Value()
			if mediaType != nil {
				response.Content[contentType] = a.convertMediaType(mediaType)
			}
		}
	}

	// Convert headers
	if resp.Headers != nil {
		for pair := resp.Headers.First(); pair != nil; pair = pair.Next() {
			headerName := pair.Key()
			header := pair.Value()
			if header != nil {
				response.Headers[headerName] = Header{
					Description: header.Description,
					Schema:      a.convertSchemaProxy(header.Schema),
				}
			}
		}
	}

	return response
}

// convertMediaType converts v3.MediaType to MediaType
func (a *LibOpenAPIAdapter) convertMediaType(mt *v3.MediaType) MediaType {
	if mt == nil {
		return MediaType{}
	}

	mediaType := MediaType{
		Schema:   a.convertSchemaProxy(mt.Schema),
		Example:  mt.Example,
		Examples: make(map[string]Example),
	}

	// Convert examples
	if mt.Examples != nil {
		for pair := mt.Examples.First(); pair != nil; pair = pair.Next() {
			exampleName := pair.Key()
			example := pair.Value()
			if example != nil {
				mediaType.Examples[exampleName] = Example{
					Summary:     example.Summary,
					Description: example.Description,
					Value:       example.Value,
				}
			}
		}
	}

	return mediaType
}

// convertComponents converts v3.Components to Components
func (a *LibOpenAPIAdapter) convertComponents(comp *v3.Components) Components {
	components := Components{
		Schemas:         make(map[string]*Schema),
		Responses:       make(map[string]Response),
		Parameters:      make(map[string]Parameter),
		Examples:        make(map[string]Example),
		RequestBodies:   make(map[string]RequestBody),
		Headers:         make(map[string]Header),
		SecuritySchemes: make(map[string]SecurityScheme),
	}

	if comp == nil {
		return components
	}

	// Convert schemas
	if comp.Schemas != nil {
		for pair := comp.Schemas.First(); pair != nil; pair = pair.Next() {
			schemaName := pair.Key()
			schemaProxy := pair.Value()
			if schemaProxy != nil {
				components.Schemas[schemaName] = a.convertSchemaProxy(schemaProxy)
			}
		}
	}

	// Convert responses
	if comp.Responses != nil {
		for pair := comp.Responses.First(); pair != nil; pair = pair.Next() {
			responseName := pair.Key()
			response := pair.Value()
			if response != nil {
				components.Responses[responseName] = a.convertResponse(response)
			}
		}
	}

	// Convert parameters
	if comp.Parameters != nil {
		for pair := comp.Parameters.First(); pair != nil; pair = pair.Next() {
			paramName := pair.Key()
			param := pair.Value()
			if param != nil {
				components.Parameters[paramName] = Parameter{
					Name:        param.Name,
					In:          param.In,
					Description: param.Description,
					Required:    param.Required != nil && *param.Required,
					Schema:      a.convertSchemaProxy(param.Schema),
				}
			}
		}
	}

	return components
}

// convertSchemaProxy converts base.SchemaProxy to Schema
func (a *LibOpenAPIAdapter) convertSchemaProxy(schemaProxy *base.SchemaProxy) *Schema {
	if schemaProxy == nil {
		return nil
	}

	// Build the schema - this resolves all references
	schema := schemaProxy.Schema()
	if schema == nil {
		return nil
	}

	return a.convertSchema(schema)
}

// convertSchema converts base.Schema to Schema
func (a *LibOpenAPIAdapter) convertSchema(schema *base.Schema) *Schema {
	if schema == nil {
		return nil
	}

	result := &Schema{
		Type:        a.convertSchemaType(schema.Type),
		Format:      schema.Format,
		Description: schema.Description,
		Required:    schema.Required,
		Enum:        a.convertEnum(schema.Enum),
		Default:     schema.Default,
		Example:     schema.Example,
		Properties:  make(map[string]*Schema),
	}

	// Convert properties
	if schema.Properties != nil {
		for pair := schema.Properties.First(); pair != nil; pair = pair.Next() {
			propName := pair.Key()
			propProxy := pair.Value()
			if propProxy != nil {
				result.Properties[propName] = a.convertSchemaProxy(propProxy)
			}
		}
	}

	// Convert items for arrays
	if schema.Items != nil && schema.Items.IsA() {
		result.Items = a.convertSchemaProxy(schema.Items.A)
	}

	// Convert composition schemas
	if schema.AllOf != nil {
		for _, schemaProxy := range schema.AllOf {
			result.AllOf = append(result.AllOf, a.convertSchemaProxy(schemaProxy))
		}
	}

	if schema.OneOf != nil {
		for _, schemaProxy := range schema.OneOf {
			result.OneOf = append(result.OneOf, a.convertSchemaProxy(schemaProxy))
		}
	}

	if schema.AnyOf != nil {
		for _, schemaProxy := range schema.AnyOf {
			result.AnyOf = append(result.AnyOf, a.convertSchemaProxy(schemaProxy))
		}
	}

	// Handle additional properties
	if schema.AdditionalProperties != nil {
		if schema.AdditionalProperties.IsA() {
			result.AdditionalProperties = a.convertSchemaProxy(schema.AdditionalProperties.A)
		} else if schema.AdditionalProperties.IsB() {
			result.AdditionalProperties = schema.AdditionalProperties.B
		}
	}

	return result
}

// convertSchemaType converts []string type to string (takes first type)
func (a *LibOpenAPIAdapter) convertSchemaType(types []string) string {
	if len(types) == 0 {
		return ""
	}
	return types[0]
}

// convertEnum converts []*yaml.Node to []interface{}
func (a *LibOpenAPIAdapter) convertEnum(enumNodes []*yaml.Node) []interface{} {
	if enumNodes == nil {
		return nil
	}

	var result []interface{}
	for _, node := range enumNodes {
		if node != nil {
			var value interface{}
			if err := node.Decode(&value); err == nil {
				result = append(result, value)
			}
		}
	}
	return result
}

// convertSecurity converts base.SecurityRequirement slice to SecurityRequirement slice
func (a *LibOpenAPIAdapter) convertSecurity(security []*base.SecurityRequirement) []SecurityRequirement {
	var result []SecurityRequirement
	for _, secReq := range security {
		if secReq == nil || secReq.Requirements == nil {
			continue
		}

		requirement := make(SecurityRequirement)
		for pair := secReq.Requirements.First(); pair != nil; pair = pair.Next() {
			key := pair.Key()
			value := pair.Value()
			requirement[key] = value
		}
		result = append(result, requirement)
	}
	return result
}
