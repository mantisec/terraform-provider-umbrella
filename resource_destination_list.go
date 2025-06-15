package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// -----------------------------------------------------------------------------
// Resource: umbrella_destination_list
// -----------------------------------------------------------------------------

type destinationListResource struct{ client *apiClient }

type destListModel struct {
	ID           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	Type         types.String `tfsdk:"type"`
	Destinations types.Set    `tfsdk:"destinations"`
}

func NewDestinationListResource() resource.Resource { return &destinationListResource{} }

func (r *destinationListResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "umbrella_destination_list"
}

func (r *destinationListResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		resp.Diagnostics.AddError("Missing provider data", "internal: no client")
		return
	}
	r.client = req.ProviderData.(*apiClient)
}

func (r *destinationListResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Umbrella Destination List (allow, block or SAML-bypass)",
		Attributes: map[string]schema.Attribute{
			"id":           schema.StringAttribute{Computed: true},
			"name":         schema.StringAttribute{Required: true},
			"type":         schema.StringAttribute{Required: true, Description: "URL | CIDR | DOMAIN"},
			"destinations": schema.SetAttribute{Optional: true, ElementType: types.StringType},
		},
	}
}

// ------------------ CRUD ------------------

func (r *destinationListResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan destListModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := map[string]string{"name": plan.Name.ValueString(), "type": plan.Type.ValueString()}
	body, _ := json.Marshal(payload)
	apiResp, err := r.client.do(ctx, http.MethodPost, fmt.Sprintf(destListPath, r.client.orgID), body)
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
		ID int `json:"id"`
	}
	if err := json.NewDecoder(apiResp.Body).Decode(&data); err != nil {
		resp.Diagnostics.AddError("decode", err.Error())
		return
	}
	plan.ID = types.StringValue(fmt.Sprintf("%d", data.ID))

	// Add destinations (if any)
	if !plan.Destinations.IsNull() {
		dests := setToStringSlice(ctx, plan.Destinations, &resp.Diagnostics)
		if len(dests) > 0 {
			if err := r.syncDestinations(ctx, plan.ID.ValueString(), nil, dests); err != nil {
				resp.Diagnostics.AddError("add destinations", err.Error())
				return
			}
		}
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *destinationListResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state destListModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	apiResp, err := r.client.do(ctx, http.MethodGet, fmt.Sprintf(destListPath+"/%s", r.client.orgID, state.ID.ValueString()), nil)
	if err != nil || apiResp.StatusCode == http.StatusNotFound {
		resp.State.RemoveResource(ctx)
		return
	}
	var dl struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}
	if err := json.NewDecoder(apiResp.Body).Decode(&dl); err != nil {
		resp.Diagnostics.AddError("decode", err.Error())
		return
	}
	state.Name = types.StringValue(dl.Name)
	state.Type = types.StringValue(dl.Type)

	// fetch destinations
	dests, err := r.getDestinations(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("destinations", err.Error())
		return
	}
	elems := []attr.Value{}
	for _, d := range dests {
		elems = append(elems, types.StringValue(d))
	}
	state.Destinations, _ = types.SetValue(types.StringType, elems)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *destinationListResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state destListModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Update name/type if changed
	if plan.Name != state.Name || plan.Type != state.Type {
		payload := map[string]string{"name": plan.Name.ValueString(), "type": plan.Type.ValueString()}
		b, _ := json.Marshal(payload)
		if _, err := r.client.do(ctx, http.MethodPut, fmt.Sprintf(destListPath+"/%s", r.client.orgID, state.ID.ValueString()), b); err != nil {
			resp.Diagnostics.AddError("update list", err.Error())
			return
		}
	}

	// ---- destinations diff logic ----
	desired := setToStringSlice(ctx, plan.Destinations, &resp.Diagnostics)
	current := setToStringSlice(ctx, state.Destinations, &resp.Diagnostics)

	toAdd, toDel := diffSlices(current, desired)
	if len(toAdd) > 0 || len(toDel) > 0 {
		if err := r.syncDestinations(ctx, state.ID.ValueString(), toDel, toAdd); err != nil {
			resp.Diagnostics.AddError("sync destinations", err.Error())
			return
		}
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *destinationListResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state destListModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.client.do(ctx, http.MethodDelete, fmt.Sprintf(destListPath+"/%s", r.client.orgID, state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("delete", err.Error())
	}
}

// ------------------ helpers ------------------

func (r *destinationListResource) getDestinations(ctx context.Context, listID string) ([]string, error) {
	path := fmt.Sprintf(destListPath+"/%s/destinations", r.client.orgID, listID)
	resp, err := r.client.do(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("destinations GET %s", resp.Status)
	}
	var out []struct {
		Destination string `json:"destination"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	vals := []string{}
	for _, v := range out {
		vals = append(vals, v.Destination)
	}
	return vals, nil
}

func (r *destinationListResource) syncDestinations(ctx context.Context, listID string, remove []string, add []string) error {
	if len(add) > 0 {
		entries := []map[string]string{}
		for _, d := range add {
			entries = append(entries, map[string]string{"destination": d})
		}
		b, _ := json.Marshal(entries)
		path := fmt.Sprintf(destListPath+"/%s/destinations", r.client.orgID, listID)
		if resp, err := r.client.do(ctx, http.MethodPost, path, b); err != nil || (resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK) {
			if err == nil {
				err = fmt.Errorf("add destinations HTTP %s", resp.Status)
			}
			return err
		}
	}
	if len(remove) > 0 {
		entries := []map[string]string{}
		for _, d := range remove {
			entries = append(entries, map[string]string{"destination": d})
		}
		b, _ := json.Marshal(entries)
		path := fmt.Sprintf(destListPath+"/%s/destinations", r.client.orgID, listID)
		if resp, err := r.client.do(ctx, http.MethodDelete, path, b); err != nil || (resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK) {
			if err == nil {
				err = fmt.Errorf("delete destinations HTTP %s", resp.Status)
			}
			return err
		}
	}
	return nil
}
