package dto

import (
	"mini-game-library/internal/constant"
	"mini-game-library/internal/models"
	"net/mail"
	"strings"
)

func (r *RegisterRequest) Validate() []ErrorDetail {
	var errors []ErrorDetail
	if r.Username == "" {
		errors = append(errors, ErrorDetail{Field: "username", Message: constant.ErrUsernameRequired})
	} else if usernameLen := len(r.Username); usernameLen < 4 || usernameLen > 20 {
		errors = append(errors, ErrorDetail{Field: "username", Message: constant.ErrUsernameLength})
	}

	if r.Email == "" {
		errors = append(errors, ErrorDetail{Field: "email", Message: constant.ErrEmailRequired})
	} else if _, err := mail.ParseAddress(r.Email); err != nil {
		errors = append(errors, ErrorDetail{Field: "email", Message: constant.ErrInvalidEmail})
	}

	if passwordLen := len(r.Password); passwordLen < 12 || passwordLen > 25 {
		errors = append(errors, ErrorDetail{Field: "password", Message: constant.ErrPasswordLength})
	}
	return errors
}

// todo make normal validation
func (r *LoginRequest) Validate() []ErrorDetail {
	var errors []ErrorDetail

	if r.Identifier == "" {
		errors = append(errors, ErrorDetail{Field: "identifier", Message: "safsafas"})
	}

	return errors
}

func (r *CreateGameRequest) Validate() []ErrorDetail {
	var errors []ErrorDetail

	titleLen := len(r.Title)
	if titleLen < 2 || titleLen > 100 {
		errors = append(errors, ErrorDetail{Field: "title", Message: constant.ErrTitleLength})
	}

	descLen := len(r.Description)
	if descLen < 10 || descLen > 1000 {
		errors = append(errors, ErrorDetail{Field: "description", Message: constant.ErrDescriptionLength})
	}

	genre := models.Genre(strings.ToUpper(r.Genre))
	if err := genre.Validate(); err != nil {
		errors = append(errors, ErrorDetail{Field: "genre", Message: constant.ErrInvalidGenre})
	}

	if r.ReleaseYear < 1900 || r.ReleaseYear > 2100 {
		errors = append(errors, ErrorDetail{Field: "release_year", Message: constant.ErrInvalidReleaseYear})
	}

	return errors
}

func (r *UpdateGameRequest) Validate() []ErrorDetail {
	var errors []ErrorDetail

	if r.Title != nil {
		titleLen := len(*r.Title)
		if titleLen < 2 || titleLen > 100 {
			errors = append(errors, ErrorDetail{Field: "title", Message: constant.ErrTitleLength})
		}
	}

	if r.Description != nil {
		descLen := len(*r.Description)
		if descLen < 10 || descLen > 1000 {
			errors = append(errors, ErrorDetail{Field: "description", Message: constant.ErrDescriptionLength})
		}
	}

	if r.Genre != nil {
		genre := models.Genre(strings.ToUpper(*r.Genre))
		if err := genre.Validate(); err != nil {
			errors = append(errors, ErrorDetail{Field: "genre", Message: constant.ErrInvalidGenre})
		}
	}

	if r.ReleaseYear != nil {
		if *r.ReleaseYear < 1900 || *r.ReleaseYear > 2100 {
			errors = append(errors, ErrorDetail{Field: "release_year", Message: constant.ErrInvalidReleaseYear})
		}
	}

	return errors
}
