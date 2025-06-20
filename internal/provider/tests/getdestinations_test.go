package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

func TestGetdestinationsDataSource_Schema(t *testing.T) {

	r := NewGetdestinationsResource()

	// Test that the resource implements the correct interface

	var _ datasource.DataSource = r

}
