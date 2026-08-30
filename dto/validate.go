package dto

import (
	"mini-game-library/constant"
	"net/mail"
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
