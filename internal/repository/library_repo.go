package repository

import (
	"context"
	"mini-game-library/internal/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LibraryRepository struct {
	pool *pgxpool.Pool
}

func NewLibraryRepository(pool *pgxpool.Pool) *LibraryRepository {
	return &LibraryRepository{pool: pool}
}

func (r *LibraryRepository) GetUserLibrary(ctx context.Context, userID uuid.UUID) ([]*models.LibraryItem, error) {
	sql := `SELECT id, user_id, game_id, status, added_at FROM library_items WHERE user_id=$1 ORDER BY added_at DESC`
	rows, err := r.pool.Query(ctx, sql, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*models.LibraryItem
	for rows.Next() {
		var item models.LibraryItem
		if err := rows.Scan(&item.Id, &item.UserId, &item.GameId, &item.Status, &item.AddedAt); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
