---
layout: "aci"
page_title: "ACI: aci_fabric_vpc_domain_policy"
sidebar_current: "docs-aci-resource-fabric-vpc-domain-policy"
description: |-
  Manages ACI Fabric VPC Domain Policy
---

# aci_fabric_vpc_domain_policy #
Manages ACI Fabric VPC Domain Policy

## Example Usage ##

```hcl
resource "aci_fabric_vpc_domain_policy" "foovpcdomainpolicy" {
  description        = "%s"
  name               = "demo_"
  peer_dead_interval = "123"
  annotation         = "tag_fabric_vpc_domain_policy"
  name_alias         = "alias_fabric_vpc_domain_policy"
}
```
## Argument Reference ##
* `name` - (Required) name of Object tenant.
* `peer_dead_interval` - (Required) Peer Dead Interval in milliseconds.
* `annotation` - (Optional) annotation for object tenant.
* `name_alias` - (Optional) name_alias for object tenant.                


## Attribute Reference

The only attribute that this resource exports is the `id`, which is set to the
Dn of the Fabric VPC Domain Policy.

## Importing ##

An existing Fabric VPC Domain Policy can be [imported][docs-import] into this resource via its Dn, via the following command:
[docs-import]: https://www.terraform.io/docs/import/index.html


```
terraform import aci_fabric_vpc_domain_policy.example <Dn>
```