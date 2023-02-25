package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ernestre/terraform-provider-teampasswordmanager/tpm"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		Description:   "Creates a project.",
		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		UpdateContext: resourceProjectUpdate,
		DeleteContext: resourceProjectDelete,
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
				Description: "Name of the project, usually use for search.",
			},
			"parent_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Parent project ID. If the project is a 'root' project then the value should be 0, otherwise set the id of the parent project.",
			},
			"tags": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "Tags which are usually used for search. Tags should be unique and in alphabetical order.",
			},
			"notes": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Notes are used to store additional information about the project.",
			},
		},
	}
}

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := getProjectClient(m)

	projectID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if err = c.Delete(projectID); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := getProjectClient(m)

	projectID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	r := tpm.UpdateProjectRequest{
		Name:  d.Get("name").(string),
		Tags:  convertListToTags(d.Get("tags").([]interface{})),
		Notes: d.Get("notes").(string),
	}

	if err = c.Update(projectID, r); err != nil {
		return diag.FromErr(err)
	}

	return resourceProjectRead(ctx, d, m)
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := getProjectClient(m)

	projectID, err := strconv.Atoi(d.Get("id").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	projectData, err := c.Get(projectID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(projectID))

	if err = d.Set("name", projectData.Name); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("parent_id", projectData.ParentID); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("notes", projectData.Notes); err != nil {
		return diag.FromErr(err)
	}

	if len(projectData.Tags) > 0 {
		if err = d.Set("tags", projectData.Tags); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func convertListToTags(l []interface{}) tpm.Tags {
	tags := tpm.Tags{}

	for _, v := range l {
		tags = append(tags, fmt.Sprint(v))
	}

	return tags
}
