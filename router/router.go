package router

import (
	"mini-game-library/handler"
	"net/http"
)

func Setup(
	userHandler *handler.UserHandler,
) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /auth/register", userHandler.RegisterUser)
	return mux
}
