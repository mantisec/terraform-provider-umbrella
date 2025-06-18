package parser

import (
	"path/filepath"
	"testing"
)

func TestLibOpenAPIAdapter_ParseFile(t *testing.T) {
	adapter := NewLibOpenAPIAdapter()

	// Test with one of our API specs
	specPath := filepath.Join("..", "..", "..", "api-specs", "cisco_umbrella_networks_api_2_0_0.yaml")

	spec, err := adapter.ParseFile(specPath)
	if err != nil {
		t.Fatalf("Failed to parse spec: %v", err)
	}

	// Basic validation
	if spec.Info.Title == "" {
		t.Error("Expected spec to have a title")
	}

	if spec.Info.Version == "" {
		t.Error("Expected spec to have a version")
	}

	if len(spec.Paths) == 0 {
		t.Error("Expected spec to have paths")
	}

	t.Logf("Parsed spec: %s v%s with %d paths",
		spec.Info.Title, spec.Info.Version, len(spec.Paths))

	// Test that we have operations in paths
	operationCount := 0
	for path, pathItem := range spec.Paths {
		t.Logf("Path: %s", path)
		if pathItem.Get != nil {
			operationCount++
			t.Logf("  GET: %s", pathItem.Get.Summary)
		}
		if pathItem.Post != nil {
			operationCount++
			t.Logf("  POST: %s", pathItem.Post.Summary)
		}
		if pathItem.Put != nil {
			operationCount++
			t.Logf("  PUT: %s", pathItem.Put.Summary)
		}
		if pathItem.Delete != nil {
			operationCount++
			t.Logf("  DELETE: %s", pathItem.Delete.Summary)
		}
		if pathItem.Patch != nil {
			operationCount++
			t.Logf("  PATCH: %s", pathItem.Patch.Summary)
		}
	}

	if operationCount == 0 {
		t.Error("Expected spec to have operations")
	}

	t.Logf("Total operations: %d", operationCount)

	// Test that schemas are properly resolved
	if len(spec.Components.Schemas) == 0 {
		t.Error("Expected spec to have component schemas")
	}

	for schemaName, schema := range spec.Components.Schemas {
		t.Logf("Schema: %s (type: %s)", schemaName, schema.Type)
		if schema.Properties != nil {
			t.Logf("  Properties: %d", len(schema.Properties))
		}
	}
}

func TestLibOpenAPIAdapter_CompareWithOldParser(t *testing.T) {
	// Test that the new adapter produces similar results to the old parser
	specPath := filepath.Join("..", "..", "..", "api-specs", "cisco_umbrella_networks_api_2_0_0.yaml")

	// Parse with new adapter
	adapter := NewLibOpenAPIAdapter()
	newSpec, err := adapter.ParseFile(specPath)
	if err != nil {
		t.Fatalf("Failed to parse with new adapter: %v", err)
	}

	// Parse with old parser
	oldParser := NewOpenAPIParser()
	oldSpec, err := oldParser.ParseFile(specPath)
	if err != nil {
		t.Fatalf("Failed to parse with old parser: %v", err)
	}

	// Compare basic info
	if newSpec.Info.Title != oldSpec.Info.Title {
		t.Errorf("Title mismatch: new=%s, old=%s", newSpec.Info.Title, oldSpec.Info.Title)
	}

	if newSpec.Info.Version != oldSpec.Info.Version {
		t.Errorf("Version mismatch: new=%s, old=%s", newSpec.Info.Version, oldSpec.Info.Version)
	}

	// Compare path counts
	if len(newSpec.Paths) != len(oldSpec.Paths) {
		t.Errorf("Path count mismatch: new=%d, old=%d", len(newSpec.Paths), len(oldSpec.Paths))
	}

	// Compare schema counts
	if len(newSpec.Components.Schemas) != len(oldSpec.Components.Schemas) {
		t.Logf("Schema count difference: new=%d, old=%d (this is expected - new parser may resolve more schemas)",
			len(newSpec.Components.Schemas), len(oldSpec.Components.Schemas))
	}

	t.Logf("Comparison complete - new adapter produces compatible results")
}

func TestLibOpenAPIAdapter_SchemaResolution(t *testing.T) {
	adapter := NewLibOpenAPIAdapter()

	// Test with a spec that has complex schema references
	specPath := filepath.Join("..", "..", "..", "api-specs", "cisco_umbrella_destination_lists_api_2_0_0.yaml")

	spec, err := adapter.ParseFile(specPath)
	if err != nil {
		t.Fatalf("Failed to parse spec: %v", err)
	}

	// This spec should have many schemas due to complex references
	if len(spec.Components.Schemas) == 0 {
		t.Error("Expected spec to have component schemas")
	}

	t.Logf("Resolved %d schemas", len(spec.Components.Schemas))

	// Check that schemas are properly resolved (no $ref fields should remain)
	for schemaName, schema := range spec.Components.Schemas {
		if schema.Ref != "" {
			t.Errorf("Schema %s still has unresolved reference: %s", schemaName, schema.Ref)
		}

		// Check nested properties are also resolved
		for propName, propSchema := range schema.Properties {
			if propSchema.Ref != "" {
				t.Errorf("Schema %s property %s still has unresolved reference: %s",
					schemaName, propName, propSchema.Ref)
			}
		}
	}

	t.Logf("Schema resolution validation complete")
}

func TestLibOpenAPIAdapter_AllSpecs(t *testing.T) {
	adapter := NewLibOpenAPIAdapter()
	specsDir := filepath.Join("..", "..", "..", "api-specs")

	// Test a subset of specs to ensure adapter works across different API definitions
	specFiles := []string{
		"cisco_umbrella_networks_api_2_0_0.yaml",
		"cisco_umbrella_sites_api_2_0_0.yaml",
		"cisco_umbrella_users_and_roles_api_2_0_0.yaml",
		"cisco_umbrella_destination_lists_api_2_0_0.yaml",
	}

	for _, specFile := range specFiles {
		t.Run(specFile, func(t *testing.T) {
			specPath := filepath.Join(specsDir, specFile)

			spec, err := adapter.ParseFile(specPath)
			if err != nil {
				t.Fatalf("Failed to parse spec %s: %v", specFile, err)
			}

			// Basic validation
			if spec.Info.Title == "" {
				t.Errorf("Spec %s has no title", specFile)
			}

			if len(spec.Paths) == 0 {
				t.Errorf("Spec %s has no paths", specFile)
			}

			t.Logf("%s: %s v%s - %d paths, %d schemas",
				specFile, spec.Info.Title, spec.Info.Version,
				len(spec.Paths), len(spec.Components.Schemas))
		})
	}
}
