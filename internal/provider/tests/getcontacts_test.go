package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

func TestGetcontactsDataSource_Schema(t *testing.T) {

	r := NewGetcontactsResource()

	// Test that the resource implements the correct interface

	var _ datasource.DataSource = r

}
