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
