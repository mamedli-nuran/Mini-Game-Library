package main

import (
	"context"
	"log/slog"
	"mini-game-library/config"
	"mini-game-library/handler"
	"mini-game-library/repository"
	"mini-game-library/router"
	"mini-game-library/service"
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

	userRepo := repository.NewUserRepository(pool)

	userService := service.NewUserService(userRepo, cfg.JWTSecret, cfg.AccessTokenExpireSeconds, cfg.RefreshTokenExpireHours)

	userHandler := handler.NewUserHandler(userService)
	mux := router.Setup(userHandler)

	slog.Info("Server starting on :8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		slog.Info("Error starting server: ", err)
	}
}
