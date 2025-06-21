package openapi

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi/datamodel"
)

func isDebugPath(p *datamodel.PathItem) bool {
	if p == nil {
		return false
	}
	for _, op := range []*datamodel.Operation{p.Get, p.Post, p.Put, p.Patch} {
		if op == nil {
			continue
		}
		for _, t := range op.Tags {
			if strings.EqualFold(t, "debug") || strings.EqualFold(t, "debugging") {
				return true
			}
		}
	}
	return false
}

func firstPayload(p *datamodel.PathItem) *SchemaSpec {
	if p.Post != nil && p.Post.RequestBody != nil {
		return extractRequestSchema(p.Post.RequestBody.Value)
	}
	if p.Put != nil && p.Put.RequestBody != nil {
		return extractRequestSchema(p.Put.RequestBody.Value)
	}
	if p.Patch != nil && p.Patch.RequestBody != nil {
		return extractRequestSchema(p.Patch.RequestBody.Value)
	}
	return nil
}

func first200Schema(p *datamodel.PathItem) *SchemaSpec {
	if p.Get == nil || p.Get.Responses == nil {
		return nil
	}
	if resp, ok := p.Get.Responses.Value["200"]; ok {
		return extractResponseSchema(resp.Value)
	}
	return nil
}

func extractRequestSchema(body *datamodel.RequestBody) *SchemaSpec {
	if body == nil || body.Content == nil {
		return nil
	}
	if mt, ok := body.Content["application/json"]; ok {
		if mt.Schema != nil && mt.Schema.Schema() != nil {
			return flattenSchema(mt.Schema.Schema())
		}
	}
	return nil
}

func extractResponseSchema(resp *datamodel.Response) *SchemaSpec {
	if resp == nil || resp.Content == nil {
		return nil
	}
	if mt, ok := resp.Content["application/json"]; ok {
		if mt.Schema != nil && mt.Schema.Schema() != nil {
			return flattenSchema(mt.Schema.Schema())
		}
	}
	return nil
}

func contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

var snakeRx = regexp.MustCompile(`([a-z0-9])([A-Z])`)

func toSnake(in string) string {
	return strings.ToLower(snakeRx.ReplaceAllString(in, "${1}_${2}"))
}

func ParseSpec(specFile string) ([]ResourceDef, error) {

	rawFile, err := os.ReadFile(specFile)

	if err != nil {
		return nil, err
	}

	document, err := libopenapi.NewDocument(rawFile)
	// if anything went wrong, an error is thrown
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	doc, err := document.BuildV3Model()
	if err != nil {
		return nil, err
	}

	paths := doc.Model.Paths.PathItems
	collectionMap := map[string]*ResourceDef{}
	for rawPath, pItem := range paths {
		if isDebugPath(pItem) {
			continue
		}

		if strings.Contains(rawPath, "{") {
			prefix := rawPath[:strings.Index(rawPath, "{")]
			if rd, ok := collectionMap[prefix]; ok {
				rd.ItemPath = rawPath
				rd.UpdatePayload = firstPayload(pItem)
				rd.HasDelete = pItem.Delete != nil
			}
			continue
		}
		parts := strings.Split(strings.Trim(rawPath, "/"), "/")
		resName := strings.TrimSuffix(parts[len(parts)-1], "s")
		collectionMap[rawPath] = &ResourceDef{
			Name:           toSnake(resName),
			CollectionPath: rawPath,
			CreatePayload:  firstPayload(pItem),
			ReadSchema:     first200Schema(pItem),
		}
	}
	var out []ResourceDef
	for _, v := range collectionMap {
		out = append(out, *v)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out, nil
}

func flattenSchema(s *datamodel.Schema) *SchemaSpec {
	if s == nil {
		return nil
	}
	var spec SchemaSpec
	for propName, prop := range s.Properties {
		spec.Fields = append(spec.Fields, convertProperty(propName, prop))
	}
	return &spec
}

func convertProperty(name string, prop *datamodel.Schema) FieldDef {
	fd := FieldDef{Name: toSnake(name), JSONName: name, Required: contains(prop.Required, name), ReadOnly: prop.ReadOnly}

	switch prop.Type {
	case "string":
		fd.Kind, fd.PrimType = KindPrimitive, PrimString
	case "integer":
		fd.Kind, fd.PrimType = KindPrimitive, PrimInt
	case "number":
		fd.Kind, fd.PrimType = KindPrimitive, PrimFloat
	case "boolean":
		fd.Kind, fd.PrimType = KindPrimitive, PrimBool
	case "array":
		if prop.Items != nil && prop.Items.Schema() != nil {
			item := prop.Items.Schema()
			// Detect primitive vs complex list
			if isPrimitive(item.Type) {
				fd.Kind = KindList
				fd.PrimType = primType(item.Type)
			} else {
				fd.Kind = KindList
				fd.Nested = flattenSchema(item)
			}
		}
	case "object":
		// primitive map or complex object
		if prop.AdditionalProperties != nil && prop.AdditionalProperties.Schema() != nil && isPrimitive(prop.AdditionalProperties.Schema().Type) {
			fd.Kind = KindMap
			fd.PrimType = primType(prop.AdditionalProperties.Schema().Type)
		} else {
			fd.Kind = KindObject
			fd.Nested = flattenSchema(prop)
		}
	default:
		fd.Kind, fd.PrimType = KindPrimitive, PrimString
	}

	for _, e := range prop.Enum {
		fd.EnumVals = append(fd.EnumVals, fmt.Sprintf("%v", e))
	}
	return fd
}

func isPrimitive(t string) bool {
	return t == "string" || t == "integer" || t == "number" || t == "boolean"
}
func primType(t string) string {
	switch t {
	case "integer":
		return PrimInt
	case "number":
		return PrimFloat
	case "boolean":
		return PrimBool
	default:
		return PrimString
	}
}
