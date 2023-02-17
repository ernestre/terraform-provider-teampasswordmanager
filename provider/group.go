package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func newResourceGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Project ID.",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the group.",
		},
		"is_ldap": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the group is a LDAP group.",
		},
		"ldap_server_id": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "LDAP server id",
		},
		"group_dn": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "LDAP group's distinguished name (DN)",
		},
		"users": {
			Type: schema.TypeList,
			Elem: &schema.Resource{
				Schema: newUserSchema(),
			},
			Computed:    true,
			Description: "Users of the group.",
		},
		"created_on": newCreatedOnSchema(),
		"created_by": newCreatedBySchema(),
		"updated_on": newUpdatedOnSchema(),
		"updated_by": newUpdatedBySchema(),
	}
}

func newDataSourceGroupSchema() map[string]*schema.Schema {
	groupSchema := newResourceGroupSchema()

	groupSchema["id"] = &schema.Schema{
		Type:        groupSchema["id"].Type,
		Required:    true,
		Description: groupSchema["id"].Description,
	}

	groupSchema["name"] = &schema.Schema{
		Type:        groupSchema["name"].Type,
		Computed:    true,
		Description: groupSchema["name"].Description,
	}

	return groupSchema
}
