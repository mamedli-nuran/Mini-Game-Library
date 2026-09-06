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
	"github.com/jackc/pgx/v5/pgconn"
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

func (r GameRepository) FindGames(ctx context.Context, filter *service.GameFilter) ([]*models.Game, int, error) {
	var (
		baseSQL = `FROM games WHERE TRUE`
		args    []interface{}
		argsN   = 1
	)

	if filter.Genre != "" {
		baseSQL += fmt.Sprintf(" AND genre = $%d", argsN)
		args = append(args, filter.Genre)
		argsN++
	}

	if filter.ReleaseYear != 0 {
		baseSQL += fmt.Sprintf(" AND release_year = $%d", argsN)
		args = append(args, filter.ReleaseYear)
		argsN++
	}

	if filter.Search != "" {
		baseSQL += fmt.Sprintf(" AND title ILIKE $%d", argsN)
		args = append(args, "%"+escapeLike(filter.Search)+"%")
		argsN++
	}

	countSQL := "SELECT COUNT(*) " + baseSQL
	var total int
	err := r.pool.QueryRow(ctx, countSQL, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	querySQL := "SELECT id, title, description, genre, release_year, created_at " + baseSQL
	querySQL += fmt.Sprintf(" ORDER BY id LIMIT $%d OFFSET $%d", argsN, argsN+1)
	queryArgs := append(args, filter.Limit, (filter.Page-1)*filter.Limit)

	rows, err := r.pool.Query(ctx, querySQL, queryArgs...)
	if err != nil {
		return nil, 0, err
	}

	var games []*models.Game
	for rows.Next() {
		var game models.Game
		err = rows.Scan(&game.Id, &game.Title, &game.Description, &game.Genre, &game.ReleaseYear, &game.CreatedAt)
		if err != nil {
			return nil, 0, err
		}
		games = append(games, &game)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return games, total, nil
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

func (r GameRepository) CreateGame(ctx context.Context, game *models.Game) error {
	sql := `INSERT INTO games (id, title, description, genre, release_year) VALUES ($1, $2, $3, $4, $5) RETURNING created_at`
	err := r.pool.QueryRow(ctx, sql, game.Id, game.Title, game.Description, game.Genre, game.ReleaseYear).Scan(&game.CreatedAt)
	if err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
			if pgErr.Code == "23505" {
				return apperror.ErrGameDuplicate
			}
		}
		return fmt.Errorf("failed to create game: %w", err)
	}
	return nil
}
