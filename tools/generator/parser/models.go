package parser

// APISpec represents a parsed OpenAPI specification
type APISpec struct {
	Info       APIInfo               `json:"info"`
	Servers    []Server              `json:"servers"`
	Paths      map[string]PathItem   `json:"paths"`
	Components Components            `json:"components"`
	Security   []SecurityRequirement `json:"security"`
}

// APIInfo contains basic information about the API
type APIInfo struct {
	Title       string `json:"title"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

// Server represents an API server
type Server struct {
	URL         string                    `json:"url"`
	Description string                    `json:"description"`
	Variables   map[string]ServerVariable `json:"variables"`
}

// ServerVariable represents a server URL variable
type ServerVariable struct {
	Default     string   `json:"default"`
	Description string   `json:"description"`
	Enum        []string `json:"enum"`
}

// PathItem represents operations available on a single path
type PathItem struct {
	Get    *Operation `json:"get"`
	Post   *Operation `json:"post"`
	Put    *Operation `json:"put"`
	Delete *Operation `json:"delete"`
	Patch  *Operation `json:"patch"`
}

// Operation represents a single API operation
type Operation struct {
	OperationID string                `json:"operationId"`
	Summary     string                `json:"summary"`
	Description string                `json:"description"`
	Tags        []string              `json:"tags"`
	Parameters  []Parameter           `json:"parameters"`
	RequestBody *RequestBody          `json:"requestBody"`
	Responses   map[string]Response   `json:"responses"`
	Security    []SecurityRequirement `json:"security"`
}

// Parameter represents an operation parameter
type Parameter struct {
	Name        string  `json:"name"`
	In          string  `json:"in"` // query, header, path, cookie
	Description string  `json:"description"`
	Required    bool    `json:"required"`
	Schema      *Schema `json:"schema"`
	Ref         string  `json:"$ref"`
}

// RequestBody represents a request body
type RequestBody struct {
	Description string               `json:"description"`
	Content     map[string]MediaType `json:"content"`
	Required    bool                 `json:"required"`
}

// Response represents an API response
type Response struct {
	Description string               `json:"description"`
	Content     map[string]MediaType `json:"content"`
	Headers     map[string]Header    `json:"headers"`
}

// MediaType represents a media type object
type MediaType struct {
	Schema   *Schema            `json:"schema"`
	Example  interface{}        `json:"example"`
	Examples map[string]Example `json:"examples"`
}

// Header represents a response header
type Header struct {
	Description string  `json:"description"`
	Schema      *Schema `json:"schema"`
}

// Example represents an example value
type Example struct {
	Summary     string      `json:"summary"`
	Description string      `json:"description"`
	Value       interface{} `json:"value"`
}

// Schema represents a JSON Schema
type Schema struct {
	Type                 string             `json:"type"`
	Format               string             `json:"format"`
	Description          string             `json:"description"`
	Properties           map[string]*Schema `json:"properties"`
	Items                *Schema            `json:"items"`
	Required             []string           `json:"required"`
	Enum                 []interface{}      `json:"enum"`
	Default              interface{}        `json:"default"`
	Example              interface{}        `json:"example"`
	Ref                  string             `json:"$ref"`
	AllOf                []*Schema          `json:"allOf"`
	OneOf                []*Schema          `json:"oneOf"`
	AnyOf                []*Schema          `json:"anyOf"`
	AdditionalProperties interface{}        `json:"additionalProperties"`
}

// Components holds reusable objects for different aspects of the OAS
type Components struct {
	Schemas         map[string]*Schema        `json:"schemas"`
	Responses       map[string]Response       `json:"responses"`
	Parameters      map[string]Parameter      `json:"parameters"`
	Examples        map[string]Example        `json:"examples"`
	RequestBodies   map[string]RequestBody    `json:"requestBodies"`
	Headers         map[string]Header         `json:"headers"`
	SecuritySchemes map[string]SecurityScheme `json:"securitySchemes"`
}

// SecurityScheme represents a security scheme
type SecurityScheme struct {
	Type             string      `json:"type"`
	Description      string      `json:"description"`
	Name             string      `json:"name"`
	In               string      `json:"in"`
	Scheme           string      `json:"scheme"`
	BearerFormat     string      `json:"bearerFormat"`
	Flows            *OAuthFlows `json:"flows"`
	OpenIDConnectURL string      `json:"openIdConnectUrl"`
}

// OAuthFlows represents OAuth2 flows
type OAuthFlows struct {
	Implicit          *OAuthFlow `json:"implicit"`
	Password          *OAuthFlow `json:"password"`
	ClientCredentials *OAuthFlow `json:"clientCredentials"`
	AuthorizationCode *OAuthFlow `json:"authorizationCode"`
}

// OAuthFlow represents a single OAuth2 flow
type OAuthFlow struct {
	AuthorizationURL string            `json:"authorizationUrl"`
	TokenURL         string            `json:"tokenUrl"`
	RefreshURL       string            `json:"refreshUrl"`
	Scopes           map[string]string `json:"scopes"`
}

// SecurityRequirement represents a security requirement
type SecurityRequirement map[string][]string

// Endpoint represents a processed API endpoint for code generation
type Endpoint struct {
	Path         string
	Method       string
	Operation    *Operation
	ResourceType string // "resource" or "data_source"
	ResourceName string
	CRUDType     string // "create", "read", "update", "delete", "list"
}

// ResourceClassification represents how an endpoint should be classified
type ResourceClassification struct {
	IsResource   bool
	IsDataSource bool
	ResourceName string
	CRUDType     string
}
