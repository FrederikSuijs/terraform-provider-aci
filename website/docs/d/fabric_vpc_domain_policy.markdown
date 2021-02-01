---
layout: "aci"
page_title: "ACI: aci_fabric_vpc_domain_policy"
sidebar_current: "docs-aci-data-source-fabric-vpc-domain-policy"
description: |-
  Data source for ACI Fabric VPC Domain Policy
---

# aci_fabric_vpc_domain_policy #
Data source for ACI Fabric VPC Domain Policy

## Example Usage ##

```hcl
data "aci_fabric_vpc_domain_policy" "example" {
  name  = "example"
}
```

## Argument Reference ##
* `name` - (Required) name of Object tenant.

## Attribute Reference

* `id` - Attribute id set to the Dn of the Tenant.
* `peer_dead_interval` - (Required) Peer Dead Interval in milliseconds.
* `annotation` - (Optional) annotation for object tenant.
* `name_alias` - (Optional) name_alias for object tenant.
