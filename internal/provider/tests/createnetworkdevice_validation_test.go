package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreatenetworkdeviceSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatenetworkdeviceConfig_createnetworkdevice_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createnetworkdevice.test", "createnetworkdevice", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatenetworkdeviceConfig_createnetworkdevice_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createnetworkdevice.test", "createnetworkdevice", "another-valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatenetworkdeviceConfig_createnetworkdevice_valid("valid-string"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createnetworkdevice.test", "createnetworkdevice", "valid-string"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatenetworkdeviceConfig_createnetworkdevice_valid("another-valid"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createnetworkdevice.test", "createnetworkdevice", "another-valid"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatenetworkdeviceConfig_createnetworkdevice_valid("valid-string"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createnetworkdevice.test", "createnetworkdevice", "valid-string"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatenetworkdeviceConfig_createnetworkdevice_valid("another-valid"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createnetworkdevice.test", "createnetworkdevice", "another-valid"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatenetworkdeviceConfig_createnetworkdevice_valid("valid-string"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createnetworkdevice.test", "createnetworkdevice", "valid-string"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatenetworkdeviceConfig_createnetworkdevice_valid("another-valid"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createnetworkdevice.test", "createnetworkdevice", "another-valid"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatenetworkdeviceConfig_createnetworkdevice_valid("valid-string"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createnetworkdevice.test", "createnetworkdevice", "valid-string"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatenetworkdeviceConfig_createnetworkdevice_valid("another-valid"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createnetworkdevice.test", "createnetworkdevice", "another-valid"),
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
				Config:      testAccCreatenetworkdeviceConfig_createnetworkdevice_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCreatenetworkdeviceConfig_createnetworkdevice_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCreatenetworkdeviceConfig_createnetworkdevice_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCreatenetworkdeviceConfig_createnetworkdevice_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCreatenetworkdeviceConfig_createnetworkdevice_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCreatenetworkdeviceConfig_createnetworkdevice_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccCreatenetworkdeviceConfig_name_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createnetworkdevice" "test" {
  name = %s
}

`, value)
}

func testAccCreatenetworkdeviceConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createnetworkdevice" "test" {
  name = %s
}

`, value)
}

func testAccCreatenetworkdeviceConfig_serialNumber_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createnetworkdevice" "test" {
  serialNumber = %s
}

`, value)
}

func testAccCreatenetworkdeviceConfig_serialNumber_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createnetworkdevice" "test" {
  serialNumber = %s
}

`, value)
}

func testAccCreatenetworkdeviceConfig_tag_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createnetworkdevice" "test" {
  tag = %s
}

`, value)
}

func testAccCreatenetworkdeviceConfig_tag_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createnetworkdevice" "test" {
  tag = %s
}

`, value)
}

func testAccCreatenetworkdeviceConfig_model_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createnetworkdevice" "test" {
  model = %s
}

`, value)
}

func testAccCreatenetworkdeviceConfig_model_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createnetworkdevice" "test" {
  model = %s
}

`, value)
}

func testAccCreatenetworkdeviceConfig_macAddress_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createnetworkdevice" "test" {
  macAddress = %s
}

`, value)
}

func testAccCreatenetworkdeviceConfig_macAddress_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createnetworkdevice" "test" {
  macAddress = %s
}

`, value)
}
