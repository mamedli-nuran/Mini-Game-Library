package handler

import (
	"context"
	"mini-game-library/internal/apperror"
	"mini-game-library/internal/config"
	"mini-game-library/internal/models"
	"mini-game-library/internal/service"
	"net/http"
	"strconv"
	"strings"
)

type GameService interface {
	FindGames(ctx context.Context, filter *service.GameFilter) ([]*models.Game, error)
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
	filter, err := filterGame(r, h)
	if err != nil {
		WriteError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	games, err := h.svc.FindGames(r.Context(), filter)
	if err != nil {
		return
	}

	writeJSON(w, http.StatusOK, games)
}

func filterGame(r *http.Request, h *GameHandler) (*service.GameFilter, error) {
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

	filter := service.GameFilter{
		Genre:       genreParam,
		ReleaseYear: releaseYear,
		Search:      search,
	}

	return &filter, nil
}
