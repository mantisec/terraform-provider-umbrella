package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreateprovidercustomerResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreateprovidercustomerResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createprovidercustomer.test", "name", "test-createprovidercustomer"),
				),
			},

			{
				Config: testAccCreateprovidercustomerResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createprovidercustomer.test", "name", "test-createprovidercustomer-updated"),
				),
			},
		},
	})
}

func TestCreateprovidercustomerResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreateprovidercustomerResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createprovidercustomer.test", "name", "test-createprovidercustomer"),
				),
			},

			{
				Config: testAccCreateprovidercustomerResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createprovidercustomer.test", "name", "test-createprovidercustomer-updated"),
				),
			},
		},
	})
}

func TestCreateprovidercustomerResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreateprovidercustomerResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createprovidercustomer.test", "name", "test-createprovidercustomer"),
				),
			},

			{
				Config: testAccCreateprovidercustomerResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createprovidercustomer.test", "name", "test-createprovidercustomer-updated"),
				),
			},
		},
	})
}

func testAccCreateprovidercustomerResourceConfig_basic() string {
	return `

resource "umbrella_createprovidercustomer" "test" {
  name = "test-createprovidercustomer"
  
  name = "test-resource"
  
}

`
}

func testAccCreateprovidercustomerResourceConfig_update() string {
	return `
resource "umbrella_createprovidercustomer" "test" {
  name = "test-createprovidercustomer-updated"
  
  name = "test-resource-updated"
  
}
`
}
