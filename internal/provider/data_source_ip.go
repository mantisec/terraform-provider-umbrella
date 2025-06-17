package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// IpDataSource implements the ip data source
type IpDataSource struct {
	client *apiClient
}

// ipDataModel represents the data source data model
type ipDataModel struct {
	Id   types.String `tfsdk:"id"`
	Data types.Set    `tfsdk:"data"`
	Meta types.String `tfsdk:"meta"`
}

// NewIpDataSource creates a new ip data source
func NewIpDataSource() datasource.DataSource {
	return &IpDataSource{}
}

// Metadata returns the data source type name
func (d *IpDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_ip"
}

// Configure configures the data source with the provider client
func (d *IpDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *IpDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "ip data source",
		Attributes: map[string]schema.Attribute{
			"id":   schema.StringAttribute{Computed: true, Description: "Data source identifier"},
			"data": schema.SetAttribute{Computed: true, ElementType: types.StringType},
			"meta": schema.StringAttribute{Computed: true},
		},
	}
}

// Read reads the ip data
func (d *IpDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ipDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /activity/ip

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
