package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetcustomeraccessrequestDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetcustomeraccessrequestDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getcustomeraccessrequest.test", "name", "test-getcustomeraccessrequest"),
				),
			},
		},
	})
}

func testAccGetcustomeraccessrequestDataSourceConfig_basic() string {
	return `

data "umbrella_getcustomeraccessrequest" "test" {
  id = "test-id-12345"
}

`
}
