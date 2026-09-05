package apperror

import "errors"

var (
	ErrInternalServerError = "Internal Server Error"
	ErrRegisterUser        = errors.New("failed to create user")
	ErrUserDuplicate       = errors.New("user already exists")
	ErrUserFind            = errors.New("can not find user in database")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrLoadingUserId       = errors.New("user id can not be load from application context")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrInvalidGenre        = errors.New("invalid genre")
	ErrInvalidYear         = errors.New("invalid release year, make sure you enter valid year")
)
