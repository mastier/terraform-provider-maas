---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "maas_vlan Data Source - terraform-provider-maas"
subcategory: ""
description: |-
  Provides details about an existing MAAS VLAN.
---

# maas_vlan (Data Source)

Provides details about an existing MAAS VLAN.

## Example Usage

```terraform
data "maas_vlan" "default" {
  fabric = data.maas_fabric.default.id
  vlan = 0
}

data "maas_vlan" "vid10" {
  fabric = data.maas_fabric.default.id
  vlan = 10
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `fabric` (String) The fabric identifier (ID or name) for the VLAN.
- `vlan` (String) The VLAN identifier (ID or traffic segregation ID).

### Read-Only

- `dhcp_on` (Boolean) Boolean value indicating if DHCP is enabled on the VLAN.
- `id` (String) The ID of this resource.
- `mtu` (Number) The MTU used on the VLAN.
- `name` (String) The VLAN name.
- `space` (String) The VLAN space.


