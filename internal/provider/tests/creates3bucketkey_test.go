package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

func TestCreates3bucketkeyResource_Schema(t *testing.T) {

	r := NewCreates3bucketkeyResource()

	// Test that the resource implements the correct interface

	var _ resource.Resource = r

}
