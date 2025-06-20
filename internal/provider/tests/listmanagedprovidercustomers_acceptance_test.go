package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestListmanagedprovidercustomersDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccListmanagedprovidercustomersDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_listmanagedprovidercustomers.test", "name", "test-listmanagedprovidercustomers"),
				),
			},
		},
	})
}

func testAccListmanagedprovidercustomersDataSourceConfig_basic() string {
	return `

data "umbrella_listmanagedprovidercustomers" "test" {
  id = "test-id-12345"
}

`
}
