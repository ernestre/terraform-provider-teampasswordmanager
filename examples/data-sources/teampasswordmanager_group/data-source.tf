resource "teampasswordmanager_group" "test_group" {
  name = "test_group"
}

data "teampasswordmanager_group" "test_group_data" {
  id = teampasswordmanager_group.test_group.id
}
