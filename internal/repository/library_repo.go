package repository

import (
	"context"
	"errors"
	"fmt"
	"mini-game-library/internal/apperror"
	"mini-game-library/internal/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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

func (r *LibraryRepository) AddToLibrary(ctx context.Context, item *models.LibraryItem) error {
	sql := `INSERT INTO library_items (id, user_id, game_id, status) VALUES ($1, $2, $3, $4) RETURNING added_at`
	err := r.pool.QueryRow(ctx, sql, item.Id, item.UserId, item.GameId, item.Status).Scan(&item.AddedAt)
	if err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
			if pgErr.Code == "23505" {
				return apperror.ErrLibraryDuplicate
			}
			if pgErr.Code == "23503" {
				return apperror.ErrGameNotFound
			}
		}
		return fmt.Errorf("failed to add to library: %w", err)
	}
	return nil
}

func (r *LibraryRepository) UpdateLibraryStatus(ctx context.Context, userID, gameID uuid.UUID, status models.LibraryStatus) (*models.LibraryItem, error) {
	sql := `UPDATE library_items SET status = $1 WHERE user_id = $2 AND game_id = $3 RETURNING id, user_id, game_id, status, added_at`
	var item models.LibraryItem
	err := r.pool.QueryRow(ctx, sql, status, userID, gameID).Scan(&item.Id, &item.UserId, &item.GameId, &item.Status, &item.AddedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperror.ErrLibraryNotFound
		}
		return nil, fmt.Errorf("failed to update library status: %w", err)
	}
	return &item, nil
}

func (r *LibraryRepository) RemoveFromLibrary(ctx context.Context, userID, gameID uuid.UUID) error {
	sql := `DELETE FROM library_items WHERE user_id = $1 AND game_id = $2`
	commandTag, err := r.pool.Exec(ctx, sql, userID, gameID)
	if err != nil {
		return fmt.Errorf("failed to remove from library: %w", err)
	}
	if commandTag.RowsAffected() == 0 {
		return apperror.ErrLibraryNotFound
	}
	return nil
}
