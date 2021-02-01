package aci

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAciFabricVpcDomainPolicy() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceAciFabricVpcDomainPolicyRead,

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"peer_dead_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"name_alias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}

func dataSourceAciFabricVpcDomainPolicyRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	name := d.Get("name").(string)

	rn := fmt.Sprintf("vpcInst-%s", name)

	dn := fmt.Sprintf("uni/fabric/%s", rn)

	fvFabricVpcDomainPolicy, err := getRemoteFabricVpcDomainPolicy(aciClient, dn)

	if err != nil {
		return err
	}
	setFabricVpcDomainPolicy(fvFabricVpcDomainPolicy, d)
	return nil
}
