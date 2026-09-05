package service

import (
	"context"
	"mini-game-library/internal/models"

	"github.com/google/uuid"
)

type GameRepository interface {
	FindGames(ctx context.Context, filter *GameFilter) ([]*models.Game, error)
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
}

func (s *GameService) FindGames(ctx context.Context, filter *GameFilter) ([]*models.Game, error) {
	games, err := s.repo.FindGames(ctx, filter)
	if err != nil {
		return nil, err
	}

	return games, nil
}

func (s *GameService) FindGameById(ctx context.Context, id uuid.UUID) (*models.Game, error) {
	game, err := s.repo.FindGameById(ctx, id)
	if err != nil {
		return nil, err
	}

	return game, nil
}
