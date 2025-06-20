package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeletetunnelResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletetunnelResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletetunnel.test", "name", "test-deletetunnel"),
				),
			},

			{
				Config: testAccDeletetunnelResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletetunnel.test", "name", "test-deletetunnel-updated"),
				),
			},
		},
	})
}

func TestDeletetunnelResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletetunnelResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletetunnel.test", "name", "test-deletetunnel"),
				),
			},

			{
				Config: testAccDeletetunnelResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletetunnel.test", "name", "test-deletetunnel-updated"),
				),
			},
		},
	})
}

func TestDeletetunnelResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletetunnelResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletetunnel.test", "name", "test-deletetunnel"),
				),
			},

			{
				Config: testAccDeletetunnelResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletetunnel.test", "name", "test-deletetunnel-updated"),
				),
			},
		},
	})
}

func testAccDeletetunnelResourceConfig_basic() string {
	return `

resource "umbrella_deletetunnel" "test" {
  name = "test-deletetunnel"
  
  name = "test-resource"
  
  detachPolicies = true
  
}

`
}

func testAccDeletetunnelResourceConfig_update() string {
	return `
resource "umbrella_deletetunnel" "test" {
  name = "test-deletetunnel-updated"
  
  name = "test-resource-updated"
  
  detachPolicies = false
  
}
`
}
