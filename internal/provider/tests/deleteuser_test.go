package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

func TestDeleteuserResource_Schema(t *testing.T) {

	r := NewDeleteuserResource()

	// Test that the resource implements the correct interface

	var _ resource.Resource = r

}
