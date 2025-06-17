package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCustomersResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomersResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_customers.test", "name", "test-customers"),
				),
			},

			{
				Config: testAccCustomersResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_customers.test", "name", "test-customers-updated"),
				),
			},
		},
	})
}

func TestCustomersResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomersResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_customers.test", "name", "test-customers"),
				),
			},

			{
				Config: testAccCustomersResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_customers.test", "name", "test-customers-updated"),
				),
			},
		},
	})
}

func TestCustomersResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomersResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_customers.test", "name", "test-customers"),
				),
			},

			{
				Config: testAccCustomersResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_customers.test", "name", "test-customers-updated"),
				),
			},
		},
	})
}

func testAccCustomersResourceConfig_basic() string {
	return `

resource "umbrella_customers" "test" {
  name = "test-customers"
  
  name = "test-resource"
  
}

`
}

func testAccCustomersResourceConfig_update() string {
	return `
resource "umbrella_customers" "test" {
  name = "test-customers-updated"
  
  name = "test-resource-updated"
  
}
`
}
