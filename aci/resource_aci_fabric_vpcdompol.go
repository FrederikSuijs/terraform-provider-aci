package aci

import (
	"fmt"
	"log"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAciFabricVpcDomainPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciFabricVpcDomainPolicyCreate,
		Update: resourceAciFabricVpcDomainPolicyUpdate,
		Read:   resourceAciFabricVpcDomainPolicyRead,
		Delete: resourceAciFabricVpcDomainPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciFabricVpcDomainPolicyImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"peer_dead_interval": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"name_alias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}
func getRemoteFabricVpcDomainPolicy(client *client.Client, dn string) (*models.FabricVpcDomainPolicy, error) {
	vpcInstPolCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	vpcInstPol := models.FabricVpcDomainPolicyFromContainer(vpcInstPolCont)

	if vpcInstPol.DistinguishedName == "" {
		return nil, fmt.Errorf("FabricVpcDomainPolicy %s not found", vpcInstPol.DistinguishedName)
	}

	return vpcInstPol, nil
}

func setFabricVpcDomainPolicy(vpcInstPol *models.FabricVpcDomainPolicy, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(vpcInstPol.DistinguishedName)
	d.Set("description", vpcInstPol.Description)
	vpcInstPolMap, _ := vpcInstPol.ToMap()

	d.Set("name", vpcInstPolMap["name"])

	d.Set("peer_dead_interval", vpcInstPolMap["deadIntvl"])
	d.Set("annotation", vpcInstPolMap["annotation"])
	d.Set("name_alias", vpcInstPolMap["nameAlias"])
	return d
}

func resourceAciFabricVpcDomainPolicyImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	log.Printf("[DEBUG] %s: Beginning Import", d.Id())
	aciClient := m.(*client.Client)

	dn := d.Id()

	vpcInstPol, err := getRemoteFabricVpcDomainPolicy(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setFabricVpcDomainPolicy(vpcInstPol, d)

	log.Printf("[DEBUG] %s: Import finished successfully", d.Id())

	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciFabricVpcDomainPolicyCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] FabricVpcDomainPolicy: Beginning Creation")
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	vpcInstPolAttr := models.FabricVpcDomainPolicyAttributes{}
	if PeerDeadInterval, ok := d.GetOk("peer_dead_interval"); ok {
		vpcInstPolAttr.PeerDeadInterval = PeerDeadInterval.(string)
	}
	if Annotation, ok := d.GetOk("annotation"); ok {
		vpcInstPolAttr.Annotation = Annotation.(string)
	} else {
		vpcInstPolAttr.Annotation = "{}"
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		vpcInstPolAttr.NameAlias = NameAlias.(string)
	}

	vpcInstPol := models.NewFabricVpcDomainPolicy(fmt.Sprintf("fabric/vpcInst-%s", name), "uni", desc, vpcInstPolAttr)

	err := aciClient.Save(vpcInstPol)
	if err != nil {
		return err
	}
	d.Partial(true)

	d.SetPartial("name")

	d.Partial(false)

	d.SetId(vpcInstPol.DistinguishedName)
	log.Printf("[DEBUG] %s: Creation finished successfully", d.Id())

	return resourceAciFabricVpcDomainPolicyRead(d, m)
}

func resourceAciFabricVpcDomainPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] FabricVpcDomainPolicy: Beginning Update")

	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	vpcInstPolAttr := models.FabricVpcDomainPolicyAttributes{}
	if PeerDeadInterval, ok := d.GetOk("peer_dead_interval"); ok {
		vpcInstPolAttr.PeerDeadInterval = PeerDeadInterval.(string)
	}
	if Annotation, ok := d.GetOk("annotation"); ok {
		vpcInstPolAttr.Annotation = Annotation.(string)
	} else {
		vpcInstPolAttr.Annotation = "{}"
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		vpcInstPolAttr.NameAlias = NameAlias.(string)
	}

	vpcInstPol := models.NewFabricVpcDomainPolicy(fmt.Sprintf("fabric/vpcInst-%s", name), "uni", desc, vpcInstPolAttr)

	vpcInstPol.Status = "modified"

	err := aciClient.Save(vpcInstPol)

	if err != nil {
		return err
	}
	d.Partial(true)

	d.SetPartial("name")

	d.Partial(false)

	d.SetId(vpcInstPol.DistinguishedName)
	log.Printf("[DEBUG] %s: Update finished successfully", d.Id())

	return resourceAciFabricVpcDomainPolicyRead(d, m)

}

func resourceAciFabricVpcDomainPolicyRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Read", d.Id())

	aciClient := m.(*client.Client)

	dn := d.Id()
	vpcInstPol, err := getRemoteFabricVpcDomainPolicy(aciClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setFabricVpcDomainPolicy(vpcInstPol, d)

	log.Printf("[DEBUG] %s: Read finished successfully", d.Id())

	return nil
}

func resourceAciFabricVpcDomainPolicyDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Destroy", d.Id())

	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "vpcInstPol")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] %s: Destroy finished successfully", d.Id())

	d.SetId("")
	return err
}
