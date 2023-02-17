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

func TestAccResourceGroup(t *testing.T) {
	checkAccountReleatedFields := func(s *terraform.State) error {
		r := s.RootModule().Resources["teampasswordmanager_group.new"]
		groupID, err := strconv.Atoi(r.Primary.ID)
		if err != nil {
			return err
		}

		groupClient := newTestGroupClient()
		group, err := groupClient.Get(groupID)
		if err != nil {
			return err
		}

		attrMap := map[string]any{
			"created_by.0.email_address": group.CreatedBy.Email,
			"created_on":                 group.CreatedOn.String(),
			"updated_by.0.email_address": group.UpdatedBy.Email,
			"updated_on":                 group.UpdatedOn.String(),
		}

		attr := r.Primary.Attributes
		for name, value := range attrMap {
			if attr[name] != value {
				return fmt.Errorf("attribute's %s value '%s' in state does not match value returned from API '%s'", name, attr[name], value)
			}
		}

		return nil
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResourceGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: `
                    resource "teampasswordmanager_group" "new" {
                        name = "new_group"
                    }
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("teampasswordmanager_group.new", "name", "new_group"),
					resource.TestCheckResourceAttr("teampasswordmanager_group.new", "is_ldap", "false"),
					resource.TestCheckResourceAttr("teampasswordmanager_group.new", "ldap_server_id", "0"),
					resource.TestCheckResourceAttr("teampasswordmanager_group.new", "group_dn", ""),
					resource.TestCheckResourceAttr("teampasswordmanager_group.new", "users.#", "0"),

					resource.TestCheckResourceAttr("teampasswordmanager_group.new", "created_by.0.id", "1"),
					resource.TestCheckResourceAttr("teampasswordmanager_group.new", "created_by.0.name", "admin"),
					resource.TestCheckResourceAttr("teampasswordmanager_group.new", "created_by.0.role", "Admin"),

					resource.TestCheckResourceAttr("teampasswordmanager_group.new", "updated_by.0.id", "1"),
					resource.TestCheckResourceAttr("teampasswordmanager_group.new", "updated_by.0.name", "admin"),
					resource.TestCheckResourceAttr("teampasswordmanager_group.new", "updated_by.0.role", "Admin"),

					checkAccountReleatedFields,
				),
			},
			{
				Config: `
                    resource "teampasswordmanager_group" "new" {
                        name = "new_group_updated"
                    }
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("teampasswordmanager_group.new", "name", "new_group_updated"),
					resource.TestCheckResourceAttr("teampasswordmanager_group.new", "is_ldap", "false"),
					resource.TestCheckResourceAttr("teampasswordmanager_group.new", "ldap_server_id", "0"),
					resource.TestCheckResourceAttr("teampasswordmanager_group.new", "group_dn", ""),
					resource.TestCheckResourceAttr("teampasswordmanager_group.new", "users.#", "0"),

					resource.TestCheckResourceAttr("teampasswordmanager_group.new", "created_by.0.id", "1"),
					resource.TestCheckResourceAttr("teampasswordmanager_group.new", "created_by.0.name", "admin"),
					resource.TestCheckResourceAttr("teampasswordmanager_group.new", "created_by.0.role", "Admin"),

					resource.TestCheckResourceAttr("teampasswordmanager_group.new", "updated_by.0.id", "1"),
					resource.TestCheckResourceAttr("teampasswordmanager_group.new", "updated_by.0.name", "admin"),
					resource.TestCheckResourceAttr("teampasswordmanager_group.new", "updated_by.0.role", "Admin"),

					checkAccountReleatedFields,
				),
			},
		},
	})
}

func testAccCheckResourceGroupDestroyed(s *terraform.State) error {
	c := getGroupClient(testAccProvider.Meta())

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "teampasswordmanager_group" {
			continue
		}

		groupID, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return err
		}

		_, err = c.Get(groupID)
		if !errors.Is(err, tpm.ErrProjectNotFound) {
			return err
		}
	}

	return nil
}
