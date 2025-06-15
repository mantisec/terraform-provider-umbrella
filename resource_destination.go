package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// -----------------------------------------------------------------------------
// Resource: umbrella_destination
// -----------------------------------------------------------------------------

type destinationResource struct {
	client *apiClient
}

type destinationModel struct {
	ID                types.String `tfsdk:"id"`
	DestinationListID types.String `tfsdk:"destination_list_id"`
	Destination       types.String `tfsdk:"destination"`
	Comment           types.String `tfsdk:"comment"`
}

func NewDestinationResource() resource.Resource {
	return &destinationResource{}
}

func (r *destinationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "umbrella_destination"
}

func (r *destinationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		resp.Diagnostics.AddError("Missing provider data", "internal: no client")
		return
	}
	r.client = req.ProviderData.(*apiClient)
}

func (r *destinationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages individual destinations within an Umbrella destination list",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for this destination",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"destination_list_id": schema.StringAttribute{
				Required:    true,
				Description: "The ID of the destination list this destination belongs to",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"destination": schema.StringAttribute{
				Required:    true,
				Description: "The destination value (URL, domain, or CIDR block)",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Description: "Optional comment for this destination",
			},
		},
	}
}

// ------------------ CRUD ------------------

func (r *destinationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan destinationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create the destination entry
	entry := map[string]string{
		"destination": plan.Destination.ValueString(),
	}
	if !plan.Comment.IsNull() && plan.Comment.ValueString() != "" {
		entry["comment"] = plan.Comment.ValueString()
	}

	entries := []map[string]string{entry}
	body, _ := json.Marshal(entries)

	path := fmt.Sprintf(destListPath+"/%s/destinations", r.client.orgID, plan.DestinationListID.ValueString())
	apiResp, err := r.client.do(ctx, http.MethodPost, path, body)
	if err != nil {
		resp.Diagnostics.AddError("API error", err.Error())
		return
	}
	defer apiResp.Body.Close()

	if apiResp.StatusCode != http.StatusCreated && apiResp.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("Create failed", fmt.Sprintf("HTTP %s", apiResp.Status))
		return
	}

	// Generate a unique ID for this destination within the list
	// Since the API doesn't return individual destination IDs, we'll use a combination
	plan.ID = types.StringValue(fmt.Sprintf("%s:%s", plan.DestinationListID.ValueString(), plan.Destination.ValueString()))

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *destinationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state destinationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get all destinations from the list and find our specific destination
	destinations, err := r.getDestinationsFromList(ctx, state.DestinationListID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to read destinations", err.Error())
		return
	}

	// Check if our destination still exists
	found := false
	for _, dest := range destinations {
		if dest.Destination == state.Destination.ValueString() {
			found = true
			// Update comment if it exists
			if dest.Comment != "" {
				state.Comment = types.StringValue(dest.Comment)
			} else {
				state.Comment = types.StringNull()
			}
			break
		}
	}

	if !found {
		// Destination no longer exists, remove from state
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *destinationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state destinationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// For updates, we need to delete the old destination and add the new one
	// since the Umbrella API doesn't support direct destination updates

	// First, remove the old destination
	if err := r.removeDestination(ctx, state.DestinationListID.ValueString(), state.Destination.ValueString()); err != nil {
		resp.Diagnostics.AddError("Failed to remove old destination", err.Error())
		return
	}

	// Then add the new destination
	entry := map[string]string{
		"destination": plan.Destination.ValueString(),
	}
	if !plan.Comment.IsNull() && plan.Comment.ValueString() != "" {
		entry["comment"] = plan.Comment.ValueString()
	}

	entries := []map[string]string{entry}
	body, _ := json.Marshal(entries)

	path := fmt.Sprintf(destListPath+"/%s/destinations", r.client.orgID, plan.DestinationListID.ValueString())
	apiResp, err := r.client.do(ctx, http.MethodPost, path, body)
	if err != nil {
		resp.Diagnostics.AddError("Failed to add updated destination", err.Error())
		return
	}
	defer apiResp.Body.Close()

	if apiResp.StatusCode != http.StatusCreated && apiResp.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("Update failed", fmt.Sprintf("HTTP %s", apiResp.Status))
		return
	}

	// Update the ID if the destination value changed
	plan.ID = types.StringValue(fmt.Sprintf("%s:%s", plan.DestinationListID.ValueString(), plan.Destination.ValueString()))

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *destinationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state destinationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.removeDestination(ctx, state.DestinationListID.ValueString(), state.Destination.ValueString()); err != nil {
		resp.Diagnostics.AddError("Failed to delete destination", err.Error())
	}
}

// ------------------ Helper Methods ------------------

// destinationEntry represents a single destination with optional comment
type destinationEntry struct {
	Destination string `json:"destination"`
	Comment     string `json:"comment,omitempty"`
}

// getDestinationsFromList retrieves all destinations from a specific destination list
func (r *destinationResource) getDestinationsFromList(ctx context.Context, listID string) ([]destinationEntry, error) {
	path := fmt.Sprintf(destListPath+"/%s/destinations", r.client.orgID, listID)
	resp, err := r.client.do(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("destinations GET %s", resp.Status)
	}

	var destinations []destinationEntry
	if err := json.NewDecoder(resp.Body).Decode(&destinations); err != nil {
		return nil, err
	}

	return destinations, nil
}

// removeDestination removes a specific destination from a destination list
func (r *destinationResource) removeDestination(ctx context.Context, listID, destination string) error {
	entries := []map[string]string{
		{"destination": destination},
	}
	body, _ := json.Marshal(entries)

	path := fmt.Sprintf(destListPath+"/%s/destinations", r.client.orgID, listID)
	resp, err := r.client.do(ctx, http.MethodDelete, path, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("delete destination HTTP %s", resp.Status)
	}

	return nil
}
