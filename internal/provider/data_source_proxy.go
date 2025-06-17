package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProxyDataSource implements the proxy data source
type ProxyDataSource struct {
	client *apiClient
}

// proxyDataModel represents the data source data model
type proxyDataModel struct {
	Id   types.String `tfsdk:"id"`
	Meta types.String `tfsdk:"meta"`
	Data types.Set    `tfsdk:"data"`
}

// NewProxyDataSource creates a new proxy data source
func NewProxyDataSource() datasource.DataSource {
	return &ProxyDataSource{}
}

// Metadata returns the data source type name
func (d *ProxyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_proxy"
}

// Configure configures the data source with the provider client
func (d *ProxyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*apiClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected DataSource Configure Type", "Expected *apiClient")
		return
	}

	d.client = client
}

// Schema defines the data source schema
func (d *ProxyDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "proxy data source",
		Attributes: map[string]schema.Attribute{
			"id":   schema.StringAttribute{Computed: true, Description: "Data source identifier"},
			"meta": schema.StringAttribute{Computed: true},
			"data": schema.SetAttribute{Computed: true, ElementType: types.StringType},
		},
	}
}

// Read reads the proxy data
func (d *ProxyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config proxyDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /activity/proxy

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
