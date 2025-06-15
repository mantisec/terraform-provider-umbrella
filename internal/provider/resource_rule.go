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

// -----------------------------------------------------------------------------
// Resource: umbrella_rule
// -----------------------------------------------------------------------------

type ruleResource struct{ client *apiClient }

type ruleModel struct {
	ID               types.String `tfsdk:"id"`
	RulesetID        types.String `tfsdk:"ruleset_id"`
	Name             types.String `tfsdk:"name"`
	Action           types.String `tfsdk:"action"`
	Rank             types.Int64  `tfsdk:"rank"`
	DestinationLists types.Set    `tfsdk:"destination_lists"`
	Applications     types.Set    `tfsdk:"applications"`
	Enabled          types.Bool   `tfsdk:"enabled"`
	CreatedAt        types.String `tfsdk:"created_at"`
	UpdatedAt        types.String `tfsdk:"updated_at"`
}

func NewRuleResource() resource.Resource { return &ruleResource{} }

func (r *ruleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "umbrella_rule"
}

func (r *ruleResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		resp.Diagnostics.AddError("Missing provider data", "internal: no client")
		return
	}
	r.client = req.ProviderData.(*apiClient)
}

func (r *ruleResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Umbrella SWG Rule within a Ruleset",
		Attributes: map[string]schema.Attribute{
			"id":                schema.StringAttribute{Computed: true, Description: "Rule ID"},
			"ruleset_id":        schema.StringAttribute{Required: true, Description: "ID of the ruleset this rule belongs to"},
			"name":              schema.StringAttribute{Required: true, Description: "Rule name"},
			"action":            schema.StringAttribute{Required: true, Description: "Rule action (ALLOW, BLOCK, DO_NOT_DECRYPT, etc.)"},
			"rank":              schema.Int64Attribute{Required: true, Description: "Rule priority/rank (lower numbers have higher priority)"},
			"destination_lists": schema.SetAttribute{Optional: true, ElementType: types.StringType, Description: "List of destination list names to apply this rule to"},
			"applications":      schema.SetAttribute{Optional: true, ElementType: types.StringType, Description: "List of applications to apply this rule to"},
			"enabled":           schema.BoolAttribute{Optional: true, Description: "Whether the rule is enabled"},
			"created_at":        schema.StringAttribute{Computed: true, Description: "Creation timestamp"},
			"updated_at":        schema.StringAttribute{Computed: true, Description: "Last update timestamp"},
		},
	}
}

