package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUpdatecnameResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecnameResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecname.test", "name", "test-updatecname"),
				),
			},

			{
				Config: testAccUpdatecnameResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatecname.test", "name", "test-updatecname-updated"),
				),
			},
		},
	})
}

func TestUpdatecnameResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecnameResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecname.test", "name", "test-updatecname"),
				),
			},

			{
				Config: testAccUpdatecnameResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatecname.test", "name", "test-updatecname-updated"),
				),
			},
		},
	})
}

func TestUpdatecnameResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecnameResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecname.test", "name", "test-updatecname"),
				),
			},

			{
				Config: testAccUpdatecnameResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatecname.test", "name", "test-updatecname-updated"),
				),
			},
		},
	})
}

func testAccUpdatecnameResourceConfig_basic() string {
	return `

resource "umbrella_updatecname" "test" {
  name = "test-updatecname"
  
  name = "test-resource"
  
}

`
}

func testAccUpdatecnameResourceConfig_update() string {
	return `
resource "umbrella_updatecname" "test" {
  name = "test-updatecname-updated"
  
  name = "test-resource-updated"
  
}
`
}
