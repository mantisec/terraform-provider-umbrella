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

// AmpRetrospectiveDataSource implements the amp_retrospective data source
type AmpRetrospectiveDataSource struct {
	client *apiClient
}

// amp_retrospectiveDataModel represents the data source data model
type amp_retrospectiveDataModel struct {
	Id   types.String `tfsdk:"id"`
	Data types.Set    `tfsdk:"data"`
	Meta types.String `tfsdk:"meta"`
}

// NewAmpRetrospectiveDataSource creates a new amp_retrospective data source
func NewAmpRetrospectiveDataSource() datasource.DataSource {
	return &AmpRetrospectiveDataSource{}
}

// Metadata returns the data source type name
func (d *AmpRetrospectiveDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_amp_retrospective"
}

// Configure configures the data source with the provider client
func (d *AmpRetrospectiveDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *AmpRetrospectiveDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "amp_retrospective data source",
		Attributes: map[string]schema.Attribute{
			"id":   schema.StringAttribute{Computed: true, Description: "Data source identifier"},
			"data": schema.SetAttribute{Computed: true, ElementType: types.StringType},
			"meta": schema.StringAttribute{Computed: true},
		},
	}
}

// Read reads the amp_retrospective data
func (d *AmpRetrospectiveDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config amp_retrospectiveDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /activity/amp-retrospective

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
