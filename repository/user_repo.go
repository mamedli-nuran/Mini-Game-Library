package repository

import (
	"context"
	"fmt"
	"log/slog"
	"mini-game-library/apperror"
	"mini-game-library/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}

func (r UserRepository) RegisterUser(ctx context.Context, user models.User, hash []byte) error {
	sql := `INSERT INTO users (id, username, email, password) 
        	VALUES ($1, $2, $3, $4)`

	_, err := r.pool.Exec(ctx, sql, user.Id, user.Username, user.Email, hash)
	if err != nil {
		slog.Info(err.Error())
		return fmt.Errorf("%w: %w", apperror.ErrRegisterUser, err)
	}
	return nil
}
