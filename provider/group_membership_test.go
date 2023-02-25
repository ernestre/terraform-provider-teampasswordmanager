package provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateGroupMemeberShipID(t *testing.T) {
	tests := []struct {
		name         string
		groupID      int
		userID       int
		membershipID string
	}{
		{
			name:         "non zero values",
			groupID:      10,
			userID:       30,
			membershipID: "10-30",
		},
		{
			name:         "zero values",
			groupID:      0,
			userID:       0,
			membershipID: "0-0",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mID := generateGroupMemeberShipID(test.groupID, test.userID)

			assert.Equal(t, test.membershipID, mID)
		})
	}
}

func TestParseGroupMemeberShipID(t *testing.T) {
	tests := []struct {
		name         string
		groupID      int
		userID       int
		membershipID string
		error        error
	}{
		{
			name:         "non zero values",
			membershipID: "10-30",
			groupID:      10,
			userID:       30,
		},
		{
			name:         "zero values",
			membershipID: "0-0",
			groupID:      0,
			userID:       0,
		},
		{
			name:         "invalid id",
			membershipID: "hello world",
			groupID:      0,
			userID:       0,
			error:        ErrFailedToParseGroupMembershipID,
		},
		{
			name:         "invalid user id",
			membershipID: "20-go",
			groupID:      0,
			userID:       0,
			error:        ErrFailedToParseGroupMembershipUserID,
		},
		{
			name:         "invalid group id",
			membershipID: "go-20",
			groupID:      0,
			userID:       0,
			error:        ErrFailedToParseGroupMembershipGroupID,
		},
		{
			name:         "invalid user and group id",
			membershipID: "go-foo",
			groupID:      0,
			userID:       0,
			error:        ErrFailedToParseGroupMembershipGroupID,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			groupID, userID, err := parseGroupMemeberShipID(test.membershipID)

			assert.ErrorIs(t, err, test.error)
			assert.Equal(t, test.groupID, groupID)
			assert.Equal(t, test.userID, userID)
		})
	}
}
