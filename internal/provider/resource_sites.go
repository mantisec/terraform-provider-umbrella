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

// SitesResource implements the sites resource
type SitesResource struct {
	client *GeneratedClient
}

// sitesModel represents the resource data model
type sitesModel struct {
	ID                   types.String `tfsdk:"id"`
	SiteID               types.Int64  `tfsdk:"site_id"`
	OriginID             types.Int64  `tfsdk:"origin_id"`
	Name                 types.String `tfsdk:"name"`
	Type                 types.String `tfsdk:"type"`
	IsDefault            types.Bool   `tfsdk:"is_default"`
	InternalNetworkCount types.Int64  `tfsdk:"internal_network_count"`
	VACount              types.Int64  `tfsdk:"va_count"`
	CreatedAt            types.String `tfsdk:"created_at"`
	ModifiedAt           types.String `tfsdk:"modified_at"`
}

// NewSitesResource creates a new sites resource
func NewSitesResource() resource.Resource {
	return &SitesResource{}
}

// Metadata returns the resource type name
func (r *SitesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "umbrella_sites"
}

// Configure configures the resource with the provider client
func (r *SitesResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *SitesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages sites in Cisco Umbrella. Sites represent organizational locations and are foundational for other Umbrella resources.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for this site (same as site_id)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"site_id": schema.Int64Attribute{
				Computed:    true,
				Description: "The ID of the site",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"origin_id": schema.Int64Attribute{
				Computed:    true,
				Description: "The origin ID of the site",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the site (1-255 characters)",
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: "The type of the site (SITE or MOBILE_DEVICE)",
			},
			"is_default": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
				Description: "Specifies whether the site is the default",
			},
			"internal_network_count": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of internal networks that are attached to the site",
			},
			"va_count": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of virtual appliances that are attached to the site",
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: "The date and time when the site was created",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"modified_at": schema.StringAttribute{
				Computed:    true,
				Description: "The date and time when the site was last modified",
			},
		},
	}
}

// Create creates a new site
func (r *SitesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan sitesModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating site", map[string]interface{}{
		"name": plan.Name.ValueString(),
	})

	// Validate required fields and constraints
	if err := r.validateSiteData(plan); err != nil {
		resp.Diagnostics.AddError("Validation Error", err.Error())
		return
	}

	// Prepare request payload
	payload := map[string]interface{}{
		"name": plan.Name.ValueString(),
	}

	// Make API call to create site
	result, err := r.client.CreateSite(ctx, payload)
	if err != nil {
		resp.Diagnostics.AddError("Failed to create site", err.Error())
		return
	}

	// Update plan with response data
	r.updateModelFromAPIResponse(&plan, result)

	tflog.Debug(ctx, "Site created successfully", map[string]interface{}{
		"id":      plan.ID.ValueString(),
		"site_id": plan.SiteID.ValueInt64(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Read reads the site resource
func (r *SitesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state sitesModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading site", map[string]interface{}{
		"id": state.ID.ValueString(),
	})

	// Make API call to get site
	result, err := r.client.GetSite(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to read site", err.Error())
		return
	}

	// Update state with response data
	r.updateModelFromAPIResponse(&state, result)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update updates the site
func (r *SitesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan sitesModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating site", map[string]interface{}{
		"id":   plan.ID.ValueString(),
		"name": plan.Name.ValueString(),
	})

	// Validate required fields and constraints
	if err := r.validateSiteData(plan); err != nil {
		resp.Diagnostics.AddError("Validation Error", err.Error())
		return
	}

	// Prepare request payload
	payload := map[string]interface{}{
		"name": plan.Name.ValueString(),
	}

	// Make API call to update site
	result, err := r.client.UpdateSite(ctx, plan.ID.ValueString(), payload)
	if err != nil {
		resp.Diagnostics.AddError("Failed to update site", err.Error())
		return
	}

	// Update plan with response data
	r.updateModelFromAPIResponse(&plan, result)

	tflog.Debug(ctx, "Site updated successfully", map[string]interface{}{
		"id": plan.ID.ValueString(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Delete deletes the site
func (r *SitesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state sitesModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting site", map[string]interface{}{
		"id": state.ID.ValueString(),
	})

	// Make API call to delete site
	err := r.client.DeleteSite(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to delete site", err.Error())
		return
	}

	tflog.Debug(ctx, "Site deleted successfully", map[string]interface{}{
		"id": state.ID.ValueString(),
	})

	// Clear cache for this resource
	r.client.clearCacheForPath(fmt.Sprintf("/deployments/v2/sites/%s", state.ID.ValueString()))
}

// validateSiteData validates the site configuration
func (r *SitesResource) validateSiteData(model sitesModel) error {
	// Validate name length
	name := model.Name.ValueString()
	if len(name) < 1 || len(name) > 255 {
		return fmt.Errorf("site name must be between 1 and 255 characters")
	}

	return nil
}

// updateModelFromAPIResponse updates the model with data from API response
func (r *SitesResource) updateModelFromAPIResponse(model *sitesModel, result map[string]interface{}) {
	if siteID, ok := result["siteId"].(float64); ok {
		model.SiteID = types.Int64Value(int64(siteID))
		model.ID = types.StringValue(strconv.FormatInt(int64(siteID), 10))
	}

	if originID, ok := result["originId"].(float64); ok {
		model.OriginID = types.Int64Value(int64(originID))
	}

	if name, ok := result["name"].(string); ok {
		model.Name = types.StringValue(name)
	}

	if siteType, ok := result["type"].(string); ok {
		model.Type = types.StringValue(siteType)
	}

	if isDefault, ok := result["isDefault"].(bool); ok {
		model.IsDefault = types.BoolValue(isDefault)
	}

	if internalNetworkCount, ok := result["internalNetworkCount"].(float64); ok {
		model.InternalNetworkCount = types.Int64Value(int64(internalNetworkCount))
	}

	if vaCount, ok := result["vaCount"].(float64); ok {
		model.VACount = types.Int64Value(int64(vaCount))
	}

	if createdAt, ok := result["createdAt"].(string); ok {
		model.CreatedAt = types.StringValue(createdAt)
	}

	if modifiedAt, ok := result["modifiedAt"].(string); ok {
		model.ModifiedAt = types.StringValue(modifiedAt)
	}
}
