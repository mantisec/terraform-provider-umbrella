package provider

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
)

// -----------------------------------------------------------------------------
// Constants
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
	c.expires = time.Now().Add(time.Duration(data.ExpiresIn-60) * time.Second) // refresh 1 min early
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
// Generic API Methods
// -----------------------------------------------------------------------------

// GenericAPIResponse represents a generic API response structure
type GenericAPIResponse struct {
	Status struct {
		Code int    `json:"code"`
		Text string `json:"text"`
	} `json:"status"`
	Data interface{} `json:"data"`
}

// CreateResource creates a new resource using the specified path and request body
func (c *apiClient) CreateResource(ctx context.Context, path string, requestBody interface{}) (*GenericAPIResponse, error) {
	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.do(ctx, http.MethodPost, path, body)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("API error: %d %s", resp.StatusCode, resp.Status)
	}

	var result GenericAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// GetResource retrieves a resource by ID using the specified path
func (c *apiClient) GetResource(ctx context.Context, path string) (*GenericAPIResponse, error) {
	resp, err := c.do(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("resource not found")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %d %s", resp.StatusCode, resp.Status)
	}

	var result GenericAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// UpdateResource updates a resource using the specified path and request body
func (c *apiClient) UpdateResource(ctx context.Context, path string, requestBody interface{}) (*GenericAPIResponse, error) {
	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.do(ctx, http.MethodPatch, path, body)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("resource not found")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %d %s", resp.StatusCode, resp.Status)
	}

	var result GenericAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// DeleteResource deletes a resource using the specified path
func (c *apiClient) DeleteResource(ctx context.Context, path string) error {
	resp, err := c.do(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		// Resource already deleted, consider this success
		return nil
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("API error: %d %s", resp.StatusCode, resp.Status)
	}

	return nil
}

// -----------------------------------------------------------------------------
// Destination Lists API Methods (Legacy - for backward compatibility)
// -----------------------------------------------------------------------------

// DestinationListCreateRequest represents a request to create a destination list
type DestinationListCreateRequest struct {
	Access       string `json:"access"`
	IsGlobal     bool   `json:"isGlobal"`
	Name         string `json:"name"`
	BundleTypeId *int64 `json:"bundleTypeId,omitempty"`
}

// DestinationListUpdateRequest represents a request to update a destination list
type DestinationListUpdateRequest struct {
	Name string `json:"name"`
}

// DestinationListObject represents a destination list object from the API
type DestinationListObject struct {
	Id                   int64  `json:"id"`
	OrganizationId       int64  `json:"organizationId"`
	Access               string `json:"access"`
	IsGlobal             bool   `json:"isGlobal"`
	Name                 string `json:"name"`
	ThirdpartyCategoryId int64  `json:"thirdpartyCategoryId"`
	CreatedAt            int64  `json:"createdAt"`
	ModifiedAt           int64  `json:"modifiedAt"`
	IsMspDefault         bool   `json:"isMspDefault"`
	MarkedForDeletion    bool   `json:"markedForDeletion"`
	BundleTypeId         int64  `json:"bundleTypeId"`
}

// DestinationListResponse represents the API response for destination list operations
type DestinationListResponse struct {
	Status struct {
		Code int    `json:"code"`
		Text string `json:"text"`
	} `json:"status"`
	Data DestinationListObject `json:"data"`
}

// CreateDestinationList creates a new destination list (legacy method)
func (c *apiClient) CreateDestinationList(ctx context.Context, req DestinationListCreateRequest) (*DestinationListResponse, error) {
	path := "/policies/v2/destinationlists"

	genericResp, err := c.CreateResource(ctx, path, req)
	if err != nil {
		return nil, err
	}

	// Convert generic response to specific type
	result := &DestinationListResponse{
		Status: genericResp.Status,
	}

	// Convert data interface{} to DestinationListObject
	if dataBytes, err := json.Marshal(genericResp.Data); err == nil {
		json.Unmarshal(dataBytes, &result.Data)
	}

	return result, nil
}

// GetDestinationList retrieves a destination list by ID (legacy method)
func (c *apiClient) GetDestinationList(ctx context.Context, id string) (*DestinationListResponse, error) {
	path := fmt.Sprintf("/policies/v2/destinationlists/%s", id)

	genericResp, err := c.GetResource(ctx, path)
	if err != nil {
		return nil, err
	}

	// Convert generic response to specific type
	result := &DestinationListResponse{
		Status: genericResp.Status,
	}

	// Convert data interface{} to DestinationListObject
	if dataBytes, err := json.Marshal(genericResp.Data); err == nil {
		json.Unmarshal(dataBytes, &result.Data)
	}

	return result, nil
}

// UpdateDestinationList updates a destination list (legacy method)
func (c *apiClient) UpdateDestinationList(ctx context.Context, id string, req DestinationListUpdateRequest) (*DestinationListResponse, error) {
	path := fmt.Sprintf("/policies/v2/destinationlists/%s", id)

	genericResp, err := c.UpdateResource(ctx, path, req)
	if err != nil {
		return nil, err
	}

	// Convert generic response to specific type
	result := &DestinationListResponse{
		Status: genericResp.Status,
	}

	// Convert data interface{} to DestinationListObject
	if dataBytes, err := json.Marshal(genericResp.Data); err == nil {
		json.Unmarshal(dataBytes, &result.Data)
	}

	return result, nil
}

// DeleteDestinationList deletes a destination list (legacy method)
func (c *apiClient) DeleteDestinationList(ctx context.Context, id string) error {
	path := fmt.Sprintf("/policies/v2/destinationlists/%s", id)
	return c.DeleteResource(ctx, path)
}
