package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// CreateNetwork creates a new network
func (c *GeneratedClient) CreateNetwork(ctx context.Context, payload map[string]interface{}) (map[string]interface{}, error) {
	path := "/deployments/v2/networks"

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

	// Validate response status (Networks API returns 200 for create, not 201)
	if err := c.validateResponse(resp, 200); err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result, nil
}

// GetNetwork retrieves a network by ID
func (c *GeneratedClient) GetNetwork(ctx context.Context, networkID string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/deployments/v2/networks/%s", networkID)

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

// UpdateNetwork updates an existing network
func (c *GeneratedClient) UpdateNetwork(ctx context.Context, networkID string, payload map[string]interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/deployments/v2/networks/%s", networkID)

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

// DeleteNetwork deletes a network
func (c *GeneratedClient) DeleteNetwork(ctx context.Context, networkID string) error {
	path := fmt.Sprintf("/deployments/v2/networks/%s", networkID)

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

	// Validate response status (Networks API returns 204 for successful delete)
	if err := c.validateResponse(resp, 204); err != nil {
		return err
	}

	return nil
}

// ListNetworks retrieves all networks
func (c *GeneratedClient) ListNetworks(ctx context.Context) ([]map[string]interface{}, error) {
	path := "/deployments/v2/networks"

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

// GetNetworkPolicies retrieves policies for a network
func (c *GeneratedClient) GetNetworkPolicies(ctx context.Context, networkID string, policyType string) ([]map[string]interface{}, error) {
	path := fmt.Sprintf("/deployments/v2/networks/%s/policies", networkID)

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
