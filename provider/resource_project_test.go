package provider

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/ernestre/terraform-provider-teampasswordmanager/tpm"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTPMProjectBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTPMProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
			                 resource "teampasswordmanager_project" "new" {
			                     name = "new_project"
			                     notes = "this is a small note"
			                 }
			             `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("teampasswordmanager_project.new", "name", "new_project"),
					resource.TestCheckResourceAttr("teampasswordmanager_project.new", "notes", "this is a small note"),
					resource.TestCheckNoResourceAttr("teampasswordmanager_project.new", "tags.#"),
					resource.TestCheckResourceAttr("teampasswordmanager_project.new", "parent_id", "0"),
					testAccCheckTPMProjectExists("teampasswordmanager_project.new"),
				),
			},
			{
				Config: `
			                 resource "teampasswordmanager_project" "new" {
			                     name = "updated_project"
			                     notes = "this is a smaller note"
			                     tags = ["a","b","c"]
			                 }
			             `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("teampasswordmanager_project.new", "name", "updated_project"),
					resource.TestCheckResourceAttr("teampasswordmanager_project.new", "notes", "this is a smaller note"),
					resource.TestCheckResourceAttr("teampasswordmanager_project.new", "tags.#", "3"),
					resource.TestCheckResourceAttr("teampasswordmanager_project.new", "tags.0", "a"),
					resource.TestCheckResourceAttr("teampasswordmanager_project.new", "tags.1", "b"),
					resource.TestCheckResourceAttr("teampasswordmanager_project.new", "tags.2", "c"),
					resource.TestCheckResourceAttr("teampasswordmanager_project.new", "parent_id", "0"),
					testAccCheckTPMProjectExists("teampasswordmanager_project.new"),
				),
			},
			{
				Config: `
			                 resource "teampasswordmanager_project" "new" {
			                     name = "updated_project"
			                     notes = "this is a smaller note"
			                 }
			             `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("teampasswordmanager_project.new", "name", "updated_project"),
					resource.TestCheckResourceAttr("teampasswordmanager_project.new", "notes", "this is a smaller note"),
					resource.TestCheckNoResourceAttr("teampasswordmanager_project.new", "tags.#"),
					resource.TestCheckResourceAttr("teampasswordmanager_project.new", "parent_id", "0"),
					testAccCheckTPMProjectExists("teampasswordmanager_project.new"),
				),
			},
			{
				Config: `
			                 resource "teampasswordmanager_project" "new" {
			                     name = "updated_project_again"
			                 }
			             `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("teampasswordmanager_project.new", "name", "updated_project_again"),
					resource.TestCheckResourceAttr("teampasswordmanager_project.new", "notes", ""),
					resource.TestCheckNoResourceAttr("teampasswordmanager_project.new", "tags.#"),
					resource.TestCheckResourceAttr("teampasswordmanager_project.new", "parent_id", "0"),
					testAccCheckTPMProjectExists("teampasswordmanager_project.new"),
				),
			},
		},
	})
}

func testAccCheckTPMProjectExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No project_id set")
		}

		c := getProjectClient(testAccProvider.Meta())

		projectID, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return err
		}

		projectData, err := c.Get(projectID)
		if err != nil {
			return err
		}

		if projectData.ID != projectID {
			return fmt.Errorf("Not found: %s", n)
		}

		return nil
	}
}

func testAccCheckTPMProjectDestroy(s *terraform.State) error {
	c := getProjectClient(testAccProvider.Meta())

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "teampasswordmanager_project" {
			continue
		}

		projectID, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return err
		}

		_, err = c.Get(projectID)
		if !errors.Is(err, tpm.ErrProjectNotFound) {
			return err
		}
	}

	return nil
}
