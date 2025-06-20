package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUpdatecustomerdealsResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecustomerdealsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecustomerdeals.test", "name", "test-updatecustomerdeals"),
				),
			},

			{
				Config: testAccUpdatecustomerdealsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatecustomerdeals.test", "name", "test-updatecustomerdeals-updated"),
				),
			},
		},
	})
}

func TestUpdatecustomerdealsResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecustomerdealsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecustomerdeals.test", "name", "test-updatecustomerdeals"),
				),
			},

			{
				Config: testAccUpdatecustomerdealsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatecustomerdeals.test", "name", "test-updatecustomerdeals-updated"),
				),
			},
		},
	})
}

func TestUpdatecustomerdealsResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecustomerdealsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecustomerdeals.test", "name", "test-updatecustomerdeals"),
				),
			},

			{
				Config: testAccUpdatecustomerdealsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatecustomerdeals.test", "name", "test-updatecustomerdeals-updated"),
				),
			},
		},
	})
}

func testAccUpdatecustomerdealsResourceConfig_basic() string {
	return `

resource "umbrella_updatecustomerdeals" "test" {
  name = "test-updatecustomerdeals"
  
  name = "test-resource"
  
  ccoid = 123
  
  customerId = 123
  
  quoteId = 123
  
  majorLineItems = ["item1", "item2"]
  
}

`
}

func testAccUpdatecustomerdealsResourceConfig_update() string {
	return `
resource "umbrella_updatecustomerdeals" "test" {
  name = "test-updatecustomerdeals-updated"
  
  name = "test-resource-updated"
  
  ccoid = 456
  
  customerId = 456
  
  quoteId = 456
  
  majorLineItems = ["updated1", "updated2"]
  
}
`
}
