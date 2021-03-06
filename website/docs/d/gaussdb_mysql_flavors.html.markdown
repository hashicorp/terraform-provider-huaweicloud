---
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_gaussdb_mysql_flavors"
sidebar_current: "docs-huaweicloud-datasource-gaussdb-mysql-flavors"
description: |-
  Get the flavor information on an HuaweiCloud gaussdb mysql.
---

# huaweicloud\_gaussdb\_mysql\_flavors

Use this data source to get available HuaweiCloud gaussdb mysql flavors.

## Example Usage

```hcl
data "huaweicloud_gaussdb_mysql_flavors" "flavors" {
}
```

## Argument Reference

* `engine` - (Optional) Specifies the database engine. Only "gaussdb-mysql" is supported now.

* `version` - (Optional) Specifies the database version. Only "8.0" is supported now.

* `availability_zone_mode` - (Optional) Specifies the availability zone mode. Currently support `single` and 'multi'. Defaults to `single`.

## Attributes Reference

In addition, the following attributes are exported:

* `flavors` -
  Indicates the flavors information. Structure is documented below.

The `flavors` block contains:

* `name` - The name of the gaussdb mysql flavor.
* `vcpus` - Indicates the CPU size.
* `memory` - Indicates the memory size in GB.
* `mode` - Indicates the database mode.
* `version` - Indicates the database version.
