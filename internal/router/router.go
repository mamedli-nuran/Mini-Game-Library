package router

import (
	"mini-game-library/internal/config"
	handler2 "mini-game-library/internal/handler"
	"mini-game-library/internal/middleware"
	"net/http"
)

func Setup(
	userHandler *handler2.UserHandler,
	gameHandler *handler2.GameHandler,
	cfg config.Config,
) *http.ServeMux {
	mux := http.NewServeMux()

	// auth
	mux.HandleFunc("POST /auth/register", userHandler.RegisterUser)
	mux.HandleFunc("POST /auth/login", userHandler.Login)
	mux.HandleFunc("GET /auth/me", middleware.JWTMiddleware(cfg.JWTSecret, userHandler.MeInfo))

	// game
	//todo add pagination
	mux.HandleFunc("GET /games", gameHandler.GetGames)
	mux.HandleFunc("GET /games/{id}", gameHandler.GetGameByID)
	return mux
}
