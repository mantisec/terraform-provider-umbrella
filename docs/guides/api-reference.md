---
page_title: "API Reference - Umbrella Provider"
subcategory: "Guides"
description: |-
  Complete API reference for the Terraform Provider for Cisco Umbrella, including all supported endpoints and operations.
---

# API Reference

This document provides a comprehensive reference for all Cisco Umbrella API endpoints used by the Terraform Provider, including request/response formats, authentication requirements, and usage patterns.

## Base URL and Authentication

### Base URL
```
https://api.umbrella.com
```

### Authentication
All API requests use OAuth2 client credentials flow:

```http
POST /auth/v2/token
Content-Type: application/x-www-form-urlencoded
Authorization: Basic {base64(client_id:client_secret)}

grant_type=client_credentials
```

**Response:**
```json
{
  "access_token": "eyJhbGciOiJSUzI1NiIs...",
  "token_type": "Bearer",
  "expires_in": 3600
}
```

## Sites API

### List Sites
```http
GET /deployments/v2/sites
Authorization: Bearer {access_token}
```

**Response:**
```json
[
  {
    "siteId": 12345,
    "originId": 12345,
    "name": "Main Office",
    "type": "SITE",
    "isDefault": true,
    "internalNetworkCount": 2,
    "vaCount": 1,
    "createdAt": "2024-01-15T10:30:00Z",
    "modifiedAt": "2024-01-15T10:30:00Z"
  }
]
```

### Get Site
```http
GET /deployments/v2/sites/{siteId}
Authorization: Bearer {access_token}
```

### Create Site
```http
POST /deployments/v2/sites
Authorization: Bearer {access_token}
Content-Type: application/json

{
  "name": "New Site",
  "isDefault": false
}
```

### Update Site
```http
PUT /deployments/v2/sites/{siteId}
Authorization: Bearer {access_token}
Content-Type: application/json

{
  "name": "Updated Site Name",
  "isDefault": false
}
```

### Delete Site
```http
DELETE /deployments/v2/sites/{siteId}
Authorization: Bearer {access_token}
```

## Networks API

### List Networks
```http
GET /deployments/v2/networks
Authorization: Bearer {access_token}
```

**Response:**
```json
[
  {
    "networkId": 67890,
    "originId": 67890,
    "name": "Corporate Network",
    "ipAddress": "10.0.0.0",
    "prefixLength": 16,
    "status": "OPEN",
    "isDynamic": false,
    "isVerified": true,
    "createdAt": "2024-01-15T10:30:00Z"
  }
]
```

### Get Network
```http
GET /deployments/v2/networks/{networkId}
Authorization: Bearer {access_token}
```

### Create Network
```http
POST /deployments/v2/networks
Authorization: Bearer {access_token}
Content-Type: application/json

{
  "name": "New Network",
  "ipAddress": "192.168.1.0",
  "prefixLength": 30,
  "isDynamic": false,
  "status": "OPEN"
}
```

**Note:** IP ranges must be verified by Cisco Support before creation.

### Update Network
```http
PUT /deployments/v2/networks/{networkId}
Authorization: Bearer {access_token}
Content-Type: application/json

{
  "name": "Updated Network Name",
  "status": "CLOSED"
}
```

### Delete Network
```http
DELETE /deployments/v2/networks/{networkId}
Authorization: Bearer {access_token}
```

## Internal Networks API

### List Internal Networks
```http
GET /deployments/v2/internalnetworks
Authorization: Bearer {access_token}
```

**Response:**
```json
[
  {
    "internalNetworkId": 11111,
    "originId": 11111,
    "name": "Office Network",
    "ipAddress": "192.168.1.0",
    "prefixLength": 24,
    "siteId": 12345,
    "siteName": "Main Office",
    "networkId": null,
    "networkName": null,
    "tunnelId": null,
    "tunnelName": null,
    "createdAt": "2024-01-15T10:30:00Z",
    "modifiedAt": "2024-01-15T10:30:00Z"
  }
]
```

### Get Internal Network
```http
GET /deployments/v2/internalnetworks/{internalNetworkId}
Authorization: Bearer {access_token}
```

### Create Internal Network
```http
POST /deployments/v2/internalnetworks
Authorization: Bearer {access_token}
Content-Type: application/json

{
  "name": "New Internal Network",
  "ipAddress": "192.168.2.0",
  "prefixLength": 24,
  "siteId": 12345
}
```

**Note:** Exactly one of `siteId`, `networkId`, or `tunnelId` must be provided.

### Update Internal Network
```http
PUT /deployments/v2/internalnetworks/{internalNetworkId}
Authorization: Bearer {access_token}
Content-Type: application/json

{
  "name": "Updated Internal Network Name"
}
```

### Delete Internal Network
```http
DELETE /deployments/v2/internalnetworks/{internalNetworkId}
Authorization: Bearer {access_token}
```

## Users API

### List Users
```http
GET /admin/v2/organizations/{orgId}/users
Authorization: Bearer {access_token}
```

