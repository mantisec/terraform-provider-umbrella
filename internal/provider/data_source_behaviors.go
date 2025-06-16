package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// BehaviorsDataSource implements the behaviors data source
type BehaviorsDataSource struct {
	client *apiClient
}

// behaviorsDataModel represents the data source data model
type behaviorsDataModel struct {
	Id types.String `tfsdk:"id"`
}

// NewBehaviorsDataSource creates a new behaviors data source
func NewBehaviorsDataSource() datasource.DataSource {
	return &BehaviorsDataSource{}
}

// Metadata returns the data source type name
func (d *BehaviorsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_behaviors"
}

// Configure configures the data source with the provider client
func (d *BehaviorsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *BehaviorsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "behaviors data source",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, Description: "Data source identifier"},
		},
	}
}

// Read reads the behaviors data
func (d *BehaviorsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config behaviorsDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /sample/{hash}/behaviors

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
