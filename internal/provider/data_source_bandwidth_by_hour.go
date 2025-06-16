package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// BandwidthByHourDataSource implements the bandwidth_by_hour data source
type BandwidthByHourDataSource struct {
	client *apiClient
}

// bandwidth_by_hourDataModel represents the data source data model
type bandwidth_by_hourDataModel struct {
	Id   types.String `tfsdk:"id"`
	Data types.Set    `tfsdk:"data"`
	Meta types.String `tfsdk:"meta"`
}

// NewBandwidthByHourDataSource creates a new bandwidth_by_hour data source
func NewBandwidthByHourDataSource() datasource.DataSource {
	return &BandwidthByHourDataSource{}
}

// Metadata returns the data source type name
func (d *BandwidthByHourDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_bandwidth_by_hour"
}

// Configure configures the data source with the provider client
func (d *BandwidthByHourDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *BandwidthByHourDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "bandwidth_by_hour data source",
		Attributes: map[string]schema.Attribute{
			"id":   schema.StringAttribute{Computed: true, Description: "Data source identifier"},
			"data": schema.SetAttribute{Computed: true, ElementType: types.StringType},
			"meta": schema.StringAttribute{Computed: true},
		},
	}
}

// Read reads the bandwidth_by_hour data
func (d *BandwidthByHourDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config bandwidth_by_hourDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /bandwidth-by-hour

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
