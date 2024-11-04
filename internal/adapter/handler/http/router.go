package http

import (
	"net/http"
)

func NewRouter(userHandler *UserHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/users", userHandler.Register)
	mux.HandleFunc("POST /api/users/login", userHandler.Login)

	return mux
}
