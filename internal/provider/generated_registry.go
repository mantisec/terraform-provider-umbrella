package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GeneratedResourceRegistry manages registration of generated resources and data sources
type GeneratedResourceRegistry struct {
	resources   []func() resource.Resource
	dataSources []func() datasource.DataSource
}

// NewGeneratedResourceRegistry creates a new registry for generated resources
func NewGeneratedResourceRegistry() *GeneratedResourceRegistry {
	return &GeneratedResourceRegistry{
		resources:   make([]func() resource.Resource, 0),
		dataSources: make([]func() datasource.DataSource, 0),
	}
}

// RegisterResource registers a generated resource constructor
func (r *GeneratedResourceRegistry) RegisterResource(constructor func() resource.Resource) {
	r.resources = append(r.resources, constructor)
}

// RegisterDataSource registers a generated data source constructor
func (r *GeneratedResourceRegistry) RegisterDataSource(constructor func() datasource.DataSource) {
	r.dataSources = append(r.dataSources, constructor)
}

// GetResources returns all registered resource constructors
func (r *GeneratedResourceRegistry) GetResources() []func() resource.Resource {
	return r.resources
}

// GetDataSources returns all registered data source constructors
func (r *GeneratedResourceRegistry) GetDataSources() []func() datasource.DataSource {
	return r.dataSources
}

// RegisterGeneratedResources registers all generated resources and data sources
// This function will be called by the code generator to register generated resources
func RegisterGeneratedResources(registry *GeneratedResourceRegistry) {
	// Register generated resources
	registry.RegisterResource(NewGeneratedDestinationListResource)

	// Additional generated resources will be registered here by the code generator
	// registry.RegisterResource(NewGeneratedNetworkResource)
	// registry.RegisterResource(NewGeneratedSiteResource)

	// Register generated data sources
	// registry.RegisterDataSource(NewGeneratedDestinationListDataSource)
	// registry.RegisterDataSource(NewGeneratedNetworkDataSource)
}

// Global registry instance
var generatedRegistry = NewGeneratedResourceRegistry()

// GetGeneratedRegistry returns the global generated resource registry
func GetGeneratedRegistry() *GeneratedResourceRegistry {
	return generatedRegistry
}

// InitializeGeneratedResources initializes all generated resources
func InitializeGeneratedResources() {
	RegisterGeneratedResources(generatedRegistry)
}
