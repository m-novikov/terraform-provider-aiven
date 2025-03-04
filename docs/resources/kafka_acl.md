---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "aiven_kafka_acl Resource - terraform-provider-aiven"
subcategory: ""
description: |-
  The Resource Kafka ACL resource allows the creation and management of ACLs for an Aiven Kafka service.
---

# aiven_kafka_acl (Resource)

The Resource Kafka ACL resource allows the creation and management of ACLs for an Aiven Kafka service.

## Example Usage

```terraform
resource "aiven_kafka_acl" "mytestacl" {
  project      = aiven_project.myproject.project
  service_name = aiven_kafka.myservice.service_name
  topic        = "<TOPIC_NAME_PATTERN>"
  permission   = "admin"
  username     = "<USERNAME_PATTERN>"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `permission` (String) Kafka permission to grant. The possible values are `admin`, `read`, `readwrite` and `write`. This property cannot be changed, doing so forces recreation of the resource.
- `project` (String) Identifies the project this resource belongs to. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.
- `service_name` (String) Specifies the name of the service that this resource belongs to. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.
- `topic` (String) Topic name pattern for the ACL entry. This property cannot be changed, doing so forces recreation of the resource.
- `username` (String) Username pattern for the ACL entry. This property cannot be changed, doing so forces recreation of the resource.

### Optional

- `acl_id` (String) Kafka ACL ID
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `id` (String) The ID of this resource.

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
terraform import aiven_kafka_acl.mytestacl project/service_name/id
```
