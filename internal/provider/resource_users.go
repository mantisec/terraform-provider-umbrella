package provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UsersResource implements the users resource
type UsersResource struct {
	client *GeneratedClient
}

// usersModel represents the resource data model
type usersModel struct {
	ID               types.String `tfsdk:"id"`
	UserID           types.Int64  `tfsdk:"user_id"`
	FirstName        types.String `tfsdk:"firstname"`
	LastName         types.String `tfsdk:"lastname"`
	Email            types.String `tfsdk:"email"`
	Password         types.String `tfsdk:"password"`
	RoleID           types.Int64  `tfsdk:"role_id"`
	Role             types.String `tfsdk:"role"`
	Timezone         types.String `tfsdk:"timezone"`
	Status           types.String `tfsdk:"status"`
	LastLoginTime    types.String `tfsdk:"last_login_time"`
	TwoFactorEnabled types.Bool   `tfsdk:"two_factor_enabled"`
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

	// Create enhanced client
	generatedClient, err := NewGeneratedClient(context.Background(), client.key, client.secret, client.orgID)
	if err != nil {
		resp.Diagnostics.AddError("Failed to create generated client", err.Error())
		return
	}

	r.client = generatedClient
}

// Schema defines the resource schema
func (r *UsersResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages user accounts in Cisco Umbrella. Users represent individual accounts with specific roles and permissions within your organization.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for this user (same as user_id)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"user_id": schema.Int64Attribute{
				Computed:    true,
				Description: "The numeric user ID",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"firstname": schema.StringAttribute{
				Required:    true,
				Description: "The user's first name (minimum 1 character)",
			},
			"lastname": schema.StringAttribute{
				Required:    true,
				Description: "The user's last name (minimum 1 character)",
			},
			"email": schema.StringAttribute{
				Required:    true,
				Description: "The user's email address (must be unique)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"password": schema.StringAttribute{
				Required:    true,
				Sensitive:   true,
				Description: "The user's password (minimum 1 character)",
			},
			"role_id": schema.Int64Attribute{
				Required:    true,
				Description: "The role ID to assign to the user (minimum 1)",
			},
			"role": schema.StringAttribute{
				Computed:    true,
				Description: "The user's role name",
			},
			"timezone": schema.StringAttribute{
				Required:    true,
				Description: "The user's timezone (minimum 1 character)",
			},
			"status": schema.StringAttribute{
				Computed:    true,
				Description: "The user's status (e.g., 'on', 'off')",
			},
			"last_login_time": schema.StringAttribute{
				Computed:    true,
				Description: "The user's last login date and time (ISO8601 timestamp)",
			},
			"two_factor_enabled": schema.BoolAttribute{
				Computed:    true,
				Description: "Specifies whether two-factor authentication is enabled",
			},
		},
	}
}

// Create creates a new user
func (r *UsersResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan usersModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating user", map[string]interface{}{
		"email":     plan.Email.ValueString(),
		"firstname": plan.FirstName.ValueString(),
		"lastname":  plan.LastName.ValueString(),
		"role_id":   plan.RoleID.ValueInt64(),
		"timezone":  plan.Timezone.ValueString(),
	})

	// Validate required fields and constraints
	if err := r.validateUserData(plan); err != nil {
		resp.Diagnostics.AddError("Validation Error", err.Error())
		return
	}

	// Prepare request payload
	payload := map[string]interface{}{
		"firstname": plan.FirstName.ValueString(),
		"lastname":  plan.LastName.ValueString(),
		"email":     plan.Email.ValueString(),
		"password":  plan.Password.ValueString(),
		"roleId":    plan.RoleID.ValueInt64(),
		"timezone":  plan.Timezone.ValueString(),
	}

	// Make API call to create user
	result, err := r.client.CreateUser(ctx, payload)
	if err != nil {
		resp.Diagnostics.AddError("Failed to create user", err.Error())
		return
	}

	// Update plan with response data
	r.updateModelFromAPIResponse(&plan, result)

	tflog.Debug(ctx, "User created successfully", map[string]interface{}{
		"id":      plan.ID.ValueString(),
		"user_id": plan.UserID.ValueInt64(),
		"email":   plan.Email.ValueString(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Read reads the user resource
func (r *UsersResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state usersModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading user", map[string]interface{}{
		"id": state.ID.ValueString(),
	})

	// Make API call to get user
	result, err := r.client.GetUser(ctx, state.ID.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "404") {
			// User was deleted outside of Terraform
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Failed to read user", err.Error())
		return
	}

	// Update state with response data
	r.updateModelFromAPIResponse(&state, result)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update updates the user
func (r *UsersResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan usersModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating user", map[string]interface{}{
		"id":        plan.ID.ValueString(),
		"firstname": plan.FirstName.ValueString(),
		"lastname":  plan.LastName.ValueString(),
	})

	// Note: The Users API doesn't have a PUT endpoint for updates
	// According to the API spec, users can only be created and deleted
	// For now, we'll return an error indicating updates are not supported
	resp.Diagnostics.AddError(
		"Update Not Supported",
		"The Umbrella Users API does not support updating existing users. "+
			"To modify a user, you must delete and recreate the user account. "+
			"This is a limitation of the Umbrella API, not the Terraform provider.",
	)
}

