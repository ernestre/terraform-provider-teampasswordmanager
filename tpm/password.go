package tpm

//go:generate go run golang.org/x/tools/cmd/stringer -type=ExpireStatus
type ExpireStatus int

const (
	NotExpired ExpireStatus = iota
	ExpiresToday
	Expired
	ExpiresSoon
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=LockingType
type LockingType int

const (
	NotLocked LockingType = iota
	RequiresReasonToUnlock
	RequiresPermissionToUnlock
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=LockingRequestNotify
type LockingRequestNotify int

const (
	// TODO: Think how to handle clashing name issue with LockingRequestNotify
	PasswordNotLocked LockingRequestNotify = iota
	NotifyManager
	NotifyAll
)

type (
	Permission struct {
		ID    int    `json:"id,omitempty"`
		Label string `json:"label,omitempty"`
	}

	User struct {
		ID       int    `json:"id,omitempty"`
		Username string `json:"username,omitempty"`
		Email    string `json:"email_address,omitempty"`
		Name     string `json:"name,omitempty"`
		Role     string `json:"Role,omitempty"`
	}

	Group struct {
		ID   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	}

	UserPermission struct {
		User       User       `json:"user,omitempty"`
		Permission Permission `json:"permission,omitempty"`
	}

	GroupPermission struct {
		Group      Group      `json:"groupt,omitempty"`
		Permission Permission `json:"permission,omitempty"`
	}

	// https://teampasswordmanager.com/docs/api-passwords/#show_password
	PasswordData struct {
		ID      int    `json:"id,omitempty"`
		Name    string `json:"name,omitempty"`
		Project struct {
			Name string `json:"name,omitempty"`
			ID   int    `json:"id,omitempty"`
		} `json:"project,omitempty"`
		Tags         Tags         `json:"tags"`
		AccessInfo   string       `json:"access_info,omitempty"`
		Username     string       `json:"username,omitempty"`
		Email        string       `json:"email,omitempty"`
		Password     string       `json:"password,omitempty"`
		ExpiryDate   ShortDate    `json:"expiry_date,omitempty"`
		ExpiryStatus ExpireStatus `json:"expiry_status,omitempty"`
		Notes        string       `json:"notes,omitempty"`

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

		UsersPermissions     []UserPermission     `json:"users_permissions,omitempty"`
		GroupsPermissions    []GroupPermission    `json:"groups_permissions,omitempty"`
		Parents              []int                `json:"parents,omitempty"`
		UserPermission       Permission           `json:"user_permission,omitempty"`
		Archived             bool                 `json:"archived,omitempty"`
		ProjectArchived      bool                 `json:"project_archived,omitempty"`
		Favorite             bool                 `json:"favorite,omitempty"`
		NumberOfFiles        int                  `json:"num_files,omitempty"`
		Locked               bool                 `json:"locked,omitempty"`
		LockingType          LockingType          `json:"locking_type,omitempty"`
		LockingRequestNotify LockingRequestNotify `json:"locking_request_notify,omitempty"`
		ExternalSharing      bool                 `json:"external_sharing,omitempty"`
		ExternalURL          string               `json:"external_url,omitempty"`
		Linked               bool                 `json:"linked,omitempty"`
		SourcePasswordID     int                  `json:"source_password_id,omitempty"`
		ManagedBy            User                 `json:"managed_by,omitempty"`
		CreatedOn            LongDate             `json:"created_on,omitempty"`
		CreatedBy            User                 `json:"created_by,omitempty"`
		UpdatedOn            LongDate             `json:"updated_on,omitempty"`
		UpdatedBy            User                 `json:"updated_by,omitempty"`
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
	// https://teampasswordmanager.com/docs/api-passwords#create_password
	CreatePasswordRequest struct {
		Name         string    `json:"name,omitempty"`
		ProjectID    int       `json:"project_id"`
		Tags         Tags      `json:"tags"`
		AccessInfo   string    `json:"access_info"`
		Username     string    `json:"username"`
		Email        string    `json:"email"`
		Password     string    `json:"password"`
		ExpiryDate   ShortDate `json:"expiry_date"`
		Notes        string    `json:"notes"`
		CustomData1  string    `json:"custom_data1"`
		CustomData2  string    `json:"custom_data2"`
		CustomData3  string    `json:"custom_data3"`
		CustomData4  string    `json:"custom_data4"`
		CustomData5  string    `json:"custom_data5"`
		CustomData6  string    `json:"custom_data6"`
		CustomData7  string    `json:"custom_data7"`
		CustomData8  string    `json:"custom_data8"`
		CustomData9  string    `json:"custom_data9"`
		CustomData10 string    `json:"custom_data10"`
	}

	CreatePasswordResponse struct {
		ID int `json:"id,omitempty"`
	}

	// https://teampasswordmanager.com/docs/api-passwords#update_password
	UpdatePasswordRequest struct {
		Name         string    `json:"name"`
		Tags         Tags      `json:"tags"`
		AccessInfo   string    `json:"access_info"`
		Username     string    `json:"username,"`
		Email        string    `json:"email,"`
		Password     string    `json:"password"`
		ExpiryDate   ShortDate `json:"expiry_date"`
		Notes        string    `json:"notes"`
		CustomData1  string    `json:"custom_data1"`
		CustomData2  string    `json:"custom_data2"`
		CustomData3  string    `json:"custom_data3"`
		CustomData4  string    `json:"custom_data4"`
		CustomData5  string    `json:"custom_data5"`
		CustomData6  string    `json:"custom_data6"`
		CustomData7  string    `json:"custom_data7"`
		CustomData8  string    `json:"custom_data8"`
		CustomData9  string    `json:"custom_data9"`
		CustomData10 string    `json:"custom_data10"`
	}

	UpdatePasswordSecurityRequest struct {
		ManagedBy         int                  `json:"managed_by,omitempty"`
		UsersPermissions  []PasswordPermission `json:"users_permissions"`
		GroupsPermissions []PasswordPermission `json:"groups_permissions"`
	}
)
