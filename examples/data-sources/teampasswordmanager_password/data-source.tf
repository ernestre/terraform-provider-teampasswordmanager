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
  password   = "feechu0W"
  project_id = teampasswordmanager_project.new.id
}

resource "teampasswordmanager_password" "sendgrid" {
  name       = "sendgrid"
  password   = "az4Oowis"
  project_id = teampasswordmanager_project.new.id
}

data "teampasswordmanager_password" "database" {
  id = teampasswordmanager_password.database.id
}

data "teampasswordmanager_password" "sendgrid" {
  id = teampasswordmanager_password.sendgrid.id
}