// Delete deletes the user
func (r *UsersResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state usersModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting user", map[string]interface{}{
		"id": state.ID.ValueString(),
	})

	// Make API call to delete user
	err := r.client.DeleteUser(ctx, state.ID.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "404") {
			// User was already deleted
			tflog.Debug(ctx, "User already deleted", map[string]interface{}{
				"id": state.ID.ValueString(),
			})
			return
		}
		resp.Diagnostics.AddError("Failed to delete user", err.Error())
		return
	}

	tflog.Debug(ctx, "User deleted successfully", map[string]interface{}{
		"id": state.ID.ValueString(),
	})

	// Clear cache for this resource
	r.client.clearCacheForPath(fmt.Sprintf("/admin/v2/users/%s", state.ID.ValueString()))
}

// validateUserData validates the user configuration
func (r *UsersResource) validateUserData(model usersModel) error {
	// Validate first name
	firstName := model.FirstName.ValueString()
	if len(strings.TrimSpace(firstName)) < 1 {
		return fmt.Errorf("firstname must be at least 1 character long")
	}

	// Validate last name
	lastName := model.LastName.ValueString()
	if len(strings.TrimSpace(lastName)) < 1 {
		return fmt.Errorf("lastname must be at least 1 character long")
	}

	// Validate email
	email := model.Email.ValueString()
	if len(strings.TrimSpace(email)) < 1 {
		return fmt.Errorf("email must be at least 1 character long")
	}
	if !strings.Contains(email, "@") {
		return fmt.Errorf("email must be a valid email address")
	}

	// Validate password
	password := model.Password.ValueString()
	if len(password) < 1 {
		return fmt.Errorf("password must be at least 1 character long")
	}

	// Validate role ID
	roleID := model.RoleID.ValueInt64()
	if roleID < 1 {
		return fmt.Errorf("role_id must be at least 1")
	}

	// Validate timezone
	timezone := model.Timezone.ValueString()
	if len(strings.TrimSpace(timezone)) < 1 {
		return fmt.Errorf("timezone must be at least 1 character long")
	}

	return nil
}

// updateModelFromAPIResponse updates the model with data from API response
func (r *UsersResource) updateModelFromAPIResponse(model *usersModel, result map[string]interface{}) {
	if id, ok := result["id"].(float64); ok {
		model.UserID = types.Int64Value(int64(id))
		model.ID = types.StringValue(strconv.FormatInt(int64(id), 10))
	}

	if firstName, ok := result["firstname"].(string); ok {
		model.FirstName = types.StringValue(firstName)
	}

	if lastName, ok := result["lastname"].(string); ok {
		model.LastName = types.StringValue(lastName)
	}

	if email, ok := result["email"].(string); ok {
		model.Email = types.StringValue(email)
	}

	if roleID, ok := result["roleId"].(float64); ok {
		model.RoleID = types.Int64Value(int64(roleID))
	}

	if role, ok := result["role"].(string); ok {
		model.Role = types.StringValue(role)
	}

	if timezone, ok := result["timezone"].(string); ok {
		model.Timezone = types.StringValue(timezone)
	}

	if status, ok := result["status"].(string); ok {
		model.Status = types.StringValue(status)
	}

	if lastLoginTime, ok := result["lastLoginTime"].(string); ok {
		model.LastLoginTime = types.StringValue(lastLoginTime)
	}

	if twoFactorEnabled, ok := result["twoFactorEnable"].(bool); ok {
		model.TwoFactorEnabled = types.BoolValue(twoFactorEnabled)
	}
}
