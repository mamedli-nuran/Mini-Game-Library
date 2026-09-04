package handler

import "net/http"

type GameService interface {
}

type GameHandler struct {
	svc GameService
}

func NewGameHandler(svc GameService) *GameHandler {
	return &GameHandler{
		svc: svc,
	}
}

func (h *GameHandler) GetGames(w http.ResponseWriter, r *http.Request) {

}
