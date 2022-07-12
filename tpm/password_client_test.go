package tpm

import (
	"testing"
)

func TestPasswordClient_GetEndpoints(t *testing.T) {
	passwordClient := NewPasswordClient(testConfig)

	tests := []struct {
		name         string
		expectedURL  string
		generateFunc func() string
	}{
		{
			name:        "generateURL",
			expectedURL: "host/index.php/foo",
			generateFunc: func() string {
				return passwordClient.generateURL("foo")
			},
		},
		{
			name:        "getProjectPasswordsEndpoint",
			expectedURL: "api/v5/projects/100/passwords.json",
			generateFunc: func() string {
				return passwordClient.getProjectPasswordsEndpoint(100)
			},
		},
		{
			name:        "getPasswordsEndpoint",
			expectedURL: "api/v5/passwords.json",
			generateFunc: func() string {
				return passwordClient.getPasswordsEndpoint()
			},
		},
		{
			name:        "getPasswordByIDEndpoint",
			expectedURL: "api/v5/passwords/100.json",
			generateFunc: func() string {
				return passwordClient.getPasswordByIDEndpoint(100)
			},
		},
		{
			name:        "getUpdatePasswordEndpoint",
			expectedURL: "api/v5/passwords/100.json",
			generateFunc: func() string {
				return passwordClient.getUpdatePasswordEndpoint(100)
			},
		},
		{
			name:        "getPasswordSearchEndpoint",
			expectedURL: "api/v5/passwords/search/hello%3Dworld.json",
			generateFunc: func() string {
				return passwordClient.getPasswordSearchEndpoint("hello=world")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			generatedURL := test.generateFunc()
			if generatedURL != test.expectedURL {
				t.Errorf(
					"Generated URL does not match expected URL. Expected %s, got %s",
					test.expectedURL,
					generatedURL,
				)
			}
		})
	}
}
