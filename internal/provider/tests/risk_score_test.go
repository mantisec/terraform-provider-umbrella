package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

func TestRiskScoreDataSource_Schema(t *testing.T) {

	r := NewRiskScoreResource()

	// Test that the resource implements the correct interface

	var _ datasource.DataSource = r

}
