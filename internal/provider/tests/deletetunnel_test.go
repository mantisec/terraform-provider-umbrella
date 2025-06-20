package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

func TestDeletetunnelResource_Schema(t *testing.T) {

	r := NewDeletetunnelResource()

	// Test that the resource implements the correct interface

	var _ resource.Resource = r

}
