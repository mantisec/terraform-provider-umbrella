package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeleteinternalnetworkSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeleteinternalnetworkConfig_deleteinternalnetwork_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deleteinternalnetwork.test", "deleteinternalnetwork", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeleteinternalnetworkConfig_deleteinternalnetwork_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deleteinternalnetwork.test", "deleteinternalnetwork", "another-valid-name"),
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
				Config:      testAccDeleteinternalnetworkConfig_deleteinternalnetwork_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDeleteinternalnetworkConfig_deleteinternalnetwork_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccDeleteinternalnetworkConfig_name_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_deleteinternalnetwork" "test" {
  name = %s
}

`, value)
}

func testAccDeleteinternalnetworkConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_deleteinternalnetwork" "test" {
  name = %s
}

`, value)
}
