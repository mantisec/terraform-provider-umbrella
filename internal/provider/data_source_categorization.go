package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// CategorizationDataSource implements the categorization data source
type CategorizationDataSource struct {
	client *apiClient
}

// categorizationDataModel represents the data source data model
type categorizationDataModel struct {
	Id types.String `tfsdk:"id"`
}

// NewCategorizationDataSource creates a new categorization data source
func NewCategorizationDataSource() datasource.DataSource {
	return &CategorizationDataSource{}
}

// Metadata returns the data source type name
func (d *CategorizationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_categorization"
}

// Configure configures the data source with the provider client
func (d *CategorizationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *CategorizationDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "categorization data source",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, Description: "Data source identifier"},
		},
	}
}

// Read reads the categorization data
func (d *CategorizationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config categorizationDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /domains/categorization/{domain}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
