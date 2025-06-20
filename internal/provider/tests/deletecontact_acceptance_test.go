package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeletecontactResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletecontactResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletecontact.test", "name", "test-deletecontact"),
				),
			},

			{
				Config: testAccDeletecontactResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletecontact.test", "name", "test-deletecontact-updated"),
				),
			},
		},
	})
}

func TestDeletecontactResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletecontactResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletecontact.test", "name", "test-deletecontact"),
				),
			},

			{
				Config: testAccDeletecontactResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletecontact.test", "name", "test-deletecontact-updated"),
				),
			},
		},
	})
}

func TestDeletecontactResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletecontactResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletecontact.test", "name", "test-deletecontact"),
				),
			},

			{
				Config: testAccDeletecontactResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletecontact.test", "name", "test-deletecontact-updated"),
				),
			},
		},
	})
}

func testAccDeletecontactResourceConfig_basic() string {
	return `

resource "umbrella_deletecontact" "test" {
  name = "test-deletecontact"
  
  name = "test-resource"
  
}

`
}

func testAccDeletecontactResourceConfig_update() string {
	return `
resource "umbrella_deletecontact" "test" {
  name = "test-deletecontact-updated"
  
  name = "test-resource-updated"
  
}
`
}
