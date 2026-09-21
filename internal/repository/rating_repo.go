package repository

import (
	"context"
	"fmt"
	"mini-game-library/internal/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RatingRepository struct {
	pool *pgxpool.Pool
}

func NewRatingRepository(pool *pgxpool.Pool) *RatingRepository {
	return &RatingRepository{
		pool: pool,
	}
}

func (r *RatingRepository) CreateOrUpdateRating(ctx context.Context, rating *models.Rating) error {
	sql := `
		INSERT INTO ratings (id, user_id, game_id, rating, created_at)
		VALUES ($1, $2, $3, $4, NOW())
		ON CONFLICT (user_id, game_id)
		DO UPDATE SET rating = EXCLUDED.rating
		RETURNING created_at, id
	`
	err := r.pool.QueryRow(ctx, sql, uuid.New(), rating.UserId, rating.GameId, rating.Rating).Scan(&rating.CreatedAt, &rating.Id)
	if err != nil {
		return fmt.Errorf("failed to upsert rating: %w", err)
	}
	return nil
}
