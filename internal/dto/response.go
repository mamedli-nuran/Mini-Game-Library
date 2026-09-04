package dto

import (
	"mini-game-library/internal/models"
	"time"

	"github.com/google/uuid"
)

type ErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Timestamp string        `json:"timestamp"`
	Status    int           `json:"status"`
	Error     string        `json:"error"`
	Message   string        `json:"message"`
	Path      string        `json:"path"`
	Details   []ErrorDetail `json:"details,omitempty"`
}

type RegisterResponse struct {
	Id uuid.UUID `json:"id"`
}

func NewRegisterResponse(user *models.User) RegisterResponse {
	return RegisterResponse{
		Id: user.Id,
	}
}

type UserResponse struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUserResponse(user *models.User) UserResponse {
	return UserResponse{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
