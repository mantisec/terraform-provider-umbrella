package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeleteuserResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeleteuserResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deleteuser.test", "name", "test-deleteuser"),
				),
			},

			{
				Config: testAccDeleteuserResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deleteuser.test", "name", "test-deleteuser-updated"),
				),
			},
		},
	})
}

func TestDeleteuserResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeleteuserResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deleteuser.test", "name", "test-deleteuser"),
				),
			},

			{
				Config: testAccDeleteuserResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deleteuser.test", "name", "test-deleteuser-updated"),
				),
			},
		},
	})
}

func TestDeleteuserResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeleteuserResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deleteuser.test", "name", "test-deleteuser"),
				),
			},

			{
				Config: testAccDeleteuserResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deleteuser.test", "name", "test-deleteuser-updated"),
				),
			},
		},
	})
}

func testAccDeleteuserResourceConfig_basic() string {
	return `

resource "umbrella_deleteuser" "test" {
  name = "test-deleteuser"
  
  name = "test-resource"
  
}

`
}

func testAccDeleteuserResourceConfig_update() string {
	return `
resource "umbrella_deleteuser" "test" {
  name = "test-deleteuser-updated"
  
  name = "test-resource-updated"
  
}
`
}
