package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceProject() *schema.Resource {
	return &schema.Resource{
		Description: "Retrieve project information.",
		ReadContext: resourceProjectRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Project ID.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the project, usually use for search.",
			},
			"parent_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Parent project ID. If the project is a 'root' project then the value should be 0, otherwise set the id of the parent project.",
			},
			"tags": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Description: "Project tags, usually used for search. Tags should be unique and in alphabetical order.",
			},
			"notes": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Notes are used to store additional information about the project.",
			},
		},
	}
}
