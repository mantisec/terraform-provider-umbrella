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

// ThreatNamesDataSource implements the threat_names data source
type ThreatNamesDataSource struct {
	client *apiClient
}

// threat_namesDataModel represents the data source data model
type threat_namesDataModel struct {
	Id   types.String `tfsdk:"id"`
	Data types.String `tfsdk:"data"`
	Meta types.String `tfsdk:"meta"`
}

// NewThreatNamesDataSource creates a new threat_names data source
func NewThreatNamesDataSource() datasource.DataSource {
	return &ThreatNamesDataSource{}
}

// Metadata returns the data source type name
func (d *ThreatNamesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_threat_names"
}

// Configure configures the data source with the provider client
func (d *ThreatNamesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *ThreatNamesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "threat_names data source",
		Attributes: map[string]schema.Attribute{
			"id":   schema.StringAttribute{Computed: true, Description: "Data source identifier"},
			"data": schema.StringAttribute{Computed: true},
			"meta": schema.StringAttribute{Computed: true},
		},
	}
}

// Read reads the threat_names data
func (d *ThreatNamesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config threat_namesDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /threat-names/{threatnameid}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
