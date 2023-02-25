package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGroupMembership() *schema.Resource {
	return &schema.Resource{
		Description:   "Group membership. https://teampasswordmanager.com/docs/api-groups/",
		CreateContext: resourceGroupMembershipCreate,
		ReadContext:   resourceGroupMembershipRead,
		UpdateContext: resourceGroupMembershipCreate,
		DeleteContext: resourceGroupMembershipDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Group ID.",
			},
			"user_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "User ID.",
			},
			"group_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Group ID.",
			},
		},
	}
}

func resourceGroupMembershipRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := getGroupClient(m)
	membershipID := d.Get("id").(string)

	groupID, userID, err := parseGroupMemeberShipID(membershipID)
	if err != nil {
		return diag.FromErr(err)
	}

	group, err := c.Get(groupID)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, user := range group.Users {
		if user.ID == userID {
			if err = d.Set("group_id", groupID); err != nil {
				return diag.FromErr(err)
			}

			if err = d.Set("user_id", userID); err != nil {
				return diag.FromErr(err)
			}

		}
	}

	d.SetId(generateGroupMemeberShipID(groupID, userID))

	return diags
}

func resourceGroupMembershipCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := getGroupClient(m)

	userID := d.Get("user_id").(int)
	groupID := d.Get("group_id").(int)

	err := c.AddUserToGroup(groupID, userID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(generateGroupMemeberShipID(groupID, userID))

	return resourceGroupMembershipRead(ctx, d, m)
}

func resourceGroupMembershipDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := getGroupClient(m)

	userID := d.Get("user_id").(int)
	groupID := d.Get("group_id").(int)

	err := c.DeleteUserFromGroup(groupID, userID)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
