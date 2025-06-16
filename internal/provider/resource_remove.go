package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// RemoveResource implements the remove resource
type RemoveResource struct {
	client *apiClient
}

// removeModel represents the resource data model
type removeModel struct {
	Id types.String `tfsdk:"id"`
}

// NewRemoveResource creates a new remove resource
func NewRemoveResource() resource.Resource {
	return &RemoveResource{}
}

// Metadata returns the resource type name
func (r *RemoveResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "umbrella_remove"
}

// Configure configures the resource with the provider client
func (r *RemoveResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *RemoveResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "remove resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, Description: "Resource identifier"},
		},
	}
}

// Create creates a new remove
func (r *RemoveResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan removeModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement create logic using DELETE /destinationlists/{destinationListId}/destinations/remove

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Read reads the remove resource
func (r *RemoveResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state removeModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /destinationlists/{destinationListId}/destinations/remove

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

// Update updates the remove resource
func (r *RemoveResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan removeModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement update logic using PUT /destinationlists/{destinationListId}/destinations/remove

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Delete deletes the remove
func (r *RemoveResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state removeModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement delete logic using DELETE /destinationlists/{destinationListId}/destinations/remove
}
