package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// --- Go Syntax ---

// Variables

var number = 1

var t int = 0

// Constants

const chicken = "chicken"

const cool bool = true

// Types

type Example struct {
	Data []int
}

type Celsius int

// Functions

func IsWarm(temp Celsius) bool {
	return temp >= 25
}

// Methods

func add(ex *Example, amt int) {
	for i := range ex.Data {
		ex.Data[i] += amt
	}
	for _, e := range ex.Data {
		fmt.Println(e)
	}
}

func (ex *Example) add(amt int) {
	for i := range ex.Data {
		ex.Data[i] += amt
	}
	for _, e := range ex.Data {
		fmt.Println(e)
	}
}

// Interfaces

type Adder interface {
	Add(amt int)
}

// Built-in Data Types

// func main() {
// 	basket := make(map[string]int)
//
// 	basket["apple"] = 3
// 	basket["book"] = 2
//
// 	fmt.Println(basket["apple"], basket["book"])
// }

// Pointers

// func main() {
// 	x := 1
// 	y := &x
// 	z := *y
//
// 	fmt.Println(x, y, z)
// }

// Error Handling

func maybeFails(x int) (int, error) {
	if x > 5 {
		return 0, errors.New("uh oh, x is too big")
	} else {
		return x + 5, nil
	}
}

// func main() {
// 	val, err := maybeFails(6)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
//
// 	fmt.Println(val)
// }

// --- Concurrency Model ---

// Goroutines

// WaitGroup

// func main() {
// 	var wg sync.WaitGroup
//
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		fmt.Println("1")
// 	}()
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		fmt.Println("2")
// 	}()
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		fmt.Println("3")
// 	}()
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		fmt.Println("4")
// 	}()
// 	wg.Wait()
// }

// Channels

// func main() {
// 	ch := make(chan int)
//
// 	go func() {
// 		ch <- 10
// 		ch <- 10
// 		ch <- 10
// 		ch <- 10
// 		close(ch)
// 	}()
//
// 	for val := range ch {
// 		fmt.Println(val)
// 	}
// }

// func main() {
// 	ch := make(chan int, 2)
//
// 	ch <- 1
// 	ch <- 2
// 	// ch <- 3
//
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

// Select

// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)
//
// 	go func() { ch1 <- "from channel 1" }()
// 	go func() { time.Sleep(1 * time.Second); ch2 <- "from channel 2" }()
//
// 	select {
// 	case msg1 := <-ch1:
// 		fmt.Println("Received:", msg1)
// 	case msg2 := <-ch2:
// 		fmt.Println("Received:", msg2)
// 	}
// }

// func main() {
// 	ch := make(chan string, 1)
//
// 	// ch <- "testing"
//
// 	select {
// 	case msg := <-ch:
// 		fmt.Println("Got:", msg)
// 	case <-time.After(2 * time.Second):
// 		fmt.Println("Timeout!")
// 	}
// }

// Context

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())
//
// 	// var wg sync.WaitGroup
// 	// wg.Add(1)
// 	go func() {
// 		// defer wg.Done()
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				fmt.Println("Goroutine stopped:", ctx.Err())
// 				return
// 			default:
// 				fmt.Println("Working...")
// 				time.Sleep(500 * time.Millisecond)
// 			}
// 		}
// 	}()
//
// 	time.Sleep(2 * time.Second)
// 	cancel()
// 	// wg.Wait()
// }
