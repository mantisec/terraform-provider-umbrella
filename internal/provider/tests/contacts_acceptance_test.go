package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestContactsResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccContactsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_contacts.test", "name", "test-contacts"),
				),
			},

			{
				Config: testAccContactsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_contacts.test", "name", "test-contacts-updated"),
				),
			},
		},
	})
}

func TestContactsResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccContactsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_contacts.test", "name", "test-contacts"),
				),
			},

			{
				Config: testAccContactsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_contacts.test", "name", "test-contacts-updated"),
				),
			},
		},
	})
}

func TestContactsResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccContactsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_contacts.test", "name", "test-contacts"),
				),
			},

			{
				Config: testAccContactsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_contacts.test", "name", "test-contacts-updated"),
				),
			},
		},
	})
}

func testAccContactsResourceConfig_basic() string {
	return `

resource "umbrella_contacts" "test" {
  name = "test-contacts"
  
  name = "test-resource"
  
}

`
}

func testAccContactsResourceConfig_update() string {
	return `
resource "umbrella_contacts" "test" {
  name = "test-contacts-updated"
  
  name = "test-resource-updated"
  
}
`
}
