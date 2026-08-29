package router

import (
	"mini-game-library/handler"
	"net/http"
)

func Setup(
	userHandler *handler.UserHandler,
) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test", userHandler.Test)

	return mux
}
