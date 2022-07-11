resource "teampasswordmanager_project" "parent" {
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

data "teampasswordmanager_project" "parent_project" {
  id = teampasswordmanager_project.new.id
}

data "teampasswordmanager_project" "child" {
  id = teampasswordmanager_project.child.id
}
