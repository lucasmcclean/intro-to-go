package routes

import (
	"net/http"
)

// TODO: Create a request struct
type DuplicateRequest struct {
	Entries []int `json:"entries"`
}

func DuplicateHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Decode the request
	var request DuplicateRequest

	// TODO: Create a response

	// TODO: Copy request entries into response
	for range request.Entries {
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// TODO: Encode the response
}
