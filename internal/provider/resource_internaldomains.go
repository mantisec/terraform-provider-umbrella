package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InternaldomainsResource implements the internaldomains resource
type InternaldomainsResource struct {
	client *apiClient
}

// internaldomainsModel represents the resource data model
type internaldomainsModel struct {
	Id types.String `tfsdk:"id"`
}

// NewInternaldomainsResource creates a new internaldomains resource
func NewInternaldomainsResource() resource.Resource {
	return &InternaldomainsResource{}
}

// Metadata returns the resource type name
func (r *InternaldomainsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "umbrella_internaldomains"
}

// Configure configures the resource with the provider client
func (r *InternaldomainsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *InternaldomainsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "internaldomains resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, Description: "Resource identifier"},
		},
	}
}

// Create creates a new internaldomains
func (r *InternaldomainsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan internaldomainsModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement create logic using DELETE /internaldomains/{internalDomainId}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Read reads the internaldomains resource
func (r *InternaldomainsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state internaldomainsModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /internaldomains/{internalDomainId}

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

// Update updates the internaldomains
func (r *InternaldomainsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan internaldomainsModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement update logic using DELETE /internaldomains/{internalDomainId}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Delete deletes the internaldomains
func (r *InternaldomainsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state internaldomainsModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement delete logic using DELETE /internaldomains/{internalDomainId}
}
