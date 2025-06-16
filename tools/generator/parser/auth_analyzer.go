package parser

import (
	"strings"
)

// AuthAnalyzer handles analysis of authentication and authorization requirements
type AuthAnalyzer struct{}

// NewAuthAnalyzer creates a new auth analyzer
func NewAuthAnalyzer() *AuthAnalyzer {
	return &AuthAnalyzer{}
}

// AuthInfo contains authentication and authorization information
type AuthInfo struct {
	SecuritySchemes map[string]SecurityScheme
	OAuth2Scopes    map[string]string
	RequiredScopes  map[string][]string // endpoint -> required scopes
}

// AnalyzeAuth analyzes authentication and authorization requirements from the API spec
func (a *AuthAnalyzer) AnalyzeAuth(spec *APISpec) (*AuthInfo, error) {
	authInfo := &AuthInfo{
		SecuritySchemes: make(map[string]SecurityScheme),
		OAuth2Scopes:    make(map[string]string),
		RequiredScopes:  make(map[string][]string),
	}

	// Extract security schemes from components
	if spec.Components.SecuritySchemes != nil {
		for name, scheme := range spec.Components.SecuritySchemes {
			authInfo.SecuritySchemes[name] = scheme

			// Extract OAuth2 scopes if present
			if scheme.Type == "oauth2" && scheme.Flows != nil {
				a.extractOAuth2Scopes(scheme.Flows, authInfo.OAuth2Scopes)
			}
		}
	}

	// Analyze security requirements for each endpoint
	for path, pathItem := range spec.Paths {
		operations := map[string]*Operation{
			"GET":    pathItem.Get,
			"POST":   pathItem.Post,
			"PUT":    pathItem.Put,
			"DELETE": pathItem.Delete,
			"PATCH":  pathItem.Patch,
		}

		for method, operation := range operations {
			if operation == nil {
				continue
			}

			endpointKey := method + " " + path
			scopes := a.extractRequiredScopes(operation, spec.Security)
			if len(scopes) > 0 {
				authInfo.RequiredScopes[endpointKey] = scopes
			}
		}
	}

	return authInfo, nil
}

// extractOAuth2Scopes extracts OAuth2 scopes from flows
func (a *AuthAnalyzer) extractOAuth2Scopes(flows *OAuthFlows, scopesMap map[string]string) {
	if flows.ClientCredentials != nil && flows.ClientCredentials.Scopes != nil {
		for scope, description := range flows.ClientCredentials.Scopes {
			scopesMap[scope] = description
		}
	}

	if flows.AuthorizationCode != nil && flows.AuthorizationCode.Scopes != nil {
		for scope, description := range flows.AuthorizationCode.Scopes {
			scopesMap[scope] = description
		}
	}

	if flows.Implicit != nil && flows.Implicit.Scopes != nil {
		for scope, description := range flows.Implicit.Scopes {
			scopesMap[scope] = description
		}
	}

	if flows.Password != nil && flows.Password.Scopes != nil {
		for scope, description := range flows.Password.Scopes {
			scopesMap[scope] = description
		}
	}
}

// extractRequiredScopes extracts required scopes for an operation
func (a *AuthAnalyzer) extractRequiredScopes(operation *Operation, globalSecurity []SecurityRequirement) []string {
	var allScopes []string

	// Check operation-level security first
	if len(operation.Security) > 0 {
		for _, secReq := range operation.Security {
			scopes := a.extractScopesFromRequirement(secReq)
			allScopes = append(allScopes, scopes...)
		}
	} else if len(globalSecurity) > 0 {
		// Fall back to global security
		for _, secReq := range globalSecurity {
			scopes := a.extractScopesFromRequirement(secReq)
			allScopes = append(allScopes, scopes...)
		}
	}

	// Remove duplicates
	return a.removeDuplicateScopes(allScopes)
}

// extractScopesFromRequirement extracts scopes from a security requirement
func (a *AuthAnalyzer) extractScopesFromRequirement(secReq SecurityRequirement) []string {
	var scopes []string

	for _, scopeList := range secReq {
		scopes = append(scopes, scopeList...)
	}

	return scopes
}

// removeDuplicateScopes removes duplicate scopes from a slice
func (a *AuthAnalyzer) removeDuplicateScopes(scopes []string) []string {
	seen := make(map[string]bool)
	var result []string

	for _, scope := range scopes {
		if !seen[scope] {
			seen[scope] = true
			result = append(result, scope)
		}
	}

	return result
}

// MapScopesToOperations maps OAuth2 scopes to CRUD operations
func (a *AuthAnalyzer) MapScopesToOperations(authInfo *AuthInfo) map[string][]string {
	scopeToOps := make(map[string][]string)

	for endpoint, scopes := range authInfo.RequiredScopes {
		parts := strings.SplitN(endpoint, " ", 2)
		if len(parts) != 2 {
			continue
		}

		method := parts[0]
		var operation string

		switch strings.ToUpper(method) {
		case "GET":
			operation = "read"
		case "POST":
			operation = "create"
		case "PUT", "PATCH":
			operation = "update"
		case "DELETE":
			operation = "delete"
		default:
			operation = "unknown"
		}

		for _, scope := range scopes {
			if !a.containsString(scopeToOps[scope], operation) {
				scopeToOps[scope] = append(scopeToOps[scope], operation)
			}
		}
	}

	return scopeToOps
}

// containsString checks if a string slice contains a specific string
func (a *AuthAnalyzer) containsString(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// GetAuthenticationMethod determines the primary authentication method
func (a *AuthAnalyzer) GetAuthenticationMethod(authInfo *AuthInfo) string {
	// Priority: OAuth2 > API Key > Basic Auth > None
	for _, scheme := range authInfo.SecuritySchemes {
		switch scheme.Type {
		case "oauth2":
			return "oauth2"
		case "apiKey":
			return "api_key"
		case "http":
			if scheme.Scheme == "basic" {
				return "basic_auth"
			} else if scheme.Scheme == "bearer" {
				return "bearer_token"
			}
		}
	}

	return "none"
}

// IsSecurityRequired checks if any security is required for the API
func (a *AuthAnalyzer) IsSecurityRequired(authInfo *AuthInfo) bool {
	return len(authInfo.SecuritySchemes) > 0 || len(authInfo.RequiredScopes) > 0
}
