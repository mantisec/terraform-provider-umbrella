package generator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mantisec/terraform-provider-umbrella/tools/generator/config"
	"github.com/mantisec/terraform-provider-umbrella/tools/generator/parser"
)

func TestEndToEndGeneration(t *testing.T) {
	// Create a temporary output directory
	tempDir, err := os.MkdirTemp("", "generator_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Load configuration
	cfg := &config.Config{
		Global: config.GlobalConfig{
			ProviderName: "umbrella",
			PackageName:  "provider",
		},
		Output: config.OutputConfig{
			ResourceFilePattern:   "resource_%s.go",
			DataSourceFilePattern: "data_source_%s.go",
		},
		Parsing: config.ParsingConfig{
			SkipEndpoints: []string{},
		},
	}

	// Test with the new libopenapi adapter
	adapter := parser.NewLibOpenAPIAdapter()
	specPath := filepath.Join("..", "..", "..", "api-specs", "cisco_umbrella_networks_api_2_0_0.yaml")

	spec, err := adapter.ParseFile(specPath)
	if err != nil {
		t.Fatalf("Failed to parse spec with new adapter: %v", err)
	}

	// Test with the new endpoint extractor
	extractorV2 := parser.NewEndpointExtractorV2()
	groups, err := extractorV2.ExtractAndConsolidate(spec)
	if err != nil {
		t.Fatalf("Failed to extract and consolidate endpoints: %v", err)
	}

	// Generate consolidation report
	report := extractorV2.GetConsolidationReport(groups)
	t.Logf("Consolidation report: %s", report.String())

	// Verify consolidation worked
	if len(groups) == 0 {
		t.Error("Expected consolidated groups, got none")
	}

	// Test the generator with new components
	gen := NewGenerator(cfg)

	// Use the main generation method
	if err := gen.GenerateFromSpec(spec, tempDir); err != nil {
		t.Errorf("Failed to generate from spec: %v", err)
	}

	// Verify files were created
	files, err := os.ReadDir(tempDir)
	if err != nil {
		t.Fatalf("Failed to read temp dir: %v", err)
	}

	if len(files) == 0 {
		t.Error("No files were generated")
	}

	// Log generated files
	for _, file := range files {
		t.Logf("Generated file: %s", file.Name())

		// Check that files are not empty
		filePath := filepath.Join(tempDir, file.Name())
		info, err := os.Stat(filePath)
		if err != nil {
			t.Errorf("Failed to stat file %s: %v", file.Name(), err)
			continue
		}

		if info.Size() == 0 {
			t.Errorf("Generated file %s is empty", file.Name())
		}
	}

	// Verify that TODO comments are not present in generated files
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".go" {
			content, err := os.ReadFile(filepath.Join(tempDir, file.Name()))
			if err != nil {
				t.Errorf("Failed to read file %s: %v", file.Name(), err)
				continue
			}

			contentStr := string(content)
			if contains(contentStr, "TODO: Implement") {
				t.Errorf("File %s still contains TODO comments", file.Name())
			}
		}
	}

	t.Logf("End-to-end test completed successfully. Generated %d files in %s", len(files), tempDir)
}

func TestNewVsOldExtractor(t *testing.T) {
	// Load a spec
	adapter := parser.NewLibOpenAPIAdapter()
	specPath := filepath.Join("..", "..", "..", "api-specs", "cisco_umbrella_networks_api_2_0_0.yaml")

	spec, err := adapter.ParseFile(specPath)
	if err != nil {
		t.Fatalf("Failed to parse spec: %v", err)
	}

	// Test old extractor
	oldExtractor := parser.NewEndpointExtractor()
	oldEndpoints, err := oldExtractor.ExtractEndpoints(spec)
	if err != nil {
		t.Fatalf("Failed to extract with old extractor: %v", err)
	}
	oldGroups := oldExtractor.GroupEndpointsByResource(oldEndpoints)

	// Test new extractor
	newExtractor := parser.NewEndpointExtractorV2()
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

	// New extractor should produce fewer or equal groups
	if len(newGroups) > len(oldGroups) {
		t.Errorf("New extractor produced more groups (%d) than old extractor (%d)",
			len(newGroups), len(oldGroups))
	}

	// Verify no duplicate resource names
	resourceNames := make(map[string]bool)
	for _, group := range newGroups {
		if group.HasResource {
			if resourceNames[group.CanonicalName] {
				t.Errorf("Duplicate resource name: %s", group.CanonicalName)
			}
			resourceNames[group.CanonicalName] = true
		}
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) &&
		(s[:len(substr)] == substr || s[len(s)-len(substr):] == substr ||
			containsAt(s, substr, 1)))
}

func containsAt(s, substr string, start int) bool {
	if start >= len(s) {
		return false
	}
	if start+len(substr) > len(s) {
		return containsAt(s, substr, start+1)
	}
	if s[start:start+len(substr)] == substr {
		return true
	}
	return containsAt(s, substr, start+1)
}
