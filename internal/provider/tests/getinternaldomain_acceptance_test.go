package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetinternaldomainDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetinternaldomainDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getinternaldomain.test", "name", "test-getinternaldomain"),
				),
			},
		},
	})
}

func testAccGetinternaldomainDataSourceConfig_basic() string {
	return `

data "umbrella_getinternaldomain" "test" {
  id = "test-id-12345"
}

`
}
