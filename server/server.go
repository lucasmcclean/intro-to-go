package server

import (
	"net/http"

	"github.com/lucasmcclean/intro-to-go/middleware"
)

func New() *http.Server {
	mux := http.NewServeMux()

	addRoutes(mux)

	handler := middleware.CORS(mux)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	return srv
}