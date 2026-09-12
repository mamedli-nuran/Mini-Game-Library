package router

import (
	"mini-game-library/internal/config"
	"mini-game-library/internal/handler"
	"mini-game-library/internal/middleware"
	"net/http"
)

func Setup(
	userHandler *handler.UserHandler,
	gameHandler *handler.GameHandler,
	libraryHandler *handler.LibraryHandler,
	cfg config.Config,
) *http.ServeMux {
	mux := http.NewServeMux()

	// auth
	mux.HandleFunc("POST /auth/register", userHandler.RegisterUser)
	mux.HandleFunc("POST /auth/login", userHandler.Login)
	mux.HandleFunc("POST /auth/refresh", userHandler.Refresh)
	mux.HandleFunc("GET /auth/me", middleware.JWTMiddleware(cfg.JWTSecret, userHandler.MeInfo))

	// game
	mux.HandleFunc("GET /games", gameHandler.GetGames)
	mux.HandleFunc("GET /games/{id}", gameHandler.GetGameByID)
	mux.HandleFunc("POST /games", middleware.JWTMiddleware(cfg.JWTSecret, gameHandler.CreateGame))
	mux.HandleFunc("PATCH /games/{id}", middleware.JWTMiddleware(cfg.JWTSecret, gameHandler.UpdateGame))
	mux.HandleFunc("DELETE /games/{id}", middleware.JWTMiddleware(cfg.JWTSecret, gameHandler.DeleteGame))

	// library
	mux.HandleFunc("GET /me/library", middleware.JWTMiddleware(cfg.JWTSecret, libraryHandler.GetLibrary))

	return mux
}
