package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

func TestGetlogoDataSource_Schema(t *testing.T) {

	r := NewGetlogoResource()

	// Test that the resource implements the correct interface

	var _ datasource.DataSource = r

}
