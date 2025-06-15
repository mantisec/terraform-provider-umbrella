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

// SamplesDataSource implements the samples data source
type SamplesDataSource struct {
	client *apiClient
}

// samplesDataModel represents the data source data model
type samplesDataModel struct {
	Id types.String `tfsdk:"id"`
}

// NewSamplesDataSource creates a new samples data source
func NewSamplesDataSource() datasource.DataSource {
	return &SamplesDataSource{}
}

// Metadata returns the data source type name
func (d *SamplesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_samples"
}

// Configure configures the data source with the provider client
func (d *SamplesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *SamplesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "samples data source",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, Description: "Data source identifier"},
		},
	}
}

// Read reads the samples data
func (d *SamplesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config samplesDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /samples/{destination}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
