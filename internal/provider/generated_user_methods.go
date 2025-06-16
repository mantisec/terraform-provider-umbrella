package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// CreateUser creates a new user
func (c *GeneratedClient) CreateUser(ctx context.Context, payload map[string]interface{}) (map[string]interface{}, error) {
	path := "/admin/v2/users"

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

	// Validate response status (Users API returns 200 for create)
	if err := c.validateResponse(resp, 200); err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result, nil
}

// GetUser retrieves a user by ID
func (c *GeneratedClient) GetUser(ctx context.Context, userID string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/admin/v2/users/%s", userID)

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

// DeleteUser deletes a user
func (c *GeneratedClient) DeleteUser(ctx context.Context, userID string) error {
	path := fmt.Sprintf("/admin/v2/users/%s", userID)

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

	// Validate response status (Users API returns 204 for successful delete)
	if err := c.validateResponse(resp, 204); err != nil {
		return err
	}

	return nil
}

// ListUsers retrieves all users
func (c *GeneratedClient) ListUsers(ctx context.Context, page, limit int) ([]map[string]interface{}, error) {
	path := "/admin/v2/users"

	// Add query parameters if specified
	if page > 0 || limit > 0 {
		path += "?"
		if page > 0 {
			path += fmt.Sprintf("page=%d", page)
		}
		if limit > 0 {
			if page > 0 {
				path += "&"
			}
			path += fmt.Sprintf("limit=%d", limit)
		}
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

// ListRoles retrieves all available roles
func (c *GeneratedClient) ListRoles(ctx context.Context) ([]map[string]interface{}, error) {
	path := "/admin/v2/roles"

	// Log the request
	c.logRequest("GET", path, nil)

	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, nil, 15*time.Minute) // Cache roles longer as they change less frequently
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
