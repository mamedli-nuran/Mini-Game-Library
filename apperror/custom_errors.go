package apperror

import "errors"

var (
	ErrRegisterUser  = errors.New("failed to create user")
	ErrUserDuplicate = errors.New("user already exists")
)
