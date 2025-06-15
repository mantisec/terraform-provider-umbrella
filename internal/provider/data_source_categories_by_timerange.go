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

// CategoriesByTimerangeDataSource implements the categories_by_timerange data source
type CategoriesByTimerangeDataSource struct {
	client *apiClient
}

// categories_by_timerangeDataModel represents the data source data model
type categories_by_timerangeDataModel struct {
	Id   types.String `tfsdk:"id"`
	Data types.Set    `tfsdk:"data"`
	Meta types.String `tfsdk:"meta"`
}

// NewCategoriesByTimerangeDataSource creates a new categories_by_timerange data source
func NewCategoriesByTimerangeDataSource() datasource.DataSource {
	return &CategoriesByTimerangeDataSource{}
}

// Metadata returns the data source type name
func (d *CategoriesByTimerangeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_categories_by_timerange"
}

// Configure configures the data source with the provider client
func (d *CategoriesByTimerangeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *CategoriesByTimerangeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "categories_by_timerange data source",
		Attributes: map[string]schema.Attribute{
			"id":   schema.StringAttribute{Computed: true, Description: "Data source identifier"},
			"data": schema.SetAttribute{Computed: true, ElementType: types.StringType},
			"meta": schema.StringAttribute{Computed: true},
		},
	}
}

// Read reads the categories_by_timerange data
func (d *CategoriesByTimerangeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config categories_by_timerangeDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /categories-by-timerange/{type}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
