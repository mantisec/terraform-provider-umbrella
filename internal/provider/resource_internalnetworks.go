package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// InternalnetworksResource implements the internalnetworks resource
type InternalnetworksResource struct {
	client *GeneratedClient
}

// internalnetworksModel represents the resource data model
type internalnetworksModel struct {
	ID           types.String `tfsdk:"id"`
	OriginID     types.Int64  `tfsdk:"origin_id"`
	Name         types.String `tfsdk:"name"`
	IPAddress    types.String `tfsdk:"ip_address"`
	PrefixLength types.Int64  `tfsdk:"prefix_length"`
	SiteID       types.Int64  `tfsdk:"site_id"`
	SiteName     types.String `tfsdk:"site_name"`
	NetworkID    types.Int64  `tfsdk:"network_id"`
	NetworkName  types.String `tfsdk:"network_name"`
	TunnelID     types.Int64  `tfsdk:"tunnel_id"`
	TunnelName   types.String `tfsdk:"tunnel_name"`
	CreatedAt    types.String `tfsdk:"created_at"`
	ModifiedAt   types.String `tfsdk:"modified_at"`
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

	// Create enhanced client
	generatedClient, err := NewGeneratedClient(context.Background(), client.key, client.secret, client.orgID)
	if err != nil {
		resp.Diagnostics.AddError("Failed to create generated client", err.Error())
		return
	}

	r.client = generatedClient
}

// Schema defines the resource schema
func (r *InternalnetworksResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages internal networks in Cisco Umbrella. Internal networks represent IP address ranges within sites and are foundational for network topology management.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for this internal network (same as origin_id)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"origin_id": schema.Int64Attribute{
				Computed:    true,
				Description: "The origin ID of the internal network",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the internal network (1-50 characters)",
			},
			"ip_address": schema.StringAttribute{
				Required:    true,
				Description: "The IPv4 address of the internal network (7-15 characters)",
			},
			"prefix_length": schema.Int64Attribute{
				Required:    true,
				Description: "The length of the prefix. Must be between 8 and 32",
			},
			"site_id": schema.Int64Attribute{
				Optional:    true,
				Description: "The site ID. For DNS policies, specify the ID of the site that is associated with internal network. Provide either site_id, network_id, or tunnel_id",
			},
			"site_name": schema.StringAttribute{
				Computed:    true,
				Description: "The name of the site associated with the internal network",
			},
			"network_id": schema.Int64Attribute{
				Optional:    true,
				Description: "The network ID. For Web policies that use proxy chaining, specify the ID of the network associated with the internal network. Provide either site_id, network_id, or tunnel_id",
			},
			"network_name": schema.StringAttribute{
				Computed:    true,
				Description: "The name of the network associated with the internal network",
			},
			"tunnel_id": schema.Int64Attribute{
				Optional:    true,
				Description: "The ID of the tunnel. For Web policies that use an IPsec tunnel, specify the ID of tunnel associated with the internal network. Provide either site_id, network_id, or tunnel_id",
			},
			"tunnel_name": schema.StringAttribute{
				Computed:    true,
				Description: "The name of the tunnel associated with the internal network",
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: "The date and time when the internal network was created",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"modified_at": schema.StringAttribute{
				Computed:    true,
				Description: "The date and time when the internal network was last modified",
			},
		},
	}
}

