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

// SubdomainsDataSource implements the subdomains data source
type SubdomainsDataSource struct {
	client *apiClient
}

// subdomainsDataModel represents the data source data model
type subdomainsDataModel struct {
	Id types.String `tfsdk:"id"`
}

// NewSubdomainsDataSource creates a new subdomains data source
func NewSubdomainsDataSource() datasource.DataSource {
	return &SubdomainsDataSource{}
}

// Metadata returns the data source type name
func (d *SubdomainsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_subdomains"
}

// Configure configures the data source with the provider client
func (d *SubdomainsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *SubdomainsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "subdomains data source",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, Description: "Data source identifier"},
		},
	}
}

// Read reads the subdomains data
func (d *SubdomainsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config subdomainsDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /subdomains/{domain}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
