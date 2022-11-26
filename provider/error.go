package provider

import "errors"

var (
	ErrInvalidExpiryDateFormat = errors.New("Invalid password expiry date format. Expected format yyyy-mm-dd.")
)