**Response:**
```json
[
  {
    "userId": 22222,
    "email": "admin@example.com",
    "firstname": "John",
    "lastname": "Doe",
    "timezone": "America/New_York",
    "role": "Full Admin",
    "roleId": 1,
    "status": "on",
    "twoFactorEnabled": true,
    "lastLoginTime": "2024-01-15T09:30:00Z"
  }
]
```

### Get User
```http
GET /admin/v2/organizations/{orgId}/users/{userId}
Authorization: Bearer {access_token}
```

### Create User
```http
POST /admin/v2/organizations/{orgId}/users
Authorization: Bearer {access_token}
Content-Type: application/json

{
  "email": "newuser@example.com",
  "firstname": "Jane",
  "lastname": "Smith",
  "password": "SecurePassword123!",
  "roleId": 2,
  "timezone": "UTC"
}
```

### Delete User
```http
DELETE /admin/v2/organizations/{orgId}/users/{userId}
Authorization: Bearer {access_token}
```

**Note:** The Users API does not support updates. Users must be deleted and recreated to modify.

## Destination Lists API

### List Destination Lists
```http
GET /policies/v2/organizations/{orgId}/destinationlists
Authorization: Bearer {access_token}
```

**Response:**
```json
[
  {
    "id": 33333,
    "name": "Blocked Domains",
    "access": "block",
    "isGlobal": false,
    "destinations": [
      "malicious-site.com",
      "phishing-domain.net"
    ],
    "createdAt": "2024-01-15T10:30:00Z",
    "modifiedAt": "2024-01-15T10:30:00Z"
  }
]
```

### Get Destination List
```http
GET /policies/v2/organizations/{orgId}/destinationlists/{id}
Authorization: Bearer {access_token}
```

### Create Destination List
```http
POST /policies/v2/organizations/{orgId}/destinationlists
Authorization: Bearer {access_token}
Content-Type: application/json

{
  "name": "New Destination List",
  "access": "block",
  "isGlobal": false,
  "destinations": [
    "example-malicious.com",
    "192.168.1.100",
    "10.0.0.0/8"
  ]
}
```

### Update Destination List
```http
PUT /policies/v2/organizations/{orgId}/destinationlists/{id}
Authorization: Bearer {access_token}
Content-Type: application/json

{
  "name": "Updated Destination List",
  "access": "allow",
  "destinations": [
    "trusted-site.com",
    "partner-domain.net"
  ]
}
```

### Delete Destination List
```http
DELETE /policies/v2/organizations/{orgId}/destinationlists/{id}
Authorization: Bearer {access_token}
```

## IPSec Tunnels API

### List Tunnels
```http
GET /v2/organizations/{orgId}/secureinternetgateway/ipsec/sites
Authorization: Bearer {access_token}
```

**Response:**
```json
[
  {
    "id": "44444444-4444-4444-4444-444444444444",
    "name": "Primary Tunnel",
    "siteOriginId": 12345,
    "deviceIp": "203.0.113.10",
    "localNetworks": ["192.168.1.0/24", "192.168.2.0/24"],
    "tunnelType": "IPSEC",
    "status": "ACTIVE",
    "tunnelEndpoint": "198.51.100.10",
    "createdAt": "2024-01-15T10:30:00Z",
    "updatedAt": "2024-01-15T10:30:00Z"
  }
]
```

### Get Tunnel
```http
GET /v2/organizations/{orgId}/secureinternetgateway/ipsec/sites/{tunnelId}
Authorization: Bearer {access_token}
```

### Create Tunnel
```http
POST /v2/organizations/{orgId}/secureinternetgateway/ipsec/sites
Authorization: Bearer {access_token}
Content-Type: application/json

{
  "name": "New Tunnel",
  "siteOriginId": 12345,
  "deviceIp": "203.0.113.20",
  "preSharedKey": "secure-psk-change-me",
  "localNetworks": ["192.168.3.0/24"],
  "tunnelType": "IPSEC"
}
```

### Update Tunnel
```http
PUT /v2/organizations/{orgId}/secureinternetgateway/ipsec/sites/{tunnelId}
Authorization: Bearer {access_token}
Content-Type: application/json

{
  "name": "Updated Tunnel Name",
  "localNetworks": ["192.168.3.0/24", "192.168.4.0/24"]
}
```

### Delete Tunnel
```http
DELETE /v2/organizations/{orgId}/secureinternetgateway/ipsec/sites/{tunnelId}
Authorization: Bearer {access_token}
```

## Error Responses

### Common HTTP Status Codes

| Status Code | Description | Common Causes |
|-------------|-------------|---------------|
| `400` | Bad Request | Invalid request format, missing required fields |
| `401` | Unauthorized | Invalid or expired authentication token |
| `403` | Forbidden | Insufficient permissions, unverified IP ranges |
| `404` | Not Found | Resource does not exist |
| `409` | Conflict | Resource name already exists, constraint violation |
| `429` | Too Many Requests | Rate limit exceeded |
| `500` | Internal Server Error | Server-side error |

