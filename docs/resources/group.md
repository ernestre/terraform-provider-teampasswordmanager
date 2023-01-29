---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "teampasswordmanager_group Resource - terraform-provider-teampasswordmanager"
subcategory: ""
description: |-
  Creates a group.
---

# teampasswordmanager_group (Resource)

Creates a group.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the group.

### Optional

- `id` (String) Project ID.

### Read-Only

- `created_by` (List of Object) User which created the password. (see [below for nested schema](#nestedatt--created_by))
- `created_on` (String) Datetime when the password was created.
- `is_ldap` (Boolean) Whether the group is a ldap group.
- `num_users` (Number) Number of users in a group.
- `updated_by` (List of Object) User which updated the password. (see [below for nested schema](#nestedatt--updated_by))
- `updated_on` (String) Datetime when the password was updated.
- `users` (List of Object) Users of the group. (see [below for nested schema](#nestedatt--users))

<a id="nestedatt--created_by"></a>
### Nested Schema for `created_by`

Read-Only:

- `email_address` (String)
- `id` (Number)
- `name` (String)
- `role` (String)
- `username` (String)


<a id="nestedatt--updated_by"></a>
### Nested Schema for `updated_by`

Read-Only:

- `email_address` (String)
- `id` (Number)
- `name` (String)
- `role` (String)
- `username` (String)


<a id="nestedatt--users"></a>
### Nested Schema for `users`

Read-Only:

- `email_address` (String)
- `id` (Number)
- `name` (String)
- `role` (String)
- `username` (String)

