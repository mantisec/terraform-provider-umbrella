package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/mantisec/terraform-provider-umbrella/tools/generator/config"
	"github.com/mantisec/terraform-provider-umbrella/tools/generator/generator"
	"github.com/mantisec/terraform-provider-umbrella/tools/generator/parser"
)

func main() {
	var (
		configPath = flag.String("config", "tools/generator/config/generation.yaml", "Path to generation configuration file")
		specsDir   = flag.String("specs", "api-specs", "Directory containing OpenAPI specification files")
		outputDir  = flag.String("output", "internal/provider", "Output directory for generated files")
		verbose    = flag.Bool("verbose", false, "Enable verbose logging")
	)
	flag.Parse()

	if *verbose {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	// Load configuration
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize parser
	apiParser := parser.NewOpenAPIParser()

	// Initialize generator
	gen := generator.NewGenerator(cfg)

	// Find OpenAPI spec files
	specFiles, err := findSpecFiles(*specsDir)
	if err != nil {
		log.Fatalf("Failed to find spec files: %v", err)
	}

	if len(specFiles) == 0 {
		log.Fatalf("No OpenAPI specification files found in %s", *specsDir)
	}

	log.Printf("Found %d OpenAPI specification files", len(specFiles))

	// Process each spec file
	for _, specFile := range specFiles {
		log.Printf("Processing %s", specFile)

		// Parse OpenAPI spec
		apiSpec, err := apiParser.ParseFile(specFile)
		if err != nil {
			log.Printf("Warning: Failed to parse %s: %v", specFile, err)
			continue
		}

		// Generate code
		if err := gen.GenerateFromSpec(apiSpec, *outputDir); err != nil {
			log.Printf("Warning: Failed to generate code for %s: %v", specFile, err)
			continue
		}

		log.Printf("Successfully generated code for %s", specFile)
	}

	log.Println("Code generation completed")
}

// findSpecFiles finds all YAML files that appear to be OpenAPI specs
func findSpecFiles(dir string) ([]string, error) {
	var specFiles []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and non-YAML files
		if info.IsDir() {
			return nil
		}

		yamlMatch, _ := filepath.Match("*.yaml", info.Name())
		ymlMatch, _ := filepath.Match("*.yml", info.Name())
		if !yamlMatch && !ymlMatch {
			return nil
		}

		// Skip files in tools directory to avoid recursion
		if filepath.HasPrefix(path, "tools/") {
			return nil
		}

		// Basic check if file contains OpenAPI content
		if isOpenAPIFile(path) {
			specFiles = append(specFiles, path)
		}

		return nil
	})

	return specFiles, err
}

// isOpenAPIFile performs a basic check to see if a file is an OpenAPI spec
func isOpenAPIFile(path string) bool {
	content, err := os.ReadFile(path)
	if err != nil {
		return false
	}

	contentStr := string(content)
	if len(contentStr) < 8 {
		return false
	}

	// Check for OpenAPI version indicators
	return contentStr[:8] == "openapi:" ||
		(len(contentStr) > 10 && contentStr[:10] == "openapi: 3")
}
