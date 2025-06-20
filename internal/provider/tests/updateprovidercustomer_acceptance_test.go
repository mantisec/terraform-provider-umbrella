package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUpdateprovidercustomerResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdateprovidercustomerResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updateprovidercustomer.test", "name", "test-updateprovidercustomer"),
				),
			},

			{
				Config: testAccUpdateprovidercustomerResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updateprovidercustomer.test", "name", "test-updateprovidercustomer-updated"),
				),
			},
		},
	})
}

func TestUpdateprovidercustomerResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdateprovidercustomerResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updateprovidercustomer.test", "name", "test-updateprovidercustomer"),
				),
			},

			{
				Config: testAccUpdateprovidercustomerResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updateprovidercustomer.test", "name", "test-updateprovidercustomer-updated"),
				),
			},
		},
	})
}

func TestUpdateprovidercustomerResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdateprovidercustomerResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updateprovidercustomer.test", "name", "test-updateprovidercustomer"),
				),
			},

			{
				Config: testAccUpdateprovidercustomerResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updateprovidercustomer.test", "name", "test-updateprovidercustomer-updated"),
				),
			},
		},
	})
}

func testAccUpdateprovidercustomerResourceConfig_basic() string {
	return `

resource "umbrella_updateprovidercustomer" "test" {
  name = "test-updateprovidercustomer"
  
  name = "test-resource"
  
}

`
}

func testAccUpdateprovidercustomerResourceConfig_update() string {
	return `
resource "umbrella_updateprovidercustomer" "test" {
  name = "test-updateprovidercustomer-updated"
  
  name = "test-resource-updated"
  
}
`
}
