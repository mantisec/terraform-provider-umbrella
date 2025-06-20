package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUpdatemanagedprovidercustomerSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatemanagedprovidercustomerConfig_updatemanagedprovidercustomer_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatemanagedprovidercustomer.test", "updatemanagedprovidercustomer", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatemanagedprovidercustomerConfig_updatemanagedprovidercustomer_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatemanagedprovidercustomer.test", "updatemanagedprovidercustomer", "another-valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatemanagedprovidercustomerConfig_updatemanagedprovidercustomer_valid("valid-string"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatemanagedprovidercustomer.test", "updatemanagedprovidercustomer", "valid-string"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatemanagedprovidercustomerConfig_updatemanagedprovidercustomer_valid("another-valid"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatemanagedprovidercustomer.test", "updatemanagedprovidercustomer", "another-valid"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatemanagedprovidercustomerConfig_updatemanagedprovidercustomer_valid(1),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatemanagedprovidercustomer.test", "updatemanagedprovidercustomer", 1),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatemanagedprovidercustomerConfig_updatemanagedprovidercustomer_valid(100),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatemanagedprovidercustomer.test", "updatemanagedprovidercustomer", 100),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatemanagedprovidercustomerConfig_updatemanagedprovidercustomer_valid(999),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatemanagedprovidercustomer.test", "updatemanagedprovidercustomer", 999),
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
				Config:      testAccUpdatemanagedprovidercustomerConfig_updatemanagedprovidercustomer_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatemanagedprovidercustomerConfig_updatemanagedprovidercustomer_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatemanagedprovidercustomerConfig_updatemanagedprovidercustomer_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatemanagedprovidercustomerConfig_updatemanagedprovidercustomer_invalid(-1),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatemanagedprovidercustomerConfig_updatemanagedprovidercustomer_invalid("not-a-number"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccUpdatemanagedprovidercustomerConfig_name_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatemanagedprovidercustomer" "test" {
  name = %s
}

`, value)
}

func testAccUpdatemanagedprovidercustomerConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatemanagedprovidercustomer" "test" {
  name = %s
}

`, value)
}

func testAccUpdatemanagedprovidercustomerConfig_customerName_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatemanagedprovidercustomer" "test" {
  customerName = %s
}

`, value)
}

func testAccUpdatemanagedprovidercustomerConfig_customerName_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatemanagedprovidercustomer" "test" {
  customerName = %s
}

`, value)
}

func testAccUpdatemanagedprovidercustomerConfig_seats_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatemanagedprovidercustomer" "test" {
  seats = %s
}

`, value)
}

func testAccUpdatemanagedprovidercustomerConfig_seats_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatemanagedprovidercustomer" "test" {
  seats = %s
}

`, value)
}
