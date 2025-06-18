package openapi

import (
	"path/filepath"
	"testing"
)

func TestLoader_LoadFromFile(t *testing.T) {
	loader := NewLoader()

	// Test with one of our API specs
	specPath := filepath.Join("..", "..", "..", "api-specs", "cisco_umbrella_networks_api_2_0_0.yaml")

	doc, err := loader.LoadFromFile(specPath)
	if err != nil {
		t.Fatalf("Failed to load spec: %v", err)
	}

	// Validate the document
	if err := loader.ValidateDocument(doc); err != nil {
		t.Fatalf("Document validation failed: %v", err)
	}

	// Get document info
	info := loader.GetDocumentInfo(doc)
	t.Logf("Loaded document: %s", info.String())

	// Basic assertions
	if info.Title == "" {
		t.Error("Expected document to have a title")
	}

	if info.Version == "" {
		t.Error("Expected document to have a version")
	}

	if info.PathCount == 0 {
		t.Error("Expected document to have paths")
	}

	if info.OperationCount == 0 {
		t.Error("Expected document to have operations")
	}

	t.Logf("Document has %d paths, %d operations, %d schemas",
		info.PathCount, info.OperationCount, info.SchemaCount)
}

func TestLoader_LoadAllSpecs(t *testing.T) {
	loader := NewLoader()
	specsDir := filepath.Join("..", "..", "..", "api-specs")

	// List of spec files to test
	specFiles := []string{
		"cisco_umbrella_networks_api_2_0_0.yaml",
		"cisco_umbrella_sites_api_2_0_0.yaml",
		"cisco_umbrella_users_and_roles_api_2_0_0.yaml",
		"cisco_umbrella_internal_networks_api_2_0_0.yaml",
		"cisco_umbrella_destination_lists_api_2_0_0.yaml",
	}

	for _, specFile := range specFiles {
		t.Run(specFile, func(t *testing.T) {
			specPath := filepath.Join(specsDir, specFile)

			doc, err := loader.LoadFromFile(specPath)
			if err != nil {
				t.Fatalf("Failed to load spec %s: %v", specFile, err)
			}

			// Validate the document
			if err := loader.ValidateDocument(doc); err != nil {
				t.Fatalf("Document validation failed for %s: %v", specFile, err)
			}

			// Get document info
			info := loader.GetDocumentInfo(doc)
			t.Logf("%s: %s", specFile, info.String())

			// Basic assertions
			if info.PathCount == 0 {
				t.Errorf("Expected %s to have paths", specFile)
			}

			if info.OperationCount == 0 {
				t.Errorf("Expected %s to have operations", specFile)
			}
		})
	}
}

func TestLoader_InvalidSpec(t *testing.T) {
	loader := NewLoader()

	// Test with invalid YAML
	_, err := loader.LoadFromBytes([]byte("invalid: yaml: content"))
	if err == nil {
		t.Error("Expected error for invalid YAML")
	}

	// Test with empty data
	_, err = loader.LoadFromBytes([]byte(""))
	if err == nil {
		t.Error("Expected error for empty data")
	}
}

func TestDocumentInfo_String(t *testing.T) {
	info := DocumentInfo{
		Title:          "Test API",
		Version:        "1.0.0",
		PathCount:      5,
		OperationCount: 15,
		SchemaCount:    10,
	}

	expected := "OpenAPI Test API v1.0.0: 5 paths, 15 operations, 10 schemas"
	if info.String() != expected {
		t.Errorf("Expected %q, got %q", expected, info.String())
	}
}
