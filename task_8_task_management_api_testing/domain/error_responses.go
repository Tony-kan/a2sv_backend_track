package domain

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserExists         = errors.New("user already exists")
	ErrInvalidUser        = errors.New("invalid user data")
	ErrInvalidUserID      = errors.New("invalid user ID")
	ErrInvalidEmail       = errors.New("invalid email format")
	ErrInvalidPassword    = errors.New("invalid password format")
	ErrInvalidUsername    = errors.New("invalid username format")
	ErrEmptyPassword      = errors.New("password cannot be empty")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrAccessDenied       = errors.New("access denied")
	ErrTokenExpired       = errors.New("token expired")
	ErrTokenInvalid       = errors.New("invalid token")
)

var (
	ErrTaskNotFound  = errors.New("task not found")
	ErrTaskExists    = errors.New("task already exists")
	ErrInvalidTask   = errors.New("invalid task data")
	ErrInvalidTaskID = errors.New("invalid task ID")
)
