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

	userService := service2.NewUserService(userRepo, cfg)
	gameService := service2.NewGameService(gameRepo)

	userHandler := handler2.NewUserHandler(userService)
	gameHandler := handler2.NewGameHandler(gameService)

	mux := router.Setup(userHandler, gameHandler, cfg)

	slog.Info("Server starting on :8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		slog.Info("Error starting server: ", err)
	}
}
