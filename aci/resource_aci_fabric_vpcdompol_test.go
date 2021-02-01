package aci

import (
	"fmt"
	"testing"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccAciFabricVpcDomainPolicy_Basic(t *testing.T) {
	var fabric_vpc_domain_policy models.FabricVpcDomainPolicy
	description := "fabric_vpc_domain_policy created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciFabricVpcDomainPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciFabricVpcDomainPolicyConfig_basic(description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciFabricVpcDomainPolicyExists("aci_fabric_vpc_domain_policy.foofabric_vpc_domain_policy", &fabric_vpc_domain_policy),
					testAccCheckAciFabricVpcDomainPolicyAttributes(description, &fabric_vpc_domain_policy),
				),
			},
			{
				ResourceName:      "aci_fabric_vpc_domain_policy",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAciFabricVpcDomainPolicy_update(t *testing.T) {
	var fabric_vpc_domain_policy models.FabricVpcDomainPolicy
	description := "fabric_vpc_domain_policy created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciFabricVpcDomainPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciFabricVpcDomainPolicyConfig_basic(description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciFabricVpcDomainPolicyExists("aci_fabric_vpc_domain_policy.foofabric_vpc_domain_policy", &fabric_vpc_domain_policy),
					testAccCheckAciFabricVpcDomainPolicyAttributes(description, &fabric_vpc_domain_policy),
				),
			},
			{
				Config: testAccCheckAciFabricVpcDomainPolicyConfig_basic(description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciFabricVpcDomainPolicyExists("aci_fabric_vpc_domain_policy.foofabric_vpc_domain_policy", &fabric_vpc_domain_policy),
					testAccCheckAciFabricVpcDomainPolicyAttributes(description, &fabric_vpc_domain_policy),
				),
			},
		},
	})
}

func testAccCheckAciFabricVpcDomainPolicyConfig_basic(description string) string {
	return fmt.Sprintf(`
	resource "aci_fabric_vpc_domain_policy" "foofabric_vpc_domain_policy" {
		description = "%s"
        name  = "example"
		annotation  = "example"
		peer_dead_interval  = "200"
		name_alias  = "example"
	}
	`, description)
}

func testAccCheckAciFabricVpcDomainPolicyExists(name string, fabric_vpc_domain_policy *models.FabricVpcDomainPolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("VPC Domain Policy %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VPC Domain Policy dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		fabric_vpc_domain_policyFound := models.FabricVpcDomainPolicyFromContainer(cont)
		if fabric_vpc_domain_policyFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("VPC Domain Policy %s not found", rs.Primary.ID)
		}
		*fabric_vpc_domain_policy = *fabric_vpc_domain_policyFound
		return nil
	}
}

func testAccCheckAciFabricVpcDomainPolicyDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_fabric_vpc_domain_policy" {
			cont, err := client.Get(rs.Primary.ID)
			fabric_vpc_domain_policy := models.FabricVpcDomainPolicyFromContainer(cont)
			if err == nil {
				return fmt.Errorf("VPC Domain Policy %s Still exists", fabric_vpc_domain_policy.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciFabricVpcDomainPolicyAttributes(description string, fabric_vpc_domain_policy *models.FabricVpcDomainPolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if description != fabric_vpc_domain_policy.Description {
			return fmt.Errorf("Bad fabric_vpc_domain_policy Description %s", fabric_vpc_domain_policy.Description)
		}

		if "example" != fabric_vpc_domain_policy.Name {
			return fmt.Errorf("Bad fabric_vpc_domain_policy name %s", fabric_vpc_domain_policy.Name)
		}

		if "200" != fabric_vpc_domain_policy.PeerDeadInterval {
			return fmt.Errorf("Bad fabric_vpc_domain_policy peer_dead_interval %s", fabric_vpc_domain_policy.PeerDeadInterval)
		}

		if "example" != fabric_vpc_domain_policy.Annotation {
			return fmt.Errorf("Bad fabric_vpc_domain_policy annotation %s", fabric_vpc_domain_policy.Annotation)
		}

		if "example" != fabric_vpc_domain_policy.NameAlias {
			return fmt.Errorf("Bad Fabric interface policy name_alias %s", fabric_vpc_domain_policy.NameAlias)
		}

		return nil
	}
}
