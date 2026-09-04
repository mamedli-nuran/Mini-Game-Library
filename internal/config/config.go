package config

import (
	"log/slog"
	"mini-game-library/internal/constant"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL              string
	AccessTokenExpireSeconds time.Duration
	RefreshTokenExpireHours  time.Duration
	JWTSecret                string
}

func Load() Config {
	var config Config
	if err := godotenv.Load(); err != nil {
		panic(constant.ErrEnvLoading)
	}

	accessExpireStr := os.Getenv("ACCESS_TOKEN_EXPIRE_SECONDS")
	accessExpire, err := strconv.Atoi(accessExpireStr)
	if err != nil {
		panic(err.Error())
	}

	refreshExpireStr := os.Getenv("REFRESH_TOKEN_EXPIRE_HOURS")
	refreshExpire, err := strconv.Atoi(refreshExpireStr)
	if err != nil {
		panic(err.Error())
	}
	slog.Info("Successfully read data from .env file")

	config.DatabaseURL = os.Getenv("DATABASE_URL")
	config.AccessTokenExpireSeconds = time.Duration(accessExpire) * time.Second
	config.RefreshTokenExpireHours = time.Duration(refreshExpire) * time.Hour
	config.JWTSecret = os.Getenv("JWT_SECRET")

	return config
}
