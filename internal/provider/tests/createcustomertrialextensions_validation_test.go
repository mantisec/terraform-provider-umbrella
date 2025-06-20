package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreatecustomertrialextensionsSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatecustomertrialextensionsConfig_createcustomertrialextensions_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createcustomertrialextensions.test", "createcustomertrialextensions", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatecustomertrialextensionsConfig_createcustomertrialextensions_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createcustomertrialextensions.test", "createcustomertrialextensions", "another-valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatecustomertrialextensionsConfig_createcustomertrialextensions_valid(1),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createcustomertrialextensions.test", "createcustomertrialextensions", 1),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatecustomertrialextensionsConfig_createcustomertrialextensions_valid(100),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createcustomertrialextensions.test", "createcustomertrialextensions", 100),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatecustomertrialextensionsConfig_createcustomertrialextensions_valid(999),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createcustomertrialextensions.test", "createcustomertrialextensions", 999),
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
				Config:      testAccCreatecustomertrialextensionsConfig_createcustomertrialextensions_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCreatecustomertrialextensionsConfig_createcustomertrialextensions_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCreatecustomertrialextensionsConfig_createcustomertrialextensions_invalid(-1),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCreatecustomertrialextensionsConfig_createcustomertrialextensions_invalid("not-a-number"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccCreatecustomertrialextensionsConfig_name_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createcustomertrialextensions" "test" {
  name = %s
}

`, value)
}

func testAccCreatecustomertrialextensionsConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createcustomertrialextensions" "test" {
  name = %s
}

`, value)
}

func testAccCreatecustomertrialextensionsConfig_trialExtensionDays_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createcustomertrialextensions" "test" {
  trialExtensionDays = %s
}

`, value)
}

func testAccCreatecustomertrialextensionsConfig_trialExtensionDays_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createcustomertrialextensions" "test" {
  trialExtensionDays = %s
}

`, value)
}
