package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreatecontactResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatecontactResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createcontact.test", "name", "test-createcontact"),
				),
			},

			{
				Config: testAccCreatecontactResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createcontact.test", "name", "test-createcontact-updated"),
				),
			},
		},
	})
}

func TestCreatecontactResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatecontactResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createcontact.test", "name", "test-createcontact"),
				),
			},

			{
				Config: testAccCreatecontactResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createcontact.test", "name", "test-createcontact-updated"),
				),
			},
		},
	})
}

func TestCreatecontactResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatecontactResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createcontact.test", "name", "test-createcontact"),
				),
			},

			{
				Config: testAccCreatecontactResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createcontact.test", "name", "test-createcontact-updated"),
				),
			},
		},
	})
}

func testAccCreatecontactResourceConfig_basic() string {
	return `

resource "umbrella_createcontact" "test" {
  name = "test-createcontact"
  
  name = "test-resource"
  
  zipCode = "test-value"
  
  countryCode = "test-value"
  
  settings = {}
  
  primaryContact = "no"
  
  firstName = "test-value"
  
  lastName = "test-value"
  
  streetAddress2 = "test-value"
  
  state = "test-value"
  
  contactType = "billing"
  
  city = "test-value"
  
  phoneNumber = "test-value"
  
  streetAddress = "test-value"
  
  phoneNumber2 = "test-value"
  
  faxNumber = "test-value"
  
  emailAddress = "test-value"
  
}

`
}

func testAccCreatecontactResourceConfig_update() string {
	return `
resource "umbrella_createcontact" "test" {
  name = "test-createcontact-updated"
  
  name = "test-resource-updated"
  
  zipCode = "updated-value"
  
  countryCode = "updated-value"
  
  settings = {}
  
  primaryContact = "yes"
  
  firstName = "updated-value"
  
  lastName = "updated-value"
  
  streetAddress2 = "updated-value"
  
  state = "updated-value"
  
  contactType = "blockfeedback"
  
  city = "updated-value"
  
  phoneNumber = "updated-value"
  
  streetAddress = "updated-value"
  
  phoneNumber2 = "updated-value"
  
  faxNumber = "updated-value"
  
  emailAddress = "updated-value"
  
}
`
}
