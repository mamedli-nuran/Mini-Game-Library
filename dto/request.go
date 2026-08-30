package dto

import (
	"mini-game-library/constant"
	"net/mail"
	"strings"
)

type AuthRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *AuthRequest) Sanitize() {
	r.Username = strings.ToLower(strings.TrimSpace(r.Username))
	r.Email = strings.ToLower(strings.TrimSpace(r.Email))
	r.Password = strings.TrimSpace(r.Password)
}

func (r *AuthRequest) Validate() []ErrorDetail {
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

	if pwlen := len(r.Password); pwlen < 12 || pwlen > 25 {
		errors = append(errors, ErrorDetail{Field: "password", Message: constant.ErrPasswordLength})
	}
	return errors
}
