package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

func TestSummaryDataSource_Schema(t *testing.T) {

	r := NewSummaryResource()

	// Test that the resource implements the correct interface

	var _ datasource.DataSource = r

}
