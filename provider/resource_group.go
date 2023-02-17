package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ernestre/terraform-provider-teampasswordmanager/tpm"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "Creates a group. https://teampasswordmanager.com/docs/api-groups/",
		CreateContext: resourceGroupCreate,
		ReadContext:   resourceGroupRead,
		UpdateContext: resourceGroupUpdate,
		DeleteContext: resourceGroupDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
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
		},
	}
}

func resourceGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := getGroupClient(m)

	r := tpm.CreateGroupRequest{
		Name: d.Get("name").(string),
	}

	resp, err := c.Create(r)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(resp.ID))

	return resourceGroupRead(ctx, d, m)
}

func resourceGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	fmt.Println("update")

	c := getGroupClient(m)

	r := tpm.UpdateGroupRequest{
		Name: d.Get("name").(string),
	}

	groupID, err := strconv.Atoi(d.Id())

	err = c.Update(groupID, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceGroupRead(ctx, d, m)
}

func resourceGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := getGroupClient(m)

	groupID, err := strconv.Atoi(d.Id())

	err = c.Delete(groupID)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	fmt.Println("read")
	var diags diag.Diagnostics
	c := getGroupClient(m)

	groupID, err := strconv.Atoi(d.Get("id").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	groupData, err := c.Get(groupID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(groupID))

	fields := map[string]any{
		"name":           groupData.Name,
		"is_ldap":        groupData.IsLdap,
		"ldap_server_id": groupData.LdapServerID,
		"group_dn":       groupData.GroupDn,
		"users":          flattenUsers(groupData.Users),
		"created_on":     groupData.CreatedOn.String(),
		"created_by":     []map[string]any{flattenUser(groupData.UpdatedBy)},
		"updated_on":     groupData.UpdatedOn.String(),
		"updated_by":     []map[string]any{flattenUser(groupData.UpdatedBy)},
	}

	for field, value := range fields {
		if err = d.Set(field, value); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}
