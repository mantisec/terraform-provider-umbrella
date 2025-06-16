package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InternalnetworksResource implements the internalnetworks resource
type InternalnetworksResource struct {
	client *apiClient
}

// internalnetworksModel represents the resource data model
type internalnetworksModel struct {
	Id types.String `tfsdk:"id"`
}

// NewInternalnetworksResource creates a new internalnetworks resource
func NewInternalnetworksResource() resource.Resource {
	return &InternalnetworksResource{}
}

// Metadata returns the resource type name
func (r *InternalnetworksResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "umbrella_internalnetworks"
}

// Configure configures the resource with the provider client
func (r *InternalnetworksResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *InternalnetworksResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "internalnetworks resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, Description: "Resource identifier"},
		},
	}
}

// Create creates a new internalnetworks
func (r *InternalnetworksResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan internalnetworksModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement create logic using DELETE /internalnetworks/{internalNetworkId}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Update updates the internalnetworks
func (r *InternalnetworksResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan internalnetworksModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement update logic using DELETE /internalnetworks/{internalNetworkId}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Delete deletes the internalnetworks
func (r *InternalnetworksResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state internalnetworksModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement delete logic using DELETE /internalnetworks/{internalNetworkId}
}
