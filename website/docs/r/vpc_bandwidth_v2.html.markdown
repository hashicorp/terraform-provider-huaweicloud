---
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_vpc_bandwidth_v2"
sidebar_current: "docs-huaweicloud-resource-vpc-bandwidth-v2"
description: |-
  Manages a V2 Shared Bandwidth resource within HuaweiCloud.
---

# huaweicloud_vpc_bandwidth_v2

Manages a Shared Bandwidth resource within HuaweiCloud.

## Example Usage

```hcl
resource "huaweicloud_vpc_bandwidth_v2" "bandwidth_1" {
  name = "bandwidth_1"
  size = 5
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the Shared Bandwidth.

* `size` - (Required) The size of the Shared Bandwidth. The value ranges from 5 to 2000 G.

* `enterprise_project_id` - (Optional) The enterprise project id of the Shared Bandwidth. Changing this creates a new bandwidth.


## Attributes Reference

The following attributes are exported:

* `id` -  ID of the Shared Bandwidth.

* `name` -  See Argument Reference above.

* `size` - See Argument Reference above.

* `enterprise_project_id` - See Argument Reference above.

## Import

Shared Bandwidths can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_vpc_bandwidth_v2.bandwidth_1 7117d38e-4c8f-4624-a505-bd96b97d024c
```
