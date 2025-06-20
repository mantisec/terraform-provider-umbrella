package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetcustomersubscriptiondetailsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetcustomersubscriptiondetailsDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getcustomersubscriptiondetails.test", "name", "test-getcustomersubscriptiondetails"),
				),
			},
		},
	})
}

func testAccGetcustomersubscriptiondetailsDataSourceConfig_basic() string {
	return `

data "umbrella_getcustomersubscriptiondetails" "test" {
  id = "test-id-12345"
}

`
}
