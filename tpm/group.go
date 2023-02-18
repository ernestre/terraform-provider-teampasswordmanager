package tpm

type (
	CreateGroupRequest struct {
		Name string `json:"name"`
	}

	UpdateGroupRequest struct {
		Name string `json:"name"`
	}

	CreateGroupResponse struct {
		ID int `json:"id"`
	}

	GetGroupResponse struct {
		ID           int      `json:"id,omitempty"`
		Name         string   `json:"name,omitempty"`
		IsLdap       bool     `json:"is_ldap,omitempty"`
		LdapServerID int      `json:"ldap_server_id,omitempty"`
		GroupDn      string   `json:"group_dn,omitempty"`
		Users        []User   `json:"users,omitempty"`
		CreatedOn    LongDate `json:"created_on,omitempty"`
		CreatedBy    User     `json:"created_by,omitempty"`
		UpdatedOn    LongDate `json:"updated_on,omitempty"`
		UpdatedBy    User     `json:"updated_by,omitempty"`
	}
)
