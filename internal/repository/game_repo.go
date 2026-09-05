package repository

import (
	"context"
	"fmt"
	"mini-game-library/internal/models"
	"mini-game-library/internal/service"

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
		sql += fmt.Sprintf(" AND search = $%d", argsN)
		args = append(args, filter.Search)
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
