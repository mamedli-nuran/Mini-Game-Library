package service

import (
	"context"
	"mini-game-library/internal/models"

	"github.com/google/uuid"
)

type RatingRepo interface {
	CreateOrUpdateRating(ctx context.Context, rating *models.Rating) error
}

type RatingService struct {
	repo RatingRepo
}

func NewRatingService(repo RatingRepo) *RatingService {
	return &RatingService{
		repo: repo,
	}
}

func (s *RatingService) SubmitRating(ctx context.Context, userId uuid.UUID, gameId uuid.UUID, ratingValue int) (*models.Rating, error) {
	rating := &models.Rating{
		UserId: userId,
		GameId: gameId,
		Rating: ratingValue,
	}
	err := s.repo.CreateOrUpdateRating(ctx, rating)
	if err != nil {
		return nil, err
	}
	return rating, nil
}
