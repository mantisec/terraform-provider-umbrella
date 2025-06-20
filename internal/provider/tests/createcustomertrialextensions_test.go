package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

func TestCreatecustomertrialextensionsResource_Schema(t *testing.T) {

	r := NewCreatecustomertrialextensionsResource()

	// Test that the resource implements the correct interface

	var _ resource.Resource = r

}
