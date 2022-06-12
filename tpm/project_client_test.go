package tpm

import (
	"encoding/json"
	"testing"
)

func areTagsEqual(a, b Tags) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestTags(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		tags           Tags
		expectedResult string
	}{
		{
			name:           "Empty list",
			tags:           Tags{},
			expectedResult: "",
		},
		{
			name:           "Already sorted list",
			tags:           Tags{"a", "b", "c"},
			expectedResult: "a,b,c",
		},
		{
			name:           "Reverse list",
			tags:           Tags{"c", "b", "a"},
			expectedResult: "a,b,c",
		},
		{
			name:           "Mixed list",
			tags:           Tags{"c", "a", "b", "10", "2", "123"},
			expectedResult: "10,123,2,a,b,c",
		},
		{
			name:           "list with empty values",
			tags:           Tags{"c", "a", "", "", "2", "", "", "", "123"},
			expectedResult: "123,2,a,c",
		},
		{
			name:           "items with whitespace",
			tags:           Tags{"c", "a", "", " ", "2    ", "", "  ", "  ", "123"},
			expectedResult: "123,2,a,c",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if test.expectedResult != test.tags.ToString() {
				t.Errorf(
					"failed to sort tags. Expected %s, got %s",
					test.expectedResult,
					test.tags,
				)
			}
		})
	}
}

func TestTagsMarshaling(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		tags           Tags
		expectedResult []byte
	}{
		{
			name:           "Empty list",
			tags:           Tags{},
			expectedResult: []byte("\"\""),
		},
		{
			name:           "Already sorted list",
			tags:           Tags{"a", "b", "c"},
			expectedResult: []byte("\"a,b,c\""),
		},
		{
			name:           "Reverse list",
			tags:           Tags{"c", "b", "a"},
			expectedResult: []byte("\"a,b,c\""),
		},
		{
			name:           "Mixed list",
			tags:           Tags{"c", "a", "b", "10", "2", "123"},
			expectedResult: []byte("\"10,123,2,a,b,c\""),
		},
		{
			name:           "list with empty values",
			tags:           Tags{"c", "a", "", "", "2", "", "", "", "123"},
			expectedResult: []byte("\"123,2,a,c\""),
		},
		{
			name:           "items with whitespace",
			tags:           Tags{"c", "a", "", " ", "2    ", "", "  ", "  ", "123"},
			expectedResult: []byte("\"123,2,a,c\""),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := json.Marshal(test.tags)
			if err != nil {
				t.Error("Failed to marshal tags")
			}

			if string(test.expectedResult) != string(result) {
				t.Errorf(
					"failed to marshal tags. Expected %s, got %s",
					test.expectedResult,
					test.tags,
				)
			}
		})
	}
}

func TestTagsUnMarshaling(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		tags           []byte
		expectedResult Tags
	}{
		{
			name:           "Empty list",
			tags:           []byte("\"\""),
			expectedResult: Tags{},
		},
		{
			name:           "Already sorted list",
			tags:           []byte("\"a,b,c\""),
			expectedResult: Tags{"a", "b", "c"},
		},
		{
			name:           "Reverse list",
			tags:           []byte("\"a,b,c\""),
			expectedResult: Tags{"a", "b", "c"},
		},
		{
			name:           "Mixed list",
			tags:           []byte("\"10,123,2,a,b,c\""),
			expectedResult: Tags{"10", "123", "2", "a", "b", "c"},
		},
		{
			name:           "list with empty values",
			tags:           []byte("\"c,a,,,2,,,,123\""),
			expectedResult: Tags{"c", "a", "", "", "2", "", "", "", "123"},
		},
		{
			name:           "items with whitespace",
			tags:           []byte("\"c,a,, ,2    ,,  ,  ,123\""),
			expectedResult: Tags{"c", "a", "", " ", "2    ", "", "  ", "  ", "123"},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result := Tags{}

			err := json.Unmarshal(test.tags, &result)
			if err != nil {
				t.Errorf("Failed to unmarshal tags: %v", err)
			}

			if !areTagsEqual(test.expectedResult, result) {
				t.Errorf(
					"failed to unmarshal tags. Expected %s, got %s",
					test.expectedResult,
					result,
				)
			}
		})
	}
}
