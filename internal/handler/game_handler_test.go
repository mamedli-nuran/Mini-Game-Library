package handler_test

import (
	"bytes"
	"context"
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

type MockGameService struct {
	FindGamesFunc    func(ctx context.Context, filter *service.GameFilter) ([]*models.Game, int, error)
	FindGameByIdFunc func(ctx context.Context, id uuid.UUID) (*models.Game, error)
	CreateGameFunc   func(ctx context.Context, req dto.CreateGameRequest) (*models.Game, error)
	UpdateGameFunc   func(ctx context.Context, id uuid.UUID, req dto.UpdateGameRequest) (*models.Game, error)
	DeleteGameFunc   func(ctx context.Context, id uuid.UUID) error
}

func (m *MockGameService) FindGames(ctx context.Context, filter *service.GameFilter) ([]*models.Game, int, error) {
	if m.FindGamesFunc != nil {
		return m.FindGamesFunc(ctx, filter)
	}
	return nil, 0, nil
}
func (m *MockGameService) FindGameById(ctx context.Context, id uuid.UUID) (*models.Game, error) {
	if m.FindGameByIdFunc != nil {
		return m.FindGameByIdFunc(ctx, id)
	}
	return nil, nil
}
func (m *MockGameService) CreateGame(ctx context.Context, req dto.CreateGameRequest) (*models.Game, error) {
	if m.CreateGameFunc != nil {
		return m.CreateGameFunc(ctx, req)
	}
	return nil, nil
}
func (m *MockGameService) UpdateGame(ctx context.Context, id uuid.UUID, req dto.UpdateGameRequest) (*models.Game, error) {
	if m.UpdateGameFunc != nil {
		return m.UpdateGameFunc(ctx, id, req)
	}
	return nil, nil
}
func (m *MockGameService) DeleteGame(ctx context.Context, id uuid.UUID) error {
	if m.DeleteGameFunc != nil {
		return m.DeleteGameFunc(ctx, id)
	}
	return nil
}

func TestCreateGame(t *testing.T) {
	mockSvc := &MockGameService{
		CreateGameFunc: func(ctx context.Context, req dto.CreateGameRequest) (*models.Game, error) {
			return &models.Game{
				Id:          uuid.New(),
				Title:       req.Title,
				Description: req.Description,
				Genre:       models.Genre(req.Genre),
				ReleaseYear: req.ReleaseYear,
				CreatedAt:   time.Now(),
			}, nil
		},
	}
	h := handler.NewGameHandler(mockSvc, config.Config{CurrentYear: time.Now().Year()})

	reqBody := `{"title": "The Witcher 3", "description": "Great game.", "genre": "RPG", "release_year": 2015}`
	req := httptest.NewRequest(http.MethodPost, "/games", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.CreateGame(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusCreated {
		t.Errorf("expected status %v, got %v", http.StatusCreated, res.StatusCode)
	}
}

func TestGetGames(t *testing.T) {
	mockSvc := &MockGameService{
		FindGamesFunc: func(ctx context.Context, filter *service.GameFilter) ([]*models.Game, int, error) {
			return []*models.Game{
				{
					Id:          uuid.New(),
					Title:       "Test Game",
					Description: "Desc",
					Genre:       models.GenreRPG,
					ReleaseYear: 2024,
					CreatedAt:   time.Now(),
				},
			}, 1, nil
		},
	}
	h := handler.NewGameHandler(mockSvc, config.Config{CurrentYear: time.Now().Year()})

	req := httptest.NewRequest(http.MethodGet, "/games", nil)
	w := httptest.NewRecorder()

	h.GetGames(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status %v, got %v", http.StatusOK, res.StatusCode)
	}
}
