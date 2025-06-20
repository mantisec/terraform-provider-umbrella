package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestListprovidercustomersDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccListprovidercustomersDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_listprovidercustomers.test", "name", "test-listprovidercustomers"),
				),
			},
		},
	})
}

func testAccListprovidercustomersDataSourceConfig_basic() string {
	return `

data "umbrella_listprovidercustomers" "test" {
  id = "test-id-12345"
}

`
}
