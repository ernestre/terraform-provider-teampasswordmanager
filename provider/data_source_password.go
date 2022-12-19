package provider

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/ernestre/terraform-provider-teampasswordmanager/tpm"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePassword() *schema.Resource {
	userShema := map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"username": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"email_address": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"role": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}

	permissionShema := map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"label": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}

	passwordSchema := map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Password ID.",
		},
		"name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Name of the password, usually used for seaching.",
		},
		"project_id": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Project ID of the project where password should be created.",
		},
		"tags": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Computed:    true,
			Description: "Tags which are usually used for search. Tags should be unique and in alphabetical order.",
		},
		"access_info": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Access information. Examples: http://site, ftp://ip-address, manual login.",
		},
		"username": {
			Type:        schema.TypeString,
			Computed:    true,
			Sensitive:   true,
			Description: "Username value.",
		},
		"email": {
			Type:        schema.TypeString,
			Computed:    true,
			Sensitive:   true,
			Description: "Email value.",
		},
		"password": {
			Type:        schema.TypeString,
			Computed:    true,
			Sensitive:   true,
			Description: "Password value.",
		},
		"expiry_date": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Expiry date of the password.",
		},
		"expiry_status": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Expiry status of the password. 0=no date or not expired, 1=expires today, 2=expired, 3=will expire soon",
		},
		"notes": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Notes are used to store additional information about the password.",
		},
		"parents": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
			Computed:    true,
			Description: "List of project ids from the root to the project of the password (in descending order), as seen by the use.",
		},
		"user_permission": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:     schema.TypeInt,
						Computed: true,
					},
					"label": {
						Type:     schema.TypeString,
						Computed: true,
					},
				},
			},
			Description: "Permission object (permission id, description) that indicates what permission has the user making the request on the password.",
		},
		"archived": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the password is archived and/or the project is archived.",
		},
		"project_archived": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the project is archived.",
		},
		"favorite": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the password is marked as favorite.",
		},
		"num_files": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Number of files.",
		},
		"locked": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the password is locked.",
		},
		"locking_type": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Locking type has the following values: 0=password not locked, 1=requires a reason to unlock, 2=requires permission to unlock.",
		},
		"locking_request_notify": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Locking request notification has the following values: 0=password not locked, 1=notify/request the password manager, 2=notify/request all the users with manage permission.",
		},
		"external_sharing": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the password is shared externaly.",
		},
		"external_url": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "External url of the password if it's shared externally.",
		},
		"linked": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the password is linked password.",
		},
		"source_password_id": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "If the password is linked, then this will be the ID of the source password.",
		},
		"managed_by": {
			Type:        schema.TypeList,
			Computed:    true,
			Elem:        &schema.Resource{Schema: userShema},
			Description: "Main manager of the password.",
		},
		"created_on": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Datetime when the password was created.",
		},
		"created_by": {
			Type:        schema.TypeList,
			Computed:    true,
			Elem:        &schema.Resource{Schema: userShema},
			Description: "User which created the password.",
		},
		"updated_on": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Datetime when the password was updated.",
		},
		"updated_by": {
			Type:        schema.TypeList,
			Computed:    true,
			Elem:        &schema.Resource{Schema: userShema},
			Description: "User which updated the password.",
		},
		"users_permissions": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"user": {
						Type:     schema.TypeList,
						Elem:     &schema.Resource{Schema: userShema},
						Computed: true,
					},
					"permission": {
						Type:     schema.TypeList,
						Elem:     &schema.Resource{Schema: permissionShema},
						Computed: true,
					},
				},
			},
			Description: "This is an array of objects of the following data: user object and permission object (permission id, description). Each object describes the permission set to the user on the password. These data are only available to users with manage permission on the password (they're set to null for users that don't have the manage permission).",
		},
		"groups_permissions": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"group": {
						Type: schema.TypeList,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"id": {
								Type:     schema.TypeInt,
								Computed: true,
							},
							"name": {
								Type:     schema.TypeString,
								Computed: true,
							},
						}},
						Computed: true,
					},
					"permission": {
						Type:     schema.TypeList,
						Elem:     &schema.Resource{Schema: permissionShema},
						Computed: true,
					},
				},
			},
			Description: "This is an array of objects of the following data: group object and permission object (permission id, description). Each object describes the permission set to the group on the password. These data are only available to users with manage permission on the password (they're set to null for users that don't have the manage permission).",
		},
	}

	for i := 1; i <= customFieldCount; i++ {
		passwordSchema[fmt.Sprintf("custom_field_%d", i)] = &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: fmt.Sprintf("Custom field %d", i),
		}
	}

	return &schema.Resource{
		Description: "Retrieve password information resource for a given project.",
		ReadContext: dataSourcePasswordRead,
		Schema:      passwordSchema,
	}
}

func dataSourcePasswordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := getPasswordClient(m)

	passwordID, err := strconv.Atoi(d.Get("id").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	passwordData, err := c.Get(passwordID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(passwordID))

	setCustomField := func(customField *tpm.CustomField, fieldName string, resourceData *schema.ResourceData) error {
		if customField != nil {
			if err = resourceData.Set(fieldName, customField.Data); err != nil {
				return err
			}
		}

		return nil
	}

	if len(passwordData.Tags) > 0 {
		if err = d.Set("tags", passwordData.Tags); err != nil {
			return diag.FromErr(err)
		}
	}

	expireDate := time.Time(passwordData.ExpiryDate)
	if !expireDate.IsZero() {
		if err = d.Set("expiry_date", expireDate.Format(tpm.ShortDateTimeFormat)); err != nil {
			return diag.FromErr(err)
		}
	}

	fields := map[string]interface{}{
		"name":          passwordData.Name,
		"project_id":    passwordData.Project.ID,
		"access_info":   passwordData.AccessInfo,
		"username":      passwordData.Username,
		"email":         passwordData.Email,
		"password":      passwordData.Password,
		"expiry_status": passwordData.ExpiryStatus,
		"notes":         passwordData.Notes,
		"parents":       passwordData.Parents,
		"user_permission": []map[string]interface{}{
			flattenPermission(passwordData.UserPermission),
		},
		"archived":               passwordData.Archived,
		"project_archived":       passwordData.ProjectArchived,
		"favorite":               passwordData.Favorite,
		"num_files":              passwordData.NumberOfFiles,
		"locked":                 passwordData.Locked,
		"locking_type":           passwordData.LockingType,
		"locking_request_notify": passwordData.LockingRequestNotify,
		"external_sharing":       passwordData.ExternalSharing,
		"external_url":           passwordData.ExternalURL,
		"linked":                 passwordData.Linked,
		"source_password_id":     passwordData.SourcePasswordID,
		"managed_by": []map[string]interface{}{
			flattenUser(passwordData.ManagedBy),
		},
		"created_on": passwordData.CreatedOn.String(),
		"created_by": []map[string]interface{}{
			flattenUser(passwordData.UpdatedBy),
		},
		"updated_on": passwordData.UpdatedOn.String(),
		"updated_by": []map[string]interface{}{
			flattenUser(passwordData.UpdatedBy),
		},
		"users_permissions":  flattenUsersPermissions(passwordData.UsersPermissions),
		"groups_permissions": flattenGroupsPermissions(passwordData.GroupsPermissions),
	}

	for field, value := range fields {
		if err = d.Set(field, value); err != nil {
			return diag.FromErr(err)
		}
	}

	customFields := map[*tpm.CustomField]string{
		passwordData.CustomField1:  "custom_field_1",
		passwordData.CustomField2:  "custom_field_2",
		passwordData.CustomField3:  "custom_field_3",
		passwordData.CustomField4:  "custom_field_4",
		passwordData.CustomField5:  "custom_field_5",
		passwordData.CustomField6:  "custom_field_6",
		passwordData.CustomField7:  "custom_field_7",
		passwordData.CustomField8:  "custom_field_8",
		passwordData.CustomField9:  "custom_field_9",
		passwordData.CustomField10: "custom_field_10",
	}

	for field, name := range customFields {
		if err = setCustomField(field, name, d); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func flattenUser(u tpm.User) map[string]interface{} {
	return map[string]interface{}{
		"id":            u.ID,
		"username":      u.Username,
		"email_address": u.Email,
		"name":          u.Name,
		"role":          u.Role,
	}
}

func flattenPermission(up tpm.Permission) map[string]interface{} {
	return map[string]interface{}{
		"id":    up.ID,
		"label": up.Label,
	}
}

func flattenUsersPermissions(up []tpm.UserPermission) []map[string]interface{} {
	userPermissions := []map[string]interface{}{}
	for _, userPermission := range up {
		user := userPermission.User
		permission := userPermission.Permission

		up := map[string]interface{}{
			"user": []map[string]interface{}{{
				"id":            user.ID,
				"username":      user.Username,
				"email_address": user.Email,
				"name":          user.Name,
				"role":          user.Role,
			}},
			"permission": []map[string]interface{}{{
				"id":    permission.ID,
				"label": permission.Label,
			}},
		}

		userPermissions = append(userPermissions, up)
	}

	return userPermissions
}

func flattenGroupsPermissions(gp []tpm.GroupPermission) []map[string]interface{} {
	groupPermissions := []map[string]interface{}{}
	for _, groupPermission := range gp {
		group := groupPermission.Group
		permission := groupPermission.Permission

		up := map[string]interface{}{
			"group": []map[string]interface{}{{
				"id":   group.ID,
				"name": group.Name,
			}},
			"permission": []map[string]interface{}{{
				"id":    permission.ID,
				"label": permission.Label,
			}},
		}

		groupPermissions = append(groupPermissions, up)
	}

	return groupPermissions
}
