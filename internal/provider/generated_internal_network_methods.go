package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// CreateInternalNetwork creates a new internal network
func (c *GeneratedClient) CreateInternalNetwork(ctx context.Context, payload map[string]interface{}) (map[string]interface{}, error) {
	path := "/deployments/v2/internalnetworks"

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Log the request
	c.logRequest("POST", path, body)

	// Clear cache for write operations
	c.clearCacheForPath(path)
	resp, err := c.do(ctx, "POST", path, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Log the response
	c.logResponse(resp)

	// Validate response status (Internal Networks API returns 200 for create)
	if err := c.validateResponse(resp, 200); err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result, nil
}

// GetInternalNetwork retrieves an internal network by ID
func (c *GeneratedClient) GetInternalNetwork(ctx context.Context, internalNetworkID string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/deployments/v2/internalnetworks/%s", internalNetworkID)

	// Log the request
	c.logRequest("GET", path, nil)

	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, nil, 5*time.Minute)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Log the response
	c.logResponse(resp)

	// Validate response status
	if err := c.validateResponse(resp, 200); err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result, nil
}

// UpdateInternalNetwork updates an existing internal network
func (c *GeneratedClient) UpdateInternalNetwork(ctx context.Context, internalNetworkID string, payload map[string]interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/deployments/v2/internalnetworks/%s", internalNetworkID)

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Log the request
	c.logRequest("PUT", path, body)

	// Clear cache for write operations
	c.clearCacheForPath(path)
	resp, err := c.do(ctx, "PUT", path, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Log the response
	c.logResponse(resp)

	// Validate response status
	if err := c.validateResponse(resp, 200); err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result, nil
}

// DeleteInternalNetwork deletes an internal network
func (c *GeneratedClient) DeleteInternalNetwork(ctx context.Context, internalNetworkID string) error {
	path := fmt.Sprintf("/deployments/v2/internalnetworks/%s", internalNetworkID)

	// Log the request
	c.logRequest("DELETE", path, nil)

	// Clear cache for write operations
	c.clearCacheForPath(path)
	resp, err := c.do(ctx, "DELETE", path, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Log the response
	c.logResponse(resp)

	// Validate response status (Internal Networks API returns 204 for successful delete)
	if err := c.validateResponse(resp, 204); err != nil {
		return err
	}

	return nil
}

// ListInternalNetworks retrieves all internal networks
func (c *GeneratedClient) ListInternalNetworks(ctx context.Context) ([]map[string]interface{}, error) {
	path := "/deployments/v2/internalnetworks"

	// Log the request
	c.logRequest("GET", path, nil)

	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, nil, 5*time.Minute)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Log the response
	c.logResponse(resp)

	// Validate response status
	if err := c.validateResponse(resp, 200); err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result, nil
}

// GetInternalNetworkPolicies retrieves policies for an internal network
func (c *GeneratedClient) GetInternalNetworkPolicies(ctx context.Context, internalNetworkID string, policyType string) ([]map[string]interface{}, error) {
	path := fmt.Sprintf("/deployments/v2/internalnetworks/%s/policies", internalNetworkID)

	// Add query parameter for policy type if specified
	if policyType != "" {
		path += fmt.Sprintf("?type=%s", policyType)
	}

	// Log the request
	c.logRequest("GET", path, nil)

	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, nil, 5*time.Minute)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Log the response
	c.logResponse(resp)

	// Validate response status
	if err := c.validateResponse(resp, 200); err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result, nil
}
