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

// ActivityDataSource implements the activity data source
type ActivityDataSource struct {
	client *apiClient
}

// activityDataModel represents the data source data model
type activityDataModel struct {
	Id   types.String `tfsdk:"id"`
	Meta types.String `tfsdk:"meta"`
	Data types.Set    `tfsdk:"data"`
}

// NewActivityDataSource creates a new activity data source
func NewActivityDataSource() datasource.DataSource {
	return &ActivityDataSource{}
}

// Metadata returns the data source type name
func (d *ActivityDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_activity"
}

// Configure configures the data source with the provider client
func (d *ActivityDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *ActivityDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "activity data source",
		Attributes: map[string]schema.Attribute{
			"id":   schema.StringAttribute{Computed: true, Description: "Data source identifier"},
			"meta": schema.StringAttribute{Computed: true},
			"data": schema.SetAttribute{Computed: true, ElementType: types.StringType},
		},
	}
}

// Read reads the activity data
func (d *ActivityDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config activityDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /activity

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
