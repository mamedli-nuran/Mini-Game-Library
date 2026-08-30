package dto

import (
	"mini-game-library/models"

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

type UserResponse struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

func NewUserResponse(user *models.User) UserResponse {
	return UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}
}
