package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUpdatetunnelSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatetunnelConfig_updatetunnel_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatetunnel.test", "updatetunnel", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatetunnelConfig_updatetunnel_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatetunnel.test", "updatetunnel", "another-valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatetunnelConfig_updatetunnel_valid(1),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatetunnel.test", "updatetunnel", 1),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatetunnelConfig_updatetunnel_valid(100),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatetunnel.test", "updatetunnel", 100),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatetunnelConfig_updatetunnel_valid(999),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatetunnel.test", "updatetunnel", 999),
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
				Config:      testAccUpdatetunnelConfig_updatetunnel_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatetunnelConfig_updatetunnel_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatetunnelConfig_updatetunnel_invalid(-1),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatetunnelConfig_updatetunnel_invalid("not-a-number"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccUpdatetunnelConfig_name_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatetunnel" "test" {
  name = %s
}

`, value)
}

func testAccUpdatetunnelConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatetunnel" "test" {
  name = %s
}

`, value)
}

func testAccUpdatetunnelConfig_siteOriginId_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatetunnel" "test" {
  siteOriginId = %s
}

`, value)
}

func testAccUpdatetunnelConfig_siteOriginId_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatetunnel" "test" {
  siteOriginId = %s
}

`, value)
}

func testAccUpdatetunnelConfig_networkCIDRs_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatetunnel" "test" {
  networkCIDRs = %s
}

`, value)
}

func testAccUpdatetunnelConfig_networkCIDRs_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatetunnel" "test" {
  networkCIDRs = %s
}

`, value)
}

func testAccUpdatetunnelConfig_client_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatetunnel" "test" {
  client = %s
}

`, value)
}

func testAccUpdatetunnelConfig_client_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatetunnel" "test" {
  client = %s
}

`, value)
}
