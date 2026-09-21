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

func (r *RatingRepository) GetRatingsByGameId(ctx context.Context, gameId uuid.UUID) ([]*models.Rating, error) {
	sql := `SELECT id, user_id, game_id, rating, created_at FROM ratings WHERE game_id = $1 ORDER BY created_at DESC`
	rows, err := r.pool.Query(ctx, sql, gameId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ratings []*models.Rating
	for rows.Next() {
		var rating models.Rating
		if err := rows.Scan(&rating.Id, &rating.UserId, &rating.GameId, &rating.Rating, &rating.CreatedAt); err != nil {
			return nil, err
		}
		ratings = append(ratings, &rating)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if ratings == nil {
		ratings = make([]*models.Rating, 0)
	}
	return ratings, nil
}
