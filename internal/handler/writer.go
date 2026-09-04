package handler

import (
	"encoding/json"
	"mini-game-library/internal/dto"
	"net/http"
	"time"
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if v != nil {
		json.NewEncoder(w).Encode(v)
	}
}

func WriteError(w http.ResponseWriter, r *http.Request, status int, message string, details ...dto.ErrorDetail) {
	resp := dto.ErrorResponse{
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    status,
		Error:     http.StatusText(status),
		Message:   message,
		Path:      r.URL.Path,
		Details:   details,
	}
	writeJSON(w, status, resp)
}
