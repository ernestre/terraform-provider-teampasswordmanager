package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePassword() *schema.Resource {
	return &schema.Resource{
		Description: "Retrieve password information resource for a given project.",
		ReadContext: resourcePasswordRead,
		Schema: map[string]*schema.Schema{
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
			"password": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Password value.",
			},
		},
	}
}
