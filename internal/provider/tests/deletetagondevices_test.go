package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

func TestDeletetagondevicesResource_Schema(t *testing.T) {

	r := NewDeletetagondevicesResource()

	// Test that the resource implements the correct interface

	var _ resource.Resource = r

}
