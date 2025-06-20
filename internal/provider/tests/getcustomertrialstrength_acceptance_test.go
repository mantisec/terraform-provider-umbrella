package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetcustomertrialstrengthDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetcustomertrialstrengthDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getcustomertrialstrength.test", "name", "test-getcustomertrialstrength"),
				),
			},
		},
	})
}

func testAccGetcustomertrialstrengthDataSourceConfig_basic() string {
	return `

data "umbrella_getcustomertrialstrength" "test" {
  id = "test-id-12345"
}

`
}
