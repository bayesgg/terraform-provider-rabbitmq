package provider_test

import (
	"testing"

	"github.com/bayesgg/terraform-provider-rabbitmq/internal/acceptance"
	rabbithole "github.com/michaelklishin/rabbit-hole/v2"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccShovel_importBasic(t *testing.T) {

	resourceName := "rabbitmq_shovel.shovelTest"
	var shovel rabbithole.ShovelInfo

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.TestAcc.PreCheck(t) },
		Providers:    acceptance.TestAcc.Providers,
		CheckDestroy: testAccShovelCheckDestroy(&shovel),
		Steps: []resource.TestStep{
			{
				Config: testAccShovelConfig_basic,
				Check: testAccShovelCheck(
					resourceName, &shovel,
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
