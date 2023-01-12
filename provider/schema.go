package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const customFieldCount = 10

func newUserSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
}

func newPermissionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"label": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func newReadOnlyPasswordSchema() map[string]*schema.Schema {
	userSchema := newUserSchema()
	permissionSchema := newPermissionSchema()

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
			Description: "Whether the password is shared externally.",
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
			Elem:        &schema.Resource{Schema: userSchema},
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
			Elem:        &schema.Resource{Schema: userSchema},
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
			Elem:        &schema.Resource{Schema: userSchema},
			Description: "User which updated the password.",
		},
		"users_permissions": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"user": {
						Type:     schema.TypeList,
						Elem:     &schema.Resource{Schema: userSchema},
						Computed: true,
					},
					"permission": {
						Type:     schema.TypeList,
						Elem:     &schema.Resource{Schema: permissionSchema},
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
						Elem:     &schema.Resource{Schema: permissionSchema},
						Computed: true,
					},
				},
			},
			Description: "This is an array of objects of the following data: group object and permission object (permission id, description). Each object describes the permission set to the group on the password. These data are only available to users with manage permission on the password (they're set to null for users that don't have the manage permission).",
		},
	}

	for i := 1; i <= customFieldCount; i++ {
		fieldName := fmt.Sprintf("custom_field_%d", i)
		description := fmt.Sprintf("Custom field %d", i)

		field := &schema.Schema{
			Type:        schema.TypeString,
			Description: description,
		}

		field.Computed = true

		passwordSchema[fieldName] = field
	}

	return passwordSchema
}

func newPasswordSchema() map[string]*schema.Schema {
	passwordSchema := newReadOnlyPasswordSchema()

	for i := 1; i <= customFieldCount; i++ {
		fieldName := fmt.Sprintf("custom_field_%d", i)
		description := fmt.Sprintf("Custom field %d", i)

		field := &schema.Schema{
			Type:        schema.TypeString,
			Description: description,
		}

		field.Optional = true

		passwordSchema[fieldName] = field
	}

	passwordSchema["id"] = &schema.Schema{
		Type:        passwordSchema["id"].Type,
		Computed:    true,
		Optional:    true,
		Description: passwordSchema["id"].Description,
	}
	passwordSchema["name"] = &schema.Schema{
		Type:        passwordSchema["name"].Type,
		Required:    true,
		Description: passwordSchema["name"].Description,
	}
	passwordSchema["project_id"] = &schema.Schema{
		Type:        passwordSchema["project_id"].Type,
		Required:    true,
		Description: passwordSchema["project_id"].Description,
	}
	passwordSchema["username"] = &schema.Schema{
		Type:        passwordSchema["username"].Type,
		Optional:    true,
		Sensitive:   true,
		Description: passwordSchema["username"].Description,
	}
	passwordSchema["email"] = &schema.Schema{
		Type:        passwordSchema["email"].Type,
		Optional:    true,
		Sensitive:   true,
		Description: passwordSchema["email"].Description,
	}
	passwordSchema["password"] = &schema.Schema{
		Type:        passwordSchema["password"].Type,
		Required:    true,
		Sensitive:   true,
		Description: passwordSchema["password"].Description,
	}
	passwordSchema["notes"] = &schema.Schema{
		Type:        passwordSchema["notes"].Type,
		Optional:    true,
		Description: passwordSchema["notes"].Description,
	}
	passwordSchema["access_info"] = &schema.Schema{
		Type:        passwordSchema["access_info"].Type,
		Optional:    true,
		Description: passwordSchema["access_info"].Description,
	}
	passwordSchema["tags"] = &schema.Schema{
		Type:        passwordSchema["tags"].Type,
		Optional:    true,
		Elem:        passwordSchema["tags"].Elem,
		Description: passwordSchema["tags"].Description,
	}
	passwordSchema["expiry_date"] = &schema.Schema{
		Type:        passwordSchema["expiry_date"].Type,
		Optional:    true,
		Description: passwordSchema["expiry_date"].Description,
	}

	return passwordSchema
}
