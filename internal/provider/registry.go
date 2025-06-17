package provider

import (
	"sync"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GeneratedResourceRegistry holds all generated resources and data sources
type GeneratedResourceRegistry struct {
	mu          sync.RWMutex
	resources   []func() resource.Resource
	dataSources []func() datasource.DataSource
	initialized bool
}

// Global registry instance
var generatedRegistry *GeneratedResourceRegistry
var registryOnce sync.Once

// NewGeneratedResourceRegistry creates a new registry for generated resources
func NewGeneratedResourceRegistry() *GeneratedResourceRegistry {
	return &GeneratedResourceRegistry{
		resources:   make([]func() resource.Resource, 0),
		dataSources: make([]func() datasource.DataSource, 0),
		initialized: false,
	}
}

// GetGeneratedRegistry returns the global generated resource registry
// This function is called by the provider to get access to generated resources
func GetGeneratedRegistry() *GeneratedResourceRegistry {
	registryOnce.Do(func() {
		generatedRegistry = NewGeneratedResourceRegistry()
	})
	return generatedRegistry
}

// RegisterResource adds a resource constructor to the registry
// This function is called by generated resource files during initialization
func (r *GeneratedResourceRegistry) RegisterResource(resourceFunc func() resource.Resource) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.resources = append(r.resources, resourceFunc)
}

// RegisterDataSource adds a data source constructor to the registry
// This function is called by generated data source files during initialization
func (r *GeneratedResourceRegistry) RegisterDataSource(dataSourceFunc func() datasource.DataSource) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.dataSources = append(r.dataSources, dataSourceFunc)
}

// GetResources returns all registered resource constructors
// This function is called by the provider to get all available resources
func (r *GeneratedResourceRegistry) GetResources() []func() resource.Resource {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Return a copy to prevent external modification
	result := make([]func() resource.Resource, len(r.resources))
	copy(result, r.resources)
	return result
}

// GetDataSources returns all registered data source constructors
// This function is called by the provider to get all available data sources
func (r *GeneratedResourceRegistry) GetDataSources() []func() datasource.DataSource {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Return a copy to prevent external modification
	result := make([]func() datasource.DataSource, len(r.dataSources))
	copy(result, r.dataSources)
	return result
}

// IsInitialized returns whether the registry has been initialized
func (r *GeneratedResourceRegistry) IsInitialized() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.initialized
}

// SetInitialized marks the registry as initialized
func (r *GeneratedResourceRegistry) SetInitialized() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.initialized = true
}

// Reset clears all registered resources and data sources
// This function is primarily used for testing
func (r *GeneratedResourceRegistry) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.resources = make([]func() resource.Resource, 0)
	r.dataSources = make([]func() datasource.DataSource, 0)
	r.initialized = false
}

// GetResourceCount returns the number of registered resources
func (r *GeneratedResourceRegistry) GetResourceCount() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.resources)
}

// GetDataSourceCount returns the number of registered data sources
func (r *GeneratedResourceRegistry) GetDataSourceCount() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.dataSources)
}
