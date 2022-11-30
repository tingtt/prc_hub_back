package user

import "errors"

// Errors
var (
	ErrUserNotFound = errors.New("user not found")
	ErrNoUpdates    = errors.New("no updates")
)
