package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

// --- Go Standard Library ---

func main() {
	x := 0
	fmt.Println("Hello, World!")
	fmt.Sprintf("value: %d", x)

	// - strings
	strings.ToUpper("go")
	strings.Split("a,b,c", ",")

	// - time
	start := time.Now()
	time.Sleep(2 * time.Second)
	time.Since(start)

	// - os
	os.Getenv("HOME")
	os.Open("file.txt")
	os.Exit(1)

	// - encoding/json
	v := struct {
		Yup int `json:"yup"`
	}{
		Yup: 1,
	}
	json.Marshal(v)

	data := make([]byte, 10)
	json.Unmarshal(data, &v)

	// - net/http
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {})
	http.ListenAndServe(":8080", nil)
	http.Get("https://example.com")

	// - context
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	<-ctx.Done()

	// - sync
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() { defer wg.Done() }()

	// - log
	log.Println("info message")
	log.Printf("formatted: %d", 1)

	// - math
	math.Sqrt(16)
	rand.Intn(10)

	// - errors
	err := errors.New("something went wrong")
	target := fmt.Errorf("wrap: %w", err)
	errors.Is(err, target)
}
