package vkcs

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccNetworkingSecGroupDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkingSecGroupDataSourceGroup,
			},
			{
				Config: testAccRenderConfig(testAccNetworkingSecGroupDataSourceBasic, map[string]string{"TestAccNetworkingSecGroupDataSourceGroup": testAccNetworkingSecGroupDataSourceGroup}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkingSecGroupDataSourceID("data.vkcs_networking_secgroup.secgroup_1"),
					resource.TestCheckResourceAttr(
						"data.vkcs_networking_secgroup.secgroup_1", "name", "secgroup_1"),
					resource.TestCheckResourceAttr(
						"data.vkcs_networking_secgroup.secgroup_1", "description", "My neutron security group"),
					resource.TestCheckResourceAttr(
						"data.vkcs_networking_secgroup.secgroup_1", "tags.#", "1"),
					resource.TestCheckResourceAttr(
						"data.vkcs_networking_secgroup.secgroup_1", "all_tags.#", "2"),
				),
			},
		},
	})
}

func TestAccNetworkingSecGroupDataSource_secGroupID(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkingSecGroupDataSourceGroup,
			},
			{
				Config: testAccRenderConfig(testAccNetworkingSecGroupDataSourceSecGroupID, map[string]string{"TestAccNetworkingSecGroupDataSourceGroup": testAccNetworkingSecGroupDataSourceGroup}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkingSecGroupDataSourceID("data.vkcs_networking_secgroup.secgroup_1"),
					resource.TestCheckResourceAttr(
						"data.vkcs_networking_secgroup.secgroup_1", "name", "secgroup_1"),
					resource.TestCheckResourceAttr(
						"data.vkcs_networking_secgroup.secgroup_1", "tags.#", "1"),
					resource.TestCheckResourceAttr(
						"data.vkcs_networking_secgroup.secgroup_1", "all_tags.#", "2"),
				),
			},
		},
	})
}

func testAccCheckNetworkingSecGroupDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find security group data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Security group data source ID not set")
		}

		return nil
	}
}

const testAccNetworkingSecGroupDataSourceGroup = `
resource "vkcs_networking_secgroup" "secgroup_1" {
  name        = "secgroup_1"
  description = "My neutron security group"
  tags = [
    "foo",
    "bar",
  ]
}
`

const testAccNetworkingSecGroupDataSourceBasic = `
{{.TestAccNetworkingSecGroupDataSourceGroup}}

data "vkcs_networking_secgroup" "secgroup_1" {
  name = vkcs_networking_secgroup.secgroup_1.name
  tags = [
    "bar",
  ]
}
`

const testAccNetworkingSecGroupDataSourceSecGroupID = `
	{{.TestAccNetworkingSecGroupDataSourceGroup}}

data "vkcs_networking_secgroup" "secgroup_1" {
  secgroup_id = vkcs_networking_secgroup.secgroup_1.id
  tags = [
    "foo",
  ]
}
`
