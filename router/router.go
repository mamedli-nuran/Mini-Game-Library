package router

import (
	"mini-game-library/config"
	"mini-game-library/handler"
	"mini-game-library/middleware"
	"net/http"
)

func Setup(
	userHandler *handler.UserHandler,
	gameHandler *handler.GameHandler,
	cfg config.Config,
) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /auth/register", userHandler.RegisterUser)
	mux.HandleFunc("POST /auth/login", userHandler.Login)
	mux.HandleFunc("GET /auth/me", middleware.JWTMiddleware(cfg.JWTSecret, userHandler.MeInfo))
	return mux
}
