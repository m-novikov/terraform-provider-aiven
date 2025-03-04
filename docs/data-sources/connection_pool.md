---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "aiven_connection_pool Data Source - terraform-provider-aiven"
subcategory: ""
description: |-
  The Connection Pool data source provides information about the existing Aiven Connection Pool.
---

# aiven_connection_pool (Data Source)

The Connection Pool data source provides information about the existing Aiven Connection Pool.

## Example Usage

```terraform
data "aiven_connection_pool" "mytestpool" {
  project      = aiven_project.myproject.project
  service_name = aiven_pg.mypg.service_name
  pool_name    = "mypool"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `pool_name` (String) The name of the created pool. This property cannot be changed, doing so forces recreation of the resource.
- `project` (String) Identifies the project this resource belongs to. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.
- `service_name` (String) Specifies the name of the service that this resource belongs to. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.

### Read-Only

- `connection_uri` (String, Sensitive) The URI for connecting to the pool
- `database_name` (String) The name of the database the pool connects to. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.
- `id` (String) The ID of this resource.
- `pool_mode` (String) The mode the pool operates in The possible values are `session`, `transaction` and `statement`. The default value is `transaction`.
- `pool_size` (Number) The number of connections the pool may create towards the backend server. This does not affect the number of incoming connections, which is always a much larger number. The default value is `10`.
- `username` (String) The name of the service user used to connect to the database. To set up proper dependencies please refer to this variable as a reference.


