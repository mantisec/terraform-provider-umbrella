package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// CreateSite creates a new site
func (c *GeneratedClient) CreateSite(ctx context.Context, payload map[string]interface{}) (map[string]interface{}, error) {
	path := "/deployments/v2/sites"

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

	// Validate response status (Sites API returns 200 for create)
	if err := c.validateResponse(resp, 200); err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result, nil
}

// GetSite retrieves a site by ID
func (c *GeneratedClient) GetSite(ctx context.Context, siteID string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/deployments/v2/sites/%s", siteID)

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

// UpdateSite updates an existing site
func (c *GeneratedClient) UpdateSite(ctx context.Context, siteID string, payload map[string]interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/deployments/v2/sites/%s", siteID)

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

// DeleteSite deletes a site
func (c *GeneratedClient) DeleteSite(ctx context.Context, siteID string) error {
	path := fmt.Sprintf("/deployments/v2/sites/%s", siteID)

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

	// Validate response status (Sites API returns 204 for successful delete)
	if err := c.validateResponse(resp, 204); err != nil {
		return err
	}

	return nil
}

// ListSites retrieves all sites
func (c *GeneratedClient) ListSites(ctx context.Context) ([]map[string]interface{}, error) {
	path := "/deployments/v2/sites"

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
