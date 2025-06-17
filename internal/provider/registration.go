package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// This file provides registration functions for generated resources and data sources
// Generated files will import this package and call these functions during init()

// RegisterGeneratedResource registers a generated resource with the global registry
func RegisterGeneratedResource(resourceFunc func() resource.Resource) {
	registry := GetGeneratedRegistry()
	registry.RegisterResource(resourceFunc)
}

// RegisterGeneratedDataSource registers a generated data source with the global registry
func RegisterGeneratedDataSource(dataSourceFunc func() datasource.DataSource) {
	registry := GetGeneratedRegistry()
	registry.RegisterDataSource(dataSourceFunc)
}
