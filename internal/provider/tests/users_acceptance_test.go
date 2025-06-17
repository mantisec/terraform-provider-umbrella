package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUsersResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUsersResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_users.test", "name", "test-users"),
				),
			},

			{
				Config: testAccUsersResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_users.test", "name", "test-users-updated"),
				),
			},
		},
	})
}

func TestUsersResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUsersResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_users.test", "name", "test-users"),
				),
			},

			{
				Config: testAccUsersResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_users.test", "name", "test-users-updated"),
				),
			},
		},
	})
}

func TestUsersResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUsersResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_users.test", "name", "test-users"),
				),
			},

			{
				Config: testAccUsersResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_users.test", "name", "test-users-updated"),
				),
			},
		},
	})
}

func testAccUsersResourceConfig_basic() string {
	return `

resource "umbrella_users" "test" {
  name = "test-users"
  
  name = "test-resource"
  
}

`
}

func testAccUsersResourceConfig_update() string {
	return `
resource "umbrella_users" "test" {
  name = "test-users-updated"
  
  name = "test-resource-updated"
  
}
`
}
