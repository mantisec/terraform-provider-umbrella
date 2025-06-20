package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUpdatecontactResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "name", "test-updatecontact"),
				),
			},

			{
				Config: testAccUpdatecontactResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "name", "test-updatecontact-updated"),
				),
			},
		},
	})
}

func TestUpdatecontactResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "name", "test-updatecontact"),
				),
			},

			{
				Config: testAccUpdatecontactResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "name", "test-updatecontact-updated"),
				),
			},
		},
	})
}

func TestUpdatecontactResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "name", "test-updatecontact"),
				),
			},

			{
				Config: testAccUpdatecontactResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "name", "test-updatecontact-updated"),
				),
			},
		},
	})
}

func testAccUpdatecontactResourceConfig_basic() string {
	return `

resource "umbrella_updatecontact" "test" {
  name = "test-updatecontact"
  
  name = "test-resource"
  
  phoneNumber = "test-value"
  
  contactType = "billing"
  
  city = "test-value"
  
  streetAddress = "test-value"
  
  emailAddress = "test-value"
  
  phoneNumber2 = "test-value"
  
  faxNumber = "test-value"
  
  lastName = "test-value"
  
  streetAddress2 = "test-value"
  
  state = "test-value"
  
  zipCode = "test-value"
  
  countryCode = "test-value"
  
  settings = {}
  
  primaryContact = "no"
  
  firstName = "test-value"
  
}

`
}

func testAccUpdatecontactResourceConfig_update() string {
	return `
resource "umbrella_updatecontact" "test" {
  name = "test-updatecontact-updated"
  
  name = "test-resource-updated"
  
  phoneNumber = "updated-value"
  
  contactType = "blockfeedback"
  
  city = "updated-value"
  
  streetAddress = "updated-value"
  
  emailAddress = "updated-value"
  
  phoneNumber2 = "updated-value"
  
  faxNumber = "updated-value"
  
  lastName = "updated-value"
  
  streetAddress2 = "updated-value"
  
  state = "updated-value"
  
  zipCode = "updated-value"
  
  countryCode = "updated-value"
  
  settings = {}
  
  primaryContact = "yes"
  
  firstName = "updated-value"
  
}
`
}
