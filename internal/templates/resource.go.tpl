{{`// Code generated from OpenAPI; DO NOT EDIT.
package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	stringvalidator "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
)

func {{ .Factory }}() resource.Resource { return &{{ .Receiver }}{} }

type {{ .Receiver }} struct{}

type {{ .Model }} struct {
	ID types.Int64 `tfsdk:"id"`
{{- range .Fields }}
	{{ .GoName }} {{ .GoType }} `tfsdk:"{{ .Name }}"`
{{- end }} }

func (r *{{ .Receiver }}) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_{{ .ResourceName }}"
}

func (r *{{ .Receiver }}) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource.Schema{Description: "Auto-generated Umbrella resource", Attributes: map[string]resource.Attribute{
		"id": resource.Int64Attribute{Computed: true},
{{- range .AttrDecls }}
		"{{ .Name }}": {{ .Decl }},
{{- end }} }}
}

// TODO: Create, Read, Update, Delete
`}}
