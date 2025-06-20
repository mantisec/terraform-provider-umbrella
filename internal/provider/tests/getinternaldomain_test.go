package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

func TestGetinternaldomainDataSource_Schema(t *testing.T) {

	r := NewGetinternaldomainResource()

	// Test that the resource implements the correct interface

	var _ datasource.DataSource = r

}
