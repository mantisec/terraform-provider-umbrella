package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUpdatemanagedprovidercustomerResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatemanagedprovidercustomerResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatemanagedprovidercustomer.test", "name", "test-updatemanagedprovidercustomer"),
				),
			},

			{
				Config: testAccUpdatemanagedprovidercustomerResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatemanagedprovidercustomer.test", "name", "test-updatemanagedprovidercustomer-updated"),
				),
			},
		},
	})
}

func TestUpdatemanagedprovidercustomerResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatemanagedprovidercustomerResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatemanagedprovidercustomer.test", "name", "test-updatemanagedprovidercustomer"),
				),
			},

			{
				Config: testAccUpdatemanagedprovidercustomerResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatemanagedprovidercustomer.test", "name", "test-updatemanagedprovidercustomer-updated"),
				),
			},
		},
	})
}

func TestUpdatemanagedprovidercustomerResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatemanagedprovidercustomerResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatemanagedprovidercustomer.test", "name", "test-updatemanagedprovidercustomer"),
				),
			},

			{
				Config: testAccUpdatemanagedprovidercustomerResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatemanagedprovidercustomer.test", "name", "test-updatemanagedprovidercustomer-updated"),
				),
			},
		},
	})
}

func testAccUpdatemanagedprovidercustomerResourceConfig_basic() string {
	return `

resource "umbrella_updatemanagedprovidercustomer" "test" {
  name = "test-updatemanagedprovidercustomer"
  
  name = "test-resource"
  
  customerName = "test-value"
  
  seats = 123
  
}

`
}

func testAccUpdatemanagedprovidercustomerResourceConfig_update() string {
	return `
resource "umbrella_updatemanagedprovidercustomer" "test" {
  name = "test-updatemanagedprovidercustomer-updated"
  
  name = "test-resource-updated"
  
  customerName = "updated-value"
  
  seats = 456
  
}
`
}
