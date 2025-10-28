package server

import (
	"net/http"

	"github.com/lucasmcclean/intro-to-go/middleware"
)

func New() *http.Server {
	mux := http.NewServeMux()

	addRoutes(mux)

	_ = middleware.CORS(mux)

	// TODO: Create the http server and return it

	return nil
}

