package tpm

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGenerateHash(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		endpoint     string
		timestamp    int64
		requestBody  interface{}
		privateKey   string
		expectedHash string
	}{
		{
			name:       "Create password request",
			endpoint:   "api/v5/projects.json",
			timestamp:  1644913719,
			privateKey: "dc6laksjddk1l23kl12j31l23j2k2jlfaa88afbdd51d7089calsdjaslkdjljc9",
			requestBody: map[string]interface{}{
				"name":      "This is a new password",
				"parent_id": 0,
			},
			expectedHash: "44105abe110f2e02b75584b939e294886a4e5cc42bc873e58d75ebb78795c2f5",
		},
		{
			name:       "Create password request with different key and data",
			endpoint:   "api/v5/projects.json",
			timestamp:  1644915273,
			privateKey: "0000000000000000002j31l23j2k2jlfaa88afbdd51d70800000000000000000",
			requestBody: map[string]interface{}{
				"name": "This is a new password",
			},
			expectedHash: "5fe43518b641d5c56a7e68d4937f2daad46c405bec2c36a0fb4563a22a51f986",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			requestBody, err := json.Marshal(test.requestBody)
			if err != nil {
				t.Error("failed to marshal request body")
			}

			hash := generateAuthHash(test.endpoint, test.timestamp, requestBody, test.privateKey)

			if hash != test.expectedHash {
				t.Errorf("expected hash %s is not equal to %s", test.expectedHash, hash)
			}
		})
	}
}

func TestGenerateAuthHeaders(t *testing.T) {
	t.Parallel()

	publicKey := "0000000000000000000000000000000000000000000000000000000000000000"
	hash := "5fe43518b641d5c56a7e68d4937f2daad46c405bec2c36a0fb4563a22a51f986"
	timestamp := int64(1644915273)

	headers := generateAuthHeaders(publicKey, hash, timestamp)

	expectedHeaders := map[string]string{
		"X-Public-Key":        publicKey,
		"X-Request-Hash":      hash,
		"X-Request-Timestamp": "1644915273",
	}

	if !reflect.DeepEqual(expectedHeaders, headers) {
		t.Errorf("expected headers don't match generated headers %#v != %#v", expectedHeaders, headers)
	}
}
