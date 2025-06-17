package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// TopCategoriesDataSource implements the top_categories data source
type TopCategoriesDataSource struct {
	client *apiClient
}

// top_categoriesDataModel represents the data source data model
type top_categoriesDataModel struct {
	Id   types.String `tfsdk:"id"`
	Data types.Set    `tfsdk:"data"`
	Meta types.String `tfsdk:"meta"`
}

// NewTopCategoriesDataSource creates a new top_categories data source
func NewTopCategoriesDataSource() datasource.DataSource {
	return &TopCategoriesDataSource{}
}

// Metadata returns the data source type name
func (d *TopCategoriesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_top_categories"
}

// Configure configures the data source with the provider client
func (d *TopCategoriesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *TopCategoriesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "top_categories data source",
		Attributes: map[string]schema.Attribute{
			"id":   schema.StringAttribute{Computed: true, Description: "Data source identifier"},
			"data": schema.SetAttribute{Computed: true, ElementType: types.StringType},
			"meta": schema.StringAttribute{Computed: true},
		},
	}
}

// Read reads the top_categories data
func (d *TopCategoriesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config top_categoriesDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /top-categories

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
