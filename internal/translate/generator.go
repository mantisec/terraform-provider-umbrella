package translate

import (
	"bytes"
	"embed"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/mantisec/terraform-provider-umbrella/internal/openapi"
)

//go:embed ../templates/*.tpl
var tplFS embed.FS

type fieldTmpl struct {
	Name   string // snake
	GoName string // Camel
	GoType string // types.String etc.
	Decl   string // attribute declaration code
}

type resTmpl struct {
	Factory      string
	Receiver     string
	Model        string
	ResourceName string
	Fields       []fieldTmpl
	AttrDecls    []fieldTmpl
}

// Helper: map FieldDef to Go struct field type in the model.
func goTypeForField(f openapi.FieldDef) string {
	switch f.Kind {
	case openapi.KindPrimitive:
		return primModelType(f.PrimType)
	case openapi.KindList:
		return "types.List" // List[primitive] stored generically in state
	case openapi.KindMap:
		return "types.Map"
	default:
		return "types.String" // fallback for complex objects
	}
}

// Terraform attribute declaration string.
func attributeDecl(f openapi.FieldDef) string {
	reqOpt := "Optional: true"
	if f.Required {
		reqOpt = "Required: true"
	}

	switch f.Kind {
	case openapi.KindPrimitive:
		attrType := primAttrType(f.PrimType)
		return fmt.Sprintf("resource.%s{%s}", attrType, reqOpt)
	case openapi.KindList:
		elemType := primTypeToken(f.PrimType)
		return fmt.Sprintf("resource.ListAttribute{ElementType: %s, %s}", elemType, reqOpt)
	case openapi.KindMap:
		elemType := primTypeToken(f.PrimType)
		return fmt.Sprintf("resource.MapAttribute{ElementType: %s, %s}", elemType, reqOpt)
	default:
		return fmt.Sprintf("resource.StringAttribute{%s}", reqOpt)
	}
}

func primAttrType(p string) string {
	switch p {
	case openapi.PrimInt:
		return "Int64Attribute"
	case openapi.PrimFloat:
		return "Float64Attribute"
	case openapi.PrimBool:
		return "BoolAttribute"
	default:
		return "StringAttribute"
	}
}

func primTypeToken(p string) string {
	switch p {
	case openapi.PrimInt:
		return "types.Int64Type"
	case openapi.PrimFloat:
		return "types.Float64Type"
	case openapi.PrimBool:
		return "types.BoolType"
	default:
		return "types.StringType"
	}
}

func toCamel(s string) string {
	parts := strings.Split(s, "_")
	for i, p := range parts {
		parts[i] = strings.Title(p)
	}
	return strings.Join(parts, "")
}

func GenerateAll(rd openapi.ResourceDef, outDir, docsDir string) (map[string][]byte, error) {
	files := make(map[string][]byte)

	// helper to render any template
	render := func(tplName string, data any) ([]byte, error) {
		tpl, err := template.ParseFS(tplFS, "../provider/templates/"+tplName)
		if err != nil {
			return nil, err
		}
		var b bytes.Buffer
		if err = tpl.Execute(&b, data); err != nil {
			return nil, err
		}
		return b.Bytes(), nil
	}

	ctx := assemble(rd) // template context (unchanged)
	if b, err := render("resource.go.tpl", ctx); err == nil {
		files[filepath.Join(outDir, "resource_"+rd.Name+"_gen.go")] = b
	} else {
		return nil, err
	}

	if b, err := render("data_source.go.tpl", ctx); err == nil {
		files[filepath.Join(outDir, "datasource_"+rd.Name+"_gen.go")] = b
	} else {
		return nil, err
	}

	if b, err := render("resource.md.tpl", ctx); err == nil {
		files[filepath.Join(docsDir, "resources", rd.Name+".md")] = b
	} else {
		return nil, err
	}

	if b, err := render("data_source.md.tpl", ctx); err == nil {
		files[filepath.Join(docsDir, "data-sources", rd.Name+".md")] = b
	} else {
		return nil, err
	}

	return files, nil
}

// GenerateResourceCode renders the resource.go.tpl template with info from rd.
func GenerateResourceCode(rd openapi.ResourceDef) (string, error) {
	tpl, err := template.ParseFS(tplFS, "../provider/templates/resource.go.tpl")
	if err != nil {
		return "", err
	}

	a := assemble(rd)
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, a); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// assemble prepares template data incl. enum validator / nested attributes.
func assemble(rd openapi.ResourceDef) resTmpl {
	camel := func(s string) string {
		parts := strings.Split(s, "_")
		for i, p := range parts {
			parts[i] = strings.Title(p)
		}
		return strings.Join(parts, "")
	}

	var ftmpl, attr []fieldTmpl
	for _, f := range rd.CreatePayload.Fields {
		goType := goModelType(f)
		decl := attrDecl(f)
		ft := fieldTmpl{Name: f.Name, GoName: camel(f.Name), GoType: goType, Decl: decl}
		ftmpl = append(ftmpl, ft)
		attr = append(attr, ft)
	}
	recv := strings.ToLower(rd.Name[:1]) + camel(rd.Name)[1:] + "Resource"
	return resTmpl{
		Factory:      "New" + camel(rd.Name) + "Resource",
		Receiver:     recv,
		Model:        camel(rd.Name) + "Model",
		ResourceName: rd.Name,
		Fields:       ftmpl,
		AttrDecls:    attr,
	}
}

// goModelType returns correct types.* symbol for field.
func goModelType(f openapi.FieldDef) string {
	switch f.Kind {
	case openapi.KindPrimitive:
		return primModelType(f.PrimType)
	case openapi.KindList:
		if f.Nested != nil {
			return "types.List"
		} // list<object>
		return "types.List"
	case openapi.KindMap:
		return "types.Map"
	default:
		return "types.Object"
	}
}

func primModelType(p string) string {
	switch p {
	case openapi.PrimInt:
		return "types.Int64"
	case openapi.PrimFloat:
		return "types.Float64"
	case openapi.PrimBool:
		return "types.Bool"
	default:
		return "types.String"
	}
}

// attrDecl builds Terraform attribute declaration code incl. enum validator.
func attrDecl(f openapi.FieldDef) string {
	reqOpt := "Optional: true"
	if f.Required {
		reqOpt = "Required: true"
	}
	var decl string
	switch f.Kind {
	case openapi.KindPrimitive:
		attrType := primAttrType(f.PrimType)
		decl = "resource." + attrType + "{" + reqOpt
		if len(f.EnumVals) > 0 {
			decl += ", Validators: []validator.StringValidator{stringvalidator.OneOf(" + quoteList(f.EnumVals) + ")}"
		}
		decl += "}"
	case openapi.KindList:
		if f.Nested == nil {
			elem := primTypeToken(f.PrimType)
			decl = "resource.ListAttribute{ElementType: " + elem + ", " + reqOpt + "}"
		} else {
			decl = "resource.ListNestedAttribute{" + reqOpt + ", NestedObject: resource.NestedAttributeObject{Attributes: map[string]resource.Attribute{" + nestedAttrMap(*f.Nested) + "}}}"
		}
	case openapi.KindMap:
		elem := primTypeToken(f.PrimType)
		decl = "resource.MapAttribute{ElementType: " + elem + ", " + reqOpt + "}"
	case openapi.KindObject:
		decl = "resource.NestedAttribute{Attributes: map[string]resource.Attribute{" + nestedAttrMap(*f.Nested) + "}, " + reqOpt + "}"
	}
	return decl
}
