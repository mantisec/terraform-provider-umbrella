package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

func TestCreatenetworkResource_Schema(t *testing.T) {

	r := NewCreatenetworkResource()

	// Test that the resource implements the correct interface

	var _ resource.Resource = r

}
