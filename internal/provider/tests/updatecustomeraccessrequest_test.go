package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

func TestUpdatecustomeraccessrequestResource_Schema(t *testing.T) {

	r := NewUpdatecustomeraccessrequestResource()

	// Test that the resource implements the correct interface

	var _ resource.Resource = r

}
