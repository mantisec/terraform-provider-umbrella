package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGettunnelpoliciesDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGettunnelpoliciesDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_gettunnelpolicies.test", "name", "test-gettunnelpolicies"),
				),
			},
		},
	})
}

func testAccGettunnelpoliciesDataSourceConfig_basic() string {
	return `

data "umbrella_gettunnelpolicies" "test" {
  id = "test-id-12345"
}

`
}
