package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents the complete generation configuration
type Config struct {
	Global      GlobalConfig      `yaml:"global"`
	Resources   ResourcesConfig   `yaml:"resources"`
	DataSources DataSourcesConfig `yaml:"data_sources"`
	Templates   TemplatesConfig   `yaml:"templates"`
	Parsing     ParsingConfig     `yaml:"parsing"`
	OAuth2      OAuth2Config      `yaml:"oauth2"`
	Output      OutputConfig      `yaml:"output"`
}

// GlobalConfig contains global generation settings
type GlobalConfig struct {
	ProviderName string `yaml:"provider_name"`
	PackageName  string `yaml:"package_name"`
	GoModule     string `yaml:"go_module"`
}

// ResourcesConfig contains resource generation settings
type ResourcesConfig struct {
	Defaults  ResourceDefaults            `yaml:"defaults"`
	Overrides map[string]ResourceDefaults `yaml:"overrides"`
}

// ResourceDefaults contains default settings for resources
type ResourceDefaults struct {
	GenerateCRUD   bool `yaml:"generate_crud"`
	GenerateImport bool `yaml:"generate_import"`
	GenerateDocs   bool `yaml:"generate_docs"`
}

// DataSourcesConfig contains data source generation settings
type DataSourcesConfig struct {
	Defaults DataSourceDefaults `yaml:"defaults"`
}

// DataSourceDefaults contains default settings for data sources
type DataSourceDefaults struct {
	GenerateRead bool `yaml:"generate_read"`
	GenerateDocs bool `yaml:"generate_docs"`
}

// TemplatesConfig contains template file paths
type TemplatesConfig struct {
	ResourceTemplate     string `yaml:"resource_template"`
	DataSourceTemplate   string `yaml:"data_source_template"`
	ClientMethodTemplate string `yaml:"client_method_template"`
}

// ParsingConfig contains API parsing settings
type ParsingConfig struct {
	SkipEndpoints []string     `yaml:"skip_endpoints"`
	Naming        NamingConfig `yaml:"naming"`
}

// NamingConfig contains naming convention settings
type NamingConfig struct {
	OperationIDToResource bool     `yaml:"operation_id_to_resource"`
	StripPrefixes         []string `yaml:"strip_prefixes"`
	StripSuffixes         []string `yaml:"strip_suffixes"`
}

// OAuth2Config contains OAuth2 scope mapping settings
type OAuth2Config struct {
	ScopeMappings map[string][]string `yaml:"scope_mappings"`
}

// OutputConfig contains output formatting settings
type OutputConfig struct {
	ResourceFilePattern   string `yaml:"resource_file_pattern"`
	DataSourceFilePattern string `yaml:"data_source_file_pattern"`
	FormatCode            bool   `yaml:"format_code"`
	AddLicenseHeader      bool   `yaml:"add_license_header"`
	AddGenerationMarker   bool   `yaml:"add_generation_marker"`
}

// LoadConfig loads configuration from a YAML file
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", path, err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file %s: %w", path, err)
	}

	// Set defaults if not specified
	if config.Global.ProviderName == "" {
		config.Global.ProviderName = "umbrella"
	}
	if config.Global.PackageName == "" {
		config.Global.PackageName = "provider"
	}
	if config.Output.ResourceFilePattern == "" {
		config.Output.ResourceFilePattern = "resource_%s.go"
	}
	if config.Output.DataSourceFilePattern == "" {
		config.Output.DataSourceFilePattern = "data_source_%s.go"
	}

	return &config, nil
}

// GetResourceConfig returns the configuration for a specific resource
func (c *Config) GetResourceConfig(resourceName string) ResourceDefaults {
	if override, exists := c.Resources.Overrides[resourceName]; exists {
		return override
	}
	return c.Resources.Defaults
}
