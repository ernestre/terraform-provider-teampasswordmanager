package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGroup(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `
                    resource "teampasswordmanager_group" "my_group" {
                        name = "test-group"
                    }

                    data "teampasswordmanager_group" "group_data" {
                        id = teampasswordmanager_group.my_group.id
                    }
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.teampasswordmanager_group.group_data", "is_ldap", "false"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_group.group_data", "ldap_server_id", "0"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_group.group_data", "group_dn", ""),
					resource.TestCheckResourceAttr("data.teampasswordmanager_group.group_data", "users.#", "0"),

					resource.TestCheckResourceAttr("data.teampasswordmanager_group.group_data", "created_by.0.id", "1"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_group.group_data", "created_by.0.name", "admin"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_group.group_data", "created_by.0.role", "Admin"),

					resource.TestCheckResourceAttr("data.teampasswordmanager_group.group_data", "updated_by.0.id", "1"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_group.group_data", "updated_by.0.name", "admin"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_group.group_data", "updated_by.0.role", "Admin"),
				),
			},
		},
	})
}
