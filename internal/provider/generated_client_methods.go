package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// ListTopCategories List the categories that received the greatest number of requests.
Order the number of requests in descending order.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) ListTopCategories(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/top-categories"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListThreatNames List the threat names.

**Access Scope:** Reports > Utilities > Read-Only
func (c *GeneratedClient) ListThreatNames(ctx context.Context) (map[string]interface{}, error) {
	path := "/threat-names"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListProvidersCategories List the Provider categories

**Access Scope:** Reports > Customer > Read-Only
func (c *GeneratedClient) ListProvidersCategories(ctx context.Context) (map[string]interface{}, error) {
	path := "/providers/categories"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListProvidersCategoryRequestsByOrg List the summary counts of all requests for each category within the timeframe.

**Access Scope:** Reports > Customer > Read-Only
func (c *GeneratedClient) ListProvidersCategoryRequestsByOrg(ctx context.Context,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/providers/category-requests-by-org"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListActivityProxy List all proxy entries within the timeframe.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) ListActivityProxy(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/activity/proxy"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListActivityFirewall List all firewall activity within the timeframe.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) ListActivityFirewall(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/activity/firewall"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListDeploymentStatus List the deployment status within the timeframe.

**Access Scope:** Reports > Granular Events > Read-Only
func (c *GeneratedClient) ListDeploymentStatus(ctx context.Context,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/deployment-status"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// GetThreatTypes List the threat types by threat ID.

**Access Scope:** Reports > Utilities > Read-Only
func (c *GeneratedClient) GetThreatTypes(ctx context.Context,  string) (map[string]interface{}, error) {
	path := "/threat-types/{threattypeid}"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// GetSummary Get the summary of requests by the traffic type.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) GetSummary(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/summary/{type}"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListRequestsByHour List the activity volume within the timeframe.

**Access Scope:** Reports > Granular Events > Read-Only
func (c *GeneratedClient) ListRequestsByHour(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/requests-by-hour"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListTopThreatTypes List the top types of threat within the timeframe. Returns both DNS and proxy data.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) ListTopThreatTypes(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/top-threat-types"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListIdentities List the identities.

**Access Scope:** Reports > Utilities > Read-Only
func (c *GeneratedClient) ListIdentities(ctx context.Context,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/identities"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// CreateIdentities Get the identities information by providing a list of identity IDs in the request body.

**Access Scope:** Reports > Utilities > Read-Only
func (c *GeneratedClient) CreateIdentities(ctx context.Context,  string) (map[string]interface{}, error) {
	path := "/identities"
	var body []byte

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
	if err := c.validateResponse(resp, 200); err != nil {
		return nil, err
	}
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result, nil
}

// ListSummariesByDestination List the summaries by destination.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) ListSummariesByDestination(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/summaries-by-destination"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListProvidersRequestsByOrg List the summary counts of all requests within the timeframe.

**Access Scope:** Reports > Customer > Read-Only
func (c *GeneratedClient) ListProvidersRequestsByOrg(ctx context.Context,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/providers/requests-by-org"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// GetTopCategories List the categories for the type of traffic that received the greatest number of requests.
Order the number of requests in descending order.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) GetTopCategories(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/top-categories/{type}"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListCategoriesByTimerange List the activity volume within the timeframe by category.

**Access Scope:** Reports > Granular Events > Read-Only
func (c *GeneratedClient) ListCategoriesByTimerange(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/categories-by-timerange"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// GetThreatNames Get the threat name by threat ID.

**Access Scope:** Reports > Utilities > Read-Only
func (c *GeneratedClient) GetThreatNames(ctx context.Context,  string) (map[string]interface{}, error) {
	path := "/threat-names/{threatnameid}"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListTopIpsInternal List the top internal IP addresses.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) ListTopIpsInternal(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/top-ips/internal"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListSummary Get the summary report.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) ListSummary(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/summary"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListIdentityDistribution List the number of requests by identity types.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) ListIdentityDistribution(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/identity-distribution"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// GetCategoriesByHour List the activity volume for the type of category within the timeframe.

**Access Scope:** Reports > Granular Events > Read-Only
func (c *GeneratedClient) GetCategoriesByHour(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/categories-by-hour/{type}"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListApplications List the applications.

**Access Scope:** Reports > Utilities > Read-Only
func (c *GeneratedClient) ListApplications(ctx context.Context,  string) (map[string]interface{}, error) {
	path := "/applications"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListSummariesByCategory List the summaries of requests by category.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) ListSummariesByCategory(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/summaries-by-category"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListProvidersRequestsByHour List the activity volume within the timeframe.

**Access Scope:** Reports > Customer > Read-Only
func (c *GeneratedClient) ListProvidersRequestsByHour(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/providers/requests-by-hour"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// GetRequestsByTimerange List the activity volume within the timeframe.

**Access Scope:** Reports > Granular Events > Read-Only
func (c *GeneratedClient) GetRequestsByTimerange(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/requests-by-timerange/{type}"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListTopFiles List the top files within the timeframe. Only returns proxy data.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) ListTopFiles(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/top-files"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListBandwidthByHour List the bandwidth in bytes within the timeframe. Only returns proxy data.

**Access Scope:** Reports > Granular Events > Read-Only
func (c *GeneratedClient) ListBandwidthByHour(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/bandwidth-by-hour"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// GetIdentities Get identity by identity ID.

**Access Scope:** Reports > Utilities > Read-Only
func (c *GeneratedClient) GetIdentities(ctx context.Context,  string) (map[string]interface{}, error) {
	path := "/identities/{identityid}"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListProvidersRequestsByCategory List the summary counts of all requests within the timeframe.

**Access Scope:** Reports > Customer > Read-Only
func (c *GeneratedClient) ListProvidersRequestsByCategory(ctx context.Context,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/providers/requests-by-category"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListActivityIp (Deprecated) List all IP activity within the timeframe.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) ListActivityIp(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/activity/ip"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListTopUrls List the top number of URLs that are requested for a certain domain.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) ListTopUrls(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/top-urls"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// GetTopIdentities List the identities for the specific traffic type by the number of requests.
Sort the results in descending order.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) GetTopIdentities(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/top-identities/{type}"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListTotalRequests Get the count of the total requests.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) ListTotalRequests(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/total-requests"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// GetTotalRequests Get the count of the total requests for the request type.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) GetTotalRequests(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/total-requests/{type}"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListTopThreats Get the top threats within the timeframe. Returns both DNS and proxy data.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) ListTopThreats(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/top-threats"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListTopIps List the top IP addresses.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) ListTopIps(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/top-ips"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListTopIdentities List the identities by the number of requests made, sorted in descending order.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) ListTopIdentities(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/top-identities"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// GetTopThreatTypes List the top threat-types within the timeframe.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) GetTopThreatTypes(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/top-threat-types/{type}"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListCategories List the categories.

**Access Scope:** Reports > Utilities > Read-Only
func (c *GeneratedClient) ListCategories(ctx context.Context) (map[string]interface{}, error) {
	path := "/categories"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// GetIdentityDistribution List the number of requests by identity for the type of traffic.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) GetIdentityDistribution(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/identity-distribution/{type}"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListTopDestinations List the destinations by the number of requests made to this destination.
Return the results in descending order.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) ListTopDestinations(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/top-destinations"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListBandwidthByTimerange List the bandwidth in bytes within the timeframe. Only returns proxy data.

**Access Scope:** Reports > Granular Events > Read-Only
func (c *GeneratedClient) ListBandwidthByTimerange(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/bandwidth-by-timerange"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// GetTopThreats Get the top threats within the timeframe.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) GetTopThreats(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/top-threats/{type}"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListActivity List all activities (dns/proxy/firewall/intrusion) within the timeframe.
**Note:** The IP activity report is not available.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) ListActivity(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/activity"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListActivityAmpRetrospective List all AMP retrospective activity within the timeframe.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) ListActivityAmpRetrospective(ctx context.Context,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/activity/amp-retrospective"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListProvidersDeployments List the summary counts of deployment status for the organization within the timeframe.

**Access Scope:** Reports > Customer > Read-Only
func (c *GeneratedClient) ListProvidersDeployments(ctx context.Context,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/providers/deployments"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListCategoriesByHour List the activity volume within the timeframe by type of category.

**Access Scope:** Reports > Granular Events > Read-Only
func (c *GeneratedClient) ListCategoriesByHour(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/categories-by-hour"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListProvidersRequestsByTimerange List the activity volume within the timeframe.

**Access Scope:** Reports > Customer > Read-Only
func (c *GeneratedClient) ListProvidersRequestsByTimerange(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/providers/requests-by-timerange"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListProvidersRequestsByDestination List the summary counts of all requests within the timeframe.

**Access Scope:** Reports > Customer > Read-Only
func (c *GeneratedClient) ListProvidersRequestsByDestination(ctx context.Context,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/providers/requests-by-destination"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// GetTopDestinations List the destinations by type of destination and the number of requests made
to this destination. Return the collection in descending order.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) GetTopDestinations(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/top-destinations/{type}"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListTopEventtypes List the top event types by the number of requests made for each type of event.
Order the number of requests in descending order.
The valid event types are: `domain_security`, `domain_integration`,
`url_security`, `url_integration`, `cisco_amp`, and `antivirus`.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) ListTopEventtypes(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/top-eventtypes"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListTopDnsQueryTypes List the top types of DNS query.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) ListTopDnsQueryTypes(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/top-dns-query-types"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// GetRequestsByHour List the activity volume within the timeframe.

**Access Scope:** Reports > Granular Events > Read-Only
func (c *GeneratedClient) GetRequestsByHour(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/requests-by-hour/{type}"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListRequestsByTimerange List the activity volume within the timeframe.

**Access Scope:** Reports > Granular Events > Read-Only
func (c *GeneratedClient) ListRequestsByTimerange(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/requests-by-timerange"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListActivityDns List all DNS entries within the timeframe.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) ListActivityDns(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/activity/dns"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListThreatTypes List the threat types.

**Access Scope:** Reports > Utilities > Read-Only
func (c *GeneratedClient) ListThreatTypes(ctx context.Context) (map[string]interface{}, error) {
	path := "/threat-types"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// GetSummariesByCategory List the summaries by category for the type of request.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) GetSummariesByCategory(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/summaries-by-category/{type}"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// GetSummariesByDestination List the summaries by destination for the type of traffic.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) GetSummariesByDestination(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/summaries-by-destination/{type}"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListActivityIntrusion List all Intrusion Prevention System (IPS) activity within the timeframe.

**Access Scope:** Reports > Aggregations > Read-Only
func (c *GeneratedClient) ListActivityIntrusion(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/activity/intrusion"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// GetCategoriesByTimerange List the activity volume within the timeframe by category.

**Access Scope:** Reports > Granular Events > Read-Only
func (c *GeneratedClient) GetCategoriesByTimerange(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/categories-by-timerange/{type}"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

// ListSummariesByRuleIntrusion List the summaries by rule for the intrusion type.

**Access Scope:** Reports > Summaries by rule > Read-Only
func (c *GeneratedClient) ListSummariesByRuleIntrusion(ctx context.Context,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string,  string) (map[string]interface{}, error) {
	path := "/summaries-by-rule/intrusion"
	var body []byte

	// Log the request
	c.logRequest("GET", path, body)
	// Use caching for GET requests
	resp, err := c.doWithCache(ctx, "GET", path, body, 5 * time.Minute)
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

