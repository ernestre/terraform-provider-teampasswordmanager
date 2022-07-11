package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTPMDataSourcePassword(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `
                    resource "teampasswordmanager_project" "my_project" {
                        name = "test_data_source_project"
                    }

                    resource "teampasswordmanager_password" "new" {
                        name = "new_password"
                        project_id = teampasswordmanager_project.my_project.id
                        password = "secure_password"
                    }

                    data "teampasswordmanager_password" "foo" {
                        id = teampasswordmanager_password.new.id
                    }
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "name", "new_password"),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "password", "secure_password"),
					// Data source
					resource.TestCheckResourceAttr("data.teampasswordmanager_password.foo", "name", "new_password"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_password.foo", "password", "secure_password"),
				),
			},
		},
	})
}
