package server

import (
	"net/http"
)

func New() *http.Server {
	mux := http.NewServeMux()

	addRoutes(mux)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return srv
}
