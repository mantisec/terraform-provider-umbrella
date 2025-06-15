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

// RequestsByTimerangeDataSource implements the requests_by_timerange data source
type RequestsByTimerangeDataSource struct {
	client *apiClient
}

// requests_by_timerangeDataModel represents the data source data model
type requests_by_timerangeDataModel struct {
	Id   types.String `tfsdk:"id"`
	Data types.Set    `tfsdk:"data"`
	Meta types.String `tfsdk:"meta"`
}

// NewRequestsByTimerangeDataSource creates a new requests_by_timerange data source
func NewRequestsByTimerangeDataSource() datasource.DataSource {
	return &RequestsByTimerangeDataSource{}
}

// Metadata returns the data source type name
func (d *RequestsByTimerangeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_requests_by_timerange"
}

// Configure configures the data source with the provider client
func (d *RequestsByTimerangeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *RequestsByTimerangeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "requests_by_timerange data source",
		Attributes: map[string]schema.Attribute{
			"id":   schema.StringAttribute{Computed: true, Description: "Data source identifier"},
			"data": schema.SetAttribute{Computed: true, ElementType: types.StringType},
			"meta": schema.StringAttribute{Computed: true},
		},
	}
}

// Read reads the requests_by_timerange data
func (d *RequestsByTimerangeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config requests_by_timerangeDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /requests-by-timerange/{type}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
