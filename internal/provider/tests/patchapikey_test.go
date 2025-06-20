package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

func TestPatchapikeyResource_Schema(t *testing.T) {

	r := NewPatchapikeyResource()

	// Test that the resource implements the correct interface

	var _ resource.Resource = r

}
