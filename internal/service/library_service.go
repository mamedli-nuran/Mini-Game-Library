package service

import (
	"context"
	"mini-game-library/internal/models"

	"github.com/google/uuid"
)

type LibraryRepository interface {
	GetUserLibrary(ctx context.Context, userID uuid.UUID) ([]*models.LibraryItem, error)
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
