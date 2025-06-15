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
// Resource: umbrella_saml
// -----------------------------------------------------------------------------

type samlResource struct{ client *apiClient }

type samlModel struct {
	ID          types.String `tfsdk:"id"`
	MetadataURL types.String `tfsdk:"metadata_url"`
	AuthType    types.String `tfsdk:"auth_type"`
	Enabled     types.Bool   `tfsdk:"enabled"`
}

func NewSAMLResource() resource.Resource { return &samlResource{} }

func (r *samlResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "umbrella_saml"
}

func (r *samlResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		resp.Diagnostics.AddError("Missing provider data", "internal: no client")
		return
	}
	r.client = req.ProviderData.(*apiClient)
}

func (r *samlResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Umbrella SAML Authentication Configuration",
		Attributes: map[string]schema.Attribute{
			"id":           schema.StringAttribute{Computed: true, Description: "SAML configuration ID"},
			"metadata_url": schema.StringAttribute{Required: true, Description: "SAML metadata URL from identity provider"},
			"auth_type":    schema.StringAttribute{Required: true, Description: "Authentication type (e.g., AzureAD, ADFS)"},
			"enabled":      schema.BoolAttribute{Computed: true, Description: "Whether SAML is enabled"},
		},
	}
}

func (r *samlResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan samlModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	payload := map[string]string{
		"metadataUrl": plan.MetadataURL.ValueString(),
		"authType":    plan.AuthType.ValueString(),
	}
	body, _ := json.Marshal(payload)

	apiResp, err := r.client.do(ctx, http.MethodPut, fmt.Sprintf(samlPath, r.client.orgID), body)
	if err != nil {
		resp.Diagnostics.AddError("API error", err.Error())
		return
	}
	defer apiResp.Body.Close()

	if apiResp.StatusCode != http.StatusOK && apiResp.StatusCode != http.StatusCreated {
		resp.Diagnostics.AddError("Create failed", fmt.Sprintf("HTTP %s", apiResp.Status))
		return
	}

	// Use org ID as the ID since SAML config is org-level
	plan.ID = types.StringValue(r.client.orgID)
	plan.Enabled = types.BoolValue(true)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *samlResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state samlModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiResp, err := r.client.do(ctx, http.MethodGet, fmt.Sprintf(samlPath, r.client.orgID), nil)
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

	var samlConfig struct {
		MetadataURL string `json:"metadataUrl"`
		AuthType    string `json:"authType"`
		Enabled     bool   `json:"enabled"`
	}
	if err := json.NewDecoder(apiResp.Body).Decode(&samlConfig); err != nil {
		resp.Diagnostics.AddError("decode", err.Error())
		return
	}

	state.MetadataURL = types.StringValue(samlConfig.MetadataURL)
	state.AuthType = types.StringValue(samlConfig.AuthType)
	state.Enabled = types.BoolValue(samlConfig.Enabled)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *samlResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state samlModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if plan.MetadataURL != state.MetadataURL || plan.AuthType != state.AuthType {
		payload := map[string]string{
			"metadataUrl": plan.MetadataURL.ValueString(),
			"authType":    plan.AuthType.ValueString(),
		}
		body, _ := json.Marshal(payload)

		apiResp, err := r.client.do(ctx, http.MethodPut, fmt.Sprintf(samlPath, r.client.orgID), body)
		if err != nil {
			resp.Diagnostics.AddError("update SAML", err.Error())
			return
		}
		defer apiResp.Body.Close()

		if apiResp.StatusCode != http.StatusOK {
			resp.Diagnostics.AddError("Update failed", fmt.Sprintf("HTTP %s", apiResp.Status))
			return
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *samlResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// SAML configuration cannot be deleted, only disabled
	// This is a no-op as the configuration remains but becomes inactive
}
