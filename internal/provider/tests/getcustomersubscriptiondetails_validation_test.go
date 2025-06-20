package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetcustomersubscriptiondetailsSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetcustomersubscriptiondetailsConfig_getcustomersubscriptiondetails_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getcustomersubscriptiondetails.test", "getcustomersubscriptiondetails", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetcustomersubscriptiondetailsConfig_getcustomersubscriptiondetails_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getcustomersubscriptiondetails.test", "getcustomersubscriptiondetails", "another-valid-name"),
				),
			},
		},
	})

	// Test invalid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccGetcustomersubscriptiondetailsConfig_getcustomersubscriptiondetails_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccGetcustomersubscriptiondetailsConfig_getcustomersubscriptiondetails_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccGetcustomersubscriptiondetailsConfig_name_valid(value string) string {
	return fmt.Sprintf(`

data "umbrella_getcustomersubscriptiondetails" "test" {
  name = %s
}

`, value)
}

func testAccGetcustomersubscriptiondetailsConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

data "umbrella_getcustomersubscriptiondetails" "test" {
  name = %s
}

`, value)
}
