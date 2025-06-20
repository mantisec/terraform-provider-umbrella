package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreatedestinationlistSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatedestinationlistConfig_createdestinationlist_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createdestinationlist.test", "createdestinationlist", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatedestinationlistConfig_createdestinationlist_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createdestinationlist.test", "createdestinationlist", "another-valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatedestinationlistConfig_createdestinationlist_valid(1),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createdestinationlist.test", "createdestinationlist", 1),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatedestinationlistConfig_createdestinationlist_valid(100),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createdestinationlist.test", "createdestinationlist", 100),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatedestinationlistConfig_createdestinationlist_valid(999),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createdestinationlist.test", "createdestinationlist", 999),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatedestinationlistConfig_createdestinationlist_valid("allow"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createdestinationlist.test", "createdestinationlist", "allow"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatedestinationlistConfig_createdestinationlist_valid("block"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createdestinationlist.test", "createdestinationlist", "block"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatedestinationlistConfig_createdestinationlist_valid(true),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createdestinationlist.test", "createdestinationlist", true),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatedestinationlistConfig_createdestinationlist_valid(false),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createdestinationlist.test", "createdestinationlist", false),
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
				Config:      testAccCreatedestinationlistConfig_createdestinationlist_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCreatedestinationlistConfig_createdestinationlist_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCreatedestinationlistConfig_createdestinationlist_invalid(-1),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCreatedestinationlistConfig_createdestinationlist_invalid("not-a-number"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCreatedestinationlistConfig_createdestinationlist_invalid("invalid-enum-value"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCreatedestinationlistConfig_createdestinationlist_invalid("not-a-boolean"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCreatedestinationlistConfig_createdestinationlist_invalid(1),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccCreatedestinationlistConfig_name_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createdestinationlist" "test" {
  name = %s
}

`, value)
}

func testAccCreatedestinationlistConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createdestinationlist" "test" {
  name = %s
}

`, value)
}

func testAccCreatedestinationlistConfig_bundleTypeId_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createdestinationlist" "test" {
  bundleTypeId = %s
}

`, value)
}

func testAccCreatedestinationlistConfig_bundleTypeId_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createdestinationlist" "test" {
  bundleTypeId = %s
}

`, value)
}

func testAccCreatedestinationlistConfig_destinations_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createdestinationlist" "test" {
  destinations = %s
}

`, value)
}

func testAccCreatedestinationlistConfig_destinations_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createdestinationlist" "test" {
  destinations = %s
}

`, value)
}

func testAccCreatedestinationlistConfig_access_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createdestinationlist" "test" {
  access = %s
}

`, value)
}

func testAccCreatedestinationlistConfig_access_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createdestinationlist" "test" {
  access = %s
}

`, value)
}

func testAccCreatedestinationlistConfig_isGlobal_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createdestinationlist" "test" {
  isGlobal = %s
}

`, value)
}

func testAccCreatedestinationlistConfig_isGlobal_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createdestinationlist" "test" {
  isGlobal = %s
}

`, value)
}
