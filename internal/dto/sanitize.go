package dto

import "strings"

func (r *RegisterRequest) Sanitize() {
	r.Username = strings.ToLower(strings.TrimSpace(r.Username))
	r.Email = strings.ToLower(strings.TrimSpace(r.Email))
	r.Password = strings.TrimSpace(r.Password)
}

func (r *LoginRequest) Sanitize() {
	r.Identifier = strings.ToLower(strings.TrimSpace(r.Identifier))
	r.Password = strings.TrimSpace(r.Password)
}

func (r *CreateGameRequest) Sanitize() {
	r.Title = strings.TrimSpace(r.Title)
	r.Description = strings.TrimSpace(r.Description)
	r.Genre = strings.TrimSpace(r.Genre)
}
