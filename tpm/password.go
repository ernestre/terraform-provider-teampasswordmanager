package tpm

type (
	PasswordData struct {
		ID      int    `json:"id,omitempty"`
		Name    string `json:"name,omitempty"`
		Project struct {
			Name string `json:"name,omitempty"`
			ID   int    `json:"id,omitempty"`
		} `json:"project,omitempty"`
		Password     string       `json:"password,omitempty"`
		CustomField1 *CustomField `json:"custom_field1,omitempty"`
		CustomField2 *CustomField `json:"custom_field2,omitempty"`
		CustomField3 *struct {
			CustomField `json:"custom_field"`
			OtpValue    string `json:"otp_value,omitempty"`
		} `json:"custom_field3,omitempty"`
		CustomField4  *CustomField `json:"custom_field4,omitempty"`
		CustomField5  interface{}  `json:"custom_field5,omitempty"`
		CustomField6  interface{}  `json:"custom_field6,omitempty"`
		CustomField7  interface{}  `json:"custom_field7,omitempty"`
		CustomField8  interface{}  `json:"custom_field8,omitempty"`
		CustomField9  interface{}  `json:"custom_field9,omitempty"`
		CustomField10 interface{}  `json:"custom_field10,omitempty"`
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
		Password     string `json:"password,omitempty"`
		CustomData1  string `json:"custom_data_1,omitempty"`
		CustomData2  string `json:"custom_data_2,omitempty"`
		CustomData4  string `json:"custom_data_4,omitempty"`
		CustomData5  string `json:"custom_data_5,omitempty"`
		CustomData6  string `json:"custom_data_6,omitempty"`
		CustomData7  string `json:"custom_data_7,omitempty"`
		CustomData8  string `json:"custom_data_8,omitempty"`
		CustomData9  string `json:"custom_data_9,omitempty"`
		CustomData10 string `json:"custom_data_10,omitempty"`
	}

	CreatePasswordResponse struct {
		ID int `json:"id,omitempty"`
	}

	UpdatePasswordRequest struct {
		Name         string `json:"name,omitempty"`
		Password     string `json:"password,omitempty"`
		CustomData1  string `json:"custom_data_1,omitempty"`
		CustomData2  string `json:"custom_data_2,omitempty"`
		CustomData4  string `json:"custom_data_4,omitempty"`
		CustomData5  string `json:"custom_data_5,omitempty"`
		CustomData6  string `json:"custom_data_6,omitempty"`
		CustomData7  string `json:"custom_data_7,omitempty"`
		CustomData8  string `json:"custom_data_8,omitempty"`
		CustomData9  string `json:"custom_data_9,omitempty"`
		CustomData10 string `json:"custom_data_10,omitempty"`
	}
)
