package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Retrieve group information.",
		ReadContext: resourceGroupRead,
		Schema:      newDataSourceGroupSchema(),
	}
}
