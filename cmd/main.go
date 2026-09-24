package main

import (
	"context"
	"log/slog"
	"mini-game-library/internal/config"
	handler2 "mini-game-library/internal/handler"
	repository2 "mini-game-library/internal/repository"
	"mini-game-library/internal/router"
	service2 "mini-game-library/internal/service"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg := config.Load()
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		panic(err.Error())
	}
	defer pool.Close()

	err = pool.Ping(ctx)
	if err != nil {
		panic(err)
	}
	slog.Info("Successfully connected to postgres")

	userRepo := repository2.NewUserRepository(pool)
	gameRepo := repository2.NewGameRepository(pool)
	libraryRepo := repository2.NewLibraryRepository(pool)
	ratingRepo := repository2.NewRatingRepository(pool)

	userService := service2.NewUserService(userRepo, cfg)
	gameService := service2.NewGameService(gameRepo)
	libraryService := service2.NewLibraryService(libraryRepo)
	ratingService := service2.NewRatingService(ratingRepo)

	userHandler := handler2.NewUserHandler(userService, cfg)
	gameHandler := handler2.NewGameHandler(gameService, cfg)
	libraryHandler := handler2.NewLibraryHandler(libraryService)
	ratingHandler := handler2.NewRatingHandler(ratingService)

	mux := router.Setup(userHandler, gameHandler, libraryHandler, ratingHandler, cfg)

	slog.Info("Server starting on :8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		slog.Info("Server stopped", slog.Any("error", err))
	}
}