func (r *ruleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ruleModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	payload := map[string]interface{}{
		"name":   plan.Name.ValueString(),
		"action": plan.Action.ValueString(),
		"rank":   plan.Rank.ValueInt64(),
	}

	if !plan.DestinationLists.IsNull() {
		destLists := setToStringSlice(ctx, plan.DestinationLists, &resp.Diagnostics)
		payload["destinationLists"] = destLists
	} else {
		payload["destinationLists"] = []string{}
	}

	if !plan.Applications.IsNull() {
		apps := setToStringSlice(ctx, plan.Applications, &resp.Diagnostics)
		payload["applications"] = apps
	} else {
		payload["applications"] = []string{}
	}

	if !plan.Enabled.IsNull() {
		payload["enabled"] = plan.Enabled.ValueBool()
	}

	body, _ := json.Marshal(payload)

	apiResp, err := r.client.do(ctx, http.MethodPost, fmt.Sprintf(rulePath, r.client.orgID, plan.RulesetID.ValueString()), body)
	if err != nil {
		resp.Diagnostics.AddError("API error", err.Error())
		return
	}
	defer apiResp.Body.Close()

	if apiResp.StatusCode != http.StatusCreated && apiResp.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("Create failed", fmt.Sprintf("HTTP %s", apiResp.Status))
		return
	}

	var data struct {
		ID               string   `json:"id"`
		Name             string   `json:"name"`
		Action           string   `json:"action"`
		Rank             int64    `json:"rank"`
		DestinationLists []string `json:"destinationLists"`
		Applications     []string `json:"applications"`
		Enabled          bool     `json:"enabled"`
		CreatedAt        string   `json:"createdAt"`
		UpdatedAt        string   `json:"updatedAt"`
	}
	if err := json.NewDecoder(apiResp.Body).Decode(&data); err != nil {
		resp.Diagnostics.AddError("decode", err.Error())
		return
	}

	plan.ID = types.StringValue(data.ID)
	plan.Enabled = types.BoolValue(data.Enabled)
	plan.CreatedAt = types.StringValue(data.CreatedAt)
	plan.UpdatedAt = types.StringValue(data.UpdatedAt)

	// Set destination lists
	destElems := stringSliceToAttrValues(data.DestinationLists)
	plan.DestinationLists, _ = types.SetValue(types.StringType, destElems)

	// Set applications
	appElems := stringSliceToAttrValues(data.Applications)
	plan.Applications, _ = types.SetValue(types.StringType, appElems)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *ruleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ruleModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiResp, err := r.client.do(ctx, http.MethodGet, fmt.Sprintf(rulePath+"/%s", r.client.orgID, state.RulesetID.ValueString(), state.ID.ValueString()), nil)
	if err != nil {
		resp.Diagnostics.AddError("API error", err.Error())
		return
	}
	defer apiResp.Body.Close()

	if apiResp.StatusCode == http.StatusNotFound {
		resp.State.RemoveResource(ctx)
		return
	}

	if apiResp.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("Read failed", fmt.Sprintf("HTTP %s", apiResp.Status))
		return
	}

	var rule struct {
		ID               string   `json:"id"`
		Name             string   `json:"name"`
		Action           string   `json:"action"`
		Rank             int64    `json:"rank"`
		DestinationLists []string `json:"destinationLists"`
		Applications     []string `json:"applications"`
		Enabled          bool     `json:"enabled"`
		CreatedAt        string   `json:"createdAt"`
		UpdatedAt        string   `json:"updatedAt"`
	}
	if err := json.NewDecoder(apiResp.Body).Decode(&rule); err != nil {
		resp.Diagnostics.AddError("decode", err.Error())
		return
	}

	state.Name = types.StringValue(rule.Name)
	state.Action = types.StringValue(rule.Action)
	state.Rank = types.Int64Value(rule.Rank)
	state.Enabled = types.BoolValue(rule.Enabled)
	state.CreatedAt = types.StringValue(rule.CreatedAt)
	state.UpdatedAt = types.StringValue(rule.UpdatedAt)

	// Set destination lists
	destElems := stringSliceToAttrValues(rule.DestinationLists)
	state.DestinationLists, _ = types.SetValue(types.StringType, destElems)

	// Set applications
	appElems := stringSliceToAttrValues(rule.Applications)
	state.Applications, _ = types.SetValue(types.StringType, appElems)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ruleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ruleModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	payload := map[string]interface{}{}
	needsUpdate := false

	if plan.Name != state.Name {
		payload["name"] = plan.Name.ValueString()
		needsUpdate = true
	}
	if plan.Action != state.Action {
		payload["action"] = plan.Action.ValueString()
		needsUpdate = true
	}
	if plan.Rank != state.Rank {
		payload["rank"] = plan.Rank.ValueInt64()
		needsUpdate = true
	}
	if plan.Enabled != state.Enabled {
		payload["enabled"] = plan.Enabled.ValueBool()
		needsUpdate = true
	}

	// Check destination lists
	planDestLists := setToStringSlice(ctx, plan.DestinationLists, &resp.Diagnostics)
	stateDestLists := setToStringSlice(ctx, state.DestinationLists, &resp.Diagnostics)
	if !stringSlicesEqual(planDestLists, stateDestLists) {
		payload["destinationLists"] = planDestLists
		needsUpdate = true
	}

	// Check applications
	planApps := setToStringSlice(ctx, plan.Applications, &resp.Diagnostics)
	stateApps := setToStringSlice(ctx, state.Applications, &resp.Diagnostics)
	if !stringSlicesEqual(planApps, stateApps) {
		payload["applications"] = planApps
		needsUpdate = true
	}

	if needsUpdate {
		body, _ := json.Marshal(payload)

		apiResp, err := r.client.do(ctx, http.MethodPut, fmt.Sprintf(rulePath+"/%s", r.client.orgID, state.RulesetID.ValueString(), state.ID.ValueString()), body)
		if err != nil {
			resp.Diagnostics.AddError("update rule", err.Error())
			return
		}
		defer apiResp.Body.Close()

		if apiResp.StatusCode != http.StatusOK {
			resp.Diagnostics.AddError("Update failed", fmt.Sprintf("HTTP %s", apiResp.Status))
			return
		}

		var data struct {
			UpdatedAt string `json:"updatedAt"`
		}
		if err := json.NewDecoder(apiResp.Body).Decode(&data); err != nil {
			resp.Diagnostics.AddError("decode", err.Error())
			return
		}

		plan.UpdatedAt = types.StringValue(data.UpdatedAt)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *ruleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ruleModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiResp, err := r.client.do(ctx, http.MethodDelete, fmt.Sprintf(rulePath+"/%s", r.client.orgID, state.RulesetID.ValueString(), state.ID.ValueString()), nil)
	if err != nil {
		resp.Diagnostics.AddError("delete", err.Error())
		return
	}
	defer apiResp.Body.Close()

	if apiResp.StatusCode != http.StatusNoContent && apiResp.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("Delete failed", fmt.Sprintf("HTTP %s", apiResp.Status))
	}
}
