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
