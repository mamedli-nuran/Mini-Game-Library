package service

import (
	"context"
	"mini-game-library/internal/dto"
	"mini-game-library/internal/models"
	"strings"

	"github.com/google/uuid"
)

type GameRepository interface {
	FindGames(ctx context.Context, filter *GameFilter) ([]*models.Game, int, error)
	FindGameById(ctx context.Context, id uuid.UUID) (*models.Game, error)
	CreateGame(ctx context.Context, game *models.Game) error
	UpdateGame(ctx context.Context, game *models.Game) error
}

type GameService struct {
	repo GameRepository
}

func NewGameService(repo GameRepository) *GameService {
	return &GameService{
		repo: repo,
	}
}

type GameFilter struct {
	Genre       string
	ReleaseYear int
	Search      string
	Page        int
	Limit       int
}

func (s *GameService) FindGames(ctx context.Context, filter *GameFilter) ([]*models.Game, int, error) {
	games, total, err := s.repo.FindGames(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return games, total, nil
}

func (s *GameService) FindGameById(ctx context.Context, id uuid.UUID) (*models.Game, error) {
	game, err := s.repo.FindGameById(ctx, id)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func (s *GameService) CreateGame(ctx context.Context, req dto.CreateGameRequest) (*models.Game, error) {
	game := &models.Game{
		Id:          uuid.New(),
		Title:       req.Title,
		Description: req.Description,
		Genre:       models.Genre(req.Genre),
		ReleaseYear: req.ReleaseYear,
	}

	err := s.repo.CreateGame(ctx, game)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func (s *GameService) UpdateGame(ctx context.Context, id uuid.UUID, req dto.UpdateGameRequest) (*models.Game, error) {
	game, err := s.repo.FindGameById(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Title != nil {
		game.Title = *req.Title
	}
	if req.Description != nil {
		game.Description = *req.Description
	}
	if req.Genre != nil {
		game.Genre = models.Genre(strings.ToUpper(*req.Genre))
	}
	if req.ReleaseYear != nil {
		game.ReleaseYear = *req.ReleaseYear
	}

	err = s.repo.UpdateGame(ctx, game)
	if err != nil {
		return nil, err
	}

	return game, nil
}
