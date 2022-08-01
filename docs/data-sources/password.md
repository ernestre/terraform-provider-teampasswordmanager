---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "teampasswordmanager_password Data Source - terraform-provider-teampasswordmanager"
subcategory: ""
description: |-
  Retrieve password information resource for a given project.
---

# teampasswordmanager_password (Data Source)

Retrieve password information resource for a given project.

## Example Usage

```terraform
resource "teampasswordmanager_project" "new" {
  name  = "wordpress"
  notes = "wordpress secrets"
  tags = [
    "e-shop",
    "wp",
  ]
}

resource "teampasswordmanager_password" "database" {
  name       = "database"
  username   = "root"
  email      = "root@example.com"
  password   = "feechu0W"
  project_id = teampasswordmanager_project.new.id
}

resource "teampasswordmanager_password" "sendgrid" {
  name       = "sendgrid"
  username   = "admin"
  email      = "admin@sendgrid.com"
  password   = "az4Oowis"
  project_id = teampasswordmanager_project.new.id
}

data "teampasswordmanager_password" "database" {
  id = teampasswordmanager_password.database.id
}

data "teampasswordmanager_password" "sendgrid" {
  id = teampasswordmanager_password.sendgrid.id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Password ID.

### Read-Only

- `custom_field_1` (String) Custom field 1
- `custom_field_10` (String) Custom field 10
- `custom_field_2` (String) Custom field 2
- `custom_field_3` (String) Custom field 3
- `custom_field_4` (String) Custom field 4
- `custom_field_5` (String) Custom field 5
- `custom_field_6` (String) Custom field 6
- `custom_field_7` (String) Custom field 7
- `custom_field_8` (String) Custom field 8
- `custom_field_9` (String) Custom field 9
- `email` (String, Sensitive) Email value.
- `name` (String) Name of the password, usually used for seaching.
- `password` (String, Sensitive) Password value.
- `project_id` (Number) Project ID of the project where password should be created.
- `username` (String, Sensitive) Username value.

