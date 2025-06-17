package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SummariesByDestinationDataSource implements the summaries_by_destination data source
type SummariesByDestinationDataSource struct {
	client *apiClient
}

// summaries_by_destinationDataModel represents the data source data model
type summaries_by_destinationDataModel struct {
	Id   types.String `tfsdk:"id"`
	Data types.Set    `tfsdk:"data"`
	Meta types.String `tfsdk:"meta"`
}

// NewSummariesByDestinationDataSource creates a new summaries_by_destination data source
func NewSummariesByDestinationDataSource() datasource.DataSource {
	return &SummariesByDestinationDataSource{}
}

// Metadata returns the data source type name
func (d *SummariesByDestinationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_summaries_by_destination"
}

// Configure configures the data source with the provider client
func (d *SummariesByDestinationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *SummariesByDestinationDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "summaries_by_destination data source",
		Attributes: map[string]schema.Attribute{
			"id":   schema.StringAttribute{Computed: true, Description: "Data source identifier"},
			"data": schema.SetAttribute{Computed: true, ElementType: types.StringType},
			"meta": schema.StringAttribute{Computed: true},
		},
	}
}

// Read reads the summaries_by_destination data
func (d *SummariesByDestinationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config summaries_by_destinationDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /summaries-by-destination

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
