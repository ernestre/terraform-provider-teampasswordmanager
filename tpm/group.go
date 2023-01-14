package tpm

type (
	CreateGroupRequest struct {
		Name string `json:"name"`
	}

	CreateGroupResponse struct {
		ID int `json:"id"`
	}
)
