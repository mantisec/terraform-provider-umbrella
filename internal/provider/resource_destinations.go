package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// DestinationsResource implements the destinations resource
type DestinationsResource struct {
	client *apiClient
}

// destinationsModel represents the resource data model
type destinationsModel struct {
	Id types.String `tfsdk:"id"`
}

// NewDestinationsResource creates a new destinations resource
func NewDestinationsResource() resource.Resource {
	return &DestinationsResource{}
}

// Metadata returns the resource type name
func (r *DestinationsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "umbrella_destinations"
}

// Configure configures the resource with the provider client
func (r *DestinationsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *DestinationsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "destinations resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, Description: "Resource identifier"},
		},
	}
}

// Create creates a new destinations
func (r *DestinationsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan destinationsModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement create logic using POST /destinationlists/{destinationListId}/destinations

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Read reads the destinations resource
func (r *DestinationsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state destinationsModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /destinationlists/{destinationListId}/destinations

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

// Update updates the destinations resource
func (r *DestinationsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan destinationsModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement update logic using PUT /destinationlists/{destinationListId}/destinations

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Delete deletes the destinations resource
func (r *DestinationsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state destinationsModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement delete logic using DELETE /destinationlists/{destinationListId}/destinations
}
