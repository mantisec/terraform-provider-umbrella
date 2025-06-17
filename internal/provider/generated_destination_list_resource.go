package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// GeneratedDestinationListResource implements the destination_list resource
type GeneratedDestinationListResource struct {
	client *GeneratedClient
}

// destinationListModel represents the resource data model
type destinationListModel struct {
	ID           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	Access       types.String `tfsdk:"access"`
	IsGlobal     types.Bool   `tfsdk:"is_global"`
	Destinations types.Set    `tfsdk:"destinations"`
	CreatedAt    types.String `tfsdk:"created_at"`
	ModifiedAt   types.String `tfsdk:"modified_at"`
}

// NewGeneratedDestinationListResource creates a new destination_list resource
func NewGeneratedDestinationListResource() resource.Resource {
	return &GeneratedDestinationListResource{}
}

// Metadata returns the resource type name
func (r *GeneratedDestinationListResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "umbrella_destination_list"
}

// Configure configures the resource with the provider client
func (r *GeneratedDestinationListResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// Schema defines the resource schema with advanced features
func (r *GeneratedDestinationListResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages destination lists in Cisco Umbrella",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for this destination list",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the destination list",
			},
			"access": schema.StringAttribute{
				Required:    true,
				Description: "Access type for the destination list (allow/block)",
			},
			"is_global": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether this is a global destination list",
			},
			"destinations": schema.SetAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "List of destinations (domains, IPs, URLs)",
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: "Timestamp when the destination list was created",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"modified_at": schema.StringAttribute{
				Computed:    true,
				Description: "Timestamp when the destination list was last modified",
			},
		},
	}
}

// Create creates a new destination_list
func (r *GeneratedDestinationListResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan destinationListModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Prepare request payload
	payload := map[string]interface{}{
		"name":   plan.Name.ValueString(),
		"access": plan.Access.ValueString(),
	}

	if !plan.IsGlobal.IsNull() {
		payload["isGlobal"] = plan.IsGlobal.ValueBool()
	}

	if !plan.Destinations.IsNull() {
		var destinations []string
		resp.Diagnostics.Append(plan.Destinations.ElementsAs(ctx, &destinations, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		payload["destinations"] = destinations
	}

	// Make API call using generated client method
	result, err := r.client.CreateDestinationList(ctx, payload)
	if err != nil {
		resp.Diagnostics.AddError("Failed to create destination list", err.Error())
		return
	}

	// Update plan with response data
	plan.ID = types.StringValue(result["id"].(string))
	if createdAt, ok := result["createdAt"].(string); ok {
		plan.CreatedAt = types.StringValue(createdAt)
	}
	if modifiedAt, ok := result["modifiedAt"].(string); ok {
		plan.ModifiedAt = types.StringValue(modifiedAt)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Read reads the destination_list
func (r *GeneratedDestinationListResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state destinationListModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Make API call using generated client method with caching
	result, err := r.client.GetDestinationList(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to read destination list", err.Error())
		return
	}

	// Update state with response data
	if name, ok := result["name"].(string); ok {
		state.Name = types.StringValue(name)
	}
	if access, ok := result["access"].(string); ok {
		state.Access = types.StringValue(access)
	}
	if isGlobal, ok := result["isGlobal"].(bool); ok {
		state.IsGlobal = types.BoolValue(isGlobal)
	}
	if modifiedAt, ok := result["modifiedAt"].(string); ok {
		state.ModifiedAt = types.StringValue(modifiedAt)
	}

	// Handle destinations array
	if destinations, ok := result["destinations"].([]interface{}); ok {
		destStrings := make([]string, len(destinations))
		for i, dest := range destinations {
			destStrings[i] = dest.(string)
		}
		destSet, diags := types.SetValueFrom(ctx, types.StringType, destStrings)
		resp.Diagnostics.Append(diags...)
		if !resp.Diagnostics.HasError() {
			state.Destinations = destSet
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update updates the destination_list
func (r *GeneratedDestinationListResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan destinationListModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Prepare request payload
	payload := map[string]interface{}{
		"name":   plan.Name.ValueString(),
		"access": plan.Access.ValueString(),
	}

	if !plan.IsGlobal.IsNull() {
		payload["isGlobal"] = plan.IsGlobal.ValueBool()
	}

	if !plan.Destinations.IsNull() {
		var destinations []string
		resp.Diagnostics.Append(plan.Destinations.ElementsAs(ctx, &destinations, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		payload["destinations"] = destinations
	}

	// Make API call using generated client method
	result, err := r.client.UpdateDestinationList(ctx, plan.ID.ValueString(), payload)
	if err != nil {
		resp.Diagnostics.AddError("Failed to update destination list", err.Error())
		return
	}

	// Update plan with response data
	if modifiedAt, ok := result["modifiedAt"].(string); ok {
		plan.ModifiedAt = types.StringValue(modifiedAt)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Delete deletes the destination_list
func (r *GeneratedDestinationListResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state destinationListModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Make API call using generated client method
	err := r.client.DeleteDestinationList(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to delete destination list", err.Error())
		return
	}

	// Clear cache for this resource
	r.client.clearCacheForPath(fmt.Sprintf("/policies/v2/organizations/%s/destinationlists/%s", r.client.orgID, state.ID.ValueString()))
}
