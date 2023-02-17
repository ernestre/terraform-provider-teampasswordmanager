package provider

import (
	"context"
	"strconv"
	"time"

	"github.com/ernestre/terraform-provider-teampasswordmanager/tpm"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePassword() *schema.Resource {
	return &schema.Resource{
		Description: "Retrieve password information resource for a given project.",
		ReadContext: dataSourcePasswordRead,
		Schema:      newReadOnlyPasswordSchema(),
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
