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

// SearchDataSource implements the search data source
type SearchDataSource struct {
	client *apiClient
}

// searchDataModel represents the data source data model
type searchDataModel struct {
	Id                types.String `tfsdk:"id"`
	Totalresults      types.Int64  `tfsdk:"totalResults"`
	Offset            types.String `tfsdk:"offset"`
	Moredataavailable types.Bool   `tfsdk:"moreDataAvailable"`
	Limit             types.Int64  `tfsdk:"limit"`
	Sortfield         types.String `tfsdk:"sortField"`
	Records           types.Set    `tfsdk:"records"`
	Expression        types.String `tfsdk:"expression"`
	Matches           types.Set    `tfsdk:"matches"`
}

// NewSearchDataSource creates a new search data source
func NewSearchDataSource() datasource.DataSource {
	return &SearchDataSource{}
}

// Metadata returns the data source type name
func (d *SearchDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_search"
}

// Configure configures the data source with the provider client
func (d *SearchDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *SearchDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "search data source",
		Attributes: map[string]schema.Attribute{
			"id":                schema.StringAttribute{Computed: true, Description: "Data source identifier"},
			"totalResults":      schema.Int64Attribute{Computed: true, Description: "The total number of results for this search."},
			"offset":            schema.StringAttribute{Computed: true},
			"moreDataAvailable": schema.BoolAttribute{Computed: true, Description: "Specifies whether there is more than 10 results for this search."},
			"limit":             schema.Int64Attribute{Computed: true, Description: "The total number of results for this page. Default limit is 10."},
			"sortField":         schema.StringAttribute{Computed: true, Description: "The field that is used to sort the collection."},
			"records":           schema.SetAttribute{Computed: true, Description: "The list of WHOIS records.", ElementType: types.StringType},
			"expression":        schema.StringAttribute{Computed: true, Description: "Specifies the regular expression used in the search."},
			"matches":           schema.SetAttribute{Computed: true, Description: "The list of matching records.", ElementType: types.StringType},
		},
	}
}

// Read reads the search data
func (d *SearchDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config searchDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /whois/search/{searchField}/{regexExpression}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
