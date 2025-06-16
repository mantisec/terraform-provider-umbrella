package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PasswordresetsResource implements the passwordresets resource
type PasswordresetsResource struct {
	client *apiClient
}

// passwordresetsModel represents the resource data model
type passwordresetsModel struct {
	Id types.String `tfsdk:"id"`
}

// NewPasswordresetsResource creates a new passwordresets resource
func NewPasswordresetsResource() resource.Resource {
	return &PasswordresetsResource{}
}

// Metadata returns the resource type name
func (r *PasswordresetsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "umbrella_passwordresets"
}

// Configure configures the resource with the provider client
func (r *PasswordresetsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *PasswordresetsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "passwordresets resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, Description: "Resource identifier"},
		},
	}
}

// Create creates a new passwordresets
func (r *PasswordresetsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan passwordresetsModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement create logic using POST /passwordResets/{customerId}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Read reads the passwordresets resource
func (r *PasswordresetsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state passwordresetsModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement read logic using GET /passwordResets/{customerId}

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

// Update updates the passwordresets resource
func (r *PasswordresetsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan passwordresetsModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement update logic using PUT /passwordResets/{customerId}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Delete deletes the passwordresets resource
func (r *PasswordresetsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state passwordresetsModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement delete logic using DELETE /passwordResets/{customerId}
}
