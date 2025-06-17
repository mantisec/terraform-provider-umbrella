package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VolumeDataSource implements the volume data source
type VolumeDataSource struct {
	client *apiClient
}

// volumeDataModel represents the data source data model
type volumeDataModel struct {
	Id      types.String `tfsdk:"id"`
	Queries types.Set    `tfsdk:"queries"`
	Dates   types.Set    `tfsdk:"dates"`
}

// NewVolumeDataSource creates a new volume data source
func NewVolumeDataSource() datasource.DataSource {
	return &VolumeDataSource{}
}

// Metadata returns the data source type name
func (d *VolumeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_volume"
}

// Configure configures the data source with the provider client
func (d *VolumeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *VolumeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "volume data source",
		Attributes: map[string]schema.Attribute{
			"id":      schema.StringAttribute{Computed: true, Description: "Data source identifier"},
			"queries": schema.SetAttribute{Computed: true, Description: "The list of the numbers of DNS queries requested for the domain in one hour, listed in ascending order.", ElementType: types.StringType},
			"dates":   schema.SetAttribute{Computed: true, Description: "The list of dates recorded for the domain.", ElementType: types.StringType},
		},
	}
}

// Read reads the volume data
func (d *VolumeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config volumeDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /domains/volume/{domain}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
