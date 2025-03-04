---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "aiven_service_user Resource - terraform-provider-aiven"
subcategory: ""
description: |-
  The Service User resource allows the creation and management of Aiven Service Users.
---

# aiven_service_user (Resource)

The Service User resource allows the creation and management of Aiven Service Users.

~> **Note:** This resource is deprecated. Please use service-specific resources instead of this one, for example: aiven_kafka_user, aiven_pg_user etc.

## Example Usage

```terraform
resource "aiven_service_user" "myserviceuser" {
  project      = aiven_project.myproject.project
  service_name = aiven_pg.mypg.service_name
  username     = "<USERNAME>"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `project` (String) Identifies the project this resource belongs to. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.
- `service_name` (String) Specifies the name of the service that this resource belongs to. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.
- `username` (String) The actual name of the service user. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.

### Optional

- `authentication` (String) Authentication details. The possible values are `caching_sha2_password` and `mysql_native_password`.
- `password` (String, Sensitive) The password of the service user ( not applicable for all services ).
- `pg_allow_replication` (Boolean) Postgres specific field, defines whether replication is allowed. This property cannot be changed, doing so forces recreation of the resource.
- `redis_acl_categories` (List of String) Redis specific field, defines command category rules. The field is required with`redis_acl_commands` and `redis_acl_keys`. This property cannot be changed, doing so forces recreation of the resource.
- `redis_acl_channels` (List of String) Redis specific field, defines the permitted pub/sub channel patterns. This property cannot be changed, doing so forces recreation of the resource.
- `redis_acl_commands` (List of String) Redis specific field, defines rules for individual commands. The field is required with`redis_acl_categories` and `redis_acl_keys`. This property cannot be changed, doing so forces recreation of the resource.
- `redis_acl_keys` (List of String) Redis specific field, defines key access rules. The field is required with`redis_acl_categories` and `redis_acl_keys`. This property cannot be changed, doing so forces recreation of the resource.
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `access_cert` (String, Sensitive) Access certificate for the user if applicable for the service in question
- `access_key` (String, Sensitive) Access certificate key for the user if applicable for the service in question
- `id` (String) The ID of this resource.
- `type` (String) Type of the user account. Tells wether the user is the primary account or a regular account.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `default` (String)
- `delete` (String)
- `update` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import aiven_service_user.myserviceuser project/service_name/username
```
