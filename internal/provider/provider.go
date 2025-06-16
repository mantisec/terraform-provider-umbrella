package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	pschema "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// -----------------------------------------------------------------------------
// Provider definition
// -----------------------------------------------------------------------------

type providerModel struct {
	APIKey    types.String `tfsdk:"api_key"`
	APISecret types.String `tfsdk:"api_secret"`
	OrgID     types.String `tfsdk:"org_id"`
}

type umbrellaProvider struct{ client *apiClient }

func NewProvider() provider.Provider { return &umbrellaProvider{} }

// -----------------------------------------------------------------------------
// Provider metadata & schema
// -----------------------------------------------------------------------------
func (p *umbrellaProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "umbrella"
}

func (p *umbrellaProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = pschema.Schema{
		Description: "Provider for Cisco Umbrella Secure Web Gateway REST API.",
		Attributes: map[string]pschema.Attribute{
			"api_key":    pschema.StringAttribute{Required: true, Sensitive: true, Description: "Umbrella API key (client ID)."},
			"api_secret": pschema.StringAttribute{Required: true, Sensitive: true, Description: "Umbrella API secret (client secret)."},
			"org_id":     pschema.StringAttribute{Required: true, Description: "Umbrella organisation ID."},
		},
	}
}

// -----------------------------------------------------------------------------
// Provider configuration – create API client
// -----------------------------------------------------------------------------
func (p *umbrellaProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var cfg providerModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &cfg)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, err := newAPIClient(ctx, cfg.APIKey.ValueString(), cfg.APISecret.ValueString(), cfg.OrgID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Unable to authenticate", err.Error())
		return
	}
	p.client = client
	resp.ResourceData = client
	resp.DataSourceData = client
}

// -----------------------------------------------------------------------------
// Provider resources & data-sources
// -----------------------------------------------------------------------------
func (p *umbrellaProvider) Resources(_ context.Context) []func() resource.Resource {
	// Initialize generated resources
	InitializeGeneratedResources()

	// Start with manually created resources (add as they are implemented)
	resources := []func() resource.Resource{
		// TODO: Add manually created resources here as they are implemented
		// NewDestinationListResource,
		// NewTunnelResource,
	}

	// Add generated resources
	resources = append(resources, GetGeneratedRegistry().GetResources()...)

	return resources
}

func (p *umbrellaProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	// Initialize generated resources (includes data sources)
	InitializeGeneratedResources()

	// Start with manually created data sources (add as they are implemented)
	dataSources := []func() datasource.DataSource{
		// TODO: Add manually created data sources here as they are implemented
	}

	// Add generated data sources
	dataSources = append(dataSources, GetGeneratedRegistry().GetDataSources()...)

	return dataSources
}
