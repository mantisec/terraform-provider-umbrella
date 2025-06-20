package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreatemanagedprovidercustomerSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatemanagedprovidercustomerConfig_createmanagedprovidercustomer_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createmanagedprovidercustomer.test", "createmanagedprovidercustomer", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatemanagedprovidercustomerConfig_createmanagedprovidercustomer_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createmanagedprovidercustomer.test", "createmanagedprovidercustomer", "another-valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatemanagedprovidercustomerConfig_createmanagedprovidercustomer_valid("valid-string"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createmanagedprovidercustomer.test", "createmanagedprovidercustomer", "valid-string"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatemanagedprovidercustomerConfig_createmanagedprovidercustomer_valid("another-valid"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createmanagedprovidercustomer.test", "createmanagedprovidercustomer", "another-valid"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatemanagedprovidercustomerConfig_createmanagedprovidercustomer_valid(1),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createmanagedprovidercustomer.test", "createmanagedprovidercustomer", 1),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatemanagedprovidercustomerConfig_createmanagedprovidercustomer_valid(100),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createmanagedprovidercustomer.test", "createmanagedprovidercustomer", 100),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatemanagedprovidercustomerConfig_createmanagedprovidercustomer_valid(999),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createmanagedprovidercustomer.test", "createmanagedprovidercustomer", 999),
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
				Config:      testAccCreatemanagedprovidercustomerConfig_createmanagedprovidercustomer_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCreatemanagedprovidercustomerConfig_createmanagedprovidercustomer_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCreatemanagedprovidercustomerConfig_createmanagedprovidercustomer_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCreatemanagedprovidercustomerConfig_createmanagedprovidercustomer_invalid(-1),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCreatemanagedprovidercustomerConfig_createmanagedprovidercustomer_invalid("not-a-number"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccCreatemanagedprovidercustomerConfig_name_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createmanagedprovidercustomer" "test" {
  name = %s
}

`, value)
}

func testAccCreatemanagedprovidercustomerConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createmanagedprovidercustomer" "test" {
  name = %s
}

`, value)
}

func testAccCreatemanagedprovidercustomerConfig_customerName_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createmanagedprovidercustomer" "test" {
  customerName = %s
}

`, value)
}

func testAccCreatemanagedprovidercustomerConfig_customerName_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createmanagedprovidercustomer" "test" {
  customerName = %s
}

`, value)
}

func testAccCreatemanagedprovidercustomerConfig_seats_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createmanagedprovidercustomer" "test" {
  seats = %s
}

`, value)
}

func testAccCreatemanagedprovidercustomerConfig_seats_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createmanagedprovidercustomer" "test" {
  seats = %s
}

`, value)
}
