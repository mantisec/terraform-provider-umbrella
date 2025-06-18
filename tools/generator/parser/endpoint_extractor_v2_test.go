package parser

import (
	"path/filepath"
	"testing"
)

func TestEndpointExtractorV2_ExtractAndConsolidate(t *testing.T) {
	// Use the new adapter to parse a spec
	adapter := NewLibOpenAPIAdapter()
	specPath := filepath.Join("..", "..", "..", "api-specs", "cisco_umbrella_networks_api_2_0_0.yaml")

	spec, err := adapter.ParseFile(specPath)
	if err != nil {
		t.Fatalf("Failed to parse spec: %v", err)
	}

	// Extract and consolidate with the new extractor
	extractor := NewEndpointExtractorV2()
	groups, err := extractor.ExtractAndConsolidate(spec)
	if err != nil {
		t.Fatalf("Failed to extract and consolidate: %v", err)
	}

	// Generate report
	report := extractor.GetConsolidationReport(groups)
	t.Logf("Consolidation report: %s", report.String())

	// Verify we have consolidated groups
	if len(groups) == 0 {
		t.Error("Expected to have consolidated groups")
	}

	// Log details of each group
	for _, group := range groups {
		t.Logf("Group: %s", group.CanonicalName)
		t.Logf("  Endpoints: %d", len(group.Endpoints))
		t.Logf("  Has Resource: %t", group.HasResource)
		t.Logf("  Has DataSource: %t", group.HasDataSource)
		t.Logf("  CRUD Operations: %v", group.CRUDOps)

		for _, endpoint := range group.Endpoints {
			t.Logf("    %s %s -> %s (%s)",
				endpoint.Method, endpoint.Path, endpoint.CRUDType, endpoint.ResourceType)
		}
	}

	// Verify that we have fewer groups than individual endpoints
	totalEndpoints := 0
	for _, group := range groups {
		totalEndpoints += len(group.Endpoints)
	}

	if len(groups) >= totalEndpoints {
		t.Errorf("Expected consolidation to reduce groups: %d groups for %d endpoints",
			len(groups), totalEndpoints)
	}

	t.Logf("Successfully consolidated %d endpoints into %d groups", totalEndpoints, len(groups))
}

