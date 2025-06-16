package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ArtifactsDataSource implements the artifacts data source
type ArtifactsDataSource struct {
	client *apiClient
}

// artifactsDataModel represents the data source data model
type artifactsDataModel struct {
	Id types.String `tfsdk:"id"`
}

// NewArtifactsDataSource creates a new artifacts data source
func NewArtifactsDataSource() datasource.DataSource {
	return &ArtifactsDataSource{}
}

// Metadata returns the data source type name
func (d *ArtifactsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_artifacts"
}

// Configure configures the data source with the provider client
func (d *ArtifactsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *ArtifactsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "artifacts data source",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, Description: "Data source identifier"},
		},
	}
}

// Read reads the artifacts data
func (d *ArtifactsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config artifactsDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /sample/{hash}/artifacts

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
