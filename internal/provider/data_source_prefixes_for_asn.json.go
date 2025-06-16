package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PrefixesForAsnDataSource implements the prefixes_for_asn.json data source
type PrefixesForAsnDataSource struct {
	client *apiClient
}

// prefixes_for_asnDataModel represents the data source data model
type prefixes_for_asnDataModel struct {
	Id   types.String `tfsdk:"id"`
	Cidr types.Set    `tfsdk:"cidr"`
	Geo  types.Object `tfsdk:"geo"`
}

// NewPrefixesForAsnDataSource creates a new prefixes_for_asn.json data source
func NewPrefixesForAsnDataSource() datasource.DataSource {
	return &PrefixesForAsnDataSource{}
}

// Metadata returns the data source type name
func (d *PrefixesForAsnDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_prefixes_for_asn.json"
}

// Configure configures the data source with the provider client
func (d *PrefixesForAsnDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *PrefixesForAsnDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "prefixes_for_asn.json data source",
		Attributes: map[string]schema.Attribute{
			"id":   schema.StringAttribute{Computed: true, Description: "Data source identifier"},
			"cidr": schema.SetAttribute{Computed: true, Description: "A list of the CIDR range of IP addresses associated with this AS. The CIDR contains the IP prefix for the ASN.", ElementType: types.StringType},
			"geo":  schema.StringAttribute{Computed: true, Description: "Geo is a hash reference with the country name and country code (the code corresponds to the country code list for ISO-3166-1 alpha-2). For more information, see [ISO 3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)."},
		},
	}
}

// Read reads the prefixes_for_asn.json data
func (d *PrefixesForAsnDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config prefixes_for_asnDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /bgp_routes/asn/{asn}/prefixes_for_asn.json

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
