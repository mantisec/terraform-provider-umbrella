package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestIdentityDistributionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIdentityDistributionDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_identity_distribution.test", "name", "test-identity_distribution"),
				),
			},
		},
	})
}

func testAccIdentityDistributionDataSourceConfig_basic() string {
	return `

data "umbrella_identity_distribution" "test" {
  id = "test-id-12345"
}

`
}
