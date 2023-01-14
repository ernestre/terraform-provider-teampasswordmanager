package provider

import (
	"context"
	"strconv"
	"time"

	"github.com/ernestre/terraform-provider-teampasswordmanager/tpm"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePassword() *schema.Resource {
	return &schema.Resource{
		Description:   "Creates a password resource for a given project.",
		CreateContext: resourcePasswordCreate,
		ReadContext:   dataSourcePasswordRead,
		UpdateContext: resourcePasswordUpdate,
		DeleteContext: resourcePasswordDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: newPasswordSchema(),
	}
}

func resourcePasswordCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := getPasswordClient(m)

	r := tpm.CreatePasswordRequest{
		Name:         d.Get("name").(string),
		ProjectID:    d.Get("project_id").(int),
		Password:     d.Get("password").(string),
		Username:     d.Get("username").(string),
		Email:        d.Get("email").(string),
		Notes:        d.Get("notes").(string),
		AccessInfo:   d.Get("access_info").(string),
		Tags:         convertListToTags(d.Get("tags").([]interface{})),
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

	expireDate := d.Get("expiry_date").(string)
	if expireDate != "" {
		parsedExpireDate, err := time.Parse(tpm.ShortDateTimeFormat, expireDate)

		if err != nil {
			return diag.FromErr(ErrInvalidExpiryDateFormat)
		}

		r.ExpiryDate = tpm.ShortDate(parsedExpireDate)
	}

	resp, err := c.Create(r)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(resp.ID))

	return dataSourcePasswordRead(ctx, d, m)
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
		Email:        d.Get("email").(string),
		Notes:        d.Get("notes").(string),
		AccessInfo:   d.Get("access_info").(string),
		Tags:         convertListToTags(d.Get("tags").([]interface{})),
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

	expireDate := d.Get("expiry_date").(string)
	if expireDate != "" {
		parsedExpireDate, err := time.Parse(tpm.ShortDateTimeFormat, expireDate)

		if err != nil {
			return diag.FromErr(ErrInvalidExpiryDateFormat)
		}

		r.ExpiryDate = tpm.ShortDate(parsedExpireDate)
	}

	if err = c.Update(passwordID, r); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(passwordID))

	return dataSourcePasswordRead(ctx, d, m)
}
