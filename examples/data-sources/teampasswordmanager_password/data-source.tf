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

output "sendgrid_password_managed_by_email_address" {
  value = data.teampasswordmanager_password.sendgrid.managed_by.0.email_address
}

output "database_password_created_by_username" {
  value = data.teampasswordmanager_password.database.created_by.0.username
}

output "database_password_created_by_user_role" {
  value = data.teampasswordmanager_password.database.created_by.0.role
}

output "database_password_created_by_user" {
  value = data.teampasswordmanager_password.database.created_by.0
}
