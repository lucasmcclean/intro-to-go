package server

import (
	"net/http"

	"github.com/lucasmcclean/intro-to-go/routes"
)

func addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/duplicate", routes.DuplicateEntries)
}
