package vkcs

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccNetworkingFloatingIPV2DataSource_address(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccRenderConfig(testAccNetworkingFloatingIPV2DataSourceFloatingIP),
			},
			{
				Config: testAccRenderConfig(testAccNetworkingFloatingIPV2DataSourceAddress, map[string]string{"TestAccNetworkingFloatingIPV2DataSourceFloatingIP": testAccRenderConfig(testAccNetworkingFloatingIPV2DataSourceFloatingIP)}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkingFloatingIPV2DataSourceID("data.vkcs_networking_floatingip.fip_1"),
					resource.TestCheckResourceAttrSet(
						"data.vkcs_networking_floatingip.fip_1", "address"),
					resource.TestCheckResourceAttrSet(
						"data.vkcs_networking_floatingip.fip_1", "pool"),
					resource.TestCheckResourceAttrSet(
						"data.vkcs_networking_floatingip.fip_1", "status"),
					resource.TestCheckResourceAttrSet(
						"data.vkcs_networking_floatingip.fip_1", "description"),
				),
			},
		},
	})
}

func testAccCheckNetworkingFloatingIPV2DataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find floating IP data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Floating IP data source ID not set")
		}

		return nil
	}
}

const testAccNetworkingFloatingIPV2DataSourceFloatingIP = `
resource "vkcs_networking_floatingip" "fip_1" {
  pool = "{{.ExtNetName}}"
  description = "test fip"
}
`

const testAccNetworkingFloatingIPV2DataSourceAddress = `
{{.TestAccNetworkingFloatingIPV2DataSourceFloatingIP}}

data "vkcs_networking_floatingip" "fip_1" {
  address = vkcs_networking_floatingip.fip_1.address
  description = "test fip"
}
`
