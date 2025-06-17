package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NetworksResource implements the networks resource
type NetworksResource struct {
	client *apiClient
}

// networksModel represents the resource data model
type networksModel struct {
	Id         types.String `tfsdk:"id"`
	Statuscode types.Int64  `tfsdk:"statusCode"`
	Error      types.String `tfsdk:"error"`
	Message    types.String `tfsdk:"message"`
}

// NewNetworksResource creates a new networks resource
func NewNetworksResource() resource.Resource {
	return &NetworksResource{}
}

// Metadata returns the resource type name
func (r *NetworksResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "umbrella_networks"
}

// Configure configures the resource with the provider client
func (r *NetworksResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *NetworksResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "networks resource",
		Attributes: map[string]schema.Attribute{
			"id":         schema.StringAttribute{Computed: true, Description: "Resource identifier"},
			"statusCode": schema.Int64Attribute{Computed: true, Description: "HTTP status code"},
			"error":      schema.StringAttribute{Computed: true, Description: "Shorthand error message"},
			"message":    schema.StringAttribute{Computed: true, Description: "Detailed error message"},
		},
	}
}

// Create creates a new networks
func (r *NetworksResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan networksModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement create logic using POST /networks

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Update updates the networks
func (r *NetworksResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan networksModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement update logic using PUT /networks/{networkId}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Delete deletes the networks
func (r *NetworksResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state networksModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement delete logic using DELETE /networks/{networkId}
}
