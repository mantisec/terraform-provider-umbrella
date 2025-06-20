package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

func TestGetorganizationinfoDataSource_Schema(t *testing.T) {

	r := NewGetorganizationinfoResource()

	// Test that the resource implements the correct interface

	var _ datasource.DataSource = r

}
