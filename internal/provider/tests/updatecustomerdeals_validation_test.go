package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUpdatecustomerdealsSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecustomerdealsConfig_updatecustomerdeals_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecustomerdeals.test", "updatecustomerdeals", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecustomerdealsConfig_updatecustomerdeals_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecustomerdeals.test", "updatecustomerdeals", "another-valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecustomerdealsConfig_updatecustomerdeals_valid(1),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecustomerdeals.test", "updatecustomerdeals", 1),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecustomerdealsConfig_updatecustomerdeals_valid(100),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecustomerdeals.test", "updatecustomerdeals", 100),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecustomerdealsConfig_updatecustomerdeals_valid(999),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecustomerdeals.test", "updatecustomerdeals", 999),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecustomerdealsConfig_updatecustomerdeals_valid(1),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecustomerdeals.test", "updatecustomerdeals", 1),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecustomerdealsConfig_updatecustomerdeals_valid(100),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecustomerdeals.test", "updatecustomerdeals", 100),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecustomerdealsConfig_updatecustomerdeals_valid(999),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecustomerdeals.test", "updatecustomerdeals", 999),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecustomerdealsConfig_updatecustomerdeals_valid(1),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecustomerdeals.test", "updatecustomerdeals", 1),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecustomerdealsConfig_updatecustomerdeals_valid(100),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecustomerdeals.test", "updatecustomerdeals", 100),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecustomerdealsConfig_updatecustomerdeals_valid(999),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecustomerdeals.test", "updatecustomerdeals", 999),
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
				Config:      testAccUpdatecustomerdealsConfig_updatecustomerdeals_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatecustomerdealsConfig_updatecustomerdeals_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatecustomerdealsConfig_updatecustomerdeals_invalid(-1),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatecustomerdealsConfig_updatecustomerdeals_invalid("not-a-number"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatecustomerdealsConfig_updatecustomerdeals_invalid(-1),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatecustomerdealsConfig_updatecustomerdeals_invalid("not-a-number"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatecustomerdealsConfig_updatecustomerdeals_invalid(-1),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUpdatecustomerdealsConfig_updatecustomerdeals_invalid("not-a-number"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccUpdatecustomerdealsConfig_name_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecustomerdeals" "test" {
  name = %s
}

`, value)
}

func testAccUpdatecustomerdealsConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecustomerdeals" "test" {
  name = %s
}

`, value)
}

func testAccUpdatecustomerdealsConfig_ccoid_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecustomerdeals" "test" {
  ccoid = %s
}

`, value)
}

func testAccUpdatecustomerdealsConfig_ccoid_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecustomerdeals" "test" {
  ccoid = %s
}

`, value)
}

func testAccUpdatecustomerdealsConfig_customerId_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecustomerdeals" "test" {
  customerId = %s
}

`, value)
}

func testAccUpdatecustomerdealsConfig_customerId_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecustomerdeals" "test" {
  customerId = %s
}

`, value)
}

func testAccUpdatecustomerdealsConfig_quoteId_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecustomerdeals" "test" {
  quoteId = %s
}

`, value)
}

func testAccUpdatecustomerdealsConfig_quoteId_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecustomerdeals" "test" {
  quoteId = %s
}

`, value)
}

func testAccUpdatecustomerdealsConfig_majorLineItems_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecustomerdeals" "test" {
  majorLineItems = %s
}

`, value)
}

func testAccUpdatecustomerdealsConfig_majorLineItems_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_updatecustomerdeals" "test" {
  majorLineItems = %s
}

`, value)
}
