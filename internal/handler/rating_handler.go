package handler

import (
	"context"
	"encoding/json"
	"mini-game-library/internal/constant"
	"mini-game-library/internal/dto"
	"mini-game-library/internal/models"
	"net/http"

	"github.com/google/uuid"
)

type RatingService interface {
	SubmitRating(ctx context.Context, userId uuid.UUID, gameId uuid.UUID, ratingValue int) (*models.Rating, error)
}

type RatingHandler struct {
	svc RatingService
}

func NewRatingHandler(svc RatingService) *RatingHandler {
	return &RatingHandler{
		svc: svc,
	}
}

func (h *RatingHandler) SubmitRating(w http.ResponseWriter, r *http.Request) {
	gameIDStr := r.PathValue("gameId")
	if gameIDStr == "" {
		WriteError(w, r, http.StatusBadRequest, constant.ErrMissingId)
		return
	}

	gameID, err := uuid.Parse(gameIDStr)
	if err != nil {
		WriteError(w, r, http.StatusBadRequest, constant.ErrParseId)
		return
	}

	userID, ok := r.Context().Value(constant.UserIDKey).(uuid.UUID)
	if !ok {
		WriteError(w, r, http.StatusUnauthorized, constant.ErrUnauthorized)
		return
	}

	var req dto.RatingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, r, http.StatusBadRequest, constant.ErrInvalidBody)
		return
	}

	if errs := req.Validate(); len(errs) > 0 {
		WriteError(w, r, http.StatusBadRequest, constant.ErrValidationFailed, errs...)
		return
	}

	rating, err := h.svc.SubmitRating(r.Context(), userID, gameID, req.Rating)
	if err != nil {
		WriteError(w, r, http.StatusInternalServerError, constant.ErrInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, rating)
}
