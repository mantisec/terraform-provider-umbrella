package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestListvirtualappliancesDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccListvirtualappliancesDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_listvirtualappliances.test", "name", "test-listvirtualappliances"),
				),
			},
		},
	})
}

func testAccListvirtualappliancesDataSourceConfig_basic() string {
	return `

data "umbrella_listvirtualappliances" "test" {
  id = "test-id-12345"
}

`
}
