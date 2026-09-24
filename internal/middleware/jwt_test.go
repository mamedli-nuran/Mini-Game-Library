package middleware_test

import (
	"mini-game-library/internal/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInvalidToken(t *testing.T) {
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	jwtSecret := "my-secret"
	mw := middleware.JWTMiddleware(jwtSecret, mockHandler)

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer invalid-token")
	w := httptest.NewRecorder()

	mw.ServeHTTP(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("expected status %v, got %v", http.StatusUnauthorized, res.StatusCode)
	}
}
