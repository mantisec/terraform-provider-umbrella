package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

func TestGetallnetworkdevicesDataSource_Schema(t *testing.T) {

	r := NewGetallnetworkdevicesResource()

	// Test that the resource implements the correct interface

	var _ datasource.DataSource = r

}
