package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUpdatenetworkResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatenetworkResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatenetwork.test", "name", "test-updatenetwork"),
				),
			},

			{
				Config: testAccUpdatenetworkResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatenetwork.test", "name", "test-updatenetwork-updated"),
				),
			},
		},
	})
}

func TestUpdatenetworkResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatenetworkResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatenetwork.test", "name", "test-updatenetwork"),
				),
			},

			{
				Config: testAccUpdatenetworkResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatenetwork.test", "name", "test-updatenetwork-updated"),
				),
			},
		},
	})
}

func TestUpdatenetworkResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatenetworkResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatenetwork.test", "name", "test-updatenetwork"),
				),
			},

			{
				Config: testAccUpdatenetworkResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatenetwork.test", "name", "test-updatenetwork-updated"),
				),
			},
		},
	})
}

func testAccUpdatenetworkResourceConfig_basic() string {
	return `

resource "umbrella_updatenetwork" "test" {
  name = "test-updatenetwork"
  
  name = "test-resource"
  
  ipAddress = "test-value"
  
  prefixLength = 123
  
  isDynamic = true
  
  status = "test-value"
  
}

`
}

func testAccUpdatenetworkResourceConfig_update() string {
	return `
resource "umbrella_updatenetwork" "test" {
  name = "test-updatenetwork-updated"
  
  name = "test-resource-updated"
  
  ipAddress = "updated-value"
  
  prefixLength = 456
  
  isDynamic = false
  
  status = "updated-value"
  
}
`
}
