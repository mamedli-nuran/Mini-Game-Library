package handler

import (
	"context"
	"encoding/json"
	"errors"
	"mini-game-library/apperror"
	"mini-game-library/constant"
	"mini-game-library/dto"
	"mini-game-library/models"
	"net/http"
)

type UserService interface {
	RegisterUser(ctx context.Context, request dto.AuthRequest) (*models.User, error)
}
type UserHandler struct {
	svc UserService
}

func NewUserHandler(svc UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req dto.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, r, http.StatusBadRequest, constant.ErrInvalidBody)
		return
	}

	req.Sanitize()
	if errorDetails := req.Validate(); len(errorDetails) > 0 {
		WriteError(w, r, http.StatusUnprocessableEntity, constant.ErrValidationFailed, errorDetails...)
		return
	}

	user, err := h.svc.RegisterUser(r.Context(), req)

	if err != nil {
		if errors.Is(err, apperror.ErrRegisterUser) {
			WriteError(w, r, http.StatusConflict, constant.ErrUserAlreadyExists)
		} else {
			WriteError(w, r, http.StatusInternalServerError, constant.ErrInternalServerError)
		}
		return
	}
	writeJSON(w, http.StatusCreated, user)

}
