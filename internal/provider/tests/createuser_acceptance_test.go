package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreateuserResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreateuserResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createuser.test", "name", "test-createuser"),
				),
			},

			{
				Config: testAccCreateuserResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createuser.test", "name", "test-createuser-updated"),
				),
			},
		},
	})
}

func TestCreateuserResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreateuserResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createuser.test", "name", "test-createuser"),
				),
			},

			{
				Config: testAccCreateuserResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createuser.test", "name", "test-createuser-updated"),
				),
			},
		},
	})
}

func TestCreateuserResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreateuserResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createuser.test", "name", "test-createuser"),
				),
			},

			{
				Config: testAccCreateuserResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createuser.test", "name", "test-createuser-updated"),
				),
			},
		},
	})
}

func testAccCreateuserResourceConfig_basic() string {
	return `

resource "umbrella_createuser" "test" {
  name = "test-createuser"
  
  name = "test-resource"
  
  firstname = "test-value"
  
  lastname = "test-value"
  
  email = "test-value"
  
  password = "test-value"
  
  roleId = 123
  
  timezone = "test-value"
  
}

`
}

func testAccCreateuserResourceConfig_update() string {
	return `
resource "umbrella_createuser" "test" {
  name = "test-createuser-updated"
  
  name = "test-resource-updated"
  
  firstname = "updated-value"
  
  lastname = "updated-value"
  
  email = "updated-value"
  
  password = "updated-value"
  
  roleId = 456
  
  timezone = "updated-value"
  
}
`
}
