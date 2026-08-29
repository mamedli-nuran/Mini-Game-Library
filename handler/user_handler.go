package handler

import (
	"encoding/json"
	"net/http"
)

type UserService interface {
}
type UserHandler struct {
	svc UserService
}

func NewUserHandler(svc UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (h *UserHandler) Test(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode("test data")
	if err != nil {
		return
	}

}
