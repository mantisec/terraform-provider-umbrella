package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeletecnameResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletecnameResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletecname.test", "name", "test-deletecname"),
				),
			},

			{
				Config: testAccDeletecnameResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletecname.test", "name", "test-deletecname-updated"),
				),
			},
		},
	})
}

func TestDeletecnameResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletecnameResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletecname.test", "name", "test-deletecname"),
				),
			},

			{
				Config: testAccDeletecnameResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletecname.test", "name", "test-deletecname-updated"),
				),
			},
		},
	})
}

func TestDeletecnameResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletecnameResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletecname.test", "name", "test-deletecname"),
				),
			},

			{
				Config: testAccDeletecnameResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletecname.test", "name", "test-deletecname-updated"),
				),
			},
		},
	})
}

func testAccDeletecnameResourceConfig_basic() string {
	return `

resource "umbrella_deletecname" "test" {
  name = "test-deletecname"
  
  name = "test-resource"
  
}

`
}

func testAccDeletecnameResourceConfig_update() string {
	return `
resource "umbrella_deletecname" "test" {
  name = "test-deletecname-updated"
  
  name = "test-resource-updated"
  
}
`
}
