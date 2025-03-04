---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "aiven_cassandra_user Data Source - terraform-provider-aiven"
subcategory: ""
description: |-
  The Cassandra User data source provides information about the existing Aiven Cassandra User.
---

# aiven_cassandra_user (Data Source)

The Cassandra User data source provides information about the existing Aiven Cassandra User.

## Example Usage

```terraform
data "aiven_cassandra_user" "user" {
  service_name = "my-service"
  project      = "my-project"
  username     = "user1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `project` (String) Identifies the project this resource belongs to. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.
- `service_name` (String) Specifies the name of the service that this resource belongs to. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.
- `username` (String) The actual name of the Cassandra User. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.

### Read-Only

- `access_cert` (String, Sensitive) Access certificate for the user if applicable for the service in question
- `access_key` (String, Sensitive) Access certificate key for the user if applicable for the service in question
- `id` (String) The ID of this resource.
- `password` (String, Sensitive) The password of the Cassandra User.
- `type` (String) Type of the user account. Tells whether the user is the primary account or a regular account.


