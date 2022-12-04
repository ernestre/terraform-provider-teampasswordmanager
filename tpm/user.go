package tpm

type UserRole string

// Admin: can do anything.
// IT: like a project manager plus access to Users/Groups, Log and Settings.
// Project manager: like normal users and also create/delete projects.
// Normal user: work with passwords and projects, but not create/delete projects.
// Read only: only read passwords on assigned projects.

const (
	// UserRoleAdmin role owners can do anything.
	UserRoleAdmin UserRole = "admin"
	// UserRoleProjectManager role owners are like normal users and also create/delete projects.
	UserRoleProjectManager UserRole = "project manager"
	// UserRoleNormalUser role owners work with passwords and projects, but not create/delete projects.
	UserRoleNormalUser UserRole = "normal user"
	// UserRoleReadOnly role owners only read passwords on assigned projects.
	UserRoleReadOnly UserRole = "read only"
	// UserRoleReadOnly role owners like a project manager plus access to Users/Groups, Log and Settings. read passwords on assigned projects.
	UserRoleIT UserRole = "it"
)

type (
	CreateUserRequest struct {
		Username                string   `json:"username"`
		EmailAddress            string   `json:"email_address"`
		Name                    string   `json:"name"`
		Role                    UserRole `json:"role"`
		Password                string   `json:"password"`
		CanCreateProjectsInRoot bool     `json:"can_create_projects_in_root,omitempty"`
	}

	CreateUserResponse struct {
		ID int `json:"id"`
	}
)