func TestEndpointExtractorV2_CanonicalNaming(t *testing.T) {
	extractor := NewEndpointExtractorV2()

	testCases := []struct {
		input    string
		expected string
	}{
		{"getNetworks", "network"},
		{"createNetwork", "network"},
		{"updateNetworkById", "network"},
		{"deleteNetwork", "network"},
		{"listUsers", "user"},
		{"getUserById", "user"},
		{"networks", "network"},
		{"policies", "policy"},
		{"destinationLists", "destination_list"},
		{"getDestinationLists", "destination_list"},
		{"createDestinationList", "destination_list"},
		{"sites", "site"},
		{"getSites", "site"},
		{"roamingComputers", "roaming_computer"},
		{"virtualAppliances", "virtual_appliance"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := extractor.normalizeToCanonical(tc.input)
			if result != tc.expected {
				t.Errorf("normalizeToCanonical(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestEndpointExtractorV2_Singularization(t *testing.T) {
	extractor := NewEndpointExtractorV2()

	testCases := []struct {
		input    string
		expected string
	}{
		{"networks", "network"},
		{"policies", "policy"},
		{"sites", "site"},
		{"users", "user"},
		{"addresses", "address"},
		{"boxes", "box"},
		{"wishes", "wish"},
		{"lives", "life"},
		{"knives", "knife"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := extractor.singularize(tc.input)
			if result != tc.expected {
				t.Errorf("singularize(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestEndpointExtractorV2_CompareWithOldExtractor(t *testing.T) {
	// Parse spec with new adapter
	adapter := NewLibOpenAPIAdapter()
	specPath := filepath.Join("..", "..", "..", "api-specs", "cisco_umbrella_networks_api_2_0_0.yaml")

	spec, err := adapter.ParseFile(specPath)
	if err != nil {
		t.Fatalf("Failed to parse spec: %v", err)
	}

	// Extract with old extractor
	oldExtractor := NewEndpointExtractor()
	oldEndpoints, err := oldExtractor.ExtractEndpoints(spec)
	if err != nil {
		t.Fatalf("Failed to extract with old extractor: %v", err)
	}
	oldGroups := oldExtractor.GroupEndpointsByResource(oldEndpoints)

	// Extract with new extractor
	newExtractor := NewEndpointExtractorV2()
	newGroups, err := newExtractor.ExtractAndConsolidate(spec)
	if err != nil {
		t.Fatalf("Failed to extract with new extractor: %v", err)
	}

	t.Logf("Old extractor: %d groups", len(oldGroups))
	for name, endpoints := range oldGroups {
		t.Logf("  %s: %d endpoints", name, len(endpoints))
	}

	t.Logf("New extractor: %d groups", len(newGroups))
	for name, group := range newGroups {
		t.Logf("  %s: %d endpoints (resource: %t, data-source: %t)",
			name, len(group.Endpoints), group.HasResource, group.HasDataSource)
	}

	// The new extractor should produce fewer or equal groups due to better consolidation
	if len(newGroups) > len(oldGroups) {
		t.Errorf("New extractor produced more groups (%d) than old extractor (%d)",
			len(newGroups), len(oldGroups))
	}
}

func TestEndpointExtractorV2_AllSpecs(t *testing.T) {
	adapter := NewLibOpenAPIAdapter()
	extractor := NewEndpointExtractorV2()
	specsDir := filepath.Join("..", "..", "..", "api-specs")

	specFiles := []string{
		"cisco_umbrella_networks_api_2_0_0.yaml",
		"cisco_umbrella_sites_api_2_0_0.yaml",
		"cisco_umbrella_users_and_roles_api_2_0_0.yaml",
		"cisco_umbrella_destination_lists_api_2_0_0.yaml",
		"cisco_umbrella_internal_networks_api_2_0_0.yaml",
	}

	for _, specFile := range specFiles {
		t.Run(specFile, func(t *testing.T) {
			specPath := filepath.Join(specsDir, specFile)

			spec, err := adapter.ParseFile(specPath)
			if err != nil {
				t.Fatalf("Failed to parse spec %s: %v", specFile, err)
			}

			groups, err := extractor.ExtractAndConsolidate(spec)
			if err != nil {
				t.Fatalf("Failed to extract from spec %s: %v", specFile, err)
			}

			report := extractor.GetConsolidationReport(groups)
			t.Logf("%s: %s", specFile, report.String())

			// Verify we have groups
			if len(groups) == 0 {
				t.Errorf("Expected %s to have groups", specFile)
			}

			// Verify no duplicate resource names would be generated
			resourceNames := make(map[string]bool)
			dataSourceNames := make(map[string]bool)

			for _, group := range groups {
				if group.HasResource {
					if resourceNames[group.CanonicalName] {
						t.Errorf("Duplicate resource name detected: %s", group.CanonicalName)
					}
					resourceNames[group.CanonicalName] = true
				}

				if group.HasDataSource {
					if dataSourceNames[group.CanonicalName] {
						t.Errorf("Duplicate data source name detected: %s", group.CanonicalName)
					}
					dataSourceNames[group.CanonicalName] = true
				}
			}

			t.Logf("%s: %d unique resource names, %d unique data source names",
				specFile, len(resourceNames), len(dataSourceNames))
		})
	}
}

func TestEndpointExtractorV2_PathExtraction(t *testing.T) {
	extractor := NewEndpointExtractorV2()

	testCases := []struct {
		path     string
		expected string
	}{
		{"/networks", "networks"},
		{"/networks/{id}", "networks"},
		{"/api/v2/networks", "networks"},
		{"/networks/{networkId}/policies", "policies"},
		{"/users/{userId}/roles", "roles"},
		{"/destination-lists", "destination-lists"},
		{"/internal_networks", "internal_networks"},
		{"/sites/{siteId}/settings", "settings"},
	}

	for _, tc := range testCases {
		t.Run(tc.path, func(t *testing.T) {
			result := extractor.extractNameFromPath(tc.path)
			if result != tc.expected {
				t.Errorf("extractNameFromPath(%q) = %q, expected %q", tc.path, result, tc.expected)
			}
		})
	}
}

func TestEndpointExtractorV2_CRUDMapping(t *testing.T) {
	adapter := NewLibOpenAPIAdapter()
	extractor := NewEndpointExtractorV2()

	specPath := filepath.Join("..", "..", "..", "api-specs", "cisco_umbrella_networks_api_2_0_0.yaml")
	spec, err := adapter.ParseFile(specPath)
	if err != nil {
		t.Fatalf("Failed to parse spec: %v", err)
	}

	groups, err := extractor.ExtractAndConsolidate(spec)
	if err != nil {
		t.Fatalf("Failed to extract: %v", err)
	}

	// Find the network group (should be the main one)
	var networkGroup *EndpointGroup
	for _, group := range groups {
		if group.CanonicalName == "network" {
			networkGroup = &group
			break
		}
	}

	if networkGroup == nil {
		t.Fatal("Expected to find 'network' group")
	}

	// Verify CRUD operations are properly mapped
	expectedCRUD := map[string]bool{
		"create": true,
		"read":   true,
		"update": true,
		"delete": true,
		"list":   true,
	}

	for crudType := range expectedCRUD {
		if _, exists := networkGroup.CRUDOps[crudType]; !exists {
			t.Errorf("Expected network group to have %s operation", crudType)
		}
	}

	t.Logf("Network group CRUD operations: %v", networkGroup.CRUDOps)
}
