package handler

import (
	"context"
	"mini-game-library/internal/constant"
	"mini-game-library/internal/models"
	"mini-game-library/internal/service"
	"net/http"
	"strconv"
)

type GameService interface {
	FindGames(ctx context.Context, filter service.GameFilter) ([]*models.Game, error)
}

type GameHandler struct {
	svc GameService
}

func NewGameHandler(svc GameService) *GameHandler {
	return &GameHandler{
		svc: svc,
	}
}

func (h *GameHandler) GetGames(w http.ResponseWriter, r *http.Request) {
	genreParam := r.URL.Query().Get("genre")
	if genreParam != "" {
		genre := models.Genre(genreParam)
		if err := genre.Validate(); err != nil {
			WriteError(w, r, http.StatusBadRequest, constant.ErrInvalidGenre)
			return
		}
	}

	releaseYearStr := r.URL.Query().Get("releaseYear")
	releaseYear := 0
	if releaseYearStr != "" {
		var err error
		releaseYear, err = strconv.Atoi(releaseYearStr)
		if err != nil {
			WriteError(w, r, http.StatusBadRequest, constant.ErrInvalidReleaseYear)
			return
		}
	}

	search := r.URL.Query().Get("search")

	filter := service.GameFilter{
		Genre:       genreParam,
		ReleaseYear: releaseYear,
		Search:      search,
	}

	games, err := h.svc.FindGames(r.Context(), filter)
	if err != nil {
		return
	}

	writeJSON(w, http.StatusOK, games)

}