// Create creates a new internal network
func (r *InternalnetworksResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan internalnetworksModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating internal network", map[string]interface{}{
		"name":          plan.Name.ValueString(),
		"ip_address":    plan.IPAddress.ValueString(),
		"prefix_length": plan.PrefixLength.ValueInt64(),
	})

	// Validate required fields and constraints
	if err := r.validateInternalNetworkData(plan); err != nil {
		resp.Diagnostics.AddError("Validation Error", err.Error())
		return
	}

	// Prepare request payload
	payload := map[string]interface{}{
		"name":         plan.Name.ValueString(),
		"ipAddress":    plan.IPAddress.ValueString(),
		"prefixLength": plan.PrefixLength.ValueInt64(),
	}

	// Add optional association fields (only one should be provided)
	if !plan.SiteID.IsNull() && !plan.SiteID.IsUnknown() {
		payload["siteId"] = plan.SiteID.ValueInt64()
	}
	if !plan.NetworkID.IsNull() && !plan.NetworkID.IsUnknown() {
		payload["networkId"] = plan.NetworkID.ValueInt64()
	}
	if !plan.TunnelID.IsNull() && !plan.TunnelID.IsUnknown() {
		payload["tunnelId"] = plan.TunnelID.ValueInt64()
	}

	// Make API call to create internal network
	result, err := r.client.CreateInternalNetwork(ctx, payload)
	if err != nil {
		resp.Diagnostics.AddError("Failed to create internal network", err.Error())
		return
	}

	// Update plan with response data
	r.updateModelFromAPIResponse(&plan, result)

	tflog.Debug(ctx, "Internal network created successfully", map[string]interface{}{
		"id":        plan.ID.ValueString(),
		"origin_id": plan.OriginID.ValueInt64(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Read reads the internal network resource
func (r *InternalnetworksResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state internalnetworksModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading internal network", map[string]interface{}{
		"id": state.ID.ValueString(),
	})

	// Make API call to get internal network
	result, err := r.client.GetInternalNetwork(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to read internal network", err.Error())
		return
	}

	// Update state with response data
	r.updateModelFromAPIResponse(&state, result)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update updates the internal network
func (r *InternalnetworksResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan internalnetworksModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating internal network", map[string]interface{}{
		"id":   plan.ID.ValueString(),
		"name": plan.Name.ValueString(),
	})

	// Validate required fields and constraints
	if err := r.validateInternalNetworkData(plan); err != nil {
		resp.Diagnostics.AddError("Validation Error", err.Error())
		return
	}

	// Prepare request payload
	payload := map[string]interface{}{
		"name":         plan.Name.ValueString(),
		"ipAddress":    plan.IPAddress.ValueString(),
		"prefixLength": plan.PrefixLength.ValueInt64(),
	}

	// Add optional association fields (only one should be provided)
	if !plan.SiteID.IsNull() && !plan.SiteID.IsUnknown() {
		payload["siteId"] = plan.SiteID.ValueInt64()
	}
	if !plan.NetworkID.IsNull() && !plan.NetworkID.IsUnknown() {
		payload["networkId"] = plan.NetworkID.ValueInt64()
	}
	if !plan.TunnelID.IsNull() && !plan.TunnelID.IsUnknown() {
		payload["tunnelId"] = plan.TunnelID.ValueInt64()
	}

	// Make API call to update internal network
	result, err := r.client.UpdateInternalNetwork(ctx, plan.ID.ValueString(), payload)
	if err != nil {
		resp.Diagnostics.AddError("Failed to update internal network", err.Error())
		return
	}

	// Update plan with response data
	r.updateModelFromAPIResponse(&plan, result)

	tflog.Debug(ctx, "Internal network updated successfully", map[string]interface{}{
		"id": plan.ID.ValueString(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Delete deletes the internal network
func (r *InternalnetworksResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state internalnetworksModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting internal network", map[string]interface{}{
		"id": state.ID.ValueString(),
	})

	// Make API call to delete internal network
	err := r.client.DeleteInternalNetwork(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to delete internal network", err.Error())
		return
	}

	tflog.Debug(ctx, "Internal network deleted successfully", map[string]interface{}{
		"id": state.ID.ValueString(),
	})

	// Clear cache for this resource
	r.client.clearCacheForPath(fmt.Sprintf("/deployments/v2/internalnetworks/%s", state.ID.ValueString()))
}

// validateInternalNetworkData validates the internal network configuration
func (r *InternalnetworksResource) validateInternalNetworkData(model internalnetworksModel) error {
	// Validate name length
	name := model.Name.ValueString()
	if len(name) < 1 || len(name) > 50 {
		return fmt.Errorf("internal network name must be between 1 and 50 characters")
	}

	// Validate IP address length
	ipAddress := model.IPAddress.ValueString()
	if len(ipAddress) < 7 || len(ipAddress) > 15 {
		return fmt.Errorf("ip_address must be between 7 and 15 characters")
	}

	// Validate prefix length
	prefixLength := model.PrefixLength.ValueInt64()
	if prefixLength < 8 || prefixLength > 32 {
		return fmt.Errorf("prefix length must be between 8 and 32")
	}

	// Validate that at least one association is provided (site_id, network_id, or tunnel_id)
	hasSiteID := !model.SiteID.IsNull() && !model.SiteID.IsUnknown()
	hasNetworkID := !model.NetworkID.IsNull() && !model.NetworkID.IsUnknown()
	hasTunnelID := !model.TunnelID.IsNull() && !model.TunnelID.IsUnknown()

	associationCount := 0
	if hasSiteID {
		associationCount++
	}
	if hasNetworkID {
		associationCount++
	}
	if hasTunnelID {
		associationCount++
	}

	if associationCount == 0 {
		return fmt.Errorf("at least one of site_id, network_id, or tunnel_id must be provided")
	}

	if associationCount > 1 {
		return fmt.Errorf("only one of site_id, network_id, or tunnel_id should be provided")
	}

	return nil
}

// updateModelFromAPIResponse updates the model with data from API response
func (r *InternalnetworksResource) updateModelFromAPIResponse(model *internalnetworksModel, result map[string]interface{}) {
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

	if siteID, ok := result["siteId"].(float64); ok {
		model.SiteID = types.Int64Value(int64(siteID))
	}

	if siteName, ok := result["siteName"].(string); ok {
		model.SiteName = types.StringValue(siteName)
	}

	if networkID, ok := result["networkId"].(float64); ok {
		model.NetworkID = types.Int64Value(int64(networkID))
	}

	if networkName, ok := result["networkName"].(string); ok {
		model.NetworkName = types.StringValue(networkName)
	}

	if tunnelID, ok := result["tunnelId"].(float64); ok {
		model.TunnelID = types.Int64Value(int64(tunnelID))
	}

	if tunnelName, ok := result["tunnelName"].(string); ok {
		model.TunnelName = types.StringValue(tunnelName)
	}

	if createdAt, ok := result["createdAt"].(string); ok {
		model.CreatedAt = types.StringValue(createdAt)
	}

	if modifiedAt, ok := result["modifiedAt"].(string); ok {
		model.ModifiedAt = types.StringValue(modifiedAt)
	}
}
