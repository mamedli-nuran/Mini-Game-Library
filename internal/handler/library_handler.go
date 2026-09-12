package handler

import (
	"context"
	"mini-game-library/internal/constant"
	"mini-game-library/internal/dto"
	"mini-game-library/internal/models"
	"net/http"

	"github.com/google/uuid"
)

type LibraryService interface {
	GetUserLibrary(ctx context.Context, userID uuid.UUID) ([]*models.LibraryItem, error)
}

type LibraryHandler struct {
	svc LibraryService
}

func NewLibraryHandler(svc LibraryService) *LibraryHandler {
	return &LibraryHandler{svc: svc}
}

func (h *LibraryHandler) GetLibrary(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(constant.UserIDKey).(uuid.UUID)
	if !ok {
		WriteError(w, r, http.StatusUnauthorized, constant.ErrUnauthorized)
		return
	}

	items, err := h.svc.GetUserLibrary(r.Context(), userID)
	if err != nil {
		WriteError(w, r, http.StatusInternalServerError, constant.ErrInternalServerError)
		return
	}

	var res = make([]dto.LibraryItemResponse, 0)
	for _, item := range items {
		res = append(res, dto.NewLibraryItemResponse(item))
	}

	writeJSON(w, http.StatusOK, res)
}
