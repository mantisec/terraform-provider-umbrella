package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NameserversDataSource implements the nameservers data source
type NameserversDataSource struct {
	client *apiClient
}

// nameserversDataModel represents the data source data model
type nameserversDataModel struct {
	Id                types.String `tfsdk:"id"`
	Totalresults      types.Int64  `tfsdk:"totalResults"`
	Moredataavailable types.String `tfsdk:"moreDataAvailable"`
	Limit             types.String `tfsdk:"limit"`
	Sortfield         types.String `tfsdk:"sortField"`
	Domains           types.Set    `tfsdk:"domains"`
}

// NewNameserversDataSource creates a new nameservers data source
func NewNameserversDataSource() datasource.DataSource {
	return &NameserversDataSource{}
}

// Metadata returns the data source type name
func (d *NameserversDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_nameservers"
}

// Configure configures the data source with the provider client
func (d *NameserversDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *NameserversDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "nameservers data source",
		Attributes: map[string]schema.Attribute{
			"id":                schema.StringAttribute{Computed: true, Description: "Data source identifier"},
			"totalResults":      schema.Int64Attribute{Computed: true, Description: "The total number of WHOIS records found for this query."},
			"moreDataAvailable": schema.StringAttribute{Computed: true},
			"limit":             schema.StringAttribute{Computed: true},
			"sortField":         schema.StringAttribute{Computed: true, Description: "The field that is used to sort the collection."},
			"domains":           schema.SetAttribute{Computed: true, Description: "The list of information about the WHOIS emails and nameservers.", ElementType: types.StringType},
		},
	}
}

// Read reads the nameservers data
func (d *NameserversDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config nameserversDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /whois/nameservers/{nameserver}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
