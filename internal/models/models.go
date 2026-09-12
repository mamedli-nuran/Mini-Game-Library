package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Game struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Genre       Genre     `json:"genre"`
	ReleaseYear int       `json:"release_year"`
	CreatedAt   time.Time `json:"created_at"`
}

type RefreshToken struct {
	Id          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"user_id"`
	HashedToken string    `json:"token"`
	IsActive    bool      `json:"is_active"`
	ExpiresAt   time.Time `json:"expires_at"`
	CreatedAt   time.Time `json:"created_at"`
}
