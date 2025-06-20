package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreatemanagedprovidercustomerResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatemanagedprovidercustomerResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createmanagedprovidercustomer.test", "name", "test-createmanagedprovidercustomer"),
				),
			},

			{
				Config: testAccCreatemanagedprovidercustomerResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createmanagedprovidercustomer.test", "name", "test-createmanagedprovidercustomer-updated"),
				),
			},
		},
	})
}

func TestCreatemanagedprovidercustomerResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatemanagedprovidercustomerResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createmanagedprovidercustomer.test", "name", "test-createmanagedprovidercustomer"),
				),
			},

			{
				Config: testAccCreatemanagedprovidercustomerResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createmanagedprovidercustomer.test", "name", "test-createmanagedprovidercustomer-updated"),
				),
			},
		},
	})
}

func TestCreatemanagedprovidercustomerResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatemanagedprovidercustomerResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createmanagedprovidercustomer.test", "name", "test-createmanagedprovidercustomer"),
				),
			},

			{
				Config: testAccCreatemanagedprovidercustomerResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createmanagedprovidercustomer.test", "name", "test-createmanagedprovidercustomer-updated"),
				),
			},
		},
	})
}

func testAccCreatemanagedprovidercustomerResourceConfig_basic() string {
	return `

resource "umbrella_createmanagedprovidercustomer" "test" {
  name = "test-createmanagedprovidercustomer"
  
  name = "test-resource"
  
  customerName = "test-value"
  
  seats = 123
  
}

`
}

func testAccCreatemanagedprovidercustomerResourceConfig_update() string {
	return `
resource "umbrella_createmanagedprovidercustomer" "test" {
  name = "test-createmanagedprovidercustomer-updated"
  
  name = "test-resource-updated"
  
  customerName = "updated-value"
  
  seats = 456
  
}
`
}
