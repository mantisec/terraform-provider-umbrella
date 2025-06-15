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

// RiskScoreDataSource implements the risk_score data source
type RiskScoreDataSource struct {
	client *apiClient
}

// risk_scoreDataModel represents the data source data model
type risk_scoreDataModel struct {
	Id types.String `tfsdk:"id"`
}

// NewRiskScoreDataSource creates a new risk_score data source
func NewRiskScoreDataSource() datasource.DataSource {
	return &RiskScoreDataSource{}
}

// Metadata returns the data source type name
func (d *RiskScoreDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_risk_score"
}

// Configure configures the data source with the provider client
func (d *RiskScoreDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *RiskScoreDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "risk_score data source",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, Description: "Data source identifier"},
		},
	}
}

// Read reads the risk_score data
func (d *RiskScoreDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config risk_scoreDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /domains/risk-score/{domain}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
