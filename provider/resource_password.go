package provider

import (
	"context"
	"strconv"

	"github.com/ernestre/terraform-provider-teampasswordmanager/tpm"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePassword() *schema.Resource {
	return &schema.Resource{
		Description:   "Creates a password resource for a given project.",
		CreateContext: resourcePasswordCreate,
		ReadContext:   resourcePasswordRead,
		UpdateContext: resourcePasswordUpdate,
		DeleteContext: resourcePasswordDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "Password ID.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the password, usually used for seaching.",
			},
			"project_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Project ID of the project where password should be created.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Password value.",
			},
		},
	}
}

func resourcePasswordCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := getPasswordClient(m)

	r := tpm.CreatePasswordRequest{
		Name:      d.Get("name").(string),
		ProjectID: d.Get("project_id").(int),
		Password:  d.Get("password").(string),
	}

	resp, err := c.Create(r)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(resp.ID))

	return diags
}

func resourcePasswordDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := getPasswordClient(m)

	passwordID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if err = c.Delete(passwordID); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourcePasswordUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := getPasswordClient(m)

	passwordID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	r := tpm.UpdatePasswordRequest{
		Name:     d.Get("name").(string),
		Password: d.Get("password").(string),
	}

	if err = c.Update(passwordID, r); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourcePasswordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	if err = d.Set("name", passwordData.Name); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("project_id", passwordData.Project.ID); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("password", passwordData.Password); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
