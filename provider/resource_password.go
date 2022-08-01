package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ernestre/terraform-provider-teampasswordmanager/tpm"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const customFieldCount = 10

func resourcePassword() *schema.Resource {
	passwordSchema := map[string]*schema.Schema{
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
		"username": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Username value.",
		},
		"password": {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			Description: "Password value.",
		},
	}

	for i := 1; i <= customFieldCount; i++ {
		passwordSchema[fmt.Sprintf("custom_field_%d", i)] = &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: fmt.Sprintf("Custom field %d", i),
		}
	}

	return &schema.Resource{
		Description:   "Creates a password resource for a given project.",
		CreateContext: resourcePasswordCreate,
		ReadContext:   resourcePasswordRead,
		UpdateContext: resourcePasswordUpdate,
		DeleteContext: resourcePasswordDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: passwordSchema,
	}
}

func resourcePasswordCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := getPasswordClient(m)

	r := tpm.CreatePasswordRequest{
		Name:         d.Get("name").(string),
		ProjectID:    d.Get("project_id").(int),
		Password:     d.Get("password").(string),
		Username:     d.Get("username").(string),
		CustomData1:  d.Get("custom_field_1").(string),
		CustomData2:  d.Get("custom_field_2").(string),
		CustomData3:  d.Get("custom_field_3").(string),
		CustomData4:  d.Get("custom_field_4").(string),
		CustomData5:  d.Get("custom_field_5").(string),
		CustomData6:  d.Get("custom_field_6").(string),
		CustomData7:  d.Get("custom_field_7").(string),
		CustomData8:  d.Get("custom_field_8").(string),
		CustomData9:  d.Get("custom_field_9").(string),
		CustomData10: d.Get("custom_field_10").(string),
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
	c := getPasswordClient(m)

	passwordID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	r := tpm.UpdatePasswordRequest{
		Name:         d.Get("name").(string),
		Password:     d.Get("password").(string),
		Username:     d.Get("username").(string),
		CustomData1:  d.Get("custom_field_1").(string),
		CustomData2:  d.Get("custom_field_2").(string),
		CustomData3:  d.Get("custom_field_3").(string),
		CustomData4:  d.Get("custom_field_4").(string),
		CustomData5:  d.Get("custom_field_5").(string),
		CustomData6:  d.Get("custom_field_6").(string),
		CustomData7:  d.Get("custom_field_7").(string),
		CustomData8:  d.Get("custom_field_8").(string),
		CustomData9:  d.Get("custom_field_9").(string),
		CustomData10: d.Get("custom_field_10").(string),
	}

	if err = c.Update(passwordID, r); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(passwordID))

	return resourcePasswordRead(ctx, d, m)
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

	setCustomField := func(customField *tpm.CustomField, fieldName string, resourceData *schema.ResourceData) error {
		if customField != nil {
			if err = resourceData.Set(fieldName, customField.Data); err != nil {
				return err
			}
		}

		return nil
	}

	if err = d.Set("name", passwordData.Name); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("project_id", passwordData.Project.ID); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("password", passwordData.Password); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("username", passwordData.Username); err != nil {
		return diag.FromErr(err)
	}

	if err = setCustomField(passwordData.CustomField1, "custom_field_1", d); err != nil {
		return diag.FromErr(err)
	}

	if err = setCustomField(passwordData.CustomField2, "custom_field_2", d); err != nil {
		return diag.FromErr(err)
	}

	if err = setCustomField(passwordData.CustomField3, "custom_field_3", d); err != nil {
		return diag.FromErr(err)
	}

	if err = setCustomField(passwordData.CustomField4, "custom_field_4", d); err != nil {
		return diag.FromErr(err)
	}

	if err = setCustomField(passwordData.CustomField5, "custom_field_5", d); err != nil {
		return diag.FromErr(err)
	}

	if err = setCustomField(passwordData.CustomField6, "custom_field_6", d); err != nil {
		return diag.FromErr(err)
	}

	if err = setCustomField(passwordData.CustomField7, "custom_field_7", d); err != nil {
		return diag.FromErr(err)
	}

	if err = setCustomField(passwordData.CustomField8, "custom_field_8", d); err != nil {
		return diag.FromErr(err)
	}

	if err = setCustomField(passwordData.CustomField9, "custom_field_9", d); err != nil {
		return diag.FromErr(err)
	}

	if err = setCustomField(passwordData.CustomField10, "custom_field_10", d); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
