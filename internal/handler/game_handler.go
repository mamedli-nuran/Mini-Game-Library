package handler

import (
	"context"
	"encoding/json"
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
	FindGames(ctx context.Context, filter *service.GameFilter) ([]*models.Game, int, error)
	FindGameById(ctx context.Context, id uuid.UUID) (*models.Game, error)
	CreateGame(ctx context.Context, title, description, genre string, releaseYear int) (*models.Game, error)
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
	games, total, err := h.svc.FindGames(r.Context(), filter)
	if err != nil {
		return
	}

	response := dto.NewGamesResponse(games, total)

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

	pageStr := r.URL.Query().Get("page")
	page := 1
	if pageStr != "" {
		p, err := strconv.Atoi(pageStr)
		if err != nil || p <= 0 {
			return nil, apperror.ErrInvalidPage
		}
		page = p
	}

	limitStr := r.URL.Query().Get("limit")
	limit := 10
	if limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err != nil || l <= 0 {
			return nil, apperror.ErrInvalidLimit
		}
		limit = l
	}
	if limit > 20 {
		limit = 20
	}

	return &service.GameFilter{
		Genre:       genreParam,
		ReleaseYear: releaseYear,
		Search:      search,
		Page:        page,
		Limit:       limit,
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

func (h *GameHandler) CreateGame(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateGameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, r, http.StatusBadRequest, constant.ErrInvalidBody)
		return
	}
	req.Sanitize()

	if errs := req.Validate(); len(errs) > 0 {
		WriteError(w, r, http.StatusBadRequest, constant.ErrValidationFailed, errs...)
		return
	}

	game, err := h.svc.CreateGame(r.Context(), req.Title, req.Description, req.Genre, req.ReleaseYear)
	if err != nil {
		if errors.Is(err, apperror.ErrGameDuplicate) {
			WriteError(w, r, http.StatusConflict, err.Error())
		} else {
			WriteError(w, r, http.StatusInternalServerError, constant.ErrInternalServerError)
		}
		return
	}

	writeJSON(w, http.StatusCreated, dto.NewGameResponse(game))
}
