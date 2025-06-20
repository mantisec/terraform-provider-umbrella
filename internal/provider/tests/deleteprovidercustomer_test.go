package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

func TestDeleteprovidercustomerResource_Schema(t *testing.T) {

	r := NewDeleteprovidercustomerResource()

	// Test that the resource implements the correct interface

	var _ resource.Resource = r

}
