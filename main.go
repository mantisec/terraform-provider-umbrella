package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	pschema "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// -----------------------------------------------------------------------------
// Provider entry point
// -----------------------------------------------------------------------------
func main() {
	providerserver.Serve(context.Background(), NewProvider, providerserver.ServeOpts{})
}

// -----------------------------------------------------------------------------
// Provider definition
// -----------------------------------------------------------------------------

const (
	apiBaseURL   = "https://api.umbrella.com"
	apiTokenURL  = apiBaseURL + "/auth/v2/token"
	userAgent    = "terraform-provider-umbrella/0.1.0"
	destListPath = "/policies/v2/organizations/%s/destinationlists"
	tunnelPath   = "/v2/organizations/%s/secureinternetgateway/ipsec/sites"
	samlPath     = "/v2/organizations/%s/saml"
	rulesetPath  = "/policies/v2/organizations/%s/rulesets"
	rulePath     = "/policies/v2/organizations/%s/rulesets/%s/rules"
)

type providerModel struct {
	APIKey    types.String `tfsdk:"api_key"`
	APISecret types.String `tfsdk:"api_secret"`
	OrgID     types.String `tfsdk:"org_id"`
}

type umbrellaProvider struct{ client *apiClient }

func NewProvider() provider.Provider { return &umbrellaProvider{} }

// -----------------------------------------------------------------------------
// Provider metadata & schema
// -----------------------------------------------------------------------------
func (p *umbrellaProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "umbrella"
}

func (p *umbrellaProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = pschema.Schema{
		Description: "Provider for Cisco Umbrella Secure Web Gateway REST API.",
		Attributes: map[string]pschema.Attribute{
			"api_key":    pschema.StringAttribute{Required: true, Sensitive: true, Description: "Umbrella API key (client ID)."},
			"api_secret": pschema.StringAttribute{Required: true, Sensitive: true, Description: "Umbrella API secret (client secret)."},
			"org_id":     pschema.StringAttribute{Required: true, Description: "Umbrella organisation ID."},
		},
	}
}

// -----------------------------------------------------------------------------
// Provider configuration – create API client
// -----------------------------------------------------------------------------
func (p *umbrellaProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var cfg providerModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &cfg)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, err := newAPIClient(ctx, cfg.APIKey.ValueString(), cfg.APISecret.ValueString(), cfg.OrgID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Unable to authenticate", err.Error())
		return
	}
	p.client = client
	resp.ResourceData = client
	resp.DataSourceData = client
}

// -----------------------------------------------------------------------------
// Provider resources & data-sources
// -----------------------------------------------------------------------------
func (p *umbrellaProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewDestinationListResource,
		NewTunnelResource,
		NewSAMLResource,
		NewRulesetResource,
		NewRuleResource,
	}
}
func (p *umbrellaProvider) DataSources(_ context.Context) []func() datasource.DataSource { return nil }

// -----------------------------------------------------------------------------
// Umbrella API client with OAuth2 token caching
// -----------------------------------------------------------------------------

type apiClient struct {
	key, secret, orgID string
	client             *http.Client
	token              string
	expires            time.Time
}

func newAPIClient(ctx context.Context, key, secret, orgID string) (*apiClient, error) {
	c := &apiClient{key: key, secret: secret, orgID: orgID, client: &http.Client{Timeout: 15 * time.Second}}
	if err := c.refreshToken(ctx); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *apiClient) refreshToken(ctx context.Context) error {
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, apiTokenURL, strings.NewReader("grant_type=client_credentials"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", userAgent)
	basic := base64.StdEncoding.EncodeToString([]byte(c.key + ":" + c.secret))
	req.Header.Set("Authorization", "Basic "+basic)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("token request failed: %s", resp.Status)
	}
	var data struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}
	if data.AccessToken == "" {
		return errors.New("no access_token returned")
	}
	c.token = data.AccessToken
	c.expires = time.Now().Add(time.Duration(data.ExpiresIn-60) * time.Second) // refresh 1 min early
	return nil
}

func (c *apiClient) do(ctx context.Context, method, path string, body []byte) (*http.Response, error) {
	if time.Now().After(c.expires) {
		if err := c.refreshToken(ctx); err != nil {
			return nil, err
		}
	}
	url := apiBaseURL + path
	req, _ := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("User-Agent", userAgent)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return c.client.Do(req)
}

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

func setToStringSlice(ctx context.Context, v types.Set, diags *diag.Diagnostics) []string {
	if v.IsNull() || v.IsUnknown() {
		return []string{}
	}
	var out []string
	diags.Append(v.ElementsAs(ctx, &out, false)...)
	return out
}

func diffSlices(old, new []string) (toAdd, toDel []string) {
	want := map[string]struct{}{}
	have := map[string]struct{}{}
	for _, n := range new {
		want[n] = struct{}{}
	}
	for _, o := range old {
		have[o] = struct{}{}
	}
	for d := range want {
		if _, ok := have[d]; !ok {
			toAdd = append(toAdd, d)
		}
	}
	for d := range have {
		if _, ok := want[d]; !ok {
			toDel = append(toDel, d)
		}
	}
	return
}

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
	destElems := []attr.Value{}
	for _, d := range data.DestinationLists {
		destElems = append(destElems, types.StringValue(d))
	}
	plan.DestinationLists, _ = types.SetValue(types.StringType, destElems)

	// Set applications
	appElems := []attr.Value{}
	for _, a := range data.Applications {
		appElems = append(appElems, types.StringValue(a))
	}
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
	destElems := []attr.Value{}
	for _, d := range rule.DestinationLists {
		destElems = append(destElems, types.StringValue(d))
	}
	state.DestinationLists, _ = types.SetValue(types.StringType, destElems)

	// Set applications
	appElems := []attr.Value{}
	for _, a := range rule.Applications {
		appElems = append(appElems, types.StringValue(a))
	}
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

// Helper function to compare string slices
func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	aMap := make(map[string]bool)
	for _, v := range a {
		aMap[v] = true
	}
	for _, v := range b {
		if !aMap[v] {
			return false
		}
	}
	return true
}