### Error Response Format

```json
{
  "error": {
    "code": "INVALID_REQUEST",
    "message": "The request is invalid",
    "details": [
      {
        "field": "name",
        "message": "Name is required"
      }
    ]
  }
}
```

### Common Error Scenarios

#### Authentication Errors
```json
{
  "error": {
    "code": "UNAUTHORIZED",
    "message": "Invalid or expired token"
  }
}
```

#### Permission Errors
```json
{
  "error": {
    "code": "FORBIDDEN",
    "message": "Insufficient permissions for this operation"
  }
}
```

#### Validation Errors
```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Request validation failed",
    "details": [
      {
        "field": "prefixLength",
        "message": "Prefix length must be between 29 and 32"
      }
    ]
  }
}
```

#### Resource Not Found
```json
{
  "error": {
    "code": "NOT_FOUND",
    "message": "Resource not found"
  }
}
```

## Rate Limiting

### Rate Limits
- **Default**: 100 requests per minute per API key
- **Burst**: Up to 10 requests per second
- **Headers**: Rate limit information is included in response headers

### Rate Limit Headers
```http
X-RateLimit-Limit: 100
X-RateLimit-Remaining: 95
X-RateLimit-Reset: 1642248000
```

### Handling Rate Limits
The provider automatically handles rate limiting with exponential backoff:

1. **Initial delay**: 1 second
2. **Maximum delay**: 60 seconds
3. **Backoff multiplier**: 2x
4. **Maximum retries**: 5

## Data Types and Validation

### Common Field Types

| Field Type | Validation Rules | Examples |
|------------|------------------|----------|
| `name` | 1-255 characters, unique within resource type | "Main Office", "Corporate Network" |
| `email` | Valid email format, unique within organization | "user@example.com" |
| `ipAddress` | Valid IPv4 address | "192.168.1.0", "10.0.0.0" |
| `prefixLength` | Integer 8-32 (networks: 29-32) | 24, 30, 16 |
| `timezone` | Valid timezone identifier | "UTC", "America/New_York" |
| `access` | "allow" or "block" | "allow", "block" |
| `status` | "OPEN" or "CLOSED" | "OPEN", "CLOSED" |

### Destination Types

| Type | Format | Examples |
|------|--------|----------|
| Domain | Valid domain name | "example.com", "sub.domain.org" |
| IP Address | Valid IPv4 address | "192.168.1.1", "203.0.113.10" |
| CIDR Block | Valid CIDR notation | "10.0.0.0/8", "192.168.1.0/24" |
| URL | Valid URL with protocol | "https://example.com/path" |

## SDK and Client Libraries

### Official Support
The Terraform Provider uses the official Cisco Umbrella REST API. No additional SDK is required.

### Custom Implementations
For custom integrations, you can use the same API endpoints documented here with any HTTP client library.

### Authentication Flow Example (Go)
```go
type AuthResponse struct {
    AccessToken string `json:"access_token"`
    TokenType   string `json:"token_type"`
    ExpiresIn   int    `json:"expires_in"`
}

func authenticate(apiKey, apiSecret string) (*AuthResponse, error) {
    client := &http.Client{}
    
    data := url.Values{}
    data.Set("grant_type", "client_credentials")
    
    req, err := http.NewRequest("POST", "https://api.umbrella.com/auth/v2/token", 
        strings.NewReader(data.Encode()))
    if err != nil {
        return nil, err
    }
    
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.SetBasicAuth(apiKey, apiSecret)
    
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    var authResp AuthResponse
    if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
        return nil, err
    }
    
    return &authResp, nil
}
```

## API Versioning

### Current Versions
- **Authentication**: v2
- **Deployments**: v2
- **Policies**: v2
- **Admin**: v2
- **Secure Internet Gateway**: v2

### Version Headers
```http
Accept: application/json
Content-Type: application/json
```

### Deprecation Policy
Cisco follows a standard deprecation policy:
1. **Notice**: 6 months advance notice for breaking changes
2. **Support**: Deprecated versions supported for 12 months
3. **Migration**: Migration guides provided for major version changes

## Best Practices

### API Usage
1. **Cache tokens**: Reuse access tokens until expiration
2. **Handle rate limits**: Implement exponential backoff
3. **Validate input**: Check data before sending requests
4. **Monitor usage**: Track API calls and quotas

### Security
1. **Secure credentials**: Store API keys securely
2. **Rotate keys**: Regularly rotate API credentials
3. **Least privilege**: Use minimum required permissions
4. **Audit access**: Monitor API key usage

### Performance
1. **Batch operations**: Group related operations when possible
2. **Parallel requests**: Use appropriate parallelism levels
3. **Efficient polling**: Use reasonable intervals for status checks
4. **Connection reuse**: Reuse HTTP connections when possible

This API reference provides the foundation for understanding how the Terraform Provider interacts with the Cisco Umbrella API. For implementation details and examples, refer to the provider source code and documentation.