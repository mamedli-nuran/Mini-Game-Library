package handler_test

import (
	"bytes"
	"context"
	"mini-game-library/internal/constant"
	"mini-game-library/internal/handler"
	"mini-game-library/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
)

type MockRatingService struct {
	SubmitRatingFunc func(ctx context.Context, userId uuid.UUID, gameId uuid.UUID, ratingValue int) (*models.Rating, error)
	GetRatingsFunc   func(ctx context.Context, gameId uuid.UUID) ([]*models.Rating, error)
}

func (m *MockRatingService) SubmitRating(ctx context.Context, userId uuid.UUID, gameId uuid.UUID, ratingValue int) (*models.Rating, error) {
	if m.SubmitRatingFunc != nil {
		return m.SubmitRatingFunc(ctx, userId, gameId, ratingValue)
	}
	return nil, nil
}
func (m *MockRatingService) GetRatings(ctx context.Context, gameId uuid.UUID) ([]*models.Rating, error) {
	if m.GetRatingsFunc != nil {
		return m.GetRatingsFunc(ctx, gameId)
	}
	return nil, nil
}

func TestRatingGame(t *testing.T) {
	mockSvc := &MockRatingService{
		SubmitRatingFunc: func(ctx context.Context, userId uuid.UUID, gameId uuid.UUID, ratingValue int) (*models.Rating, error) {
			return &models.Rating{
				Id:        uuid.New(),
				UserId:    userId,
				GameId:    gameId,
				Rating:    ratingValue,
				CreatedAt: time.Now(),
			}, nil
		},
	}
	h := handler.NewRatingHandler(mockSvc)

	reqBody := `{"rating": 5}`
	req := httptest.NewRequest(http.MethodPost, "/games/123e4567-e89b-12d3-a456-426614174000/rating", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("gameId", "123e4567-e89b-12d3-a456-426614174000")

	ctx := context.WithValue(req.Context(), constant.UserIDKey, uuid.New())
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	h.SubmitRating(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status %v, got %v", http.StatusOK, res.StatusCode)
	}
}

func TestRejectInvalidRating(t *testing.T) {
	h := handler.NewRatingHandler(&MockRatingService{})

	reqBody := `{"rating": 10}`
	req := httptest.NewRequest(http.MethodPost, "/games/123e4567-e89b-12d3-a456-426614174000/rating", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("gameId", "123e4567-e89b-12d3-a456-426614174000")

	ctx := context.WithValue(req.Context(), constant.UserIDKey, uuid.New())
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	h.SubmitRating(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status %v, got %v", http.StatusBadRequest, res.StatusCode)
	}
}

func TestUpdateExistingRating(t *testing.T) {
	// Our handler does not differentiate between create and update,
	// it uses CreateOrUpdate via the service layer.
	// We can test it exactly like TestRatingGame, ensuring it returns 200 OK.
	mockSvc := &MockRatingService{
		SubmitRatingFunc: func(ctx context.Context, userId uuid.UUID, gameId uuid.UUID, ratingValue int) (*models.Rating, error) {
			return &models.Rating{
				Id:        uuid.New(),
				UserId:    userId,
				GameId:    gameId,
				Rating:    ratingValue,
				CreatedAt: time.Now(), // time could be older here if mock was advanced
			}, nil
		},
	}
	h := handler.NewRatingHandler(mockSvc)

	reqBody := `{"rating": 4}`
	req := httptest.NewRequest(http.MethodPost, "/games/123e4567-e89b-12d3-a456-426614174000/rating", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("gameId", "123e4567-e89b-12d3-a456-426614174000")

	ctx := context.WithValue(req.Context(), constant.UserIDKey, uuid.New())
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	h.SubmitRating(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status %v, got %v", http.StatusOK, res.StatusCode)
	}
}
