package service

import (
	"context"
	"fmt"
	"mini-game-library/internal/dto"
	"mini-game-library/internal/models"

	"github.com/google/uuid"
)

type LibraryRepository interface {
	GetUserLibrary(ctx context.Context, userID uuid.UUID) ([]*models.LibraryItem, error)
	AddToLibrary(ctx context.Context, item *models.LibraryItem) error
}

type LibraryService struct {
	repo LibraryRepository
}

func NewLibraryService(repo LibraryRepository) *LibraryService {
	return &LibraryService{repo: repo}
}

func (s *LibraryService) GetUserLibrary(ctx context.Context, userID uuid.UUID) ([]*models.LibraryItem, error) {
	return s.repo.GetUserLibrary(ctx, userID)
}

func (s *LibraryService) AddToLibrary(ctx context.Context, userID uuid.UUID, req dto.AddToLibraryRequest) (*models.LibraryItem, error) {
	gameID, err := uuid.Parse(req.GameId)
	if err != nil {
		return nil, fmt.Errorf("invalid game id")
	}

	item := &models.LibraryItem{
		Id:     uuid.Must(uuid.NewV7()),
		UserId: userID,
		GameId: gameID,
		Status: models.LibraryStatus(req.Status),
	}

	if err := s.repo.AddToLibrary(ctx, item); err != nil {
		return nil, err
	}

	return item, nil
}
