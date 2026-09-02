package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"mini-game-library/apperror"
	"mini-game-library/config"
	"mini-game-library/constant"
	"mini-game-library/dto"
	"mini-game-library/models"
	"net/mail"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	RegisterUser(ctx context.Context, user models.User, hash []byte) error
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	FindUserByUsername(ctx context.Context, username string) (*models.User, error)
	FindUserById(ctx context.Context, userID uuid.UUID) (*models.User, error)
}

type UserService struct {
	repo UserRepo
	cfg  config.Config
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

func NewUserService(repo UserRepo, cfg config.Config) *UserService {
	return &UserService{
		repo: repo,
		cfg:  cfg,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, request dto.RegisterRequest) (*models.User, error) {
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

func (s *UserService) LoginUser(ctx context.Context, request dto.LoginRequest) (*TokenPair, error) {
	isEmail := false
	_, err := mail.ParseAddress(request.Identifier)
	if err == nil {
		isEmail = true
	}

	var user *models.User
	if isEmail {
		user, err = s.repo.FindUserByEmail(ctx, request.Identifier)
	} else {
		user, err = s.repo.FindUserByUsername(ctx, request.Identifier)
	}

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperror.ErrUserFind
		}
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return nil, apperror.ErrInvalidCredentials
	}

	accessToken, err := s.generateAccessToken(user.Id)
	if err != nil {
		return nil, err
	}

	refreshToken, err := GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}

func (s *UserService) GetMeInfo(ctx context.Context) (*models.User, error) {
	userID, ok := ctx.Value(constant.UserIDKey).(uuid.UUID)
	if !ok {
		return nil, apperror.ErrLoadingUserId
	}

	user, err := s.repo.FindUserById(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) generateAccessToken(id uuid.UUID) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(now.Add(s.cfg.AccessTokenExpireSeconds)),
		Subject:   id.String(),
		IssuedAt:  jwt.NewNumericDate(now),
	})
	return token.SignedString([]byte(s.cfg.JWTSecret))
}

func GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
