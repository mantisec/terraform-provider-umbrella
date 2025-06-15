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

// BandwidthByTimerangeDataSource implements the bandwidth_by_timerange data source
type BandwidthByTimerangeDataSource struct {
	client *apiClient
}

// bandwidth_by_timerangeDataModel represents the data source data model
type bandwidth_by_timerangeDataModel struct {
	Id   types.String `tfsdk:"id"`
	Data types.Set    `tfsdk:"data"`
	Meta types.String `tfsdk:"meta"`
}

// NewBandwidthByTimerangeDataSource creates a new bandwidth_by_timerange data source
func NewBandwidthByTimerangeDataSource() datasource.DataSource {
	return &BandwidthByTimerangeDataSource{}
}

// Metadata returns the data source type name
func (d *BandwidthByTimerangeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_bandwidth_by_timerange"
}

// Configure configures the data source with the provider client
func (d *BandwidthByTimerangeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *BandwidthByTimerangeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "bandwidth_by_timerange data source",
		Attributes: map[string]schema.Attribute{
			"id":   schema.StringAttribute{Computed: true, Description: "Data source identifier"},
			"data": schema.SetAttribute{Computed: true, ElementType: types.StringType},
			"meta": schema.StringAttribute{Computed: true},
		},
	}
}

// Read reads the bandwidth_by_timerange data
func (d *BandwidthByTimerangeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config bandwidth_by_timerangeDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /bandwidth-by-timerange

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
