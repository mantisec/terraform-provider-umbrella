package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// AsForIpDataSource implements the as_for_ip.json data source
type AsForIpDataSource struct {
	client *apiClient
}

// as_for_ipDataModel represents the data source data model
type as_for_ipDataModel struct {
	Id           types.String `tfsdk:"id"`
	Description  types.String `tfsdk:"description"`
	Asn          types.String `tfsdk:"asn"`
	Cidr         types.String `tfsdk:"cidr"`
	CreationDate types.String `tfsdk:"creation_date"`
	Ir           types.Int64  `tfsdk:"ir"`
}

// NewAsForIpDataSource creates a new as_for_ip.json data source
func NewAsForIpDataSource() datasource.DataSource {
	return &AsForIpDataSource{}
}

// Metadata returns the data source type name
func (d *AsForIpDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_as_for_ip.json"
}

// Configure configures the data source with the provider client
func (d *AsForIpDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *AsForIpDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "as_for_ip.json data source",
		Attributes: map[string]schema.Attribute{
			"id":            schema.StringAttribute{Computed: true, Description: "Data source identifier"},
			"description":   schema.StringAttribute{Computed: true, Description: "Network Owner Description as provided by the network owner."},
			"asn":           schema.StringAttribute{Computed: true, Description: "The autonomous system number (ASN) associated with the IP address."},
			"cidr":          schema.StringAttribute{Computed: true, Description: "The IP CIDR for the ASN."},
			"creation_date": schema.StringAttribute{Computed: true, Description: "The date when the AS was first created."},
			"ir":            schema.Int64Attribute{Computed: true, Description: "The IR number corresponds to one of the 5 Regional Internet Registries (RIR).  Registry  Number  Region    Registry  1  AfriNIC: Africa   Registry  2  APNIC: Asia, Australia, New Zealand, and neighboring countries.   Registry  3  ARIN: United States, Canada, several parts of the Caribbean region, and Antarctica.   Registry  4  LACNIC: Latin America and parts of the Caribbean region.   Registry  5  RIPE NCC: Europe, Russia, the Middle East, and Central Asia.   Registry  0  Unknown / Not Available"},
		},
	}
}

// Read reads the as_for_ip.json data
func (d *AsForIpDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config as_for_ipDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /bgp_routes/ip/{ip}/as_for_ip.json

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
