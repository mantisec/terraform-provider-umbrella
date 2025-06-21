package openapi

// PrimitiveType enumerates the scalar kinds we support.
const (
	PrimString = "string"
	PrimInt    = "int"
	PrimFloat  = "float"
	PrimBool   = "bool"
)

// FieldKind represents the overall attribute flavour.
type FieldKind string

const (
	KindPrimitive FieldKind = "primitive" // uses Prim* for PrimType
	KindList      FieldKind = "list"      // element is primitive
	KindMap       FieldKind = "map"       // element is primitive
	KindObject    FieldKind = "object"    // free‑form object, currently map[string]string
)

// ResourceDef captures the information we need to emit Terraform code for
// a single REST resource discovered in an OpenAPI spec.
type ResourceDef struct {
	Name           string      // singular resource name, e.g. "internal_domain"
	CollectionPath string      // e.g. "/deployments/v2/internaldomains"
	ItemPath       string      // e.g. "/deployments/v2/internaldomains/{id}"
	CreatePayload  *SchemaSpec // schema for POST body (nil if create not allowed)
	ReadSchema     *SchemaSpec // schema for GET 200 response
	UpdatePayload  *SchemaSpec // schema for PUT/PATCH body (nil if immutable)
	HasDelete      bool        // true if DELETE operation exists
}

type SchemaSpec struct {
	Fields []FieldDef
}

type FieldDef struct {
	Name     string // snake_case
	JSONName string
	Kind     FieldKind
	PrimType string      // for primitive/list/map
	Nested   *SchemaSpec // for KindObject or list<object>
	EnumVals []string    // allowed values
	Required bool
	ReadOnly bool
}
