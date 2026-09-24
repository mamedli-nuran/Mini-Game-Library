package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"mini-game-library/internal/config"
	"mini-game-library/internal/dto"
	"mini-game-library/internal/handler"
	"mini-game-library/internal/models"
	"mini-game-library/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
)

type MockUserService struct {
	RegisterUserFunc  func(ctx context.Context, req dto.RegisterRequest) (*models.User, error)
	LoginUserFunc     func(ctx context.Context, req dto.LoginRequest) (*service.TokenPair, error)
	RefreshTokensFunc func(ctx context.Context, refreshToken string) (*service.TokenPair, error)
	GetMeInfoFunc     func(ctx context.Context) (*models.User, error)
}

func (m *MockUserService) RegisterUser(ctx context.Context, req dto.RegisterRequest) (*models.User, error) {
	if m.RegisterUserFunc != nil {
		return m.RegisterUserFunc(ctx, req)
	}
	return nil, nil
}

func (m *MockUserService) LoginUser(ctx context.Context, req dto.LoginRequest) (*service.TokenPair, error) {
	if m.LoginUserFunc != nil {
		return m.LoginUserFunc(ctx, req)
	}
	return nil, nil
}

func (m *MockUserService) RefreshTokens(ctx context.Context, refreshToken string) (*service.TokenPair, error) {
	if m.RefreshTokensFunc != nil {
		return m.RefreshTokensFunc(ctx, refreshToken)
	}
	return nil, nil
}

func (m *MockUserService) GetMeInfo(ctx context.Context) (*models.User, error) {
	if m.GetMeInfoFunc != nil {
		return m.GetMeInfoFunc(ctx)
	}
	return nil, nil
}

func TestRegistration(t *testing.T) {
	mockSvc := &MockUserService{
		RegisterUserFunc: func(ctx context.Context, req dto.RegisterRequest) (*models.User, error) {
			return &models.User{
				Id:        uuid.New(),
				Username:  req.Username,
				Email:     req.Email,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil
		},
	}
	h := handler.NewUserHandler(mockSvc, config.Config{})

	reqBody := `{"username": "testuser", "email": "test@example.com", "password": "password123"}`
	req := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.RegisterUser(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusCreated {
		t.Errorf("expected status %v, got %v", http.StatusCreated, res.StatusCode)
	}
}

func TestLogin(t *testing.T) {
	mockSvc := &MockUserService{
		LoginUserFunc: func(ctx context.Context, req dto.LoginRequest) (*service.TokenPair, error) {
			return &service.TokenPair{AccessToken: "mock-jwt-token", RefreshToken: "mock-refresh-token"}, nil
		},
	}
	h := handler.NewUserHandler(mockSvc, config.Config{})

	reqBody := `{"identifier": "test@example.com", "password": "password123"}`
	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.Login(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status %v, got %v", http.StatusOK, res.StatusCode)
	}

	var response map[string]string
	json.NewDecoder(res.Body).Decode(&response)
	if response["access_token"] != "mock-jwt-token" {
		t.Errorf("expected token 'mock-jwt-token', got '%v'", response["access_token"])
	}
}
