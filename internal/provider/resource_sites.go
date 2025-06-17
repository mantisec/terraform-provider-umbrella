package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SitesResource implements the sites resource
type SitesResource struct {
	client *apiClient
}

// sitesModel represents the resource data model
type sitesModel struct {
	Id types.String `tfsdk:"id"`
}

// NewSitesResource creates a new sites resource
func NewSitesResource() resource.Resource {
	return &SitesResource{}
}

// Metadata returns the resource type name
func (r *SitesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "umbrella_sites"
}

// Configure configures the resource with the provider client
func (r *SitesResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*apiClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", "Expected *apiClient")
		return
	}

	r.client = client
}

// Schema defines the resource schema
func (r *SitesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "sites resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, Description: "Resource identifier"},
		},
	}
}

// Create creates a new sites
func (r *SitesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan sitesModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement create logic using POST /sites

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Update updates the sites
func (r *SitesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan sitesModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement update logic using PUT /sites/{siteId}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Delete deletes the sites
func (r *SitesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state sitesModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement delete logic using DELETE /sites/{siteId}
}
