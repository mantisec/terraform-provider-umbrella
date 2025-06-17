package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// UsersResource implements the users resource
type UsersResource struct {
	client *apiClient
}

// usersModel represents the resource data model
type usersModel struct {
	Id types.String `tfsdk:"id"`
}

// NewUsersResource creates a new users resource
func NewUsersResource() resource.Resource {
	return &UsersResource{}
}

// Metadata returns the resource type name
func (r *UsersResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "umbrella_users"
}

// Configure configures the resource with the provider client
func (r *UsersResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *UsersResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "users resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, Description: "Resource identifier"},
		},
	}
}

// Create creates a new users
func (r *UsersResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan usersModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement create logic using POST /users

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Delete deletes the users
func (r *UsersResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state usersModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement delete logic using DELETE /users/{userId}
}
