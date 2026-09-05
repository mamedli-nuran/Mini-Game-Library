package repository

import (
	"context"
	"errors"
	"fmt"
	"mini-game-library/internal/apperror"
	"mini-game-library/internal/models"
	"mini-game-library/internal/service"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GameRepository struct {
	pool *pgxpool.Pool
}

func NewGameRepository(pool *pgxpool.Pool) *GameRepository {
	return &GameRepository{
		pool: pool,
	}
}

func (r GameRepository) FindGames(ctx context.Context, filter *service.GameFilter) ([]*models.Game, error) {
	var (
		sql   = `SELECT id, title, description, genre, release_year, created_at FROM games WHERE TRUE`
		args  []interface{}
		argsN = 1
	)

	if filter.Genre != "" {
		sql += fmt.Sprintf(" AND genre = $%d", argsN)
		args = append(args, filter.Genre)
		argsN++
	}

	if filter.ReleaseYear != 0 {
		sql += fmt.Sprintf(" AND release_year = $%d", argsN)
		args = append(args, filter.ReleaseYear)
		argsN++
	}

	if filter.Search != "" {
		sql += fmt.Sprintf(" AND title ILIKE $%d", argsN)
		args = append(args, escapeLike(filter.Search)+"%")
		argsN++
	}
	rows, err := r.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	var games []*models.Game
	for rows.Next() {
		var game models.Game
		err = rows.Scan(&game.Id, &game.Title, &game.Description, &game.Genre, &game.ReleaseYear, &game.CreatedAt)
		if err != nil {
			return nil, err
		}
		games = append(games, &game)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return games, nil
}

func escapeLike(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `%`, `\%`)
	s = strings.ReplaceAll(s, `_`, `\_`)
	return s
}

func (r GameRepository) FindGameById(ctx context.Context, id uuid.UUID) (*models.Game, error) {
	sql := `SELECT id, title, description, genre, release_year, created_at FROM games WHERE id=$1`

	var game models.Game
	err := r.pool.QueryRow(ctx, sql, id).
		Scan(&game.Id, &game.Title, &game.Description, &game.Genre, &game.ReleaseYear, &game.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperror.ErrGameNotFound
		}
		return nil, fmt.Errorf("%w: %w", apperror.ErrGetGame, err)
	}
	return &game, nil
}
