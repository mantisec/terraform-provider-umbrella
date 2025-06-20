package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

func TestUpdatenetworkdeviceResource_Schema(t *testing.T) {

	r := NewUpdatenetworkdeviceResource()

	// Test that the resource implements the correct interface

	var _ resource.Resource = r

}
