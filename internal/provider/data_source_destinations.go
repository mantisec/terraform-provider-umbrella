package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// DestinationsDataSource implements the destinations data source
type DestinationsDataSource struct {
	client *apiClient
}

// destinationsDataModel represents the data source data model
type destinationsDataModel struct {
	Id types.String `tfsdk:"id"`
}

// NewDestinationsDataSource creates a new destinations data source
func NewDestinationsDataSource() datasource.DataSource {
	return &DestinationsDataSource{}
}

// Metadata returns the data source type name
func (d *DestinationsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_destinations"
}

// Configure configures the data source with the provider client
func (d *DestinationsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *DestinationsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "destinations data source",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, Description: "Data source identifier"},
		},
	}
}

// Read reads the destinations data
func (d *DestinationsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config destinationsDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /destinationlists/{destinationListId}/destinations

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
