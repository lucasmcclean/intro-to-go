package main

import (
	"sync"

	"github.com/lucasmcclean/intro-to-go/server"
)

func main() {
	// ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	// defer cancel()

	_ = server.New()

	// TODO: Run the server in a goroutine
	go func() {
	}()

	var wg sync.WaitGroup

	// TODO: Run the shutdown in a goroutine
	go func() {
		defer wg.Done()
	}()

	wg.Wait()
}
