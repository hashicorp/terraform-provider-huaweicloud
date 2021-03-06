---
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_css_cluster_v1"
sidebar_current: "docs-huaweicloud-resource-css-cluster-v1"
description: |-
 CSS cluster management
---

# huaweicloud\_css\_cluster\_v1

CSS cluster management

## Example Usage

### create a cluster

```hcl
resource "huaweicloud_networking_secgroup_v2" "secgroup" {
  name        = "terraform_test_security_group"
  description = "terraform security group acceptance test"
}

resource "huaweicloud_css_cluster_v1" "cluster" {
  expect_node_num = 1
  name            = "terraform_test_cluster"
  engine_version  = "7.1.1"
  node_config {
    flavor = "ess.spec-2u16g"
    network_info {
      security_group_id = "${huaweicloud_networking_secgroup_v2.secgroup.id}"
      subnet_id         = "{{ network_id }}"
      vpc_id            = "{{ vpc_id }}"
    }
    volume {
      volume_type = "COMMON"
      size        = 40
    }
    availability_zone = "{{ availability_zone }}"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` -
  (Required)
  Cluster name. It contains 4 to 32 characters. Only letters, digits,
  hyphens (-), and underscores (_) are allowed. The value must start
  with a letter. Changing this parameter will create a new resource.

* `engine_type` -
  (Optional)
  Engine type. The default value is "elasticsearch". Currently, the value
  can only be "elasticsearch". Changing this parameter will create a new resource.

* `engine_version` -
  (Required)
  Engine version. Versions 5.5.1, 6.2.3, 6.5.4 and 7.1.1 are supported. Changing this parameter will create a new resource.

* `expect_node_num` -
  (Optional)
  Number of cluster instances. The value range is 1 to 32. Defaults to 1.

* `node_config` -
  (Required)
  Node configuration. Structure is documented below. Changing this parameter will create a new resource.

* `backup_strategy` - (Optional) Specifies the advanced backup policy. Structure is documented below.

* `tags` - (Optional) The key/value pairs to associate with the cluster.
  Changing this parameter will create a new resource.

The `node_config` block supports:

* `availability_zone` -
  (Optional)
  Availability zone (AZ).  Changing this parameter will create a new resource.

* `flavor` -
  (Required)
  Instance flavor name. For example: value range of flavor ess.spec-2u8g:
  40 GB to 800 GB, value range of flavor ess.spec-4u16g: 40 GB to 1600 GB,
  value range of flavor ess.spec-8u32g: 80 GB to 3200 GB, value range of
  flavor ess.spec-16u64g: 100 GB to 6400 GB, value range of
  flavor ess.spec-32u128g: 100 GB to 10240 GB.
  Changing this parameter will create a new resource.

* `network_info` -
  (Required)
  Network information. Structure is documented below. Changing this parameter will create a new resource.

* `volume` -
  (Required)
  Information about the volume. Structure is documented below. Changing this parameter will create a new resource.

The `network_info` block supports:

* `vpc_id` -
  (Required)
  VPC ID, which is used for configuring cluster network. Changing this parameter will create a new resource.

* `subnet_id` -
  (Required)
  Subnet ID. All instances in a cluster must have the same subnet which should be configured with a *DNS address*.
  Changing this parameter will create a new resource.

* `security_group_id` -
  (Required)
  Security group ID. All instances in a cluster must have the same security group.
  Changing this parameter will create a new resource.

The `volume` block supports:

* `size` -
  (Required)
  Specifies the volume size in GB, which must be a multiple of 10.

* `volume_type` -
  (Required)
  Specifies the volume type. COMMON: Common I/O. The SATA disk is used. HIGH: High I/O.
  The SAS disk is used. ULTRAHIGH: Ultra-high I/O. The
  solid-state drive (SSD) is used. Changing this parameter will create a new resource.


The `backup_strategy` block supports:

* `start_time` - (Required) Specifies the time when a snapshot is automatically
  created everyday. Snapshots can only be created on the hour. The time format is
  the time followed by the time zone, specifically, **HH:mm z**. In the format, HH:mm
  refers to the hour time and z refers to the time zone. For example, "00:00 GMT+08:00"
  and "01:00 GMT+08:00".

* `keep_days` - (Optional) Specifies the number of days to retain the generated
   snapshots. Snapshots are reserved for seven days by default.

* `prefix` - (Optional) Specifies the prefix of the snapshot that is automatically
  created. The default value is "snapshot".


## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `endpoint` -
  Indicates the IP address and port number.

* `created` -
  Time when a cluster is created. The format is ISO8601:
  CCYY-MM-DDThh:mm:ss.

* `nodes` -
  List of node objects. Structure is documented below.

The `nodes` block contains:

* `id` - Instance ID.

* `name` - Instance name.

* `type` - Supported type: ess (indicating the Elasticsearch node).

## Timeouts

This resource provides the following timeouts configuration options:

- `create` - Default is 30 minute.
- `update` - Default is 30 minute.
