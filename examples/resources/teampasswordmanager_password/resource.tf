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

  custom_field_1 = "Admin user"
  custom_field_2 = "Marketing"
}

output "sendgrid_password_created_by_username" {
  value = teampasswordmanager_password.sendgrid.created_by.0.username
}

output "sendgrid_password_created_by_user_role" {
  value = teampasswordmanager_password.sendgrid.created_by.0.role
}

output "sendgrid_password_created_by_user" {
  value = teampasswordmanager_password.sendgrid.created_by.0
}
