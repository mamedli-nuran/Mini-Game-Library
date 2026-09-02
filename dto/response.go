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

type RegisterResponse struct {
	Id uuid.UUID `json:"id"`
}

func NewRegisterResponse(user *models.User) RegisterResponse {
	return RegisterResponse{
		Id: user.Id,
	}
}
