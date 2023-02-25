resource "teampasswordmanager_group" "new" {
  name = "new_group"
}

resource "teampasswordmanager_group_membership" "new_group_user" {
  group_id = teampasswordmanager_group.new.id
  user_id  = 1
}
