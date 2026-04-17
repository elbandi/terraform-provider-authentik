package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"goauthentik.io/terraform-provider-authentik/pkg/helpers"
)

func dataSourcePolicyPassword() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePolicyPasswordRead,
		Description: "Customization --- Get Policy Password by name",
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"execution_logging": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"password_field": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"check_static_rules": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"check_have_i_been_pwned": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"check_zxcvbn": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"error_message": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"amount_uppercase": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"amount_lowercase": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"amount_symbols": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"amount_digits": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"length_min": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"symbol_charset": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"hibp_allowed_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"zxcvbn_score_threshold": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourcePolicyPasswordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*APIClient)

	req := c.client.PoliciesApi.PoliciesPasswordList(ctx)
	if name, ok := d.GetOk("name"); ok {
		req = req.Name(name.(string))
	}

	res, hr, err := req.Execute()
	if err != nil {
		return helpers.HTTPToDiag(d, hr, err)
	}

	if len(res.Results) < 1 {
		return diag.Errorf("No matching policy password found")
	}
	f := res.Results[0]
	d.SetId(f.Pk)
	helpers.SetWrapper(d, "name", f.Name)
	helpers.SetWrapper(d, "execution_logging", f.ExecutionLogging)
	helpers.SetWrapper(d, "password_field", f.PasswordField)
	helpers.SetWrapper(d, "check_static_rules", f.CheckStaticRules)
	helpers.SetWrapper(d, "check_have_i_been_pwned", f.CheckHaveIBeenPwned)
	helpers.SetWrapper(d, "check_zxcvbn", f.CheckZxcvbn)
	helpers.SetWrapper(d, "error_message", f.ErrorMessage)
	helpers.SetWrapper(d, "amount_uppercase", f.AmountUppercase)
	helpers.SetWrapper(d, "amount_lowercase", f.AmountLowercase)
	helpers.SetWrapper(d, "amount_symbols", f.AmountSymbols)
	helpers.SetWrapper(d, "amount_digits", f.AmountDigits)
	helpers.SetWrapper(d, "length_min", f.LengthMin)
	helpers.SetWrapper(d, "symbol_charset", f.SymbolCharset)
	helpers.SetWrapper(d, "hibp_allowed_count", f.HibpAllowedCount)
	helpers.SetWrapper(d, "zxcvbn_score_threshold", f.ZxcvbnScoreThreshold)
	//	helpers.SetWrapper(d, "uuid", f.Pk)
	return diags
}
