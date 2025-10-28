package routes

import (
	"net/http"
)

// TODO: Create a request struct

func DuplicateHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Decode the request

	// TODO: Create a response

	// TODO: Copy request entries into response
	// for _, entry := range request.Entries {
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// TODO: Encode the response
}
