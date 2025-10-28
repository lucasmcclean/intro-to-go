package routes

import (
	"encoding/json"
	"net/http"
)

// TODO: Create the mixer interface

// TODO: Create the left, right, and weave mixer structs

// TODO: Add mix method for each struct

type MixRequest struct {
	Left  string `json:"left"`
	Right string `json:"right"`
	How   string `json:"how"`
}

func MixHandler(w http.ResponseWriter, r *http.Request) {
	var req MixRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	// var mixer Mixer

	// TODO: Switch on "how" and run the mixer

	// result := mixer.Mix()

	var response map[string]string

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "error encoding JSON", http.StatusInternalServerError)
	}
}
