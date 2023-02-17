package provider

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/ernestre/terraform-provider-teampasswordmanager/tpm"
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
                        username = "secret_username"
                        email = "foo@example.com"
                        notes = "additinal information about password"
                        access_info = "ftp://ip-address"
                        tags = ["a","b","c"]
                        expiry_date = "2022-11-26"

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
					resource.TestCheckResourceAttr("data.teampasswordmanager_password.foo", "username", "secret_username"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_password.foo", "email", "foo@example.com"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_password.foo", "notes", "additinal information about password"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_password.foo", "access_info", "ftp://ip-address"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_password.foo", "tags.#", "3"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_password.foo", "tags.0", "a"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_password.foo", "tags.1", "b"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_password.foo", "tags.2", "c"),
					resource.TestCheckResourceAttr("data.teampasswordmanager_password.foo", "expiry_date", "2022-11-26"),
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

func TestAccTPMDataSourcePasswordAddionalFields(t *testing.T) {
	projectClient := newTestProjectClient()
	passwordClient := newTestPasswordClient()
	userClient := newTestUserClient()
	groupClient := newTestGroupClient()

	project, err := projectClient.Create(tpm.CreateProjectRequest{
		Name: "foo_proj",
	})
	if err != nil {
		t.Fatal(err)
	}

	defer projectClient.Delete(project.ID)

	user, err := userClient.Create(tpm.CreateUserRequest{
		Username:     "test-user",
		EmailAddress: "test@example.com",
		Name:         "test-user",
		Role:         tpm.UserRoleProjectManager,
		Password:     "jahsah4bei0F",
	})
	if err != nil {
		t.Fatal(err)
	}

	defer userClient.Delete(user.ID)

	group, err := groupClient.Create(tpm.CreateGroupRequest{
		Name: "test-group",
	})
	if err != nil {
		t.Fatal(err)
	}

	defer groupClient.Delete(group.ID)

	req := tpm.CreatePasswordRequest{
		Name:        fmt.Sprintf("pass_%d", time.Now().UnixMicro()),
		ProjectID:   project.ID,
		Password:    "top-secret",
		Username:    "foobar",
		Email:       "foo@bar.io",
		Tags:        []string{"hello", "world"},
		AccessInfo:  "https://ftp.example.com",
		ExpiryDate:  tpm.ShortDate(time.Now().Add(time.Hour * 24 * 2)),
		Notes:       "some notes",
		CustomData1: "1",
		CustomData2: "2",
		CustomData3: "3",
		CustomData4: "4",
	}

	password, err := passwordClient.Create(req)
	if err != nil {
		t.Fatal(err)
	}

	defer passwordClient.Delete(password.ID)

	err = passwordClient.UpdatePasswordSecurity(
		password.ID,
		tpm.UpdatePasswordSecurityRequest{
			UsersPermissions: []tpm.PasswordPermission{
				tpm.NewPasswordPermission(user.ID, tpm.PasswordAccessManage),
			},
			GroupsPermissions: []tpm.PasswordPermission{
				tpm.NewPasswordPermission(group.ID, tpm.PasswordAccessEdit),
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	passwordFromRemote, err := passwordClient.Get(password.ID)
	if err != nil {
		t.Fatal(err)
	}

	checkDataSourceField := func(field string, expectedValue string) resource.TestCheckFunc {
		return resource.TestCheckResourceAttr("data.teampasswordmanager_password.foo", field, expectedValue)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
                    data "teampasswordmanager_password" "foo" {
                        id = %d
                    }
                `, password.ID),
				Check: resource.ComposeTestCheckFunc(
					checkDataSourceField("id", strconv.Itoa(passwordFromRemote.ID)),
					checkDataSourceField("name", passwordFromRemote.Name),
					checkDataSourceField("project_id", strconv.Itoa(passwordFromRemote.Project.ID)),
					checkDataSourceField("access_info", passwordFromRemote.AccessInfo),
					checkDataSourceField("username", passwordFromRemote.Username),
					checkDataSourceField("email", passwordFromRemote.Email),
					checkDataSourceField("password", passwordFromRemote.Password),
					checkDataSourceField("expiry_date", passwordFromRemote.ExpiryDate.String()),
					checkDataSourceField("expiry_status", strconv.Itoa(int(tpm.ExpiresSoon))),
					checkDataSourceField("notes", passwordFromRemote.Notes),
					checkDataSourceField("custom_field_1", "1"),
					checkDataSourceField("custom_field_2", "2"),
					checkDataSourceField("custom_field_3", "3"),
					checkDataSourceField("custom_field_4", "4"),
					checkDataSourceField("parents.#", "1"),
					checkDataSourceField("parents.0", strconv.Itoa(project.ID)),
					checkDataSourceField("user_permission.#", "1"),
					checkDataSourceField("user_permission.0.id", strconv.Itoa(passwordFromRemote.UserPermission.ID)),
					checkDataSourceField("user_permission.0.label", passwordFromRemote.UserPermission.Label),
					// TODO: Create a separate test for archiving
					checkDataSourceField("archived", "false"),
					checkDataSourceField("project_archived", "false"),
					// TODO: Create a separate test to check favorite
					checkDataSourceField("favorite", "false"),
					checkDataSourceField("num_files", "0"),
					// TODO: Create a separate test Locking
					checkDataSourceField("locked", "false"),
					checkDataSourceField("locking_type", strconv.Itoa(int(tpm.NotLocked))),
					checkDataSourceField("locking_request_notify", strconv.Itoa(int(tpm.PasswordNotLocked))),
					checkDataSourceField("external_sharing", "false"),
					checkDataSourceField("external_url", ""),
					// TODO: Test linked passwords
					checkDataSourceField("linked", "false"),
					checkDataSourceField("source_password_id", "0"),
					checkDataSourceField("created_on", passwordFromRemote.CreatedOn.String()),
					checkDataSourceField("updated_on", passwordFromRemote.UpdatedOn.String()),

					checkDataSourceField("created_by.0.id", strconv.Itoa(passwordFromRemote.CreatedBy.ID)),
					checkDataSourceField("created_by.0.username", passwordFromRemote.CreatedBy.Username),
					checkDataSourceField("created_by.0.email_address", passwordFromRemote.CreatedBy.Email),
					checkDataSourceField("created_by.0.name", passwordFromRemote.CreatedBy.Name),
					checkDataSourceField("created_by.0.role", passwordFromRemote.CreatedBy.Role),

					checkDataSourceField("updated_by.0.id", strconv.Itoa(passwordFromRemote.UpdatedBy.ID)),
					checkDataSourceField("updated_by.0.username", passwordFromRemote.UpdatedBy.Username),
					checkDataSourceField("updated_by.0.email_address", passwordFromRemote.UpdatedBy.Email),
					checkDataSourceField("updated_by.0.name", passwordFromRemote.UpdatedBy.Name),
					checkDataSourceField("updated_by.0.role", passwordFromRemote.UpdatedBy.Role),
					// Tags
					checkDataSourceField("tags.#", "2"),
					checkDataSourceField("tags.0", "hello"),
					checkDataSourceField("tags.1", "world"),
					// Managed by
					checkDataSourceField("managed_by.0.id", strconv.Itoa(passwordFromRemote.ManagedBy.ID)),
					checkDataSourceField("managed_by.0.username", passwordFromRemote.ManagedBy.Username),
					checkDataSourceField("managed_by.0.email_address", passwordFromRemote.ManagedBy.Email),
					checkDataSourceField("managed_by.0.name", passwordFromRemote.ManagedBy.Name),
					checkDataSourceField("managed_by.0.role", passwordFromRemote.ManagedBy.Role),
					// Users permissions
					checkDataSourceField("users_permissions.0.user.0.id", strconv.Itoa(passwordFromRemote.UsersPermissions[0].User.ID)),
					checkDataSourceField("users_permissions.0.user.0.username", passwordFromRemote.UsersPermissions[0].User.Username),
					checkDataSourceField("users_permissions.0.user.0.email_address", passwordFromRemote.UsersPermissions[0].User.Email),
					checkDataSourceField("users_permissions.0.user.0.name", passwordFromRemote.UsersPermissions[0].User.Name),
					checkDataSourceField("users_permissions.0.user.0.role", passwordFromRemote.UsersPermissions[0].User.Role),
					checkDataSourceField("users_permissions.0.permission.0.id", strconv.Itoa(int(tpm.PasswordAccessManage))),
					// Groups permissions
					checkDataSourceField("groups_permissions.0.group.0.id", strconv.Itoa(passwordFromRemote.GroupsPermissions[0].Group.ID)),
					checkDataSourceField("groups_permissions.0.group.0.name", passwordFromRemote.GroupsPermissions[0].Group.Name),
					checkDataSourceField("groups_permissions.0.permission.0.id", strconv.Itoa(int(tpm.PasswordAccessEdit))),
				),
			},
		},
	})
}
