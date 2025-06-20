package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUpdatecontactSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "another-valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("valid-string"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "valid-string"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("another-valid"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "another-valid"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("billing"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "billing"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("blockfeedback"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "blockfeedback"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("report"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "report"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("serviceupdate"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "serviceupdate"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("support"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "support"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("distributor"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "distributor"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("valid-string"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "valid-string"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("another-valid"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "another-valid"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("valid-string"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "valid-string"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("another-valid"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "another-valid"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("valid-string"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "valid-string"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("another-valid"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "another-valid"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("valid-string"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "valid-string"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("another-valid"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "another-valid"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("valid-string"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "valid-string"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("another-valid"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "another-valid"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("valid-string"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "valid-string"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("another-valid"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "another-valid"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("valid-string"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "valid-string"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("another-valid"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "another-valid"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("valid-string"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "valid-string"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("another-valid"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "another-valid"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("valid-string"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "valid-string"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("another-valid"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "another-valid"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("valid-string"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "valid-string"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("another-valid"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "another-valid"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("no"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "no"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("yes"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "yes"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("valid-string"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "valid-string"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecontactConfig_updatecontact_valid("another-valid"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecontact.test", "updatecontact", "another-valid"),
				),
			},
		},
	})

	// Test invalid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatecontactConfig_updatecontact_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatecontactConfig_updatecontact_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatecontactConfig_updatecontact_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatecontactConfig_updatecontact_invalid("invalid-enum-value"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatecontactConfig_updatecontact_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatecontactConfig_updatecontact_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatecontactConfig_updatecontact_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatecontactConfig_updatecontact_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatecontactConfig_updatecontact_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatecontactConfig_updatecontact_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatecontactConfig_updatecontact_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatecontactConfig_updatecontact_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatecontactConfig_updatecontact_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatecontactConfig_updatecontact_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatecontactConfig_updatecontact_invalid("invalid-enum-value"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatecontactConfig_updatecontact_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccUpdatecontactConfig_name_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  name = %s
}

`, value)
}

func testAccUpdatecontactConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  name = %s
}

`, value)
}

func testAccUpdatecontactConfig_phoneNumber_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  phoneNumber = %s
}

`, value)
}

func testAccUpdatecontactConfig_phoneNumber_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  phoneNumber = %s
}

`, value)
}

func testAccUpdatecontactConfig_contactType_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  contactType = %s
}

`, value)
}

func testAccUpdatecontactConfig_contactType_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  contactType = %s
}

`, value)
}

func testAccUpdatecontactConfig_city_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  city = %s
}

`, value)
}

func testAccUpdatecontactConfig_city_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  city = %s
}

`, value)
}

func testAccUpdatecontactConfig_streetAddress_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  streetAddress = %s
}

`, value)
}

func testAccUpdatecontactConfig_streetAddress_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  streetAddress = %s
}

`, value)
}

func testAccUpdatecontactConfig_emailAddress_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  emailAddress = %s
}

`, value)
}

func testAccUpdatecontactConfig_emailAddress_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  emailAddress = %s
}

`, value)
}

func testAccUpdatecontactConfig_phoneNumber2_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  phoneNumber2 = %s
}

`, value)
}

func testAccUpdatecontactConfig_phoneNumber2_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  phoneNumber2 = %s
}

`, value)
}

func testAccUpdatecontactConfig_faxNumber_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  faxNumber = %s
}

`, value)
}

func testAccUpdatecontactConfig_faxNumber_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  faxNumber = %s
}

`, value)
}

func testAccUpdatecontactConfig_lastName_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  lastName = %s
}

`, value)
}

func testAccUpdatecontactConfig_lastName_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  lastName = %s
}

`, value)
}

func testAccUpdatecontactConfig_streetAddress2_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  streetAddress2 = %s
}

`, value)
}

func testAccUpdatecontactConfig_streetAddress2_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  streetAddress2 = %s
}

`, value)
}

func testAccUpdatecontactConfig_state_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  state = %s
}

`, value)
}

func testAccUpdatecontactConfig_state_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  state = %s
}

`, value)
}

func testAccUpdatecontactConfig_zipCode_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  zipCode = %s
}

`, value)
}

func testAccUpdatecontactConfig_zipCode_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  zipCode = %s
}

`, value)
}

func testAccUpdatecontactConfig_countryCode_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  countryCode = %s
}

`, value)
}

func testAccUpdatecontactConfig_countryCode_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  countryCode = %s
}

`, value)
}

func testAccUpdatecontactConfig_settings_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  settings = %s
}

`, value)
}

func testAccUpdatecontactConfig_settings_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  settings = %s
}

`, value)
}

func testAccUpdatecontactConfig_primaryContact_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  primaryContact = %s
}

`, value)
}

func testAccUpdatecontactConfig_primaryContact_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  primaryContact = %s
}

`, value)
}

func testAccUpdatecontactConfig_firstName_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  firstName = %s
}

`, value)
}

func testAccUpdatecontactConfig_firstName_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecontact" "test" {
  firstName = %s
}

`, value)
}
