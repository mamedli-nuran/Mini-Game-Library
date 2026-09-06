package service

import (
	"context"
	"mini-game-library/internal/models"

	"github.com/google/uuid"
)

type GameRepository interface {
	FindGames(ctx context.Context, filter *GameFilter) ([]*models.Game, int, error)
	FindGameById(ctx context.Context, id uuid.UUID) (*models.Game, error)
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
