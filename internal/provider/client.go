package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type apiClient struct {
	orgID   string
	baseURL string
	token   string
	http    *http.Client
}

func newAPIClient(ctx context.Context, apiKey, apiSecret, orgID string) (*apiClient, error) {
	c := &apiClient{
		orgID:   orgID,
		baseURL: "https://api.umbrella.com",
		http:    &http.Client{Timeout: 15 * time.Second},
	}

	// Exchange API key/secret for bearer token
	payload := map[string]string{"grant_type": "client_credentials", "client_id": apiKey, "client_secret": apiSecret}
	b, _ := json.Marshal(payload)
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/auth/v2/token", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("token request failed: %s", res.Status)
	}
	var tok struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(res.Body).Decode(&tok); err != nil {
		return nil, err
	}
	c.token = tok.AccessToken
	return c, nil
}

func (c *apiClient) do(ctx context.Context, method, path string, body any, out any) error {
	var rdr io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}
		rdr = bytes.NewReader(b)
	}
	req, _ := http.NewRequestWithContext(ctx, method, c.baseURL+path, rdr)
	req.Header.Set("Authorization", "Bearer "+c.token)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	res, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		return fmt.Errorf("umbrella API %s %s: %s", method, path, res.Status)
	}
	if out != nil {
		return json.NewDecoder(res.Body).Decode(out)
	}
	return nil
}
