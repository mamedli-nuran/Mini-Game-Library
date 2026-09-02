package repository

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"mini-game-library/apperror"
	"mini-game-library/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
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

func (r *UserRepository) RegisterUser(ctx context.Context, user models.User, hash []byte) error {
	sql := `INSERT INTO users (id, username, email, password) 
        	VALUES ($1, $2, $3, $4)`

	_, err := r.pool.Exec(ctx, sql, user.Id, user.Username, user.Email, hash)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return apperror.ErrUserDuplicate
			}
		}

		slog.Error("Database error during user registration", "error", err)
		return fmt.Errorf("%w: %w", apperror.ErrRegisterUser, err)
	}
	return nil
}

func (r *UserRepository) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	sql := `SELECT id, username, email, password, created_at, updated_at 
			FROM users WHERE email=$1`

	var user models.User
	var password []byte
	err := r.pool.QueryRow(ctx, sql, email).Scan(
		&user.Id, &user.Username, &user.Email, &password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	user.Password = string(password)
	return &user, nil

}

func (r *UserRepository) FindUserByUsername(ctx context.Context, username string) (*models.User, error) {
	sql := `SELECT id, username, email, password, created_at, updated_at 
			FROM users WHERE username=$1`

	var user models.User
	var password []byte
	err := r.pool.QueryRow(ctx, sql, username).Scan(
		&user.Id, &user.Username, &user.Email, &password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	user.Password = string(password)
	return &user, nil
}

func (r *UserRepository) FindUserById(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	sql := `SELECT id, username, email, created_at, updated_at 
			FROM users WHERE id=$1`

	var user models.User
	err := r.pool.QueryRow(ctx, sql, userID).Scan(
		&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
