package handler

import (
	"context"
	"encoding/json"
	"errors"
	"mini-game-library/internal/apperror"
	"mini-game-library/internal/constant"
	"mini-game-library/internal/dto"
	"mini-game-library/internal/models"
	"net/http"

	"github.com/google/uuid"
)

type LibraryService interface {
	GetUserLibrary(ctx context.Context, userID uuid.UUID) ([]*models.LibraryItem, error)
	AddToLibrary(ctx context.Context, userID uuid.UUID, req dto.AddToLibraryRequest) (*models.LibraryItem, error)
	UpdateLibraryStatus(ctx context.Context, userID, gameID uuid.UUID, req dto.UpdateLibraryStatusRequest) (*models.LibraryItem, error)
	RemoveFromLibrary(ctx context.Context, userID, gameID uuid.UUID) error
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

func (h *LibraryHandler) AddToLibrary(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(constant.UserIDKey).(uuid.UUID)
	if !ok {
		WriteError(w, r, http.StatusUnauthorized, constant.ErrUnauthorized)
		return
	}

	var req dto.AddToLibraryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, r, http.StatusBadRequest, constant.ErrInvalidBody)
		return
	}

	if errs := req.Validate(); len(errs) > 0 {
		WriteError(w, r, http.StatusBadRequest, constant.ErrValidationFailed, errs...)
		return
	}

	item, err := h.svc.AddToLibrary(r.Context(), userID, req)
	if err != nil {
		if errors.Is(err, apperror.ErrLibraryDuplicate) {
			WriteError(w, r, http.StatusConflict, apperror.ErrLibraryDuplicate.Error())
			return
		}
		if errors.Is(err, apperror.ErrGameNotFound) {
			WriteError(w, r, http.StatusNotFound, apperror.ErrGameNotFound.Error())
			return
		}
		WriteError(w, r, http.StatusInternalServerError, constant.ErrInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, dto.NewLibraryItemResponse(item))
}

func (h *LibraryHandler) UpdateLibraryStatus(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(constant.UserIDKey).(uuid.UUID)
	if !ok {
		WriteError(w, r, http.StatusUnauthorized, constant.ErrUnauthorized)
		return
	}

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

	var req dto.UpdateLibraryStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, r, http.StatusBadRequest, constant.ErrInvalidBody)
		return
	}
	req.Sanitize()

	if errs := req.Validate(); len(errs) > 0 {
		WriteError(w, r, http.StatusBadRequest, constant.ErrValidationFailed, errs...)
		return
	}

	item, err := h.svc.UpdateLibraryStatus(r.Context(), userID, gameID, req)
	if err != nil {
		if errors.Is(err, apperror.ErrLibraryNotFound) {
			WriteError(w, r, http.StatusNotFound, apperror.ErrLibraryNotFound.Error())
			return
		}
		WriteError(w, r, http.StatusInternalServerError, constant.ErrInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, dto.NewLibraryItemResponse(item))
}

func (h *LibraryHandler) RemoveFromLibrary(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(constant.UserIDKey).(uuid.UUID)
	if !ok {
		WriteError(w, r, http.StatusUnauthorized, constant.ErrUnauthorized)
		return
	}

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

	err = h.svc.RemoveFromLibrary(r.Context(), userID, gameID)
	if err != nil {
		if errors.Is(err, apperror.ErrLibraryNotFound) {
			WriteError(w, r, http.StatusNotFound, apperror.ErrLibraryNotFound.Error())
			return
		}
		WriteError(w, r, http.StatusInternalServerError, constant.ErrInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
