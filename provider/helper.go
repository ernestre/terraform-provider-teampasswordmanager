package provider

import "github.com/ernestre/terraform-provider-teampasswordmanager/tpm"

func flattenUsers(u []tpm.User) []map[string]any {
	results := []map[string]any{}
	for _, v := range u {
		results = append(results, flattenUser(v))
	}

	return results
}

func flattenUser(u tpm.User) map[string]interface{} {
	return map[string]interface{}{
		"id":            u.ID,
		"username":      u.Username,
		"email_address": u.Email,
		"name":          u.Name,
		"role":          u.Role,
	}
}

func flattenPermission(up tpm.Permission) map[string]interface{} {
	return map[string]interface{}{
		"id":    up.ID,
		"label": up.Label,
	}
}

func flattenUsersPermissions(up []tpm.UserPermission) []map[string]interface{} {
	userPermissions := []map[string]interface{}{}
	for _, userPermission := range up {
		user := userPermission.User
		permission := userPermission.Permission

		up := map[string]interface{}{
			"user": []map[string]interface{}{{
				"id":            user.ID,
				"username":      user.Username,
				"email_address": user.Email,
				"name":          user.Name,
				"role":          user.Role,
			}},
			"permission": []map[string]interface{}{{
				"id":    permission.ID,
				"label": permission.Label,
			}},
		}

		userPermissions = append(userPermissions, up)
	}

	return userPermissions
}

func flattenGroupsPermissions(gp []tpm.GroupPermission) []map[string]interface{} {
	groupPermissions := []map[string]interface{}{}
	for _, groupPermission := range gp {
		group := groupPermission.Group
		permission := groupPermission.Permission

		up := map[string]interface{}{
			"group": []map[string]interface{}{{
				"id":   group.ID,
				"name": group.Name,
			}},
			"permission": []map[string]interface{}{{
				"id":    permission.ID,
				"label": permission.Label,
			}},
		}

		groupPermissions = append(groupPermissions, up)
	}

	return groupPermissions
}
