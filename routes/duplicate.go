package routes

import (
	"encoding/json"
	"net/http"
)

type DuplicateRequest struct {
	Entries []int `json:"entries"`
}

func DuplicateEntries(w http.ResponseWriter, r *http.Request) {
	var request DuplicateRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := struct {
		DuplicatedEntries []int `json:"entries"`
	}{
		DuplicatedEntries: make([]int, 0, len(request.Entries)*2),
	}

	for _, entry := range request.Entries {
		response.DuplicatedEntries = append(response.DuplicatedEntries, entry, entry)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
