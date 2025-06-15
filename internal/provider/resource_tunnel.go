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
// Resource: umbrella_tunnel
// -----------------------------------------------------------------------------

type tunnelResource struct{ client *apiClient }

type tunnelModel struct {
	ID           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	DeviceIP     types.String `tfsdk:"device_ip"`
	PreSharedKey types.String `tfsdk:"pre_shared_key"`
	Status       types.String `tfsdk:"status"`
	CreatedAt    types.String `tfsdk:"created_at"`
	UpdatedAt    types.String `tfsdk:"updated_at"`
}

func NewTunnelResource() resource.Resource { return &tunnelResource{} }

func (r *tunnelResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "umbrella_tunnel"
}

func (r *tunnelResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		resp.Diagnostics.AddError("Missing provider data", "internal: no client")
		return
	}
	r.client = req.ProviderData.(*apiClient)
}

func (r *tunnelResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Umbrella IPSec Tunnel for Secure Internet Gateway",
		Attributes: map[string]schema.Attribute{
			"id":             schema.StringAttribute{Computed: true, Description: "Tunnel ID"},
			"name":           schema.StringAttribute{Required: true, Description: "Tunnel name"},
			"device_ip":      schema.StringAttribute{Required: true, Description: "Device IP address for the tunnel"},
			"pre_shared_key": schema.StringAttribute{Required: true, Sensitive: true, Description: "Pre-shared key for IPSec tunnel"},
			"status":         schema.StringAttribute{Computed: true, Description: "Tunnel status"},
			"created_at":     schema.StringAttribute{Computed: true, Description: "Creation timestamp"},
			"updated_at":     schema.StringAttribute{Computed: true, Description: "Last update timestamp"},
		},
	}
}

// ------------------ CRUD ------------------

func (r *tunnelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan tunnelModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	payload := map[string]string{
		"name":         plan.Name.ValueString(),
		"deviceIp":     plan.DeviceIP.ValueString(),
		"preSharedKey": plan.PreSharedKey.ValueString(),
	}
	body, _ := json.Marshal(payload)

	apiResp, err := r.client.do(ctx, http.MethodPost, fmt.Sprintf(tunnelPath, r.client.orgID), body)
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
		ID        string `json:"id"`
		Name      string `json:"name"`
		DeviceIP  string `json:"deviceIp"`
		Status    string `json:"status"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	}
	if err := json.NewDecoder(apiResp.Body).Decode(&data); err != nil {
		resp.Diagnostics.AddError("decode", err.Error())
		return
	}

	plan.ID = types.StringValue(data.ID)
	plan.Status = types.StringValue(data.Status)
	plan.CreatedAt = types.StringValue(data.CreatedAt)
	plan.UpdatedAt = types.StringValue(data.UpdatedAt)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *tunnelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state tunnelModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiResp, err := r.client.do(ctx, http.MethodGet, fmt.Sprintf(tunnelPath+"/%s", r.client.orgID, state.ID.ValueString()), nil)
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

	var tunnel struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		DeviceIP  string `json:"deviceIp"`
		Status    string `json:"status"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	}
	if err := json.NewDecoder(apiResp.Body).Decode(&tunnel); err != nil {
		resp.Diagnostics.AddError("decode", err.Error())
		return
	}

	state.Name = types.StringValue(tunnel.Name)
	state.DeviceIP = types.StringValue(tunnel.DeviceIP)
	state.Status = types.StringValue(tunnel.Status)
	state.CreatedAt = types.StringValue(tunnel.CreatedAt)
	state.UpdatedAt = types.StringValue(tunnel.UpdatedAt)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *tunnelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state tunnelModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check if any updateable fields have changed
	if plan.Name != state.Name || plan.DeviceIP != state.DeviceIP || plan.PreSharedKey != state.PreSharedKey {
		payload := map[string]string{
			"name":         plan.Name.ValueString(),
			"deviceIp":     plan.DeviceIP.ValueString(),
			"preSharedKey": plan.PreSharedKey.ValueString(),
		}
		body, _ := json.Marshal(payload)

		apiResp, err := r.client.do(ctx, http.MethodPut, fmt.Sprintf(tunnelPath+"/%s", r.client.orgID, state.ID.ValueString()), body)
		if err != nil {
			resp.Diagnostics.AddError("update tunnel", err.Error())
			return
		}
		defer apiResp.Body.Close()

		if apiResp.StatusCode != http.StatusOK {
			resp.Diagnostics.AddError("Update failed", fmt.Sprintf("HTTP %s", apiResp.Status))
			return
		}

		var data struct {
			Status    string `json:"status"`
			UpdatedAt string `json:"updatedAt"`
		}
		if err := json.NewDecoder(apiResp.Body).Decode(&data); err != nil {
			resp.Diagnostics.AddError("decode", err.Error())
			return
		}

		plan.Status = types.StringValue(data.Status)
		plan.UpdatedAt = types.StringValue(data.UpdatedAt)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *tunnelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state tunnelModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiResp, err := r.client.do(ctx, http.MethodDelete, fmt.Sprintf(tunnelPath+"/%s", r.client.orgID, state.ID.ValueString()), nil)
	if err != nil {
		resp.Diagnostics.AddError("delete", err.Error())
		return
	}
	defer apiResp.Body.Close()

	if apiResp.StatusCode != http.StatusNoContent && apiResp.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("Delete failed", fmt.Sprintf("HTTP %s", apiResp.Status))
	}
}
