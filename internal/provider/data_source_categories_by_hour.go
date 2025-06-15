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

// CategoriesByHourDataSource implements the categories_by_hour data source
type CategoriesByHourDataSource struct {
	client *apiClient
}

// categories_by_hourDataModel represents the data source data model
type categories_by_hourDataModel struct {
	Id   types.String `tfsdk:"id"`
	Meta types.String `tfsdk:"meta"`
	Data types.Set    `tfsdk:"data"`
}

// NewCategoriesByHourDataSource creates a new categories_by_hour data source
func NewCategoriesByHourDataSource() datasource.DataSource {
	return &CategoriesByHourDataSource{}
}

// Metadata returns the data source type name
func (d *CategoriesByHourDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_categories_by_hour"
}

// Configure configures the data source with the provider client
func (d *CategoriesByHourDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *CategoriesByHourDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "categories_by_hour data source",
		Attributes: map[string]schema.Attribute{
			"id":   schema.StringAttribute{Computed: true, Description: "Data source identifier"},
			"meta": schema.StringAttribute{Computed: true},
			"data": schema.SetAttribute{Computed: true, ElementType: types.StringType},
		},
	}
}

// Read reads the categories_by_hour data
func (d *CategoriesByHourDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config categories_by_hourDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /categories-by-hour

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
