package azurerm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceAzureRMApiManagement_basic(t *testing.T) {
	dataSourceName := "data.azurerm_api_management.test"
	rInt := acctest.RandInt()
	location := testLocation()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceApiManagement_basic(rInt, location),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "publisher_email", "pub1@email.com"),
					resource.TestCheckResourceAttr(dataSourceName, "publisher_name", "pub1"),
					resource.TestCheckResourceAttr(dataSourceName, "sku.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "sku.0.capacity", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "sku.0.name", "Developer"),
					resource.TestCheckResourceAttr(dataSourceName, "tags.%", "0"),
					resource.TestCheckResourceAttrSet(dataSourceName, "public_ip_addresses.#"),
				),
			},
		},
	})
}

func testAccDataSourceApiManagement_basic(rInt int, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "amtestRG-%d"
  location = "%s"
}

resource "azurerm_api_management" "test" {
  name            = "acctestAM-%d"
  publisher_name  = "pub1"
  publisher_email = "pub1@email.com"

  sku {
    name     = "Developer"
    capacity = 1
  }

  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
}

data "azurerm_api_management" "test" {
  name                = "${azurerm_api_management.test.name}"
  resource_group_name = "${azurerm_api_management.test.resource_group_name}"
}
`, rInt, location, rInt)
}
