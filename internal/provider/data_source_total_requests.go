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

// TotalRequestsDataSource implements the total_requests data source
type TotalRequestsDataSource struct {
	client *apiClient
}

// total_requestsDataModel represents the data source data model
type total_requestsDataModel struct {
	Id   types.String `tfsdk:"id"`
	Data types.String `tfsdk:"data"`
	Meta types.String `tfsdk:"meta"`
}

// NewTotalRequestsDataSource creates a new total_requests data source
func NewTotalRequestsDataSource() datasource.DataSource {
	return &TotalRequestsDataSource{}
}

// Metadata returns the data source type name
func (d *TotalRequestsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_total_requests"
}

// Configure configures the data source with the provider client
func (d *TotalRequestsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *TotalRequestsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "total_requests data source",
		Attributes: map[string]schema.Attribute{
			"id":   schema.StringAttribute{Computed: true, Description: "Data source identifier"},
			"data": schema.StringAttribute{Computed: true},
			"meta": schema.StringAttribute{Computed: true},
		},
	}
}

// Read reads the total_requests data
func (d *TotalRequestsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config total_requestsDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /total-requests

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
