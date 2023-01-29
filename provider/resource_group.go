package provider

import (
	"context"
	"strconv"

	"github.com/ernestre/terraform-provider-teampasswordmanager/tpm"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Creates a group.",
		// CreateContext: resourceGroupCreate,
		ReadContext: resourceGroupRead,
		// UpdateContext: resourceProjectUpdate,
		// DeleteContext: resourceProjectDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "Project ID.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the group.",
			},
			"num_users": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of users in a group.",
			},
			"is_ldap": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the group is a ldap group.",
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
		},
	}
}

func resourceGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := getProjectClient(m)

	r := tpm.CreateProjectRequest{
		Name: d.Get("name").(string),
	}

	tags := d.Get("tags").([]interface{})

	if len(tags) > 0 {
		r.Tags = convertListToTags(tags)
	}

	if notes := d.Get("notes"); notes != nil {
		r.Notes = notes.(string)
	}

	if parentID := d.Get("parent_id"); parentID != nil {
		r.ParentID = parentID.(int)
	}

	resp, err := c.Create(r)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(resp.ID))

	return diags
}
