package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// CategoryRequestsByOrgDataSource implements the category_requests_by_org data source
type CategoryRequestsByOrgDataSource struct {
	client *apiClient
}

// category_requests_by_orgDataModel represents the data source data model
type category_requests_by_orgDataModel struct {
	Id   types.String `tfsdk:"id"`
	Data types.Set    `tfsdk:"data"`
	Meta types.String `tfsdk:"meta"`
}

// NewCategoryRequestsByOrgDataSource creates a new category_requests_by_org data source
func NewCategoryRequestsByOrgDataSource() datasource.DataSource {
	return &CategoryRequestsByOrgDataSource{}
}

// Metadata returns the data source type name
func (d *CategoryRequestsByOrgDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_category_requests_by_org"
}

// Configure configures the data source with the provider client
func (d *CategoryRequestsByOrgDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *CategoryRequestsByOrgDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "category_requests_by_org data source",
		Attributes: map[string]schema.Attribute{
			"id":   schema.StringAttribute{Computed: true, Description: "Data source identifier"},
			"data": schema.SetAttribute{Computed: true, ElementType: types.StringType},
			"meta": schema.StringAttribute{Computed: true},
		},
	}
}

// Read reads the category_requests_by_org data
func (d *CategoryRequestsByOrgDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config category_requests_by_orgDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /providers/category-requests-by-org

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
