package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

func TestGettagsDataSource_Schema(t *testing.T) {

	r := NewGettagsResource()

	// Test that the resource implements the correct interface

	var _ datasource.DataSource = r

}
