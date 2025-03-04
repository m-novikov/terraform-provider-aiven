---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "aiven_account_team_member Data Source - terraform-provider-aiven"
subcategory: ""
description: |-
  The Account Team Member data source provides information about the existing Aiven Account Team Member.
---

# aiven_account_team_member (Data Source)

The Account Team Member data source provides information about the existing Aiven Account Team Member.

## Example Usage

```terraform
data "aiven_account_team_member" "foo" {
  account_id = aiven_account.<ACCOUNT_RESOURCE>.account_id
  team_id    = aiven_account_team.<TEAM_RESOURCE>.team_id
  user_email = "user+1@example.com"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) The unique account id This property cannot be changed, doing so forces recreation of the resource.
- `team_id` (String) An account team id This property cannot be changed, doing so forces recreation of the resource.
- `user_email` (String) Is a user email address that first will be invited, and after accepting an invitation, he or she becomes a member of a team. This property cannot be changed, doing so forces recreation of the resource.

### Read-Only

- `accepted` (Boolean) is a boolean flag that determines whether an invitation was accepted or not by the user. `false` value means that the invitation was sent to the user but not yet accepted. `true` means that the user accepted the invitation and now a member of an account team.
- `create_time` (String) Time of creation
- `id` (String) The ID of this resource.
- `invited_by_user_email` (String) The email address that invited this user.


