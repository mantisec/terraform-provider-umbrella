package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NameDataSource implements the name data source
type NameDataSource struct {
	client *apiClient
}

// nameDataModel represents the data source data model
type nameDataModel struct {
	Id                     types.String `tfsdk:"id"`
	Pfs2                   types.Set    `tfsdk:"pfs2"`
	Found                  types.Bool   `tfsdk:"found"`
	Tb1                    types.Set    `tfsdk:"tb1"`
	Popularity             types.Int64  `tfsdk:"popularity"`
	Geodiversity           types.Set    `tfsdk:"geodiversity"`
	Attack                 types.String `tfsdk:"attack"`
	Entropy                types.Int64  `tfsdk:"entropy"`
	AsnScore               types.Int64  `tfsdk:"asn_score"`
	PrefixScore            types.Int64  `tfsdk:"prefix_score"`
	RipScore               types.Int64  `tfsdk:"rip_score"`
	Perplexity             types.Int64  `tfsdk:"perplexity"`
	Securerank2            types.Int64  `tfsdk:"securerank2"`
	KsTest                 types.Int64  `tfsdk:"ks_test"`
	Geoscore               types.Int64  `tfsdk:"geoscore"`
	ThreatType             types.String `tfsdk:"threat_type"`
	DgaScore               types.Int64  `tfsdk:"dga_score"`
	Pagerank               types.Int64  `tfsdk:"pagerank"`
	GeodiversityNormalized types.Set    `tfsdk:"geodiversity_normalized"`
	TldGeodiversity        types.Set    `tfsdk:"tld_geodiversity"`
}

// NewNameDataSource creates a new name data source
func NewNameDataSource() datasource.DataSource {
	return &NameDataSource{}
}

// Metadata returns the data source type name
func (d *NameDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "umbrella_name"
}

// Configure configures the data source with the provider client
func (d *NameDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *NameDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "name data source",
		Attributes: map[string]schema.Attribute{
			"id":                      schema.StringAttribute{Computed: true, Description: "Data source identifier"},
			"pfs2":                    schema.SetAttribute{Computed: true, Description: "The list of the co-occurring domains.", ElementType: types.StringType},
			"found":                   schema.BoolAttribute{Computed: true, Description: "Specify whether the domain is co-occurring."},
			"tb1":                     schema.SetAttribute{Computed: true, Description: "The list of domain name and score pairs where score is the number of client IP requests to the site around the same time that the site is looked up.", ElementType: types.StringType},
			"popularity":              schema.Int64Attribute{Computed: true, Description: "The number of unique client IPs visiting this site, relative to all requests to all sites. A score of how many different client or unique IPs requested to this domain compared to others."},
			"geodiversity":            schema.SetAttribute{Computed: true, Description: "The list of scores that represent the number of queries from clients visiting the domain, broken down by country.", ElementType: types.StringType},
			"attack":                  schema.StringAttribute{Computed: true, Description: "The name of any known attacks associated with this domain. Returns an empty string if no known threat associated with domain."},
			"entropy":                 schema.Int64Attribute{Computed: true, Description: "The number of bits required to encode the domain name as a score. This score is used in conjunction with DGA and Perplexity."},
			"asn_score":               schema.Int64Attribute{Computed: true, Description: "The ASN reputation score ranges from -100 to 0 where -100 is very suspicious."},
			"prefix_score":            schema.Int64Attribute{Computed: true, Description: "The prefix ranks domains given their IP prefixes (an IP prefix is the first three octets in an IP address) and the reputation score of these prefixes. The scores range from -100 to 0 where -100 is very suspicious."},
			"rip_score":               schema.Int64Attribute{Computed: true, Description: "The RIP ranks domains given their IP addresses and the reputation score of these IP addresses. The scores ranges from -100 to 0 where -100 is very suspicious."},
			"perplexity":              schema.Int64Attribute{Computed: true, Description: "A second score on the likeliness of the name to be algorithmically generated, on a scale from 0 to 100. This score is used in conjunction with DGA."},
			"securerank2":             schema.Int64Attribute{Computed: true, Description: "The suspicious rank for a domain that reviews base on the lookup behavior of client IP for the domain. Securerank is designed to identify hostnames requested by known infected clients but never requested by clean clients, assuming these domains are more likely to be bad. Scores returned range from -100 (suspicious) to 100 (benign)."},
			"ks_test":                 schema.Int64Attribute{Computed: true, Description: "A number that represents the Kolmogorov-Smirnov test on geodiversity. Zero indicates that the client traffic matches what is expected for this top-level domain."},
			"geoscore":                schema.Int64Attribute{Computed: true, Description: "A score that represents how far the different physical locations serving this name are from each other."},
			"threat_type":             schema.StringAttribute{Computed: true, Description: "The type of the known attack, such as botnet or APT. Returns an empty string if no known threat associated with domain."},
			"dga_score":               schema.Int64Attribute{Computed: true, Description: "A domain generation algorithm (DGA) is used by malware to generate large lists of domain names. This score is created based on the likeliness of the domain name being generated by an algorithm rather than a human. This algorithm is designed to identify domains which have been created using an automated randomization strategy, which is a common evasion technique in malware kits or botnets. This score ranges from -100 (suspicious) to 0 (benign)."},
			"pagerank":                schema.Int64Attribute{Computed: true, Description: "A popularity score according to Google's PageRank algorithm."},
			"geodiversity_normalized": schema.SetAttribute{Computed: true, Description: "The list of scores that represents the amount of queries for clients visiting the domain, broken down by country.", ElementType: types.StringType},
			"tld_geodiversity":        schema.SetAttribute{Computed: true, Description: "The list of scores that represent the top-level domain country code geodiversity as a percentage of clients visiting the domain.", ElementType: types.StringType},
		},
	}
}

// Read reads the name data
func (d *NameDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config nameDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /pdns/name/{domain}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
