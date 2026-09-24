package handler_test

import (
	"bytes"
	"context"
	"mini-game-library/internal/apperror"
	"mini-game-library/internal/constant"
	"mini-game-library/internal/dto"
	"mini-game-library/internal/handler"
	"mini-game-library/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
)

type MockLibraryService struct {
	GetLibraryFunc          func(ctx context.Context, userId uuid.UUID) ([]*models.LibraryItem, error)
	AddToLibraryFunc        func(ctx context.Context, userId uuid.UUID, req dto.AddToLibraryRequest) (*models.LibraryItem, error)
	UpdateLibraryStatusFunc func(ctx context.Context, userId uuid.UUID, gameId uuid.UUID, req dto.UpdateLibraryStatusRequest) (*models.LibraryItem, error)
	RemoveFromLibraryFunc   func(ctx context.Context, userId uuid.UUID, gameId uuid.UUID) error
}

func (m *MockLibraryService) GetUserLibrary(ctx context.Context, userId uuid.UUID) ([]*models.LibraryItem, error) {
	if m.GetLibraryFunc != nil {
		return m.GetLibraryFunc(ctx, userId)
	}
	return nil, nil
}
func (m *MockLibraryService) AddToLibrary(ctx context.Context, userId uuid.UUID, req dto.AddToLibraryRequest) (*models.LibraryItem, error) {
	if m.AddToLibraryFunc != nil {
		return m.AddToLibraryFunc(ctx, userId, req)
	}
	return nil, nil
}
func (m *MockLibraryService) UpdateLibraryStatus(ctx context.Context, userId uuid.UUID, gameId uuid.UUID, req dto.UpdateLibraryStatusRequest) (*models.LibraryItem, error) {
	if m.UpdateLibraryStatusFunc != nil {
		return m.UpdateLibraryStatusFunc(ctx, userId, gameId, req)
	}
	return nil, nil
}
func (m *MockLibraryService) RemoveFromLibrary(ctx context.Context, userId uuid.UUID, gameId uuid.UUID) error {
	if m.RemoveFromLibraryFunc != nil {
		return m.RemoveFromLibraryFunc(ctx, userId, gameId)
	}
	return nil
}

func TestAddGameToLibrary(t *testing.T) {
	mockSvc := &MockLibraryService{
		AddToLibraryFunc: func(ctx context.Context, userId uuid.UUID, req dto.AddToLibraryRequest) (*models.LibraryItem, error) {
			return &models.LibraryItem{
				Id:      uuid.New(),
				UserId:  userId,
				GameId:  uuid.MustParse(req.GameId),
				Status:  models.LibraryStatus(req.Status),
				AddedAt: time.Now(),
			}, nil
		},
	}
	h := handler.NewLibraryHandler(mockSvc)

	reqBody := `{"game_id": "123e4567-e89b-12d3-a456-426614174000", "status": "PLAYING"}`
	req := httptest.NewRequest(http.MethodPost, "/me/library", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Inject user id into context
	ctx := context.WithValue(req.Context(), constant.UserIDKey, uuid.New())
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	h.AddToLibrary(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusCreated {
		t.Errorf("expected status %v, got %v", http.StatusCreated, res.StatusCode)
	}
}

func TestPreventDuplicateLibraryEntry(t *testing.T) {
	mockSvc := &MockLibraryService{
		AddToLibraryFunc: func(ctx context.Context, userId uuid.UUID, req dto.AddToLibraryRequest) (*models.LibraryItem, error) {
			return nil, apperror.ErrLibraryDuplicate
		},
	}
	h := handler.NewLibraryHandler(mockSvc)

	reqBody := `{"game_id": "123e4567-e89b-12d3-a456-426614174000", "status": "PLAYING"}`
	req := httptest.NewRequest(http.MethodPost, "/me/library", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Inject user id into context
	ctx := context.WithValue(req.Context(), constant.UserIDKey, uuid.New())
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	h.AddToLibrary(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusConflict {
		t.Errorf("expected status %v, got %v", http.StatusConflict, res.StatusCode)
	}
}
