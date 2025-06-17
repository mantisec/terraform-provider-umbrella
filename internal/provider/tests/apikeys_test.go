package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

func TestApikeysResource_Schema(t *testing.T) {

	r := NewApikeysResource()

	// Test that the resource implements the correct interface

	var _ resource.Resource = r

}
