package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestPuttrialconversionsSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPuttrialconversionsConfig_puttrialconversions_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_puttrialconversions.test", "puttrialconversions", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPuttrialconversionsConfig_puttrialconversions_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_puttrialconversions.test", "puttrialconversions", "another-valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPuttrialconversionsConfig_puttrialconversions_valid(1),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_puttrialconversions.test", "puttrialconversions", 1),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPuttrialconversionsConfig_puttrialconversions_valid(100),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_puttrialconversions.test", "puttrialconversions", 100),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPuttrialconversionsConfig_puttrialconversions_valid(999),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_puttrialconversions.test", "puttrialconversions", 999),
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
				Config:      testAccPuttrialconversionsConfig_puttrialconversions_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccPuttrialconversionsConfig_puttrialconversions_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccPuttrialconversionsConfig_puttrialconversions_invalid(-1),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccPuttrialconversionsConfig_puttrialconversions_invalid("not-a-number"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccPuttrialconversionsConfig_name_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_puttrialconversions" "test" {
  name = %s
}

`, value)
}

func testAccPuttrialconversionsConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_puttrialconversions" "test" {
  name = %s
}

`, value)
}

func testAccPuttrialconversionsConfig_packageId_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_puttrialconversions" "test" {
  packageId = %s
}

`, value)
}

func testAccPuttrialconversionsConfig_packageId_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_puttrialconversions" "test" {
  packageId = %s
}

`, value)
}
