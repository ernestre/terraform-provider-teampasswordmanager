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

func TestAccTPMPasswordBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTPMPasswordDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
                    resource "teampasswordmanager_project" "my_project" {
                        name = "test_project"
                    }
                    resource "teampasswordmanager_password" "new" {
                        name = "new_password"
                        project_id = teampasswordmanager_project.my_project.id
                        password = "secure_password"
                        username = "secret_username"
                        email = "foo@bar.com"
                        notes = "additinal information about password"

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
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "name", "new_password"),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "password", "secure_password"),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "username", "secret_username"),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "email", "foo@bar.com"),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "notes", "additinal information about password"),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "custom_field_1", "custom data 1"),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "custom_field_2", "custom data 2"),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "custom_field_3", "custom data 3"),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "custom_field_4", "custom data 4"),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "custom_field_5", "custom data 5"),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "custom_field_6", "custom data 6"),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "custom_field_7", "custom data 7"),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "custom_field_8", "custom data 8"),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "custom_field_9", "custom data 9"),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "custom_field_10", "custom data 10"),
					testAccCheckTPMPasswordExists("teampasswordmanager_password.new", "teampasswordmanager_project.my_project"),
				),
			},
			{
				Config: `
                    resource "teampasswordmanager_project" "my_project" {
                        name = "test_project"
                    }
                    resource "teampasswordmanager_password" "new" {
                        name = "the_new_old_passwowrd"
                        project_id = teampasswordmanager_project.my_project.id
                        password = "foobar"
                    }
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "name", "the_new_old_passwowrd"),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "password", "foobar"),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "username", ""),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "email", ""),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "notes", ""),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "custom_field_1", ""),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "custom_field_2", ""),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "custom_field_3", ""),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "custom_field_4", ""),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "custom_field_5", ""),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "custom_field_6", ""),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "custom_field_7", ""),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "custom_field_8", ""),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "custom_field_9", ""),
					resource.TestCheckResourceAttr("teampasswordmanager_password.new", "custom_field_10", ""),
					testAccCheckTPMPasswordExists("teampasswordmanager_password.new", "teampasswordmanager_project.my_project"),
				),
			},
		},
	})
}

func testAccCheckTPMPasswordExists(passwordResourceName string, projectResourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		passwordResource, ok := s.RootModule().Resources[passwordResourceName]

		if !ok {
			return fmt.Errorf("Not found: %s", passwordResourceName)
		}

		if passwordResource.Primary.ID == "" {
			return fmt.Errorf("project ID is not set")
		}

		projectResource, ok := s.RootModule().Resources[projectResourceName]

		if !ok {
			return fmt.Errorf("Not found: %s", projectResourceName)
		}

		if projectResource.Primary.ID == "" {
			return fmt.Errorf("project ID is not set")
		}

		passwordProjectID, err := strconv.Atoi(passwordResource.Primary.Attributes["project_id"])
		if err != nil {
			return err
		}
		projectID, err := strconv.Atoi(projectResource.Primary.ID)
		if err != nil {
			return err
		}

		if passwordProjectID != projectID {
			return fmt.Errorf(
				"password has invalid project id assigned. Got expected %d, got %d",
				projectID,
				passwordProjectID,
			)
		}

		c := getPasswordClient(testAccProvider.Meta())

		passwordID, err := strconv.Atoi(passwordResource.Primary.ID)
		if err != nil {
			return err
		}

		password, err := c.Get(passwordID)
		if err != nil {
			return err
		}

		if password.ID != passwordID {
			return fmt.Errorf("remote password ID does not match the password id in state")
		}

		if password.Project.ID != projectID {
			return fmt.Errorf("remote password's project ID does not match the project id in state")
		}

		return nil
	}
}

func testAccCheckTPMPasswordDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(clientRegistry)[clientPassword].(tpm.PasswordClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "teampasswordmanager_password" {
			continue
		}

		passwordID, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return err
		}

		_, err = c.Get(passwordID)
		if !errors.Is(err, tpm.ErrPasswordNotFound) {
			return err
		}
	}

	return nil
}
