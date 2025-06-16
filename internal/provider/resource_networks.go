package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NetworksResource implements the networks resource
type NetworksResource struct {
	client *GeneratedClient
}

// networksModel represents the resource data model
type networksModel struct {
	ID           types.String `tfsdk:"id"`
	OriginID     types.Int64  `tfsdk:"origin_id"`
	Name         types.String `tfsdk:"name"`
	IPAddress    types.String `tfsdk:"ip_address"`
	PrefixLength types.Int64  `tfsdk:"prefix_length"`
	IsDynamic    types.Bool   `tfsdk:"is_dynamic"`
	IsVerified   types.Bool   `tfsdk:"is_verified"`
	Status       types.String `tfsdk:"status"`
	CreatedAt    types.String `tfsdk:"created_at"`
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

	// Create enhanced client
	generatedClient, err := NewGeneratedClient(context.Background(), client.key, client.secret, client.orgID)
	if err != nil {
		resp.Diagnostics.AddError("Failed to create generated client", err.Error())
		return
	}

	r.client = generatedClient
}

// Schema defines the resource schema
func (r *NetworksResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages networks in Cisco Umbrella. Networks represent IP address ranges that are associated with your organization.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for this network (same as origin_id)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"origin_id": schema.Int64Attribute{
				Computed:    true,
				Description: "The origin ID of the network",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the network (1-50 characters)",
			},
			"ip_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The IP address of the network. Required for static networks, optional for dynamic networks",
			},
			"prefix_length": schema.Int64Attribute{
				Required:    true,
				Description: "The length of the prefix. Must be between 29 and 32",
			},
			"is_dynamic": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
				Description: "Specifies whether the IP address is dynamic",
			},
			"is_verified": schema.BoolAttribute{
				Computed:    true,
				Description: "Specifies whether the network is verified",
			},
			"status": schema.StringAttribute{
				Required:    true,
				Description: "The status of the network (OPEN or CLOSED)",
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: "The date and time when the network was created",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

// Create creates a new network
func (r *NetworksResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan networksModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating network", map[string]interface{}{
		"name":          plan.Name.ValueString(),
		"prefix_length": plan.PrefixLength.ValueInt64(),
		"is_dynamic":    plan.IsDynamic.ValueBool(),
		"status":        plan.Status.ValueString(),
	})

	// Validate required fields and constraints
	if err := r.validateNetworkData(plan); err != nil {
		resp.Diagnostics.AddError("Validation Error", err.Error())
		return
	}

	// Prepare request payload
	payload := map[string]interface{}{
		"name":         plan.Name.ValueString(),
		"prefixLength": plan.PrefixLength.ValueInt64(),
		"isDynamic":    plan.IsDynamic.ValueBool(),
		"status":       plan.Status.ValueString(),
	}

	// Add IP address if provided (required for static networks)
	if !plan.IPAddress.IsNull() && !plan.IPAddress.IsUnknown() {
		payload["ipAddress"] = plan.IPAddress.ValueString()
	}

	// Make API call to create network
	result, err := r.client.CreateNetwork(ctx, payload)
	if err != nil {
		resp.Diagnostics.AddError("Failed to create network", err.Error())
		return
	}

	// Update plan with response data
	r.updateModelFromAPIResponse(&plan, result)

	tflog.Debug(ctx, "Network created successfully", map[string]interface{}{
		"id":        plan.ID.ValueString(),
		"origin_id": plan.OriginID.ValueInt64(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Read reads the network resource
func (r *NetworksResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state networksModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading network", map[string]interface{}{
		"id": state.ID.ValueString(),
	})

	// Make API call to get network
	result, err := r.client.GetNetwork(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to read network", err.Error())
		return
	}

	// Update state with response data
	r.updateModelFromAPIResponse(&state, result)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update updates the network
func (r *NetworksResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan networksModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating network", map[string]interface{}{
		"id":   plan.ID.ValueString(),
		"name": plan.Name.ValueString(),
	})

	// Validate required fields and constraints
	if err := r.validateNetworkData(plan); err != nil {
		resp.Diagnostics.AddError("Validation Error", err.Error())
		return
	}

	// Prepare request payload
	payload := map[string]interface{}{
		"name":      plan.Name.ValueString(),
		"isDynamic": plan.IsDynamic.ValueBool(),
		"status":    plan.Status.ValueString(),
	}

	// Add optional fields if provided
	if !plan.IPAddress.IsNull() && !plan.IPAddress.IsUnknown() {
		payload["ipAddress"] = plan.IPAddress.ValueString()
	}
	if !plan.PrefixLength.IsNull() && !plan.PrefixLength.IsUnknown() {
		payload["prefixLength"] = plan.PrefixLength.ValueInt64()
	}

	// Make API call to update network
	result, err := r.client.UpdateNetwork(ctx, plan.ID.ValueString(), payload)
	if err != nil {
		resp.Diagnostics.AddError("Failed to update network", err.Error())
		return
	}

	// Update plan with response data
	r.updateModelFromAPIResponse(&plan, result)

	tflog.Debug(ctx, "Network updated successfully", map[string]interface{}{
		"id": plan.ID.ValueString(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Delete deletes the network
func (r *NetworksResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state networksModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting network", map[string]interface{}{
		"id": state.ID.ValueString(),
	})

	// Make API call to delete network
	err := r.client.DeleteNetwork(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to delete network", err.Error())
		return
	}

	tflog.Debug(ctx, "Network deleted successfully", map[string]interface{}{
		"id": state.ID.ValueString(),
	})

	// Clear cache for this resource
	r.client.clearCacheForPath(fmt.Sprintf("/deployments/v2/networks/%s", state.ID.ValueString()))
}

// validateNetworkData validates the network configuration
func (r *NetworksResource) validateNetworkData(model networksModel) error {
	// Validate name length
	name := model.Name.ValueString()
	if len(name) < 1 || len(name) > 50 {
		return fmt.Errorf("network name must be between 1 and 50 characters")
	}

	// Validate prefix length
	prefixLength := model.PrefixLength.ValueInt64()
	if prefixLength < 29 || prefixLength > 32 {
		return fmt.Errorf("prefix length must be between 29 and 32")
	}

	// Validate status
	status := model.Status.ValueString()
	if status != "OPEN" && status != "CLOSED" {
		return fmt.Errorf("status must be either 'OPEN' or 'CLOSED'")
	}

	// Validate IP address for static networks
	if !model.IsDynamic.ValueBool() {
		if model.IPAddress.IsNull() || model.IPAddress.IsUnknown() || model.IPAddress.ValueString() == "" {
			return fmt.Errorf("ip_address is required for static networks (is_dynamic = false)")
		}

		ipAddress := model.IPAddress.ValueString()
		if len(ipAddress) < 7 {
			return fmt.Errorf("ip_address must be at least 7 characters long")
		}
	}

	return nil
}

// updateModelFromAPIResponse updates the model with data from API response
func (r *NetworksResource) updateModelFromAPIResponse(model *networksModel, result map[string]interface{}) {
	if originID, ok := result["originId"].(float64); ok {
		model.OriginID = types.Int64Value(int64(originID))
		model.ID = types.StringValue(strconv.FormatInt(int64(originID), 10))
	}

	if name, ok := result["name"].(string); ok {
		model.Name = types.StringValue(name)
	}

	if ipAddress, ok := result["ipAddress"].(string); ok {
		model.IPAddress = types.StringValue(ipAddress)
	}

	if prefixLength, ok := result["prefixLength"].(float64); ok {
		model.PrefixLength = types.Int64Value(int64(prefixLength))
	}

	if isDynamic, ok := result["isDynamic"].(bool); ok {
		model.IsDynamic = types.BoolValue(isDynamic)
	}

	if isVerified, ok := result["isVerified"].(bool); ok {
		model.IsVerified = types.BoolValue(isVerified)
	}

	if status, ok := result["status"].(string); ok {
		model.Status = types.StringValue(status)
	}

	if createdAt, ok := result["createdAt"].(string); ok {
		model.CreatedAt = types.StringValue(createdAt)
	}
}
