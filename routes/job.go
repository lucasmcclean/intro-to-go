package routes

import (
	"net/http"
)

// Create a request with "workers" and "jobs"

func JobsHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the request

	// Create a context with 2 second timeout

	// Create a jobs channel

	// Create a waitgroup

	// Create a results slice and mutex

	// Create a worker func that:
	// - Stops on context cancel
	// - Receives jobs

	// For number of workers, create a worker (add to wg)

	// For number of jobs
	// - Place a job in the queue
	// - Cancel after .2 seconds if queue full

	// Close jobs and wait for workers to finish

	// Append complete message

	// Set headers

	// Encode the response
}
