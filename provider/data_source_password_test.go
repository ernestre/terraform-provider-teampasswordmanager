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

                        custom_field_1 = "custom data 1"
                        custom_field_2 = "custom data 2"
                        custom_field_3 = "custom data 3"
                        custom_field_4 = "custom data 4"
                        custom_field_5 = "custom data 5"
                        custom_field_6 = "custom data 6"
                        custom_field_7 = "custom data 7"
                        custom_field_8 = "custom data 8"
                        custom_field_9 = "custom data 9"
                        custom_field_10 = "custom data 10"
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
					resource.TestCheckResourceAttr("data.teampasswordmanager_password.foo", "custom_field_1", "custom data 1"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_password.foo", "custom_field_2", "custom data 2"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_password.foo", "custom_field_3", "custom data 3"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_password.foo", "custom_field_4", "custom data 4"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_password.foo", "custom_field_5", "custom data 5"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_password.foo", "custom_field_6", "custom data 6"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_password.foo", "custom_field_7", "custom data 7"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_password.foo", "custom_field_8", "custom data 8"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_password.foo", "custom_field_9", "custom data 9"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_password.foo", "custom_field_10", "custom data 10"),
				),
			},
		},
	})
}
