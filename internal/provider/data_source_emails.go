package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// EmailsDataSource implements the emails data source
type EmailsDataSource struct {
	client *apiClient
}

// emailsDataModel represents the data source data model
type emailsDataModel struct {
	Id                types.String `tfsdk:"id"`
	Totalresults      types.Int64  `tfsdk:"totalResults"`
	Offset            types.String `tfsdk:"offset"`
	Moredataavailable types.Bool   `tfsdk:"moreDataAvailable"`
	Limit             types.Int64  `tfsdk:"limit"`
	Sortfield         types.String `tfsdk:"sortField"`
	Domains           types.Set    `tfsdk:"domains"`
}

// NewEmailsDataSource creates a new emails data source
func NewEmailsDataSource() datasource.DataSource {
	return &EmailsDataSource{}
}

// Metadata returns the data source type name
func (d *EmailsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_emails"
}

// Configure configures the data source with the provider client
func (d *EmailsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *EmailsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "emails data source",
		Attributes: map[string]schema.Attribute{
			"id":                schema.StringAttribute{Computed: true, Description: "Data source identifier"},
			"totalResults":      schema.Int64Attribute{Computed: true, Description: "The total number of results for this email address."},
			"offset":            schema.StringAttribute{Computed: true},
			"moreDataAvailable": schema.BoolAttribute{Computed: true, Description: "Specifies whether there is more than 500 results for this email."},
			"limit":             schema.Int64Attribute{Computed: true, Description: "The number of results returned in the response. The default limit is 500."},
			"sortField":         schema.StringAttribute{Computed: true, Description: "The field that is used to sort the collection."},
			"domains":           schema.SetAttribute{Computed: true, Description: "The list of domains registered by this email and if the domain is currently registered by this email address.", ElementType: types.StringType},
		},
	}
}

// Read reads the emails data
func (d *EmailsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config emailsDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /whois/emails/{email}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
