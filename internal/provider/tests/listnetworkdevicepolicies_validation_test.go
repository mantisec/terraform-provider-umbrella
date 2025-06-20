package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestListnetworkdevicepoliciesSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccListnetworkdevicepoliciesConfig_listnetworkdevicepolicies_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_listnetworkdevicepolicies.test", "listnetworkdevicepolicies", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccListnetworkdevicepoliciesConfig_listnetworkdevicepolicies_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_listnetworkdevicepolicies.test", "listnetworkdevicepolicies", "another-valid-name"),
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
				Config:      testAccListnetworkdevicepoliciesConfig_listnetworkdevicepolicies_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccListnetworkdevicepoliciesConfig_listnetworkdevicepolicies_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccListnetworkdevicepoliciesConfig_name_valid(value string) string {
	return fmt.Sprintf(`

data "umbrella_listnetworkdevicepolicies" "test" {
  name = %s
}

`, value)
}

func testAccListnetworkdevicepoliciesConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

data "umbrella_listnetworkdevicepolicies" "test" {
  name = %s
}

`, value)
}
