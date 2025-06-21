{{`// Code generated from OpenAPI; DO NOT EDIT.
package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func {{ .Factory }}() datasource.DataSource { return &{{ .Receiver }}{} }

type {{ .Receiver }} struct{}

type {{ .Model }} struct {
{{- range .Fields }}
	{{ .GoName }} {{ .GoType }} `tfsdk:"{{ .Name }}"`
{{- end }} }

func (d *{{ .Receiver }}) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_{{ .ResourceName }}"
}

func (d *{{ .Receiver }}) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource.Schema{Description: "Auto‑generated Umbrella data source", Attributes: map[string]datasource.Attribute{
{{- range .AttrDecls }}
		"{{ .Name }}": {{ .Decl }},
{{- end }} }}
}

// TODO: Read() implementation invokes Umbrella GET list/find endpoint.
`}}
