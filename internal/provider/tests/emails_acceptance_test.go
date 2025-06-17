package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestEmailsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccEmailsDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_emails.test", "name", "test-emails"),
				),
			},
		},
	})
}

func testAccEmailsDataSourceConfig_basic() string {
	return `

data "umbrella_emails" "test" {
  id = "test-id-12345"
}

`
}
