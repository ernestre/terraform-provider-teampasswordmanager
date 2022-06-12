package tpm

type (
	CreateProjectRequest struct {
		Name     string `json:"name,omitempty"`
		ParentID int    `json:"parent_id"`
		Tags     Tags   `json:"tags,omitempty"`
		Notes    string `json:"notes,omitempty"`
	}

	UpdateProjectRequest struct {
		Name  string `json:"name,omitempty"`
		Tags  Tags   `json:"tags"`
		Notes string `json:"notes"`
	}

	CreateProjectResponse struct {
		ID int `json:"id,omitempty"`
	}

	ProjectData struct {
		ID       int    `json:"id,omitempty"`
		Name     string `json:"name,omitempty"`
		ParentID int    `json:"parent_id,omitempty"`
		Tags     Tags   `json:"tags,omitempty"`
		Notes    string `json:"notes,omitempty"`
	}
)
