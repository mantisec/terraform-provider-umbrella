//go:build gen
// +build gen

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	_ "github.com/mantisec/terraform-provider-umbrella/generated"
)

// Ensure the implementation satisfies the expected interfaces.
var _ provider.Provider = &umbrellaProvider{}

// New instantiates the provider (entry‑point called by Terraform).
func New() provider.Provider { return &umbrellaProvider{} }

type umbrellaProvider struct{}

type umbrellaProviderModel struct {
	ApiKey    types.String `tfsdk:"api_key"`
	ApiSecret types.String `tfsdk:"api_secret"`
	OrgID     types.String `tfsdk:"org_id"`
}

func (p *umbrellaProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "umbrella"
}

func (p *umbrellaProvider) GetSchema(_ context.Context) (provider.Schema, diag.Diagnostics) {
	return provider.Schema{
		Attributes: map[string]provider.Attribute{
			"api_key":    provider.StringAttribute{Required: true, Description: "Umbrella API key (client ID)."},
			"api_secret": provider.StringAttribute{Required: true, Sensitive: true, Description: "Umbrella API secret (client secret)."},
			"org_id":     provider.StringAttribute{Required: true, Description: "Umbrella organization ID."},
		},
	}, nil
}

func (p *umbrellaProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var conf umbrellaProviderModel
	diags := req.Config.Get(ctx, &conf)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, err := newAPIClient(ctx, conf.ApiKey.ValueString(), conf.ApiSecret.ValueString(), conf.OrgID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Unable to authenticate", err.Error())
		return
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *umbrellaProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}

func (p *umbrellaProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}
