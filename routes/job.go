package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
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

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	jobs := make(chan int, 3)

	var wg sync.WaitGroup

	results := make([]string, 0)
	resultsMu := sync.Mutex{}

	worker := func(id int) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				resultsMu.Lock()
				results = append(results, fmt.Sprintf("worker %d: stopping", id))
				resultsMu.Unlock()
				return
			case job, ok := <-jobs:
				if !ok {
					return
				}
				resultsMu.Lock()
				results = append(results, fmt.Sprintf("worker %d: got job %d", id, job))
				resultsMu.Unlock()
				time.Sleep(300 * time.Millisecond)
				results = append(results, fmt.Sprintf("worker %d: completed job %d", id, job))
			}
		}
	}

	for i := 1; i <= req.Workers; i++ {
		wg.Add(1)
		go worker(i)
	}

	for j := 1; j <= req.Jobs; j++ {
		select {
		case jobs <- j:
			resultsMu.Lock()
			results = append(results, fmt.Sprintf("sent job %d", j))
			resultsMu.Unlock()
		case <-time.After(200 * time.Millisecond):
			resultsMu.Lock()
			results = append(results, fmt.Sprintf("queue full, dropping job %d", j))
			resultsMu.Unlock()
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