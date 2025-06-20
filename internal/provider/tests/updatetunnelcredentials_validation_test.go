package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUpdatetunnelcredentialsSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatetunnelcredentialsConfig_updatetunnelcredentials_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatetunnelcredentials.test", "updatetunnelcredentials", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatetunnelcredentialsConfig_updatetunnelcredentials_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatetunnelcredentials.test", "updatetunnelcredentials", "another-valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatetunnelcredentialsConfig_updatetunnelcredentials_valid(true),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatetunnelcredentials.test", "updatetunnelcredentials", true),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatetunnelcredentialsConfig_updatetunnelcredentials_valid(false),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatetunnelcredentials.test", "updatetunnelcredentials", false),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatetunnelcredentialsConfig_updatetunnelcredentials_valid(true),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatetunnelcredentials.test", "updatetunnelcredentials", true),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatetunnelcredentialsConfig_updatetunnelcredentials_valid(false),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatetunnelcredentials.test", "updatetunnelcredentials", false),
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
				Config:      testAccUpdatetunnelcredentialsConfig_updatetunnelcredentials_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatetunnelcredentialsConfig_updatetunnelcredentials_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatetunnelcredentialsConfig_updatetunnelcredentials_invalid("not-a-boolean"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatetunnelcredentialsConfig_updatetunnelcredentials_invalid(1),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatetunnelcredentialsConfig_updatetunnelcredentials_invalid("not-a-boolean"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatetunnelcredentialsConfig_updatetunnelcredentials_invalid(1),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccUpdatetunnelcredentialsConfig_name_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatetunnelcredentials" "test" {
  name = %s
}

`, value)
}

func testAccUpdatetunnelcredentialsConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatetunnelcredentials" "test" {
  name = %s
}

`, value)
}

func testAccUpdatetunnelcredentialsConfig_deprecateCurrentKeys_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatetunnelcredentials" "test" {
  deprecateCurrentKeys = %s
}

`, value)
}

func testAccUpdatetunnelcredentialsConfig_deprecateCurrentKeys_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatetunnelcredentials" "test" {
  deprecateCurrentKeys = %s
}

`, value)
}

func testAccUpdatetunnelcredentialsConfig_autoRotate_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatetunnelcredentials" "test" {
  autoRotate = %s
}

`, value)
}

func testAccUpdatetunnelcredentialsConfig_autoRotate_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatetunnelcredentials" "test" {
  autoRotate = %s
}

`, value)
}

func testAccUpdatetunnelcredentialsConfig_psk_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatetunnelcredentials" "test" {
  psk = %s
}

`, value)
}

func testAccUpdatetunnelcredentialsConfig_psk_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatetunnelcredentials" "test" {
  psk = %s
}

`, value)
}
