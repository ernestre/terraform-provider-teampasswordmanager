package tpm

import "errors"

type ApiError struct {
	Type    string
	Message string
}

func (e ApiError) Error() string {
	return e.Message
}

var (
	ErrPasswordNameIsRequired = errors.New("password name is a required")
	ErrProjectIDIsRequired    = errors.New("project id is a required")
	ErrPasswordNotFound       = errors.New("the password does not exist or you cannot access it")
	ErrProjectNotFound        = errors.New("the project does not exist or you cannot access it")

	ErrProjectNameIsRequired = errors.New("project name is a required")
)
