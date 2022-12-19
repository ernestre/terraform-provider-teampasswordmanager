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

output "sendgrid" {
  value     = data.teampasswordmanager_password.sendgrid.managed_by.0.email_address
  sensitive = true
}

output "database_password_users_permissions_user_email" {
  value     = data.teampasswordmanager_password.foo.users_permissions.0.user.0.email_address
  sensitive = true
}

output "database_password_users_permissions_permission_id" {
  value     = data.teampasswordmanager_password.foo.users_permissions.0.permission.0.id
  sensitive = true
}
