package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUpdateapplicationResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdateapplicationResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updateapplication.test", "name", "test-updateapplication"),
				),
			},

			{
				Config: testAccUpdateapplicationResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updateapplication.test", "name", "test-updateapplication-updated"),
				),
			},
		},
	})
}

func TestUpdateapplicationResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdateapplicationResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updateapplication.test", "name", "test-updateapplication"),
				),
			},

			{
				Config: testAccUpdateapplicationResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updateapplication.test", "name", "test-updateapplication-updated"),
				),
			},
		},
	})
}

func TestUpdateapplicationResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdateapplicationResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updateapplication.test", "name", "test-updateapplication"),
				),
			},

			{
				Config: testAccUpdateapplicationResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updateapplication.test", "name", "test-updateapplication-updated"),
				),
			},
		},
	})
}

func testAccUpdateapplicationResourceConfig_basic() string {
	return `

resource "umbrella_updateapplication" "test" {
  name = "test-updateapplication"
  
  name = "test-resource"
  
  label = "unreviewed"
  
}

`
}

func testAccUpdateapplicationResourceConfig_update() string {
	return `
resource "umbrella_updateapplication" "test" {
  name = "test-updateapplication-updated"
  
  name = "test-resource-updated"
  
  label = "approved"
  
}
`
}
