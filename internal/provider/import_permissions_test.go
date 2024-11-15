package provider_test

import (
	"testing"

	"github.com/bayesgg/terraform-provider-rabbitmq/internal/acceptance"
	rabbithole "github.com/michaelklishin/rabbit-hole/v2"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccPermissions_importBasic(t *testing.T) {
	resourceName := "rabbitmq_permissions.test"
	var permissionInfo rabbithole.PermissionInfo

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.TestAcc.PreCheck(t) },
		Providers:    acceptance.TestAcc.Providers,
		CheckDestroy: testAccPermissionsCheckDestroy(&permissionInfo),
		Steps: []resource.TestStep{
			{
				Config: testAccPermissionsConfig_basic,
				Check: testAccPermissionsCheck(
					resourceName, &permissionInfo,
				),
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
