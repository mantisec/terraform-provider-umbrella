package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAddtunnelSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAddtunnelConfig_addtunnel_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_addtunnel.test", "addtunnel", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAddtunnelConfig_addtunnel_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_addtunnel.test", "addtunnel", "another-valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAddtunnelConfig_addtunnel_valid(1),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_addtunnel.test", "addtunnel", 1),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAddtunnelConfig_addtunnel_valid(100),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_addtunnel.test", "addtunnel", 100),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAddtunnelConfig_addtunnel_valid(999),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_addtunnel.test", "addtunnel", 999),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAddtunnelConfig_addtunnel_valid("ASA"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_addtunnel.test", "addtunnel", "ASA"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAddtunnelConfig_addtunnel_valid("FTD"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_addtunnel.test", "addtunnel", "FTD"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAddtunnelConfig_addtunnel_valid("ISR"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_addtunnel.test", "addtunnel", "ISR"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAddtunnelConfig_addtunnel_valid("Meraki MX"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_addtunnel.test", "addtunnel", "Meraki MX"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAddtunnelConfig_addtunnel_valid("Viptela cEdge"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_addtunnel.test", "addtunnel", "Viptela cEdge"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAddtunnelConfig_addtunnel_valid("Viptela vEdge"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_addtunnel.test", "addtunnel", "Viptela vEdge"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAddtunnelConfig_addtunnel_valid("other"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_addtunnel.test", "addtunnel", "other"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAddtunnelConfig_addtunnel_valid("SIG"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_addtunnel.test", "addtunnel", "SIG"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAddtunnelConfig_addtunnel_valid("Private Access"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_addtunnel.test", "addtunnel", "Private Access"),
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
				Config:      testAccAddtunnelConfig_addtunnel_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccAddtunnelConfig_addtunnel_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccAddtunnelConfig_addtunnel_invalid(-1),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccAddtunnelConfig_addtunnel_invalid("not-a-number"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccAddtunnelConfig_addtunnel_invalid("invalid-enum-value"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccAddtunnelConfig_addtunnel_invalid("invalid-enum-value"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccAddtunnelConfig_name_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_addtunnel" "test" {
  name = %s
}

`, value)
}

func testAccAddtunnelConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_addtunnel" "test" {
  name = %s
}

`, value)
}

func testAccAddtunnelConfig_siteOriginId_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_addtunnel" "test" {
  siteOriginId = %s
}

`, value)
}

func testAccAddtunnelConfig_siteOriginId_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_addtunnel" "test" {
  siteOriginId = %s
}

`, value)
}

func testAccAddtunnelConfig_deviceType_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_addtunnel" "test" {
  deviceType = %s
}

`, value)
}

func testAccAddtunnelConfig_deviceType_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_addtunnel" "test" {
  deviceType = %s
}

`, value)
}

func testAccAddtunnelConfig_serviceType_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_addtunnel" "test" {
  serviceType = %s
}

`, value)
}

func testAccAddtunnelConfig_serviceType_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_addtunnel" "test" {
  serviceType = %s
}

`, value)
}

func testAccAddtunnelConfig_networkCIDRs_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_addtunnel" "test" {
  networkCIDRs = %s
}

`, value)
}

func testAccAddtunnelConfig_networkCIDRs_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_addtunnel" "test" {
  networkCIDRs = %s
}

`, value)
}

func testAccAddtunnelConfig_transport_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_addtunnel" "test" {
  transport = %s
}

`, value)
}

func testAccAddtunnelConfig_transport_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_addtunnel" "test" {
  transport = %s
}

`, value)
}

func testAccAddtunnelConfig_authentication_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_addtunnel" "test" {
  authentication = %s
}

`, value)
}

func testAccAddtunnelConfig_authentication_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_addtunnel" "test" {
  authentication = %s
}

`, value)
}
