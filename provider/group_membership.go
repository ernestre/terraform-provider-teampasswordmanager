package provider

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrFailedToParseGroupMembershipID      = errors.New("failed to parse group membership id")
	ErrFailedToParseGroupMembershipUserID  = errors.New("failed to parse group membership's user id")
	ErrFailedToParseGroupMembershipGroupID = errors.New("failed to parse group membership's group id")
)

func generateGroupMemeberShipID(groupID, userID int) string {
	return fmt.Sprintf("%d-%d", groupID, userID)
}

func parseGroupMemeberShipID(membershipID string) (int, int, error) {
	ids := strings.Split(membershipID, "-")

	if len(ids) < 2 {
		return 0, 0, ErrFailedToParseGroupMembershipID
	}

	groupID, err := strconv.Atoi(ids[0])
	if err != nil {
		return 0, 0, ErrFailedToParseGroupMembershipGroupID
	}

	userID, err := strconv.Atoi(ids[1])
	if err != nil {
		return 0, 0, ErrFailedToParseGroupMembershipUserID
	}

	return groupID, userID, nil
}
