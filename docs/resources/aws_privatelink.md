---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "aiven_aws_privatelink Resource - terraform-provider-aiven"
subcategory: ""
description: |-
  The AWS Privatelink resource allows the creation and management of Aiven AWS Privatelink for a services.
---

# aiven_aws_privatelink (Resource)

The AWS Privatelink resource allows the creation and management of Aiven AWS Privatelink for a services.

## Example Usage

```terraform
resource "aiven_aws_privatelink" "foo" {
  project      = data.aiven_project.foo.project
  service_name = aiven_kafka.bar.service_name

  principals = [
    "arn:aws:iam::012345678901:user/mwf"
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `principals` (Set of String) List of allowed principals
- `project` (String) Identifies the project this resource belongs to. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.
- `service_name` (String) Specifies the name of the service that this resource belongs to. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `aws_service_id` (String) AWS service ID
- `aws_service_name` (String) AWS service name
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
terraform import aiven_aws_privatelink.foo project/service_name
```
