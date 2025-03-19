package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"goauthentik.io/terraform-provider-authentik/pkg/helpers"
)

func dataSourcePolicyExpression() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePolicyExpressionRead,
		Description: "Customization --- Get Policy Expression by name",
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"execution_logging": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"expression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourcePolicyExpressionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*APIClient)

	req := c.client.PoliciesApi.PoliciesExpressionList(ctx)
	if name, ok := d.GetOk("name"); ok {
		req = req.Name(name.(string))
	}

	res, hr, err := req.Execute()
	if err != nil {
		return helpers.HTTPToDiag(d, hr, err)
	}

	if len(res.Results) < 1 {
		return diag.Errorf("No matching policy expression found")
	}
	f := res.Results[0]
	d.SetId(f.Pk)
	helpers.SetWrapper(d, "name", f.Name)
	helpers.SetWrapper(d, "execution_logging", f.ExecutionLogging)
	helpers.SetWrapper(d, "expression", f.Expression)
	//	helpers.SetWrapper(d, "uuid", f.Pk)
	return diags
}
