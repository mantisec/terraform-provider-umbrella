package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// TopmillionDataSource implements the topmillion data source
type TopmillionDataSource struct {
	client *apiClient
}

// topmillionDataModel represents the data source data model
type topmillionDataModel struct {
	Id types.String `tfsdk:"id"`
}

// NewTopmillionDataSource creates a new topmillion data source
func NewTopmillionDataSource() datasource.DataSource {
	return &TopmillionDataSource{}
}

// Metadata returns the data source type name
func (d *TopmillionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_topmillion"
}

// Configure configures the data source with the provider client
func (d *TopmillionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *TopmillionDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "topmillion data source",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, Description: "Data source identifier"},
		},
	}
}

// Read reads the topmillion data
func (d *TopmillionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config topmillionDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /topmillion

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
