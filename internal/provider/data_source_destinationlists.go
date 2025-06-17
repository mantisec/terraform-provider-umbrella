package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// DestinationlistsDataSource implements the destinationlists data source
type DestinationlistsDataSource struct {
	client *apiClient
}

// destinationlistsDataModel represents the data source data model
type destinationlistsDataModel struct {
	Id types.String `tfsdk:"id"`
}

// NewDestinationlistsDataSource creates a new destinationlists data source
func NewDestinationlistsDataSource() datasource.DataSource {
	return &DestinationlistsDataSource{}
}

// Metadata returns the data source type name
func (d *DestinationlistsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_destinationlists"
}

// Configure configures the data source with the provider client
func (d *DestinationlistsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *DestinationlistsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "destinationlists data source",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, Description: "Data source identifier"},
		},
	}
}

// Read reads the destinationlists data
func (d *DestinationlistsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config destinationlistsDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /destinationlists/{destinationListId}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
