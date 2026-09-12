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
	ErrInvalidPage         = errors.New("invalid page parameter")
	ErrInvalidLimit        = errors.New("invalid limit parameter")

	ErrGetGame          = errors.New("failed to get user")
	ErrGameNotFound     = errors.New("game not found")
	ErrGameDuplicate    = errors.New("game with this title already exists")
	ErrLibraryDuplicate = errors.New("game already exists in library")
	ErrLibraryNotFound  = errors.New("game not found in library")
)
