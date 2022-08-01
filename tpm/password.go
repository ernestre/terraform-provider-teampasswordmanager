package tpm

type (
	PasswordData struct {
		ID      int    `json:"id,omitempty"`
		Name    string `json:"name,omitempty"`
		Project struct {
			Name string `json:"name,omitempty"`
			ID   int    `json:"id,omitempty"`
		} `json:"project,omitempty"`
		Password      string       `json:"password,omitempty"`
		Username      string       `json:"username,omitempty"`
		Email         string       `json:"email,omitempty"`
		CustomField1  *CustomField `json:"custom_field1,omitempty"`
		CustomField2  *CustomField `json:"custom_field2,omitempty"`
		CustomField3  *CustomField `json:"custom_field3,omitempty"`
		CustomField4  *CustomField `json:"custom_field4,omitempty"`
		CustomField5  *CustomField `json:"custom_field5,omitempty"`
		CustomField6  *CustomField `json:"custom_field6,omitempty"`
		CustomField7  *CustomField `json:"custom_field7,omitempty"`
		CustomField8  *CustomField `json:"custom_field8,omitempty"`
		CustomField9  *CustomField `json:"custom_field9,omitempty"`
		CustomField10 *CustomField `json:"custom_field10,omitempty"`
	}

	CustomField struct {
		Type  string `json:"type,omitempty"`
		Label string `json:"label,omitempty"`
		Data  string `json:"data,omitempty"`
	}

	Password struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Project struct {
			Name string `json:"name"`
		} `json:"project"`
	}

	ErrorResponse struct {
		Error   bool   `json:"error"`
		Type    string `json:"type"`
		Message string `json:"message"`
	}
)

type (
	CreatePasswordRequest struct {
		Name         string `json:"name,omitempty"`
		ProjectID    int    `json:"project_id,omitempty"`
		Password     string `json:"password"`
		Username     string `json:"username"`
		Email        string `json:"email"`
		CustomData1  string `json:"custom_data1"`
		CustomData2  string `json:"custom_data2"`
		CustomData3  string `json:"custom_data3"`
		CustomData4  string `json:"custom_data4"`
		CustomData5  string `json:"custom_data5"`
		CustomData6  string `json:"custom_data6"`
		CustomData7  string `json:"custom_data7"`
		CustomData8  string `json:"custom_data8"`
		CustomData9  string `json:"custom_data9"`
		CustomData10 string `json:"custom_data10"`
	}

	CreatePasswordResponse struct {
		ID int `json:"id,omitempty"`
	}

	UpdatePasswordRequest struct {
		Name         string `json:"name,omitempty"`
		Password     string `json:"password,omitempty"`
		Username     string `json:"username,"`
		Email        string `json:"email,"`
		CustomData1  string `json:"custom_data1"`
		CustomData2  string `json:"custom_data2"`
		CustomData3  string `json:"custom_data3"`
		CustomData4  string `json:"custom_data4"`
		CustomData5  string `json:"custom_data5"`
		CustomData6  string `json:"custom_data6"`
		CustomData7  string `json:"custom_data7"`
		CustomData8  string `json:"custom_data8"`
		CustomData9  string `json:"custom_data9"`
		CustomData10 string `json:"custom_data10"`
	}
)
