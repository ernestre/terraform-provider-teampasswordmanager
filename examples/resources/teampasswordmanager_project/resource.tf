resource "teampasswordmanager_project" "new" {
  name  = "wordpress"
  notes = "wordpress secrets"
  tags = [
    "e-shop",
    "wp",
  ]
}

resource "teampasswordmanager_project" "child" {
  name      = "wordpress-copy"
  parent_id = teampasswordmanager_project.new.id
}
