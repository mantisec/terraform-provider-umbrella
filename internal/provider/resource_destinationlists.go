package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// DestinationlistsResource implements the destinationlists resource
type DestinationlistsResource struct {
	client *apiClient
}

// destinationlistsModel represents the resource data model
type destinationlistsModel struct {
	Id types.String `tfsdk:"id"`
}

// NewDestinationlistsResource creates a new destinationlists resource
func NewDestinationlistsResource() resource.Resource {
	return &DestinationlistsResource{}
}

// Metadata returns the resource type name
func (r *DestinationlistsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "umbrella_destinationlists"
}

// Configure configures the resource with the provider client
func (r *DestinationlistsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *DestinationlistsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "destinationlists resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, Description: "Resource identifier"},
		},
	}
}

// Create creates a new destinationlists
func (r *DestinationlistsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan destinationlistsModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement create logic using PATCH /destinationlists/{destinationListId}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Update updates the destinationlists
func (r *DestinationlistsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan destinationlistsModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement update logic using PATCH /destinationlists/{destinationListId}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Delete deletes the destinationlists
func (r *DestinationlistsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state destinationlistsModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement delete logic using PATCH /destinationlists/{destinationListId}
}
