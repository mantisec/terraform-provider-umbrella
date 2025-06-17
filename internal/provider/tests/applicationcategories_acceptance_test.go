package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestApplicationcategoriesDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationcategoriesDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_applicationcategories.test", "name", "test-applicationcategories"),
				),
			},
		},
	})
}

func testAccApplicationcategoriesDataSourceConfig_basic() string {
	return `

data "umbrella_applicationcategories" "test" {
  id = "test-id-12345"
}

`
}
