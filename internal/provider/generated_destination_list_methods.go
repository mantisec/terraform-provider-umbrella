package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// CreateDestinationList creates a new destination list
func (c *GeneratedClient) CreateDestinationList(ctx context.Context, payload map[string]interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/policies/v2/organizations/%s/destinationlists", c.orgID)

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

	// Validate response status
	if err := c.validateResponse(resp, 201); err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result, nil
}

// GetDestinationList retrieves a destination list by ID
func (c *GeneratedClient) GetDestinationList(ctx context.Context, id string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/policies/v2/organizations/%s/destinationlists/%s", c.orgID, id)

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

// UpdateDestinationList updates an existing destination list
func (c *GeneratedClient) UpdateDestinationList(ctx context.Context, id string, payload map[string]interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/policies/v2/organizations/%s/destinationlists/%s", c.orgID, id)

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

// DeleteDestinationList deletes a destination list
func (c *GeneratedClient) DeleteDestinationList(ctx context.Context, id string) error {
	path := fmt.Sprintf("/policies/v2/organizations/%s/destinationlists/%s", c.orgID, id)

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

	// Validate response status
	if err := c.validateResponse(resp, 204); err != nil {
		return err
	}

	return nil
}

// ListDestinationLists retrieves all destination lists
func (c *GeneratedClient) ListDestinationLists(ctx context.Context) ([]map[string]interface{}, error) {
	path := fmt.Sprintf("/policies/v2/organizations/%s/destinationlists", c.orgID)

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
