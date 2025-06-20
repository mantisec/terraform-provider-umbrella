package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetmanagedprovidercustomerDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetmanagedprovidercustomerDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getmanagedprovidercustomer.test", "name", "test-getmanagedprovidercustomer"),
				),
			},
		},
	})
}

func testAccGetmanagedprovidercustomerDataSourceConfig_basic() string {
	return `

data "umbrella_getmanagedprovidercustomer" "test" {
  id = "test-id-12345"
}

`
}
