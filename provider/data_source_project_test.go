package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTPMDataSourceProject(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `
                    resource "teampasswordmanager_project" "my_project" {
                        name = "project_data_source_test"
                        notes = "this is a small note"
                        tags = ["a","b","c"]
                    }

                    data "teampasswordmanager_project" "foo" {
                        id = teampasswordmanager_project.my_project.id
                    }
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.teampasswordmanager_project.foo", "name", "project_data_source_test"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_project.foo", "notes", "this is a small note"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_project.foo", "tags.#", "3"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_project.foo", "tags.0", "a"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_project.foo", "tags.1", "b"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_project.foo", "tags.2", "c"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_project.foo", "parent_id", "0"),
				),
			},
		},
	})
}
