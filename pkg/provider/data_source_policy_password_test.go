package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourcePolicyPassword(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourcePolicyPasswordSimple,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.authentik_policy_password.inbuilt", "name", "default-password-change-password-policy"),
					resource.TestCheckResourceAttrSet("data.authentik_policy_password.inbuilt", "password_field"),
				),
			},
		},
	})
}

const testAccDataSourcePolicyPasswordSimple = `
data "authentik_policy_password" "inbuilt" {
  name = "default-password-change-password-policy"
}
`
