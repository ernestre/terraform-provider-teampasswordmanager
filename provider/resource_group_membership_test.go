package provider

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceGroupMembership(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `
                    resource "teampasswordmanager_group" "new" {
                        name = "new_group"
                    }
                    resource "teampasswordmanager_group_membership" "new" {
                        group_id = teampasswordmanager_group.new.id
                        user_id = 1
                    }
                `,
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						groupResource := s.RootModule().Resources["teampasswordmanager_group.new"]
						groupID, err := strconv.Atoi(groupResource.Primary.ID)
						if err != nil {
							return err
						}

						groupMembershipResource := s.RootModule().Resources["teampasswordmanager_group_membership.new"]
						groupID, userID, err := parseGroupMemeberShipID(groupMembershipResource.Primary.ID)
						if err != nil {
							return err
						}

						groupClient := newTestGroupClient()
						group, err := groupClient.Get(groupID)
						if err != nil {
							return err
						}

						if groupMembershipResource.Primary.Attributes["group_id"] != fmt.Sprint(group.ID) {
							return fmt.Errorf(
								"group id(%s) in state does not match the group id(%s) returned from API",
								groupMembershipResource.Primary.Attributes["group_id"],
								fmt.Sprint(group.ID),
							)
						}

						for _, user := range group.Users {
							if user.ID == userID {
								if groupMembershipResource.Primary.Attributes["user_id"] != fmt.Sprint(user.ID) {
									return fmt.Errorf(
										"user id(%s) in state does not match the user id(%s) returned from API",
										groupMembershipResource.Primary.Attributes["user_id"],
										fmt.Sprint(group.ID),
									)
								}

								return nil
							}
						}

						return fmt.Errorf("user(%d) not found in group (%d)", userID, groupID)
					},
				),
			},
		},
	})
}
