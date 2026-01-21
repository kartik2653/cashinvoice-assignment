package errors

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUnauthorized      = errors.New("you are unauthorozed to perform this action")
	ErrInvalidRoleValue  = errors.New("invalid role value")
	ErrInvalidTaskStatus = errors.New("invalid status: must be 'pending', 'in_progress', or 'completed'")
)
