package handler

import (
	"context"
	"errors"
	"mini-game-library/internal/apperror"
	"mini-game-library/internal/config"
	"mini-game-library/internal/constant"
	"mini-game-library/internal/dto"
	"mini-game-library/internal/models"
	"mini-game-library/internal/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type GameService interface {
	FindGames(ctx context.Context, filter *service.GameFilter) ([]*models.Game, error)
	FindGameById(ctx context.Context, id uuid.UUID) (*models.Game, error)
}

type GameHandler struct {
	svc GameService
	cfg config.Config
}

func NewGameHandler(svc GameService, cfg config.Config) *GameHandler {
	return &GameHandler{
		svc: svc,
		cfg: cfg,
	}
}

func (h *GameHandler) GetGames(w http.ResponseWriter, r *http.Request) {
	filter, err := h.parseGameFilter(r)
	if err != nil {
		WriteError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	games, err := h.svc.FindGames(r.Context(), filter)
	if err != nil {
		return
	}

	var response []dto.GameResponse
	for _, game := range games {
		response = append(response, dto.NewGameResponse(game))
	}

	writeJSON(w, http.StatusOK, response)
}

func (h *GameHandler) parseGameFilter(r *http.Request) (*service.GameFilter, error) {
	genreParam := r.URL.Query().Get("genre")
	if genreParam != "" {
		genreParam = strings.ToUpper(genreParam)
		genre := models.Genre(genreParam)
		if err := genre.Validate(); err != nil {
			return nil, apperror.ErrInvalidGenre
		}
	}

	releaseYearStr := r.URL.Query().Get("releaseYear")
	releaseYear := 0
	if releaseYearStr != "" {
		var err error
		releaseYear, err = strconv.Atoi(releaseYearStr)
		if err != nil {
			return nil, apperror.ErrInvalidYear
		}
		if releaseYear > h.cfg.CurrentYear || releaseYear < 1900 {
			return nil, apperror.ErrInvalidYear
		}
	}

	search := r.URL.Query().Get("search")

	return &service.GameFilter{
		Genre:       genreParam,
		ReleaseYear: releaseYear,
		Search:      search,
	}, nil
}

func (h *GameHandler) GetGameByID(w http.ResponseWriter, r *http.Request) {
	gameIDStr := r.PathValue("id")
	if gameIDStr == "" {
		WriteError(w, r, http.StatusBadRequest, constant.ErrMissingId)
		return
	}

	gameID, err := uuid.Parse(gameIDStr)
	if err != nil {
		WriteError(w, r, http.StatusBadRequest, constant.ErrParseId)
		return
	}
	game, err := h.svc.FindGameById(r.Context(), gameID)
	if err != nil {
		if errors.Is(err, apperror.ErrGameNotFound) {
			WriteError(w, r, http.StatusNotFound, err.Error())
			return
		}
		WriteError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, dto.NewGameResponse(game))

}
