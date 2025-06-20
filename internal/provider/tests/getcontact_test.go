package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

func TestGetcontactDataSource_Schema(t *testing.T) {

	r := NewGetcontactResource()

	// Test that the resource implements the correct interface

	var _ datasource.DataSource = r

}
