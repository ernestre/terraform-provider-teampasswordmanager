package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePassword() *schema.Resource {
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
		"notes": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Notes are used to store additional information about the password.",
		},
		"access_info": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Access information. Examples: http://site, ftp://ip-address, manual login.",
		},
		"tags": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional:    true,
			Description: "Tags which are usually used for search. Tags should be unique and in alphabetical order.",
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
		ReadContext: resourcePasswordRead,
		Schema:      passwordSchema,
	}
}
