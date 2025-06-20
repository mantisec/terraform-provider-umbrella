package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeletemangedprovidercustomerResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletemangedprovidercustomerResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletemangedprovidercustomer.test", "name", "test-deletemangedprovidercustomer"),
				),
			},

			{
				Config: testAccDeletemangedprovidercustomerResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletemangedprovidercustomer.test", "name", "test-deletemangedprovidercustomer-updated"),
				),
			},
		},
	})
}

func TestDeletemangedprovidercustomerResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletemangedprovidercustomerResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletemangedprovidercustomer.test", "name", "test-deletemangedprovidercustomer"),
				),
			},

			{
				Config: testAccDeletemangedprovidercustomerResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletemangedprovidercustomer.test", "name", "test-deletemangedprovidercustomer-updated"),
				),
			},
		},
	})
}

func TestDeletemangedprovidercustomerResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletemangedprovidercustomerResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletemangedprovidercustomer.test", "name", "test-deletemangedprovidercustomer"),
				),
			},

			{
				Config: testAccDeletemangedprovidercustomerResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletemangedprovidercustomer.test", "name", "test-deletemangedprovidercustomer-updated"),
				),
			},
		},
	})
}

func testAccDeletemangedprovidercustomerResourceConfig_basic() string {
	return `

resource "umbrella_deletemangedprovidercustomer" "test" {
  name = "test-deletemangedprovidercustomer"
  
  name = "test-resource"
  
}

`
}

func testAccDeletemangedprovidercustomerResourceConfig_update() string {
	return `
resource "umbrella_deletemangedprovidercustomer" "test" {
  name = "test-deletemangedprovidercustomer-updated"
  
  name = "test-resource-updated"
  
}
`
}
