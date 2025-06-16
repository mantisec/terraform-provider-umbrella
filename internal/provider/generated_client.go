package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"strings"
	"time"
)

// GeneratedClient extends the base apiClient with generated methods for all API endpoints
type GeneratedClient struct {
	*apiClient
	cache map[string]cacheEntry
}

// cacheEntry represents a cached response for read-only operations
type cacheEntry struct {
	data      interface{}
	timestamp time.Time
	ttl       time.Duration
}

// NewGeneratedClient creates a new generated client with caching support
func NewGeneratedClient(ctx context.Context, key, secret, orgID string) (*GeneratedClient, error) {
	baseClient, err := newAPIClient(ctx, key, secret, orgID)
	if err != nil {
		return nil, err
	}

	return &GeneratedClient{
		apiClient: baseClient,
		cache:     make(map[string]cacheEntry),
	}, nil
}

// getCachedResponse retrieves a cached response if valid
func (c *GeneratedClient) getCachedResponse(key string) (interface{}, bool) {
	entry, exists := c.cache[key]
	if !exists {
		return nil, false
	}

	if time.Since(entry.timestamp) > entry.ttl {
		delete(c.cache, key)
		return nil, false
	}

	return entry.data, true
}

// setCachedResponse stores a response in cache
func (c *GeneratedClient) setCachedResponse(key string, data interface{}, ttl time.Duration) {
	c.cache[key] = cacheEntry{
		data:      data,
		timestamp: time.Now(),
		ttl:       ttl,
	}
}

// doWithCache performs a request with optional caching for GET operations
func (c *GeneratedClient) doWithCache(ctx context.Context, method, path string, body []byte, cacheTTL time.Duration) (*http.Response, error) {
	// For GET requests, check cache first
	if method == "GET" && cacheTTL > 0 {
		cacheKey := fmt.Sprintf("%s:%s", method, path)
		if cached, found := c.getCachedResponse(cacheKey); found {
			// Return a mock response with cached data
			return c.createCachedResponse(cached)
		}
	}

	// Perform the actual request
	resp, err := c.do(ctx, method, path, body)
	if err != nil {
		return nil, err
	}

	// Cache successful GET responses
	if method == "GET" && cacheTTL > 0 && resp.StatusCode == http.StatusOK {
		// Read and cache the response body
		var responseData interface{}
		if err := json.NewDecoder(resp.Body).Decode(&responseData); err == nil {
			cacheKey := fmt.Sprintf("%s:%s", method, path)
			c.setCachedResponse(cacheKey, responseData, cacheTTL)
		}
		resp.Body.Close()

		// Return a new response with the cached data
		return c.createCachedResponse(responseData)
	}

	return resp, nil
}

// createCachedResponse creates an HTTP response from cached data
func (c *GeneratedClient) createCachedResponse(data interface{}) (*http.Response, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal cached data: %w", err)
	}

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Status:     "200 OK",
		Body:       &responseBody{data: jsonData},
		Header:     make(http.Header),
	}
	resp.Header.Set("Content-Type", "application/json")

	return resp, nil
}

// responseBody implements io.ReadCloser for cached responses
type responseBody struct {
	data   []byte
	offset int
}

func (rb *responseBody) Read(p []byte) (n int, err error) {
	if rb.offset >= len(rb.data) {
		return 0, fmt.Errorf("EOF")
	}

	n = copy(p, rb.data[rb.offset:])
	rb.offset += n
	return n, nil
}

func (rb *responseBody) Close() error {
	return nil
}

// resolveScopeForEndpoint determines the required OAuth2 scope for an endpoint
func (c *GeneratedClient) resolveScopeForEndpoint(method, path string) []string {
	// Default scope mappings based on HTTP method
	switch method {
	case "GET":
		return []string{"read"}
	case "POST":
		return []string{"write"}
	case "PUT", "PATCH":
		return []string{"write"}
	case "DELETE":
		return []string{"delete"}
	default:
		return []string{"read"}
	}
}

// validateResponse validates the HTTP response and handles common error cases
func (c *GeneratedClient) validateResponse(resp *http.Response, expectedStatuses ...int) error {
	if len(expectedStatuses) == 0 {
		expectedStatuses = []int{http.StatusOK}
	}

	for _, status := range expectedStatuses {
		if resp.StatusCode == status {
			return nil
		}
	}

	// Handle common error responses
	switch resp.StatusCode {
	case http.StatusUnauthorized:
		return fmt.Errorf("authentication failed: %s", resp.Status)
	case http.StatusForbidden:
		return fmt.Errorf("insufficient permissions: %s", resp.Status)
	case http.StatusNotFound:
		return fmt.Errorf("resource not found: %s", resp.Status)
	case http.StatusTooManyRequests:
		return fmt.Errorf("rate limit exceeded: %s", resp.Status)
	case http.StatusInternalServerError:
		return fmt.Errorf("server error: %s", resp.Status)
	default:
		return fmt.Errorf("API request failed with status: %s", resp.Status)
	}
}

// buildPath constructs the full API path with organization ID and parameters
func (c *GeneratedClient) buildPath(pathTemplate string, pathParams map[string]string) string {
	path := pathTemplate

	// Replace organization ID placeholder
	path = strings.ReplaceAll(path, "{orgId}", c.orgID)
	path = strings.ReplaceAll(path, "{organizationId}", c.orgID)

	// Replace other path parameters
	for key, value := range pathParams {
		placeholder := fmt.Sprintf("{%s}", key)
		path = strings.ReplaceAll(path, placeholder, value)
	}

	return path
}

// logRequest logs API requests for debugging
func (c *GeneratedClient) logRequest(method, path string, body []byte) {
	// In a production implementation, this would use a proper logger
	// For now, we'll keep it simple
	if len(body) > 0 {
		fmt.Printf("API Request: %s %s (body: %d bytes)\n", method, path, len(body))
	} else {
		fmt.Printf("API Request: %s %s\n", method, path)
	}
}

// logResponse logs API responses for debugging
func (c *GeneratedClient) logResponse(resp *http.Response) {
	fmt.Printf("API Response: %s\n", resp.Status)
}

// clearCache clears all cached responses
func (c *GeneratedClient) clearCache() {
	c.cache = make(map[string]cacheEntry)
}

// clearCacheForPath clears cached responses for a specific path pattern
func (c *GeneratedClient) clearCacheForPath(pathPattern string) {
	for key := range c.cache {
		if strings.Contains(key, pathPattern) {
			delete(c.cache, key)
		}
	}
}
