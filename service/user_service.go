package service

import (
	"context"
	"mini-game-library/dto"
	"mini-game-library/models"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	RegisterUser(ctx context.Context, user models.User, hash []byte) error
}

type UserService struct {
	repo          UserRepo
	jwtSecret     string
	accessExpire  time.Duration
	refreshExpire time.Duration
}

func NewUserService(repo UserRepo, jwtSecret string, accessExpire, refreshExpire time.Duration) *UserService {
	return &UserService{
		repo:          repo,
		jwtSecret:     jwtSecret,
		accessExpire:  accessExpire,
		refreshExpire: refreshExpire,
	}
}

func (s UserService) RegisterUser(ctx context.Context, request dto.AuthRequest) (*models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Id:       uuid.New(),
		Username: request.Username,
		Email:    request.Email,
	}

	err = s.repo.RegisterUser(ctx, user, hash)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
