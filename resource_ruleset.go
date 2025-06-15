package main

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
// Resource: umbrella_ruleset
// -----------------------------------------------------------------------------

type rulesetResource struct{ client *apiClient }

type rulesetModel struct {
	ID                   types.String `tfsdk:"id"`
	Name                 types.String `tfsdk:"name"`
	Description          types.String `tfsdk:"description"`
	SAMLEnabled          types.Bool   `tfsdk:"saml_enabled"`
	SSLDecryptionEnabled types.Bool   `tfsdk:"ssl_decryption_enabled"`
	CreatedAt            types.String `tfsdk:"created_at"`
	UpdatedAt            types.String `tfsdk:"updated_at"`
}

func NewRulesetResource() resource.Resource { return &rulesetResource{} }

func (r *rulesetResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "umbrella_ruleset"
}

func (r *rulesetResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		resp.Diagnostics.AddError("Missing provider data", "internal: no client")
		return
	}
	r.client = req.ProviderData.(*apiClient)
}

func (r *rulesetResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Umbrella SWG Ruleset Configuration",
		Attributes: map[string]schema.Attribute{
			"id":                     schema.StringAttribute{Computed: true, Description: "Ruleset ID"},
			"name":                   schema.StringAttribute{Required: true, Description: "Ruleset name"},
			"description":            schema.StringAttribute{Optional: true, Description: "Ruleset description"},
			"saml_enabled":           schema.BoolAttribute{Optional: true, Description: "Enable SAML authentication for this ruleset"},
			"ssl_decryption_enabled": schema.BoolAttribute{Optional: true, Description: "Enable SSL decryption for this ruleset"},
			"created_at":             schema.StringAttribute{Computed: true, Description: "Creation timestamp"},
			"updated_at":             schema.StringAttribute{Computed: true, Description: "Last update timestamp"},
		},
	}
}

func (r *rulesetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan rulesetModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	payload := map[string]interface{}{
		"name": plan.Name.ValueString(),
	}

	if !plan.Description.IsNull() {
		payload["description"] = plan.Description.ValueString()
	}
	if !plan.SAMLEnabled.IsNull() {
		payload["samlEnabled"] = plan.SAMLEnabled.ValueBool()
	}
	if !plan.SSLDecryptionEnabled.IsNull() {
		payload["sslDecryptionEnabled"] = plan.SSLDecryptionEnabled.ValueBool()
	}

	body, _ := json.Marshal(payload)

	apiResp, err := r.client.do(ctx, http.MethodPost, fmt.Sprintf(rulesetPath, r.client.orgID), body)
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
		ID                   string `json:"id"`
		Name                 string `json:"name"`
		Description          string `json:"description"`
		SAMLEnabled          bool   `json:"samlEnabled"`
		SSLDecryptionEnabled bool   `json:"sslDecryptionEnabled"`
		CreatedAt            string `json:"createdAt"`
		UpdatedAt            string `json:"updatedAt"`
	}
	if err := json.NewDecoder(apiResp.Body).Decode(&data); err != nil {
		resp.Diagnostics.AddError("decode", err.Error())
		return
	}

	plan.ID = types.StringValue(data.ID)
	plan.Description = types.StringValue(data.Description)
	plan.SAMLEnabled = types.BoolValue(data.SAMLEnabled)
	plan.SSLDecryptionEnabled = types.BoolValue(data.SSLDecryptionEnabled)
	plan.CreatedAt = types.StringValue(data.CreatedAt)
	plan.UpdatedAt = types.StringValue(data.UpdatedAt)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *rulesetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state rulesetModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiResp, err := r.client.do(ctx, http.MethodGet, fmt.Sprintf(rulesetPath+"/%s", r.client.orgID, state.ID.ValueString()), nil)
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

	var ruleset struct {
		ID                   string `json:"id"`
		Name                 string `json:"name"`
		Description          string `json:"description"`
		SAMLEnabled          bool   `json:"samlEnabled"`
		SSLDecryptionEnabled bool   `json:"sslDecryptionEnabled"`
		CreatedAt            string `json:"createdAt"`
		UpdatedAt            string `json:"updatedAt"`
	}
	if err := json.NewDecoder(apiResp.Body).Decode(&ruleset); err != nil {
		resp.Diagnostics.AddError("decode", err.Error())
		return
	}

	state.Name = types.StringValue(ruleset.Name)
	state.Description = types.StringValue(ruleset.Description)
	state.SAMLEnabled = types.BoolValue(ruleset.SAMLEnabled)
	state.SSLDecryptionEnabled = types.BoolValue(ruleset.SSLDecryptionEnabled)
	state.CreatedAt = types.StringValue(ruleset.CreatedAt)
	state.UpdatedAt = types.StringValue(ruleset.UpdatedAt)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *rulesetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state rulesetModel
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
	if plan.Description != state.Description {
		payload["description"] = plan.Description.ValueString()
		needsUpdate = true
	}
	if plan.SAMLEnabled != state.SAMLEnabled {
		payload["samlEnabled"] = plan.SAMLEnabled.ValueBool()
		needsUpdate = true
	}
	if plan.SSLDecryptionEnabled != state.SSLDecryptionEnabled {
		payload["sslDecryptionEnabled"] = plan.SSLDecryptionEnabled.ValueBool()
		needsUpdate = true
	}

	if needsUpdate {
		body, _ := json.Marshal(payload)

		apiResp, err := r.client.do(ctx, http.MethodPatch, fmt.Sprintf(rulesetPath+"/%s", r.client.orgID, state.ID.ValueString()), body)
		if err != nil {
			resp.Diagnostics.AddError("update ruleset", err.Error())
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

func (r *rulesetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state rulesetModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiResp, err := r.client.do(ctx, http.MethodDelete, fmt.Sprintf(rulesetPath+"/%s", r.client.orgID, state.ID.ValueString()), nil)
	if err != nil {
		resp.Diagnostics.AddError("delete", err.Error())
		return
	}
	defer apiResp.Body.Close()

	if apiResp.StatusCode != http.StatusNoContent && apiResp.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("Delete failed", fmt.Sprintf("HTTP %s", apiResp.Status))
	}
}
