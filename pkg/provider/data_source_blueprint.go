package provider

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"goauthentik.io/terraform-provider-authentik/pkg/helpers"
)

func dataSourceBlueprint() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBlueprintRead,
		Description: "Blueprints ---",
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Find blueprint by name",
			},
			"path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Find blueprint by path",
			},
			"content": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"context": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_applied": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceBlueprintRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*APIClient)

	req := c.client.ManagedApi.ManagedBlueprintsList(ctx)
	if s, ok := d.GetOk("name"); ok {
		req = req.Name(s.(string))
	}
	if s, ok := d.GetOk("path"); ok {
		req = req.Path(s.(string))
	}

	res, hr, err := req.Execute()
	if err != nil {
		return helpers.HTTPToDiag(d, hr, err)
	}

	if len(res.Results) < 1 {
		return diag.Errorf("No matching blueprint found")
	}
	f := res.Results[0]
	d.SetId(f.Pk)
	helpers.SetWrapper(d, "name", f.Name)
	helpers.SetWrapper(d, "path", f.Path)
	if f.Content != nil {
		helpers.SetWrapper(d, "content", *f.Content)
	}
	helpers.SetWrapper(d, "enabled", f.Enabled)
	b, err := json.Marshal(f.Context)
	if err != nil {
		return diag.FromErr(err)
	}
	helpers.SetWrapper(d, "context", string(b))
	helpers.SetWrapper(d, "last_applied", f.LastApplied.String())
	helpers.SetWrapper(d, "status", f.Status)
	return diags
}
