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

// InternalnetworksDataSource implements the internalnetworks data source
type InternalnetworksDataSource struct {
	client *apiClient
}

// internalnetworksDataModel represents the data source data model
type internalnetworksDataModel struct {
	Id types.String `tfsdk:"id"`
}

// NewInternalnetworksDataSource creates a new internalnetworks data source
func NewInternalnetworksDataSource() datasource.DataSource {
	return &InternalnetworksDataSource{}
}

// Metadata returns the data source type name
func (d *InternalnetworksDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_internalnetworks"
}

// Configure configures the data source with the provider client
func (d *InternalnetworksDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *InternalnetworksDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "internalnetworks data source",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, Description: "Data source identifier"},
		},
	}
}

// Read reads the internalnetworks data
func (d *InternalnetworksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config internalnetworksDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /internalnetworks

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
