package service

import (
	"context"
	"mini-game-library/internal/models"
)

type GameRepository interface {
	FindGames(ctx context.Context, filter GameFilter) ([]*models.Game, error)
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

func (s *GameService) FindGames(ctx context.Context, filter GameFilter) ([]*models.Game, error) {
	games, err := s.repo.FindGames(ctx, filter)
	if err != nil {
		return nil, err
	}

	return games, nil
}
