package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type JobRequest struct {
	Workers int `json:"workers"`
	Jobs    int `json:"jobs"`
}

func JobsHandler(w http.ResponseWriter, r *http.Request) {
	var req JobRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	// TODO: Create a context

	// TODO: Make a job queue
	var jobs chan int

	results := make([]string, 0)
	resultsMu := sync.Mutex{}

	// TODO: Create a worker function to handle jobs
	var wg sync.WaitGroup
	worker := func(_ int) {
		defer wg.Done()
	}

	for i := 1; i <= req.Workers; i++ {
		wg.Add(1)
		go worker(i)
	}

	// Add jobs to the queue if there's room
	for j := 1; j <= req.Jobs; j++ {
		select {
		case jobs <- j:
			resultsMu.Lock()
			results = append(results, fmt.Sprintf("sent job %d", j))
			resultsMu.Unlock()
			// TODO: Handle the case of no room
		}
	}

	close(jobs)
	wg.Wait()

	results = append(results, "all workers done")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(results); err != nil {
		http.Error(w, "error encoding JSON", http.StatusInternalServerError)
	}
}
