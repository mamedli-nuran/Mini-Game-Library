package handler

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
